package databases

import (
	fmt "fmt"
	time "time"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
)

func (PromotionFlow) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (PromotionFlow) Indexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"I_payment_flow_id": []string{
			"PaymentFlowID",
		},
		"I_settlement": []string{
			"SettlementID",
		},
		"I_user_id": []string{
			"UserID",
			"CreatedAt",
		},
	}
}

func (PromotionFlow) UniqueIndexUFlowID() string {
	return "U_flow_id"
}

func (PromotionFlow) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_flow_id": []string{
			"FlowID",
			"DeletedAt",
		},
	}
}

func (PromotionFlow) Comments() map[string]string {
	return map[string]string{
		"Amount":          "订单金额",
		"FlowID":          "流水ID",
		"PaymentFlowID":   "关联的支付流水",
		"RefererID":       "奖励来源用户ID",
		"RefererNickName": "奖励来源的用户昵称",
		"SettlementID":    "关联的结算单ID",
		"UserID":          "获得奖励的用户ID",
		"UserNickName":    "获得奖励的用户昵称",
	}
}

var PromotionFlowTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	PromotionFlowTable = Config.DB.Register(&PromotionFlow{})
}

type PromotionFlowIterator struct {
}

func (PromotionFlowIterator) New() interface{} {
	return &PromotionFlow{}
}

func (PromotionFlowIterator) Resolve(v interface{}) *PromotionFlow {
	return v.(*PromotionFlow)
}

func (PromotionFlow) TableName() string {
	return "t_promotion_flow"
}

func (PromotionFlow) ColDescriptions() map[string][]string {
	return map[string][]string{
		"Amount": []string{
			"订单金额",
		},
		"FlowID": []string{
			"流水ID",
		},
		"PaymentFlowID": []string{
			"关联的支付流水",
		},
		"RefererID": []string{
			"奖励来源用户ID",
		},
		"RefererNickName": []string{
			"奖励来源的用户昵称",
		},
		"SettlementID": []string{
			"关联的结算单ID",
		},
		"UserID": []string{
			"获得奖励的用户ID",
		},
		"UserNickName": []string{
			"获得奖励的用户昵称",
		},
	}
}

func (PromotionFlow) FieldKeyID() string {
	return "ID"
}

func (m *PromotionFlow) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return PromotionFlowTable.F(m.FieldKeyID())
}

func (PromotionFlow) FieldKeyFlowID() string {
	return "FlowID"
}

func (m *PromotionFlow) FieldFlowID() *github_com_eden_framework_sqlx_builder.Column {
	return PromotionFlowTable.F(m.FieldKeyFlowID())
}

func (PromotionFlow) FieldKeyUserID() string {
	return "UserID"
}

func (m *PromotionFlow) FieldUserID() *github_com_eden_framework_sqlx_builder.Column {
	return PromotionFlowTable.F(m.FieldKeyUserID())
}

func (PromotionFlow) FieldKeyUserNickName() string {
	return "UserNickName"
}

func (m *PromotionFlow) FieldUserNickName() *github_com_eden_framework_sqlx_builder.Column {
	return PromotionFlowTable.F(m.FieldKeyUserNickName())
}

func (PromotionFlow) FieldKeyRefererID() string {
	return "RefererID"
}

func (m *PromotionFlow) FieldRefererID() *github_com_eden_framework_sqlx_builder.Column {
	return PromotionFlowTable.F(m.FieldKeyRefererID())
}

func (PromotionFlow) FieldKeyRefererNickName() string {
	return "RefererNickName"
}

func (m *PromotionFlow) FieldRefererNickName() *github_com_eden_framework_sqlx_builder.Column {
	return PromotionFlowTable.F(m.FieldKeyRefererNickName())
}

func (PromotionFlow) FieldKeyAmount() string {
	return "Amount"
}

func (m *PromotionFlow) FieldAmount() *github_com_eden_framework_sqlx_builder.Column {
	return PromotionFlowTable.F(m.FieldKeyAmount())
}

func (PromotionFlow) FieldKeyPaymentFlowID() string {
	return "PaymentFlowID"
}

func (m *PromotionFlow) FieldPaymentFlowID() *github_com_eden_framework_sqlx_builder.Column {
	return PromotionFlowTable.F(m.FieldKeyPaymentFlowID())
}

func (PromotionFlow) FieldKeySettlementID() string {
	return "SettlementID"
}

func (m *PromotionFlow) FieldSettlementID() *github_com_eden_framework_sqlx_builder.Column {
	return PromotionFlowTable.F(m.FieldKeySettlementID())
}

func (PromotionFlow) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *PromotionFlow) FieldCreatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return PromotionFlowTable.F(m.FieldKeyCreatedAt())
}

func (PromotionFlow) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *PromotionFlow) FieldUpdatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return PromotionFlowTable.F(m.FieldKeyUpdatedAt())
}

func (PromotionFlow) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *PromotionFlow) FieldDeletedAt() *github_com_eden_framework_sqlx_builder.Column {
	return PromotionFlowTable.F(m.FieldKeyDeletedAt())
}

func (PromotionFlow) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *PromotionFlow) IndexFieldNames() []string {
	return []string{
		"CreatedAt",
		"FlowID",
		"ID",
		"PaymentFlowID",
		"SettlementID",
		"UserID",
	}
}

func (m *PromotionFlow) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
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

func (m *PromotionFlow) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *PromotionFlow) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

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

func (m *PromotionFlow) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("PromotionFlow.DeleteByStruct"),
			),
	)

	return err
}

func (m *PromotionFlow) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("PromotionFlow.FetchByID"),
			),
		m,
	)

	return err
}

func (m *PromotionFlow) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("PromotionFlow.UpdateByIDWithMap"),
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

func (m *PromotionFlow) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *PromotionFlow) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("PromotionFlow.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *PromotionFlow) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("PromotionFlow.DeleteByID"),
			))

	return err
}

func (m *PromotionFlow) SoftDeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("PromotionFlow.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *PromotionFlow) FetchByFlowID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("FlowID").Eq(m.FlowID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("PromotionFlow.FetchByFlowID"),
			),
		m,
	)

	return err
}

func (m *PromotionFlow) UpdateByFlowIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("PromotionFlow.UpdateByFlowIDWithMap"),
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

func (m *PromotionFlow) UpdateByFlowIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByFlowIDWithMap(db, fieldValues)

}

func (m *PromotionFlow) FetchByFlowIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("PromotionFlow.FetchByFlowIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *PromotionFlow) DeleteByFlowID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("FlowID").Eq(m.FlowID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("PromotionFlow.DeleteByFlowID"),
			))

	return err
}

func (m *PromotionFlow) SoftDeleteByFlowID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("PromotionFlow.SoftDeleteByFlowID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *PromotionFlow) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]PromotionFlow, error) {

	list := make([]PromotionFlow, 0)

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("PromotionFlow.List"),
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

func (m *PromotionFlow) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("PromotionFlow.Count"),
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

func (m *PromotionFlow) BatchFetchByCreatedAtList(db github_com_eden_framework_sqlx.DBExecutor, values []github_com_eden_framework_sqlx_datatypes.Timestamp) ([]PromotionFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("CreatedAt").In(values)

	return m.List(db, condition)

}

func (m *PromotionFlow) BatchFetchByFlowIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]PromotionFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("FlowID").In(values)

	return m.List(db, condition)

}

func (m *PromotionFlow) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]PromotionFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *PromotionFlow) BatchFetchByPaymentFlowIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]PromotionFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("PaymentFlowID").In(values)

	return m.List(db, condition)

}

func (m *PromotionFlow) BatchFetchBySettlementIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]PromotionFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("SettlementID").In(values)

	return m.List(db, condition)

}

func (m *PromotionFlow) BatchFetchByUserIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]PromotionFlow, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("UserID").In(values)

	return m.List(db, condition)

}
