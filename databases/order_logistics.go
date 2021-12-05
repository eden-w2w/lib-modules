package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
)

//go:generate eden generate model OrderLogistics --database Config.DB --with-comments
//go:generate eden generate tag OrderLogistics --defaults=true
// @def primary ID
// @def unique_index U_logistics_id LogisticsID
// @def unique_index U_order_id OrderID
// @def index I_number CourierNumber
type OrderLogistics struct {
	datatypes.PrimaryID
	// 业务ID
	LogisticsID uint64 `json:"logisticsID,string" db:"f_logistics_id"`
	// 订单号
	OrderID uint64 `json:"orderID,string" db:"f_order_id"`
	// 收件地址ID
	ShippingID uint64 `json:"shippingID,string" db:"f_shipping_id"`
	// 收件人
	Recipients string `json:"recipients" db:"f_recipients"`
	// 收货地址
	ShippingAddr string `json:"shippingAddr" db:"f_shipping_addr"`
	// 联系电话
	Mobile string `json:"mobile" db:"f_mobile"`
	// 快递公司
	CourierCompany string `json:"courierCompany" db:"f_courier_company,default=''"`
	// 快递单号
	CourierNumber string `json:"courierNumber" db:"f_courier_number,default=''"`

	datatypes.OperateTime
}
