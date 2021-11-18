package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidBookingStatus = errors.New("invalid BookingStatus")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("BookingStatus", map[string]string{
		"COMPLETE": "已结束",
		"PROCESS":  "进行中",
		"READY":    "待开始",
	})
}

func ParseBookingStatusFromString(s string) (BookingStatus, error) {
	switch s {
	case "":
		return BOOKING_STATUS_UNKNOWN, nil
	case "COMPLETE":
		return BOOKING_STATUS__COMPLETE, nil
	case "PROCESS":
		return BOOKING_STATUS__PROCESS, nil
	case "READY":
		return BOOKING_STATUS__READY, nil
	}
	return BOOKING_STATUS_UNKNOWN, InvalidBookingStatus
}

func ParseBookingStatusFromLabelString(s string) (BookingStatus, error) {
	switch s {
	case "":
		return BOOKING_STATUS_UNKNOWN, nil
	case "已结束":
		return BOOKING_STATUS__COMPLETE, nil
	case "进行中":
		return BOOKING_STATUS__PROCESS, nil
	case "待开始":
		return BOOKING_STATUS__READY, nil
	}
	return BOOKING_STATUS_UNKNOWN, InvalidBookingStatus
}

func (BookingStatus) EnumType() string {
	return "BookingStatus"
}

func (BookingStatus) Enums() map[int][]string {
	return map[int][]string{
		int(BOOKING_STATUS__COMPLETE): {"COMPLETE", "已结束"},
		int(BOOKING_STATUS__PROCESS):  {"PROCESS", "进行中"},
		int(BOOKING_STATUS__READY):    {"READY", "待开始"},
	}
}

func (v BookingStatus) String() string {
	switch v {
	case BOOKING_STATUS_UNKNOWN:
		return ""
	case BOOKING_STATUS__COMPLETE:
		return "COMPLETE"
	case BOOKING_STATUS__PROCESS:
		return "PROCESS"
	case BOOKING_STATUS__READY:
		return "READY"
	}
	return "UNKNOWN"
}

func (v BookingStatus) Label() string {
	switch v {
	case BOOKING_STATUS_UNKNOWN:
		return ""
	case BOOKING_STATUS__COMPLETE:
		return "已结束"
	case BOOKING_STATUS__PROCESS:
		return "进行中"
	case BOOKING_STATUS__READY:
		return "待开始"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*BookingStatus)(nil)

func (v BookingStatus) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidBookingStatus
	}
	return []byte(str), nil
}

func (v *BookingStatus) UnmarshalText(data []byte) (err error) {
	*v, err = ParseBookingStatusFromString(string(bytes.ToUpper(data)))
	return
}
