package databases

import (
	fmt "fmt"
	time "time"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
)

func (FreightTemplateRules) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (FreightTemplateRules) Indexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"I_template_id": []string{
			"TemplateID",
			"IsFreeFreight",
		},
	}
}

func (FreightTemplateRules) UniqueIndexURuleID() string {
	return "U_rule_id"
}

func (FreightTemplateRules) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_rule_id": []string{
			"RuleID",
			"DeletedAt",
		},
	}
}

func (FreightTemplateRules) Comments() map[string]string {
	return map[string]string{
		"Area":          "包含区域",
		"ContinuePrice": "续重/续件价格",
		"ContinueRange": "续重（克）/续件（个）范围",
		"Description":   "展示话术",
		"FirstPrice":    "首重/首件价格",
		"FirstRange":    "-------------------------------------------------",
		"IsFreeFreight": "是否包邮",
		"RuleID":        "业务ID",
		"TemplateID":    "模板ID",
	}
}

var FreightTemplateRulesTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	FreightTemplateRulesTable = Config.DB.Register(&FreightTemplateRules{})
}

type FreightTemplateRulesIterator struct {
}

func (FreightTemplateRulesIterator) New() interface{} {
	return &FreightTemplateRules{}
}

func (FreightTemplateRulesIterator) Resolve(v interface{}) *FreightTemplateRules {
	return v.(*FreightTemplateRules)
}

func (FreightTemplateRules) TableName() string {
	return "t_freight_template_rules"
}

func (FreightTemplateRules) ColDescriptions() map[string][]string {
	return map[string][]string{
		"Area": []string{
			"包含区域",
		},
		"ContinuePrice": []string{
			"续重/续件价格",
		},
		"ContinueRange": []string{
			"续重（克）/续件（个）范围",
		},
		"Description": []string{
			"展示话术",
		},
		"FirstPrice": []string{
			"首重/首件价格",
		},
		"FirstRange": []string{
			"-------------------------------------------------",
			"运费设置",
			"首重（克）/首件（个）范围",
		},
		"IsFreeFreight": []string{
			"是否包邮",
		},
		"RuleID": []string{
			"业务ID",
		},
		"TemplateID": []string{
			"模板ID",
		},
	}
}

func (FreightTemplateRules) FieldKeyID() string {
	return "ID"
}

func (m *FreightTemplateRules) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateRulesTable.F(m.FieldKeyID())
}

func (FreightTemplateRules) FieldKeyRuleID() string {
	return "RuleID"
}

func (m *FreightTemplateRules) FieldRuleID() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateRulesTable.F(m.FieldKeyRuleID())
}

func (FreightTemplateRules) FieldKeyTemplateID() string {
	return "TemplateID"
}

func (m *FreightTemplateRules) FieldTemplateID() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateRulesTable.F(m.FieldKeyTemplateID())
}

func (FreightTemplateRules) FieldKeyArea() string {
	return "Area"
}

func (m *FreightTemplateRules) FieldArea() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateRulesTable.F(m.FieldKeyArea())
}

func (FreightTemplateRules) FieldKeyIsFreeFreight() string {
	return "IsFreeFreight"
}

func (m *FreightTemplateRules) FieldIsFreeFreight() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateRulesTable.F(m.FieldKeyIsFreeFreight())
}

func (FreightTemplateRules) FieldKeyDescription() string {
	return "Description"
}

func (m *FreightTemplateRules) FieldDescription() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateRulesTable.F(m.FieldKeyDescription())
}

func (FreightTemplateRules) FieldKeyFirstRange() string {
	return "FirstRange"
}

func (m *FreightTemplateRules) FieldFirstRange() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateRulesTable.F(m.FieldKeyFirstRange())
}

func (FreightTemplateRules) FieldKeyFirstPrice() string {
	return "FirstPrice"
}

func (m *FreightTemplateRules) FieldFirstPrice() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateRulesTable.F(m.FieldKeyFirstPrice())
}

func (FreightTemplateRules) FieldKeyContinueRange() string {
	return "ContinueRange"
}

func (m *FreightTemplateRules) FieldContinueRange() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateRulesTable.F(m.FieldKeyContinueRange())
}

func (FreightTemplateRules) FieldKeyContinuePrice() string {
	return "ContinuePrice"
}

func (m *FreightTemplateRules) FieldContinuePrice() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateRulesTable.F(m.FieldKeyContinuePrice())
}

func (FreightTemplateRules) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *FreightTemplateRules) FieldCreatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateRulesTable.F(m.FieldKeyCreatedAt())
}

func (FreightTemplateRules) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *FreightTemplateRules) FieldUpdatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateRulesTable.F(m.FieldKeyUpdatedAt())
}

func (FreightTemplateRules) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *FreightTemplateRules) FieldDeletedAt() *github_com_eden_framework_sqlx_builder.Column {
	return FreightTemplateRulesTable.F(m.FieldKeyDeletedAt())
}

func (FreightTemplateRules) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *FreightTemplateRules) IndexFieldNames() []string {
	return []string{
		"ID",
		"IsFreeFreight",
		"RuleID",
		"TemplateID",
	}
}

func (m *FreightTemplateRules) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
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

func (m *FreightTemplateRules) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *FreightTemplateRules) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

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

func (m *FreightTemplateRules) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplateRules.DeleteByStruct"),
			),
	)

	return err
}

func (m *FreightTemplateRules) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplateRules.FetchByID"),
			),
		m,
	)

	return err
}

func (m *FreightTemplateRules) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplateRules.UpdateByIDWithMap"),
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

func (m *FreightTemplateRules) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *FreightTemplateRules) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplateRules.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *FreightTemplateRules) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplateRules.DeleteByID"),
			))

	return err
}

func (m *FreightTemplateRules) SoftDeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplateRules.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *FreightTemplateRules) FetchByRuleID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("RuleID").Eq(m.RuleID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplateRules.FetchByRuleID"),
			),
		m,
	)

	return err
}

func (m *FreightTemplateRules) UpdateByRuleIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("RuleID").Eq(m.RuleID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplateRules.UpdateByRuleIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByRuleID(db)
	}

	return nil

}

func (m *FreightTemplateRules) UpdateByRuleIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByRuleIDWithMap(db, fieldValues)

}

func (m *FreightTemplateRules) FetchByRuleIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("RuleID").Eq(m.RuleID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplateRules.FetchByRuleIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *FreightTemplateRules) DeleteByRuleID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("RuleID").Eq(m.RuleID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplateRules.DeleteByRuleID"),
			))

	return err
}

func (m *FreightTemplateRules) SoftDeleteByRuleID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
					table.F("RuleID").Eq(m.RuleID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("FreightTemplateRules.SoftDeleteByRuleID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *FreightTemplateRules) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]FreightTemplateRules, error) {

	list := make([]FreightTemplateRules, 0)

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("FreightTemplateRules.List"),
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

func (m *FreightTemplateRules) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("FreightTemplateRules.Count"),
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

func (m *FreightTemplateRules) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]FreightTemplateRules, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *FreightTemplateRules) BatchFetchByIsFreeFreightList(db github_com_eden_framework_sqlx.DBExecutor, values []github_com_eden_framework_sqlx_datatypes.Bool) ([]FreightTemplateRules, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("IsFreeFreight").In(values)

	return m.List(db, condition)

}

func (m *FreightTemplateRules) BatchFetchByRuleIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]FreightTemplateRules, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("RuleID").In(values)

	return m.List(db, condition)

}

func (m *FreightTemplateRules) BatchFetchByTemplateIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]FreightTemplateRules, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("TemplateID").In(values)

	return m.List(db, condition)

}
