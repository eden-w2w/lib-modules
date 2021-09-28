package enums

//go:generate eden generate enum --type-name=SettlementStatus
// api:enum
type SettlementStatus uint8

// 结算状态
const (
	SETTLEMENT_STATUS_UNKNOWN   SettlementStatus = iota
	SETTLEMENT_STATUS__CREATED                   // 待结算
	SETTLEMENT_STATUS__COMPLETE                  // 已结算
)
