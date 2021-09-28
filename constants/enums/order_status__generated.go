package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidOrderStatus = errors.New("invalid OrderStatus")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("OrderStatus", map[string]string{
		"CLOSED":   "已关闭",
		"REFUND":   "退款中",
		"COMPLETE": "已完成",
		"DISPATCH": "已发货",
		"CONFIRM":  "待发货",
		"PAID":     "已支付",
		"CREATED":  "待支付",
	})
}

func ParseOrderStatusFromString(s string) (OrderStatus, error) {
	switch s {
	case "":
		return ORDER_STATUS_UNKNOWN, nil
	case "CLOSED":
		return ORDER_STATUS__CLOSED, nil
	case "REFUND":
		return ORDER_STATUS__REFUND, nil
	case "COMPLETE":
		return ORDER_STATUS__COMPLETE, nil
	case "DISPATCH":
		return ORDER_STATUS__DISPATCH, nil
	case "CONFIRM":
		return ORDER_STATUS__CONFIRM, nil
	case "PAID":
		return ORDER_STATUS__PAID, nil
	case "CREATED":
		return ORDER_STATUS__CREATED, nil
	}
	return ORDER_STATUS_UNKNOWN, InvalidOrderStatus
}

func ParseOrderStatusFromLabelString(s string) (OrderStatus, error) {
	switch s {
	case "":
		return ORDER_STATUS_UNKNOWN, nil
	case "已关闭":
		return ORDER_STATUS__CLOSED, nil
	case "退款中":
		return ORDER_STATUS__REFUND, nil
	case "已完成":
		return ORDER_STATUS__COMPLETE, nil
	case "已发货":
		return ORDER_STATUS__DISPATCH, nil
	case "待发货":
		return ORDER_STATUS__CONFIRM, nil
	case "已支付":
		return ORDER_STATUS__PAID, nil
	case "待支付":
		return ORDER_STATUS__CREATED, nil
	}
	return ORDER_STATUS_UNKNOWN, InvalidOrderStatus
}

func (OrderStatus) EnumType() string {
	return "OrderStatus"
}

func (OrderStatus) Enums() map[int][]string {
	return map[int][]string{
		int(ORDER_STATUS__CLOSED):   {"CLOSED", "已关闭"},
		int(ORDER_STATUS__REFUND):   {"REFUND", "退款中"},
		int(ORDER_STATUS__COMPLETE): {"COMPLETE", "已完成"},
		int(ORDER_STATUS__DISPATCH): {"DISPATCH", "已发货"},
		int(ORDER_STATUS__CONFIRM):  {"CONFIRM", "待发货"},
		int(ORDER_STATUS__PAID):     {"PAID", "已支付"},
		int(ORDER_STATUS__CREATED):  {"CREATED", "待支付"},
	}
}

func (v OrderStatus) String() string {
	switch v {
	case ORDER_STATUS_UNKNOWN:
		return ""
	case ORDER_STATUS__CLOSED:
		return "CLOSED"
	case ORDER_STATUS__REFUND:
		return "REFUND"
	case ORDER_STATUS__COMPLETE:
		return "COMPLETE"
	case ORDER_STATUS__DISPATCH:
		return "DISPATCH"
	case ORDER_STATUS__CONFIRM:
		return "CONFIRM"
	case ORDER_STATUS__PAID:
		return "PAID"
	case ORDER_STATUS__CREATED:
		return "CREATED"
	}
	return "UNKNOWN"
}

func (v OrderStatus) Label() string {
	switch v {
	case ORDER_STATUS_UNKNOWN:
		return ""
	case ORDER_STATUS__CLOSED:
		return "已关闭"
	case ORDER_STATUS__REFUND:
		return "退款中"
	case ORDER_STATUS__COMPLETE:
		return "已完成"
	case ORDER_STATUS__DISPATCH:
		return "已发货"
	case ORDER_STATUS__CONFIRM:
		return "待发货"
	case ORDER_STATUS__PAID:
		return "已支付"
	case ORDER_STATUS__CREATED:
		return "待支付"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*OrderStatus)(nil)

func (v OrderStatus) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidOrderStatus
	}
	return []byte(str), nil
}

func (v *OrderStatus) UnmarshalText(data []byte) (err error) {
	*v, err = ParseOrderStatusFromString(string(bytes.ToUpper(data)))
	return
}
