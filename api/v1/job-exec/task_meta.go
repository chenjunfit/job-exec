package job_exec

import (
	"github.com/gogf/gf/v2/frame/g"
	"job-exec/api/v1/common"
	"job-exec/internal/model"
	"job-exec/internal/model/entity"
)

// TaskMetaSearchReq 搜索,查询请求参数
type TaskMetaSearchReq struct {
	g.Meta   `path:"/taskmeta/list" tags:"任务信息" method:"get" summary:"任务列表"`
	KeyWords string `p:"keyWords"`
	common.PageReq
}

type TaskMetaSearchRes struct {
	g.Meta   `mime:"application/json"`
	TaskList []*entity.TaskMeta `json:"taskList"`
	common.ListRes
}

// TaskMetaOperateRes 添加,删除，修改任务接口返回体
type TaskMetaOperateRes struct {
	g.Meta `mime:"application/json"`
}

// TaskMetaAddReq 添加任务接口
type TaskMetaAddReq struct {
	g.Meta `path:"/taskmeta/add" tags:"任务信息" method:"post" summary:"添加任务"`
	model.TaskMetaAddReq
}

// TaskEditReq 编辑任务接口
type TaskEditReq struct {
	g.Meta `path:"/taskmeta/edit" tags:"任务信息" method:"post" summary:"编辑任务"`
	Id     string `v:"required" p:"id"`
	model.TaskMetaAddReq
}

// TaskDeleteReq 删除任务接口
type TaskDeleteReq struct {
	g.Meta `path:"/taskmeta/del" tags:"任务信息" method:"post" summary:"删除任务"`
	Id     string `v:"required" p:"id"`
}
