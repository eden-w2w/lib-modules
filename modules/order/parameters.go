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
	OnOrderCreateEvent(db sqlx.DBExecutor, order *databases.Order, goodsList []databases.OrderGoods) error
	OnOrderPaidEvent(db sqlx.DBExecutor, order *databases.Order, payment *databases.PaymentFlow) error
	OnOrderConfirmEvent(db sqlx.DBExecutor, order *databases.Order) error
	OnOrderDispatchEvent(
		db sqlx.DBExecutor,
		order *databases.Order,
		logistics *databases.OrderLogistics,
		goodsList []databases.OrderGoods,
	) error
	OnOrderCompleteEvent(
		db sqlx.DBExecutor,
		order *databases.Order,
		logistics *databases.OrderLogistics,
		goods []databases.OrderGoods,
	) error
	OnOrderCloseEvent(db sqlx.DBExecutor, order *databases.Order, goodsList []databases.OrderGoods) error
}

type PreCreateOrderParams struct {
	// 用户ID
	UserID uint64 `in:"body" default:"" json:"userID,string"`
	// 优惠信息
	Discounts types.Uint64List `json:"discounts" default:""`
	// 物料信息
	Goods []CreateOrderGoodsParams `in:"body" json:"goods"`
	// 收件地址ID
	ShippingID uint64 `in:"body" json:"shippingID,string" default:""`
}

type CreateOrderParams struct {
	// 用户ID
	UserID uint64 `in:"body" default:"" json:"userID,string"`
	// 订单总额
	TotalPrice uint64 `in:"body" json:"totalPrice"`
	// 运费
	FreightAmount uint64 `json:"freightAmount" default:""`
	// 支付方式
	PaymentMethod enums.PaymentMethod `in:"body" json:"paymentMethod"`
	// 备注
	Remark string `in:"body" default:"" json:"remark"`
	// 收件地址ID
	ShippingID uint64 `in:"body" json:"shippingID,string"`
	// 收件人
	Recipients string `in:"body" json:"recipients"`
	// 收货地址
	ShippingAddr string `in:"body" json:"shippingAddr"`
	// 联系电话
	Mobile string `in:"body" json:"mobile"`
	// 优惠信息
	Discounts types.Uint64List `json:"discounts" default:""`
	// 物料信息
	Goods []CreateOrderGoodsParams `in:"body" json:"goods"`
}

type CreateOrderGoodsParams struct {
	// 商品ID
	GoodsID uint64 `in:"body" json:"goodsID,string"`
	// 数量
	Amount uint32 `in:"body" json:"amount"`
	// 是否预订
	IsBooking *bool `json:"isBooking"`
}

type PreCreateOrderGoodsParams struct {
	// 商品ID
	GoodsID uint64 `in:"body" json:"goodsID,string"`
	// 数量
	Amount uint32 `in:"body" json:"amount"`
	// 是否预订
	IsBooking *bool `json:"isBooking"`
	// 商品金额
	Price uint64 `json:"price"`
	// 折扣金额
	DiscountPrice uint64 `json:"discountPrice"`
}

type CreateOrderGoodsModelParams struct {
	databases.Goods
	Amount        uint32
	IsBooking     datatypes.Bool
	BookingFlowID uint64
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
			limit = limit.Offset(int64(p.Offset))
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
	// 简介
	Comment string `json:"comment"`
	// 标题图片
	MainPicture string `json:"mainPicture"`
	// 规格
	Specifications types.JsonArrayString `json:"specifications"`
	// 价格
	Price uint64 `json:"price"`
	// 数量
	Amount uint32 `json:"amount"`
	// 是否预订
	IsBooking datatypes.Bool `json:"isBooking"`
}

type GetMyOrdersResponse struct {
	// 业务ID
	OrderID uint64 `json:"orderID,string"`
	// 用户ID
	UserID uint64 `json:"userID,string"`
	// 订单总额
	TotalPrice uint64 `json:"totalPrice"`
	// 运费
	FreightAmount uint64 `json:"freightAmount"`
	// 优惠金额
	DiscountAmount uint64 `json:"discountAmount"`
	// 实际金额
	ActualAmount uint64 `json:"actualAmount"`
	// 支付方式
	PaymentMethod enums.PaymentMethod `json:"paymentMethod"`
	// 订单状态
	Status enums.OrderStatus `json:"status"`
	// 创建时间
	CreatedAt datatypes.MySQLTimestamp `json:"createdAt"`
	// 物料
	Goods []GoodsListResponse `json:"goods"`
	// 物流信息
	Logistics *databases.OrderLogistics `json:"logistics"`
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
	// 运费
	FreightAmount uint64 `json:"freightAmount"`
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
	// 运费
	FreightAmount *uint64 `json:"freightAmount" default:""`
	// 优惠金额
	DiscountAmount uint64 `json:"discountAmount" default:""`
	// 备注
	Remark string `json:"remark" default:""`
	// 收件地址ID
	ShippingID uint64 `json:"shippingID" default:""`
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
	// 物料信息
	Goods []CreateOrderGoodsParams `json:"goods" default:""`
}
