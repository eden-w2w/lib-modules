package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidPaymentStatus = errors.New("invalid PaymentStatus")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("PaymentStatus", map[string]string{
		"CLOSED":  "已关闭",
		"FAIL":    "支付失败",
		"SUCCESS": "支付成功",
		"PROCESS": "处理中",
		"CREATED": "未支付",
	})
}

func ParsePaymentStatusFromString(s string) (PaymentStatus, error) {
	switch s {
	case "":
		return PAYMENT_STATUS_UNKNOWN, nil
	case "CLOSED":
		return PAYMENT_STATUS__CLOSED, nil
	case "FAIL":
		return PAYMENT_STATUS__FAIL, nil
	case "SUCCESS":
		return PAYMENT_STATUS__SUCCESS, nil
	case "PROCESS":
		return PAYMENT_STATUS__PROCESS, nil
	case "CREATED":
		return PAYMENT_STATUS__CREATED, nil
	}
	return PAYMENT_STATUS_UNKNOWN, InvalidPaymentStatus
}

func ParsePaymentStatusFromLabelString(s string) (PaymentStatus, error) {
	switch s {
	case "":
		return PAYMENT_STATUS_UNKNOWN, nil
	case "已关闭":
		return PAYMENT_STATUS__CLOSED, nil
	case "支付失败":
		return PAYMENT_STATUS__FAIL, nil
	case "支付成功":
		return PAYMENT_STATUS__SUCCESS, nil
	case "处理中":
		return PAYMENT_STATUS__PROCESS, nil
	case "未支付":
		return PAYMENT_STATUS__CREATED, nil
	}
	return PAYMENT_STATUS_UNKNOWN, InvalidPaymentStatus
}

func (PaymentStatus) EnumType() string {
	return "PaymentStatus"
}

func (PaymentStatus) Enums() map[int][]string {
	return map[int][]string{
		int(PAYMENT_STATUS__CLOSED):  {"CLOSED", "已关闭"},
		int(PAYMENT_STATUS__FAIL):    {"FAIL", "支付失败"},
		int(PAYMENT_STATUS__SUCCESS): {"SUCCESS", "支付成功"},
		int(PAYMENT_STATUS__PROCESS): {"PROCESS", "处理中"},
		int(PAYMENT_STATUS__CREATED): {"CREATED", "未支付"},
	}
}

func (v PaymentStatus) String() string {
	switch v {
	case PAYMENT_STATUS_UNKNOWN:
		return ""
	case PAYMENT_STATUS__CLOSED:
		return "CLOSED"
	case PAYMENT_STATUS__FAIL:
		return "FAIL"
	case PAYMENT_STATUS__SUCCESS:
		return "SUCCESS"
	case PAYMENT_STATUS__PROCESS:
		return "PROCESS"
	case PAYMENT_STATUS__CREATED:
		return "CREATED"
	}
	return "UNKNOWN"
}

func (v PaymentStatus) Label() string {
	switch v {
	case PAYMENT_STATUS_UNKNOWN:
		return ""
	case PAYMENT_STATUS__CLOSED:
		return "已关闭"
	case PAYMENT_STATUS__FAIL:
		return "支付失败"
	case PAYMENT_STATUS__SUCCESS:
		return "支付成功"
	case PAYMENT_STATUS__PROCESS:
		return "处理中"
	case PAYMENT_STATUS__CREATED:
		return "未支付"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*PaymentStatus)(nil)

func (v PaymentStatus) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidPaymentStatus
	}
	return []byte(str), nil
}

func (v *PaymentStatus) UnmarshalText(data []byte) (err error) {
	*v, err = ParsePaymentStatusFromString(string(bytes.ToUpper(data)))
	return
}
