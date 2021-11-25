package enums

//go:generate eden generate enum --type-name=DiscountStatus
// api:enum
type DiscountStatus uint8

// 营销状态
const (
	DISCOUNT_STATUS_UNKNOWN  DiscountStatus = iota
	DISCOUNT_STATUS__READY                  // 就绪
	DISCOUNT_STATUS__PROCESS                // 进行中
	DISCOUNT_STATUS__STOP                   // 已停止
)
