// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// OperateLogDao is the manager for logic model data accessing and custom defined data operations functions management.
type OperateLogDao struct {
	Table   string            // Table is the underlying table name of the DAO.
	Group   string            // Group is the database configuration group name of current DAO.
	Columns OperateLogColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// OperateLogColumns defines and stores column names for table operate_log.
type OperateLogColumns struct {
	Id            string //
	Module        string // 操作的模块的名称
	OptType       string // 操作类型(0->其他, 1->新增, 2->修改, 3->删除)
	OptPerson     string // 操作员姓名
	Ip            string // ip地址
	RequestUrl    string // 请求的URL
	RequestMethod string // 请求方法名
	RequestParam  string // 请求的参数
	Status        string // 操作状态(0->正常, 1->异常)
	OptTime       string //
}

//  operateLogColumns holds the columns for table operate_log.
var operateLogColumns = OperateLogColumns{
	Id:            "id",
	Module:        "module",
	OptType:       "opt_type",
	OptPerson:     "opt_person",
	Ip:            "ip",
	RequestUrl:    "request_url",
	RequestMethod: "request_method",
	RequestParam:  "request_param",
	Status:        "status",
	OptTime:       "opt_time",
}

// NewOperateLogDao creates and returns a new DAO object for table data access.
func NewOperateLogDao() *OperateLogDao {
	return &OperateLogDao{
		Group:   "default",
		Table:   "operate_log",
		Columns: operateLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OperateLogDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OperateLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OperateLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}