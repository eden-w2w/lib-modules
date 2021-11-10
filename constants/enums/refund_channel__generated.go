package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidRefundChannel = errors.New("invalid RefundChannel")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("RefundChannel", map[string]string{
		"OTHER_BANKCARD": "原银行卡异常退到其他银行卡",
		"OTHER_BALANCE":  "原账户异常退到其他余额账户",
		"BALANCE":        "退款到余额",
		"ORIGINAL":       "原路退款",
	})
}

func ParseRefundChannelFromString(s string) (RefundChannel, error) {
	switch s {
	case "":
		return REFUND_CHANNEL_UNKNOWN, nil
	case "OTHER_BANKCARD":
		return REFUND_CHANNEL__OTHER_BANKCARD, nil
	case "OTHER_BALANCE":
		return REFUND_CHANNEL__OTHER_BALANCE, nil
	case "BALANCE":
		return REFUND_CHANNEL__BALANCE, nil
	case "ORIGINAL":
		return REFUND_CHANNEL__ORIGINAL, nil
	}
	return REFUND_CHANNEL_UNKNOWN, InvalidRefundChannel
}

func ParseRefundChannelFromLabelString(s string) (RefundChannel, error) {
	switch s {
	case "":
		return REFUND_CHANNEL_UNKNOWN, nil
	case "原银行卡异常退到其他银行卡":
		return REFUND_CHANNEL__OTHER_BANKCARD, nil
	case "原账户异常退到其他余额账户":
		return REFUND_CHANNEL__OTHER_BALANCE, nil
	case "退款到余额":
		return REFUND_CHANNEL__BALANCE, nil
	case "原路退款":
		return REFUND_CHANNEL__ORIGINAL, nil
	}
	return REFUND_CHANNEL_UNKNOWN, InvalidRefundChannel
}

func (RefundChannel) EnumType() string {
	return "RefundChannel"
}

func (RefundChannel) Enums() map[int][]string {
	return map[int][]string{
		int(REFUND_CHANNEL__OTHER_BANKCARD): {"OTHER_BANKCARD", "原银行卡异常退到其他银行卡"},
		int(REFUND_CHANNEL__OTHER_BALANCE):  {"OTHER_BALANCE", "原账户异常退到其他余额账户"},
		int(REFUND_CHANNEL__BALANCE):        {"BALANCE", "退款到余额"},
		int(REFUND_CHANNEL__ORIGINAL):       {"ORIGINAL", "原路退款"},
	}
}

func (v RefundChannel) String() string {
	switch v {
	case REFUND_CHANNEL_UNKNOWN:
		return ""
	case REFUND_CHANNEL__OTHER_BANKCARD:
		return "OTHER_BANKCARD"
	case REFUND_CHANNEL__OTHER_BALANCE:
		return "OTHER_BALANCE"
	case REFUND_CHANNEL__BALANCE:
		return "BALANCE"
	case REFUND_CHANNEL__ORIGINAL:
		return "ORIGINAL"
	}
	return "UNKNOWN"
}

func (v RefundChannel) Label() string {
	switch v {
	case REFUND_CHANNEL_UNKNOWN:
		return ""
	case REFUND_CHANNEL__OTHER_BANKCARD:
		return "原银行卡异常退到其他银行卡"
	case REFUND_CHANNEL__OTHER_BALANCE:
		return "原账户异常退到其他余额账户"
	case REFUND_CHANNEL__BALANCE:
		return "退款到余额"
	case REFUND_CHANNEL__ORIGINAL:
		return "原路退款"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*RefundChannel)(nil)

func (v RefundChannel) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidRefundChannel
	}
	return []byte(str), nil
}

func (v *RefundChannel) UnmarshalText(data []byte) (err error) {
	*v, err = ParseRefundChannelFromString(string(bytes.ToUpper(data)))
	return
}
