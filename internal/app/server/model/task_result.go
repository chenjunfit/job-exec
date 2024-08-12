package model

type TaskResultAddReq struct {
	TaskId int64  `json:"taskId"   orm:"task_id"   description:"任务id"`
	Host   string `json:"host"     orm:"host"      description:"执行的具体机器信息"`
	Status string `json:"status"   orm:"status"    description:"任务执行结果"`
	Stdout string `json:"stdout"   orm:"stdout"    description:"标准输出"`
	StdErr string `json:"stdErr"   orm:"std_err"   description:"标准错误"`
}
