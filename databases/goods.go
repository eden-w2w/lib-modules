package databases

import (
	"database/sql"
	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/types"
)

//go:generate eden generate model Goods --database Config.DB --with-comments
//go:generate eden generate tag Goods --defaults=true
// @def primary ID
// @def unique_index U_goods_id GoodsID
// @def index I_freight FreightTemplateID
type Goods struct {
	datatypes.PrimaryID
	// 业务ID
	GoodsID uint64 `json:"goodsID,string" db:"f_goods_id"`
	// 名称
	Name string `json:"name" db:"f_name"`
	// 描述
	Comment string `json:"comment" db:"f_comment,default=''"`
	// 运费模板
	FreightTemplateID uint64 `json:"freightTemplateID,string" db:"f_freight_template_id"`
	// 单品净重，单位克（用于邮费计算）
	UnitNetWeight uint32 `json:"unitNetWeight" db:"f_unit_net_weight,default=0"`
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
	// 是否开启无货后预订模式
	IsAllowBooking datatypes.Bool `json:"isAllowBooking" db:"f_is_allow_booking,default=0"`

	datatypes.OperateTime
}

func (m *Goods) MaxGoodsID(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (maxID uint64, err error) {
	table := db.T(m)

	id := sql.NullInt64{}

	condition = github_com_eden_framework_sqlx_builder.And(condition, m.FieldDeletedAt().Eq(0))
	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("Goods.MaxGoodsID"),
	}

	if len(additions) > 0 {
		finalAdditions = append(finalAdditions, additions...)
	}

	err = db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(
			github_com_eden_framework_sqlx_builder.Max(m.FieldGoodsID()),
		).
			From(table, finalAdditions...),
		&id,
	)

	return uint64(id.Int64), err

}
