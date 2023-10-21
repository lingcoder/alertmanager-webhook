package custom

import (
	"context"
	v1 "github.com/lingcoder/prometheus-notifier/api/v1"
	"github.com/lingcoder/prometheus-notifier/internal/service"
)

type sCustom struct {
}

func init() {
	service.RegisterCustom(New())
}

func New() *sCustom {
	return &sCustom{}
}

func (u *sCustom) init(ctx context.Context) {

}

func (u *sCustom) Send(ctx context.Context, req *v1.CustomReq) error {

	return nil
}
