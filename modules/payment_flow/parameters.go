package payment_flow

import (
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

type CreatePaymentFlowParams struct {
	// 用户ID
	UserID uint64 `in:"body" default:"" json:"userID,string"`
	// 关联订单号
	OrderID uint64 `in:"body" json:"orderID,string"`
	// 支付金额
	Amount uint64 `in:"body" default:"" json:"amount"`
	// 支付方式
	PaymentMethod enums.PaymentMethod `in:"body" json:"paymentMethod"`
}

type CreatePaymentFlowResponse struct {
	PaymentFlow  *databases.PaymentFlow                  `json:"paymentFlow"`
	WechatPrepay *jsapi.PrepayWithRequestPaymentResponse `json:"prepay"`
}
