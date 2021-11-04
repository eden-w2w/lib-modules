package task_flow

import (
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules"
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

type GetTaskParams struct {
	// 开始时间大于等于
	StartTimeGte datatypes.MySQLTimestamp `name:"startTimeGte" in:"query" default:""`
	// 开始时间小于
	StartTimeLt datatypes.MySQLTimestamp `name:"startTimeLt" in:"query" default:""`
	// 任务执行状态
	Status enums.TaskProcessStatus `json:"status" in:"query" default:""`
	modules.Pagination
}

func (p GetTaskParams) Conditions() builder.SqlCondition {
	var condition builder.SqlCondition
	model := databases.TaskFlow{}

	if p.Status != enums.TASK_PROCESS_STATUS_UNKNOWN {
		condition = builder.And(condition, model.FieldStatus().Eq(p.Status))
	}
	if p.StartTimeGte != datatypes.TimestampZero {
		condition = builder.And(condition, model.FieldCreatedAt().Gte(p.StartTimeGte))
	}
	if p.StartTimeLt != datatypes.TimestampZero {
		condition = builder.And(condition, model.FieldCreatedAt().Lt(p.StartTimeLt))
	}

	return condition
}

func (p GetTaskParams) Additions() []builder.Addition {
	var additions = make([]builder.Addition, 0)

	if p.Size != 0 {
		limit := builder.Limit(int64(p.Size))
		if p.Offset != 0 {
			limit = limit.Offset(int64(p.Offset))
		}
		additions = append(additions, limit)
	}

	additions = append(additions, builder.OrderBy(builder.DescOrder((&databases.Order{}).FieldCreatedAt())))

	return additions
}
