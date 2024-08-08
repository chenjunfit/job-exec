// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TaskResult is the golang structure for table task_result.
type TaskResult struct {
	Id       int64  `json:"id"       orm:"id"        description:""`
	TaskId   int64  `json:"taskId"   orm:"task_id"   description:"任务id"`
	Host     string `json:"host"     orm:"host"      description:"执行的具体机器信息"`
	Status   string `json:"status"   orm:"status"    description:"任务执行结果"`
	Stdout   string `json:"stdout"   orm:"stdout"    description:"标准输出"`
	StdErr   string `json:"stdErr"   orm:"std_err"   description:"标准错误"`
	CreateAt string `json:"createAt" orm:"create_at" description:""`
	UpdateAt string `json:"updateAt" orm:"update_at" description:""`
	DeleteAt string `json:"deleteAt" orm:"delete_at" description:""`
}
