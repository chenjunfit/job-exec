package job_exec

import "github.com/gogf/gf/v2/frame/g"

// StudentSearchReq UserSearchReq 用户搜索请求参数
type StudentSearchReq struct {
	g.Meta   `path:"/student/list" tags:"学生管理" method:"get" summary:"学生列表"`
	ClassId  string `p:"classId"` //班级id
	Status   string `p:"status"`
	KeyWords string `p:"keyWords"`
	commonApi.PageReq
	commonApi.Author
}

type StudentSearchRes struct {
	g.Meta      `mime:"application/json"`
	StudentList []*entity.SysStudent `json:"studentList"`
	commonApi.ListRes
}

// StudentDetailReq 获取班级详情接口
type StudentDetailReq struct {
	g.Meta `path:"/student/detail" tags:"学生管理" method:"get" summary:"学生详情"`
	Id     string `p:"id" v:"required#学生ID不能为空"`
	commonApi.Author
}

// StudentDetailRes 学生详情返回
type StudentDetailRes struct {
	g.Meta `mime:"application/json"`
	Detail *entity.SysStudent `json:"detail"`
}

// StudentOperateRes 添加接口返回体
type StudentOperateRes struct {
	g.Meta `mime:"application/json"`
}

// StudentAddReq 添加学生接口
type StudentAddReq struct {
	g.Meta `path:"/student/add" tags:"学生管理" method:"post" summary:"添加学生"`
	model.StudentAddReq
	commonApi.Author
}

// StudentEditReq 编辑学生接口
type StudentEditReq struct {
	g.Meta `path:"/student/edit" tags:"学生管理" method:"post" summary:"编辑学生"`
	Id     string `p:"id"`
	model.StudentAddReq
	commonApi.Author
}

// StudentDeleteReq 删除班级接口
type StudentDeleteReq struct {
	g.Meta `path:"/student/del" tags:"学生管理" method:"post" summary:"删除学生"`
	Id     string `p:"id"`
	commonApi.Author
}

type StudentStatusChangeReq struct {
	g.Meta `path:"/student/status" tags:"学生管理" method:"post" summary:"学生禁用状态更换"`
	Id     string `p:"id"`
	Status string `p:"status"`
	commonApi.Author
}
