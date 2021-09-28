package databases

import (
	fmt "fmt"
	time "time"

	github_com_eden_framework_sqlx "github.com/eden-framework/sqlx"
	github_com_eden_framework_sqlx_builder "github.com/eden-framework/sqlx/builder"
	github_com_eden_framework_sqlx_datatypes "github.com/eden-framework/sqlx/datatypes"
)

func (OrderGoods) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (OrderGoods) UniqueIndexUOrderGoodsID() string {
	return "U_order_goods_id"
}

func (OrderGoods) UniqueIndexes() github_com_eden_framework_sqlx_builder.Indexes {
	return github_com_eden_framework_sqlx_builder.Indexes{
		"U_order_goods_id": []string{
			"OrderID",
			"GoodsID",
			"DeletedAt",
		},
	}
}

func (OrderGoods) Comments() map[string]string {
	return map[string]string{
		"Activities":     "活动",
		"Amount":         "-------------------------------------------",
		"Comment":        "描述",
		"Detail":         "详细介绍",
		"DispatchAddr":   "发货地",
		"GoodsID":        "商品ID",
		"Inventory":      "库存",
		"LogisticPolicy": "物流政策",
		"MainPicture":    "标题图片",
		"Name":           "---------------- 商品快照 ------------------",
		"OrderID":        "订单ID",
		"Pictures":       "所有展示图片",
		"Price":          "价格",
		"Sales":          "销量",
		"Specifications": "规格",
	}
}

var OrderGoodsTable *github_com_eden_framework_sqlx_builder.Table

func init() {
	OrderGoodsTable = Config.DB.Register(&OrderGoods{})
}

type OrderGoodsIterator struct {
}

func (OrderGoodsIterator) New() interface{} {
	return &OrderGoods{}
}

func (OrderGoodsIterator) Resolve(v interface{}) *OrderGoods {
	return v.(*OrderGoods)
}

func (OrderGoods) TableName() string {
	return "t_order_goods"
}

func (OrderGoods) ColDescriptions() map[string][]string {
	return map[string][]string{
		"Activities": []string{
			"活动",
		},
		"Amount": []string{
			"-------------------------------------------",
			"数量",
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
			"商品ID",
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
			"---------------- 商品快照 ------------------",
			"名称",
		},
		"OrderID": []string{
			"订单ID",
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

func (OrderGoods) FieldKeyID() string {
	return "ID"
}

func (m *OrderGoods) FieldID() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyID())
}

func (OrderGoods) FieldKeyOrderID() string {
	return "OrderID"
}

func (m *OrderGoods) FieldOrderID() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyOrderID())
}

func (OrderGoods) FieldKeyGoodsID() string {
	return "GoodsID"
}

func (m *OrderGoods) FieldGoodsID() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyGoodsID())
}

func (OrderGoods) FieldKeyName() string {
	return "Name"
}

func (m *OrderGoods) FieldName() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyName())
}

func (OrderGoods) FieldKeyComment() string {
	return "Comment"
}

func (m *OrderGoods) FieldComment() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyComment())
}

func (OrderGoods) FieldKeyDispatchAddr() string {
	return "DispatchAddr"
}

func (m *OrderGoods) FieldDispatchAddr() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyDispatchAddr())
}

func (OrderGoods) FieldKeySales() string {
	return "Sales"
}

func (m *OrderGoods) FieldSales() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeySales())
}

func (OrderGoods) FieldKeyMainPicture() string {
	return "MainPicture"
}

func (m *OrderGoods) FieldMainPicture() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyMainPicture())
}

func (OrderGoods) FieldKeyPictures() string {
	return "Pictures"
}

func (m *OrderGoods) FieldPictures() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyPictures())
}

func (OrderGoods) FieldKeySpecifications() string {
	return "Specifications"
}

func (m *OrderGoods) FieldSpecifications() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeySpecifications())
}

func (OrderGoods) FieldKeyActivities() string {
	return "Activities"
}

func (m *OrderGoods) FieldActivities() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyActivities())
}

func (OrderGoods) FieldKeyLogisticPolicy() string {
	return "LogisticPolicy"
}

func (m *OrderGoods) FieldLogisticPolicy() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyLogisticPolicy())
}

func (OrderGoods) FieldKeyPrice() string {
	return "Price"
}

func (m *OrderGoods) FieldPrice() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyPrice())
}

func (OrderGoods) FieldKeyInventory() string {
	return "Inventory"
}

func (m *OrderGoods) FieldInventory() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyInventory())
}

func (OrderGoods) FieldKeyDetail() string {
	return "Detail"
}

func (m *OrderGoods) FieldDetail() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyDetail())
}

func (OrderGoods) FieldKeyAmount() string {
	return "Amount"
}

func (m *OrderGoods) FieldAmount() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyAmount())
}

func (OrderGoods) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *OrderGoods) FieldCreatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyCreatedAt())
}

func (OrderGoods) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *OrderGoods) FieldUpdatedAt() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyUpdatedAt())
}

func (OrderGoods) FieldKeyDeletedAt() string {
	return "DeletedAt"
}

func (m *OrderGoods) FieldDeletedAt() *github_com_eden_framework_sqlx_builder.Column {
	return OrderGoodsTable.F(m.FieldKeyDeletedAt())
}

func (OrderGoods) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *OrderGoods) IndexFieldNames() []string {
	return []string{
		"GoodsID",
		"ID",
		"OrderID",
	}
}

func (m *OrderGoods) ConditionByStruct(db github_com_eden_framework_sqlx.DBExecutor) github_com_eden_framework_sqlx_builder.SqlCondition {
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

func (m *OrderGoods) Create(db github_com_eden_framework_sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_eden_framework_sqlx.InsertToDB(db, m, nil))
	return err

}

func (m *OrderGoods) CreateOnDuplicateWithUpdateFields(db github_com_eden_framework_sqlx.DBExecutor, updateFields []string) error {

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

func (m *OrderGoods) DeleteByStruct(db github_com_eden_framework_sqlx.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(m.ConditionByStruct(db)),
				github_com_eden_framework_sqlx_builder.Comment("OrderGoods.DeleteByStruct"),
			),
	)

	return err
}

func (m *OrderGoods) FetchByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("OrderGoods.FetchByID"),
			),
		m,
	)

	return err
}

func (m *OrderGoods) UpdateByIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

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
				github_com_eden_framework_sqlx_builder.Comment("OrderGoods.UpdateByIDWithMap"),
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

func (m *OrderGoods) UpdateByIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *OrderGoods) FetchByIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("OrderGoods.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *OrderGoods) DeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("ID").Eq(m.ID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("OrderGoods.DeleteByID"),
			))

	return err
}

func (m *OrderGoods) SoftDeleteByID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
				github_com_eden_framework_sqlx_builder.Comment("OrderGoods.SoftDeleteByID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *OrderGoods) FetchByOrderIDAndGoodsID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("OrderID").Eq(m.OrderID),
					table.F("GoodsID").Eq(m.GoodsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("OrderGoods.FetchByOrderIDAndGoodsID"),
			),
		m,
	)

	return err
}

func (m *OrderGoods) UpdateByOrderIDAndGoodsIDWithMap(db github_com_eden_framework_sqlx.DBExecutor, fieldValues github_com_eden_framework_sqlx_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_eden_framework_sqlx_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Update(db.T(m)).
			Where(
				github_com_eden_framework_sqlx_builder.And(
					table.F("OrderID").Eq(m.OrderID),
					table.F("GoodsID").Eq(m.GoodsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("OrderGoods.UpdateByOrderIDAndGoodsIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByOrderIDAndGoodsID(db)
	}

	return nil

}

func (m *OrderGoods) UpdateByOrderIDAndGoodsIDWithStruct(db github_com_eden_framework_sqlx.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_eden_framework_sqlx_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByOrderIDAndGoodsIDWithMap(db, fieldValues)

}

func (m *OrderGoods) FetchByOrderIDAndGoodsIDForUpdate(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_eden_framework_sqlx_builder.Select(nil).
			From(
				db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("OrderID").Eq(m.OrderID),
					table.F("GoodsID").Eq(m.GoodsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.ForUpdate(),
				github_com_eden_framework_sqlx_builder.Comment("OrderGoods.FetchByOrderIDAndGoodsIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *OrderGoods) DeleteByOrderIDAndGoodsID(db github_com_eden_framework_sqlx.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_eden_framework_sqlx_builder.Delete().
			From(db.T(m),
				github_com_eden_framework_sqlx_builder.Where(github_com_eden_framework_sqlx_builder.And(
					table.F("OrderID").Eq(m.OrderID),
					table.F("GoodsID").Eq(m.GoodsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				)),
				github_com_eden_framework_sqlx_builder.Comment("OrderGoods.DeleteByOrderIDAndGoodsID"),
			))

	return err
}

func (m *OrderGoods) SoftDeleteByOrderIDAndGoodsID(db github_com_eden_framework_sqlx.DBExecutor) error {

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
					table.F("GoodsID").Eq(m.GoodsID),
					table.F("DeletedAt").Eq(m.DeletedAt),
				),
				github_com_eden_framework_sqlx_builder.Comment("OrderGoods.SoftDeleteByOrderIDAndGoodsID"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	return err

}

func (m *OrderGoods) List(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) ([]OrderGoods, error) {

	list := make([]OrderGoods, 0)

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("OrderGoods.List"),
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

func (m *OrderGoods) Count(db github_com_eden_framework_sqlx.DBExecutor, condition github_com_eden_framework_sqlx_builder.SqlCondition, additions ...github_com_eden_framework_sqlx_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	condition = github_com_eden_framework_sqlx_builder.And(condition, table.F("DeletedAt").Eq(0))

	finalAdditions := []github_com_eden_framework_sqlx_builder.Addition{
		github_com_eden_framework_sqlx_builder.Where(condition),
		github_com_eden_framework_sqlx_builder.Comment("OrderGoods.Count"),
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

func (m *OrderGoods) BatchFetchByGoodsIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]OrderGoods, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("GoodsID").In(values)

	return m.List(db, condition)

}

func (m *OrderGoods) BatchFetchByIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]OrderGoods, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *OrderGoods) BatchFetchByOrderIDList(db github_com_eden_framework_sqlx.DBExecutor, values []uint64) ([]OrderGoods, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("OrderID").In(values)

	return m.List(db, condition)

}
