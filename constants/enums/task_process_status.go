package enums

//go:generate eden generate enum --type-name=TaskProcessStatus
// api:enum
type TaskProcessStatus uint8

// 任务执行状态
const (
	TASK_PROCESS_STATUS_UNKNOWN   TaskProcessStatus = iota
	TASK_PROCESS_STATUS__CREATED                    // 待执行
	TASK_PROCESS_STATUS__PROCESS                    // 执行中
	TASK_PROCESS_STATUS__COMPLETE                   // 已完成
	TASK_PROCESS_STATUS__FAIL                       // 失败
)
