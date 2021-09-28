package enums

//go:generate eden generate enum --type-name=PaymentStatus
// api:enum
type PaymentStatus uint8

// 支付状态
const (
	PAYMENT_STATUS_UNKNOWN  PaymentStatus = iota
	PAYMENT_STATUS__CREATED               // 未支付
	PAYMENT_STATUS__PROCESS               // 处理中
	PAYMENT_STATUS__SUCCESS               // 支付成功
	PAYMENT_STATUS__FAIL                  // 支付失败
	PAYMENT_STATUS__CLOSED                // 已关闭
)

var paymentStatusNext = map[PaymentStatus][]PaymentStatus{
	PAYMENT_STATUS__CREATED: {PAYMENT_STATUS__PROCESS, PAYMENT_STATUS__SUCCESS, PAYMENT_STATUS__FAIL, PAYMENT_STATUS__CLOSED},
	PAYMENT_STATUS__PROCESS: {PAYMENT_STATUS__SUCCESS, PAYMENT_STATUS__FAIL},
}

func (v PaymentStatus) CheckNextStatusIsValid(next PaymentStatus) bool {
	if v == PAYMENT_STATUS_UNKNOWN {
		return false
	}
	nextStatus := paymentStatusNext[v]
	for _, status := range nextStatus {
		if status == next {
			return true
		}
	}
	return false
}
