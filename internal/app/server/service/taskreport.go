// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "job-exec/api/taskreport/v1"
)

type (
	ITaskReport interface {
		TaskReport(ctx context.Context, req *v1.TaskReportReq) (res *v1.TaskReportRes, err error)
	}
)

var (
	localTaskReport ITaskReport
)

func TaskReport() ITaskReport {
	if localTaskReport == nil {
		panic("implement not found for interface ITaskReport, forgot register?")
	}
	return localTaskReport
}

func RegisterTaskReport(i ITaskReport) {
	localTaskReport = i
}
