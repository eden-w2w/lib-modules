package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidWechatTradeState = errors.New("invalid WechatTradeState")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("WechatTradeState", map[string]string{
		"PAYERROR":   "支付失败",
		"USERPAYING": "支付中",
		"REVOKED":    "已撤销",
		"CLOSED":     "已关闭",
		"NOTPAY":     "未支付",
		"REFUND":     "转入退款",
		"SUCCESS":    "支付成功",
	})
}

func ParseWechatTradeStateFromString(s string) (WechatTradeState, error) {
	switch s {
	case "":
		return WECHAT_TRADE_STATE_UNKNOWN, nil
	case "PAYERROR":
		return WECHAT_TRADE_STATE__PAYERROR, nil
	case "USERPAYING":
		return WECHAT_TRADE_STATE__USERPAYING, nil
	case "REVOKED":
		return WECHAT_TRADE_STATE__REVOKED, nil
	case "CLOSED":
		return WECHAT_TRADE_STATE__CLOSED, nil
	case "NOTPAY":
		return WECHAT_TRADE_STATE__NOTPAY, nil
	case "REFUND":
		return WECHAT_TRADE_STATE__REFUND, nil
	case "SUCCESS":
		return WECHAT_TRADE_STATE__SUCCESS, nil
	}
	return WECHAT_TRADE_STATE_UNKNOWN, InvalidWechatTradeState
}

func ParseWechatTradeStateFromLabelString(s string) (WechatTradeState, error) {
	switch s {
	case "":
		return WECHAT_TRADE_STATE_UNKNOWN, nil
	case "支付失败":
		return WECHAT_TRADE_STATE__PAYERROR, nil
	case "支付中":
		return WECHAT_TRADE_STATE__USERPAYING, nil
	case "已撤销":
		return WECHAT_TRADE_STATE__REVOKED, nil
	case "已关闭":
		return WECHAT_TRADE_STATE__CLOSED, nil
	case "未支付":
		return WECHAT_TRADE_STATE__NOTPAY, nil
	case "转入退款":
		return WECHAT_TRADE_STATE__REFUND, nil
	case "支付成功":
		return WECHAT_TRADE_STATE__SUCCESS, nil
	}
	return WECHAT_TRADE_STATE_UNKNOWN, InvalidWechatTradeState
}

func (WechatTradeState) EnumType() string {
	return "WechatTradeState"
}

func (WechatTradeState) Enums() map[int][]string {
	return map[int][]string{
		int(WECHAT_TRADE_STATE__PAYERROR):   {"PAYERROR", "支付失败"},
		int(WECHAT_TRADE_STATE__USERPAYING): {"USERPAYING", "支付中"},
		int(WECHAT_TRADE_STATE__REVOKED):    {"REVOKED", "已撤销"},
		int(WECHAT_TRADE_STATE__CLOSED):     {"CLOSED", "已关闭"},
		int(WECHAT_TRADE_STATE__NOTPAY):     {"NOTPAY", "未支付"},
		int(WECHAT_TRADE_STATE__REFUND):     {"REFUND", "转入退款"},
		int(WECHAT_TRADE_STATE__SUCCESS):    {"SUCCESS", "支付成功"},
	}
}

func (v WechatTradeState) String() string {
	switch v {
	case WECHAT_TRADE_STATE_UNKNOWN:
		return ""
	case WECHAT_TRADE_STATE__PAYERROR:
		return "PAYERROR"
	case WECHAT_TRADE_STATE__USERPAYING:
		return "USERPAYING"
	case WECHAT_TRADE_STATE__REVOKED:
		return "REVOKED"
	case WECHAT_TRADE_STATE__CLOSED:
		return "CLOSED"
	case WECHAT_TRADE_STATE__NOTPAY:
		return "NOTPAY"
	case WECHAT_TRADE_STATE__REFUND:
		return "REFUND"
	case WECHAT_TRADE_STATE__SUCCESS:
		return "SUCCESS"
	}
	return "UNKNOWN"
}

func (v WechatTradeState) Label() string {
	switch v {
	case WECHAT_TRADE_STATE_UNKNOWN:
		return ""
	case WECHAT_TRADE_STATE__PAYERROR:
		return "支付失败"
	case WECHAT_TRADE_STATE__USERPAYING:
		return "支付中"
	case WECHAT_TRADE_STATE__REVOKED:
		return "已撤销"
	case WECHAT_TRADE_STATE__CLOSED:
		return "已关闭"
	case WECHAT_TRADE_STATE__NOTPAY:
		return "未支付"
	case WECHAT_TRADE_STATE__REFUND:
		return "转入退款"
	case WECHAT_TRADE_STATE__SUCCESS:
		return "支付成功"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*WechatTradeState)(nil)

func (v WechatTradeState) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidWechatTradeState
	}
	return []byte(str), nil
}

func (v *WechatTradeState) UnmarshalText(data []byte) (err error) {
	*v, err = ParseWechatTradeStateFromString(string(bytes.ToUpper(data)))
	return
}
