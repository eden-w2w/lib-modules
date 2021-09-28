package enums

//go:generate eden generate enum --type-name=SettlementType
// api:enum
type SettlementType uint8

// 结算周期类型
const (
	SETTLEMENT_TYPE_UNKNOWN SettlementType = iota
	SETTLEMENT_TYPE__WEEK                  // 周
	SETTLEMENT_TYPE__MONTH                 // 月
)
