package types

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/eden-w2w/lib-modules/constants/enums"
)

type GoodsPicture struct {
	ID   uint32            `json:"id"`
	Type enums.PictureType `json:"type"`
	Url  string            `json:"url"`
}

type GoodsPictures []GoodsPicture

func (g *GoodsPictures) DataType(driverName string) string {
	if driverName == "mysql" {
		return "text"
	}
	return ""
}

func (g GoodsPictures) Value() (driver.Value, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (g *GoodsPictures) Scan(src interface{}) error {
	*g = make([]GoodsPicture, 0)
	if data, ok := src.([]byte); ok {
		return json.Unmarshal(data, &g)
	}
	return nil
}
