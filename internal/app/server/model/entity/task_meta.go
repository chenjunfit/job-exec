// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskMeta is the golang structure for table task_meta.
type TaskMeta struct {
	Id         int         `json:"id"         orm:"id"          description:"任务id"`
	Title      string      `json:"title"      orm:"title"       description:"标题"`
	Account    string      `json:"account"    orm:"account"     description:"脚本执行账号"`
	ExecHosts  string      `json:"execHosts"  orm:"exec_hosts"  description:"执行的机器列表"`
	Script     string      `json:"script"     orm:"script"      description:"执行的脚本"`
	ScriptArgs string      `json:"scriptArgs" orm:"script_args" description:"脚本参数"`
	Creator    string      `json:"creator"    orm:"creator"     description:"创建者"`
	Done       int         `json:"done"       orm:"done"        description:"执行是否结束：0:没结束 1:结束"`
	CreateAt   *gtime.Time `json:"createAt"   orm:"create_at"   description:""`
	UpdateAt   *gtime.Time `json:"updateAt"   orm:"update_at"   description:""`
	DeleteAt   *gtime.Time `json:"deleteAt"   orm:"delete_at"   description:""`
}
