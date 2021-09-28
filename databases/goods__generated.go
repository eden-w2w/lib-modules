package databases

import (
	fmt "fmt"
	time "time"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
)

func (Goods) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (Goods) UniqueIndexUGoodsID() string {
	return "U_goods_id"
}

func (Goods) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_goods_id": []string{
			"GoodsID",
			"DeletedAt",
		},
	}
}

func (Goods) Comments() map[string]string {
	return map[string]string{
		"Activities":     "活动",
		"Comment":        "描述",
		"Detail":         "详细介绍",
		"DispatchAddr":   "发货地",
		"GoodsID":        "业务ID",
		"Inventory":      "库存",
		"LogisticPolicy": "物流政策",
		"MainPicture":    "标题图片",
		"Name":           "名称",
		"Pictures":       "所有展示图片",
		"Price":          "价格",
		"Sales":          "销量",
		"Specifications": "规格",
	}
}

var GoodsTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	GoodsTable = Config.DB.Register(&Goods{})
}

type GoodsIterator struct {
}

func (GoodsIterator) New() interface{} {
	return &Goods{}
}

func (GoodsIterator) Resolve(v interface{}) *Goods {
	return v.(*Goods)
}

func (Goods) TableName() string {
	return "t_goods"
}

func (Goods) ColDescriptions() map[string][]string {
	return map[string][]string{
		"Activities": []string{
			"活动",
		},
		"Comment": []string{
			"描述",
		},
		"Detail": []string{
			"详细介绍",
		},
		"DispatchAddr": []string{
			"发货地",
		},
		"GoodsID": []string{
			"业务ID",
		},
		"Inventory": []string{
			"库存",
		},
		"LogisticPolicy": []string{
			"物流政策",
		},
		"MainPicture": []string{
			"标题图片",
		},
		"Name": []string{
			"名称",
		},
		"Pictures": []string{
			"所有展示图片",
		},
		"Price": []string{
			"价格",
		},
		"Sales": []string{
			"销量",
		},
		"Specifications": []string{
			"规格",
		},
	}
}

func (Goods) FieldKeyID() string {
	return "ID"
}

func (m *Goods) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeyID())
}

func (Goods) FieldKeyGoodsID() string {
	return "GoodsID"
}

func (m *Goods) FieldGoodsID() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeyGoodsID())
}

func (Goods) FieldKeyName() string {
	return "Name"
}

func (m *Goods) FieldName() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeyName())
}

func (Goods) FieldKeyComment() string {
	return "Comment"
}

func (m *Goods) FieldComment() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeyComment())
}

func (Goods) FieldKeyDispatchAddr() string {
	return "DispatchAddr"
}

func (m *Goods) FieldDispatchAddr() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeyDispatchAddr())
}

func (Goods) FieldKeySales() string {
	return "Sales"
}

func (m *Goods) FieldSales() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeySales())
}

func (Goods) FieldKeyMainPicture() string {
	return "MainPicture"
}

func (m *Goods) FieldMainPicture() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeyMainPicture())
}

func (Goods) FieldKeyPictures() string {
	return "Pictures"
}

func (m *Goods) FieldPictures() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeyPictures())
}

func (Goods) FieldKeySpecifications() string {
	return "Specifications"
}

func (m *Goods) FieldSpecifications() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeySpecifications())
}

func (Goods) FieldKeyActivities() string {
	return "Activities"
}

func (m *Goods) FieldActivities() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeyActivities())
}

func (Goods) FieldKeyLogisticPolicy() string {
	return "LogisticPolicy"
}

func (m *Goods) FieldLogisticPolicy() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeyLogisticPolicy())
}

func (Goods) FieldKeyPrice() string {
	return "Price"
}

func (m *Goods) FieldPrice() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeyPrice())
}

func (Goods) FieldKeyInventory() string {
	return "Inventory"
}

func (m *Goods) FieldInventory() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeyInventory())
}

func (Goods) FieldKeyDetail() string {
	return "Detail"
}

func (m *Goods) FieldDetail() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeyDetail())
}

func (Goods) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *Goods) FieldCreatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeyCreatedAt())
}

func (Goods) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *Goods) FieldUpdatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeyUpdatedAt())
}

func (Goods) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *Goods) FieldDeletedAt() *github_com_eden_framework_sqlx_builder.Column {
	return GoodsTable.F(m.FieldKeyDeletedAt())
}

func (Goods) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *Goods) IndexFieldNames() []string {
	return []string{
		"GoodsID",
		"ID",
	}
}

func (m *Goods) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
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

func (m *Goods) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *Goods) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

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

func (m *Goods) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("Goods.DeleteByStruct"),
			),
	)

	return err
}

func (m *Goods) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Goods.FetchByID"),
			),
		m,
	)

	return err
}

func (m *Goods) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("Goods.UpdateByIDWithMap"),
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

func (m *Goods) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *Goods) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("Goods.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Goods) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Goods.DeleteByID"),
			))

	return err
}

func (m *Goods) SoftDeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("Goods.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *Goods) FetchByGoodsID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("GoodsID").Eq(m.GoodsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Goods.FetchByGoodsID"),
			),
		m,
	)

	return err
}

func (m *Goods) UpdateByGoodsIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("GoodsID").Eq(m.GoodsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("Goods.UpdateByGoodsIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByGoodsID(db)
	}

	return nil

}

func (m *Goods) UpdateByGoodsIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByGoodsIDWithMap(db, fieldValues)

}

func (m *Goods) FetchByGoodsIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("GoodsID").Eq(m.GoodsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("Goods.FetchByGoodsIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Goods) DeleteByGoodsID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("GoodsID").Eq(m.GoodsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("Goods.DeleteByGoodsID"),
			))

	return err
}

func (m *Goods) SoftDeleteByGoodsID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
					table.F("GoodsID").Eq(m.GoodsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("Goods.SoftDeleteByGoodsID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *Goods) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]Goods, error) {

	list := make([]Goods, 0)

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("Goods.List"),
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

func (m *Goods) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("Goods.Count"),
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

func (m *Goods) BatchFetchByGoodsIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]Goods, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("GoodsID").In(values)

	return m.List(db, condition)

}

func (m *Goods) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]Goods, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}
