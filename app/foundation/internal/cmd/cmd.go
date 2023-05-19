package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"go-metablox/app/foundation/internal/consts"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  consts.AppName,
		Usage: consts.AppUsage,
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := router()
			g.Log().Info(ctx, "server bootstrap!")
			s.Run()
			return nil
		},
	}
)
