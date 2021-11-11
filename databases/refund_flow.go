package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
)

//go:generate eden generate model RefundFlow --database Config.DB --with-comments
//go:generate eden generate tag RefundFlow --defaults=true
// @def primary ID
// @def unique_index U_flow_id FlowID
// @def index I_status Status
// @def index I_remote_id RemoteFlowID PaymentFlowID RemotePaymentFlowID
type RefundFlow struct {
	datatypes.PrimaryID
	// 业务ID
	FlowID uint64 `json:"flowID,string" db:"f_flow_id"`
	// 支付系统退款单号
	RemoteFlowID string `json:"remoteFlowID" db:"f_remote_flow_id,default=''"`
	// 交易单号
	PaymentFlowID uint64 `json:"paymentFlowID,string" db:"f_payment_flow_id"`
	// 支付系统交易单号
	RemotePaymentFlowID string `json:"remotePaymentFlowID" db:"f_remote_payment_flow_id"`
	// 退款渠道
	Channel enums.RefundChannel `json:"channel" db:"f_channel,default=0"`
	// 退款账户
	Account string `json:"account" db:"f_account,null"`
	// 退款状态
	Status enums.RefundStatus `json:"status" db:"f_status"`
	// 交易总额
	TotalAmount uint64 `json:"totalAmount" db:"f_total_amount"`
	// 退款总额
	RefundAmount uint64 `json:"refundAmount" db:"f_refund_amount"`
	datatypes.OperateTime
}
