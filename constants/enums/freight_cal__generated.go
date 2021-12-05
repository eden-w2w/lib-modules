package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidFreightCal = errors.New("invalid FreightCal")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("FreightCal", map[string]string{
		"WEIGHT": "按重量",
		"COUNT":  "按数量",
	})
}

func ParseFreightCalFromString(s string) (FreightCal, error) {
	switch s {
	case "":
		return FREIGHT_CAL_UNKNOWN, nil
	case "WEIGHT":
		return FREIGHT_CAL__WEIGHT, nil
	case "COUNT":
		return FREIGHT_CAL__COUNT, nil
	}
	return FREIGHT_CAL_UNKNOWN, InvalidFreightCal
}

func ParseFreightCalFromLabelString(s string) (FreightCal, error) {
	switch s {
	case "":
		return FREIGHT_CAL_UNKNOWN, nil
	case "按重量":
		return FREIGHT_CAL__WEIGHT, nil
	case "按数量":
		return FREIGHT_CAL__COUNT, nil
	}
	return FREIGHT_CAL_UNKNOWN, InvalidFreightCal
}

func (FreightCal) EnumType() string {
	return "FreightCal"
}

func (FreightCal) Enums() map[int][]string {
	return map[int][]string{
		int(FREIGHT_CAL__WEIGHT): {"WEIGHT", "按重量"},
		int(FREIGHT_CAL__COUNT):  {"COUNT", "按数量"},
	}
}

func (v FreightCal) String() string {
	switch v {
	case FREIGHT_CAL_UNKNOWN:
		return ""
	case FREIGHT_CAL__WEIGHT:
		return "WEIGHT"
	case FREIGHT_CAL__COUNT:
		return "COUNT"
	}
	return "UNKNOWN"
}

func (v FreightCal) Label() string {
	switch v {
	case FREIGHT_CAL_UNKNOWN:
		return ""
	case FREIGHT_CAL__WEIGHT:
		return "按重量"
	case FREIGHT_CAL__COUNT:
		return "按数量"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*FreightCal)(nil)

func (v FreightCal) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidFreightCal
	}
	return []byte(str), nil
}

func (v *FreightCal) UnmarshalText(data []byte) (err error) {
	*v, err = ParseFreightCalFromString(string(bytes.ToUpper(data)))
	return
}
