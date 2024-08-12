package jobexec

import (
	"context"
	job_exec "job-exec/api/v1/jobexec"
	"job-exec/internal/app/server/model/entity"
	"job-exec/internal/app/server/service"
)

var StaskResult = Snew()

type sTaskResultController struct {
}

func Snew() *sTaskResultController {
	return &sTaskResultController{}
}

func (s *sTaskResultController) Add(ctx context.Context, req *job_exec.TaskResultAddReq) (res *job_exec.TaskResultOperateRes, err error) {
	res = new(job_exec.TaskResultOperateRes)
	err = service.TaskResult().Add(ctx, req)
	return
}

func (s *sTaskResultController) Del(ctx context.Context, req *job_exec.TaskResultDeleteReq) (res *job_exec.TaskResultOperateRes, err error) {
	res = new(job_exec.TaskResultOperateRes)
	err = service.TaskResult().Del(ctx, req)
	return
}

func (s *sTaskResultController) Update(ctx context.Context, req *job_exec.TaskResultEditReq) (res *job_exec.TaskResultOperateRes, err error) {
	res = new(job_exec.TaskResultOperateRes)
	err = service.TaskResult().Update(ctx, req)
	return
}
func (s *sTaskResultController) List(ctx context.Context, req *job_exec.TaskResultSearchReq) (res *job_exec.TaskResultSearchRes, err error) {
	var (
		total int
		list  []*entity.TaskResult
	)
	res = new(job_exec.TaskResultSearchRes)
	total, list, err = service.TaskResult().List(ctx, req)
	res.ResultList = list
	res.Total = total
	return
}
