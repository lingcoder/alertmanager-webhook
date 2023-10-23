package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gsession"
	"github.com/lingcoder/prometheus-notifier/internal/controller"
	"github.com/lingcoder/prometheus-notifier/internal/service"
	"github.com/lingcoder/prometheus-notifier/pkg/base"
	"github.com/lingcoder/prometheus-notifier/pkg/middleware"
)

func router(ctx context.Context) *ghttp.Server {
	s := g.Server()
	WrapperOpenAPI(s)
	s.SetSessionStorage(gsession.NewStorageMemory())
	s.Use(
		middleware.CORS,
		middleware.Log,
		service.Middleware().Ctx,
		service.Middleware().I18n,
		middleware.ResponseHandler,
	)
	v1 := s.Group("/notifier")
	v1.POST("/welcome", controller.Index.Welcome)
	v1.POST("/feishu", controller.Feishu.Send)
	v1.POST("/custom", controller.Custom.Send)

	return s
}

func WrapperOpenAPI(s *ghttp.Server) {
	api := s.GetOpenApi()
	api.Config.CommonResponse = &base.Result{}
	api.Config.CommonResponseDataField = "Data"
	api.Config.IgnorePkgPath = true
}
