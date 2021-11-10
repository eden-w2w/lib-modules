package enums

//go:generate eden generate enum --type-name=RefundStatus
// api:enum
type RefundStatus uint8

// 退款状态
const (
	REFUND_STATUS_UNKNOWN     RefundStatus = iota
	REFUND_STATUS__SUCCESS                 // 退款成功
	REFUND_STATUS__CLOSED                  // 退款关闭
	REFUND_STATUS__PROCESSING              // 退款处理中
	REFUND_STATUS__ABNORMAL                // 退款异常
)
