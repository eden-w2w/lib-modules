package settlement_flow

import (
	"errors"
	"fmt"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"strconv"
	"strings"
)

type SettlementConfig struct {
	// 结算周期
	SettlementType enums.SettlementType
	// 结算节点 周：0-6，月：1-31
	SettlementDate uint8
	// 提成比例规则
	SettlementRules []SettlementRule
}

func (c SettlementConfig) ToSettlementCronRule() string {
	if c.SettlementType == enums.SETTLEMENT_TYPE__WEEK {
		return fmt.Sprintf("0 0 0 * * %d", c.SettlementDate)
	} else if c.SettlementType == enums.SETTLEMENT_TYPE__MONTH {
		return fmt.Sprintf("0 0 0 %d * *", c.SettlementDate)
	}
	return ""
}

func (s *SettlementRule) UnmarshalText(text []byte) (err error) {
	strList := strings.Split(string(text), "|")
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
	s.Proportion, err = strconv.ParseFloat(strList[1], 64)
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
	// 名称
	Name string `in:"body" json:"name"`
	// 销售总额
	TotalSales uint64 `in:"body" json:"totalSales"`
	// 计算比例
	Proportion float64 `in:"body" json:"proportion"`
	// 结算金额
	Amount uint64 `in:"body" json:"amount"`
}
