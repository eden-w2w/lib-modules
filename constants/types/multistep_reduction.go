package types

import (
	"database/sql/driver"
	"encoding/json"
)

type MultiStepReduction struct {
	// 总价下限>=
	Min uint64 `json:"min"`
	// 总价上限<
	Max uint64 `json:"max" default:""`
	// 无上限
	NoMax bool `json:"noMax" default:""`
	// 立减
	Reduction uint64 `json:"reduction"`
}

type MultiStepReductionConfig []MultiStepReduction

func (m MultiStepReductionConfig) DiscountAmount(totalPrice uint64) uint64 {
	for _, reduction := range m {
		if reduction.NoMax {
			if totalPrice >= reduction.Min {
				return reduction.Reduction
			}
		} else {
			if totalPrice >= reduction.Min && totalPrice < reduction.Max {
				return reduction.Reduction
			}
		}
	}
	return 0
}

func (m *MultiStepReductionConfig) DataType(driverName string) string {
	if driverName == "mysql" {
		return "text"
	}
	return ""
}

func (m MultiStepReductionConfig) Value() (driver.Value, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (m *MultiStepReductionConfig) Scan(src interface{}) error {
	if data, ok := src.([]byte); ok {
		if len(data) == 0 {
			*m = MultiStepReductionConfig{}
			return nil
		}
		return json.Unmarshal(data, m)
	}
	return nil
}
