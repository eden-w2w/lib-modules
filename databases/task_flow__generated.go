package databases

import (
	fmt "fmt"
	time "time"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
	github_com_eden_w2_w_lib_modules_constants_enums "github.com/eden-w2w/lib-modules/constants/enums"
)

func (TaskFlow) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (TaskFlow) Indexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"I_type": []string{
			"Type",
		},
	}
}

func (TaskFlow) UniqueIndexUTaskFlowID() string {
	return "U_task_flow_id"
}

func (TaskFlow) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_task_flow_id": []string{
			"FlowID",
			"DeletedAt",
		},
	}
}

func (TaskFlow) Comments() map[string]string {
	return map[string]string{
		"EndedAt":   "任务结束时间",
		"FlowID":    "业务ID",
		"Message":   "任务上报信息",
		"Name":      "任务名称",
		"StartedAt": "任务开始时间",
		"Status":    "任务执行状态",
		"Type":      "任务类型",
	}
}

var TaskFlowTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	TaskFlowTable = Config.DB.Register(&TaskFlow{})
}

type TaskFlowIterator struct {
}

func (TaskFlowIterator) New() interface{} {
	return &TaskFlow{}
}

func (TaskFlowIterator) Resolve(v interface{}) *TaskFlow {
	return v.(*TaskFlow)
}

func (TaskFlow) TableName() string {
	return "t_task_flow"
}

func (TaskFlow) ColDescriptions() map[string][]string {
	return map[string][]string{
		"EndedAt": []string{
			"任务结束时间",
		},
		"FlowID": []string{
			"业务ID",
		},
		"Message": []string{
			"任务上报信息",
		},
		"Name": []string{
			"任务名称",
		},
		"StartedAt": []string{
			"任务开始时间",
		},
		"Status": []string{
			"任务执行状态",
		},
		"Type": []string{
			"任务类型",
		},
	}
}

func (TaskFlow) FieldKeyID() string {
	return "ID"
}

func (m *TaskFlow) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return TaskFlowTable.F(m.FieldKeyID())
}

func (TaskFlow) FieldKeyFlowID() string {
	return "FlowID"
}

func (m *TaskFlow) FieldFlowID() *github_com_eden_framework_sqlx_builder.Column {
	return TaskFlowTable.F(m.FieldKeyFlowID())
}

func (TaskFlow) FieldKeyName() string {
	return "Name"
}

func (m *TaskFlow) FieldName() *github_com_eden_framework_sqlx_builder.Column {
	return TaskFlowTable.F(m.FieldKeyName())
}

func (TaskFlow) FieldKeyStartedAt() string {
	return "StartedAt"
}

func (m *TaskFlow) FieldStartedAt() *github_com_eden_framework_sqlx_builder.Column {
	return TaskFlowTable.F(m.FieldKeyStartedAt())
}

func (TaskFlow) FieldKeyEndedAt() string {
	return "EndedAt"
}

func (m *TaskFlow) FieldEndedAt() *github_com_eden_framework_sqlx_builder.Column {
	return TaskFlowTable.F(m.FieldKeyEndedAt())
}

func (TaskFlow) FieldKeyStatus() string {
	return "Status"
}

func (m *TaskFlow) FieldStatus() *github_com_eden_framework_sqlx_builder.Column {
	return TaskFlowTable.F(m.FieldKeyStatus())
}

func (TaskFlow) FieldKeyMessage() string {
	return "Message"
}

func (m *TaskFlow) FieldMessage() *github_com_eden_framework_sqlx_builder.Column {
	return TaskFlowTable.F(m.FieldKeyMessage())
}

func (TaskFlow) FieldKeyType() string {
	return "Type"
}

func (m *TaskFlow) FieldType() *github_com_eden_framework_sqlx_builder.Column {
	return TaskFlowTable.F(m.FieldKeyType())
}

func (TaskFlow) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *TaskFlow) FieldCreatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return TaskFlowTable.F(m.FieldKeyCreatedAt())
}

func (TaskFlow) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *TaskFlow) FieldUpdatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return TaskFlowTable.F(m.FieldKeyUpdatedAt())
}

func (TaskFlow) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *TaskFlow) FieldDeletedAt() *github_com_eden_framework_sqlx_builder.Column {
	return TaskFlowTable.F(m.FieldKeyDeletedAt())
}

func (TaskFlow) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *TaskFlow) IndexFieldNames() []string {
	return []string{
		"FlowID",
		"ID",
		"Type",
	}
}

func (m *TaskFlow) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
	table := db.T(m)
	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m)

	conditions := make([]github_com_eden_framework_sqlx_builder.SqlCondition, 0)

	for _, fieldName := range m.IndexFieldNames() {
		if v, exists := fieldValues[fieldName]; exists {
			conditions = append(conditions, table.F(fieldName).Eq(v))
			delete(fieldValues, fieldName)
		}
	}

	if len(conditions) == 0 {
		panic(fmt.Errorf("at least one of field for indexes has value"))
	}

	for fieldName, v := range fieldValues {
		conditions = append(conditions, table.F(fieldName).Eq(v))
	}

	condition := github_com_eden_framework_sqlx_builder.And(conditions...)

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))
	return condition
}

func (m *TaskFlow) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *TaskFlow) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

	if len(updateFields) == 0 {
		panic(fmt.Errorf("must have update fields"))
	}

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, updateFields...)

	delete(fieldValues, "ID")

	table := db.T(m)

	cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

	fields := make(map[string]bool, len(updateFields))
	for _, field := range updateFields {
		fields[field] = true
	}

	for _, fieldNames := range m.UniqueIndexes() {
		for _, field := range fieldNames {
			delete(fields, field)
		}
	}

	if len(fields) == 0 {
		panic(fmt.Errorf("no fields for updates"))
	}

	for field := range fieldValues {
		if !fields[field] {
			delete(fieldValues, field)
		}
	}

	additions := github_com_eden_framework_sqlx_builder.Additions{}

	switch db.Dialect().DriverName() {
	case "mysql":
		additions = append(additions, github_com_eden_framework_sqlx_builder.OnDuplicateKeyUpdate(table.AssignmentsByFieldValues(fieldValues)...))
	case "postgres":
		indexes := m.UniqueIndexes()
		fields := make([]string, 0)
		for _, fs := range indexes {
			fields = append(fields, fs...)
		}
		indexFields, _ := db.T(m).Fields(fields...)

		additions = append(additions,
			github_com_eden_framework_sqlx_builder.OnConflict(indexFields).
				DoUpdateSet(table.AssignmentsByFieldValues(fieldValues)...))
	}

	additions = append(additions, github_com_eden_framework_sqlx_builder.Comment("User.CreateOnDuplicateWithUpdateFields"))

	expr := github_com_eden_framework_sqlx_builder.Insert().Into(table, additions...).Values(cols, vals...)

	_, err := db.ExecExpr(expr)
	return err

}

func (m *TaskFlow) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("TaskFlow.DeleteByStruct"),
			),
	)

	return err
}

func (m *TaskFlow) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("TaskFlow.FetchByID"),
			),
		m,
	)

	return err
}

func (m *TaskFlow) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("TaskFlow.UpdateByIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByID(db)
	}

	return nil

}

func (m *TaskFlow) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *TaskFlow) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("TaskFlow.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *TaskFlow) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("TaskFlow.DeleteByID"),
			))

	return err
}

func (m *TaskFlow) SoftDeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValues{}
	if _, ok := fieldValues["DeletedAt"]; !ok {
		fieldValues["DeletedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("TaskFlow.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *TaskFlow) FetchByFlowID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("FlowID").Eq(m.FlowID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("TaskFlow.FetchByFlowID"),
			),
		m,
	)

	return err
}

func (m *TaskFlow) UpdateByFlowIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("FlowID").Eq(m.FlowID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("TaskFlow.UpdateByFlowIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByFlowID(db)
	}

	return nil

}

func (m *TaskFlow) UpdateByFlowIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByFlowIDWithMap(db, fieldValues)

}

func (m *TaskFlow) FetchByFlowIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("FlowID").Eq(m.FlowID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("TaskFlow.FetchByFlowIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *TaskFlow) DeleteByFlowID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("FlowID").Eq(m.FlowID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("TaskFlow.DeleteByFlowID"),
			))

	return err
}

func (m *TaskFlow) SoftDeleteByFlowID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValues{}
	if _, ok := fieldValues["DeletedAt"]; !ok {
		fieldValues["DeletedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("FlowID").Eq(m.FlowID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("TaskFlow.SoftDeleteByFlowID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *TaskFlow) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]TaskFlow, error) {

	list := make([]TaskFlow, 0)

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("TaskFlow.List"),
	}

	if len(additions) > 0 {
		finalAdditions = append(finalAdditions, additions...)
	}

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(db.T(m), finalAdditions...),
		&list,
	)

	return list, err

}

func (m *TaskFlow) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("TaskFlow.Count"),
	}

	if len(additions) > 0 {
		finalAdditions = append(finalAdditions, additions...)
	}

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(
			github_com_eden_framework_sqlx_builder.Count(),
		).
			From(db.T(m), finalAdditions...),
		&count,
	)

	return count, err

}

func (m *TaskFlow) BatchFetchByFlowIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]TaskFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("FlowID").In(values)

	return m.List(db, condition)

}

func (m *TaskFlow) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]TaskFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *TaskFlow) BatchFetchByTypeList(db github_com_eden_framework_sqlx.DBExecutor, values []github_com_eden_w2_w_lib_modules_constants_enums.TaskType) ([]TaskFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Type").In(values)

	return m.List(db, condition)

}
