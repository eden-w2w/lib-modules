package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidPaymentMethod = errors.New("invalid PaymentMethod")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("PaymentMethod", map[string]string{
		"WECHAT": "微信支付",
	})
}

func ParsePaymentMethodFromString(s string) (PaymentMethod, error) {
	switch s {
	case "":
		return PAYMENT_METHOD_UNKNOWN, nil
	case "WECHAT":
		return PAYMENT_METHOD__WECHAT, nil
	}
	return PAYMENT_METHOD_UNKNOWN, InvalidPaymentMethod
}

func ParsePaymentMethodFromLabelString(s string) (PaymentMethod, error) {
	switch s {
	case "":
		return PAYMENT_METHOD_UNKNOWN, nil
	case "微信支付":
		return PAYMENT_METHOD__WECHAT, nil
	}
	return PAYMENT_METHOD_UNKNOWN, InvalidPaymentMethod
}

func (PaymentMethod) EnumType() string {
	return "PaymentMethod"
}

func (PaymentMethod) Enums() map[int][]string {
	return map[int][]string{
		int(PAYMENT_METHOD__WECHAT): {"WECHAT", "微信支付"},
	}
}

func (v PaymentMethod) String() string {
	switch v {
	case PAYMENT_METHOD_UNKNOWN:
		return ""
	case PAYMENT_METHOD__WECHAT:
		return "WECHAT"
	}
	return "UNKNOWN"
}

func (v PaymentMethod) Label() string {
	switch v {
	case PAYMENT_METHOD_UNKNOWN:
		return ""
	case PAYMENT_METHOD__WECHAT:
		return "微信支付"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*PaymentMethod)(nil)

func (v PaymentMethod) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidPaymentMethod
	}
	return []byte(str), nil
}

func (v *PaymentMethod) UnmarshalText(data []byte) (err error) {
	*v, err = ParsePaymentMethodFromString(string(bytes.ToUpper(data)))
	return
}
