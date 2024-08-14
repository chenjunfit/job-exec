package taskworker

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "job-exec/api/taskreport/v1"
)

type LocalTasksT struct {
	M       map[int64]*Task
	MetaDir string
}

var Locals *LocalTasksT

func InitLocals(metaDir string) {
	Locals = &LocalTasksT{
		M:       make(map[int64]*Task),
		MetaDir: metaDir,
	}
}

func (lt *LocalTasksT) ReportTasks() []*v1.TaskResult {
	ret := make([]*v1.TaskResult, 0, len(lt.M))
	clean := make(map[int64]struct{})
	for id, task := range lt.M {
		rt := &v1.TaskResult{
			Id:     id,
			TaskId: task.JobId,
		}
		rt.Status = task.GetStatus()
		if rt.Status == "running" || rt.Status == "killing" {
			continue
		}
		rt.Stdout = task.GetStdOut()
		rt.StdErr = task.GetStdErr()

		stdOutLen := len(rt.Stdout)
		stdErrLen := len(rt.StdErr)

		if stdOutLen > 65535 {
			start := stdOutLen - 65535
			rt.Stdout = rt.Stdout[start:]
		}

		if stdErrLen > 65535 {
			start := stdErrLen - 65535
			rt.StdErr = rt.StdErr[start:]
		}
		ret = append(ret, rt)
		//汇报完清理，successed清理，failed清理
		clean[id] = struct{}{}
		lt.Clean(clean)
	}

	return ret
}

func (lt *LocalTasksT) GetTask(id int64) (*Task, bool) {
	t, found := lt.M[id]
	return t, found
}

func (lt *LocalTasksT) SetTask(t *Task) {
	lt.M[t.Id] = t
}

func (lt *LocalTasksT) AssignTask(at *v1.TaskMetaFix) {
	local, found := lt.M[gconv.Int64(at.Id)]
	if found {
		if local.Clock == at.Clock && local.Action == at.Action {
			return
		}
	} else {
		if at.Action == "kill" {
			return
		}
		local = &Task{
			JobId:   gconv.Int64(at.Id),
			Id:      gconv.Int64(at.Id),
			Clock:   at.Clock,
			Action:  at.Action,
			MetaDir: lt.MetaDir,
			Scripts: at.Script,
			Args:    at.ScriptArgs,
			Account: at.Account,
		}
		lt.SetTask(local)
		if local.DoneBefore() {
			local.LoadResult()
			return
		}
	}

	if local.Action == "kill" {
		local.SetStatus("killing")
		local.Kill()
	} else if local.Action == "start" {
		local.SetStatus("running")
		local.Start()
	} else {
		glog.Error(context.TODO(), "unkown actions %s taskid %d", local.Action, at.Id)
	}
}
func (lt *LocalTasksT) Clean(assigined map[int64]struct{}) {
	del := make(map[int64]struct{})
	for id := range lt.M {
		if _, found := assigined[id]; found {
			del[id] = struct{}{}
		}
	}
	for id := range del {
		if lt.M[id].GetStatus() == "running" {
			continue
		}
		lt.M[id].RestBuff()
		cmd := lt.M[id].Cmd
		delete(lt.M, id)
		if cmd != nil && cmd.Process != nil {
			cmd.Process.Release()
		}
	}

}
