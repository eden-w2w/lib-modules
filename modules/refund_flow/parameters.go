package refund_flow

import (
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules"
)

type CreateRefundFlowRequest struct {
	// 交易单号
	PaymentFlowID uint64 `json:"paymentFlowID,string"`
	// 支付系统交易单号
	RemotePaymentFlowID string `json:"remotePaymentFlowID"`
	// 交易总额
	TotalAmount uint64 `json:"totalAmount"`
	// 退款总额
	RefundAmount uint64 `json:"refundAmount"`
}

type UpdateRefundFlowRequest struct {
	// 交易总额
	TotalAmount uint64 `json:"totalAmount" default:"0"`
	// 退款总额
	RefundAmount uint64 `json:"refundAmount" default:"0"`
	// 退款渠道
	Channel enums.RefundChannel `json:"channel" default:""`
	// 退款账户
	Account string `json:"account" default:""`
	// 退款状态
	Status enums.RefundStatus `json:"status" default:""`
}

type GetRefundFlowsRequest struct {
	// 支付系统退款单号
	RemoteFlowID string `in:"query" name:"remoteFlowID" default:""`
	// 交易单号
	PaymentFlowID uint64 `in:"query" name:"paymentFlowID,string" default:""`
	// 支付系统交易单号
	RemotePaymentFlowID string `in:"query" name:"remotePaymentFlowID" default:""`
	// 退款状态
	Status enums.RefundStatus `in:"query" name:"status" default:""`
	modules.Pagination
}

func (p GetRefundFlowsRequest) Conditions() builder.SqlCondition {
	var condition builder.SqlCondition
	var model = databases.RefundFlow{}
	if p.RemoteFlowID != "" {
		condition = builder.And(condition, model.FieldRemoteFlowID().Eq(p.RemoteFlowID))
	}
	if p.PaymentFlowID != 0 {
		condition = builder.And(condition, model.FieldPaymentFlowID().Eq(p.PaymentFlowID))
	}
	if p.RemotePaymentFlowID != "" {
		condition = builder.And(condition, model.FieldRemotePaymentFlowID().Eq(p.RemotePaymentFlowID))
	}
	if p.Status != enums.REFUND_STATUS_UNKNOWN {
		condition = builder.And(condition, model.FieldStatus().Eq(p.Status))
	}
	return condition
}

func (p GetRefundFlowsRequest) Additions() []builder.Addition {
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
