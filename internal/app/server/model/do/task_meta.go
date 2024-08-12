// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskMeta is the golang structure of table task_meta for DAO operations like Where/Data.
type TaskMeta struct {
	g.Meta     `orm:"table:task_meta, do:true"`
	Id         interface{} // 任务id
	Title      interface{} // 标题
	Account    interface{} // 脚本执行账号
	ExecHosts  interface{} // 执行的机器列表
	Script     interface{} // 执行的脚本
	ScriptArgs interface{} // 脚本参数
	Creator    interface{} // 创建者
	Done       interface{} // 执行是否结束：0:没结束 1:结束
	CreateAt   *gtime.Time //
	UpdateAt   *gtime.Time //
	DeleteAt   *gtime.Time //
}
