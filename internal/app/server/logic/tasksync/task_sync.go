package tasksync

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "job-exec/api/taskreport/v1"
	"job-exec/internal/app/server/dao"
	"job-exec/utility/liberr"
	"sync"
	"time"
)

//1.从mysql读取所有done=0的task
//2.map[ip][]taskMeta
//3.定时每5秒去读取一次

func init() {
	TaskCache = &TaskSync{
		TasksMap: make(map[string][]*v1.TaskMetaFix),
	}
}

var TaskCache *TaskSync

type TaskSync struct {
	sync.Mutex
	TasksMap map[string][]*v1.TaskMetaFix
}

func (t *TaskSync) GetTasksByIp(ip string) []*v1.TaskMetaFix {
	t.Lock()
	defer t.Unlock()
	res, ok := t.TasksMap[ip]
	if !ok {
		res = make([]*v1.TaskMetaFix, 0)
	}
	return res
}
func (t *TaskSync) SyncManager(ctx context.Context) error {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	t.DoSyncTask(ctx)
	for {
		select {
		case <-ticker.C:
			t.DoSyncTask(ctx)
		case <-ctx.Done():
			return gerror.New("Context Done")
		}

	}
}

func (t *TaskSync) DoSyncTask(ctx context.Context) (err error) {
	//获取未完成的任务
	taskMetas := make([]*v1.TaskMetaFix, 0)
	m := make(map[string][]*v1.TaskMetaFix)
	err = dao.TaskMeta.Ctx(ctx).WhereNot("done", 1).Scan(&taskMetas)
	liberr.ErrIsNil(ctx, err, "获取任务列表失败")
	for _, taskMeta := range taskMetas {
		execHosts := taskMeta.ExecHosts
		if len(execHosts) == 0 {
			continue
		}
		for _, host := range execHosts {
			hostTaskMetas, ok := m[host]
			if !ok {
				hostTaskMetas = make([]*v1.TaskMetaFix, 0)
			}
			task := &v1.TaskMetaFix{
				Id:         taskMeta.Id,
				Account:    taskMeta.Account,
				Script:     taskMeta.Script,
				ScriptArgs: taskMeta.ScriptArgs,
				Done:       taskMeta.Done,
			}
			task.Action = "start"
			hostTaskMetas = append(hostTaskMetas, task)
			m[host] = hostTaskMetas
		}

	}
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	t.TasksMap = m
	return
}
