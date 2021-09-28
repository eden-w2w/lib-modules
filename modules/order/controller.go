package order

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/modules/id_generator"
	"github.com/eden-w2w/lib-modules/modules/payment_flow"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/eden-w2w/lib-modules/contants/enums"
	"github.com/eden-w2w/lib-modules/contants/errors"
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
	// 获取订单总额与库中物料进行比对
	var totalPrice uint64 = 0
	var goodsList = make([]CreateOrderGoodsModelParams, 0)
	for _, g := range p.Goods {
		goods := databases.Goods{GoodsID: g.GoodsID}
		err := goods.FetchByGoodsID(c.db)
		if err != nil {
			logrus.Errorf("[CreateOrder] goods.FetchByGoodsID(c.db) err: %v, goodsID: %d", err, g.GoodsID)
			return nil, errors.NotFound.StatusError().WithDesc("商品无法找到")
		}
		totalPrice += goods.Price * uint64(g.Amount)
		goodsList = append(goodsList, CreateOrderGoodsModelParams{
			Goods:  goods,
			Amount: g.Amount,
		})
	}
	if totalPrice != p.TotalPrice {
		logrus.Errorf("[CreateOrder] totalPrice != p.TotalPrice totalPrice: %d, p.TotalPrice: %d", totalPrice, p.TotalPrice)
		return nil, errors.BadRequest.StatusError().WithDesc("订单总额与商品总价不一致")
	}
	if len(goodsList) == 0 {
		logrus.Errorf("[CreateOrder] len(goodsList) == 0")
		return nil, errors.BadRequest.StatusError().WithDesc("商品列表为空")
	}

	// 创建订单
	id, _ := id_generator.GetGenerator().GenerateUniqueID()
	order := &databases.Order{
		OrderID:       id,
		UserID:        p.UserID,
		RefererID:     p.RefererID,
		TotalPrice:    p.TotalPrice,
		PaymentMethod: p.PaymentMethod,
		Remark:        p.Remark,
		Recipients:    p.Recipients,
		ShippingAddr:  p.ShippingAddr,
		Mobile:        p.Mobile,
		Status:        enums.ORDER_STATUS__CREATED,
		ExpiredAt:     datatypes.MySQLTimestamp(time.Now().Add(c.orderExpireDuration)),
	}

	tx := sqlx.NewTasks(c.db)
	tx = tx.With(func(db sqlx.DBExecutor) error {
		return order.Create(db)
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
				OrderID:        id,
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

	err := tx.Do()
	if err != nil {
		logrus.Errorf("[CreateOrder] err: %v, params: %+v", err, p)
		return nil, errors.InternalError
	}

	return order, nil
}

func (c Controller) GetOrder(orderID, userID uint64, db sqlx.DBExecutor, forUpdate bool) (order *databases.Order, err error) {
	if !c.isInit {
		logrus.Panicf("[OrderController] not Init")
	}
	order = &databases.Order{OrderID: orderID}
	if forUpdate {
		err = order.FetchByOrderIDForUpdate(db)
	} else {
		err = order.FetchByOrderID(c.db)
	}
	if err != nil {
		if sqlx.DBErr(err).IsNotFound() {
			return nil, errors.OrderNotFound
		}
		logrus.Errorf("[GetOrder] err: %v, orderID: %d", err, orderID)
		return nil, errors.InternalError
	}
	if order.UserID != userID {
		logrus.Errorf("[GetOrder] order.UserID != userID, order.UserID: %d, userID: %d", order.UserID, userID)
		return nil, errors.Forbidden
	}
	return order, nil
}

func (c Controller) GetOrders(p GetOrdersParams) ([]databases.Order, error) {
	if !c.isInit {
		logrus.Panicf("[OrderController] not Init")
	}
	order := databases.Order{}
	orders, err := order.List(c.db, p.Conditions(), p.Additions()...)
	if err != nil {
		logrus.Errorf("[GetOrders] order.List err: %v, params: %+v", err, p)
		return nil, errors.InternalError
	}
	return orders, nil
}

func (c Controller) GetOrderGoods(orderID uint64) ([]databases.OrderGoods, error) {
	if !c.isInit {
		logrus.Panicf("[OrderController] not Init")
	}
	og := databases.OrderGoods{}
	goods, err := og.BatchFetchByOrderIDList(c.db, []uint64{orderID})
	if err != nil {
		logrus.Errorf("[GetOrderGoods] og.BatchFetchByOrderIDList err: %v, orderID: %d", err, orderID)
		return nil, errors.InternalError
	}
	return goods, nil
}

func (c Controller) UpdateOrderStatusWithDB(db sqlx.DBExecutor, order *databases.Order, status enums.OrderStatus) error {
	if !c.isInit {
		logrus.Panicf("[OrderController] not Init")
	}
	if order.Status == status {
		return nil
	}

	// 状态流转检查
	if !order.Status.CheckNextStatusIsValid(status) {
		logrus.Errorf("[UpdateOrderStatusWithDB] !order.Status.CheckNextStatusIsValid(status) currentStatus: %s, nextStatus: %s", order.Status.String(), status.String())
		return errors.OrderStatusFlowIncorrect
	}

	// 变更订单状态
	f := builder.FieldValues{
		"Status": status,
	}
	order.Status = status
	err := order.UpdateByIDWithMap(db, f)
	if err != nil {
		logrus.Errorf("[UpdateOrderStatusWithDB] order.UpdateByIDWithMap err: %v, orderID: %d, status: %s", err, order.OrderID, status.String())
		return errors.InternalError
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
		order, err = c.GetOrder(orderID, userID, db, true)
		if err != nil {
			return err
		}

		if order.Status == enums.ORDER_STATUS__CLOSED {
			return errors.OrderCanceled
		}
		return nil
	})

	tx = tx.With(func(db sqlx.DBExecutor) error {
		return c.UpdateOrderStatusWithDB(db, order, enums.ORDER_STATUS__CLOSED)
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
