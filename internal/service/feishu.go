// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/lingcoder/prometheus-notifier/api/v1"
)

type (
	IFeishu interface {
		Send(ctx context.Context, req *v1.FeishuReq) error
	}
)

var (
	localFeishu IFeishu
)

func Feishu() IFeishu {
	if localFeishu == nil {
		panic("implement not found for interface IFeishu, forgot register?")
	}
	return localFeishu
}

func RegisterFeishu(i IFeishu) {
	localFeishu = i
}
