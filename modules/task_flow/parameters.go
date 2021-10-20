package task_flow

import (
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
)

type CreateTaskFlowParams struct {
	// 任务名称
	Name string `json:"name"`
}

type UpdateTaskParams struct {
	// 任务开始时间
	StartedAt datatypes.MySQLTimestamp `json:"startedAt" default:""`
	// 任务结束时间
	EndedAt datatypes.MySQLTimestamp `json:"endedAt" default:""`
	// 任务执行状态
	Status enums.TaskProcessStatus `json:"status" default:""`
	// 任务上报信息
	Message string `json:"message" default:""`
}
