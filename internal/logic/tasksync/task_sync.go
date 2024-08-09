package tasksync

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"job-exec/internal/dao"
	"job-exec/internal/model"
	"job-exec/utility/liberr"
	"sync"
	"time"
)

//1.从mysql读取所有done=0的task
//2.map[ip][]taskMeta
//3.定时每5秒去读取一次

func init() {
	TaskCache = TaskSync{
		tasksMap: make(map[string][]*model.Task),
	}
}

var TaskCache TaskSync

type TaskSync struct {
	sync.Mutex
	tasksMap map[string][]*model.Task
}

func (t *TaskSync) SyncManager(ctx context.Context) error {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			t.doSyncTask(ctx)
			for ip, task := range t.tasksMap {
				g.Dump(ip, task)
			}
		case <-ctx.Done():
			return gerror.New("Context Done")
		}

	}
}

func (t *TaskSync) doSyncTask(ctx context.Context) (err error) {
	//获取未完成的任务
	taskMetas := make([]*model.TaskMetaAddReq, 0)
	m := make(map[string][]*model.Task)
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
				hostTaskMetas = make([]*model.Task, 0)
			}
			task := &model.Task{
				Id:         taskMeta.Id,
				Title:      taskMeta.Title,
				Account:    taskMeta.Account,
				Script:     taskMeta.Script,
				ScriptArgs: taskMeta.ScriptArgs,
				Creator:    taskMeta.Creator,
				Done:       taskMeta.Done,
			}
			hostTaskMetas = append(hostTaskMetas, task)
			m[host] = hostTaskMetas
		}

	}
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	t.tasksMap = m
	return
}
