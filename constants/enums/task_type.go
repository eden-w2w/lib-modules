package enums

//go:generate eden generate enum --type-name=TaskType
// api:enum
type TaskType uint8

// 任务类型
const (
	TASK_TYPE_UNKNOWN         TaskType = iota
	TASK_TYPE__SETTLEMENT              // 结算任务
	TASK_TYPE__RECONCILIATION          // 对账任务
)
