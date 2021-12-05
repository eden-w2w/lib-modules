package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
)

//go:generate eden generate model Order --database Config.DB --with-comments
//go:generate eden generate tag Order --defaults=true
// @def primary ID
// @def unique_index U_order_id OrderID
// @def index I_index UserID Status
// @def index I_expire Status ExpiredAt
type Order struct {
	datatypes.PrimaryID
	// 业务ID
	OrderID uint64 `json:"orderID,string" db:"f_order_id"`
	// 用户ID
	UserID uint64 `json:"userID,string" db:"f_user_id"`
	// 用户昵称
	NickName string `json:"nickName" db:"f_nick_name,default=''"`
	// 微信OpenID
	UserOpenID string `json:"userOpenID" db:"f_user_open_id,default=''"`
	// 订单总额
	TotalPrice uint64 `json:"totalPrice" db:"f_total_price"`
	// 优惠金额
	DiscountAmount uint64 `json:"discountAmount" db:"f_discount_amount,default=0"`
	// 运费
	FreightAmount uint64 `json:"freightAmount" db:"f_freight_amount,default=0"`
	// 实际金额
	ActualAmount uint64 `json:"actualAmount" db:"f_actual_amount"`
	// 支付方式
	PaymentMethod enums.PaymentMethod `json:"paymentMethod" db:"f_payment_method"`
	// 备注
	Remark string `json:"remark" db:"f_remark,default='',size=1024"`
	// 订单状态
	Status enums.OrderStatus `json:"status" db:"f_status"`
	// 过期时间
	ExpiredAt datatypes.MySQLTimestamp `db:"f_expired_at" json:"expiredAt"`
	datatypes.OperateTime
}
