package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/contants/types"
)

//go:generate eden generate model OrderGoods --database Config.DB --with-comments
//go:generate eden generate tag OrderGoods --defaults=true
// @def primary ID
// @def unique_index U_order_goods_id OrderID GoodsID
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
	Comment string `json:"comment" db:"f_comment"`
	// 发货地
	DispatchAddr string `json:"dispatchAddr" db:"f_dispatch_addr"`
	// 销量
	Sales int `json:"sales" db:"f_sales"`
	// 标题图片
	MainPicture string `json:"mainPicture" db:"f_main_picture,size=1024"`
	// 所有展示图片
	Pictures types.GoodsPictures `json:"pictures" db:"f_pictures,size=65535"`
	// 规格
	Specifications types.JsonArrayString `json:"specifications" db:"f_specification,size=1024"`
	// 活动
	Activities types.JsonArrayString `json:"activities" db:"f_activities"`
	// 物流政策
	LogisticPolicy string `json:"logisticPolicy" db:"f_logistic_policy,size=512"`
	// 价格
	Price uint64 `json:"price" db:"f_price"`
	// 库存
	Inventory uint64 `json:"inventory" db:"f_inventory"`
	// 详细介绍
	Detail string `json:"detail" db:"f_detail,size=65535"`
	// -------------------------------------------
	// 数量
	Amount uint32 `json:"amount" db:"f_amount"`

	datatypes.OperateTime
}
