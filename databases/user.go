package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
)

//go:generate eden generate model User --database Config.DB --with-comments
//go:generate eden generate tag User --defaults=true
// @def primary ID
// @def unique_index U_user_id UserID
// @def unique_index U_open_id OpenID
// @def index I_union_id UnionID
// @def unique_index U_token Token
type User struct {
	datatypes.PrimaryID
	// 业务ID
	UserID uint64 `json:"userID,string" db:"f_user_id"`
	// 用户名
	UserName string `json:"userName" db:"f_user_name,default=''"`
	// 手机号
	Mobile string `json:"mobile" db:"f_mobile,default=''"`
	// 昵称
	NickName string `json:"nickName" db:"f_nick_name,default=''"`
	// 头像地址
	AvatarUrl string `json:"avatarUrl" db:"f_avatar_url,size=1024,default=''"`
	// 推荐人ID
	RefererID uint64 `json:"refererID,string" db:"f_referer_id,default='0'"`
	// 访问令牌
	Token string `json:"token" db:"f_token,default=''"`
	// 微信OpenID
	OpenID string `json:"openID" db:"f_open_id,default=''"`
	// 微信UnionID
	UnionID string `json:"unionID" db:"f_union_id,default=''"`
	// 微信SessionKey
	SessionKey string `json:"sessionKey" db:"f_session_key,default=''"`

	datatypes.OperateTime
}
