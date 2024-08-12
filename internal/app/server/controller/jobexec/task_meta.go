package jobexec

import (
	"context"
	job_exec "job-exec/api/v1/jobexec"
	"job-exec/internal/app/server/model/entity"
	"job-exec/internal/app/server/service"
)

var StaskMeta = New()

type sTaskMetaController struct {
}

func New() *sTaskMetaController {
	return &sTaskMetaController{}
}

func (s *sTaskMetaController) Add(ctx context.Context, req *job_exec.TaskMetaAddReq) (res *job_exec.TaskMetaOperateRes, err error) {
	res = new(job_exec.TaskMetaOperateRes)
	err = service.TaskMeta().Add(ctx, req)
	return
}

func (s *sTaskMetaController) Del(ctx context.Context, req *job_exec.TaskMetaDeleteReq) (res *job_exec.TaskMetaOperateRes, err error) {
	res = new(job_exec.TaskMetaOperateRes)
	err = service.TaskMeta().Del(ctx, req)
	return
}

func (s *sTaskMetaController) Update(ctx context.Context, req *job_exec.TaskMetaEditReq) (res *job_exec.TaskMetaOperateRes, err error) {
	res = new(job_exec.TaskMetaOperateRes)
	err = service.TaskMeta().Update(ctx, req)
	return
}
func (s *sTaskMetaController) List(ctx context.Context, req *job_exec.TaskMetaSearchReq) (res *job_exec.TaskMetaSearchRes, err error) {
	var (
		total int
		list  []*entity.TaskMeta
	)
	res = new(job_exec.TaskMetaSearchRes)
	total, list, err = service.TaskMeta().List(ctx, req)
	res.TaskMetaList = list
	res.Total = total
	return
}
