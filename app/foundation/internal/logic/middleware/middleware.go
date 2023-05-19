package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"go-metablox/app/foundation/internal/service"
)

const AuthHeader = "Authorization"

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

func (s *sMiddleware) Ctx(r *ghttp.Request) {
	service.BizCtx().Init(r)
	r.Middleware.Next()
}
