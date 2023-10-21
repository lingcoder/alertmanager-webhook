package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/lingcoder/prometheus-notifier/internal/service"
)

type sMiddleware struct {
}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *sMiddleware) Log(r *ghttp.Request) {
	ctx := r.GetCtx()
	g.Log().Debugf(ctx, "Request %s %s %s %s", r.Request.Method, r.Request.URL, r.Request.Host, r.GetBodyString())
	r.Middleware.Next()
	g.Log().Debugf(ctx, "Response %d %s", r.Response.Status, r.Response.BufferString())
}

func (s *sMiddleware) Ctx(r *ghttp.Request) {
	service.BizCtx().Init(r)
	r.Middleware.Next()
}

func (s *sMiddleware) I18n(r *ghttp.Request) {
	lang := r.GetQuery("lang").String()
	if lang != "" {
		r.SetCtx(gi18n.WithLanguage(r.Context(), lang))
	}
	r.Middleware.Next()
}
