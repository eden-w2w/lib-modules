package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidWebhookMsgType = errors.New("invalid WebhookMsgType")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("WebhookMsgType", map[string]string{
		"markdown": "markdown",
		"text":     "text",
	})
}

func ParseWebhookMsgTypeFromString(s string) (WebhookMsgType, error) {
	switch s {
	case "":
		return WEBHOOK_MSG_TYPE_UNKNOWN, nil
	case "markdown":
		return WEBHOOK_MSG_TYPE__markdown, nil
	case "text":
		return WEBHOOK_MSG_TYPE__text, nil
	}
	return WEBHOOK_MSG_TYPE_UNKNOWN, InvalidWebhookMsgType
}

func ParseWebhookMsgTypeFromLabelString(s string) (WebhookMsgType, error) {
	switch s {
	case "":
		return WEBHOOK_MSG_TYPE_UNKNOWN, nil
	case "markdown":
		return WEBHOOK_MSG_TYPE__markdown, nil
	case "text":
		return WEBHOOK_MSG_TYPE__text, nil
	}
	return WEBHOOK_MSG_TYPE_UNKNOWN, InvalidWebhookMsgType
}

func (WebhookMsgType) EnumType() string {
	return "WebhookMsgType"
}

func (WebhookMsgType) Enums() map[int][]string {
	return map[int][]string{
		int(WEBHOOK_MSG_TYPE__markdown): {"markdown", "markdown"},
		int(WEBHOOK_MSG_TYPE__text):     {"text", "text"},
	}
}

func (v WebhookMsgType) String() string {
	switch v {
	case WEBHOOK_MSG_TYPE_UNKNOWN:
		return ""
	case WEBHOOK_MSG_TYPE__markdown:
		return "markdown"
	case WEBHOOK_MSG_TYPE__text:
		return "text"
	}
	return "UNKNOWN"
}

func (v WebhookMsgType) Label() string {
	switch v {
	case WEBHOOK_MSG_TYPE_UNKNOWN:
		return ""
	case WEBHOOK_MSG_TYPE__markdown:
		return "markdown"
	case WEBHOOK_MSG_TYPE__text:
		return "text"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*WebhookMsgType)(nil)

func (v WebhookMsgType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidWebhookMsgType
	}
	return []byte(str), nil
}

func (v *WebhookMsgType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseWebhookMsgTypeFromString(string(bytes.ToUpper(data)))
	return
}
