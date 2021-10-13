package order

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/modules/id_generator"
	"github.com/eden-w2w/lib-modules/modules/payment_flow"
	"github.com/eden-w2w/lib-modules/modules/user"
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

func (c Controller) CreateOrder(p CreateOrderParams, locker InventoryLock) (*databases.Order, error) {
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

	// 获取订单总额与库中物料进行比对
	var totalPrice uint64 = 0
	var goodsList = make([]CreateOrderGoodsModelParams, 0)
	for _, g := range p.Goods {
		goods := databases.Goods{GoodsID: g.GoodsID}
		err := goods.FetchByGoodsID(c.db)
		if err != nil {
			logrus.Errorf("[CreateOrder] goods.FetchByGoodsID(c.db) err: %v, goodsID: %d", err, g.GoodsID)
			return nil, general_errors.NotFound.StatusError().WithDesc("商品无法找到")
		}
		totalPrice += goods.Price * uint64(g.Amount)
		goodsList = append(goodsList, CreateOrderGoodsModelParams{
			Goods:  goods,
			Amount: g.Amount,
		})
	}
	if totalPrice != p.TotalPrice {
		logrus.Errorf("[CreateOrder] totalPrice != p.TotalPrice totalPrice: %d, p.TotalPrice: %d", totalPrice, p.TotalPrice)
		return nil, general_errors.BadRequest.StatusError().WithDesc("订单总额与商品总价不一致")
	}
	if len(goodsList) == 0 {
		logrus.Errorf("[CreateOrder] len(goodsList) == 0")
		return nil, general_errors.BadRequest.StatusError().WithDesc("商品列表为空")
	}
	if p.TotalPrice-p.DiscountAmount != p.ActualAmount {
		logrus.Errorf("[CreateOrder] p.TotalPrice - p.DiscountAmount != p.ActualAmount p.TotalPrice: %d, p.DiscountAmount: %d, p.ActualAmount: %d", p.TotalPrice, p.DiscountAmount, p.ActualAmount)
		return nil, general_errors.BadRequest.StatusError().WithDesc("订单实际支付金额错误")
	}

	// 创建订单
	var order *databases.Order

	tx := sqlx.NewTasks(c.db)
	tx = tx.With(func(db sqlx.DBExecutor) error {
		id, _ := id_generator.GetGenerator().GenerateUniqueID()
		order = &databases.Order{
			OrderID:        id,
			UserID:         p.UserID,
			NickName:       u.NickName,
			UserOpenID:     u.OpenID,
			TotalPrice:     p.TotalPrice,
			DiscountAmount: p.DiscountAmount,
			ActualAmount:   p.ActualAmount,
			PaymentMethod:  p.PaymentMethod,
			Remark:         p.Remark,
			Status:         enums.ORDER_STATUS__CREATED,
			ExpiredAt:      datatypes.MySQLTimestamp(time.Now().Add(c.orderExpireDuration)),
		}
		err := order.Create(db)
		if err != nil {
			return err
		}

		id, _ = id_generator.GetGenerator().GenerateUniqueID()
		courier := &databases.OrderLogistics{
			PrimaryID:      datatypes.PrimaryID{},
			LogisticsID:    id,
			OrderID:        order.OrderID,
			Recipients:     p.Recipients,
			ShippingAddr:   p.ShippingAddr,
			Mobile:         p.Mobile,
			CourierCompany: "",
			CourierNumber:  "",
		}
		return courier.Create(db)
	})

	// 锁定库存
	tx = tx.With(func(db sqlx.DBExecutor) error {
		for _, item := range goodsList {
			err := locker(db, item.GoodsID, item.Amount)
			if err != nil {
				return err
			}
		}
		return nil
	})

	// 创建订单物料
	tx = tx.With(func(db sqlx.DBExecutor) error {
		for _, item := range goodsList {
			orderGoods := &databases.OrderGoods{
				OrderID:        order.OrderID,
				GoodsID:        item.GoodsID,
				Name:           item.Name,
				Comment:        item.Comment,
				DispatchAddr:   item.DispatchAddr,
				Sales:          item.Sales,
				MainPicture:    item.MainPicture,
				Pictures:       item.Pictures,
				Specifications: item.Specifications,
				Activities:     item.Activities,
				LogisticPolicy: item.LogisticPolicy,
				Price:          item.Price,
				Inventory:      item.Inventory,
				Detail:         item.Detail,
				Amount:         item.Amount,
			}
			err := orderGoods.Create(db)
			if err != nil {
				return err
			}
		}
		return nil
	})

	// 执行创建事件
	tx = tx.With(func(db sqlx.DBExecutor) error {
		return c.eventHandler.OnOrderCreateEvent(db, order)
	})

	err = tx.Do()
	if err != nil {
		logrus.Errorf("[CreateOrder] err: %v, params: %+v", err, p)
		return nil, general_errors.InternalError
	}

	return order, nil
}

func (c Controller) GetOrder(orderID, userID uint64, db sqlx.DBExecutor, forUpdate bool) (order *databases.Order, logistics *databases.OrderLogistics, err error) {
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

func (c Controller) GetOrderGoods(orderID uint64) ([]databases.OrderGoods, error) {
	if !c.isInit {
		logrus.Panicf("[OrderController] not Init")
	}
	og := databases.OrderGoods{}
	goods, err := og.BatchFetchByOrderIDList(c.db, []uint64{orderID})
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
		logrus.Errorf("[updateOrderStatus] !order.Status.CheckNextStatusIsValid(status) currentStatus: %s, nextStatus: %s", order.Status.String(), status.String())
		return general_errors.OrderStatusFlowIncorrect
	}

	// 变更订单状态
	f := builder.FieldValues{
		"Status": status,
	}
	order.Status = status
	err := order.UpdateByIDWithMap(db, f)
	if err != nil {
		logrus.Errorf("[updateOrderStatus] order.UpdateByIDWithMap err: %v, orderID: %d, status: %s", err, order.OrderID, status.String())
		return general_errors.InternalError
	}

	return nil
}

func (c Controller) updateOrderDiscount(db sqlx.DBExecutor, order *databases.Order, discountAmount uint64) error {
	if order.DiscountAmount == discountAmount {
		return nil
	}
	if order.Status != enums.ORDER_STATUS__CREATED {
		return general_errors.NotAllowedChangeAmount
	}
	if discountAmount > order.TotalPrice {
		return general_errors.DiscountAmountOverflow
	}

	order.ActualAmount = order.TotalPrice - discountAmount
	f := builder.FieldValues{
		"DiscountAmount": discountAmount,
		"ActualAmount":   order.ActualAmount,
	}
	err := order.UpdateByIDWithMap(db, f)
	if err != nil {
		logrus.Errorf("[updateOrderDiscount] order.UpdateByIDWithMap err: %v, orderID: %d, discount: %d", err, order.OrderID, discountAmount)
		return general_errors.InternalError
	}
	return nil
}

func (c Controller) updateOrderLogistics(db sqlx.DBExecutor, order *databases.Order, logistics *databases.OrderLogistics, recipients, address, mobile string) (err error) {
	if logistics.Recipients == recipients && logistics.ShippingAddr == address && logistics.Mobile == mobile {
		return
	}
	if order.Status >= enums.ORDER_STATUS__DISPATCH {
		return general_errors.NotAllowedChangeLogistics
	}
	logistics.Recipients = recipients
	logistics.ShippingAddr = address
	logistics.Mobile = mobile
	f := builder.FieldValues{
		"Recipients":   logistics.Recipients,
		"ShippingAddr": logistics.ShippingAddr,
		"Mobile":       logistics.Mobile,
	}
	err = logistics.UpdateByIDWithMap(db, f)
	if err != nil {
		logrus.Errorf("[updateOrderLogistics] logistics.UpdateByIDWithMap err: %v, orderID: %d, logisticsID: %d, recipients: %s, address: %s, mobile: %s", err, order.OrderID, logistics.LogisticsID, recipients, address, mobile)
		return general_errors.InternalError
	}
	return
}

func (c Controller) updateCourierInfo(db sqlx.DBExecutor, logistics *databases.OrderLogistics, courierCompany, courierNumber string) error {
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
		logrus.Errorf("[updateCourierInfo] logistics.UpdateByIDWithMap err: %v, logisticsID: %d, courierCompany: %s, courierNumber: %s", err, logistics.LogisticsID, courierCompany, courierNumber)
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
		logrus.Errorf("[updateOrderRemark] order.UpdateByIDWithMap err: %v, orderID: %d, remark: %s", err, order.OrderID, remark)
		return general_errors.InternalError
	}
	return nil
}

func (c Controller) UpdateOrder(order *databases.Order, logistics *databases.OrderLogistics, params UpdateOrderParams, db sqlx.DBExecutor) (err error) {
	if !c.isInit {
		logrus.Panicf("[OrderController] not Init")
	}
	if db == nil {
		db = c.db
	}
	if params.Status != enums.ORDER_STATUS_UNKNOWN && params.Status != enums.ORDER_STATUS__CLOSED {
		if err = c.updateOrderStatus(db, order, params.Status); err != nil {
			return err
		}
	}
	if params.DiscountAmount != 0 {
		if err = c.updateOrderDiscount(db, order, params.DiscountAmount); err != nil {
			return err
		}
	}
	if params.Recipients != "" || params.ShippingAddr != "" || params.Mobile != "" {
		if err = c.updateOrderLogistics(db, order, logistics, params.Recipients, params.ShippingAddr, params.Mobile); err != nil {
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

	// 执行状态变更事件
	switch order.Status {
	case enums.ORDER_STATUS__PAID:
		// 获取支付流水
		flow, err := payment_flow.GetController().GetFlowByOrderAndUserID(order.OrderID, order.UserID, db)
		if err != nil {
			return err
		}
		err = c.eventHandler.OnOrderPaidEvent(db, order, flow)
	case enums.ORDER_STATUS__COMPLETE:
		err = c.eventHandler.OnOrderCompleteEvent(db, order)
	}

	return err
}

func (c Controller) CancelOrder(orderID, userID uint64, unlocker InventoryUnlock) error {
	if !c.isInit {
		logrus.Panicf("[OrderController] not Init")
	}
	var order *databases.Order
	var err error
	tx := sqlx.NewTasks(c.db)
	tx = tx.With(func(db sqlx.DBExecutor) error {
		order, _, err = c.GetOrder(orderID, userID, db, true)
		if err != nil {
			return err
		}

		if order.Status == enums.ORDER_STATUS__CLOSED {
			return general_errors.OrderCanceled
		}
		return nil
	})

	tx = tx.With(func(db sqlx.DBExecutor) error {
		return c.updateOrderStatus(db, order, enums.ORDER_STATUS__CLOSED)
	})

	tx = tx.With(func(db sqlx.DBExecutor) error {
		goods, err := c.GetOrderGoods(orderID)
		if err != nil {
			return err
		}

		for _, g := range goods {
			err = unlocker(db, g.GoodsID, g.Amount)
			if err != nil {
				return err
			}
		}

		return nil
	})

	err = tx.Do()
	if err != nil {
		logrus.Errorf("[CancelOrder] tx.Do() err: %v, orderID: %d, userID: %d", err, orderID, userID)
	}
	return err
}
