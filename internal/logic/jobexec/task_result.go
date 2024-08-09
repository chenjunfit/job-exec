package jobexec

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	job_exec "job-exec/api/v1/jobexec"
	"job-exec/internal/consts"
	"job-exec/internal/dao"
	"job-exec/internal/model/do"
	"job-exec/internal/model/entity"
	"job-exec/internal/service"
	"job-exec/utility/liberr"
)

type sTaskResult struct {
}

func init() {
	service.RegisterTaskResult(Snew())
}
func Snew() *sTaskResult {
	return &sTaskResult{}
}

func (s *sTaskResult) Add(ctx context.Context, req *job_exec.TaskResultAddReq) (err error) {
	var count int
	err = g.Try(ctx, func(ctx context.Context) {
		count, err = dao.TaskMeta.Ctx(ctx).Where("id", req.TaskId).Count()
		if count == 0 {
			liberr.ErrIsNil(ctx, err, "任务id不存在")
		}
		_, err = dao.TaskResult.Ctx(ctx).Data(do.TaskResult{
			TaskId: req.TaskId,
			Host:   req.Host,
			Status: req.Status,
			Stdout: req.Stdout,
			StdErr: req.StdErr,
		}).Save()
		liberr.ErrIsNil(ctx, err, "插入失败")
	})
	return
}
func (s *sTaskResult) Update(ctx context.Context, req *job_exec.TaskResultEditReq) (err error) {
	var (
		count int
		id    interface{}
	)

	err = g.Try(ctx, func(ctx context.Context) {
		id, err = dao.TaskResult.Ctx(ctx).Fields("id").Value()
		if id != nil && gconv.String(id) != req.Id {
			liberr.ErrIsNil(ctx, err, "id不存在")
		}
		count, err = dao.TaskMeta.Ctx(ctx).Where("id", req.TaskId).Count()
		if count == 0 {
			liberr.ErrIsNil(ctx, err, "任务id不存在")
		}
		_, err = dao.TaskResult.Ctx(ctx).Data(do.TaskResult{
			TaskId: req.TaskId,
			Host:   req.Host,
			Status: req.Status,
			Stdout: req.Stdout,
			StdErr: req.StdErr,
			Id:     req.Id,
		}).Where("id", req.Id).Update()
		liberr.ErrIsNil(ctx, err, "插入失败")
	})
	return
}
func (s *sTaskResult) Del(ctx context.Context, req *job_exec.TaskResultDeleteReq) (err error) {
	var id interface{}
	err = g.Try(ctx, func(ctx context.Context) {
		id, err = dao.TaskResult.Ctx(ctx).Fields("id").Value()
		if id != nil && gconv.String(id) != req.Id {
			liberr.ErrIsNil(ctx, err, "id不存在")
		}
		_, err = dao.TaskResult.Ctx(ctx).Where("id", req.Id).Delete()
		liberr.ErrIsNil(ctx, err, "插入失败")
	})
	return
}
func (s *sTaskResult) List(ctx context.Context, req *job_exec.TaskResultSearchReq) (total int, list []*entity.TaskResult, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.TaskResult.Ctx(ctx)
		if req.KeyWords != "" {
			keyWords := "%" + req.KeyWords + "%"
			m = m.Where("host like ? or  status like ? or taskId like ?", keyWords, keyWords, keyWords)
		}

		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		total, err = m.Count()
		err = m.FieldsEx(dao.TaskResult.Columns().DeleteAt).Page(req.PageNum, req.PageSize).Order("id DESC").Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取任务执行结果列表失败")
	})
	return
}
