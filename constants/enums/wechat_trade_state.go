package enums

//go:generate eden generate enum --type-name=WechatTradeState
// api:enum
type WechatTradeState uint8

// 微信支付状态
const (
	WECHAT_TRADE_STATE_UNKNOWN     WechatTradeState = iota
	WECHAT_TRADE_STATE__SUCCESS                     // 支付成功
	WECHAT_TRADE_STATE__REFUND                      // 转入退款
	WECHAT_TRADE_STATE__NOTPAY                      // 未支付
	WECHAT_TRADE_STATE__CLOSED                      // 已关闭
	WECHAT_TRADE_STATE__REVOKED                     // 已撤销
	WECHAT_TRADE_STATE__USERPAYING                  // 支付中
	WECHAT_TRADE_STATE__PAYERROR                    // 支付失败
)

func (v WechatTradeState) IsEqual(status PaymentStatus) bool {
	if v == WECHAT_TRADE_STATE__NOTPAY && status == PAYMENT_STATUS__CREATED {
		return true
	}
	if v == WECHAT_TRADE_STATE__SUCCESS && status == PAYMENT_STATUS__SUCCESS {
		return true
	}
	if v == WECHAT_TRADE_STATE__USERPAYING && status == PAYMENT_STATUS__PROCESS {
		return true
	}
	if v == WECHAT_TRADE_STATE__PAYERROR && status == PAYMENT_STATUS__FAIL {
		return true
	}
	if v == WECHAT_TRADE_STATE__CLOSED && status == PAYMENT_STATUS__CLOSED {
		return true
	}
	return false
}

func (v WechatTradeState) IsEnding() bool {
	if v == WECHAT_TRADE_STATE__NOTPAY ||
		v == WECHAT_TRADE_STATE__USERPAYING {
		return false
	}
	return true
}

func (v WechatTradeState) IsSuccess() bool {
	if v == WECHAT_TRADE_STATE__SUCCESS {
		return true
	}
	return false
}

func (v WechatTradeState) IsFail() bool {
	if v == WECHAT_TRADE_STATE__PAYERROR ||
		v == WECHAT_TRADE_STATE__CLOSED ||
		v == WECHAT_TRADE_STATE__REVOKED {
		return true
	}
	return false
}

func (v WechatTradeState) IsRefund() bool {
	if v == WECHAT_TRADE_STATE__REFUND {
		return true
	}
	return false
}
