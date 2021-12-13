package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
)

//go:generate eden generate model Settings --database Config.DB --with-comments
//go:generate eden generate tag Settings --defaults=true
// @def primary ID
// @def unique_index U_settings_id SettingsID
type Settings struct {
	datatypes.PrimaryID
	// 业务ID
	SettingsID uint64 `json:"settingsID,string" db:"f_settings_id"`
	// 推荐有礼分享标题图片
	PromotionMainPicture string `json:"promotionMainPicture" db:"f_promotion_main_picture,default=''"`
	// 消息推送
	// 是否启用消息推送
	WebhookEnabled datatypes.Bool `json:"webhookEnabled" db:"f_webhook_enabled,default=2"`
	// 地址
	WebhookURL string `json:"webhookURL" db:"f_webhook_url,default=''"`

	datatypes.OperateTime
}
