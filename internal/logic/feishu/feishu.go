package feishu

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	v1 "github.com/lingcoder/prometheus-notifier/api/v1"
	"github.com/lingcoder/prometheus-notifier/internal/consts"
	"github.com/lingcoder/prometheus-notifier/internal/model"
	"github.com/lingcoder/prometheus-notifier/internal/service"
	"time"
)

type sFeishu struct {
}

func init() {
	service.RegisterFeishu(New())
}

func New() *sFeishu {
	return &sFeishu{}
}

func (u *sFeishu) init(ctx context.Context) {

}

func (u *sFeishu) Send(ctx context.Context, req *v1.FeishuReq) error {

	// 重新解析模板
	alerts := req.WebhookBody.Alerts
	webhookAlert := model.WebhookAlert{
		IsRecovered: req.WebhookBody.Status == "resolved",
		SendingTime: time.Now().Format(time.RFC3339),
	}

	encodeString := gjson.MustEncodeString(req)
	fmt.Println(encodeString)

	for i := 0; i < len(alerts); i++ {
		alert := alerts[i]

		webhookAlert.Severity = alert.Labels["severity"]
		webhookAlert.RuleName = alert.Labels["alertname"]
		webhookAlert.RuleNotes = alert.Annotations["summary"]
		webhookAlert.TriggerTime = alert.StartsAt.Format(time.RFC3339)
		webhookAlert.LastEvalTime = alert.EndsAt.Format(time.RFC3339)
		webhookAlert.TriggerValue = alert.Annotations["value"]
		webhookAlert.TagsJSON = alert.Labels["tags"]

		content, err := g.View().ParseContent(ctx, consts.DefaultFeishuTemplate, gconv.Map(webhookAlert))

		if err != nil {

		}
		// 创建卡片配置，设置宽度和固定显示

		card := larkcard.NewMessageCard().
			Config(larkcard.NewMessageCardConfig().
				WideScreenMode(true).EnableForward(true)).
			Header(larkcard.NewMessageCardHeader().
				Template(func() string {
					if webhookAlert.IsRecovered {
						return larkcard.TemplateGreen
					} else {
						return larkcard.TemplateRed
					}
				}()).
				Title(larkcard.NewMessageCardPlainText().Content(g.I18n().T(ctx, "alert-notification")))).
			Elements([]larkcard.MessageCardElement{
				larkcard.NewMessageCardDiv().Text(larkcard.NewMessageCardLarkMd().Content(content)),
			})

		_, _ = g.Client().Post(ctx, req.URL, g.Map{
			"msg_type": "interactive",
			"card":     card,
		})

	}

	// 解析模板

	return nil
}
