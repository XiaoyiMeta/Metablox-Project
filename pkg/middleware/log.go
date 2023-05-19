package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Log(r *ghttp.Request) {
	ctx := r.GetCtx()
	var params interface{}
	if r.GetMultipartForm() != nil {
		params = ""
	} else {
		params = r.GetBodyString()
	}
	g.Log().Debugf(ctx, "Req %s %s %s %s", r.Request.Method, r.Request.URL, r.GetClientIp(), params)
	r.Middleware.Next()
	g.Log().Debugf(ctx, "Res %d %s", r.Response.Status, r.Response.BufferString())

}
