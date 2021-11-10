package databases

import (
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
)

//go:generate eden generate model TaskFlow --database Config.DB --with-comments
//go:generate eden generate tag TaskFlow --defaults=true
// @def primary ID
// @def unique_index U_task_flow_id FlowID
// @def index I_type Type
type TaskFlow struct {
	datatypes.PrimaryID
	// 业务ID
	FlowID uint64 `json:"flowID,string" db:"f_task_flow_id"`
	// 任务名称
	Name string `json:"name" db:"f_name"`
	// 任务开始时间
	StartedAt datatypes.MySQLTimestamp `json:"startedAt" db:"f_started_at"`
	// 任务结束时间
	EndedAt datatypes.MySQLTimestamp `json:"endedAt" db:"f_ended_at,null"`
	// 任务执行状态
	Status enums.TaskProcessStatus `json:"status" db:"f_status"`
	// 任务上报信息
	Message string `json:"message" db:"f_message,null,size=65535"`
	// 任务类型
	Type enums.TaskType `json:"type" db:"f_type"`

	datatypes.OperateTime
}
