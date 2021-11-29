package discounts

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

type CreateDiscountParams struct {
	// 营销名称
	Name string `json:"name"`
	// 营销类型
	Type enums.DiscountType `json:"type"`
	// 计算方式
	Cal enums.DiscountCal `json:"cal"`
	// 有效期开始
	ValidityStart datatypes.MySQLTimestamp `json:"validityStart"`
	// 有效期结束
	ValidityEnd datatypes.MySQLTimestamp `json:"validityEnd"`
	// 单用户优惠次数上限
	UserLimit uint32 `json:"userLimit" default:""`
	// 总数限制
	Limit uint32 `json:"limit" default:""`
	// 优惠上限
	DiscountLimit uint64 `json:"discountLimit" default:""`

	// 总价满额单价优惠阈值
	MinTotalPrice uint64 `json:"minTotalPrice" default:""`
	// 单价折扣比例
	DiscountRate float64 `json:"discountRate" default:""`
	// 阶梯式折扣比例
	MultiStepRate types.MultiStepRateConfig `json:"multiStepRate" default:""`

	// 单价立减金额
	DiscountAmount uint64 `json:"discountAmount" default:""`
	// 阶梯式立减金额
	MultiStepReduction types.MultiStepReductionConfig `json:"multiStepReduction" default:""`
}

func (p CreateDiscountParams) Model() (model *databases.MarketingDiscount, err error) {
	if p.Type == enums.DISCOUNT_TYPE__ALL {
		if p.Cal == enums.DISCOUNT_CAL__UNIT {
			if p.DiscountAmount == 0 {
				err = general_errors.BadRequest.StatusError().WithErrTalk().WithMsg("需要指定单价立减金额")
				return
			}
		} else if p.Cal == enums.DISCOUNT_CAL__MULTISTEP {
			if len(p.MultiStepReduction) == 0 {
				err = general_errors.BadRequest.StatusError().WithErrTalk().WithMsg("需要指定阶梯式立减金额")
				return
			}
		}
	} else if p.Type == enums.DISCOUNT_TYPE__ALL_PERCENT {
		if p.Cal == enums.DISCOUNT_CAL__UNIT {
			if p.DiscountRate == 0 {
				err = general_errors.BadRequest.StatusError().WithErrTalk().WithMsg("需要指定单价立减金额")
				return
			}
		} else if p.Cal == enums.DISCOUNT_CAL__MULTISTEP {
			if len(p.MultiStepRate) == 0 {
				err = general_errors.BadRequest.StatusError().WithErrTalk().WithMsg("需要指定阶梯式立减金额")
				return
			}
		}
	}
	id, _ := id_generator.GetGenerator().GenerateUniqueID()
	model = &databases.MarketingDiscount{
		DiscountID:         id,
		Name:               p.Name,
		Type:               p.Type,
		Status:             enums.DISCOUNT_STATUS__READY,
		Cal:                p.Cal,
		ValidityStart:      p.ValidityStart,
		ValidityEnd:        p.ValidityEnd,
		UserLimit:          p.UserLimit,
		Limit:              p.Limit,
		DiscountLimit:      p.DiscountLimit,
		MinTotalPrice:      p.MinTotalPrice,
		DiscountRate:       p.DiscountRate,
		MultiStepRate:      p.MultiStepRate,
		DiscountAmount:     p.DiscountAmount,
		MultiStepReduction: p.MultiStepReduction,
	}
	return
}

type GetDiscountsParams struct {
	// 营销类型
	Type enums.DiscountType `in:"query" name:"type" default:""`
	// 营销状态
	Status enums.DiscountStatus `in:"query" name:"status" default:""`
	// 计算方式
	Cal enums.DiscountCal `in:"query" name:"cal" default:""`
	modules.Pagination
}

func (p GetDiscountsParams) Conditions() builder.SqlCondition {
	var condition builder.SqlCondition
	var model = databases.MarketingDiscount{}
	if p.Type != enums.DISCOUNT_TYPE_UNKNOWN {
		condition = builder.And(condition, model.FieldType().Eq(p.Type))
	}
	if p.Status != enums.DISCOUNT_STATUS_UNKNOWN {
		condition = builder.And(condition, model.FieldStatus().Eq(p.Status))
	}
	if p.Cal != enums.DISCOUNT_CAL_UNKNOWN {
		condition = builder.And(condition, model.FieldCal().Eq(p.Cal))
	}
	return condition
}

func (p GetDiscountsParams) Additions() []builder.Addition {
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

type UpdateDiscountParams struct {
	// 营销名称
	Name string `json:"name" default:""`
	// 营销状态
	Status enums.DiscountStatus `json:"status" default:""`
	// 有效期开始
	ValidityStart datatypes.MySQLTimestamp `json:"validityStart" default:""`
	// 有效期结束
	ValidityEnd datatypes.MySQLTimestamp `json:"validityEnd" default:""`
	// 单用户优惠次数上限
	UserLimit *uint32 `json:"userLimit" default:""`
	// 总数限制
	Limit *uint32 `json:"limit" default:""`
	// 已优惠次数
	Times *uint64 `json:"times" default:""`
	// 优惠上限
	DiscountLimit *uint64 `json:"discountLimit" default:""`

	// 总价满额单价优惠阈值
	MinTotalPrice *uint64 `json:"minTotalPrice" default:""`
	// 单价折扣比例
	DiscountRate float64 `json:"discountRate" default:""`
	// 阶梯式折扣比例
	MultiStepRate types.MultiStepRateConfig `json:"multiStepRate" default:""`

	// 单价立减金额
	DiscountAmount uint64 `json:"discountAmount" default:""`
	// 阶梯式立减金额
	MultiStepReduction types.MultiStepReductionConfig `json:"multiStepReduction" default:""`
}

func (p *UpdateDiscountParams) Fill(model *databases.MarketingDiscount) {
	if p.Name != "" {
		model.Name = p.Name
	}
	if p.Status != enums.DISCOUNT_STATUS_UNKNOWN {
		model.Status = p.Status
	}
	if p.ValidityStart != datatypes.TimestampZero {
		model.ValidityStart = p.ValidityStart
	}
	if p.ValidityEnd != datatypes.TimestampZero {
		model.ValidityEnd = p.ValidityEnd
	}
	if p.UserLimit != nil {
		model.UserLimit = *p.UserLimit
	}
	if p.Limit != nil {
		model.Limit = *p.Limit
	}
	if p.Times != nil {
		model.Times = *p.Times
	}
	if p.DiscountLimit != nil {
		model.DiscountLimit = *p.DiscountLimit
	}
	if p.MinTotalPrice != nil {
		model.MinTotalPrice = *p.MinTotalPrice
	}
	if p.DiscountRate != 0 {
		model.DiscountRate = p.DiscountRate
	}
	if len(p.MultiStepRate) > 0 {
		model.MultiStepRate = p.MultiStepRate
	}
	if p.DiscountAmount != 0 {
		model.DiscountAmount = p.DiscountAmount
	}
	if len(p.MultiStepReduction) > 0 {
		model.MultiStepReduction = p.MultiStepReduction
	}
}
