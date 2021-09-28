package databases

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
)

//go:generate eden generate model PaymentFlow --database Config.DB --with-comments
//go:generate eden generate tag PaymentFlow --defaults=true
// @def primary ID
// @def unique_index U_flow_id FlowID
// @def index I_order_id OrderID UserID Status
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

func (m PaymentFlow) BatchFetchByOrderAndUserID(db sqlx.DBExecutor, orderID, userID uint64, status enums.PaymentStatus) ([]PaymentFlow, error) {
	table := db.T(m)

	condition := table.F(m.FieldKeyOrderID()).Eq(orderID)
	condition = condition.And(table.F(m.FieldKeyUserID()).Eq(userID))
	if status != enums.PAYMENT_STATUS_UNKNOWN {
		condition = condition.And(table.F(m.FieldKeyStatus()).Eq(status))
	}

	return m.List(db, condition)
}
