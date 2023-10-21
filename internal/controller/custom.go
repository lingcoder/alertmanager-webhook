package controller

import (
	"context"
	v1 "github.com/lingcoder/prometheus-notifier/api/v1"
	"github.com/lingcoder/prometheus-notifier/internal/service"
)

type cCustom struct{}

var Custom = &cCustom{}

func (c *cCustom) Send(ctx context.Context, req *v1.CustomReq) (*v1.BaseRes, error) {
	err := service.Custom().Send(ctx, req)
	return &v1.BaseRes{}, err
}
