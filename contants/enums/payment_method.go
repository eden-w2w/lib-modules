package enums

//go:generate eden generate enum --type-name=PaymentMethod
// api:enum
type PaymentMethod uint8

// Test
const (
	PAYMENT_METHOD_UNKNOWN PaymentMethod = iota
	PAYMENT_METHOD__WECHAT               // 微信支付
)
