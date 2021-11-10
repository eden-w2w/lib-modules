package enums

//go:generate eden generate enum --type-name=RefundChannel
// api:enum
type RefundChannel uint8

// 退款渠道
const (
	REFUND_CHANNEL_UNKNOWN         RefundChannel = iota
	REFUND_CHANNEL__ORIGINAL                     // 原路退款
	REFUND_CHANNEL__BALANCE                      // 退款到余额
	REFUND_CHANNEL__OTHER_BALANCE                // 原账户异常退到其他余额账户
	REFUND_CHANNEL__OTHER_BANKCARD               // 原银行卡异常退到其他银行卡
)
