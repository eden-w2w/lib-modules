package webhook

import "github.com/eden-w2w/lib-modules/constants/enums"

type SendMessageRequest struct {
	Key  string          `in:"query" name:"key"`
	Data SendMessageBody `in:"body"`
}

func (s *SendMessageRequest) SetKey(key string) {
	s.Key = key
}

type SendMessageBody struct {
	MsgType  enums.WebhookMsgType `json:"msgtype"`
	Text     *SendMessageText     `json:"text" default:""`
	Markdown *SendMessageMarkdown `json:"markdown" default:""`
}

type SendMessageText struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list" default:""`
	MentionedMobileList []string `json:"mentioned_mobile_list" default:""`
}

type SendMessageMarkdown struct {
	Content string `json:"content"`
}
