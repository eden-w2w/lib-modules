package databases

import (
	fmt "fmt"
	time "time"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
)

func (Administrators) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (Administrators) Indexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"I_expire": []string{
			"ExpiredAt",
		},
	}
}

func (Administrators) UniqueIndexUAdministratorsID() string {
	return "U_administrators_id"
}

func (Administrators) UniqueIndexUToken() string {
	return "U_token"
}

func (Administrators) UniqueIndexUUserName() string {
	return "U_user_name"
}

func (Administrators) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_administrators_id": []string{
			"AdministratorsID",
			"DeletedAt",
		},
		"U_token": []string{
			"Token",
			"DeletedAt",
		},
		"U_user_name": []string{
			"UserName",
			"DeletedAt",
		},
	}
}

func (Administrators) Comments() map[string]string {
	return map[string]string{
		"AdministratorsID": "业务ID",
		"ExpiredAt":        "访问令牌过期时间",
		"Password":         "密码",
		"Token":            "访问令牌",
		"UserName":         "用户名",
	}
}

var AdministratorsTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	AdministratorsTable = Config.DB.Register(&Administrators{})
}

type AdministratorsIterator struct {
}

func (AdministratorsIterator) New() interface{} {
	return &Administrators{}
}

func (AdministratorsIterator) Resolve(v interface{}) *Administrators {
	return v.(*Administrators)
}

func (Administrators) TableName() string {
	return "t_administrators"
}

func (Administrators) ColDescriptions() map[string][]string {
	return map[string][]string{
		"AdministratorsID": []string{
			"业务ID",
		},
		"ExpiredAt": []string{
			"访问令牌过期时间",
		},
		"Password": []string{
			"密码",
		},
		"Token": []string{
			"访问令牌",
		},
		"UserName": []string{
			"用户名",
		},
	}
}

func (Administrators) FieldKeyID() string {
	return "ID"
}

func (m *Administrators) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return AdministratorsTable.F(m.FieldKeyID())
}

func (Administrators) FieldKeyAdministratorsID() string {
	return "AdministratorsID"
}

func (m *Administrators) FieldAdministratorsID() *github_com_eden_framework_sqlx_builder.Column {
	return AdministratorsTable.F(m.FieldKeyAdministratorsID())
}

func (Administrators) FieldKeyUserName() string {
	return "UserName"
}

func (m *Administrators) FieldUserName() *github_com_eden_framework_sqlx_builder.Column {
	return AdministratorsTable.F(m.FieldKeyUserName())
}

func (Administrators) FieldKeyPassword() string {
	return "Password"
}

func (m *Administrators) FieldPassword() *github_com_eden_framework_sqlx_builder.Column {
	return AdministratorsTable.F(m.FieldKeyPassword())
}

func (Administrators) FieldKeyToken() string {
	return "Token"
}

func (m *Administrators) FieldToken() *github_com_eden_framework_sqlx_builder.Column {
	return AdministratorsTable.F(m.FieldKeyToken())
}

func (Administrators) FieldKeyExpiredAt() string {
	return "ExpiredAt"
}

func (m *Administrators) FieldExpiredAt() *github_com_eden_framework_sqlx_builder.Column {
	return AdministratorsTable.F(m.FieldKeyExpiredAt())
}

func (Administrators) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *Administrators) FieldCreatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return AdministratorsTable.F(m.FieldKeyCreatedAt())
}

func (Administrators) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *Administrators) FieldUpdatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return AdministratorsTable.F(m.FieldKeyUpdatedAt())
}

func (Administrators) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *Administrators) FieldDeletedAt() *github_com_eden_framework_sqlx_builder.Column {
	return AdministratorsTable.F(m.FieldKeyDeletedAt())
}

func (Administrators) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *Administrators) IndexFieldNames() []string {
	return []string{
		"AdministratorsID",
		"ExpiredAt",
		"ID",
		"Token",
		"UserName",
	}
}

func (m *Administrators) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
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

func (m *Administrators) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *Administrators) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

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

func (m *Administrators) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.DeleteByStruct"),
			),
	)

	return err
}

func (m *Administrators) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.FetchByID"),
			),
		m,
	)

	return err
}

func (m *Administrators) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("Administrators.UpdateByIDWithMap"),
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

func (m *Administrators) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *Administrators) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("Administrators.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Administrators) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.DeleteByID"),
			))

	return err
}

func (m *Administrators) SoftDeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("Administrators.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *Administrators) FetchByAdministratorsID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("AdministratorsID").Eq(m.AdministratorsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.FetchByAdministratorsID"),
			),
		m,
	)

	return err
}

func (m *Administrators) UpdateByAdministratorsIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("AdministratorsID").Eq(m.AdministratorsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.UpdateByAdministratorsIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByAdministratorsID(db)
	}

	return nil

}

func (m *Administrators) UpdateByAdministratorsIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByAdministratorsIDWithMap(db, fieldValues)

}

func (m *Administrators) FetchByAdministratorsIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("AdministratorsID").Eq(m.AdministratorsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.FetchByAdministratorsIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Administrators) DeleteByAdministratorsID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("AdministratorsID").Eq(m.AdministratorsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.DeleteByAdministratorsID"),
			))

	return err
}

func (m *Administrators) SoftDeleteByAdministratorsID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
					table.F("AdministratorsID").Eq(m.AdministratorsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.SoftDeleteByAdministratorsID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *Administrators) FetchByToken(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("Token").Eq(m.Token),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.FetchByToken"),
			),
		m,
	)

	return err
}

func (m *Administrators) UpdateByTokenWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("Token").Eq(m.Token),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.UpdateByTokenWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByToken(db)
	}

	return nil

}

func (m *Administrators) UpdateByTokenWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByTokenWithMap(db, fieldValues)

}

func (m *Administrators) FetchByTokenForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("Token").Eq(m.Token),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.FetchByTokenForUpdate"),
			),
		m,
	)

	return err
}

func (m *Administrators) DeleteByToken(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("Token").Eq(m.Token),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.DeleteByToken"),
			))

	return err
}

func (m *Administrators) SoftDeleteByToken(db github_com_eden_framework_sqlx.DBExecutor) error {

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
					table.F("Token").Eq(m.Token),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.SoftDeleteByToken"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *Administrators) FetchByUserName(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("UserName").Eq(m.UserName),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.FetchByUserName"),
			),
		m,
	)

	return err
}

func (m *Administrators) UpdateByUserNameWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("UserName").Eq(m.UserName),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.UpdateByUserNameWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByUserName(db)
	}

	return nil

}

func (m *Administrators) UpdateByUserNameWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByUserNameWithMap(db, fieldValues)

}

func (m *Administrators) FetchByUserNameForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("UserName").Eq(m.UserName),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.FetchByUserNameForUpdate"),
			),
		m,
	)

	return err
}

func (m *Administrators) DeleteByUserName(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("UserName").Eq(m.UserName),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.DeleteByUserName"),
			))

	return err
}

func (m *Administrators) SoftDeleteByUserName(db github_com_eden_framework_sqlx.DBExecutor) error {

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
					table.F("UserName").Eq(m.UserName),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("Administrators.SoftDeleteByUserName"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *Administrators) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]Administrators, error) {

	list := make([]Administrators, 0)

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("Administrators.List"),
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

func (m *Administrators) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("Administrators.Count"),
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

func (m *Administrators) BatchFetchByAdministratorsIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]Administrators, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("AdministratorsID").In(values)

	return m.List(db, condition)

}

func (m *Administrators) BatchFetchByExpiredAtList(db github_com_eden_framework_sqlx.DBExecutor, values []github_com_eden_framework_sqlx_datatypes.Timestamp) ([]Administrators, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ExpiredAt").In(values)

	return m.List(db, condition)

}

func (m *Administrators) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]Administrators, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *Administrators) BatchFetchByTokenList(db github_com_eden_framework_sqlx.DBExecutor, values []string) ([]Administrators, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Token").In(values)

	return m.List(db, condition)

}

func (m *Administrators) BatchFetchByUserNameList(db github_com_eden_framework_sqlx.DBExecutor, values []string) ([]Administrators, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("UserName").In(values)

	return m.List(db, condition)

}
