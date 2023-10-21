package controller

import (
	"context"
	v1 "github.com/lingcoder/prometheus-notifier/api/v1"
	"github.com/lingcoder/prometheus-notifier/internal/service"
)

type cFeishu struct{}

var Feishu = &cFeishu{}

func (c *cFeishu) Send(ctx context.Context, req *v1.FeishuReq) (*v1.BaseRes, error) {
	err := service.Feishu().Send(ctx, req)
	return &v1.BaseRes{}, err
}
