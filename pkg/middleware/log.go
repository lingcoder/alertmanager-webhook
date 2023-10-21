package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"strings"
)

func Log(r *ghttp.Request) {
	ctx := r.GetCtx()
	var params interface{}
	if r.GetMultipartForm() != nil {
		params = ""
	} else {
		params = r.GetBodyString()
	}
	g.Log().Debugf(ctx, "Req %s %s IP=%s %s", r.Request.Method, r.Request.URL, r.GetClientIp(), params)
	r.Middleware.Next()
	contentDisposition := r.Response.Header().Get("Content-Disposition")
	response := ""
	if !strings.HasPrefix(contentDisposition, "attachment") || !strings.HasPrefix(contentDisposition, "text/event-stream") {
		response = r.Response.BufferString()
	}
	g.Log().Debugf(ctx, "Res %d %s", r.Response.Status, response)

}
