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
	ICustom interface {
		Send(ctx context.Context, req *v1.CustomReq) error
	}
)

var (
	localCustom ICustom
)

func Custom() ICustom {
	if localCustom == nil {
		panic("implement not found for interface ICustom, forgot register?")
	}
	return localCustom
}

func RegisterCustom(i ICustom) {
	localCustom = i
}
