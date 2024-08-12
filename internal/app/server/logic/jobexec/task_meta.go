package jobexec

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	job_exec "job-exec/api/v1/jobexec"
	"job-exec/internal/app/server/dao"
	"job-exec/internal/app/server/model/do"
	"job-exec/internal/app/server/model/entity"
	"job-exec/internal/app/server/service"
	"job-exec/internal/consts"
	"job-exec/utility/liberr"
)

type sTaskMeta struct {
}

func init() {
	service.RegisterTaskMeta(New())
}
func New() *sTaskMeta {
	return &sTaskMeta{}
}

func (s *sTaskMeta) Add(ctx context.Context, req *job_exec.TaskMetaAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.TaskMeta.Ctx(ctx).Data(do.TaskMeta{
			Title:      req.Title,
			Account:    req.Account,
			ExecHosts:  req.ExecHosts,
			Script:     req.Script,
			ScriptArgs: req.ScriptArgs,
			Creator:    req.Creator,
			Done:       req.Done,
		}).Save()
		liberr.ErrIsNil(ctx, err, "插入失败")
	})
	return
}
func (s *sTaskMeta) Update(ctx context.Context, req *job_exec.TaskMetaEditReq) (err error) {
	var id interface{}
	err = g.Try(ctx, func(ctx context.Context) {
		id, err = dao.TaskMeta.Ctx(ctx).Fields("id").Value()
		if id != nil && gconv.String(id) != req.Id {
			liberr.ErrIsNil(ctx, err, "id不存在")
		}
		_, err = dao.TaskMeta.Ctx(ctx).Data(do.TaskMeta{
			Id:         req.Id,
			Title:      req.Title,
			Account:    req.Account,
			ExecHosts:  req.ExecHosts,
			Script:     req.Script,
			ScriptArgs: req.ScriptArgs,
			Creator:    req.Creator,
			Done:       req.Done,
		}).Where("id", req.Id).Update()
		liberr.ErrIsNil(ctx, err, "插入失败")
	})
	return
}
func (s *sTaskMeta) Del(ctx context.Context, req *job_exec.TaskMetaDeleteReq) (err error) {
	var id interface{}
	err = g.Try(ctx, func(ctx context.Context) {
		id, err = dao.TaskMeta.Ctx(ctx).Fields("id").Value()
		if id != nil && gconv.String(id) != req.Id {
			liberr.ErrIsNil(ctx, err, "id不存在")
		}
		_, err = dao.TaskMeta.Ctx(ctx).Where("id", req.Id).Delete()
		liberr.ErrIsNil(ctx, err, "插入失败")
	})
	return
}
func (s *sTaskMeta) List(ctx context.Context, req *job_exec.TaskMetaSearchReq) (total int, list []*entity.TaskMeta, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.TaskMeta.Ctx(ctx)
		if req.KeyWords != "" {
			keyWords := "%" + req.KeyWords + "%"
			m = m.Where("title like ? or  creator like ? or script like ?", keyWords, keyWords, keyWords)
		}

		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		total, err = m.Count()
		err = m.FieldsEx(dao.TaskMeta.Columns().DeleteAt).Page(req.PageNum, req.PageSize).Order("id DESC").Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取任务列表失败")
	})
	return
}
