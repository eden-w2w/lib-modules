package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
)

//go:generate eden generate model SettlementFlow --database Config.DB --with-comments
//go:generate eden generate tag SettlementFlow --defaults=true
// @def primary ID
// @def unique_index U_settlement_id SettlementID
// @def unique_index U_interval UserID Name
type SettlementFlow struct {
	datatypes.PrimaryID
	// 结算单ID
	SettlementID uint64 `json:"settlementID,string" db:"f_settlement_id"`
	// 用户ID
	UserID uint64 `json:"userID,string" db:"f_user_id"`
	// 昵称
	NickName string `json:"nickName" db:"f_nick_name,default=''"`
	// OpenID
	OpenID string `json:"openID" db:"f_open_id"`
	// 名称
	Name string `json:"name" db:"f_name"`
	// 销售总额
	TotalSales uint64 `json:"totalSales" db:"f_total_sales"`
	// 计算比例
	Proportion float64 `json:"proportion" db:"f_proportion"`
	// 结算金额
	Amount uint64 `json:"amount" db:"f_amount"`
	// 结算状态
	Status enums.SettlementStatus `json:"status" db:"f_status"`

	datatypes.OperateTime
}
