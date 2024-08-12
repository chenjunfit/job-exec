package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/os/gcmd"
	"job-exec/internal/app/server/controller/taskreport"
	"job-exec/internal/app/server/router"
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
			gsvc.SetRegistry(etcd.New(`127.0.0.1:2379`))
			s := grpcx.Server.New()
			taskreport.Register(s)
			s.Run()
			return nil
		},
	}
)
