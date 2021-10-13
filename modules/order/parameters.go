package order

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-framework/sqlx/datatypes"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/constants/types"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules"
)

type InventoryLock func(db sqlx.DBExecutor, goodsID uint64, amount uint32) error
type InventoryUnlock func(db sqlx.DBExecutor, goodsID uint64, amount uint32) error

type EventHandler interface {
	OnOrderCreateEvent(db sqlx.DBExecutor, order *databases.Order) error
	OnOrderPaidEvent(db sqlx.DBExecutor, order *databases.Order, payment *databases.PaymentFlow) error
	OnOrderCompleteEvent(db sqlx.DBExecutor, order *databases.Order) error
}

type CreateOrderParams struct {
	// 用户ID
	UserID uint64 `in:"body" default:"" json:"userID,string"`
	// 订单总额
	TotalPrice uint64 `in:"body" json:"totalPrice"`
	// 优惠金额
	DiscountAmount uint64 `in:"body" default:"" json:"discountAmount"`
	// 实际金额
	ActualAmount uint64 `in:"body" json:"actualAmount"`
	// 支付方式
	PaymentMethod enums.PaymentMethod `in:"body" json:"paymentMethod"`
	// 备注
	Remark string `in:"body" default:"" json:"remark"`
	// 收件人
	Recipients string `in:"body" json:"recipients"`
	// 收货地址
	ShippingAddr string `in:"body" json:"shippingAddr"`
	// 联系电话
	Mobile string `in:"body" json:"mobile"`
	// 物料信息
	Goods []CreateOrderGoodsParams `in:"body" json:"goods"`
}

type CreateOrderGoodsParams struct {
	// 商品ID
	GoodsID uint64 `in:"body" json:"goodsID,string"`
	// 数量
	Amount uint32 `in:"body" json:"amount"`
}

type CreateOrderGoodsModelParams struct {
	databases.Goods
	Amount uint32
}

type GetOrdersParams struct {
	// 用户ID
	UserID uint64 `in:"query" default:"" name:"userID,string" json:"userID,string"`
	// 支付方式
	PaymentMethod enums.PaymentMethod `in:"query" default:"" name:"paymentMethod" json:"paymentMethod"`
	// 订单状态
	Status enums.OrderStatus `in:"query" default:"" name:"status" json:"status"`

	modules.Pagination
}

func (p GetOrdersParams) Conditions() builder.SqlCondition {
	var condition builder.SqlCondition
	model := &databases.Order{}

	if p.UserID != 0 {
		condition = builder.And(condition, model.FieldUserID().Eq(p.UserID))
	}
	if p.PaymentMethod != enums.PAYMENT_METHOD_UNKNOWN {
		condition = builder.And(condition, model.FieldPaymentMethod().Eq(p.PaymentMethod))
	}
	if p.Status != enums.ORDER_STATUS_UNKNOWN {
		condition = builder.And(condition, model.FieldStatus().Eq(p.Status))
	}

	return condition
}

func (p GetOrdersParams) Additions() []builder.Addition {
	var additions = make([]builder.Addition, 0)

	if p.Size != 0 {
		limit := builder.Limit(int64(p.Size))
		if p.Offset != 0 {
			limit.Offset(int64(p.Offset))
		}
		additions = append(additions, limit)
	}

	additions = append(additions, builder.OrderBy(builder.DescOrder((&databases.Order{}).FieldCreatedAt())))

	return additions
}

type GoodsListResponse struct {
	// 商品ID
	GoodsID uint64 `json:"goodsID,string"`
	// 名称
	Name string `json:"name"`
	// 标题图片
	MainPicture string `json:"mainPicture"`
	// 规格
	Specifications types.JsonArrayString `json:"specifications"`
	// 价格
	Price uint64 `json:"price"`
	// 数量
	Amount uint32 `json:"amount"`
}

type GetMyOrdersResponse struct {
	// 业务ID
	OrderID uint64 `json:"orderID,string"`
	// 用户ID
	UserID uint64 `json:"userID,string"`
	// 订单总额
	TotalPrice uint64 `json:"totalPrice"`
	// 支付方式
	PaymentMethod enums.PaymentMethod `json:"paymentMethod"`
	// 订单状态
	Status enums.OrderStatus `json:"status"`
	// 创建时间
	CreatedAt datatypes.MySQLTimestamp `json:"createdAt"`
	// 物料
	Goods []GoodsListResponse `json:"goods"`
}

type GetOrderByIDResponse struct {
	// 业务ID
	OrderID uint64 `json:"orderID,string"`
	// 用户ID
	UserID uint64 `json:"userID,string"`
	// 昵称
	NickName string `json:"nickName"`
	// 微信OpenID
	OpenID string `json:"openID"`
	// 订单总额
	TotalPrice uint64 `json:"totalPrice"`
	// 优惠金额
	DiscountAmount uint64 `json:"discountAmount"`
	// 实际金额
	ActualAmount uint64 `json:"actualAmount"`
	// 支付方式
	PaymentMethod enums.PaymentMethod `json:"paymentMethod"`
	// 备注
	Remark string `json:"remark"`
	// 收件人
	Recipients string `json:"recipients"`
	// 收货地址
	ShippingAddr string `json:"shippingAddr"`
	// 联系电话
	Mobile string `json:"mobile"`
	// 快递公司
	CourierCompany string `json:"courierCompany"`
	// 快递单号
	CourierNumber string `json:"courierNumber"`
	// 订单状态
	Status enums.OrderStatus `json:"status"`
	// 过期时间
	ExpiredAt datatypes.MySQLTimestamp `json:"expiredAt"`
	// 创建时间
	CreatedAt datatypes.MySQLTimestamp `json:"createdAt" `
	// 更新时间
	UpdatedAt datatypes.MySQLTimestamp `json:"updatedAt"`
	// 物料
	Goods []GoodsListResponse `json:"goods"`
}

type UpdateOrderParams struct {
	// 订单状态
	Status enums.OrderStatus `json:"status" default:""`
	// 优惠金额
	DiscountAmount uint64 `json:"discountAmount" default:""`
	// 备注
	Remark string `json:"remark" default:""`
	// 收件人
	Recipients string `json:"recipients" default:""`
	// 收货地址
	ShippingAddr string `json:"shippingAddr" default:""`
	// 联系电话
	Mobile string `json:"mobile" default:""`
	// 快递公司
	CourierCompany string `json:"courierCompany" default:""`
	// 快递单号
	CourierNumber string `json:"courierNumber" default:""`
}
