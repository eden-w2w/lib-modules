package wechat

type WechatNotifyResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type WechatUserInfo struct {
	// 微信的加密开放数据
	EncryptedData string `in:"body" json:"encryptedData"`
	// 微信的加密原始数据
	RawData string `in:"body" json:"rawData" default:""`
	// 微信的签名数据
	Signature string `json:"signature" default:""`
	// 微信的初始向量
	IV string `json:"iv"`
}
