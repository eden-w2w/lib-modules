package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidDiscountStatus = errors.New("invalid DiscountStatus")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("DiscountStatus", map[string]string{
		"STOP":    "已停止",
		"PROCESS": "进行中",
		"READY":   "就绪",
	})
}

func ParseDiscountStatusFromString(s string) (DiscountStatus, error) {
	switch s {
	case "":
		return DISCOUNT_STATUS_UNKNOWN, nil
	case "STOP":
		return DISCOUNT_STATUS__STOP, nil
	case "PROCESS":
		return DISCOUNT_STATUS__PROCESS, nil
	case "READY":
		return DISCOUNT_STATUS__READY, nil
	}
	return DISCOUNT_STATUS_UNKNOWN, InvalidDiscountStatus
}

func ParseDiscountStatusFromLabelString(s string) (DiscountStatus, error) {
	switch s {
	case "":
		return DISCOUNT_STATUS_UNKNOWN, nil
	case "已停止":
		return DISCOUNT_STATUS__STOP, nil
	case "进行中":
		return DISCOUNT_STATUS__PROCESS, nil
	case "就绪":
		return DISCOUNT_STATUS__READY, nil
	}
	return DISCOUNT_STATUS_UNKNOWN, InvalidDiscountStatus
}

func (DiscountStatus) EnumType() string {
	return "DiscountStatus"
}

func (DiscountStatus) Enums() map[int][]string {
	return map[int][]string{
		int(DISCOUNT_STATUS__STOP):    {"STOP", "已停止"},
		int(DISCOUNT_STATUS__PROCESS): {"PROCESS", "进行中"},
		int(DISCOUNT_STATUS__READY):   {"READY", "就绪"},
	}
}

func (v DiscountStatus) String() string {
	switch v {
	case DISCOUNT_STATUS_UNKNOWN:
		return ""
	case DISCOUNT_STATUS__STOP:
		return "STOP"
	case DISCOUNT_STATUS__PROCESS:
		return "PROCESS"
	case DISCOUNT_STATUS__READY:
		return "READY"
	}
	return "UNKNOWN"
}

func (v DiscountStatus) Label() string {
	switch v {
	case DISCOUNT_STATUS_UNKNOWN:
		return ""
	case DISCOUNT_STATUS__STOP:
		return "已停止"
	case DISCOUNT_STATUS__PROCESS:
		return "进行中"
	case DISCOUNT_STATUS__READY:
		return "就绪"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*DiscountStatus)(nil)

func (v DiscountStatus) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidDiscountStatus
	}
	return []byte(str), nil
}

func (v *DiscountStatus) UnmarshalText(data []byte) (err error) {
	*v, err = ParseDiscountStatusFromString(string(bytes.ToUpper(data)))
	return
}
