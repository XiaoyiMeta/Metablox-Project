package bizctx

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"go-metablox/app/foundation/internal/consts"
	"go-metablox/app/foundation/internal/model"
	"go-metablox/app/foundation/internal/service"
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
	r.SetCtxVar(consts.BizCtxUserKey, nil)
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

func (s *sBizCtx) SetUser(ctx context.Context, user *model.CurrentUser) {
	g.RequestFromCtx(ctx).SetCtxVar(consts.BizCtxUserKey, user)
}

func (s *sBizCtx) CurrentUser(ctx context.Context) *model.CurrentUser {
	value := ctx.Value(consts.BizCtxUserKey)
	if value == nil {
		return nil
	}
	if ctxUser, ok := value.(*model.CurrentUser); ok {
		return ctxUser
	}
	return nil
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
