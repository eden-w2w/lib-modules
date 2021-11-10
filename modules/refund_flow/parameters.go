package refund_flow

type CreateRefundFlowRequest struct {
	// 支付系统退款单号
	RemoteFlowID string `json:"remoteFlowID"`
	// 交易单号
	PaymentFlowID uint64 `json:"paymentFlowID,string"`
	// 支付系统交易单号
	RemotePaymentFlowID string `json:"remotePaymentFlowID"`
	// 交易总额
	TotalAmount uint64 `json:"totalAmount"`
	// 退款总额
	RefundAmount uint64 `json:"refundAmount"`
}
