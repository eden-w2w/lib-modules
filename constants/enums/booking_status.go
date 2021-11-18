package enums

//go:generate eden generate enum --type-name=BookingStatus
// api:enum
type BookingStatus uint8

// 预售状态
const (
	BOOKING_STATUS_UNKNOWN   BookingStatus = iota
	BOOKING_STATUS__READY                  // 待开始
	BOOKING_STATUS__PROCESS                // 进行中
	BOOKING_STATUS__COMPLETE               // 已结束
)
