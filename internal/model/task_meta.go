package model

type TaskMetaAddReq struct {
	Title      string `json:"title"   description:"标题"`
	Account    string `json:"account"    description:"脚本执行账号"`
	ExecHosts  string `json:"execHosts"  description:"执行的机器列表"`
	Script     string `json:"script"    description:"执行的脚本"`
	ScriptArgs string `json:"scriptArgs"  description:"脚本参数"`
	Creator    string `v:"required" json:"creator"      description:"创建者"`
	Done       int    `json:"done"      description:"执行是否结束：0:没结束 1:结束"`
}
