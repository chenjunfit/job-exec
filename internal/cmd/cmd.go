package cmd

import (
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"job-exec/internal/router"
)

var (
	Http = gcmd.Command{
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				router.R.BindController(ctx, group)
			})
			s.Run()
			return nil
		},
	}
	Rpc = gcmd.Command{
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := grpcx.Server.New()
			s.Run()
			return nil
		},
	}
)
