package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidDiscountCal = errors.New("invalid DiscountCal")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("DiscountCal", map[string]string{
		"MULTISTEP_UNIT": "总价阶梯式单价立减",
		"MULTISTEP":      "总价阶梯式立减",
		"UNIT":           "单价立减",
	})
}

func ParseDiscountCalFromString(s string) (DiscountCal, error) {
	switch s {
	case "":
		return DISCOUNT_CAL_UNKNOWN, nil
	case "MULTISTEP_UNIT":
		return DISCOUNT_CAL__MULTISTEP_UNIT, nil
	case "MULTISTEP":
		return DISCOUNT_CAL__MULTISTEP, nil
	case "UNIT":
		return DISCOUNT_CAL__UNIT, nil
	}
	return DISCOUNT_CAL_UNKNOWN, InvalidDiscountCal
}

func ParseDiscountCalFromLabelString(s string) (DiscountCal, error) {
	switch s {
	case "":
		return DISCOUNT_CAL_UNKNOWN, nil
	case "总价阶梯式单价立减":
		return DISCOUNT_CAL__MULTISTEP_UNIT, nil
	case "总价阶梯式立减":
		return DISCOUNT_CAL__MULTISTEP, nil
	case "单价立减":
		return DISCOUNT_CAL__UNIT, nil
	}
	return DISCOUNT_CAL_UNKNOWN, InvalidDiscountCal
}

func (DiscountCal) EnumType() string {
	return "DiscountCal"
}

func (DiscountCal) Enums() map[int][]string {
	return map[int][]string{
		int(DISCOUNT_CAL__MULTISTEP_UNIT): {"MULTISTEP_UNIT", "总价阶梯式单价立减"},
		int(DISCOUNT_CAL__MULTISTEP):      {"MULTISTEP", "总价阶梯式立减"},
		int(DISCOUNT_CAL__UNIT):           {"UNIT", "单价立减"},
	}
}

func (v DiscountCal) String() string {
	switch v {
	case DISCOUNT_CAL_UNKNOWN:
		return ""
	case DISCOUNT_CAL__MULTISTEP_UNIT:
		return "MULTISTEP_UNIT"
	case DISCOUNT_CAL__MULTISTEP:
		return "MULTISTEP"
	case DISCOUNT_CAL__UNIT:
		return "UNIT"
	}
	return "UNKNOWN"
}

func (v DiscountCal) Label() string {
	switch v {
	case DISCOUNT_CAL_UNKNOWN:
		return ""
	case DISCOUNT_CAL__MULTISTEP_UNIT:
		return "总价阶梯式单价立减"
	case DISCOUNT_CAL__MULTISTEP:
		return "总价阶梯式立减"
	case DISCOUNT_CAL__UNIT:
		return "单价立减"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*DiscountCal)(nil)

func (v DiscountCal) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidDiscountCal
	}
	return []byte(str), nil
}

func (v *DiscountCal) UnmarshalText(data []byte) (err error) {
	*v, err = ParseDiscountCalFromString(string(bytes.ToUpper(data)))
	return
}
