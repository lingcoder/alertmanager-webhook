package bizctx

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/lingcoder/prometheus-notifier/internal/consts"
	"github.com/lingcoder/prometheus-notifier/internal/service"
	"sync"
)

type sBizCtx struct {
}

func init() {
	service.RegisterBizCtx(New())
}

func New() *sBizCtx {
	return &sBizCtx{}
}

func (s *sBizCtx) Init(r *ghttp.Request) {
	r.SetCtxVar(consts.BizCtxKey, sync.Map{})
}

func (s *sBizCtx) bizCtx(ctx context.Context) sync.Map {
	value := ctx.Value(consts.BizCtxKey)
	if value == nil {
		return sync.Map{}
	}
	if localCtx, ok := value.(sync.Map); ok {
		return localCtx
	}
	return sync.Map{}
}

func (s *sBizCtx) SetValue(ctx context.Context, key string, value any) {
	bizCtx := s.bizCtx(ctx)
	bizCtx.Store(key, value)
}

func (s *sBizCtx) GetValue(ctx context.Context, key string) any {
	bizCtx := s.bizCtx(ctx)
	value, ok := bizCtx.Load(key)
	if !ok {
		return nil
	}
	return value
}
