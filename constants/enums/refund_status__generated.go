package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidRefundStatus = errors.New("invalid RefundStatus")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("RefundStatus", map[string]string{
		"ABNORMAL":   "退款异常",
		"PROCESSING": "退款处理中",
		"CLOSED":     "退款关闭",
		"SUCCESS":    "退款成功",
	})
}

func ParseRefundStatusFromString(s string) (RefundStatus, error) {
	switch s {
	case "":
		return REFUND_STATUS_UNKNOWN, nil
	case "ABNORMAL":
		return REFUND_STATUS__ABNORMAL, nil
	case "PROCESSING":
		return REFUND_STATUS__PROCESSING, nil
	case "CLOSED":
		return REFUND_STATUS__CLOSED, nil
	case "SUCCESS":
		return REFUND_STATUS__SUCCESS, nil
	}
	return REFUND_STATUS_UNKNOWN, InvalidRefundStatus
}

func ParseRefundStatusFromLabelString(s string) (RefundStatus, error) {
	switch s {
	case "":
		return REFUND_STATUS_UNKNOWN, nil
	case "退款异常":
		return REFUND_STATUS__ABNORMAL, nil
	case "退款处理中":
		return REFUND_STATUS__PROCESSING, nil
	case "退款关闭":
		return REFUND_STATUS__CLOSED, nil
	case "退款成功":
		return REFUND_STATUS__SUCCESS, nil
	}
	return REFUND_STATUS_UNKNOWN, InvalidRefundStatus
}

func (RefundStatus) EnumType() string {
	return "RefundStatus"
}

func (RefundStatus) Enums() map[int][]string {
	return map[int][]string{
		int(REFUND_STATUS__ABNORMAL):   {"ABNORMAL", "退款异常"},
		int(REFUND_STATUS__PROCESSING): {"PROCESSING", "退款处理中"},
		int(REFUND_STATUS__CLOSED):     {"CLOSED", "退款关闭"},
		int(REFUND_STATUS__SUCCESS):    {"SUCCESS", "退款成功"},
	}
}

func (v RefundStatus) String() string {
	switch v {
	case REFUND_STATUS_UNKNOWN:
		return ""
	case REFUND_STATUS__ABNORMAL:
		return "ABNORMAL"
	case REFUND_STATUS__PROCESSING:
		return "PROCESSING"
	case REFUND_STATUS__CLOSED:
		return "CLOSED"
	case REFUND_STATUS__SUCCESS:
		return "SUCCESS"
	}
	return "UNKNOWN"
}

func (v RefundStatus) Label() string {
	switch v {
	case REFUND_STATUS_UNKNOWN:
		return ""
	case REFUND_STATUS__ABNORMAL:
		return "退款异常"
	case REFUND_STATUS__PROCESSING:
		return "退款处理中"
	case REFUND_STATUS__CLOSED:
		return "退款关闭"
	case REFUND_STATUS__SUCCESS:
		return "退款成功"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*RefundStatus)(nil)

func (v RefundStatus) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidRefundStatus
	}
	return []byte(str), nil
}

func (v *RefundStatus) UnmarshalText(data []byte) (err error) {
	*v, err = ParseRefundStatusFromString(string(bytes.ToUpper(data)))
	return
}
