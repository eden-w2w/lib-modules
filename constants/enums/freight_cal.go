package enums

//go:generate eden generate enum --type-name=FreightCal
// api:enum
type FreightCal uint8

// 运费计费方式
const (
	FREIGHT_CAL_UNKNOWN FreightCal = iota
	FREIGHT_CAL__COUNT             // 按数量
	FREIGHT_CAL__WEIGHT            // 按重量
)
