// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskMetaDao is the data access object for table task_meta.
type TaskMetaDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns TaskMetaColumns // columns contains all the column names of Table for convenient usage.
}

// TaskMetaColumns defines and stores column names for table task_meta.
type TaskMetaColumns struct {
	Id         string // 任务id
	Title      string // 标题
	Account    string // 脚本执行账号
	ExecHosts  string // 执行的机器列表
	Script     string // 执行的脚本
	ScriptArgs string // 脚本参数
	Creator    string // 创建者
	Done       string // 执行是否结束：0:没结束 1:结束
	CreateAt   string //
	UpdateAt   string //
	DeleteAt   string //
}

// taskMetaColumns holds the columns for table task_meta.
var taskMetaColumns = TaskMetaColumns{
	Id:         "id",
	Title:      "title",
	Account:    "account",
	ExecHosts:  "exec_hosts",
	Script:     "script",
	ScriptArgs: "script_args",
	Creator:    "creator",
	Done:       "done",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
	DeleteAt:   "delete_at",
}

// NewTaskMetaDao creates and returns a new DAO object for table data access.
func NewTaskMetaDao() *TaskMetaDao {
	return &TaskMetaDao{
		group:   "default",
		table:   "task_meta",
		columns: taskMetaColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskMetaDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskMetaDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskMetaDao) Columns() TaskMetaColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskMetaDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskMetaDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskMetaDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
