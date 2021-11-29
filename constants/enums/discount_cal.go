package enums

//go:generate eden generate enum --type-name=DiscountCal
// api:enum
type DiscountCal uint8

// 计算方式
const (
	DISCOUNT_CAL_UNKNOWN         DiscountCal = iota
	DISCOUNT_CAL__UNIT                       // 单价立减
	DISCOUNT_CAL__MULTISTEP                  // 总价阶梯式立减
	DISCOUNT_CAL__MULTISTEP_UNIT             // 总价阶梯式单价立减
)
