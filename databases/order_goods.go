package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/types"
)

//go:generate eden generate model OrderGoods --database Config.DB --with-comments
//go:generate eden generate tag OrderGoods --defaults=true
// @def primary ID
// @def unique_index U_order_goods_id OrderID GoodsID
// @def index I_booking BookingFlowID
type OrderGoods struct {
	datatypes.PrimaryID
	// 订单ID
	OrderID uint64 `json:"orderID,string" db:"f_order_id"`
	// 商品ID
	GoodsID uint64 `json:"goodsID,string" db:"f_goods_id"`
	// ---------------- 商品快照 ------------------
	// 名称
	Name string `json:"name" db:"f_name"`
	// 描述
	Comment string `json:"comment" db:"f_comment,default=''"`
	// 运费模板
	FreightTemplateID uint64 `json:"freightTemplateID,string" db:"f_freight_template_id"`
	// 销量
	Sales uint32 `json:"sales" db:"f_sales,default=0"`
	// 标题图片
	MainPicture string `json:"mainPicture" db:"f_main_picture,size=1024"`
	// 所有展示图片
	Pictures types.GoodsPictures `json:"pictures" db:"f_pictures,size=65535"`
	// 规格
	Specifications types.JsonArrayString `json:"specifications" db:"f_specification,size=1024"`
	// 价格
	Price uint64 `json:"price" db:"f_price"`
	// 库存
	Inventory uint64 `json:"inventory" db:"f_inventory,default=0"`
	// 详细介绍
	Detail string `json:"detail" db:"f_detail,size=65535"`
	// -------------------------------------------
	// 数量
	Amount uint32 `json:"amount" db:"f_amount"`
	// 是否预订
	IsBooking datatypes.Bool `json:"isBooking" db:"f_is_booking,default=2"`
	// 预售单号
	BookingFlowID uint64 `json:"bookingFlowID,string" db:"f_booking_flow_id,null"`

	datatypes.OperateTime
}
