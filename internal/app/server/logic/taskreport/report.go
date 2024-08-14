package taskreport

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	v1 "job-exec/api/taskreport/v1"
	dao2 "job-exec/internal/app/server/dao"
	"job-exec/internal/app/server/logic/tasksync"
	"job-exec/internal/app/server/service"
	"job-exec/utility/liberr"
)

type sTaskReport struct {
}

func New() *sTaskReport {
	return &sTaskReport{}
}
func init() {
	service.RegisterTaskReport(New())
}

func (s *sTaskReport) TaskReport(ctx context.Context, req *v1.TaskReportReq) (res *v1.TaskReportRes, err error) {
	doneTaskIds := make([]int64, 0)
	//处理agent发送过来的任务,写入task_result,更新task_meta状态
	if len(req.Results) > 0 {
		err = g.Try(ctx, func(ctx context.Context) {
			for _, task := range req.Results {
				id, err := dao2.TaskResult.Ctx(ctx).Where("host", task.Host).Where("task_id", task.TaskId).Value("id")
				data := v1.TaskResult{
					TaskId: task.TaskId,
					Host:   req.AgentIp,
					Status: task.Status,
					Stdout: task.Stdout,
					StdErr: task.StdErr,
				}
				if id != nil {
					_, err = dao2.TaskResult.Ctx(ctx).Where("host", task.Host).Where("task_id", task.TaskId).Data(data).Update()
					liberr.ErrIsNil(ctx, err, "更新执行结果失败")
				}
				_, err = dao2.TaskResult.Ctx(ctx).Data(data).Insert()
				liberr.ErrIsNil(ctx, err, "插入执行结果失败")
				doneTaskIds = append(doneTaskIds, task.TaskId)
			}
		})
	}
	dao2.TaskMeta.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = tx.Model("task_meta").WhereIn("id", doneTaskIds).Data(g.Map{"done": 1}).Update()
		return err
	})
	//给agent发送带待处理的任务'
	res = &v1.TaskReportRes{}
	tasks := tasksync.TaskCache.GetTasksByIp(req.AgentIp)
	res.AssignTasks = append(res.AssignTasks, tasks...)
	return res, nil
}
