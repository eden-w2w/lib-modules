package events

import (
	"fmt"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/booking_flow"
	"github.com/eden-w2w/lib-modules/modules/goods"
	"github.com/eden-w2w/lib-modules/modules/order"
	"github.com/eden-w2w/lib-modules/modules/payment_flow"
	"github.com/eden-w2w/lib-modules/modules/promotion_flow"
	"github.com/eden-w2w/lib-modules/modules/refund_flow"
	"github.com/eden-w2w/lib-modules/modules/user"
	"github.com/eden-w2w/lib-modules/modules/wechat"
	"github.com/eden-w2w/lib-modules/pkg/strings"
	"github.com/eden-w2w/wechatpay-go/core"
	"github.com/eden-w2w/wechatpay-go/services/payments/jsapi"
	"github.com/eden-w2w/wechatpay-go/services/refunddomestic"
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
	"github.com/sirupsen/logrus"
	"sync"
)

type InventoryLock func(db sqlx.DBExecutor, goodsID uint64, amount uint32) error
type InventoryUnlock func(db sqlx.DBExecutor, goodsID uint64, amount uint32) error

type OrderEvent struct {
	config   wechat.Wechat
	locker   InventoryLock
	unlocker InventoryUnlock
	sync.Mutex
}

func NewOrderEvent(config wechat.Wechat, inventoryLocker InventoryLock, inventoryUnlocker InventoryUnlock) *OrderEvent {
	return &OrderEvent{
		config:   config,
		locker:   inventoryLocker,
		unlocker: inventoryUnlocker,
	}
}

func (o *OrderEvent) OnOrderCreateEvent(
	db sqlx.DBExecutor,
	order *databases.Order,
	goodsList []databases.OrderGoods,
) error {
	for _, item := range goodsList {
		if item.IsBooking != datatypes.BOOL_TRUE {
			err := o.locker(db, item.GoodsID, item.Amount)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *OrderEvent) OnOrderConfirmEvent(db sqlx.DBExecutor, oModel *databases.Order) error {
	// 防止并发导致预售销量错误，手动加锁
	o.Lock()
	defer o.Unlock()

	// 增加预售销量
	goodsList, err := order.GetController().GetOrderGoods(oModel.OrderID, db)
	if err != nil {
		return err
	}

	for _, orderGoods := range goodsList {
		if orderGoods.IsBooking != datatypes.BOOL_TRUE {
			continue
		}

		flows, err := booking_flow.GetController().GetBookingFlowByGoodsID(orderGoods.GoodsID)
		if err != nil {
			return err
		}

		if len(flows) == 0 {
			logrus.Warningf(
				"[OnOrderConfirmEvent] booking_flow.GetController().GetBookingFlowByGoodsIDAndStatus not found, goodsID: %d",
				orderGoods.GoodsID,
			)
			continue
		}

		var sales = flows[0].Sales + orderGoods.Amount
		err = booking_flow.GetController().UpdateBookingFlow(
			&flows[0], booking_flow.UpdateBookingFlowParams{
				Sales: &sales,
			}, db,
		)
		if err != nil {
			return err
		}
	}

	msg := &subscribe.Message{
		ToUser:     oModel.UserOpenID,
		TemplateID: o.config.ConfirmMessageTemplateID,
		Page:       fmt.Sprintf("%s?orderID=%d", o.config.OrderPage, oModel.OrderID),
		Data: map[string]*subscribe.DataItem{
			"character_string1": {Value: oModel.OrderID},
			"thing2":            {Value: "已确认"},
			"thing4":            {Value: "已收到您的付款，正在备货哟"},
		},
		MiniprogramState: o.config.ProgramState,
	}
	_ = wechat.GetController().SendSubscribeMessage(msg)
	return nil
}

func (o *OrderEvent) OnOrderPaidEvent(
	db sqlx.DBExecutor,
	order *databases.Order,
	payment *databases.PaymentFlow,
) error {
	return nil
}

func (o *OrderEvent) OnOrderDispatchEvent(
	db sqlx.DBExecutor,
	order *databases.Order,
	logistics *databases.OrderLogistics,
	goodsList []databases.OrderGoods,
) error {
	// 若为预售则需要验证商品是否库存充足并扣减库存
	for _, orderGoods := range goodsList {
		if orderGoods.IsBooking == datatypes.BOOL_TRUE {
			// 检查商品库存
			gModel, err := goods.GetController().GetGoodsByID(orderGoods.GoodsID, db, true)
			if err != nil {
				return err
			}
			if gModel.Inventory < uint64(orderGoods.Amount) {
				return general_errors.GoodsInventoryShortage.StatusError().WithMsg("商品库存不足，请先增加库存")
			}
			err = o.locker(db, gModel.GoodsID, orderGoods.Amount)
			if err != nil {
				return err
			}
		}
	}

	msg := &subscribe.Message{
		ToUser:     order.UserOpenID,
		TemplateID: o.config.DispatchMessageTemplateID,
		Page:       fmt.Sprintf("%s?orderID=%d", o.config.OrderPage, order.OrderID),
		Data: map[string]*subscribe.DataItem{
			"character_string5": {Value: order.OrderID},
			"thing7":            {Value: strings.ShortenString(logistics.Recipients, 17, "...")},
			"thing8":            {Value: strings.ShortenString(logistics.ShippingAddr, 17, "...")},
			"thing2":            {Value: strings.ShortenString(logistics.CourierCompany, 17, "...")},
			"character_string3": {Value: logistics.CourierNumber},
		},
		MiniprogramState: o.config.ProgramState,
	}
	_ = wechat.GetController().SendSubscribeMessage(msg)
	return nil
}

func (o *OrderEvent) OnOrderCompleteEvent(
	db sqlx.DBExecutor,
	order *databases.Order,
	logistics *databases.OrderLogistics,
	goodsList []databases.OrderGoods,
) error {
	// 更新商品销量
	for _, g := range goodsList {
		gModel, err := goods.GetController().GetGoodsByID(g.GoodsID, db, true)
		if err != nil {
			return err
		}

		err = goods.GetController().UpdateGoods(
			gModel.GoodsID, goods.UpdateGoodsParams{
				Sales: gModel.Sales + g.Amount,
			}, db,
		)
		if err != nil {
			return err
		}
	}

	var goodsName = ""
	if len(goodsList) > 0 {
		goodsName = goodsList[0].Name
	}
	msg := &subscribe.Message{
		ToUser:     order.UserOpenID,
		TemplateID: o.config.CompleteMessageTemplateID,
		Page:       fmt.Sprintf("%s?orderID=%d", o.config.OrderPage, order.OrderID),
		Data: map[string]*subscribe.DataItem{
			"character_string5": {Value: order.OrderID},
			"thing1":            {Value: strings.ShortenString(goodsName, 17, "...")},
			"thing2":            {Value: strings.ShortenString(logistics.Recipients, 17, "...")},
			"phone_number3":     {Value: strings.ShortenString(logistics.Mobile, 14, "...")},
			"time7":             {Value: order.UpdatedAt.Format("2006-01-02 15:04:05")},
		},
		MiniprogramState: o.config.ProgramState,
	}
	_ = wechat.GetController().SendSubscribeMessage(msg)

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

	// 获取订单创建者
	orderUser, err := user.GetController().GetUserByUserID(order.UserID, db, false)
	if err != nil {
		return err
	}

	// 如果创建者没有推荐者信息则无需计算提成
	if orderUser.RefererID == 0 {
		return nil
	}
	// 获取推荐者
	refererUser, err := user.GetController().GetUserByUserID(orderUser.RefererID, db, false)
	if err != nil {
		return err
	}

	// 创建提成流水
	proCtrl := promotion_flow.GetController()
	_, err = proCtrl.CreatePromotionFlow(
		promotion_flow.CreatePromotionFlowParams{
			UserID:          refererUser.UserID,
			UserNickName:    refererUser.NickName,
			UserOpenID:      refererUser.OpenID,
			RefererID:       orderUser.UserID,
			RefererNickName: orderUser.NickName,
			RefererOpenID:   orderUser.OpenID,
			Amount:          flow.Amount,
			PaymentFlowID:   flow.FlowID,
		}, db,
	)

	return err
}

func (o *OrderEvent) OnOrderCloseEvent(
	db sqlx.DBExecutor,
	order *databases.Order,
	goodsList []databases.OrderGoods,
) error {
	// 解锁物料
	for _, g := range goodsList {
		if g.IsBooking != datatypes.BOOL_TRUE {
			err := o.unlocker(db, g.GoodsID, g.Amount)
			if err != nil {
				return err
			}
		} else {
			o.Lock()
			// 恢复预售销量
			flows, err := booking_flow.GetController().GetBookingFlowByGoodsID(g.GoodsID)
			if err != nil {
				o.Unlock()
				return err
			}

			if len(flows) == 0 {
				logrus.Warningf(
					"[OnOrderCloseEvent] booking_flow.GetController().GetBookingFlowByGoodsIDAndStatus not found, goodsID: %d",
					g.GoodsID,
				)
				o.Unlock()
				continue
			}

			if flows[0].Sales < g.Amount {
				flows[0].Sales = g.Amount
			}
			var sales = flows[0].Sales - g.Amount
			err = booking_flow.GetController().UpdateBookingFlow(
				&flows[0], booking_flow.UpdateBookingFlowParams{
					Sales: &sales,
				}, db,
			)
			if err != nil {
				o.Unlock()
				return err
			}
			o.Unlock()
		}
	}

	// 获取支付流水
	flows, err := payment_flow.GetController().GetFlowByOrderIDAndStatus(
		order.OrderID,
		order.UserID,
		[]enums.PaymentStatus{enums.PAYMENT_STATUS__SUCCESS, enums.PAYMENT_STATUS__CREATED},
		db,
	)
	if err != nil {
		return err
	}

	if flows == nil || len(flows) == 0 {
		return nil
	}

	for _, flow := range flows {
		if flow.Status == enums.PAYMENT_STATUS__SUCCESS {
			// 查询是否存在关联的佣金流水单
			proCtrl := promotion_flow.GetController()
			promotions, _, err := proCtrl.GetPromotionFlows(
				promotion_flow.GetPromotionFlowParams{
					PaymentFlowID:   flow.FlowID,
					IsNotSettlement: datatypes.BOOL_TRUE,
				}, false,
			)
			if err != nil {
				return err
			}

			if len(promotions) > 0 {
				err = promotions[0].SoftDeleteByFlowID(db)
				if err != nil {
					logrus.Errorf(
						"[OnOrderCloseEvent] promotions[0].SoftDeleteByFlowID(db) err: %v, flowID: %d",
						err,
						flow.FlowID,
					)
					return general_errors.InternalError
				}
			}

			err = o.refundPayment(flow, db)
			if err != nil {
				return err
			}
		} else if flow.Status == enums.PAYMENT_STATUS__CREATED {
			// 微信关单
			err = wechat.GetController().CloseOrder(
				jsapi.CloseOrderRequest{
					OutTradeNo: core.String(fmt.Sprintf("%d", flow.FlowID)),
					Mchid:      core.String(o.config.MerchantID),
				},
			)
			if err != nil {
				return general_errors.InternalError
			}
			tran, err := wechat.GetController().QueryOrderByOutTradeNo(
				jsapi.QueryOrderByOutTradeNoRequest{
					OutTradeNo: core.String(fmt.Sprintf("%d", flow.FlowID)),
					Mchid:      core.String(o.config.MerchantID),
				},
			)
			if err != nil {
				return general_errors.InternalError
			}
			tradeState, err := enums.ParseWechatTradeStateFromString(*tran.TradeState)
			if err != nil {
				logrus.Errorf(
					"[OnOrderCloseEvent] enums.ParseWechatTradeStateFromString err: %v, TradeState: %s",
					err,
					*tran.TradeState,
				)
				return err
			}
			// 更新支付单
			err = payment_flow.GetController().UpdatePaymentFlowStatus(
				&flow,
				tradeState.ToPaymentStatus(),
				tran,
				db,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *OrderEvent) refundPayment(flow databases.PaymentFlow, db sqlx.DBExecutor) error {
	// 变更支付单状态为转入退款
	err := payment_flow.GetController().UpdatePaymentFlowStatus(&flow, enums.PAYMENT_STATUS__REFUND, nil, db)
	if err != nil {
		return err
	}

	// 创建退款单
	refundFlow, err := refund_flow.GetController().CreateRefundFlow(
		refund_flow.CreateRefundFlowRequest{
			PaymentFlowID:       flow.FlowID,
			RemotePaymentFlowID: flow.RemoteFlowID,
			TotalAmount:         flow.Amount,
			RefundAmount:        flow.Amount,
		}, db,
	)
	if err != nil {
		return err
	}

	// 微信支付退款
	wechatRefund, err := wechat.GetController().CreateRefund(
		refunddomestic.CreateRequest{
			SubMchid:      nil,
			TransactionId: core.String(flow.RemoteFlowID),
			OutTradeNo:    core.String(fmt.Sprintf("%d", flow.FlowID)),
			OutRefundNo:   core.String(fmt.Sprintf("%d", refundFlow.FlowID)),
			Reason:        nil,
			NotifyUrl:     core.String(o.config.RefundNotifyUrl),
			FundsAccount:  nil,
			Amount: &refunddomestic.AmountReq{
				Refund:   core.Int64(int64(refundFlow.RefundAmount)),
				From:     nil,
				Total:    core.Int64(int64(flow.Amount)),
				Currency: core.String("CNY"),
			},
			GoodsDetail: nil,
		},
	)
	if err != nil {
		return err
	}

	return refund_flow.GetController().UpdateRefundFlowRemoteID(refundFlow.FlowID, *wechatRefund.RefundId, db)
}
