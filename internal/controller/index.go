package controller

import (
	"context"
	v1 "github.com/lingcoder/prometheus-notifier/api/v1"
)

type cIndex struct{}

var Index = &cIndex{}

func (c *cIndex) Welcome(ctx context.Context, req *v1.WelcomeReq) (*v1.WelcomeRes, error) {

	return &v1.WelcomeRes{}, nil
}
