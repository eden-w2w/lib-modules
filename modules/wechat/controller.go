package wechat

import (
	"context"
	"fmt"
	errors "github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/wechatpay-go/core"
	"github.com/eden-w2w/wechatpay-go/core/auth/verifiers"
	"github.com/eden-w2w/wechatpay-go/core/downloader"
	"github.com/eden-w2w/wechatpay-go/core/notify"
	"github.com/eden-w2w/wechatpay-go/core/option"
	"github.com/eden-w2w/wechatpay-go/services/payments"
	"github.com/eden-w2w/wechatpay-go/services/payments/jsapi"
	"github.com/eden-w2w/wechatpay-go/utils"
	w "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	programConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/encryptor"
	"github.com/silenceper/wechat/v2/miniprogram/qrcode"
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var controller *Controller

func GetController() *Controller {
	if controller == nil {
		controller = &Controller{isInit: false}
	}
	return controller
}

type Wechat struct {
	// 小程序AppID
	AppID string
	// 小程序AppSecret
	AppSecret string
	// 微信商户ID
	MerchantID string
	// 微信商户证书序列号
	MerchantCertSerialNo string
	// 微信商户证书私钥
	MerchantPK string
	// 微信商户APIv3密钥
	MerchantSecret string
	// 微信支付商品描述
	ProductionDesc string
	// 微信支付回调地址
	NotifyUrl string
	// 启用微信支付
	EnableWechatPay bool

	// 定时查单任务配置
	FetchWechatPaymentStatusTask string

	// 订单确认通知订阅消息模板
	ConfirmMessageTemplateID string
	// 发货通知订阅消息模板
	DispatchMessageTemplateID string
	// 订单完成通知订阅消息模板
	CompleteMessageTemplateID string
	// 订单相关通知页面地址
	OrderPage string
	// 跳转小程序类型
	ProgramState string
}

type Controller struct {
	wc           *w.Wechat
	program      *miniprogram.MiniProgram
	payClient    *core.Client
	jsapiService jsapi.JsapiApiService
	config       Wechat
	isInit       bool
}

func (c *Controller) Init(wechatConfig Wechat) {
	wc := w.NewWechat()
	memory := cache.NewMemory()
	program := wc.GetMiniProgram(
		&programConfig.Config{
			AppID:     wechatConfig.AppID,
			AppSecret: wechatConfig.AppSecret,
			Cache:     memory,
		},
	)

	var client *core.Client
	var service jsapi.JsapiApiService
	if wechatConfig.EnableWechatPay {
		mchPK, err := utils.LoadPrivateKey(wechatConfig.MerchantPK)
		if err != nil {
			logrus.Panicf("[wechat.newController] utils.LoadPrivateKey err: %v", err)
		}
		ctx := context.Background()
		opts := []core.ClientOption{
			option.WithWechatPayAutoAuthCipher(
				wechatConfig.MerchantID,
				wechatConfig.MerchantCertSerialNo,
				mchPK,
				wechatConfig.MerchantSecret,
			),
		}
		client, err = core.NewClient(ctx, opts...)
		if err != nil {
			logrus.Panicf("[wechat.newController] core.NewClient err: %v", err)
		}

		service = jsapi.JsapiApiService{
			Client: client,
		}
	}

	c.wc = wc
	c.program = program
	c.payClient = client
	c.jsapiService = service
	c.config = wechatConfig
	c.isInit = true
}

func (c Controller) Code2Session(code string) (*auth.ResCode2Session, error) {
	if !c.isInit {
		logrus.Panicf("[WechatController] not Init")
	}
	resp, err := c.program.GetAuth().Code2Session(code)
	if err != nil {
		logrus.Errorf("[Code2Session] c.program.GetAuth().Code2Session(code) err: %v, code: %s", err, code)
		return nil, errors.BadGateway
	}
	return &resp, nil
}

func (c Controller) ExchangeEncryptedData(sessionKey string, params WechatUserInfo) (*encryptor.PlainData, error) {
	if !c.isInit {
		logrus.Panicf("[WechatController] not Init")
	}
	plain, err := c.program.GetEncryptor().Decrypt(sessionKey, params.EncryptedData, params.IV)
	if err != nil {
		logrus.Errorf(
			"[ExchangeEncryptedData] program.GetEncryptor().Decrypt err: %v, sessionKey: %s, params: %+v",
			err,
			sessionKey,
			params,
		)
		return nil, errors.InternalError
	}

	return plain, nil
}

func (c Controller) GetUnlimitedQrCode(params qrcode.QRCoder) (buffer []byte, err error) {
	if !c.isInit {
		logrus.Panicf("[WechatController] not Init")
	}
	buffer, err = c.program.GetQRCode().GetWXACodeUnlimit(params)
	if err != nil {
		logrus.Errorf(
			"[GetUnlimitedQrCode] c.program.GetQRCode().GetWXACodeUnlimit(params) err: %v, params: %+v",
			err,
			params,
		)
		return nil, errors.BadGateway
	}

	return
}

func (c Controller) CreatePrePayment(
	ctx context.Context,
	order *databases.Order,
	flow *databases.PaymentFlow,
	payer *databases.User,
) (resp *jsapi.PrepayWithRequestPaymentResponse, err error) {
	if !c.isInit {
		logrus.Panicf("[WechatController] not Init")
	}
	if c.payClient == nil {
		return
	}
	if !c.config.EnableWechatPay {
		return
	}
	request := jsapi.PrepayRequest{
		Appid:         core.String(c.config.AppID),
		Mchid:         core.String(c.config.MerchantID),
		Description:   core.String(c.config.ProductionDesc),
		OutTradeNo:    core.String(fmt.Sprintf("%d", flow.FlowID)),
		TimeExpire:    core.Time(time.Time(flow.ExpiredAt)),
		Attach:        nil,
		NotifyUrl:     core.String(c.config.NotifyUrl),
		GoodsTag:      nil,
		LimitPay:      nil,
		SupportFapiao: nil,
		Amount: &jsapi.Amount{
			Total:    core.Int64(int64(flow.Amount)),
			Currency: nil,
		},
		Payer: &jsapi.Payer{
			Openid: core.String(payer.OpenID),
		},
		Detail:     nil,
		SceneInfo:  nil,
		SettleInfo: nil,
	}
	resp, _, err = c.jsapiService.PrepayWithRequestPayment(ctx, request)
	if err != nil {
		logrus.Errorf("[CreatePrePayment] service.PrepayWithRequestPayment err: %v, request: %+v", err, request)
		return nil, errors.BadGateway
	}
	return
}

func (c Controller) ParseWechatPaymentNotify(ctx context.Context, request *http.Request) (
	*notify.Request,
	*payments.Transaction,
	error,
) {
	if !c.isInit {
		logrus.Panicf("[WechatController] not Init")
	}
	certVisitor := downloader.MgrInstance().GetCertificateVisitor(c.config.MerchantID)
	handler := notify.NewNotifyHandler(c.config.MerchantSecret, verifiers.NewSHA256WithRSAVerifier(certVisitor))

	transaction := new(payments.Transaction)
	notifyReq, err := handler.ParseNotifyRequest(ctx, request, transaction)
	if err != nil {
		logrus.Errorf("[ParseWechatPaymentNotify] handler.ParseNotifyRequest err: %v", err)
		return nil, nil, errors.InternalError
	}

	return notifyReq, transaction, nil
}

func (c Controller) QueryOrderByOutTradeNo(req jsapi.QueryOrderByOutTradeNoRequest) (
	resp *payments.Transaction,
	err error,
) {
	resp, _, err = c.jsapiService.QueryOrderByOutTradeNo(context.Background(), req)
	if err != nil {
		logrus.Warningf(
			"[QueryOrderByOutTradeNo] c.jsapiService.QueryOrderByOutTradeNo err: %v, request: %+v",
			err,
			req,
		)
	}
	return
}

func (c Controller) CloseOrder(req jsapi.CloseOrderRequest) error {
	_, err := c.jsapiService.CloseOrder(context.Background(), req)
	if err != nil {
		logrus.Warningf(
			"[CloseOrder] c.jsapiService.CloseOrder err: %v, request: %+v",
			err,
			req,
		)
	}
	return err
}

func (c Controller) SendSubscribeMessage(msg *subscribe.Message) error {
	err := c.program.GetSubscribe().Send(msg)
	if err != nil {
		logrus.Warningf("[SendSubscribeMessage] c.program.GetSubscribe().Send err: %v, msg: %+v", err, msg)
	}
	return err
}

func (c Controller) GetTradeBill(req jsapi.TradeBillRequest) (*jsapi.BillResponse, error) {
	resp, _, err := c.jsapiService.TradeBill(context.Background(), req)
	if err != nil {
		logrus.Errorf("[GetTradeBill] c.jsapiService.TradeBill err: %v, req: %+v", err, req)
		return nil, errors.BadGateway
	}
	return resp, nil
}

func (c Controller) DownloadURL(bill jsapi.BillResponse) (data []byte, err error) {
	data, _, err = c.jsapiService.DownloadURL(context.Background(), bill)
	if err != nil {
		logrus.Errorf("[DownloadURL] c.jsapiService.DownloadURL err: %v, req: %+v", err, bill)
		return nil, errors.BadGateway
	}
	return
}
