package general_errors

import (
	"github.com/eden-framework/courier/status_error"
)

func init() {
	status_error.StatusErrorCodes.Register("OrderCanceled", 403000101, "订单不可重复取消", "", true)
	status_error.StatusErrorCodes.Register("SettlementFlowNotFound", 404000106, "结算单未找到", "", true)
	status_error.StatusErrorCodes.Register("FreightRuleNotFound", 404000109, "运费模板规则未找到", "", true)
	status_error.StatusErrorCodes.Register("InvalidUserNamePassword", 401000101, "用户名或密码错误", "", true)
	status_error.StatusErrorCodes.Register("GoodsNotFound", 404000105, "商品未找到", "", true)
	status_error.StatusErrorCodes.Register("BadGateway", 502000100, "上游错误", "", false)
	status_error.StatusErrorCodes.Register("DiscountNotStart", 400000103, "优惠活动未开始", "", true)
	status_error.StatusErrorCodes.Register("PaymentFlowNotFound", 404000104, "支付流水号未找到", "", true)
	status_error.StatusErrorCodes.Register("GoodsInventorySufficient", 400000102, "商品库存充足无法创建预订订单，请刷新页面重试", "", true)
	status_error.StatusErrorCodes.Register("NotAllowedChangeAmount", 403000102, "订单已支付不允许变更金额", "", true)
	status_error.StatusErrorCodes.Register("Forbidden", 403000100, "不允许操作", "", true)
	status_error.StatusErrorCodes.Register("NotAllowedChangeLogistics", 403000103, "订单状态不允许变更收件信息", "", true)
	status_error.StatusErrorCodes.Register("FlowAmountIncorrect", 409000101, "支付金额与交易单金额不一致", "", true)
	status_error.StatusErrorCodes.Register("DiscountAmountOverflow", 409000104, "折扣金额超过订单总额", "", true)
	status_error.StatusErrorCodes.Register("AdminNotFound", 404000101, "管理员未找到", "", true)
	status_error.StatusErrorCodes.Register("RegionsConflict", 400000105, "行政区划冲突", "", true)
	status_error.StatusErrorCodes.Register("OrderNotFound", 404000103, "订单未找到", "", true)
	status_error.StatusErrorCodes.Register("GoodsInventoryShortage", 400000101, "商品库存不足无法创建订单，请刷新页面重试", "", true)
	status_error.StatusErrorCodes.Register("FreightTemplateNotFound", 404000108, "运费模板未找到", "", true)
	status_error.StatusErrorCodes.Register("Unauthorized", 401000100, "未授权", "", true)
	status_error.StatusErrorCodes.Register("BadRequest", 400000100, "请求参数错误", "", false)
	status_error.StatusErrorCodes.Register("DiscountEnd", 400000104, "优惠活动已结束", "", true)
	status_error.StatusErrorCodes.Register("Conflict", 409000100, "操作冲突", "", true)
	status_error.StatusErrorCodes.Register("InternalError", 500000100, "内部处理错误", "", false)
	status_error.StatusErrorCodes.Register("OrderStatusFlowIncorrect", 409000102, "订单状态流转错误", "", true)
	status_error.StatusErrorCodes.Register("PaymentStatusFlowIncorrect", 409000103, "支付状态流转错误", "", true)
	status_error.StatusErrorCodes.Register("AdminTokenExpired", 404000107, "管理员登录已失效", "", true)
	status_error.StatusErrorCodes.Register("NotFound", 404000100, "未找到", "", false)
	status_error.StatusErrorCodes.Register("UserNotFound", 404000102, "用户未找到", "", true)

}
