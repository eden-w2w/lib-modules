package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/contants/enums"
)

//go:generate eden generate model Order --database Config.DB --with-comments
//go:generate eden generate tag Order --defaults=true
// @def primary ID
// @def unique_index U_order_id OrderID
// @def index I_index UserID Status
// @def index I_expire ExpiredAt
type Order struct {
	datatypes.PrimaryID
	// 业务ID
	OrderID uint64 `json:"orderID,string" db:"f_order_id"`
	// 用户ID
	UserID uint64 `json:"userID,string" db:"f_user_id"`
	// 推荐人ID
	RefererID uint64 `json:"refererID,string" db:"f_referer_id,default='0'"`
	// 订单总额
	TotalPrice uint64 `json:"totalPrice" db:"f_total_price"`
	// 支付方式
	PaymentMethod enums.PaymentMethod `json:"paymentMethod" db:"f_payment_method"`
	// 备注
	Remark string `json:"remark" db:"f_remark,default='',size=1024"`
	// 收件人
	Recipients string `json:"recipients" db:"f_recipients"`
	// 收货地址
	ShippingAddr string `json:"shippingAddr" db:"f_shipping_addr"`
	// 联系电话
	Mobile string `json:"mobile" db:"f_mobile"`
	// 订单状态
	Status enums.OrderStatus `json:"status" db:"f_status"`
	// 过期时间
	ExpiredAt datatypes.MySQLTimestamp `db:"f_expired_at" json:"expiredAt"`
	datatypes.OperateTime
}
