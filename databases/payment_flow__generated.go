package databases

import (
	fmt "fmt"
	time "time"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
	github_com_eden_w2_w_srv_w2_w_internal_contants_enums "github.com/eden-w2w/lib-modules/constants/enums"
)

func (PaymentFlow) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (PaymentFlow) Indexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"I_expire": []string{
			"ExpiredAt",
		},
		"I_order_id": []string{
			"OrderID",
			"UserID",
			"Status",
		},
	}
}

func (PaymentFlow) UniqueIndexUFlowID() string {
	return "U_flow_id"
}

func (PaymentFlow) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_flow_id": []string{
			"FlowID",
			"DeletedAt",
		},
	}
}

func (PaymentFlow) Comments() map[string]string {
	return map[string]string{
		"Amount":        "支付金额",
		"ExpiredAt":     "超时时间",
		"FlowID":        "流水ID",
		"OrderID":       "关联订单号",
		"PaymentMethod": "支付方式",
		"RemoteData":    "支付系统回调原始报文",
		"RemoteFlowID":  "支付系统流水号",
		"Status":        "支付状态",
		"UserID":        "用户ID",
	}
}

var PaymentFlowTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	PaymentFlowTable = Config.DB.Register(&PaymentFlow{})
}

type PaymentFlowIterator struct {
}

func (PaymentFlowIterator) New() interface{} {
	return &PaymentFlow{}
}

func (PaymentFlowIterator) Resolve(v interface{}) *PaymentFlow {
	return v.(*PaymentFlow)
}

func (PaymentFlow) TableName() string {
	return "t_payment_flow"
}

func (PaymentFlow) ColDescriptions() map[string][]string {
	return map[string][]string{
		"Amount": []string{
			"支付金额",
		},
		"ExpiredAt": []string{
			"超时时间",
		},
		"FlowID": []string{
			"流水ID",
		},
		"OrderID": []string{
			"关联订单号",
		},
		"PaymentMethod": []string{
			"支付方式",
		},
		"RemoteData": []string{
			"支付系统回调原始报文",
		},
		"RemoteFlowID": []string{
			"支付系统流水号",
		},
		"Status": []string{
			"支付状态",
		},
		"UserID": []string{
			"用户ID",
		},
	}
}

func (PaymentFlow) FieldKeyID() string {
	return "ID"
}

func (m *PaymentFlow) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return PaymentFlowTable.F(m.FieldKeyID())
}

func (PaymentFlow) FieldKeyFlowID() string {
	return "FlowID"
}

func (m *PaymentFlow) FieldFlowID() *github_com_eden_framework_sqlx_builder.Column {
	return PaymentFlowTable.F(m.FieldKeyFlowID())
}

func (PaymentFlow) FieldKeyUserID() string {
	return "UserID"
}

func (m *PaymentFlow) FieldUserID() *github_com_eden_framework_sqlx_builder.Column {
	return PaymentFlowTable.F(m.FieldKeyUserID())
}

func (PaymentFlow) FieldKeyOrderID() string {
	return "OrderID"
}

func (m *PaymentFlow) FieldOrderID() *github_com_eden_framework_sqlx_builder.Column {
	return PaymentFlowTable.F(m.FieldKeyOrderID())
}

func (PaymentFlow) FieldKeyAmount() string {
	return "Amount"
}

func (m *PaymentFlow) FieldAmount() *github_com_eden_framework_sqlx_builder.Column {
	return PaymentFlowTable.F(m.FieldKeyAmount())
}

func (PaymentFlow) FieldKeyPaymentMethod() string {
	return "PaymentMethod"
}

func (m *PaymentFlow) FieldPaymentMethod() *github_com_eden_framework_sqlx_builder.Column {
	return PaymentFlowTable.F(m.FieldKeyPaymentMethod())
}

func (PaymentFlow) FieldKeyRemoteFlowID() string {
	return "RemoteFlowID"
}

func (m *PaymentFlow) FieldRemoteFlowID() *github_com_eden_framework_sqlx_builder.Column {
	return PaymentFlowTable.F(m.FieldKeyRemoteFlowID())
}

func (PaymentFlow) FieldKeyStatus() string {
	return "Status"
}

func (m *PaymentFlow) FieldStatus() *github_com_eden_framework_sqlx_builder.Column {
	return PaymentFlowTable.F(m.FieldKeyStatus())
}

func (PaymentFlow) FieldKeyExpiredAt() string {
	return "ExpiredAt"
}

func (m *PaymentFlow) FieldExpiredAt() *github_com_eden_framework_sqlx_builder.Column {
	return PaymentFlowTable.F(m.FieldKeyExpiredAt())
}

func (PaymentFlow) FieldKeyRemoteData() string {
	return "RemoteData"
}

func (m *PaymentFlow) FieldRemoteData() *github_com_eden_framework_sqlx_builder.Column {
	return PaymentFlowTable.F(m.FieldKeyRemoteData())
}

func (PaymentFlow) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *PaymentFlow) FieldCreatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return PaymentFlowTable.F(m.FieldKeyCreatedAt())
}

func (PaymentFlow) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *PaymentFlow) FieldUpdatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return PaymentFlowTable.F(m.FieldKeyUpdatedAt())
}

func (PaymentFlow) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *PaymentFlow) FieldDeletedAt() *github_com_eden_framework_sqlx_builder.Column {
	return PaymentFlowTable.F(m.FieldKeyDeletedAt())
}

func (PaymentFlow) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *PaymentFlow) IndexFieldNames() []string {
	return []string{
		"ExpiredAt",
		"FlowID",
		"ID",
		"OrderID",
		"Status",
		"UserID",
	}
}

func (m *PaymentFlow) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
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

func (m *PaymentFlow) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *PaymentFlow) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

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

func (m *PaymentFlow) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("PaymentFlow.DeleteByStruct"),
			),
	)

	return err
}

func (m *PaymentFlow) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("PaymentFlow.FetchByID"),
			),
		m,
	)

	return err
}

func (m *PaymentFlow) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("PaymentFlow.UpdateByIDWithMap"),
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

func (m *PaymentFlow) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *PaymentFlow) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("PaymentFlow.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *PaymentFlow) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("PaymentFlow.DeleteByID"),
			))

	return err
}

func (m *PaymentFlow) SoftDeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("PaymentFlow.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *PaymentFlow) FetchByFlowID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("FlowID").Eq(m.FlowID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("PaymentFlow.FetchByFlowID"),
			),
		m,
	)

	return err
}

func (m *PaymentFlow) UpdateByFlowIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("PaymentFlow.UpdateByFlowIDWithMap"),
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

func (m *PaymentFlow) UpdateByFlowIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByFlowIDWithMap(db, fieldValues)

}

func (m *PaymentFlow) FetchByFlowIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("PaymentFlow.FetchByFlowIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *PaymentFlow) DeleteByFlowID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("FlowID").Eq(m.FlowID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("PaymentFlow.DeleteByFlowID"),
			))

	return err
}

func (m *PaymentFlow) SoftDeleteByFlowID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("PaymentFlow.SoftDeleteByFlowID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *PaymentFlow) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]PaymentFlow, error) {

	list := make([]PaymentFlow, 0)

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("PaymentFlow.List"),
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

func (m *PaymentFlow) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("PaymentFlow.Count"),
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

func (m *PaymentFlow) BatchFetchByExpiredAtList(db github_com_eden_framework_sqlx.DBExecutor, values []github_com_eden_framework_sqlx_datatypes.Timestamp) ([]PaymentFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ExpiredAt").In(values)

	return m.List(db, condition)

}

func (m *PaymentFlow) BatchFetchByFlowIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]PaymentFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("FlowID").In(values)

	return m.List(db, condition)

}

func (m *PaymentFlow) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]PaymentFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *PaymentFlow) BatchFetchByOrderIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]PaymentFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("OrderID").In(values)

	return m.List(db, condition)

}

func (m *PaymentFlow) BatchFetchByStatusList(db github_com_eden_framework_sqlx.DBExecutor, values []github_com_eden_w2_w_srv_w2_w_internal_contants_enums.PaymentStatus) ([]PaymentFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Status").In(values)

	return m.List(db, condition)

}

func (m *PaymentFlow) BatchFetchByUserIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]PaymentFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("UserID").In(values)

	return m.List(db, condition)

}
