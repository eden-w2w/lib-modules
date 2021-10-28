package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
)

//go:generate eden generate model ShippingAddress --database Config.DB --with-comments
//go:generate eden generate tag ShippingAddress --defaults=true
// @def primary ID
// @def unique_index U_shipping_id ShippingID
// @def index I_user UserID
type ShippingAddress struct {
	datatypes.PrimaryID
	// 业务ID
	ShippingID uint64 `json:"shippingID,string" db:"f_shipping_id"`
	// 用户ID
	UserID uint64 `json:"userID,string" db:"f_user_id"`
	// 收件人
	Recipients string `json:"recipients" db:"f_recipients"`
	// 省市区街道
	District string `json:"district" db:"f_district"`
	// 详细地址
	Address string `json:"address" db:"f_address"`
	// 联系电话
	Mobile string `json:"mobile" db:"f_mobile"`

	datatypes.OperateTime
}
