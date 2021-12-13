package webhook

import (
	"errors"
	"fmt"
	"github.com/eden-framework/courier/client"
	"github.com/eden-framework/sqlx"
	"github.com/eden-w2w/lib-modules/constants/enums"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules/settings"
	"github.com/sirupsen/logrus"
	"net/url"
	"strconv"
)

var webhook *Client

type Request interface {
	SetKey(key string)
}

func GetInstance() *Client {
	if webhook == nil {
		webhook = &Client{
			isInit:    false,
			isEnabled: false,
		}
	}
	return webhook
}

type Client struct {
	RawURL     string
	webhookURL string
	key        string

	client    *client.Client
	db        sqlx.DBExecutor
	isInit    bool
	isEnabled bool
}

func (c *Client) MustInit(db sqlx.DBExecutor) {
	err := c.Init(db)
	if err != nil {
		logrus.Panicf("[WebhookClient] MustInit err: %v", err)
	}
}

func (c *Client) Init(db sqlx.DBExecutor) error {
	c.db = db
	c.isInit = true

	model := settings.GetController().GetSetting()
	if model == nil {
		return errors.New("setting model not exist")
	}

	if model.WebhookEnabled.False() {
		return nil
	}
	if model.WebhookURL == "" {
		return nil
	}
	c.RawURL = model.WebhookURL
	webhookUrl, err := url.Parse(model.WebhookURL)
	if err != nil {
		return fmt.Errorf("url.Parse(model.WebhookURL) err: %v, url: %s", err, model.WebhookURL)
	}
	c.webhookURL = webhookUrl.Path
	var port = int64(0)
	if webhookUrl.Scheme == "http" {
		port = 80
	} else {
		port = 443
	}
	if webhookUrl.Port() != "" {
		port, err = strconv.ParseInt(webhookUrl.Port(), 10, 64)
		if err != nil {
			return fmt.Errorf("strconv.ParseInt(webhookUrl.Port(), 10, 64) err: %v, port: %s", err, webhookUrl.Port())
		}
	}
	c.key = webhookUrl.Query().Get("key")
	c.client = &client.Client{
		Name: "webhook-client",
		Host: webhookUrl.Host,
		Mode: webhookUrl.Scheme,
		Port: int16(port),
	}
	c.isEnabled = model.WebhookEnabled.True()
	return nil
}

func (c Client) send(req Request) error {
	if !c.isEnabled {
		return nil
	}

	req.SetKey(c.key)
	request := c.client.Request("", "POST", c.webhookURL, req)
	return request.Do().Err
}

func (c Client) SendMessage(msgType enums.WebhookMsgType, content string) error {
	req := &SendMessageRequest{
		Data: SendMessageBody{
			MsgType: msgType,
		},
	}
	if msgType == enums.WEBHOOK_MSG_TYPE_UNKNOWN {
		return fmt.Errorf("[WebhookClient] SendMessage err: unknown msgType: %s", msgType)
	} else if msgType == enums.WEBHOOK_MSG_TYPE__text {
		req.Data.Text = &SendMessageText{
			Content:       content,
			MentionedList: []string{"@all"},
		}
	} else if msgType == enums.WEBHOOK_MSG_TYPE__markdown {
		req.Data.Markdown = &SendMessageMarkdown{
			Content: content,
		}
	}
	return c.send(req)
}

func (c Client) SendCreateOrder(model *databases.Order) error {
	return c.SendMessage(
		enums.WEBHOOK_MSG_TYPE__markdown,
		fmt.Sprintf(
			"**收到新创建订单请注意处理**\n"+
				">订单号：<font color=\"warning\">%d</font>\n"+
				"金额：<font color=\"warning\">%.2f</font>", model.OrderID, float64(model.TotalPrice)/100.0,
		),
	)
}

func (c Client) SendPayment(
	model *databases.Order,
	paymentFlow *databases.PaymentFlow,
	logistics *databases.OrderLogistics,
) error {
	return c.SendMessage(
		enums.WEBHOOK_MSG_TYPE__markdown,
		fmt.Sprintf(
			"**收到订单付款信息请注意处理**\n\n"+
				"### 订单信息\n"+
				">订单号：<font color=\"warning\">%d</font>\n"+
				">订单金额：<font color=\"warning\">%.2f</font>\n"+
				">优惠金额：<font color=\"warning\">%.2f</font>\n"+
				">运费金额：<font color=\"warning\">%.2f</font>\n"+
				">实际金额：<font color=\"warning\">%.2f</font>\n"+
				">支付金额：<font color=\"warning\">%.2f</font>\n\n"+
				"### 物流信息\n"+
				">收件人：<font color=\"comment\">%s</font>\n"+
				">收件地址：<font color=\"comment\">%s</font>\n"+
				">手机号：<font color=\"comment\">%s</font>",
			model.OrderID,
			float64(model.TotalPrice)/100.0,
			float64(model.DiscountAmount)/100.0,
			float64(model.FreightAmount)/100.0,
			float64(model.ActualAmount)/100.0,
			float64(paymentFlow.ActualAmount)/100.0,
			logistics.Recipients,
			logistics.ShippingAddr,
			logistics.Mobile,
		),
	)
}
