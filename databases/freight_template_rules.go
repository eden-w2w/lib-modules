package databases

import (
	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/types"
	"time"
)

//go:generate eden generate model FreightTemplateRules --database Config.DB --with-comments
//go:generate eden generate tag FreightTemplateRules --defaults=true
// @def primary ID
// @def unique_index U_rule_id RuleID
// @def index I_template_id TemplateID IsFreeFreight
type FreightTemplateRules struct {
	datatypes.PrimaryID
	// 业务ID
	RuleID uint64 `json:"ruleID,string" db:"f_rule_id"`
	// 模板ID
	TemplateID uint64 `json:"templateID,string" db:"f_template_id"`
	// 包含区域
	Area types.FreightRuleAreas `json:"area" db:"f_area"`
	// 是否包邮
	IsFreeFreight datatypes.Bool `json:"isFreeFreight" db:"f_is_free_freight"`
	// 展示话术
	Description string `json:"description" db:"f_description"`
	// -------------------------------------------------
	// 运费设置
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

func (m *FreightTemplateRules) SoftDeleteByTemplateID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValues{}
	if _, ok := fieldValues["DeletedAt"]; !ok {
		fieldValues["DeletedAt"] = datatypes.Timestamp(time.Now())
	}

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					m.FieldTemplateID().Eq(m.TemplateID),
					m.FieldDeletedAt().Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplateRules.SoftDeleteByRuleID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}
