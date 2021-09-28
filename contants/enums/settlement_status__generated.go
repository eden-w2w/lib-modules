package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidSettlementStatus = errors.New("invalid SettlementStatus")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("SettlementStatus", map[string]string{
		"COMPLETE": "已结算",
		"CREATED":  "待结算",
	})
}

func ParseSettlementStatusFromString(s string) (SettlementStatus, error) {
	switch s {
	case "":
		return SETTLEMENT_STATUS_UNKNOWN, nil
	case "COMPLETE":
		return SETTLEMENT_STATUS__COMPLETE, nil
	case "CREATED":
		return SETTLEMENT_STATUS__CREATED, nil
	}
	return SETTLEMENT_STATUS_UNKNOWN, InvalidSettlementStatus
}

func ParseSettlementStatusFromLabelString(s string) (SettlementStatus, error) {
	switch s {
	case "":
		return SETTLEMENT_STATUS_UNKNOWN, nil
	case "已结算":
		return SETTLEMENT_STATUS__COMPLETE, nil
	case "待结算":
		return SETTLEMENT_STATUS__CREATED, nil
	}
	return SETTLEMENT_STATUS_UNKNOWN, InvalidSettlementStatus
}

func (SettlementStatus) EnumType() string {
	return "SettlementStatus"
}

func (SettlementStatus) Enums() map[int][]string {
	return map[int][]string{
		int(SETTLEMENT_STATUS__COMPLETE): {"COMPLETE", "已结算"},
		int(SETTLEMENT_STATUS__CREATED):  {"CREATED", "待结算"},
	}
}

func (v SettlementStatus) String() string {
	switch v {
	case SETTLEMENT_STATUS_UNKNOWN:
		return ""
	case SETTLEMENT_STATUS__COMPLETE:
		return "COMPLETE"
	case SETTLEMENT_STATUS__CREATED:
		return "CREATED"
	}
	return "UNKNOWN"
}

func (v SettlementStatus) Label() string {
	switch v {
	case SETTLEMENT_STATUS_UNKNOWN:
		return ""
	case SETTLEMENT_STATUS__COMPLETE:
		return "已结算"
	case SETTLEMENT_STATUS__CREATED:
		return "待结算"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*SettlementStatus)(nil)

func (v SettlementStatus) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidSettlementStatus
	}
	return []byte(str), nil
}

func (v *SettlementStatus) UnmarshalText(data []byte) (err error) {
	*v, err = ParseSettlementStatusFromString(string(bytes.ToUpper(data)))
	return
}
