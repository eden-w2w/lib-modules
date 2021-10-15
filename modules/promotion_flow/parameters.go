package promotion_flow

import (
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules"
)

type CreatePromotionFlowParams struct {
	// 获得奖励的用户ID
	UserID uint64 `json:"userID,string"`
	// 获得奖励的用户昵称
	UserNickName string `json:"userNickName" default:""`
	// 获取奖励的用户OpenID
	UserOpenID string `json:"userOpenID"`
	// 奖励来源用户ID
	RefererID uint64 `json:"refererID,string"`
	// 奖励来源的用户昵称
	RefererNickName string `json:"refererNickName" default:""`
	// 奖励来源的用户OpenID
	RefererOpenID string `json:"refererOpenID"`
	// 奖励金额
	Amount uint64 `json:"amount"`
	// 关联的支付流水
	PaymentFlowID uint64 `json:"paymentFlowID"`
}

type GetPromotionFlowParams struct {
	// 获得奖励的用户ID
	UserID uint64 `name:"userID,string" in:"query" default:""`
	// 奖励来源用户ID
	RefererID uint64 `name:"refererID,string" in:"query" default:""`
	// 关联的支付流水
	PaymentFlowID uint64 `name:"paymentFlowID,string" in:"query" default:""`
	// 关联的结算单ID
	SettlementID uint64 `name:"settlementID,string" in:"query" default:""`
	// 是否只查询未结算的流水
	IsNotSettlement datatypes.Bool `name:"isNotSettlement" in:"query" default:""`
	// 创建时间大于等于
	CreateGte datatypes.MySQLTimestamp `name:"createGte" in:"query" default:""`
	// 创建时间小于
	CreateLt datatypes.MySQLTimestamp `name:"createLt" in:"query" default:""`
	modules.Pagination
}

func (p GetPromotionFlowParams) Conditions() builder.SqlCondition {
	var condition builder.SqlCondition
	var model = databases.PromotionFlow{}
	if p.UserID != 0 {
		condition = builder.And(condition, model.FieldUserID().Eq(p.UserID))
	}
	if p.RefererID != 0 {
		condition = builder.And(condition, model.FieldRefererID().Eq(p.RefererID))
	}
	if p.PaymentFlowID != 0 {
		condition = builder.And(condition, model.FieldPaymentFlowID().Eq(p.PaymentFlowID))
	}
	if p.CreateGte != datatypes.TimestampZero {
		condition = builder.And(condition, model.FieldCreatedAt().Gte(p.CreateGte))
	}
	if p.CreateLt != datatypes.TimestampZero {
		condition = builder.And(condition, model.FieldCreatedAt().Lt(p.CreateLt))
	}
	if p.IsNotSettlement != datatypes.BOOL_UNKNOWN {
		if p.IsNotSettlement == datatypes.BOOL_TRUE {
			condition = builder.And(condition, model.FieldSettlementID().Eq(0))
		} else {
			if p.SettlementID != 0 {
				condition = builder.And(condition, model.FieldSettlementID().Eq(p.SettlementID))
			} else {
				condition = builder.And(condition, model.FieldSettlementID().Neq(0))
			}
		}
	}
	return condition
}

func (p GetPromotionFlowParams) Additions() []builder.Addition {
	var additions = make([]builder.Addition, 0)

	if p.Size != 0 {
		limit := builder.Limit(int64(p.Size))
		if p.Offset != 0 {
			limit.Offset(int64(p.Offset))
		}
		additions = append(additions, limit)
	}

	additions = append(additions, builder.OrderBy(builder.DescOrder((&databases.Order{}).FieldCreatedAt())))

	return additions
}
