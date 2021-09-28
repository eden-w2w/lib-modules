package types

import (
	"database/sql/driver"
	"encoding/json"
)

type JsonArrayString []string

func (v *JsonArrayString) DataType(driverName string) string {
	if driverName == "mysql" {
		return "text"
	}
	return ""
}

func (v JsonArrayString) Value() (driver.Value, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (v *JsonArrayString) Scan(src interface{}) error {
	if data, ok := src.([]byte); ok {
		return json.Unmarshal(data, v)
	}
	return nil
}
