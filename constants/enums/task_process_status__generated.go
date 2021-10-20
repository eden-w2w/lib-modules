package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidTaskProcessStatus = errors.New("invalid TaskProcessStatus")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("TaskProcessStatus", map[string]string{
		"FAIL":     "失败",
		"COMPLETE": "已完成",
		"PROCESS":  "执行中",
		"CREATED":  "待执行",
	})
}

func ParseTaskProcessStatusFromString(s string) (TaskProcessStatus, error) {
	switch s {
	case "":
		return TASK_PROCESS_STATUS_UNKNOWN, nil
	case "FAIL":
		return TASK_PROCESS_STATUS__FAIL, nil
	case "COMPLETE":
		return TASK_PROCESS_STATUS__COMPLETE, nil
	case "PROCESS":
		return TASK_PROCESS_STATUS__PROCESS, nil
	case "CREATED":
		return TASK_PROCESS_STATUS__CREATED, nil
	}
	return TASK_PROCESS_STATUS_UNKNOWN, InvalidTaskProcessStatus
}

func ParseTaskProcessStatusFromLabelString(s string) (TaskProcessStatus, error) {
	switch s {
	case "":
		return TASK_PROCESS_STATUS_UNKNOWN, nil
	case "失败":
		return TASK_PROCESS_STATUS__FAIL, nil
	case "已完成":
		return TASK_PROCESS_STATUS__COMPLETE, nil
	case "执行中":
		return TASK_PROCESS_STATUS__PROCESS, nil
	case "待执行":
		return TASK_PROCESS_STATUS__CREATED, nil
	}
	return TASK_PROCESS_STATUS_UNKNOWN, InvalidTaskProcessStatus
}

func (TaskProcessStatus) EnumType() string {
	return "TaskProcessStatus"
}

func (TaskProcessStatus) Enums() map[int][]string {
	return map[int][]string{
		int(TASK_PROCESS_STATUS__FAIL):     {"FAIL", "失败"},
		int(TASK_PROCESS_STATUS__COMPLETE): {"COMPLETE", "已完成"},
		int(TASK_PROCESS_STATUS__PROCESS):  {"PROCESS", "执行中"},
		int(TASK_PROCESS_STATUS__CREATED):  {"CREATED", "待执行"},
	}
}

func (v TaskProcessStatus) String() string {
	switch v {
	case TASK_PROCESS_STATUS_UNKNOWN:
		return ""
	case TASK_PROCESS_STATUS__FAIL:
		return "FAIL"
	case TASK_PROCESS_STATUS__COMPLETE:
		return "COMPLETE"
	case TASK_PROCESS_STATUS__PROCESS:
		return "PROCESS"
	case TASK_PROCESS_STATUS__CREATED:
		return "CREATED"
	}
	return "UNKNOWN"
}

func (v TaskProcessStatus) Label() string {
	switch v {
	case TASK_PROCESS_STATUS_UNKNOWN:
		return ""
	case TASK_PROCESS_STATUS__FAIL:
		return "失败"
	case TASK_PROCESS_STATUS__COMPLETE:
		return "已完成"
	case TASK_PROCESS_STATUS__PROCESS:
		return "执行中"
	case TASK_PROCESS_STATUS__CREATED:
		return "待执行"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*TaskProcessStatus)(nil)

func (v TaskProcessStatus) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidTaskProcessStatus
	}
	return []byte(str), nil
}

func (v *TaskProcessStatus) UnmarshalText(data []byte) (err error) {
	*v, err = ParseTaskProcessStatusFromString(string(bytes.ToUpper(data)))
	return
}
