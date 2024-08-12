// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskResultDao is the data access object for table task_result.
type TaskResultDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns TaskResultColumns // columns contains all the column names of Table for convenient usage.
}

// TaskResultColumns defines and stores column names for table task_result.
type TaskResultColumns struct {
	Id       string //
	TaskId   string // 任务id
	Host     string // 执行的具体机器信息
	Status   string // 任务执行结果
	Stdout   string // 标准输出
	StdErr   string // 标准错误
	CreateAt string //
	UpdateAt string //
	DeleteAt string //
}

// taskResultColumns holds the columns for table task_result.
var taskResultColumns = TaskResultColumns{
	Id:       "id",
	TaskId:   "task_id",
	Host:     "host",
	Status:   "status",
	Stdout:   "stdout",
	StdErr:   "std_err",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	DeleteAt: "delete_at",
}

// NewTaskResultDao creates and returns a new DAO object for table data access.
func NewTaskResultDao() *TaskResultDao {
	return &TaskResultDao{
		group:   "default",
		table:   "task_result",
		columns: taskResultColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskResultDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskResultDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskResultDao) Columns() TaskResultColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskResultDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskResultDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskResultDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
