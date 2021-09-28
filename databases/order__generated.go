package databases

import (
	fmt "fmt"
	time "time"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
	github_com_eden_w2_w_lib_modules_constants_enums "github.com/eden-w2w/lib-modules/constants/enums"
)

func (Order) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (Order) Indexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"I_expire": []string{
			"ExpiredAt",
		},
		"I_index": []string{
			"UserID",
			"Status",
		},
	}
}

func (Order) UniqueIndexUOrderID() string {
	return "U_order_id"
}

func (Order) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_order_id": []string{
			"OrderID",
			"DeletedAt",
		},
	}
}

func (Order) Comments() map[string]string {
	return map[string]string{
		"ActualAmount":   "实际金额",
		"DiscountAmount": "优惠金额",
		"ExpiredAt":      "过期时间",
		"OrderID":        "业务ID",
		"PaymentMethod":  "支付方式",
		"RefererID":      "推荐人ID",
		"Remark":         "备注",
		"Status":         "订单状态",
		"TotalPrice":     "订单总额",
		"UserID":         "用户ID",
	}
}

var OrderTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	OrderTable = Config.DB.Register(&Order{})
}

type OrderIterator struct {
}

func (OrderIterator) New() interface{} {
	return &Order{}
}

func (OrderIterator) Resolve(v interface{}) *Order {
	return v.(*Order)
}

func (Order) TableName() string {
	return "t_order"
}

func (Order) ColDescriptions() map[string][]string {
	return map[string][]string{
		"ActualAmount": []string{
			"实际金额",
		},
		"DiscountAmount": []string{
			"优惠金额",
		},
		"ExpiredAt": []string{
			"过期时间",
		},
		"OrderID": []string{
			"业务ID",
		},
		"PaymentMethod": []string{
			"支付方式",
		},
		"RefererID": []string{
			"推荐人ID",
		},
		"Remark": []string{
			"备注",
		},
		"Status": []string{
			"订单状态",
		},
		"TotalPrice": []string{
			"订单总额",
		},
		"UserID": []string{
			"用户ID",
		},
	}
}

func (Order) FieldKeyID() string {
	return "ID"
}

func (m *Order) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return OrderTable.F(m.FieldKeyID())
}

func (Order) FieldKeyOrderID() string {
	return "OrderID"
}

func (m *Order) FieldOrderID() *github_com_eden_framework_sqlx_builder.Column {
	return OrderTable.F(m.FieldKeyOrderID())
}

func (Order) FieldKeyUserID() string {
	return "UserID"
}

func (m *Order) FieldUserID() *github_com_eden_framework_sqlx_builder.Column {
	return OrderTable.F(m.FieldKeyUserID())
}

func (Order) FieldKeyRefererID() string {
	return "RefererID"
}

func (m *Order) FieldRefererID() *github_com_eden_framework_sqlx_builder.Column {
	return OrderTable.F(m.FieldKeyRefererID())
}

func (Order) FieldKeyTotalPrice() string {
	return "TotalPrice"
}

func (m *Order) FieldTotalPrice() *github_com_eden_framework_sqlx_builder.Column {
	return OrderTable.F(m.FieldKeyTotalPrice())
}

func (Order) FieldKeyDiscountAmount() string {
	return "DiscountAmount"
}

func (m *Order) FieldDiscountAmount() *github_com_eden_framework_sqlx_builder.Column {
	return OrderTable.F(m.FieldKeyDiscountAmount())
}

func (Order) FieldKeyActualAmount() string {
	return "ActualAmount"
}

func (m *Order) FieldActualAmount() *github_com_eden_framework_sqlx_builder.Column {
	return OrderTable.F(m.FieldKeyActualAmount())
}

func (Order) FieldKeyPaymentMethod() string {
	return "PaymentMethod"
}

func (m *Order) FieldPaymentMethod() *github_com_eden_framework_sqlx_builder.Column {
	return OrderTable.F(m.FieldKeyPaymentMethod())
}

func (Order) FieldKeyRemark() string {
	return "Remark"
}

func (m *Order) FieldRemark() *github_com_eden_framework_sqlx_builder.Column {
	return OrderTable.F(m.FieldKeyRemark())
}

func (Order) FieldKeyStatus() string {
	return "Status"
}

func (m *Order) FieldStatus() *github_com_eden_framework_sqlx_builder.Column {
	return OrderTable.F(m.FieldKeyStatus())
}

func (Order) FieldKeyExpiredAt() string {
	return "ExpiredAt"
}

func (m *Order) FieldExpiredAt() *github_com_eden_framework_sqlx_builder.Column {
	return OrderTable.F(m.FieldKeyExpiredAt())
}

func (Order) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *Order) FieldCreatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return OrderTable.F(m.FieldKeyCreatedAt())
}

func (Order) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *Order) FieldUpdatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return OrderTable.F(m.FieldKeyUpdatedAt())
}

func (Order) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *Order) FieldDeletedAt() *github_com_eden_framework_sqlx_builder.Column {
	return OrderTable.F(m.FieldKeyDeletedAt())
}

func (Order) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *Order) IndexFieldNames() []string {
	return []string{
		"ExpiredAt",
		"ID",
		"OrderID",
		"Status",
		"UserID",
	}
}

func (m *Order) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
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

func (m *Order) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *Order) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

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

func (m *Order) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("Order.DeleteByStruct"),
			),
	)

	return err
}

func (m *Order) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Order.FetchByID"),
			),
		m,
	)

	return err
}

func (m *Order) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("Order.UpdateByIDWithMap"),
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

func (m *Order) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *Order) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("Order.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Order) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Order.DeleteByID"),
			))

	return err
}

func (m *Order) SoftDeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("Order.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *Order) FetchByOrderID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("OrderID").Eq(m.OrderID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Order.FetchByOrderID"),
			),
		m,
	)

	return err
}

func (m *Order) UpdateByOrderIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("Order.UpdateByOrderIDWithMap"),
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

func (m *Order) UpdateByOrderIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByOrderIDWithMap(db, fieldValues)

}

func (m *Order) FetchByOrderIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("Order.FetchByOrderIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Order) DeleteByOrderID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("OrderID").Eq(m.OrderID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Order.DeleteByOrderID"),
			))

	return err
}

func (m *Order) SoftDeleteByOrderID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("Order.SoftDeleteByOrderID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *Order) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]Order, error) {

	list := make([]Order, 0)

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("Order.List"),
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

func (m *Order) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("Order.Count"),
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

func (m *Order) BatchFetchByExpiredAtList(db github_com_eden_framework_sqlx.DBExecutor, values []github_com_eden_framework_sqlx_datatypes.Timestamp) ([]Order, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ExpiredAt").In(values)

	return m.List(db, condition)

}

func (m *Order) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]Order, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *Order) BatchFetchByOrderIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]Order, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("OrderID").In(values)

	return m.List(db, condition)

}

func (m *Order) BatchFetchByStatusList(db github_com_eden_framework_sqlx.DBExecutor, values []github_com_eden_w2_w_lib_modules_constants_enums.OrderStatus) ([]Order, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Status").In(values)

	return m.List(db, condition)

}

func (m *Order) BatchFetchByUserIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]Order, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("UserID").In(values)

	return m.List(db, condition)

}
