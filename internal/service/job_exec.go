// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	job_exec "job-exec/api/v1/jobexec"
	"job-exec/internal/model/entity"
)

type (
	ITaskMeta interface {
		Add(ctx context.Context, req *job_exec.TaskMetaAddReq) (err error)
		Update(ctx context.Context, req *job_exec.TaskMetaEditReq) (err error)
		Del(ctx context.Context, req *job_exec.TaskMetaDeleteReq) (err error)
		List(ctx context.Context, req *job_exec.TaskMetaSearchReq) (total int, list []*entity.TaskMeta, err error)
	}
	ITaskResult interface {
		Add(ctx context.Context, req *job_exec.TaskResultAddReq) (err error)
		Update(ctx context.Context, req *job_exec.TaskResultEditReq) (err error)
		Del(ctx context.Context, req *job_exec.TaskResultDeleteReq) (err error)
		List(ctx context.Context, req *job_exec.TaskResultSearchReq) (total int, list []*entity.TaskResult, err error)
	}
)

var (
	localTaskMeta   ITaskMeta
	localTaskResult ITaskResult
)

func TaskMeta() ITaskMeta {
	if localTaskMeta == nil {
		panic("implement not found for interface ITaskMeta, forgot register?")
	}
	return localTaskMeta
}

func RegisterTaskMeta(i ITaskMeta) {
	localTaskMeta = i
}

func TaskResult() ITaskResult {
	if localTaskResult == nil {
		panic("implement not found for interface ITaskResult, forgot register?")
	}
	return localTaskResult
}

func RegisterTaskResult(i ITaskResult) {
	localTaskResult = i
}
