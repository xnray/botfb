package controller

import (
	"botfb/internal/consts"
	"botfb/internal/model"
	"botfb/internal/model/entity"
	"botfb/internal/service"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"html/template"
	"strings"
)

type eventHandle struct {
	event    AppEvent
	endPoint string
}

type AppEvent struct {
	EventType   int
	PageID      string
	RecipientID int64
}

func NewEventHandle(ctx context.Context) *eventHandle {
	baseUrl := g.Cfg().MustGet(ctx, "fb.baseUrl").String()
	pageID := g.Cfg().MustGet(ctx, "fb.pageID").String()
	accessToken := g.Cfg().MustGet(ctx, "fb.pageAccessToken").String()
	endpoint := baseUrl + "/" + pageID + "/messages?" + accessToken
	return &eventHandle{
		endPoint: endpoint,
	}
}

func (e *eventHandle) Dispatch(ctx context.Context, payload string) error {
	var data AppEvent
	if err := gjson.Unmarshal([]byte(payload), &data); err != nil {
		return err
	}
	e.event = data
	event := data.EventType
	switch event {
	case int(consts.EventAddToCart):
		// 处理添加到购物车事件
		break
	case int(consts.EventRemoveFromCart):
		// 处理从购物车移除事件
		break
	case int(consts.EventCheckout):
		// 处理结账事件
		break
	case int(consts.EventPaymentSuccess):
		// 处理支付成功事件
		break
	case int(consts.EventPaymentFailure):
		// 处理支付失败事件
		break
	case int(consts.EventOrderPlaced):
		// 处理订单已下单事件
		break
	case int(consts.EventOrderCancelled):
		// 处理订单取消事件
		break
	case int(consts.EventOrderShipped):
		// 处理订单已发货事件
		break
	default:
		return gerror.New("Unsupported event types")
	}

	return nil
}

// templateData := MessageTemplate{
//		TemplateID:      "1",
//		TemplateName:    "Welcome",
//		TemplateContent: "Hello {{.name}}, welcome to our platform. Your ID is {{.id}}.",
//		Variables:       []byte(`{"name": "John", "id": "123"}`),
//	}

func (e *eventHandle) addToCart(ctx context.Context) error {

	//curl -X POST -H "Content-Type: application/json" -d '{
	//  "recipient":{
	//    "id":"PSID"
	//  },
	//  "messaging_type": "RESPONSE",
	//  "message":{
	//    "text":"Hello, world!"
	//  }
	//}' "https://graph.facebook.com/LATEST-API-VERSION/PAGE-ID/messages?access_token=PAGE-ACCESS-TOKEN"

	// get data from template database
	templateID, err := getTemplateIDByEvent(e.event.EventType)
	if err != nil {
		return err
	}
	// query template
	msgTp, err := service.MessageTemplate().View(ctx, templateID)
	if err != nil {
		return err
	}
	// parse template
	content, err := parseTemplate(msgTp)
	if err != nil {
		return err
	}
	sendMsg := model.SendMessage{
		MessagingType: "RESPONSE",
		Recipient:     model.Recipient{ID: e.event.RecipientID},
		Message: model.MessageData{
			Text: content,
		},
	}
	// send message to customer
	r, err := g.Client().Post(
		ctx,
		e.endPoint,
		sendMsg,
	)
	if err != nil {
		return err
	}
	defer r.Close()
	g.Log().Info(ctx, r.ReadAllString())
	return nil
}

func parseTemplate(src entity.MessageTemplate) (string, error) {
	// 解析模板
	tmpl, err := template.New("tp").Parse(src.TemplateContent)
	if err != nil {
		return "", err
	}

	// 解析变量值
	var variables map[string]interface{}
	err = gjson.Unmarshal([]byte(src.Variables), &variables)
	if err != nil {
		return "", err
	}

	// 执行渲染
	var parsedContent strings.Builder
	err = tmpl.Execute(&parsedContent, struct{ Name string }{"John"})
	if err != nil {
		return "", err
	}

	// 获取渲染结果
	result := parsedContent.String()
	return result, nil
}

func getTemplateIDByEvent(eventType int) (string, error) {
	var templateID string
	switch eventType {
	case int(consts.EventAddToCart):
		// 处理添加到购物车事件
		templateID = "1"
		break
	case int(consts.EventRemoveFromCart):
		// 处理从购物车移除事件
		templateID = "2"
		break
	case int(consts.EventCheckout):
		// 处理结账事件
		templateID = "3"
		break
	case int(consts.EventPaymentSuccess):
		// 处理支付成功事件
		templateID = "3"
		break
	case int(consts.EventPaymentFailure):
		// 处理支付失败事件
		templateID = "4"
		break
	case int(consts.EventOrderPlaced):
		// 处理订单已下单事件
		templateID = "5"
		break
	case int(consts.EventOrderCancelled):
		// 处理订单取消事件
		templateID = "6"
		break
	case int(consts.EventOrderShipped):
		// 处理订单已发货事件
		templateID = "7"
		break
	default:
		return "", gerror.New("Unsupported event types")
	}
	return templateID, nil
}
