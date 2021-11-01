package settlement_flow

import (
	"errors"
	"fmt"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules"
	"strconv"
	"strings"
	"time"
)

type SettlementConfig struct {
	// 结算周期
	SettlementType enums.SettlementType
	// 结算节点 周：0-6，月：1-31
	SettlementDate uint8
	// 提成比例规则
	SettlementRules []SettlementRule
	// 结算等待期，同时也是退货有效期
	SettlementDuration time.Duration
}

func (c SettlementConfig) ToSettlementCronRule() string {
	if c.SettlementType == enums.SETTLEMENT_TYPE__WEEK {
		return fmt.Sprintf("0 0 0 * * %d", c.SettlementDate)
	} else if c.SettlementType == enums.SETTLEMENT_TYPE__MONTH {
		return fmt.Sprintf("0 0 0 %d * *", c.SettlementDate)
	}
	return ""
}

func (c SettlementConfig) GetProportion(sales uint64) float64 {
	for _, rule := range c.SettlementRules {
		if sales >= rule.MinSales && sales < rule.MaxSales {
			return rule.Proportion
		}
	}
	return 0
}

func (s *SettlementRule) UnmarshalText(text []byte) (err error) {
	strList := strings.Split(strings.Trim(string(text), " "), "|")
	if len(strList) != 3 {
		return errors.New("SettlementRule not support more than 3 args")
	}
	s.MinSales, err = strconv.ParseUint(strList[0], 10, 64)
	if err != nil {
		return
	}
	s.MaxSales, err = strconv.ParseUint(strList[1], 10, 64)
	if err != nil {
		return
	}
	s.Proportion, err = strconv.ParseFloat(strList[2], 64)
	return
}

func (s SettlementRule) MarshalText() (text []byte, err error) {
	str := fmt.Sprintf("%d|%d|%f", s.MinSales, s.MaxSales, s.Proportion)
	return []byte(str), nil
}

type SettlementRule struct {
	// 最小销售量（闭区间）
	MinSales uint64
	// 最大销售量（开区间）
	MaxSales uint64
	// 计提比例
	Proportion float64
}

func (s SettlementRule) String() string {
	str, _ := s.MarshalText()
	return string(str)
}

type CreateSettlementParams struct {
	// 用户ID
	UserID uint64 `in:"body" json:"userID,string" default:""`
	// 昵称
	NickName string `in:"body" json:"nickName" default:""`
	// OpenID
	OpenID string `in:"body" json:"openID"`
	// 名称
	Name string `in:"body" json:"name"`
	// 销售总额
	TotalSales uint64 `in:"body" json:"totalSales"`
	// 计算比例
	Proportion float64 `in:"body" json:"proportion"`
	// 结算金额
	Amount uint64 `in:"body" json:"amount"`
}

type GetSettlementFlowsParams struct {
	// 用户ID
	UserID uint64 `name:"userID,string" in:"query" default:""`
	// 名称
	Name string `name:"name" in:"query" default:""`
	// 结算状态
	Status enums.SettlementStatus `name:"status" in:"query" default:""`
	// 创建时间大于等于
	CreateGte datatypes.MySQLTimestamp `name:"createGte" in:"query" default:""`
	// 创建时间小于
	CreateLt datatypes.MySQLTimestamp `name:"createLt" in:"query" default:""`
	modules.Pagination
}

func (p GetSettlementFlowsParams) Conditions() builder.SqlCondition {
	var condition builder.SqlCondition
	model := databases.SettlementFlow{}

	if p.UserID != 0 {
		condition = builder.And(condition, model.FieldUserID().Eq(p.UserID))
	}
	if p.Name != "" {
		condition = builder.And(condition, model.FieldName().Eq(p.Name))
	}
	if p.Status != enums.SETTLEMENT_STATUS_UNKNOWN {
		condition = builder.And(condition, model.FieldStatus().Eq(p.Status))
	}
	if p.CreateGte != datatypes.TimestampZero {
		condition = builder.And(condition, model.FieldCreatedAt().Gte(p.CreateGte))
	}
	if p.CreateLt != datatypes.TimestampZero {
		condition = builder.And(condition, model.FieldCreatedAt().Lt(p.CreateLt))
	}

	return condition
}

func (p GetSettlementFlowsParams) Additions() []builder.Addition {
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

type UpdateSettlementParams struct {
	// 销售总额
	TotalSales uint64 `json:"totalSales" default:""`
	// 计算比例
	Proportion float64 `json:"proportion" default:""`
	// 结算金额
	Amount uint64 `json:"amount" default:""`
	// 结算状态
	Status enums.SettlementStatus `json:"status" default:""`
}
