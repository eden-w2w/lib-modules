package errors

import (
	"github.com/eden-framework/courier/status_error"
)

func init() {
	status_error.StatusErrorCodes.Register("InternalError", 500103000, "内部处理错误", "", false)
	status_error.StatusErrorCodes.Register("Forbidden", 403103000, "不允许操作", "", true)
	status_error.StatusErrorCodes.Register("PaymentStatusFlowIncorrect", 409103003, "支付状态流转错误", "", true)
	status_error.StatusErrorCodes.Register("OrderCanceled", 403103001, "订单不可重复取消", "", true)
	status_error.StatusErrorCodes.Register("OrderNotFound", 404103002, "订单未找到", "", true)
	status_error.StatusErrorCodes.Register("Unauthorized", 401103000, "未授权", "", true)
	status_error.StatusErrorCodes.Register("UserNotFound", 404103001, "用户未找到", "", true)
	status_error.StatusErrorCodes.Register("PaymentFlowNotFound", 404103003, "支付流水号未找到", "", true)
	status_error.StatusErrorCodes.Register("OrderStatusFlowIncorrect", 409103002, "订单状态流转错误", "", true)
	status_error.StatusErrorCodes.Register("BadGateway", 502103000, "上游错误", "", false)
	status_error.StatusErrorCodes.Register("NotFound", 404103000, "未找到", "", false)
	status_error.StatusErrorCodes.Register("FlowAmountIncorrect", 409103001, "支付金额与交易单金额不一致", "", true)
	status_error.StatusErrorCodes.Register("Conflict", 409103000, "操作冲突", "", true)
	status_error.StatusErrorCodes.Register("BadRequest", 400103000, "请求参数错误", "", false)

}
