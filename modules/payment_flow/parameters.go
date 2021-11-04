package payment_flow

import (
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules"
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

type GetPaymentFlowsParams struct {
	// 用户ID
	UserID uint64 `in:"query" default:"" name:"userID,string"`
	// 关联订单号
	OrderID uint64 `in:"query" default:"" name:"orderID,string"`
	// 支付方式
	PaymentMethod enums.PaymentMethod `in:"query" default:"" name:"paymentMethod"`
	// 支付状态
	Status enums.PaymentStatus `in:"query" default:"" name:"status"`

	modules.Pagination
}

func (p GetPaymentFlowsParams) Conditions() builder.SqlCondition {
	var condition builder.SqlCondition
	var model = databases.PaymentFlow{}
	if p.UserID != 0 {
		condition = builder.And(condition, model.FieldUserID().Eq(p.UserID))
	}
	if p.OrderID != 0 {
		condition = builder.And(condition, model.FieldOrderID().Eq(p.OrderID))
	}
	if p.PaymentMethod != enums.PAYMENT_METHOD_UNKNOWN {
		condition = builder.And(condition, model.FieldPaymentMethod().Eq(p.PaymentMethod))
	}
	if p.Status != enums.PAYMENT_STATUS_UNKNOWN {
		condition = builder.And(condition, model.FieldStatus().Eq(p.Status))
	}
	return condition
}

func (p GetPaymentFlowsParams) Additions() []builder.Addition {
	var additions = make([]builder.Addition, 0)

	if p.Size != 0 {
		limit := builder.Limit(int64(p.Size))
		if p.Offset != 0 {
			limit = limit.Offset(int64(p.Offset))
		}
		additions = append(additions, limit)
	}

	additions = append(additions, builder.OrderBy(builder.DescOrder((&databases.Order{}).FieldCreatedAt())))

	return additions
}
