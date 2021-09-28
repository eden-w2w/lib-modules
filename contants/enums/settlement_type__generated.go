package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidSettlementType = errors.New("invalid SettlementType")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("SettlementType", map[string]string{
		"MONTH": "月",
		"WEEK":  "周",
	})
}

func ParseSettlementTypeFromString(s string) (SettlementType, error) {
	switch s {
	case "":
		return SETTLEMENT_TYPE_UNKNOWN, nil
	case "MONTH":
		return SETTLEMENT_TYPE__MONTH, nil
	case "WEEK":
		return SETTLEMENT_TYPE__WEEK, nil
	}
	return SETTLEMENT_TYPE_UNKNOWN, InvalidSettlementType
}

func ParseSettlementTypeFromLabelString(s string) (SettlementType, error) {
	switch s {
	case "":
		return SETTLEMENT_TYPE_UNKNOWN, nil
	case "月":
		return SETTLEMENT_TYPE__MONTH, nil
	case "周":
		return SETTLEMENT_TYPE__WEEK, nil
	}
	return SETTLEMENT_TYPE_UNKNOWN, InvalidSettlementType
}

func (SettlementType) EnumType() string {
	return "SettlementType"
}

func (SettlementType) Enums() map[int][]string {
	return map[int][]string{
		int(SETTLEMENT_TYPE__MONTH): {"MONTH", "月"},
		int(SETTLEMENT_TYPE__WEEK):  {"WEEK", "周"},
	}
}

func (v SettlementType) String() string {
	switch v {
	case SETTLEMENT_TYPE_UNKNOWN:
		return ""
	case SETTLEMENT_TYPE__MONTH:
		return "MONTH"
	case SETTLEMENT_TYPE__WEEK:
		return "WEEK"
	}
	return "UNKNOWN"
}

func (v SettlementType) Label() string {
	switch v {
	case SETTLEMENT_TYPE_UNKNOWN:
		return ""
	case SETTLEMENT_TYPE__MONTH:
		return "月"
	case SETTLEMENT_TYPE__WEEK:
		return "周"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*SettlementType)(nil)

func (v SettlementType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidSettlementType
	}
	return []byte(str), nil
}

func (v *SettlementType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseSettlementTypeFromString(string(bytes.ToUpper(data)))
	return
}
