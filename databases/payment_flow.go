package databases

import (
	"errors"
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
)

//go:generate eden generate model PaymentFlow --database Config.DB --with-comments
//go:generate eden generate tag PaymentFlow --defaults=true
// @def primary ID
// @def unique_index U_flow_id FlowID
// @def index I_order_id OrderID Status
// @def index I_expire ExpiredAt
type PaymentFlow struct {
	datatypes.PrimaryID
	// 流水ID
	FlowID uint64 `json:"flowID,string" db:"f_flow_id"`
	// 用户ID
	UserID uint64 `json:"userID,string" db:"f_user_id"`
	// 关联订单号
	OrderID uint64 `json:"orderID,string" db:"f_order_id"`
	// 支付金额
	Amount uint64 `json:"amount" db:"f_amount"`
	// 支付方式
	PaymentMethod enums.PaymentMethod `json:"paymentMethod" db:"f_payment_method"`
	// 支付系统流水号
	RemoteFlowID string `json:"remoteFlowID" db:"f_remote_flow_id,size=255,default=''"`
	// 支付状态
	Status enums.PaymentStatus `json:"status" db:"f_status"`
	// 超时时间
	ExpiredAt datatypes.MySQLTimestamp `db:"f_expired_at,default='0'" json:"expiredAt"`
	// 支付系统回调原始报文
	RemoteData string `json:"-" db:"f_remote_data,size=65535"`
	datatypes.OperateTime
}

func (m PaymentFlow) BatchFetchByOrderAndStatus(db sqlx.DBExecutor, orderID uint64, status []enums.PaymentStatus) ([]PaymentFlow, error) {
	if orderID == 0 && (status == nil || len(status) == 0) {
		return nil, errors.New("invalid orderID and status")
	}
	table := db.T(m)
	var condition builder.SqlCondition
	if orderID != 0 {
		condition = builder.And(condition, table.F(m.FieldKeyOrderID()).Eq(orderID))
	}
	if status != nil && len(status) > 0 {
		condition = builder.And(condition, table.F(m.FieldKeyStatus()).In(status))
	}

	return m.List(db, condition)
}
