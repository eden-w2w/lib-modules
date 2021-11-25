package enums

//go:generate eden generate enum --type-name=DiscountType
// api:enum
type DiscountType uint8

// 立减类型
const (
	DISCOUNT_TYPE_UNKNOWN      DiscountType = iota
	DISCOUNT_TYPE__ALL                      // 全场立减
	DISCOUNT_TYPE__ALL_PERCENT              // 全场折扣
)
