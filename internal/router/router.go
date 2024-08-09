package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"job-exec/internal/controller/jobexec"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/jobexec", func(group *ghttp.RouterGroup) {
		group.Middleware(func(r *ghttp.Request) {
			r.SetCtx(r.GetNeverDoneCtx())
			r.Middleware.Next()
		})
		group.Bind(
			jobexec.StaskMeta,
			jobexec.StaskResult,
		)
	})

}
