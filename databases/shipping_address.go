package databases

import (
	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/sqlx/datatypes"
	"time"
)

//go:generate eden generate model ShippingAddress --database Config.DB --with-comments
//go:generate eden generate tag ShippingAddress --defaults=true
// @def primary ID
// @def unique_index U_shipping_id ShippingID
// @def index I_user UserID Default
type ShippingAddress struct {
	datatypes.PrimaryID
	// 业务ID
	ShippingID uint64 `json:"shippingID,string" db:"f_shipping_id"`
	// 用户ID
	UserID uint64 `json:"userID,string" db:"f_user_id"`
	// 收件人
	Recipients string `json:"recipients" db:"f_recipients"`
	// 省名称
	Province string `json:"province" db:"f_province"`
	// 省编码
	ProvinceCode string `json:"provinceCode" db:"f_province_code"`
	// 市名称
	City string `json:"city" db:"f_city,default=''"`
	// 市编码
	CityCode string `json:"cityCode" db:"f_city_code,default=''"`
	// 区县名称
	District string `json:"district" db:"f_district,default=''"`
	// 区县编码
	DistrictCode string `json:"districtCode" db:"f_district_code,default=''"`
	// 街道名称
	Street string `json:"street" db:"f_street,default=''"`
	// 街道编码
	StreetCode string `json:"streetCode" db:"f_street_code,default=''"`
	// 详细地址
	Address string `json:"address" db:"f_address"`
	// 联系电话
	Mobile string `json:"mobile" db:"f_mobile"`
	// 是否默认
	Default datatypes.Bool `json:"default" db:"f_default"`

	datatypes.OperateTime
}

func (m *ShippingAddress) ResetAllDefault(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValues{
		m.FieldKeyDefault():   datatypes.BOOL_FALSE,
		m.FieldKeyUpdatedAt(): datatypes.Timestamp(time.Now()),
	}

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					m.FieldUserID().Eq(m.UserID),
					m.FieldDeletedAt().Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("ShippingAddress.ResetAllDefault"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}
