package feishu

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	v1 "github.com/lingcoder/prometheus-notifier/api/v1"
	"github.com/lingcoder/prometheus-notifier/internal/service"
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

	// 根据语言选择合适的模板
	template := g.I18n().T(ctx, "feishu-template")

	// 解析模板
	content, err := g.View().ParseContent(ctx, template, gconv.Map(req.Body))
	// 创建卡片配置，设置宽度和固定显示

	card := larkcard.NewMessageCard().
		Config(larkcard.NewMessageCardConfig().
			WideScreenMode(true).EnableForward(true)).
		Header(larkcard.NewMessageCardHeader().
			Title(larkcard.NewMessageCardPlainText().Content(g.I18n().T(ctx, "feishu-title")))).
		Elements([]larkcard.MessageCardElement{
			larkcard.NewMessageCardDiv().Text(larkcard.NewMessageCardPlainText().Content(content)),
		})

	_, err = g.Client().Post(ctx, req.URL, g.Map{
		"msg_type": "interactive",
		"card":     card,
	})
	return err
}
