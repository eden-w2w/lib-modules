package types

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/shopspring/decimal"
)

type MultiStepRate struct {
	// 总价下限>=
	Min uint64 `json:"min"`
	// 总价上限<
	Max uint64 `json:"max" default:""`
	// 无上限
	NoMax bool `json:"noMax" default:""`
	// 折扣
	Rate float64 `json:"rate"`
}

type MultiStepRateConfig []MultiStepRate

func (m MultiStepRateConfig) DiscountAmount(totalPrice uint64) uint64 {
	for _, reduction := range m {
		if reduction.NoMax {
			if totalPrice >= reduction.Min {
				return uint64(decimal.NewFromInt(int64(totalPrice)).Mul(decimal.NewFromFloat(reduction.Rate)).IntPart())
			}
		} else {
			if totalPrice >= reduction.Min && totalPrice < reduction.Max {
				return uint64(decimal.NewFromInt(1).Sub(decimal.NewFromFloat(reduction.Rate)).Mul(decimal.NewFromInt(int64(totalPrice))).IntPart())
			}
		}
	}
	return 0
}

func (m MultiStepRateConfig) DiscountAmountUnit(totalPrice, unitPrice uint64) uint64 {
	for _, reduction := range m {
		if reduction.NoMax {
			if totalPrice >= reduction.Min {
				return uint64(decimal.NewFromInt(int64(unitPrice)).Mul(decimal.NewFromFloat(reduction.Rate)).IntPart())
			}
		} else {
			if totalPrice >= reduction.Min && totalPrice < reduction.Max {
				return uint64(decimal.NewFromInt(1).Sub(decimal.NewFromFloat(reduction.Rate)).Mul(decimal.NewFromInt(int64(unitPrice))).IntPart())
			}
		}
	}
	return 0
}

func (m *MultiStepRateConfig) DataType(driverName string) string {
	if driverName == "mysql" {
		return "text"
	}
	return ""
}

func (m MultiStepRateConfig) Value() (driver.Value, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (m *MultiStepRateConfig) Scan(src interface{}) error {
	if data, ok := src.([]byte); ok {
		if len(data) == 0 {
			*m = MultiStepRateConfig{}
			return nil
		}
		return json.Unmarshal(data, m)
	}
	return nil
}
