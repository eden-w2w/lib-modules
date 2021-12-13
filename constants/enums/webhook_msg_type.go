package enums

//go:generate eden generate enum --type-name=WebhookMsgType
// api:enum
type WebhookMsgType uint8

// webhook消息类型
const (
	WEBHOOK_MSG_TYPE_UNKNOWN   WebhookMsgType = iota
	WEBHOOK_MSG_TYPE__text                    // text
	WEBHOOK_MSG_TYPE__markdown                // markdown
)
