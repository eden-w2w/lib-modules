package databases

import (
	fmt "fmt"
	time "time"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
)

func (ShippingAddress) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (ShippingAddress) Indexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"I_user": []string{
			"UserID",
			"Default",
		},
	}
}

func (ShippingAddress) UniqueIndexUShippingID() string {
	return "U_shipping_id"
}

func (ShippingAddress) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_shipping_id": []string{
			"ShippingID",
			"DeletedAt",
		},
	}
}

func (ShippingAddress) Comments() map[string]string {
	return map[string]string{
		"Address":    "详细地址",
		"Default":    "是否默认",
		"District":   "省市区街道",
		"Mobile":     "联系电话",
		"Recipients": "收件人",
		"ShippingID": "业务ID",
		"UserID":     "用户ID",
	}
}

var ShippingAddressTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	ShippingAddressTable = Config.DB.Register(&ShippingAddress{})
}

type ShippingAddressIterator struct {
}

func (ShippingAddressIterator) New() interface{} {
	return &ShippingAddress{}
}

func (ShippingAddressIterator) Resolve(v interface{}) *ShippingAddress {
	return v.(*ShippingAddress)
}

func (ShippingAddress) TableName() string {
	return "t_shipping_address"
}

func (ShippingAddress) ColDescriptions() map[string][]string {
	return map[string][]string{
		"Address": []string{
			"详细地址",
		},
		"Default": []string{
			"是否默认",
		},
		"District": []string{
			"省市区街道",
		},
		"Mobile": []string{
			"联系电话",
		},
		"Recipients": []string{
			"收件人",
		},
		"ShippingID": []string{
			"业务ID",
		},
		"UserID": []string{
			"用户ID",
		},
	}
}

func (ShippingAddress) FieldKeyID() string {
	return "ID"
}

func (m *ShippingAddress) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return ShippingAddressTable.F(m.FieldKeyID())
}

func (ShippingAddress) FieldKeyShippingID() string {
	return "ShippingID"
}

func (m *ShippingAddress) FieldShippingID() *github_com_eden_framework_sqlx_builder.Column {
	return ShippingAddressTable.F(m.FieldKeyShippingID())
}

func (ShippingAddress) FieldKeyUserID() string {
	return "UserID"
}

func (m *ShippingAddress) FieldUserID() *github_com_eden_framework_sqlx_builder.Column {
	return ShippingAddressTable.F(m.FieldKeyUserID())
}

func (ShippingAddress) FieldKeyRecipients() string {
	return "Recipients"
}

func (m *ShippingAddress) FieldRecipients() *github_com_eden_framework_sqlx_builder.Column {
	return ShippingAddressTable.F(m.FieldKeyRecipients())
}

func (ShippingAddress) FieldKeyDistrict() string {
	return "District"
}

func (m *ShippingAddress) FieldDistrict() *github_com_eden_framework_sqlx_builder.Column {
	return ShippingAddressTable.F(m.FieldKeyDistrict())
}

func (ShippingAddress) FieldKeyAddress() string {
	return "Address"
}

func (m *ShippingAddress) FieldAddress() *github_com_eden_framework_sqlx_builder.Column {
	return ShippingAddressTable.F(m.FieldKeyAddress())
}

func (ShippingAddress) FieldKeyMobile() string {
	return "Mobile"
}

func (m *ShippingAddress) FieldMobile() *github_com_eden_framework_sqlx_builder.Column {
	return ShippingAddressTable.F(m.FieldKeyMobile())
}

func (ShippingAddress) FieldKeyDefault() string {
	return "Default"
}

func (m *ShippingAddress) FieldDefault() *github_com_eden_framework_sqlx_builder.Column {
	return ShippingAddressTable.F(m.FieldKeyDefault())
}

func (ShippingAddress) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *ShippingAddress) FieldCreatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return ShippingAddressTable.F(m.FieldKeyCreatedAt())
}

func (ShippingAddress) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *ShippingAddress) FieldUpdatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return ShippingAddressTable.F(m.FieldKeyUpdatedAt())
}

func (ShippingAddress) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *ShippingAddress) FieldDeletedAt() *github_com_eden_framework_sqlx_builder.Column {
	return ShippingAddressTable.F(m.FieldKeyDeletedAt())
}

func (ShippingAddress) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *ShippingAddress) IndexFieldNames() []string {
	return []string{
		"Default",
		"ID",
		"ShippingID",
		"UserID",
	}
}

func (m *ShippingAddress) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
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

func (m *ShippingAddress) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *ShippingAddress) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

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

func (m *ShippingAddress) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("ShippingAddress.DeleteByStruct"),
			),
	)

	return err
}

func (m *ShippingAddress) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("ShippingAddress.FetchByID"),
			),
		m,
	)

	return err
}

func (m *ShippingAddress) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("ShippingAddress.UpdateByIDWithMap"),
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

func (m *ShippingAddress) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *ShippingAddress) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("ShippingAddress.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *ShippingAddress) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("ShippingAddress.DeleteByID"),
			))

	return err
}

func (m *ShippingAddress) SoftDeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("ShippingAddress.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *ShippingAddress) FetchByShippingID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ShippingID").Eq(m.ShippingID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("ShippingAddress.FetchByShippingID"),
			),
		m,
	)

	return err
}

func (m *ShippingAddress) UpdateByShippingIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("ShippingID").Eq(m.ShippingID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("ShippingAddress.UpdateByShippingIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByShippingID(db)
	}

	return nil

}

func (m *ShippingAddress) UpdateByShippingIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByShippingIDWithMap(db, fieldValues)

}

func (m *ShippingAddress) FetchByShippingIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ShippingID").Eq(m.ShippingID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("ShippingAddress.FetchByShippingIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *ShippingAddress) DeleteByShippingID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ShippingID").Eq(m.ShippingID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("ShippingAddress.DeleteByShippingID"),
			))

	return err
}

func (m *ShippingAddress) SoftDeleteByShippingID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
					table.F("ShippingID").Eq(m.ShippingID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("ShippingAddress.SoftDeleteByShippingID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *ShippingAddress) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]ShippingAddress, error) {

	list := make([]ShippingAddress, 0)

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("ShippingAddress.List"),
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

func (m *ShippingAddress) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("ShippingAddress.Count"),
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

func (m *ShippingAddress) BatchFetchByDefaultList(db github_com_eden_framework_sqlx.DBExecutor, values []github_com_eden_framework_sqlx_datatypes.Bool) ([]ShippingAddress, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Default").In(values)

	return m.List(db, condition)

}

func (m *ShippingAddress) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]ShippingAddress, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *ShippingAddress) BatchFetchByShippingIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]ShippingAddress, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ShippingID").In(values)

	return m.List(db, condition)

}

func (m *ShippingAddress) BatchFetchByUserIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]ShippingAddress, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("UserID").In(values)

	return m.List(db, condition)

}
