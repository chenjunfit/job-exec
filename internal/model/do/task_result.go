// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskResult is the golang structure of table task_result for DAO operations like Where/Data.
type TaskResult struct {
	g.Meta   `orm:"table:task_result, do:true"`
	Id       interface{} //
	TaskId   interface{} // 任务id
	Host     interface{} // 执行的具体机器信息
	Status   interface{} // 任务执行结果
	Stdout   interface{} // 标准输出
	StdErr   interface{} // 标准错误
	CreateAt *gtime.Time //
	UpdateAt *gtime.Time //
	DeleteAt *gtime.Time //
}
