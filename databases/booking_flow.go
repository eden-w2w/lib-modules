package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
)

//go:generate eden generate model BookingFlow --database Config.DB --with-comments
//go:generate eden generate tag BookingFlow --defaults=true
// @def primary ID
// @def unique_index U_booking_flow_id FlowID
// @def index I_default GoodsID Status Type
// @def index I_time StartTime EndTime
type BookingFlow struct {
	datatypes.PrimaryID
	// 业务ID
	FlowID uint64 `json:"flowID,string" db:"f_flow_id"`
	// 商品ID
	GoodsID uint64 `json:"goodsID,string" db:"f_goods_id"`
	// 预售销量
	Sales uint32 `json:"sales" db:"f_sales,default=0"`
	// 预售限量
	Limit uint32 `json:"limit" db:"f_limit,default=0"`
	// 预售模式
	Type enums.BookingType `json:"type" db:"f_type"`
	// 预售状态
	Status enums.BookingStatus `json:"status" db:"f_status"`
	// 预售开始时间
	StartTime datatypes.MySQLTimestamp `json:"startTime" db:"f_start_time"`
	// 预售结束时间
	EndTime datatypes.MySQLTimestamp `json:"endTime" db:"f_end_time,null"`
	// 预计到货时间
	EstimatedTimeArrival datatypes.MySQLTimestamp `json:"eta" db:"f_eta,null"`

	datatypes.OperateTime
}
