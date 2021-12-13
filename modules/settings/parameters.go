package settings

import (
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/databases"
)

type UpdateSettingParams struct {
	PromotionMainPicture string `json:"promotionMainPicture" default:""`
	// 消息推送
	// 是否启用消息推送
	WebhookEnabled datatypes.Bool `json:"webhookEnabled" default:""`
	// 地址
	WebhookURL string `json:"webhookURL" default:""`
}

func (p UpdateSettingParams) Fill(model *databases.Settings) {
	if p.PromotionMainPicture != "" {
		model.PromotionMainPicture = p.PromotionMainPicture
	}
	if p.WebhookEnabled != datatypes.BOOL_UNKNOWN {
		model.WebhookEnabled = p.WebhookEnabled
	}
	if p.WebhookURL != "" {
		model.WebhookURL = p.WebhookURL
	}
}
