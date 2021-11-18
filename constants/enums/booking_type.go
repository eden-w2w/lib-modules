package enums

//go:generate eden generate enum --type-name=BookingType
// api:enum
type BookingType uint8

// 预售模式
const (
	BOOKING_TYPE_UNKNOWN BookingType = iota
	BOOKING_TYPE__MANUAL             // 手动
	BOOKING_TYPE__AUTO               // 自动
)
