package response

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"go-metablox/pkg/biz"
)

// ResultExit Return Result and exit the current HTTP execution function
func ResultExit(r *ghttp.Request, result *biz.Result) {
	JSON(r, result.Code, result.Msg, result.Data)
	r.Exit()
}

// Result Return Result
func Result(r *ghttp.Request, result *biz.Result) {
	JSON(r, result.Code, result.Msg, result.Data)
}

// JSON Return standard JSON data。
func JSON(r *ghttp.Request, code int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = g.Map{}
	}
	r.Response.WriteJson(biz.Result{
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

// JSONRedirect Return standard JSON data to guide the client to jump, and exit the current HTTP execution function
func JSONRedirect(r *ghttp.Request, code int, message, redirect string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(biz.Result{
		Code:     code,
		Msg:      message,
		Data:     responseData,
		Redirect: redirect,
	})
}

// JSONRedirectExit Return standard JSON data to guide the client to jump and exit the current HTTP execution function.
func JSONRedirectExit(r *ghttp.Request, code int, message, redirect string, data ...interface{}) {
	JSONRedirect(r, code, message, redirect, data...)
	r.Exit()
}
