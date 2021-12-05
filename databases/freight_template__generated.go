package databases

import (
	fmt "fmt"
	time "time"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
	github_com_eden_w2_w_lib_modules_constants_enums "github.com/eden-w2w/lib-modules/constants/enums"
)

func (FreightTemplate) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (FreightTemplate) Indexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"I_search": []string{
			"Name",
			"IsFreeFreight",
			"Cal",
		},
	}
}

func (FreightTemplate) UniqueIndexUTemplateID() string {
	return "U_template_id"
}

func (FreightTemplate) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_template_id": []string{
			"TemplateID",
			"DeletedAt",
		},
	}
}

func (FreightTemplate) Comments() map[string]string {
	return map[string]string{
		"Cal":           "计费方式",
		"ContinuePrice": "续重/续件价格",
		"ContinueRange": "续重（克）/续件（个）范围",
		"DispatchAddr":  "发货地",
		"DispatchTime":  "发货时间",
		"FirstPrice":    "首重/首件价格",
		"FirstRange":    "-------------------------------------------------",
		"IsFreeFreight": "是否全场包邮",
		"Name":          "名称",
		"TemplateID":    "业务ID",
	}
}

var FreightTemplateTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	FreightTemplateTable = Config.DB.Register(&FreightTemplate{})
}

type FreightTemplateIterator struct {
}

func (FreightTemplateIterator) New() interface{} {
	return &FreightTemplate{}
}

func (FreightTemplateIterator) Resolve(v interface{}) *FreightTemplate {
	return v.(*FreightTemplate)
}

func (FreightTemplate) TableName() string {
	return "t_freight_template"
}

func (FreightTemplate) ColDescriptions() map[string][]string {
	return map[string][]string{
		"Cal": []string{
			"计费方式",
		},
		"ContinuePrice": []string{
			"续重/续件价格",
		},
		"ContinueRange": []string{
			"续重（克）/续件（个）范围",
		},
		"DispatchAddr": []string{
			"发货地",
		},
		"DispatchTime": []string{
			"发货时间",
		},
		"FirstPrice": []string{
			"首重/首件价格",
		},
		"FirstRange": []string{
			"-------------------------------------------------",
			"默认运费设置",
			"首重（克）/首件（个）范围",
		},
		"IsFreeFreight": []string{
			"是否全场包邮",
		},
		"Name": []string{
			"名称",
		},
		"TemplateID": []string{
			"业务ID",
		},
	}
}

func (FreightTemplate) FieldKeyID() string {
	return "ID"
}

func (m *FreightTemplate) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateTable.F(m.FieldKeyID())
}

func (FreightTemplate) FieldKeyTemplateID() string {
	return "TemplateID"
}

func (m *FreightTemplate) FieldTemplateID() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateTable.F(m.FieldKeyTemplateID())
}

func (FreightTemplate) FieldKeyName() string {
	return "Name"
}

func (m *FreightTemplate) FieldName() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateTable.F(m.FieldKeyName())
}

func (FreightTemplate) FieldKeyDispatchAddr() string {
	return "DispatchAddr"
}

func (m *FreightTemplate) FieldDispatchAddr() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateTable.F(m.FieldKeyDispatchAddr())
}

func (FreightTemplate) FieldKeyDispatchTime() string {
	return "DispatchTime"
}

func (m *FreightTemplate) FieldDispatchTime() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateTable.F(m.FieldKeyDispatchTime())
}

func (FreightTemplate) FieldKeyIsFreeFreight() string {
	return "IsFreeFreight"
}

func (m *FreightTemplate) FieldIsFreeFreight() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateTable.F(m.FieldKeyIsFreeFreight())
}

func (FreightTemplate) FieldKeyCal() string {
	return "Cal"
}

func (m *FreightTemplate) FieldCal() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateTable.F(m.FieldKeyCal())
}

func (FreightTemplate) FieldKeyFirstRange() string {
	return "FirstRange"
}

func (m *FreightTemplate) FieldFirstRange() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateTable.F(m.FieldKeyFirstRange())
}

func (FreightTemplate) FieldKeyFirstPrice() string {
	return "FirstPrice"
}

func (m *FreightTemplate) FieldFirstPrice() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateTable.F(m.FieldKeyFirstPrice())
}

func (FreightTemplate) FieldKeyContinueRange() string {
	return "ContinueRange"
}

func (m *FreightTemplate) FieldContinueRange() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateTable.F(m.FieldKeyContinueRange())
}

func (FreightTemplate) FieldKeyContinuePrice() string {
	return "ContinuePrice"
}

func (m *FreightTemplate) FieldContinuePrice() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateTable.F(m.FieldKeyContinuePrice())
}

func (FreightTemplate) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *FreightTemplate) FieldCreatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateTable.F(m.FieldKeyCreatedAt())
}

func (FreightTemplate) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *FreightTemplate) FieldUpdatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateTable.F(m.FieldKeyUpdatedAt())
}

func (FreightTemplate) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *FreightTemplate) FieldDeletedAt() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateTable.F(m.FieldKeyDeletedAt())
}

func (FreightTemplate) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *FreightTemplate) IndexFieldNames() []string {
	return []string{
		"Cal",
		"ID",
		"IsFreeFreight",
		"Name",
		"TemplateID",
	}
}

func (m *FreightTemplate) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
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

func (m *FreightTemplate) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *FreightTemplate) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

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

func (m *FreightTemplate) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplate.DeleteByStruct"),
			),
	)

	return err
}

func (m *FreightTemplate) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplate.FetchByID"),
			),
		m,
	)

	return err
}

func (m *FreightTemplate) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplate.UpdateByIDWithMap"),
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

func (m *FreightTemplate) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *FreightTemplate) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplate.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *FreightTemplate) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplate.DeleteByID"),
			))

	return err
}

func (m *FreightTemplate) SoftDeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplate.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *FreightTemplate) FetchByTemplateID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("TemplateID").Eq(m.TemplateID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplate.FetchByTemplateID"),
			),
		m,
	)

	return err
}

func (m *FreightTemplate) UpdateByTemplateIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("TemplateID").Eq(m.TemplateID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplate.UpdateByTemplateIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByTemplateID(db)
	}

	return nil

}

func (m *FreightTemplate) UpdateByTemplateIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByTemplateIDWithMap(db, fieldValues)

}

func (m *FreightTemplate) FetchByTemplateIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("TemplateID").Eq(m.TemplateID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplate.FetchByTemplateIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *FreightTemplate) DeleteByTemplateID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("TemplateID").Eq(m.TemplateID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplate.DeleteByTemplateID"),
			))

	return err
}

func (m *FreightTemplate) SoftDeleteByTemplateID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
					table.F("TemplateID").Eq(m.TemplateID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplate.SoftDeleteByTemplateID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *FreightTemplate) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]FreightTemplate, error) {

	list := make([]FreightTemplate, 0)

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("FreightTemplate.List"),
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

func (m *FreightTemplate) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("FreightTemplate.Count"),
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

func (m *FreightTemplate) BatchFetchByCalList(db github_com_eden_framework_sqlx.DBExecutor, values []github_com_eden_w2_w_lib_modules_constants_enums.FreightCal) ([]FreightTemplate, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Cal").In(values)

	return m.List(db, condition)

}

func (m *FreightTemplate) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]FreightTemplate, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *FreightTemplate) BatchFetchByIsFreeFreightList(db github_com_eden_framework_sqlx.DBExecutor, values []github_com_eden_framework_sqlx_datatypes.Bool) ([]FreightTemplate, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("IsFreeFreight").In(values)

	return m.List(db, condition)

}

func (m *FreightTemplate) BatchFetchByNameList(db github_com_eden_framework_sqlx.DBExecutor, values []string) ([]FreightTemplate, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Name").In(values)

	return m.List(db, condition)

}

func (m *FreightTemplate) BatchFetchByTemplateIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]FreightTemplate, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("TemplateID").In(values)

	return m.List(db, condition)

}
