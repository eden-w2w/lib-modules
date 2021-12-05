package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidAreaLevel = errors.New("invalid AreaLevel")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("AreaLevel", map[string]string{
		"STREET":   "街道级",
		"DISTRICT": "区县级",
		"CITY":     "市级",
		"PROVINCE": "省级",
		"COUNTRY":  "国家级",
	})
}

func ParseAreaLevelFromString(s string) (AreaLevel, error) {
	switch s {
	case "":
		return AREA_LEVEL_UNKNOWN, nil
	case "STREET":
		return AREA_LEVEL__STREET, nil
	case "DISTRICT":
		return AREA_LEVEL__DISTRICT, nil
	case "CITY":
		return AREA_LEVEL__CITY, nil
	case "PROVINCE":
		return AREA_LEVEL__PROVINCE, nil
	case "COUNTRY":
		return AREA_LEVEL__COUNTRY, nil
	}
	return AREA_LEVEL_UNKNOWN, InvalidAreaLevel
}

func ParseAreaLevelFromLabelString(s string) (AreaLevel, error) {
	switch s {
	case "":
		return AREA_LEVEL_UNKNOWN, nil
	case "街道级":
		return AREA_LEVEL__STREET, nil
	case "区县级":
		return AREA_LEVEL__DISTRICT, nil
	case "市级":
		return AREA_LEVEL__CITY, nil
	case "省级":
		return AREA_LEVEL__PROVINCE, nil
	case "国家级":
		return AREA_LEVEL__COUNTRY, nil
	}
	return AREA_LEVEL_UNKNOWN, InvalidAreaLevel
}

func (AreaLevel) EnumType() string {
	return "AreaLevel"
}

func (AreaLevel) Enums() map[int][]string {
	return map[int][]string{
		int(AREA_LEVEL__STREET):   {"STREET", "街道级"},
		int(AREA_LEVEL__DISTRICT): {"DISTRICT", "区县级"},
		int(AREA_LEVEL__CITY):     {"CITY", "市级"},
		int(AREA_LEVEL__PROVINCE): {"PROVINCE", "省级"},
		int(AREA_LEVEL__COUNTRY):  {"COUNTRY", "国家级"},
	}
}

func (v AreaLevel) String() string {
	switch v {
	case AREA_LEVEL_UNKNOWN:
		return ""
	case AREA_LEVEL__STREET:
		return "STREET"
	case AREA_LEVEL__DISTRICT:
		return "DISTRICT"
	case AREA_LEVEL__CITY:
		return "CITY"
	case AREA_LEVEL__PROVINCE:
		return "PROVINCE"
	case AREA_LEVEL__COUNTRY:
		return "COUNTRY"
	}
	return "UNKNOWN"
}

func (v AreaLevel) Label() string {
	switch v {
	case AREA_LEVEL_UNKNOWN:
		return ""
	case AREA_LEVEL__STREET:
		return "街道级"
	case AREA_LEVEL__DISTRICT:
		return "区县级"
	case AREA_LEVEL__CITY:
		return "市级"
	case AREA_LEVEL__PROVINCE:
		return "省级"
	case AREA_LEVEL__COUNTRY:
		return "国家级"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*AreaLevel)(nil)

func (v AreaLevel) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidAreaLevel
	}
	return []byte(str), nil
}

func (v *AreaLevel) UnmarshalText(data []byte) (err error) {
	*v, err = ParseAreaLevelFromString(string(bytes.ToUpper(data)))
	return
}
