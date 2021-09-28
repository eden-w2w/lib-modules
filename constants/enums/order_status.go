package enums

//go:generate eden generate enum --type-name=OrderStatus
// api:enum
type OrderStatus uint8

// 订单状态
const (
	ORDER_STATUS_UNKNOWN   OrderStatus = iota
	ORDER_STATUS__CREATED              // 待支付
	ORDER_STATUS__PAID                 // 已支付
	ORDER_STATUS__CONFIRM              // 待发货
	ORDER_STATUS__DISPATCH             // 已发货
	ORDER_STATUS__COMPLETE             // 已完成
	ORDER_STATUS__REFUND               // 退款中
	ORDER_STATUS__CLOSED               // 已关闭
)

var orderStatusNext = map[OrderStatus][]OrderStatus{
	ORDER_STATUS__CREATED:  {ORDER_STATUS__PAID, ORDER_STATUS__CLOSED},
	ORDER_STATUS__PAID:     {ORDER_STATUS__CONFIRM, ORDER_STATUS__REFUND},
	ORDER_STATUS__CONFIRM:  {ORDER_STATUS__DISPATCH, ORDER_STATUS__REFUND},
	ORDER_STATUS__DISPATCH: {ORDER_STATUS__COMPLETE, ORDER_STATUS__REFUND},
	ORDER_STATUS__COMPLETE: {ORDER_STATUS__REFUND},
	ORDER_STATUS__REFUND:   {ORDER_STATUS__CLOSED},
	ORDER_STATUS__CLOSED:   {},
}

func (v OrderStatus) CheckNextStatusIsValid(next OrderStatus) bool {
	if v == ORDER_STATUS_UNKNOWN {
		return false
	}
	nextStatus := orderStatusNext[v]
	for _, status := range nextStatus {
		if status == next {
			return true
		}
	}
	return false
}
