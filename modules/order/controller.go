package order

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/modules/booking_flow"
	"github.com/eden-w2w/lib-modules/modules/discounts"
	"github.com/eden-w2w/lib-modules/modules/freight_trial"
	"github.com/eden-w2w/lib-modules/modules/goods"
	"github.com/eden-w2w/lib-modules/modules/id_generator"
	"github.com/eden-w2w/lib-modules/modules/payment_flow"
	"github.com/eden-w2w/lib-modules/modules/user"
	"github.com/eden-w2w/lib-modules/pkg/search"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/databases"
)

var controller *Controller

func GetController() *Controller {
	if controller == nil {
		controller = &Controller{isInit: false}
	}
	return controller
}

type Controller struct {
	isInit              bool
	db                  sqlx.DBExecutor
	orderExpireDuration time.Duration
	eventHandler        EventHandler
}

func (c *Controller) Init(db sqlx.DBExecutor, d time.Duration, h EventHandler) {
	c.db = db
	c.orderExpireDuration = d
	c.eventHandler = h
	c.isInit = true
}

func (c Controller) PreCreateOrder(p PreCreateOrderParams) (
	preGoodsList []PreCreateOrderGoodsParams,
	totalPrice, freightPrice, discountPrice, actualPrice uint64,
	freightName string,
	err error,
) {
	if !c.isInit {
		logrus.Panicf("[OrderController] not Init")
	}
	if p.UserID == 0 {
		logrus.Error("[CreateOrder] userID cannot be empty")
		return nil, 0, 0, 0, 0, "", general_errors.BadRequest
	}
	// 获取用户信息
	_, err = user.GetController().GetUserByUserID(p.UserID, nil, false)
	if err != nil {
		return
	}

	// 获取订单总额与库中物料进行比对，物料库存检查及优惠计算
	var goodsList = make([]freight_trial.FreightTrialParams, 0)
	for _, g := range p.Goods {
		gModel, err := goods.GetController().GetGoodsByID(g.GoodsID, nil, false)
		if err != nil {
			return nil, 0, 0, 0, 0, "", err
		}

		var isBooking = datatypes.BOOL_FALSE
		var bookingFlowID = uint64(0)
		flows, err := booking_flow.GetController().GetBookingFlowByGoodsID(gModel.GoodsID)
		if err != nil {
			return nil, 0, 0, 0, 0, "", err
		}
		if len(flows) > 0 {
			isBooking = datatypes.BOOL_TRUE
			bookingFlowID = flows[0].FlowID
		}
		if isBooking != datatypes.BOOL_TRUE && uint64(g.Amount) > gModel.Inventory {
			return nil, 0, 0, 0, 0, "", general_errors.GoodsInventoryShortage
		}
		if *g.IsBooking && isBooking == datatypes.BOOL_FALSE {
			return nil, 0, 0, 0, 0, "", general_errors.GoodsInventorySufficient
		}
		if !*g.IsBooking && isBooking == datatypes.BOOL_TRUE {
			return nil, 0, 0, 0, 0, "", general_errors.GoodsInventoryShortage
		}

		totalPrice += gModel.Price * uint64(g.Amount)
		goodsList = append(
			goodsList, freight_trial.FreightTrialParams{
				Goods:         *gModel,
				Amount:        g.Amount,
				IsBooking:     isBooking,
				BookingFlowID: bookingFlowID,
			},
		)
	}
	if len(goodsList) == 0 {
		logrus.Errorf("[CreateOrder] len(goodsList) == 0")
		return nil, 0, 0, 0, 0, "", general_errors.BadRequest.StatusError().WithDesc("商品列表为空")
	}

	// 试算运费
	if p.ShippingID != 0 {
		shipping, err := user.GetController().GetShippingAddressByShippingID(p.ShippingID, p.UserID)
		if err != nil {
			return nil, 0, 0, 0, 0, "", err
		}
		freightPrice, freightName, err = freight_trial.FreightTrial(goodsList, shipping)
		if err != nil {
			return nil, 0, 0, 0, 0, "", err
		}
	}

	// 计算优惠，目前暂时只支持同时进行一种优惠
	if len(p.Discounts) > 0 {
		discountID := p.Discounts[0]
		err := func() error {
			_ = discounts.GetController().RLock(discountID)
			defer discounts.GetController().RUnlock(discountID)
			discount, err := discounts.GetController().GetDiscountByID(discountID, nil, false)
			if err != nil {
				return err
			}
			if discount.Status != enums.DISCOUNT_STATUS__PROCESS {
				return general_errors.DiscountNotStart
			}
			if discount.Limit > 0 && discount.Times >= uint64(discount.Limit) {
				return general_errors.DiscountEnd
			}
			current := time.Now()
			if current.Before(time.Time(discount.ValidityStart)) {
				return general_errors.DiscountNotStart
			} else if current.After(time.Time(discount.ValidityEnd)) {
				return general_errors.DiscountEnd
			}

			var discountAmount uint64
			preGoodsList, _, discountAmount = ToDiscountAmount(discount, goodsList)
			discountPrice += discountAmount
			return nil
		}()
		if err != nil {
			return nil, 0, 0, 0, 0, "", err
		}
	}
	actualPrice = totalPrice + freightPrice - discountPrice
	return
}

func (c Controller) CreateOrder(p CreateOrderParams) (*databases.Order, error) {
	if !c.isInit {
		logrus.Panicf("[OrderController] not Init")
	}
	if p.UserID == 0 {
		logrus.Error("[CreateOrder] userID cannot be empty")
		return nil, general_errors.BadRequest
	}
	// 获取用户信息
	u, err := user.GetController().GetUserByUserID(p.UserID, nil, false)
	if err != nil {
		return nil, err
	}

	var order *databases.Order
	var orderGoodsList = make([]databases.OrderGoods, 0)
	var goodsList = make([]freight_trial.FreightTrialParams, 0)

	tx := sqlx.NewTasks(c.db)

	// 获取订单总额与库中物料进行比对，物料库存检查及优惠计算
	var discountAmountTotal, freightAmountTotal, actualAmountTotal uint64
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			var totalPrice uint64 = 0
			for _, g := range p.Goods {
				gModel, err := goods.GetController().GetGoodsByID(g.GoodsID, db, true)
				if err != nil {
					return err
				}

				var isBooking = datatypes.BOOL_FALSE
				var bookingFlowID = uint64(0)
				flows, err := booking_flow.GetController().GetBookingFlowByGoodsID(gModel.GoodsID)
				if err != nil {
					return err
				}
				if len(flows) > 0 {
					isBooking = datatypes.BOOL_TRUE
					bookingFlowID = flows[0].FlowID
				}
				if isBooking != datatypes.BOOL_TRUE && uint64(g.Amount) > gModel.Inventory {
					return general_errors.GoodsInventoryShortage
				}
				if *g.IsBooking && isBooking == datatypes.BOOL_FALSE {
					return general_errors.GoodsInventorySufficient
				}
				if !*g.IsBooking && isBooking == datatypes.BOOL_TRUE {
					return general_errors.GoodsInventoryShortage
				}

				totalPrice += gModel.Price * uint64(g.Amount)
				goodsList = append(
					goodsList, freight_trial.FreightTrialParams{
						Goods:         *gModel,
						Amount:        g.Amount,
						IsBooking:     isBooking,
						BookingFlowID: bookingFlowID,
					},
				)
			}

			if len(goodsList) == 0 {
				logrus.Errorf("[CreateOrder] len(goodsList) == 0")
				return general_errors.BadRequest.StatusError().WithDesc("商品列表为空")
			}

			if totalPrice != p.TotalPrice {
				logrus.Errorf(
					"[CreateOrder] totalPrice != p.TotalPrice totalPrice: %d, p.TotalPrice: %d",
					totalPrice,
					p.TotalPrice,
				)
				return general_errors.BadRequest.StatusError().WithoutErrTalk().WithMsg("订单总额与商品总价不一致，请尝试返回重新下单")
			}

			shipping, err := user.GetController().GetShippingAddressByShippingID(p.ShippingID, p.UserID)
			if err != nil {
				return err
			}
			freightAmountTotal, _, err = freight_trial.FreightTrial(goodsList, shipping)
			if err != nil {
				return err
			}
			if freightAmountTotal != p.FreightAmount {
				logrus.Errorf(
					"[CreateOrder] freightAmountTotal != p.FreightAmount freightAmountTotal: %d, p.FreightAmount: %d",
					freightAmountTotal,
					p.FreightAmount,
				)
				return general_errors.BadRequest.StatusError().WithoutErrTalk().WithMsg("订单运费核验失败，请尝试返回重新下单")
			}

			// 计算优惠，目前暂时只支持同时进行一种优惠
			if len(p.Discounts) > 0 {
				discountID := p.Discounts[0]
				err := func() error {
					_ = discounts.GetController().RLock(discountID)
					defer discounts.GetController().RUnlock(discountID)
					discount, err := discounts.GetController().GetDiscountByID(discountID, db, true)
					if err != nil {
						return err
					}
					if discount.Status != enums.DISCOUNT_STATUS__PROCESS {
						return general_errors.DiscountNotStart
					}
					if discount.Limit > 0 && discount.Times >= uint64(discount.Limit) {
						return general_errors.DiscountEnd
					}
					current := time.Now()
					if current.Before(time.Time(discount.ValidityStart)) {
						return general_errors.DiscountNotStart
					} else if current.After(time.Time(discount.ValidityEnd)) {
						return general_errors.DiscountEnd
					}
					_, _, discountAmount := ToDiscountAmount(discount, goodsList)
					discountAmountTotal += discountAmount
					return nil
				}()
				if err != nil {
					return err
				}
			}

			// 计算实际订单金额
			actualAmountTotal = totalPrice + freightAmountTotal - discountAmountTotal
			return nil
		},
	)

	// 创建订单
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			id := id_generator.GetGenerator().GenerateUniqueID()
			order = &databases.Order{
				OrderID:        id,
				UserID:         p.UserID,
				NickName:       u.NickName,
				UserOpenID:     u.OpenID,
				TotalPrice:     p.TotalPrice,
				DiscountAmount: discountAmountTotal,
				FreightAmount:  p.FreightAmount,
				ActualAmount:   actualAmountTotal,
				PaymentMethod:  p.PaymentMethod,
				Remark:         p.Remark,
				Status:         enums.ORDER_STATUS__CREATED,
				ExpiredAt:      datatypes.MySQLTimestamp(time.Now().Add(c.orderExpireDuration)),
			}
			err := order.Create(db)
			if err != nil {
				return err
			}

			id = id_generator.GetGenerator().GenerateUniqueID()
			courier := &databases.OrderLogistics{
				PrimaryID:      datatypes.PrimaryID{},
				LogisticsID:    id,
				OrderID:        order.OrderID,
				ShippingID:     p.ShippingID,
				Recipients:     p.Recipients,
				ShippingAddr:   p.ShippingAddr,
				Mobile:         p.Mobile,
				CourierCompany: "",
				CourierNumber:  "",
			}
			return courier.Create(db)
		},
	)

	// 创建订单物料
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			for _, item := range goodsList {
				orderGoods := databases.OrderGoods{
					OrderID:           order.OrderID,
					GoodsID:           item.GoodsID,
					Name:              item.Name,
					Comment:           item.Comment,
					FreightTemplateID: item.FreightTemplateID,
					Sales:             item.Sales,
					MainPicture:       item.MainPicture,
					Pictures:          item.Pictures,
					Specifications:    item.Specifications,
					Price:             item.Price,
					Inventory:         item.Inventory,
					Detail:            item.Detail,
					Amount:            item.Amount,
					IsBooking:         item.IsBooking,
				}
				err := orderGoods.Create(db)
				if err != nil {
					return err
				}

				orderGoodsList = append(orderGoodsList, orderGoods)
			}
			return nil
		},
	)

	// 执行创建事件
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			return c.eventHandler.OnOrderCreateEvent(db, order, orderGoodsList)
		},
	)

	err = tx.Do()
	if err != nil {
		logrus.Errorf("[CreateOrder] err: %v, params: %+v", err, p)
		return nil, general_errors.InternalError
	}

	return order, nil
}

func (c Controller) GetOrder(orderID, userID uint64, db sqlx.DBExecutor, forUpdate bool) (
	order *databases.Order,
	logistics *databases.OrderLogistics,
	err error,
) {
	if !c.isInit {
		logrus.Panicf("[OrderController] not Init")
	}
	if db == nil {
		db = c.db
	}
	order = &databases.Order{OrderID: orderID}
	if forUpdate {
		err = order.FetchByOrderIDForUpdate(db)
	} else {
		err = order.FetchByOrderID(db)
	}
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, nil, general_errors.OrderNotFound
		}
		logrus.Errorf("[GetOrder] err: %v, orderID: %d", err, orderID)
		return nil, nil, general_errors.InternalError
	}
	if userID != 0 && order.UserID != userID {
		logrus.Errorf("[GetOrder] order.UserID != userID, order.UserID: %d, userID: %d", order.UserID, userID)
		return nil, nil, general_errors.Forbidden
	}

	logistics = &databases.OrderLogistics{
		OrderID: orderID,
	}
	err = logistics.FetchByOrderID(db)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, nil, general_errors.OrderNotFound
		}
		logrus.Errorf("[GetOrder] logistics.FetchByOrderID err: %v, orderID: %d", err, orderID)
		return nil, nil, general_errors.InternalError
	}
	return order, logistics, nil
}

func (c Controller) GetOrders(p GetOrdersParams, withCount bool) (orders []databases.Order, count int, err error) {
	if !c.isInit {
		logrus.Panicf("[OrderController] not Init")
	}
	order := databases.Order{}
	orders, err = order.List(c.db, p.Conditions(), p.Additions()...)
	if err != nil {
		logrus.Errorf("[GetOrders] order.List err: %v, params: %+v", err, p)
		return nil, 0, general_errors.InternalError
	}
	if withCount {
		count, err = order.Count(c.db, p.Conditions())
		if err != nil {
			logrus.Errorf("[GetOrders] order.Count err: %v, params: %+v", err, p)
			return nil, 0, general_errors.InternalError
		}
	}
	return
}

func (c Controller) GetOrderLogistics(orderID uint64) (*databases.OrderLogistics, error) {
	logistics := &databases.OrderLogistics{
		OrderID: orderID,
	}
	err := logistics.FetchByOrderID(c.db)
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, general_errors.OrderNotFound
		}
		logrus.Errorf("[GetOrderLogistics] logistics.FetchByOrderID err: %v, orderID: %d", err, orderID)
		return nil, general_errors.InternalError
	}
	return logistics, nil
}

func (c Controller) GetOrderGoods(orderID uint64, db sqlx.DBExecutor) ([]databases.OrderGoods, error) {
	if !c.isInit {
		logrus.Panicf("[OrderController] not Init")
	}
	if db == nil {
		db = c.db
	}
	og := databases.OrderGoods{}
	goods, err := og.BatchFetchByOrderIDList(db, []uint64{orderID})
	if err != nil {
		logrus.Errorf("[GetOrderGoods] og.BatchFetchByOrderIDList err: %v, orderID: %d", err, orderID)
		return nil, general_errors.InternalError
	}
	return goods, nil
}

func (c Controller) updateOrderStatus(db sqlx.DBExecutor, order *databases.Order, status enums.OrderStatus) error {
	if order.Status == status {
		return nil
	}

	// 状态流转检查
	if !order.Status.CheckNextStatusIsValid(status) {
		logrus.Errorf(
			"[updateOrderStatus] !order.Status.CheckNextStatusIsValid(status) currentStatus: %s, nextStatus: %s",
			order.Status.String(),
			status.String(),
		)
		return general_errors.OrderStatusFlowIncorrect
	}

	// 变更订单状态
	f := builder.FieldValues{
		"Status": status,
	}
	order.Status = status
	err := order.UpdateByIDWithMap(db, f)
	if err != nil {
		logrus.Errorf(
			"[updateOrderStatus] order.UpdateByIDWithMap err: %v, orderID: %d, status: %s",
			err,
			order.OrderID,
			status.String(),
		)
		return general_errors.InternalError
	}

	return nil
}

func (c Controller) updateOrderDiscount(db sqlx.DBExecutor, order *databases.Order, params UpdateOrderParams) error {
	if order.DiscountAmount == params.DiscountAmount {
		return nil
	}
	if order.Status != enums.ORDER_STATUS__CREATED {
		if params.Status != enums.ORDER_STATUS__PAID {
			return general_errors.NotAllowedChangeAmount
		}
	}
	if params.DiscountAmount > order.TotalPrice {
		return general_errors.DiscountAmountOverflow
	}

	order.DiscountAmount = params.DiscountAmount
	order.ActualAmount = order.TotalPrice - params.DiscountAmount
	f := builder.FieldValues{
		"DiscountAmount": params.DiscountAmount,
		"ActualAmount":   order.ActualAmount,
	}
	err := order.UpdateByIDWithMap(db, f)
	if err != nil {
		logrus.Errorf(
			"[updateOrderDiscount] order.UpdateByIDWithMap err: %v, orderID: %d, discount: %d",
			err,
			order.OrderID,
			params.DiscountAmount,
		)
		return general_errors.InternalError
	}
	return nil
}

func (c Controller) updateOrderLogistics(
	db sqlx.DBExecutor,
	order *databases.Order,
	logistics *databases.OrderLogistics,
	shippingID uint64, recipients, address, mobile string,
) (err error) {
	if logistics.ShippingID == shippingID && logistics.Recipients == recipients && logistics.ShippingAddr == address && logistics.Mobile == mobile {
		return
	}
	if order.Status >= enums.ORDER_STATUS__DISPATCH {
		return general_errors.NotAllowedChangeLogistics
	}
	logistics.ShippingID = shippingID
	logistics.Recipients = recipients
	logistics.ShippingAddr = address
	logistics.Mobile = mobile
	f := builder.FieldValues{
		"ShippingID":   logistics.ShippingID,
		"Recipients":   logistics.Recipients,
		"ShippingAddr": logistics.ShippingAddr,
		"Mobile":       logistics.Mobile,
	}
	err = logistics.UpdateByIDWithMap(db, f)
	if err != nil {
		logrus.Errorf(
			"[updateOrderLogistics] logistics.UpdateByIDWithMap err: %v, orderID: %d, logisticsID: %d, recipients: %s, address: %s, mobile: %s",
			err,
			order.OrderID,
			logistics.LogisticsID,
			recipients,
			address,
			mobile,
		)
		return general_errors.InternalError
	}
	return
}

func (c Controller) updateCourierInfo(
	db sqlx.DBExecutor,
	logistics *databases.OrderLogistics,
	courierCompany, courierNumber string,
) error {
	if logistics.CourierCompany == courierCompany && logistics.CourierNumber == courierNumber {
		return nil
	}
	logistics.CourierCompany = courierCompany
	logistics.CourierNumber = courierNumber
	f := builder.FieldValues{
		"CourierCompany": logistics.CourierCompany,
		"CourierNumber":  logistics.CourierNumber,
	}
	err := logistics.UpdateByIDWithMap(db, f)
	if err != nil {
		logrus.Errorf(
			"[updateCourierInfo] logistics.UpdateByIDWithMap err: %v, logisticsID: %d, courierCompany: %s, courierNumber: %s",
			err,
			logistics.LogisticsID,
			courierCompany,
			courierNumber,
		)
		return general_errors.InternalError
	}
	return nil
}

func (c Controller) updateOrderRemark(db sqlx.DBExecutor, order *databases.Order, remark string) error {
	if order.Remark == remark {
		return nil
	}
	order.Remark = remark
	f := builder.FieldValues{
		"Remark": remark,
	}
	err := order.UpdateByIDWithMap(db, f)
	if err != nil {
		logrus.Errorf(
			"[updateOrderRemark] order.UpdateByIDWithMap err: %v, orderID: %d, remark: %s",
			err,
			order.OrderID,
			remark,
		)
		return general_errors.InternalError
	}
	return nil
}

func (c Controller) updateOrderGoods(
	db sqlx.DBExecutor,
	order *databases.Order,
	goods []databases.OrderGoods,
	params []CreateOrderGoodsParams,
	locker InventoryLock,
	unlocker InventoryUnlock,
) error {
	if order.Status != enums.ORDER_STATUS__CREATED {
		return general_errors.NotAllowedChangeAmount
	}

	type modifiedGoodsParams struct {
		databases.OrderGoods
		modifiedAmount int32
	}

	var deleteGoods = make([]databases.OrderGoods, 0)
	var newGoods = make([]CreateOrderGoodsParams, 0)
	var modifiedGoods = make([]modifiedGoodsParams, 0)

	for _, g := range goods {
		ok, i, err := search.In(
			params, g.GoodsID, func(current interface{}, needle interface{}) bool {
				var p = current.(CreateOrderGoodsParams)
				if p.GoodsID == needle {
					return true
				}
				return false
			},
		)
		if err != nil {
			logrus.Errorf("[updateOrderGoods] search.In params err: %v", err)
			return err
		}
		if ok {
			if g.Amount != params[i].Amount {
				offset := int32(params[i].Amount) - int32(g.Amount)
				g.Amount = params[i].Amount
				modifiedGoods = append(
					modifiedGoods, modifiedGoodsParams{
						g,
						offset,
					},
				)
			}
		} else {
			deleteGoods = append(deleteGoods, g)
		}
	}

	for _, param := range params {
		ok, _, err := search.In(
			goods, param.GoodsID, func(current interface{}, needle interface{}) bool {
				var g = current.(databases.OrderGoods)
				if g.GoodsID == needle {
					return true
				}
				return false
			},
		)
		if err != nil {
			logrus.Errorf("[updateOrderGoods] search.In goods err: %v", err)
			return err
		}
		if !ok {
			newGoods = append(newGoods, param)
		}
	}

	for _, g := range modifiedGoods {
		err := g.OrderGoods.UpdateByOrderIDAndGoodsIDWithStruct(db)
		if err != nil {
			logrus.Errorf("[updateOrderGoods] g.UpdateByOrderIDAndGoodsIDWithStruct err: %v", err)
			return err
		}
		if g.IsBooking != datatypes.BOOL_TRUE {
			if g.modifiedAmount < 0 {
				err = unlocker(db, g.GoodsID, uint32(-g.modifiedAmount))
				if err != nil {
					logrus.Errorf(
						"[updateOrderGoods] unlocker err: %v, goodsID: %d, unlockAmount: %d",
						err,
						g.GoodsID,
						-g.modifiedAmount,
					)
					return err
				}
			} else {
				err = locker(db, g.GoodsID, uint32(g.modifiedAmount))
				if err != nil {
					logrus.Errorf(
						"[updateOrderGoods] locker err: %v, goodsID: %d, lockAmount: %d",
						err,
						g.GoodsID,
						g.modifiedAmount,
					)
					return err
				}
			}
		}
	}
	for _, g := range deleteGoods {
		if g.IsBooking != datatypes.BOOL_TRUE {
			// 释放库存
			err := unlocker(db, g.GoodsID, g.Amount)
			if err != nil {
				logrus.Errorf("[updateOrderGoods] unlocker err: %v", err)
				return err
			}
		}

		err := g.DeleteByOrderIDAndGoodsID(db)
		if err != nil {
			logrus.Errorf("[updateOrderGoods] g.DeleteByOrderIDAndGoodsID err: %v", err)
			return err
		}
	}
	for _, g := range newGoods {
		model := databases.Goods{GoodsID: g.GoodsID}
		err := model.FetchByGoodsIDForUpdate(db)
		if err != nil {
			logrus.Errorf("[updateOrderGoods] model.FetchByGoodsID(db) err: %v, goodsID: %d", err, g.GoodsID)
			return general_errors.GoodsNotFound
		}

		if model.Inventory > 0 && uint64(g.Amount) > model.Inventory {
			return general_errors.GoodsInventoryShortage
		}

		var isBooking = datatypes.BOOL_FALSE
		if model.Inventory == 0 {
			if model.IsAllowBooking == datatypes.BOOL_TRUE {
				isBooking = datatypes.BOOL_TRUE
			} else {
				return general_errors.GoodsInventoryShortage
			}
		}

		if isBooking != datatypes.BOOL_TRUE {
			// 锁定库存
			err = locker(db, g.GoodsID, g.Amount)
			if err != nil {
				logrus.Errorf("[updateOrderGoods] locker err: %v, goodsID: %d", err, g.GoodsID)
				return err
			}
		}

		// 创建物料
		orderGoods := &databases.OrderGoods{
			OrderID:           order.OrderID,
			GoodsID:           model.GoodsID,
			Name:              model.Name,
			Comment:           model.Comment,
			FreightTemplateID: model.FreightTemplateID,
			Sales:             model.Sales,
			MainPicture:       model.MainPicture,
			Pictures:          model.Pictures,
			Specifications:    model.Specifications,
			Price:             model.Price,
			Inventory:         model.Inventory,
			Detail:            model.Detail,
			Amount:            g.Amount,
			IsBooking:         isBooking,
		}
		err = orderGoods.Create(db)
		if err != nil {
			logrus.Errorf("[updateOrderGoods] orderGoods.Create err: %v, orderGoods: %+v", err, orderGoods)
			return err
		}
	}

	// 重新计算订单价格
	goodsList, err := c.GetOrderGoods(order.OrderID, db)
	if err != nil {
		logrus.Errorf("[updateOrderGoods] GetOrderGoods err: %v, orderID: %d", err, order.OrderID)
		return err
	}

	var totalPrice uint64 = 0
	for _, g := range goodsList {
		totalPrice += g.Price * uint64(g.Amount)
	}
	order.TotalPrice = totalPrice
	order.ActualAmount = order.TotalPrice - order.DiscountAmount
	f := builder.FieldValues{
		"TotalPrice":   totalPrice,
		"ActualAmount": order.ActualAmount,
	}
	err = order.UpdateByIDWithMap(db, f)
	if err != nil {
		logrus.Errorf(
			"[updateOrderGoods] order.UpdateByIDWithMap err: %v, orderID: %d, fields: %+v",
			err,
			order.OrderID,
			f,
		)
		return general_errors.InternalError
	}
	return nil
}

func (c Controller) UpdateOrder(
	order *databases.Order,
	logistics *databases.OrderLogistics,
	orderGoods []databases.OrderGoods,
	params UpdateOrderParams,
	locker InventoryLock,
	unlocker InventoryUnlock,
	db sqlx.DBExecutor,
) (err error) {
	if !c.isInit {
		logrus.Panicf("[OrderController] not Init")
	}
	if db == nil {
		db = c.db
	}

	if params.DiscountAmount != 0 {
		if err = c.updateOrderDiscount(db, order, params); err != nil {
			return err
		}
	}

	if params.CourierCompany != "" || params.CourierNumber != "" {
		if err = c.updateCourierInfo(db, logistics, params.CourierCompany, params.CourierNumber); err != nil {
			return err
		}
	}

	if params.Remark != "" {
		if err = c.updateOrderRemark(db, order, params.Remark); err != nil {
			return err
		}
	}

	if len(params.Goods) > 0 {
		if err = c.updateOrderGoods(db, order, orderGoods, params.Goods, locker, unlocker); err != nil {
			return err
		}
	}

	if params.ShippingID != 0 || params.Recipients != "" || params.ShippingAddr != "" || params.Mobile != "" {
		if err = c.updateOrderLogistics(
			db,
			order,
			logistics,
			params.ShippingID,
			params.Recipients,
			params.ShippingAddr,
			params.Mobile,
		); err != nil {
			return err
		}
	}

	orderStatus := order.Status
	if params.Status != enums.ORDER_STATUS_UNKNOWN && params.Status != enums.ORDER_STATUS__CLOSED {
		if err = c.updateOrderStatus(db, order, params.Status); err != nil {
			return err
		}
	}

	if orderStatus != order.Status {
		// 状态发生变更，执行状态变更事件
		switch order.Status {
		case enums.ORDER_STATUS__CONFIRM:
			err = c.eventHandler.OnOrderConfirmEvent(db, order)
		case enums.ORDER_STATUS__PAID:
			// 获取支付流水
			flows, err := payment_flow.GetController().MustGetFlowByOrderIDAndStatus(
				order.OrderID,
				order.UserID,
				[]enums.PaymentStatus{enums.PAYMENT_STATUS__SUCCESS},
				db,
			)
			if err != nil {
				return err
			}
			flow := flows[0]
			err = c.eventHandler.OnOrderPaidEvent(db, order, &flow)
			if err != nil {
				return err
			}
		case enums.ORDER_STATUS__DISPATCH:
			// 获取商品列表
			goodsList, err := c.GetOrderGoods(order.OrderID, db)
			if err != nil {
				return err
			}
			err = c.eventHandler.OnOrderDispatchEvent(db, order, logistics, goodsList)
			if err != nil {
				return err
			}
		case enums.ORDER_STATUS__COMPLETE:
			// 获取商品列表
			goodsList, err := c.GetOrderGoods(order.OrderID, db)
			if err != nil {
				return err
			}
			err = c.eventHandler.OnOrderCompleteEvent(db, order, logistics, goodsList)
			if err != nil {
				return err
			}
		}
	}

	return err
}

func (c Controller) CancelOrder(orderID, userID uint64) error {
	if !c.isInit {
		logrus.Panicf("[OrderController] not Init")
	}
	var order *databases.Order
	var err error
	tx := sqlx.NewTasks(c.db)
	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			order, _, err = c.GetOrder(orderID, userID, db, true)
			if err != nil {
				return err
			}

			if order.Status == enums.ORDER_STATUS__CLOSED {
				return general_errors.OrderCanceled
			}
			return nil
		},
	)

	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			return c.updateOrderStatus(db, order, enums.ORDER_STATUS__CLOSED)
		},
	)

	tx = tx.With(
		func(db sqlx.DBExecutor) error {
			goods, err := c.GetOrderGoods(orderID, db)
			if err != nil {
				return err
			}
			// 执行订单取消事件
			return c.eventHandler.OnOrderCloseEvent(db, order, goods)
		},
	)

	err = tx.Do()
	if err != nil {
		logrus.Errorf("[CancelOrder] tx.Do() err: %v, orderID: %d, userID: %d", err, orderID, userID)
	}
	return err
}
