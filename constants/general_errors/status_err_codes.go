package general_errors

import (
	"net/http"

	"github.com/eden-framework/courier/status_error"
)

//go:generate eden generate error
const ServiceStatusErrorCode = 100

const (
	// 请求参数错误
	BadRequest status_error.StatusErrorCode = http.StatusBadRequest*1e6 + ServiceStatusErrorCode + iota
	// @errTalk 商品库存不足无法创建订单，请刷新页面重试
	GoodsInventoryShortage
	// @errTalk 商品库存充足无法创建预订订单，请刷新页面重试
	GoodsInventorySufficient
	// @errTalk 优惠活动未开始
	DiscountNotStart
	// @errTalk 优惠活动已结束
	DiscountEnd
	// @errTalk 行政区划冲突
	RegionsConflict
)

const (
	// 未找到
	NotFound status_error.StatusErrorCode = http.StatusNotFound*1e6 + ServiceStatusErrorCode + iota
	// @errTalk 管理员未找到
	AdminNotFound
	// @errTalk 用户未找到
	UserNotFound
	// @errTalk 订单未找到
	OrderNotFound
	// @errTalk 支付流水号未找到
	PaymentFlowNotFound
	// @errTalk 商品未找到
	GoodsNotFound
	// @errTalk 结算单未找到
	SettlementFlowNotFound
	// @errTalk 管理员登录已失效
	AdminTokenExpired
	// @errTalk 运费模板未找到
	FreightTemplateNotFound
	// @errTalk 运费模板规则未找到
	FreightRuleNotFound
)

const (
	// @errTalk 未授权
	Unauthorized status_error.StatusErrorCode = http.StatusUnauthorized*1e6 + ServiceStatusErrorCode + iota
	// @errTalk 用户名或密码错误
	InvalidUserNamePassword
)

const (
	// @errTalk 操作冲突
	Conflict status_error.StatusErrorCode = http.StatusConflict*1e6 + ServiceStatusErrorCode + iota
	// @errTalk 支付金额与交易单金额不一致
	FlowAmountIncorrect
	// @errTalk 订单状态流转错误
	OrderStatusFlowIncorrect
	// @errTalk 支付状态流转错误
	PaymentStatusFlowIncorrect
	// @errTalk 折扣金额超过订单总额
	DiscountAmountOverflow
)

const (
	// @errTalk 不允许操作
	Forbidden status_error.StatusErrorCode = http.StatusForbidden*1e6 + ServiceStatusErrorCode + iota
	// @errTalk 订单不可重复取消
	OrderCanceled
	// @errTalk 订单已支付不允许变更金额
	NotAllowedChangeAmount
	// @errTalk 订单状态不允许变更收件信息
	NotAllowedChangeLogistics
)

const (
	// 内部处理错误
	InternalError status_error.StatusErrorCode = http.StatusInternalServerError*1e6 + ServiceStatusErrorCode + iota
)

const (
	// 上游错误
	BadGateway status_error.StatusErrorCode = http.StatusBadGateway*1e6 + ServiceStatusErrorCode + iota
)
