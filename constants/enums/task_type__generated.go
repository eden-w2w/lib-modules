package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidTaskType = errors.New("invalid TaskType")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("TaskType", map[string]string{
		"RECONCILIATION": "对账任务",
		"SETTLEMENT":     "结算任务",
	})
}

func ParseTaskTypeFromString(s string) (TaskType, error) {
	switch s {
	case "":
		return TASK_TYPE_UNKNOWN, nil
	case "RECONCILIATION":
		return TASK_TYPE__RECONCILIATION, nil
	case "SETTLEMENT":
		return TASK_TYPE__SETTLEMENT, nil
	}
	return TASK_TYPE_UNKNOWN, InvalidTaskType
}

func ParseTaskTypeFromLabelString(s string) (TaskType, error) {
	switch s {
	case "":
		return TASK_TYPE_UNKNOWN, nil
	case "对账任务":
		return TASK_TYPE__RECONCILIATION, nil
	case "结算任务":
		return TASK_TYPE__SETTLEMENT, nil
	}
	return TASK_TYPE_UNKNOWN, InvalidTaskType
}

func (TaskType) EnumType() string {
	return "TaskType"
}

func (TaskType) Enums() map[int][]string {
	return map[int][]string{
		int(TASK_TYPE__RECONCILIATION): {"RECONCILIATION", "对账任务"},
		int(TASK_TYPE__SETTLEMENT):     {"SETTLEMENT", "结算任务"},
	}
}

func (v TaskType) String() string {
	switch v {
	case TASK_TYPE_UNKNOWN:
		return ""
	case TASK_TYPE__RECONCILIATION:
		return "RECONCILIATION"
	case TASK_TYPE__SETTLEMENT:
		return "SETTLEMENT"
	}
	return "UNKNOWN"
}

func (v TaskType) Label() string {
	switch v {
	case TASK_TYPE_UNKNOWN:
		return ""
	case TASK_TYPE__RECONCILIATION:
		return "对账任务"
	case TASK_TYPE__SETTLEMENT:
		return "结算任务"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*TaskType)(nil)

func (v TaskType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidTaskType
	}
	return []byte(str), nil
}

func (v *TaskType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseTaskTypeFromString(string(bytes.ToUpper(data)))
	return
}
