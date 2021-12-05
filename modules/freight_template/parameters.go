package freight_template

import (
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/constants/types"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules"
	"github.com/eden-w2w/lib-modules/modules/id_generator"
)

type CreateTemplateRuleParams struct {
	// 包含区域
	Area types.FreightRuleAreas `json:"area"`
	// 是否包邮
	IsFreeFreight datatypes.Bool `json:"isFreeFreight"`
	// 展示话术
	Description string `json:"description"`
	// -------------------------------------------------
	// 运费设置
	// 首重（克）/首件（个）范围
	FirstRange uint32 `json:"firstRange" default:""`
	// 首重/首件价格
	FirstPrice uint32 `json:"firstPrice" default:""`
	// 续重（克）/续件（个）范围
	ContinueRange uint32 `json:"continueRange" default:""`
	// 续重/续件价格
	ContinuePrice uint32 `json:"continuePrice" default:""`
}

func (p CreateTemplateRuleParams) Model() (rule *databases.FreightTemplateRules, err error) {
	if p.IsFreeFreight.False() {
		if p.FirstRange == 0 {
			err = general_errors.BadRequest.StatusError().WithErrTalk().WithMsg("请填写首重（克）/首件（个）范围")
			return
		}
		if p.FirstPrice == 0 {
			err = general_errors.BadRequest.StatusError().WithErrTalk().WithMsg("请填写首重/首件价格")
			return
		}
		if p.ContinueRange == 0 {
			err = general_errors.BadRequest.StatusError().WithErrTalk().WithMsg("请填写续重（克）/续件（个）范围")
			return
		}
		if p.ContinuePrice == 0 {
			err = general_errors.BadRequest.StatusError().WithErrTalk().WithMsg("请填写续重/续件价格")
			return
		}
	}
	id := id_generator.GetGenerator().GenerateUniqueID()
	rule = &databases.FreightTemplateRules{
		RuleID:        id,
		Area:          p.Area,
		IsFreeFreight: p.IsFreeFreight,
		Description:   p.Description,
		FirstRange:    p.FirstRange,
		FirstPrice:    p.FirstPrice,
		ContinueRange: p.ContinueRange,
		ContinuePrice: p.ContinuePrice,
	}

	return
}

type CreateTemplateParams struct {
	// 名称
	Name string `json:"name"`
	// 发货地
	DispatchAddr string `json:"dispatchAddr"`
	// 发货时间
	DispatchTime uint32 `json:"dispatchTime" default:""`
	// 是否全场包邮
	IsFreeFreight datatypes.Bool `json:"isFreeFreight"`
	// 计费方式
	Cal enums.FreightCal `json:"cal" default:""`
	// -------------------------------------------------
	// 默认运费设置
	// 首重（克）/首件（个）范围
	FirstRange uint32 `json:"firstRange" default:""`
	// 首重/首件价格
	FirstPrice uint32 `json:"firstPrice" default:""`
	// 续重（克）/续件（个）范围
	ContinueRange uint32 `json:"continueRange" default:""`
	// 续重/续件价格
	ContinuePrice uint32 `json:"continuePrice" default:""`
}

func (p CreateTemplateParams) Model() (template *databases.FreightTemplate, err error) {
	if p.IsFreeFreight.False() {
		if p.FirstRange == 0 {
			err = general_errors.BadRequest.StatusError().WithErrTalk().WithMsg("请填写首重（克）/首件（个）范围")
			return
		}
		if p.FirstPrice == 0 {
			err = general_errors.BadRequest.StatusError().WithErrTalk().WithMsg("请填写首重/首件价格")
			return
		}
		if p.ContinueRange == 0 {
			err = general_errors.BadRequest.StatusError().WithErrTalk().WithMsg("请填写续重（克）/续件（个）范围")
			return
		}
		if p.ContinuePrice == 0 {
			err = general_errors.BadRequest.StatusError().WithErrTalk().WithMsg("请填写续重/续件价格")
			return
		}
	}
	id := id_generator.GetGenerator().GenerateUniqueID()
	template = &databases.FreightTemplate{
		TemplateID:    id,
		Name:          p.Name,
		DispatchAddr:  p.DispatchAddr,
		DispatchTime:  p.DispatchTime,
		IsFreeFreight: p.IsFreeFreight,
		Cal:           p.Cal,
		FirstRange:    p.FirstRange,
		FirstPrice:    p.FirstPrice,
		ContinueRange: p.ContinueRange,
		ContinuePrice: p.ContinuePrice,
	}

	return
}

type UpdateTemplateParams struct {
	// 名称
	Name string `json:"name" default:""`
	// 发货地
	DispatchAddr string `json:"dispatchAddr" default:""`
	// 发货时间
	DispatchTime uint32 `json:"dispatchTime" default:""`
	// 是否全场包邮
	IsFreeFreight datatypes.Bool `json:"isFreeFreight" default:""`
	// 计费方式
	Cal enums.FreightCal `json:"cal" default:""`
	// -------------------------------------------------
	// 默认运费设置
	// 首重（克）/首件（个）范围
	FirstRange uint32 `json:"firstRange" default:""`
	// 首重/首件价格
	FirstPrice uint32 `json:"firstPrice" default:""`
	// 续重（克）/续件（个）范围
	ContinueRange uint32 `json:"continueRange" default:""`
	// 续重/续件价格
	ContinuePrice uint32 `json:"continuePrice" default:""`
}

func (p UpdateTemplateParams) Fill(model *databases.FreightTemplate) {
	model.Name = p.Name
	model.DispatchAddr = p.DispatchAddr
	model.DispatchTime = p.DispatchTime
	model.IsFreeFreight = p.IsFreeFreight
	model.Cal = p.Cal
	model.FirstRange = p.FirstRange
	model.FirstPrice = p.FirstPrice
	model.ContinueRange = p.ContinueRange
	model.ContinuePrice = p.ContinuePrice
}

type GetTemplatesParams struct {
	// 名称
	Name string `in:"query" name:"name" default:""`
	// 是否全场包邮
	IsFreeFreight datatypes.Bool `in:"query" name:"isFreeFreight" default:""`
	// 计费方式
	Cal enums.FreightCal `in:"query" name:"cal" default:""`
	modules.Pagination
}

func (p GetTemplatesParams) Conditions() builder.SqlCondition {
	var condition builder.SqlCondition
	model := &databases.FreightTemplate{}

	if p.Name != "" {
		condition = builder.And(condition, model.FieldName().Eq(p.Name))
	}
	if p.IsFreeFreight != datatypes.BOOL_UNKNOWN {
		condition = builder.And(condition, model.FieldIsFreeFreight().Eq(p.IsFreeFreight))
	}
	if p.Cal != enums.FREIGHT_CAL_UNKNOWN {
		condition = builder.And(condition, model.FieldCal().Eq(p.Cal))
	}

	return condition
}

func (p GetTemplatesParams) Additions() []builder.Addition {
	var additions = make([]builder.Addition, 0)

	if p.Size != 0 {
		limit := builder.Limit(int64(p.Size))
		if p.Offset != 0 {
			limit = limit.Offset(int64(p.Offset))
		}
		additions = append(additions, limit)
	}

	additions = append(additions, builder.OrderBy(builder.DescOrder((&databases.Order{}).FieldCreatedAt())))

	return additions
}

type UpdateTemplateRuleParams struct {
	// 包含区域
	Area types.FreightRuleAreas `json:"area" default:""`
	// 是否包邮
	IsFreeFreight datatypes.Bool `json:"isFreeFreight" default:""`
	// 展示话术
	Description string `json:"description" default:""`
	// -------------------------------------------------
	// 运费设置
	// 首重（克）/首件（个）范围
	FirstRange uint32 `json:"firstRange" default:""`
	// 首重/首件价格
	FirstPrice uint32 `json:"firstPrice" default:""`
	// 续重（克）/续件（个）范围
	ContinueRange uint32 `json:"continueRange" default:""`
	// 续重/续件价格
	ContinuePrice uint32 `json:"continuePrice" default:""`
}

func (p UpdateTemplateRuleParams) Fill(model *databases.FreightTemplateRules) {
	model.Area = p.Area
	model.IsFreeFreight = p.IsFreeFreight
	model.Description = p.Description
	model.FirstRange = p.FirstRange
	model.FirstPrice = p.FirstPrice
	model.ContinueRange = p.ContinueRange
	model.ContinuePrice = p.ContinuePrice
}

type GetTemplateRuleParams struct {
	// 是否包邮
	IsFreeFreight datatypes.Bool `in:"query" name:"isFreeFreight" default:""`
}

func (p GetTemplateRuleParams) Conditions() builder.SqlCondition {
	var condition builder.SqlCondition
	model := &databases.FreightTemplate{}

	if p.IsFreeFreight != datatypes.BOOL_UNKNOWN {
		condition = builder.And(condition, model.FieldIsFreeFreight().Eq(p.IsFreeFreight))
	}

	return condition
}
