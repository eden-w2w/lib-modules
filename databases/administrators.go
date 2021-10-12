package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
)

//go:generate eden generate model Administrators --database Config.DB --with-comments
//go:generate eden generate tag Administrators --defaults=true
// @def primary ID
// @def unique_index U_administrators_id AdministratorsID
// @def unique_index U_token Token
// @def unique_index U_user_name UserName
// @def index I_expire ExpiredAt
type Administrators struct {
	datatypes.PrimaryID
	// 业务ID
	AdministratorsID uint64 `json:"administratorsID,string" db:"f_administrators_id"`
	// 用户名
	UserName string `json:"userName" db:"f_username"`
	// 密码
	Password string `json:"-" db:"f_password"`
	// 访问令牌
	Token string `json:"token" db:"f_token,default=''"`
	// 访问令牌过期时间
	ExpiredAt datatypes.MySQLTimestamp `json:"expiredAt" db:"f_expired_at,default='0'"`

	datatypes.OperateTime
}
