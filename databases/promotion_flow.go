package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
)

//go:generate eden generate model PromotionFlow --database Config.DB --with-comments
//go:generate eden generate tag PromotionFlow --defaults=true
// @def primary ID
// @def unique_index U_flow_id FlowID
// @def index I_user_id UserID CreatedAt
// @def index I_payment_flow_id PaymentFlowID
// @def index I_settlement SettlementID
type PromotionFlow struct {
	datatypes.PrimaryID
	// 流水ID
	FlowID uint64 `json:"flowID,string" db:"f_flow_id"`
	// 获得奖励的用户ID
	UserID uint64 `json:"userID,string" db:"f_user_id"`
	// 获得奖励的用户昵称
	UserNickName string `json:"userNickName" db:"f_user_nick_name,default=''"`
	// 奖励来源用户ID
	RefererID uint64 `json:"refererID,string" db:"f_referer_id"`
	// 奖励来源的用户昵称
	RefererNickName string `json:"refererNickName" db:"f_referer_nick_name,default=''"`
	// 订单金额
	Amount uint64 `json:"amount" db:"f_amount"`
	// 关联的支付流水
	PaymentFlowID uint64 `json:"paymentFlowID,string" db:"f_payment_flow_id,default='0'"`
	// 关联的结算单ID
	SettlementID uint64 `json:"settlementID,string" db:"f_settlement_id,default='0'"`

	datatypes.OperateTime
}
