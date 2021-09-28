package databases

import (
	fmt "fmt"
	time "time"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
)

func (OrderLogistics) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (OrderLogistics) Indexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"I_number": []string{
			"CourierNumber",
		},
	}
}

func (OrderLogistics) UniqueIndexULogisticsID() string {
	return "U_logistics_id"
}

func (OrderLogistics) UniqueIndexUOrderID() string {
	return "U_order_id"
}

func (OrderLogistics) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_logistics_id": []string{
			"LogisticsID",
			"DeletedAt",
		},
		"U_order_id": []string{
			"OrderID",
			"DeletedAt",
		},
	}
}

func (OrderLogistics) Comments() map[string]string {
	return map[string]string{
		"CourierCompany": "快递公司",
		"CourierNumber":  "快递单号",
		"LogisticsID":    "业务ID",
		"Mobile":         "联系电话",
		"OrderID":        "订单号",
		"Recipients":     "收件人",
		"ShippingAddr":   "收货地址",
	}
}

var OrderLogisticsTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	OrderLogisticsTable = Config.DB.Register(&OrderLogistics{})
}

type OrderLogisticsIterator struct {
}

func (OrderLogisticsIterator) New() interface{} {
	return &OrderLogistics{}
}

func (OrderLogisticsIterator) Resolve(v interface{}) *OrderLogistics {
	return v.(*OrderLogistics)
}

func (OrderLogistics) TableName() string {
	return "t_order_logistics"
}

func (OrderLogistics) ColDescriptions() map[string][]string {
	return map[string][]string{
		"CourierCompany": []string{
			"快递公司",
		},
		"CourierNumber": []string{
			"快递单号",
		},
		"LogisticsID": []string{
			"业务ID",
		},
		"Mobile": []string{
			"联系电话",
		},
		"OrderID": []string{
			"订单号",
		},
		"Recipients": []string{
			"收件人",
		},
		"ShippingAddr": []string{
			"收货地址",
		},
	}
}

func (OrderLogistics) FieldKeyID() string {
	return "ID"
}

func (m *OrderLogistics) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return OrderLogisticsTable.F(m.FieldKeyID())
}

func (OrderLogistics) FieldKeyLogisticsID() string {
	return "LogisticsID"
}

func (m *OrderLogistics) FieldLogisticsID() *github_com_eden_framework_sqlx_builder.Column {
	return OrderLogisticsTable.F(m.FieldKeyLogisticsID())
}

func (OrderLogistics) FieldKeyOrderID() string {
	return "OrderID"
}

func (m *OrderLogistics) FieldOrderID() *github_com_eden_framework_sqlx_builder.Column {
	return OrderLogisticsTable.F(m.FieldKeyOrderID())
}

func (OrderLogistics) FieldKeyRecipients() string {
	return "Recipients"
}

func (m *OrderLogistics) FieldRecipients() *github_com_eden_framework_sqlx_builder.Column {
	return OrderLogisticsTable.F(m.FieldKeyRecipients())
}

func (OrderLogistics) FieldKeyShippingAddr() string {
	return "ShippingAddr"
}

func (m *OrderLogistics) FieldShippingAddr() *github_com_eden_framework_sqlx_builder.Column {
	return OrderLogisticsTable.F(m.FieldKeyShippingAddr())
}

func (OrderLogistics) FieldKeyMobile() string {
	return "Mobile"
}

func (m *OrderLogistics) FieldMobile() *github_com_eden_framework_sqlx_builder.Column {
	return OrderLogisticsTable.F(m.FieldKeyMobile())
}

func (OrderLogistics) FieldKeyCourierCompany() string {
	return "CourierCompany"
}

func (m *OrderLogistics) FieldCourierCompany() *github_com_eden_framework_sqlx_builder.Column {
	return OrderLogisticsTable.F(m.FieldKeyCourierCompany())
}

func (OrderLogistics) FieldKeyCourierNumber() string {
	return "CourierNumber"
}

func (m *OrderLogistics) FieldCourierNumber() *github_com_eden_framework_sqlx_builder.Column {
	return OrderLogisticsTable.F(m.FieldKeyCourierNumber())
}

func (OrderLogistics) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *OrderLogistics) FieldCreatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return OrderLogisticsTable.F(m.FieldKeyCreatedAt())
}

func (OrderLogistics) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *OrderLogistics) FieldUpdatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return OrderLogisticsTable.F(m.FieldKeyUpdatedAt())
}

func (OrderLogistics) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *OrderLogistics) FieldDeletedAt() *github_com_eden_framework_sqlx_builder.Column {
	return OrderLogisticsTable.F(m.FieldKeyDeletedAt())
}

func (OrderLogistics) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *OrderLogistics) IndexFieldNames() []string {
	return []string{
		"CourierNumber",
		"ID",
		"LogisticsID",
		"OrderID",
	}
}

func (m *OrderLogistics) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
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

func (m *OrderLogistics) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *OrderLogistics) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

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

func (m *OrderLogistics) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.DeleteByStruct"),
			),
	)

	return err
}

func (m *OrderLogistics) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.FetchByID"),
			),
		m,
	)

	return err
}

func (m *OrderLogistics) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.UpdateByIDWithMap"),
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

func (m *OrderLogistics) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *OrderLogistics) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *OrderLogistics) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.DeleteByID"),
			))

	return err
}

func (m *OrderLogistics) SoftDeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *OrderLogistics) FetchByLogisticsID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("LogisticsID").Eq(m.LogisticsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.FetchByLogisticsID"),
			),
		m,
	)

	return err
}

func (m *OrderLogistics) UpdateByLogisticsIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("LogisticsID").Eq(m.LogisticsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.UpdateByLogisticsIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByLogisticsID(db)
	}

	return nil

}

func (m *OrderLogistics) UpdateByLogisticsIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByLogisticsIDWithMap(db, fieldValues)

}

func (m *OrderLogistics) FetchByLogisticsIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("LogisticsID").Eq(m.LogisticsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.FetchByLogisticsIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *OrderLogistics) DeleteByLogisticsID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("LogisticsID").Eq(m.LogisticsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.DeleteByLogisticsID"),
			))

	return err
}

func (m *OrderLogistics) SoftDeleteByLogisticsID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
					table.F("LogisticsID").Eq(m.LogisticsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.SoftDeleteByLogisticsID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *OrderLogistics) FetchByOrderID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("OrderID").Eq(m.OrderID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.FetchByOrderID"),
			),
		m,
	)

	return err
}

func (m *OrderLogistics) UpdateByOrderIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("OrderID").Eq(m.OrderID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.UpdateByOrderIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByOrderID(db)
	}

	return nil

}

func (m *OrderLogistics) UpdateByOrderIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByOrderIDWithMap(db, fieldValues)

}

func (m *OrderLogistics) FetchByOrderIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("OrderID").Eq(m.OrderID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.FetchByOrderIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *OrderLogistics) DeleteByOrderID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("OrderID").Eq(m.OrderID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.DeleteByOrderID"),
			))

	return err
}

func (m *OrderLogistics) SoftDeleteByOrderID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
					table.F("OrderID").Eq(m.OrderID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.SoftDeleteByOrderID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *OrderLogistics) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]OrderLogistics, error) {

	list := make([]OrderLogistics, 0)

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.List"),
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

func (m *OrderLogistics) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("OrderLogistics.Count"),
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

func (m *OrderLogistics) BatchFetchByCourierNumberList(db github_com_eden_framework_sqlx.DBExecutor, values []string) ([]OrderLogistics, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("CourierNumber").In(values)

	return m.List(db, condition)

}

func (m *OrderLogistics) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]OrderLogistics, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *OrderLogistics) BatchFetchByLogisticsIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]OrderLogistics, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("LogisticsID").In(values)

	return m.List(db, condition)

}

func (m *OrderLogistics) BatchFetchByOrderIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]OrderLogistics, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("OrderID").In(values)

	return m.List(db, condition)

}
