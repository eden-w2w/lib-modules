package types

import (
	"database/sql/driver"
	"encoding/json"
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
				return uint64(float64(totalPrice) * reduction.Rate)
			}
		} else {
			if totalPrice >= reduction.Min && totalPrice < reduction.Max {
				return uint64(float64(totalPrice) * (1.0 - reduction.Rate))
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
