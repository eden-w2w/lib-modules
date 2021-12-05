package enums

//go:generate eden generate enum --type-name=AreaLevel
// api:enum
type AreaLevel uint8

// 行政区级别
const (
	AREA_LEVEL_UNKNOWN   AreaLevel = iota
	AREA_LEVEL__COUNTRY            // 国家级
	AREA_LEVEL__PROVINCE           // 省级
	AREA_LEVEL__CITY               // 市级
	AREA_LEVEL__DISTRICT           // 区县级
	AREA_LEVEL__STREET             // 街道级
)
