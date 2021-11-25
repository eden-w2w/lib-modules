package databases

import (
	fmt "fmt"
	time "time"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
)

func (MarketingDiscount) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (MarketingDiscount) UniqueIndexUDiscountID() string {
	return "U_discount_id"
}

func (MarketingDiscount) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_discount_id": []string{
			"DiscountID",
			"DeletedAt",
		},
	}
}

func (MarketingDiscount) Comments() map[string]string {
	return map[string]string{
		"Cal":                "计算方式",
		"DiscountAmount":     "单价立减金额",
		"DiscountID":         "业务ID",
		"DiscountLimit":      "优惠上限",
		"DiscountRate":       "单价折扣比例",
		"Limit":              "总数限制",
		"MultiStepRate":      "阶梯式折扣比例",
		"MultiStepReduction": "阶梯式立减金额",
		"Name":               "营销名称",
		"Status":             "营销状态",
		"Times":              "已优惠次数",
		"Type":               "营销类型",
		"UserLimit":          "单用户优惠次数上限",
		"ValidityEnd":        "有效期结束",
		"ValidityStart":      "有效期开始",
	}
}

var MarketingDiscountTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	MarketingDiscountTable = Config.DB.Register(&MarketingDiscount{})
}

type MarketingDiscountIterator struct {
}

func (MarketingDiscountIterator) New() interface{} {
	return &MarketingDiscount{}
}

func (MarketingDiscountIterator) Resolve(v interface{}) *MarketingDiscount {
	return v.(*MarketingDiscount)
}

func (MarketingDiscount) TableName() string {
	return "t_marketing_discount"
}

func (MarketingDiscount) ColDescriptions() map[string][]string {
	return map[string][]string{
		"Cal": []string{
			"计算方式",
		},
		"DiscountAmount": []string{
			"单价立减金额",
		},
		"DiscountID": []string{
			"业务ID",
		},
		"DiscountLimit": []string{
			"优惠上限",
		},
		"DiscountRate": []string{
			"单价折扣比例",
		},
		"Limit": []string{
			"总数限制",
		},
		"MultiStepRate": []string{
			"阶梯式折扣比例",
		},
		"MultiStepReduction": []string{
			"阶梯式立减金额",
		},
		"Name": []string{
			"营销名称",
		},
		"Status": []string{
			"营销状态",
		},
		"Times": []string{
			"已优惠次数",
		},
		"Type": []string{
			"营销类型",
		},
		"UserLimit": []string{
			"单用户优惠次数上限",
		},
		"ValidityEnd": []string{
			"有效期结束",
		},
		"ValidityStart": []string{
			"有效期开始",
		},
	}
}

func (MarketingDiscount) FieldKeyID() string {
	return "ID"
}

func (m *MarketingDiscount) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyID())
}

func (MarketingDiscount) FieldKeyDiscountID() string {
	return "DiscountID"
}

func (m *MarketingDiscount) FieldDiscountID() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyDiscountID())
}

func (MarketingDiscount) FieldKeyName() string {
	return "Name"
}

func (m *MarketingDiscount) FieldName() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyName())
}

func (MarketingDiscount) FieldKeyType() string {
	return "Type"
}

func (m *MarketingDiscount) FieldType() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyType())
}

func (MarketingDiscount) FieldKeyStatus() string {
	return "Status"
}

func (m *MarketingDiscount) FieldStatus() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyStatus())
}

func (MarketingDiscount) FieldKeyCal() string {
	return "Cal"
}

func (m *MarketingDiscount) FieldCal() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyCal())
}

func (MarketingDiscount) FieldKeyValidityStart() string {
	return "ValidityStart"
}

func (m *MarketingDiscount) FieldValidityStart() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyValidityStart())
}

func (MarketingDiscount) FieldKeyValidityEnd() string {
	return "ValidityEnd"
}

func (m *MarketingDiscount) FieldValidityEnd() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyValidityEnd())
}

func (MarketingDiscount) FieldKeyUserLimit() string {
	return "UserLimit"
}

func (m *MarketingDiscount) FieldUserLimit() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyUserLimit())
}

func (MarketingDiscount) FieldKeyLimit() string {
	return "Limit"
}

func (m *MarketingDiscount) FieldLimit() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyLimit())
}

func (MarketingDiscount) FieldKeyTimes() string {
	return "Times"
}

func (m *MarketingDiscount) FieldTimes() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyTimes())
}

func (MarketingDiscount) FieldKeyDiscountLimit() string {
	return "DiscountLimit"
}

func (m *MarketingDiscount) FieldDiscountLimit() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyDiscountLimit())
}

func (MarketingDiscount) FieldKeyDiscountRate() string {
	return "DiscountRate"
}

func (m *MarketingDiscount) FieldDiscountRate() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyDiscountRate())
}

func (MarketingDiscount) FieldKeyMultiStepRate() string {
	return "MultiStepRate"
}

func (m *MarketingDiscount) FieldMultiStepRate() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyMultiStepRate())
}

func (MarketingDiscount) FieldKeyDiscountAmount() string {
	return "DiscountAmount"
}

func (m *MarketingDiscount) FieldDiscountAmount() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyDiscountAmount())
}

func (MarketingDiscount) FieldKeyMultiStepReduction() string {
	return "MultiStepReduction"
}

func (m *MarketingDiscount) FieldMultiStepReduction() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyMultiStepReduction())
}

func (MarketingDiscount) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *MarketingDiscount) FieldCreatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyCreatedAt())
}

func (MarketingDiscount) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *MarketingDiscount) FieldUpdatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyUpdatedAt())
}

func (MarketingDiscount) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *MarketingDiscount) FieldDeletedAt() *github_com_eden_framework_sqlx_builder.Column {
	return MarketingDiscountTable.F(m.FieldKeyDeletedAt())
}

func (MarketingDiscount) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *MarketingDiscount) IndexFieldNames() []string {
	return []string{
		"DiscountID",
		"ID",
	}
}

func (m *MarketingDiscount) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
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

func (m *MarketingDiscount) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *MarketingDiscount) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

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

func (m *MarketingDiscount) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("MarketingDiscount.DeleteByStruct"),
			),
	)

	return err
}

func (m *MarketingDiscount) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("MarketingDiscount.FetchByID"),
			),
		m,
	)

	return err
}

func (m *MarketingDiscount) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("MarketingDiscount.UpdateByIDWithMap"),
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

func (m *MarketingDiscount) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *MarketingDiscount) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("MarketingDiscount.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *MarketingDiscount) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("MarketingDiscount.DeleteByID"),
			))

	return err
}

func (m *MarketingDiscount) SoftDeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("MarketingDiscount.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *MarketingDiscount) FetchByDiscountID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("DiscountID").Eq(m.DiscountID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("MarketingDiscount.FetchByDiscountID"),
			),
		m,
	)

	return err
}

func (m *MarketingDiscount) UpdateByDiscountIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("DiscountID").Eq(m.DiscountID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("MarketingDiscount.UpdateByDiscountIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByDiscountID(db)
	}

	return nil

}

func (m *MarketingDiscount) UpdateByDiscountIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByDiscountIDWithMap(db, fieldValues)

}

func (m *MarketingDiscount) FetchByDiscountIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("DiscountID").Eq(m.DiscountID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("MarketingDiscount.FetchByDiscountIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *MarketingDiscount) DeleteByDiscountID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("DiscountID").Eq(m.DiscountID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("MarketingDiscount.DeleteByDiscountID"),
			))

	return err
}

func (m *MarketingDiscount) SoftDeleteByDiscountID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
					table.F("DiscountID").Eq(m.DiscountID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("MarketingDiscount.SoftDeleteByDiscountID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *MarketingDiscount) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]MarketingDiscount, error) {

	list := make([]MarketingDiscount, 0)

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("MarketingDiscount.List"),
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

func (m *MarketingDiscount) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("MarketingDiscount.Count"),
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

func (m *MarketingDiscount) BatchFetchByDiscountIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]MarketingDiscount, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("DiscountID").In(values)

	return m.List(db, condition)

}

func (m *MarketingDiscount) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]MarketingDiscount, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}
