package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"go-metablox/pkg/biz"
	"go-metablox/pkg/response"
	"net/http"
)

func ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()
	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		g.Log().Info(r.GetCtx(), err.Error())
		switch code {
		case gcode.CodeNil:
			code = biz.CodeError
			msg = code.Detail().(string)
		case gcode.CodeDbOperationError, gcode.CodeInternalError, gcode.CodeNotImplemented:
			msg = code.Message()
		default:
			msg = err.Error()
		}

	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = biz.CodeNotFound
			case http.StatusForbidden:
				code = biz.CodeUnauthorized
			default:
				code = biz.CodeError
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = biz.CodeOK
			msg = biz.CodeOK.Message()
		}
	}

	response.JSON(r, code.Code(), msg, res)
}
