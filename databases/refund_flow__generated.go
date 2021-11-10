package databases

import (
	fmt "fmt"
	time "time"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
)

func (RefundFlow) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (RefundFlow) UniqueIndexUFlowID() string {
	return "U_flow_id"
}

func (RefundFlow) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_flow_id": []string{
			"FlowID",
			"DeletedAt",
		},
	}
}

func (RefundFlow) Comments() map[string]string {
	return map[string]string{
		"Account":             "退款账户",
		"Channel":             "退款渠道",
		"FlowID":              "业务ID",
		"PaymentFlowID":       "交易单号",
		"RefundAmount":        "退款总额",
		"RemoteFlowID":        "支付系统退款单号",
		"RemotePaymentFlowID": "支付系统交易单号",
		"Status":              "退款状态",
		"TotalAmount":         "交易总额",
	}
}

var RefundFlowTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	RefundFlowTable = Config.DB.Register(&RefundFlow{})
}

type RefundFlowIterator struct {
}

func (RefundFlowIterator) New() interface{} {
	return &RefundFlow{}
}

func (RefundFlowIterator) Resolve(v interface{}) *RefundFlow {
	return v.(*RefundFlow)
}

func (RefundFlow) TableName() string {
	return "t_refund_flow"
}

func (RefundFlow) ColDescriptions() map[string][]string {
	return map[string][]string{
		"Account": []string{
			"退款账户",
		},
		"Channel": []string{
			"退款渠道",
		},
		"FlowID": []string{
			"业务ID",
		},
		"PaymentFlowID": []string{
			"交易单号",
		},
		"RefundAmount": []string{
			"退款总额",
		},
		"RemoteFlowID": []string{
			"支付系统退款单号",
		},
		"RemotePaymentFlowID": []string{
			"支付系统交易单号",
		},
		"Status": []string{
			"退款状态",
		},
		"TotalAmount": []string{
			"交易总额",
		},
	}
}

func (RefundFlow) FieldKeyID() string {
	return "ID"
}

func (m *RefundFlow) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return RefundFlowTable.F(m.FieldKeyID())
}

func (RefundFlow) FieldKeyFlowID() string {
	return "FlowID"
}

func (m *RefundFlow) FieldFlowID() *github_com_eden_framework_sqlx_builder.Column {
	return RefundFlowTable.F(m.FieldKeyFlowID())
}

func (RefundFlow) FieldKeyRemoteFlowID() string {
	return "RemoteFlowID"
}

func (m *RefundFlow) FieldRemoteFlowID() *github_com_eden_framework_sqlx_builder.Column {
	return RefundFlowTable.F(m.FieldKeyRemoteFlowID())
}

func (RefundFlow) FieldKeyPaymentFlowID() string {
	return "PaymentFlowID"
}

func (m *RefundFlow) FieldPaymentFlowID() *github_com_eden_framework_sqlx_builder.Column {
	return RefundFlowTable.F(m.FieldKeyPaymentFlowID())
}

func (RefundFlow) FieldKeyRemotePaymentFlowID() string {
	return "RemotePaymentFlowID"
}

func (m *RefundFlow) FieldRemotePaymentFlowID() *github_com_eden_framework_sqlx_builder.Column {
	return RefundFlowTable.F(m.FieldKeyRemotePaymentFlowID())
}

func (RefundFlow) FieldKeyChannel() string {
	return "Channel"
}

func (m *RefundFlow) FieldChannel() *github_com_eden_framework_sqlx_builder.Column {
	return RefundFlowTable.F(m.FieldKeyChannel())
}

func (RefundFlow) FieldKeyAccount() string {
	return "Account"
}

func (m *RefundFlow) FieldAccount() *github_com_eden_framework_sqlx_builder.Column {
	return RefundFlowTable.F(m.FieldKeyAccount())
}

func (RefundFlow) FieldKeyStatus() string {
	return "Status"
}

func (m *RefundFlow) FieldStatus() *github_com_eden_framework_sqlx_builder.Column {
	return RefundFlowTable.F(m.FieldKeyStatus())
}

func (RefundFlow) FieldKeyTotalAmount() string {
	return "TotalAmount"
}

func (m *RefundFlow) FieldTotalAmount() *github_com_eden_framework_sqlx_builder.Column {
	return RefundFlowTable.F(m.FieldKeyTotalAmount())
}

func (RefundFlow) FieldKeyRefundAmount() string {
	return "RefundAmount"
}

func (m *RefundFlow) FieldRefundAmount() *github_com_eden_framework_sqlx_builder.Column {
	return RefundFlowTable.F(m.FieldKeyRefundAmount())
}

func (RefundFlow) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *RefundFlow) FieldCreatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return RefundFlowTable.F(m.FieldKeyCreatedAt())
}

func (RefundFlow) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *RefundFlow) FieldUpdatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return RefundFlowTable.F(m.FieldKeyUpdatedAt())
}

func (RefundFlow) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *RefundFlow) FieldDeletedAt() *github_com_eden_framework_sqlx_builder.Column {
	return RefundFlowTable.F(m.FieldKeyDeletedAt())
}

func (RefundFlow) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *RefundFlow) IndexFieldNames() []string {
	return []string{
		"FlowID",
		"ID",
	}
}

func (m *RefundFlow) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
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

func (m *RefundFlow) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *RefundFlow) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

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

func (m *RefundFlow) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("RefundFlow.DeleteByStruct"),
			),
	)

	return err
}

func (m *RefundFlow) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("RefundFlow.FetchByID"),
			),
		m,
	)

	return err
}

func (m *RefundFlow) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("RefundFlow.UpdateByIDWithMap"),
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

func (m *RefundFlow) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *RefundFlow) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("RefundFlow.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *RefundFlow) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("RefundFlow.DeleteByID"),
			))

	return err
}

func (m *RefundFlow) SoftDeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("RefundFlow.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *RefundFlow) FetchByFlowID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("FlowID").Eq(m.FlowID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("RefundFlow.FetchByFlowID"),
			),
		m,
	)

	return err
}

func (m *RefundFlow) UpdateByFlowIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("RefundFlow.UpdateByFlowIDWithMap"),
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

func (m *RefundFlow) UpdateByFlowIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByFlowIDWithMap(db, fieldValues)

}

func (m *RefundFlow) FetchByFlowIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("RefundFlow.FetchByFlowIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *RefundFlow) DeleteByFlowID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("FlowID").Eq(m.FlowID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("RefundFlow.DeleteByFlowID"),
			))

	return err
}

func (m *RefundFlow) SoftDeleteByFlowID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("RefundFlow.SoftDeleteByFlowID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *RefundFlow) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]RefundFlow, error) {

	list := make([]RefundFlow, 0)

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("RefundFlow.List"),
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

func (m *RefundFlow) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("RefundFlow.Count"),
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

func (m *RefundFlow) BatchFetchByFlowIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]RefundFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("FlowID").In(values)

	return m.List(db, condition)

}

func (m *RefundFlow) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]RefundFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}
