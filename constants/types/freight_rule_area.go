package types

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/pkg/search"
)

type FreightRuleArea struct {
	ADCode string           `json:"adCode"`
	Name   string           `json:"name"`
	Level  enums.AreaLevel  `json:"level"`
	Parent *FreightRuleArea `json:"parent" default:""`
}

type FreightRuleAreas []FreightRuleArea

func (v *FreightRuleAreas) DataType(driverName string) string {
	if driverName == "mysql" {
		return "text"
	}
	return ""
}

func (v FreightRuleAreas) Value() (driver.Value, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (v *FreightRuleAreas) Scan(src interface{}) error {
	if data, ok := src.([]byte); ok {
		if len(data) == 0 {
			*v = FreightRuleAreas{}
			return nil
		}
		return json.Unmarshal(data, v)
	}
	return nil
}

func (v FreightRuleAreas) Contain(area FreightRuleArea) bool {
	ok, _, err := search.In(
		v, area, func(current interface{}, needle interface{}) bool {
			var p = current.(FreightRuleArea)
			var t = needle.(FreightRuleArea)
			if p.ADCode == t.ADCode {
				return true
			}
			return false
		},
	)
	if !ok || err != nil {
		return false
	}
	return true
}

func (v FreightRuleAreas) ContainsAll(areas FreightRuleAreas) bool {
	for _, area := range areas {
		if !v.Contain(area) {
			return false
		}
	}
	return true
}

func (v FreightRuleAreas) ContainsOne(areas FreightRuleAreas) bool {
	for _, area := range areas {
		if v.Contain(area) {
			return true
		}
	}
	return false
}
