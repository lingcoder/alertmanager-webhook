package response

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/lingcoder/prometheus-notifier/pkg/base"
	"net/http"
)

// JSON Return standard JSON data。
func JSON(r *ghttp.Request, code int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = g.Map{}
	}
	r.Response.WriteJson(base.Result{
		Code: code,
		Msg:  message,
		Data: responseData,
	})
}

// JSONExit Return standard JSON data and exit the current HTTP execution function
func JSONExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	JSON(r, code, message, data...)
	r.Exit()
}

// AuthorizationExit Return standard JSON data and exit the HTTP execution function
func AuthorizationExit(r *ghttp.Request, message string, data ...interface{}) {
	JSON(r, http.StatusUnauthorized, message, data...)
	r.Exit()
}
