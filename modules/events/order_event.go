package events

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/payment_flow"
	"github.com/eden-w2w/lib-modules/modules/promotion_flow"
	"github.com/eden-w2w/lib-modules/modules/user"
	"github.com/sirupsen/logrus"
)

type OrderEvent struct {
}

func (o *OrderEvent) OnOrderCreateEvent(db sqlx.DBExecutor, order *databases.Order) error {
	return nil
}

func (o *OrderEvent) OnOrderPaidEvent(db sqlx.DBExecutor, order *databases.Order, payment *databases.PaymentFlow) error {
	return nil
}

func (o *OrderEvent) OnOrderCompleteEvent(db sqlx.DBExecutor, order *databases.Order) error {
	// 获取支付流水
	flows, err := payment_flow.GetController().MustGetFlowByOrderIDAndStatus(order.OrderID, order.UserID, []enums.PaymentStatus{enums.PAYMENT_STATUS__SUCCESS}, db)
	if err != nil {
		return err
	}

	flow := flows[0]

	// 获取订单创建者
	orderUser, err := user.GetController().GetUserByUserID(order.UserID, db, true)
	if err != nil {
		return err
	}

	// 如果创建者没有推荐者信息则无需计算提成
	if orderUser.RefererID == 0 {
		return nil
	}
	// 获取推荐者
	refererUser, err := user.GetController().GetUserByUserID(orderUser.RefererID, db, true)
	if err != nil {
		return err
	}

	// 创建提成流水
	proCtrl := promotion_flow.GetController()
	_, err = proCtrl.CreatePromotionFlow(promotion_flow.CreatePromotionFlowParams{
		UserID:          refererUser.UserID,
		UserNickName:    refererUser.NickName,
		UserOpenID:      refererUser.OpenID,
		RefererID:       orderUser.UserID,
		RefererNickName: orderUser.NickName,
		RefererOpenID:   orderUser.OpenID,
		Amount:          flow.Amount,
		PaymentFlowID:   flow.FlowID,
	}, db)

	// 更新提成概况
	return err
}

func (o *OrderEvent) OnOrderCloseEvent(db sqlx.DBExecutor, order *databases.Order) error {
	// 获取支付流水
	flows, err := payment_flow.GetController().MustGetFlowByOrderIDAndStatus(order.OrderID, order.UserID, []enums.PaymentStatus{enums.PAYMENT_STATUS__SUCCESS}, db)
	if err != nil {
		return err
	}

	flow := flows[0]

	// 查询是否存在关联的佣金流水单
	proCtrl := promotion_flow.GetController()
	promotions, _, err := proCtrl.GetPromotionFlows(promotion_flow.GetPromotionFlowParams{
		PaymentFlowID:   flow.FlowID,
		IsNotSettlement: datatypes.BOOL_TRUE,
	}, false)
	if err != nil {
		return err
	}

	if len(promotions) > 0 {
		err = promotions[0].SoftDeleteByFlowID(db)
		if err != nil {
			logrus.Errorf("[OnOrderCloseEvent] promotions[0].SoftDeleteByFlowID(db) err: %v, flowID: %d", err, flow.FlowID)
			return general_errors.InternalError
		}
	}

	return nil
}

func NewOrderEvent() *OrderEvent {
	return &OrderEvent{}
}
