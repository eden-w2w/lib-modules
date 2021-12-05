package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
)

//go:generate eden generate model FreightTemplate --database Config.DB --with-comments
//go:generate eden generate tag FreightTemplate --defaults=true
// @def primary ID
// @def unique_index U_template_id TemplateID
// @def index I_search Name IsFreeFreight Cal
type FreightTemplate struct {
	datatypes.PrimaryID
	// 业务ID
	TemplateID uint64 `json:"templateID,string" db:"f_template_id"`
	// 名称
	Name string `json:"name" db:"f_name"`
	// 发货地
	DispatchAddr string `json:"dispatchAddr" db:"f_dispatch_addr"`
	// 发货时间
	DispatchTime uint32 `json:"dispatchTime" db:"f_dispatch_time,null"`
	// 是否全场包邮
	IsFreeFreight datatypes.Bool `json:"isFreeFreight" db:"f_is_free_freight"`
	// 计费方式
	Cal enums.FreightCal `json:"cal" db:"f_cal,null"`
	// -------------------------------------------------
	// 默认运费设置
	// 首重（克）/首件（个）范围
	FirstRange uint32 `json:"firstRange" db:"f_first_range,null"`
	// 首重/首件价格
	FirstPrice uint32 `json:"firstPrice" db:"f_first_price,null"`
	// 续重（克）/续件（个）范围
	ContinueRange uint32 `json:"continueRange" db:"f_continue_range,null"`
	// 续重/续件价格
	ContinuePrice uint32 `json:"continuePrice" db:"f_continue_price,null"`

	datatypes.OperateTime
}
