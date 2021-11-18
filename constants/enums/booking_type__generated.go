package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidBookingType = errors.New("invalid BookingType")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("BookingType", map[string]string{
		"AUTO":   "自动",
		"MANUAL": "手动",
	})
}

func ParseBookingTypeFromString(s string) (BookingType, error) {
	switch s {
	case "":
		return BOOKING_TYPE_UNKNOWN, nil
	case "AUTO":
		return BOOKING_TYPE__AUTO, nil
	case "MANUAL":
		return BOOKING_TYPE__MANUAL, nil
	}
	return BOOKING_TYPE_UNKNOWN, InvalidBookingType
}

func ParseBookingTypeFromLabelString(s string) (BookingType, error) {
	switch s {
	case "":
		return BOOKING_TYPE_UNKNOWN, nil
	case "自动":
		return BOOKING_TYPE__AUTO, nil
	case "手动":
		return BOOKING_TYPE__MANUAL, nil
	}
	return BOOKING_TYPE_UNKNOWN, InvalidBookingType
}

func (BookingType) EnumType() string {
	return "BookingType"
}

func (BookingType) Enums() map[int][]string {
	return map[int][]string{
		int(BOOKING_TYPE__AUTO):   {"AUTO", "自动"},
		int(BOOKING_TYPE__MANUAL): {"MANUAL", "手动"},
	}
}

func (v BookingType) String() string {
	switch v {
	case BOOKING_TYPE_UNKNOWN:
		return ""
	case BOOKING_TYPE__AUTO:
		return "AUTO"
	case BOOKING_TYPE__MANUAL:
		return "MANUAL"
	}
	return "UNKNOWN"
}

func (v BookingType) Label() string {
	switch v {
	case BOOKING_TYPE_UNKNOWN:
		return ""
	case BOOKING_TYPE__AUTO:
		return "自动"
	case BOOKING_TYPE__MANUAL:
		return "手动"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*BookingType)(nil)

func (v BookingType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidBookingType
	}
	return []byte(str), nil
}

func (v *BookingType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseBookingTypeFromString(string(bytes.ToUpper(data)))
	return
}
