package taskreport

import (
	"context"
	v1 "job-exec/api/taskreport/v1"
	"job-exec/internal/app/server/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedServiceServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterServiceServer(s.Server, &Controller{})
}

func (*Controller) TaskReport(ctx context.Context, req *v1.TaskReportReq) (res *v1.TaskReportRes, err error) {
	res, err = service.TaskReport().TaskReport(ctx, req)
	return
}
