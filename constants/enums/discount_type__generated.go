package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidDiscountType = errors.New("invalid DiscountType")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("DiscountType", map[string]string{
		"ALL_PERCENT": "全场折扣",
		"ALL":         "全场立减",
	})
}

func ParseDiscountTypeFromString(s string) (DiscountType, error) {
	switch s {
	case "":
		return DISCOUNT_TYPE_UNKNOWN, nil
	case "ALL_PERCENT":
		return DISCOUNT_TYPE__ALL_PERCENT, nil
	case "ALL":
		return DISCOUNT_TYPE__ALL, nil
	}
	return DISCOUNT_TYPE_UNKNOWN, InvalidDiscountType
}

func ParseDiscountTypeFromLabelString(s string) (DiscountType, error) {
	switch s {
	case "":
		return DISCOUNT_TYPE_UNKNOWN, nil
	case "全场折扣":
		return DISCOUNT_TYPE__ALL_PERCENT, nil
	case "全场立减":
		return DISCOUNT_TYPE__ALL, nil
	}
	return DISCOUNT_TYPE_UNKNOWN, InvalidDiscountType
}

func (DiscountType) EnumType() string {
	return "DiscountType"
}

func (DiscountType) Enums() map[int][]string {
	return map[int][]string{
		int(DISCOUNT_TYPE__ALL_PERCENT): {"ALL_PERCENT", "全场折扣"},
		int(DISCOUNT_TYPE__ALL):         {"ALL", "全场立减"},
	}
}

func (v DiscountType) String() string {
	switch v {
	case DISCOUNT_TYPE_UNKNOWN:
		return ""
	case DISCOUNT_TYPE__ALL_PERCENT:
		return "ALL_PERCENT"
	case DISCOUNT_TYPE__ALL:
		return "ALL"
	}
	return "UNKNOWN"
}

func (v DiscountType) Label() string {
	switch v {
	case DISCOUNT_TYPE_UNKNOWN:
		return ""
	case DISCOUNT_TYPE__ALL_PERCENT:
		return "全场折扣"
	case DISCOUNT_TYPE__ALL:
		return "全场立减"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*DiscountType)(nil)

func (v DiscountType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidDiscountType
	}
	return []byte(str), nil
}

func (v *DiscountType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseDiscountTypeFromString(string(bytes.ToUpper(data)))
	return
}
