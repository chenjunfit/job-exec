package jobexec

import (
	"github.com/gogf/gf/v2/frame/g"
	"job-exec/api/v1/common"
	"job-exec/internal/app/server/model"
	"job-exec/internal/app/server/model/entity"
)

// TaskResultSearchReq 搜索,查询请求参数
type TaskResultSearchReq struct {
	g.Meta   `path:"/taskresult/list" tags:"执行结果信息" method:"get" summary:"执行结果列表"`
	KeyWords string `p:"keyWords"`
	common.PageReq
}

type TaskResultSearchRes struct {
	g.Meta     `mime:"application/json"`
	ResultList []*entity.TaskResult `json:"resultList"`
	common.ListRes
}

// TaskResultOperateRes 添加,删除，修改执行结果接口返回体
type TaskResultOperateRes struct {
	g.Meta `mime:"application/json"`
}

// TaskResultAddReq 添加执行结果接口
type TaskResultAddReq struct {
	g.Meta `path:"/taskresult/add" tags:"执行结果信息" method:"post" summary:"添加执行结果"`
	model.TaskResultAddReq
}

// TaskResultEditReq 编辑执行结果接口
type TaskResultEditReq struct {
	g.Meta `path:"/taskresult/edit" tags:"执行结果信息" method:"post" summary:"编辑执行结果"`
	Id     string `v:"required" p:"id"`
	model.TaskResultAddReq
}

// TaskResultDeleteReq 删除执行结果接口
type TaskResultDeleteReq struct {
	g.Meta `path:"/taskresult/del" tags:"执行结果信息" method:"post" summary:"删除执行结果"`
	Id     string `v:"required" p:"id"`
}
