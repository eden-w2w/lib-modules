package databases

import (
	fmt "fmt"
	time "time"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
	github_com_eden_w2_w_lib_modules_constants_enums "github.com/eden-w2w/lib-modules/constants/enums"
)

func (BookingFlow) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (BookingFlow) Indexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"I_default": []string{
			"GoodsID",
			"Status",
			"Type",
		},
		"I_time": []string{
			"StartTime",
			"EndTime",
		},
	}
}

func (BookingFlow) UniqueIndexUBookingFlowID() string {
	return "U_booking_flow_id"
}

func (BookingFlow) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_booking_flow_id": []string{
			"FlowID",
			"DeletedAt",
		},
	}
}

func (BookingFlow) Comments() map[string]string {
	return map[string]string{
		"EndTime":   "预售结束时间",
		"FlowID":    "业务ID",
		"GoodsID":   "商品ID",
		"Limit":     "预售限量",
		"Sales":     "预售销量",
		"StartTime": "预售开始时间",
		"Status":    "预售状态",
		"Type":      "预售模式",
	}
}

var BookingFlowTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	BookingFlowTable = Config.DB.Register(&BookingFlow{})
}

type BookingFlowIterator struct {
}

func (BookingFlowIterator) New() interface{} {
	return &BookingFlow{}
}

func (BookingFlowIterator) Resolve(v interface{}) *BookingFlow {
	return v.(*BookingFlow)
}

func (BookingFlow) TableName() string {
	return "t_booking_flow"
}

func (BookingFlow) ColDescriptions() map[string][]string {
	return map[string][]string{
		"EndTime": []string{
			"预售结束时间",
		},
		"FlowID": []string{
			"业务ID",
		},
		"GoodsID": []string{
			"商品ID",
		},
		"Limit": []string{
			"预售限量",
		},
		"Sales": []string{
			"预售销量",
		},
		"StartTime": []string{
			"预售开始时间",
		},
		"Status": []string{
			"预售状态",
		},
		"Type": []string{
			"预售模式",
		},
	}
}

func (BookingFlow) FieldKeyID() string {
	return "ID"
}

func (m *BookingFlow) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return BookingFlowTable.F(m.FieldKeyID())
}

func (BookingFlow) FieldKeyFlowID() string {
	return "FlowID"
}

func (m *BookingFlow) FieldFlowID() *github_com_eden_framework_sqlx_builder.Column {
	return BookingFlowTable.F(m.FieldKeyFlowID())
}

func (BookingFlow) FieldKeyGoodsID() string {
	return "GoodsID"
}

func (m *BookingFlow) FieldGoodsID() *github_com_eden_framework_sqlx_builder.Column {
	return BookingFlowTable.F(m.FieldKeyGoodsID())
}

func (BookingFlow) FieldKeySales() string {
	return "Sales"
}

func (m *BookingFlow) FieldSales() *github_com_eden_framework_sqlx_builder.Column {
	return BookingFlowTable.F(m.FieldKeySales())
}

func (BookingFlow) FieldKeyLimit() string {
	return "Limit"
}

func (m *BookingFlow) FieldLimit() *github_com_eden_framework_sqlx_builder.Column {
	return BookingFlowTable.F(m.FieldKeyLimit())
}

func (BookingFlow) FieldKeyType() string {
	return "Type"
}

func (m *BookingFlow) FieldType() *github_com_eden_framework_sqlx_builder.Column {
	return BookingFlowTable.F(m.FieldKeyType())
}

func (BookingFlow) FieldKeyStatus() string {
	return "Status"
}

func (m *BookingFlow) FieldStatus() *github_com_eden_framework_sqlx_builder.Column {
	return BookingFlowTable.F(m.FieldKeyStatus())
}

func (BookingFlow) FieldKeyStartTime() string {
	return "StartTime"
}

func (m *BookingFlow) FieldStartTime() *github_com_eden_framework_sqlx_builder.Column {
	return BookingFlowTable.F(m.FieldKeyStartTime())
}

func (BookingFlow) FieldKeyEndTime() string {
	return "EndTime"
}

func (m *BookingFlow) FieldEndTime() *github_com_eden_framework_sqlx_builder.Column {
	return BookingFlowTable.F(m.FieldKeyEndTime())
}

func (BookingFlow) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *BookingFlow) FieldCreatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return BookingFlowTable.F(m.FieldKeyCreatedAt())
}

func (BookingFlow) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *BookingFlow) FieldUpdatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return BookingFlowTable.F(m.FieldKeyUpdatedAt())
}

func (BookingFlow) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *BookingFlow) FieldDeletedAt() *github_com_eden_framework_sqlx_builder.Column {
	return BookingFlowTable.F(m.FieldKeyDeletedAt())
}

func (BookingFlow) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *BookingFlow) IndexFieldNames() []string {
	return []string{
		"EndTime",
		"FlowID",
		"GoodsID",
		"ID",
		"StartTime",
		"Status",
		"Type",
	}
}

func (m *BookingFlow) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
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

func (m *BookingFlow) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *BookingFlow) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

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

func (m *BookingFlow) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("BookingFlow.DeleteByStruct"),
			),
	)

	return err
}

func (m *BookingFlow) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("BookingFlow.FetchByID"),
			),
		m,
	)

	return err
}

func (m *BookingFlow) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("BookingFlow.UpdateByIDWithMap"),
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

func (m *BookingFlow) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *BookingFlow) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("BookingFlow.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *BookingFlow) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("BookingFlow.DeleteByID"),
			))

	return err
}

func (m *BookingFlow) SoftDeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("BookingFlow.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *BookingFlow) FetchByFlowID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("FlowID").Eq(m.FlowID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("BookingFlow.FetchByFlowID"),
			),
		m,
	)

	return err
}

func (m *BookingFlow) UpdateByFlowIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("BookingFlow.UpdateByFlowIDWithMap"),
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

func (m *BookingFlow) UpdateByFlowIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByFlowIDWithMap(db, fieldValues)

}

func (m *BookingFlow) FetchByFlowIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("BookingFlow.FetchByFlowIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *BookingFlow) DeleteByFlowID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("FlowID").Eq(m.FlowID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("BookingFlow.DeleteByFlowID"),
			))

	return err
}

func (m *BookingFlow) SoftDeleteByFlowID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("BookingFlow.SoftDeleteByFlowID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *BookingFlow) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]BookingFlow, error) {

	list := make([]BookingFlow, 0)

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("BookingFlow.List"),
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

func (m *BookingFlow) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("BookingFlow.Count"),
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

func (m *BookingFlow) BatchFetchByEndTimeList(db github_com_eden_framework_sqlx.DBExecutor, values []github_com_eden_framework_sqlx_datatypes.Timestamp) ([]BookingFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("EndTime").In(values)

	return m.List(db, condition)

}

func (m *BookingFlow) BatchFetchByFlowIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]BookingFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("FlowID").In(values)

	return m.List(db, condition)

}

func (m *BookingFlow) BatchFetchByGoodsIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]BookingFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("GoodsID").In(values)

	return m.List(db, condition)

}

func (m *BookingFlow) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]BookingFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *BookingFlow) BatchFetchByStartTimeList(db github_com_eden_framework_sqlx.DBExecutor, values []github_com_eden_framework_sqlx_datatypes.Timestamp) ([]BookingFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("StartTime").In(values)

	return m.List(db, condition)

}

func (m *BookingFlow) BatchFetchByStatusList(db github_com_eden_framework_sqlx.DBExecutor, values []github_com_eden_w2_w_lib_modules_constants_enums.BookingStatus) ([]BookingFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Status").In(values)

	return m.List(db, condition)

}

func (m *BookingFlow) BatchFetchByTypeList(db github_com_eden_framework_sqlx.DBExecutor, values []github_com_eden_w2_w_lib_modules_constants_enums.BookingType) ([]BookingFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Type").In(values)

	return m.List(db, condition)

}
