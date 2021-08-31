// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// AbSystemNodeDao is the manager for logic model data accessing and custom defined data operations functions management.
type AbSystemNodeDao struct {
	Table   string              // Table is the underlying table name of the DAO.
	Group   string              // Group is the database configuration group name of current DAO.
	Columns AbSystemNodeColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// AbSystemNodeColumns defines and stores column names for table ab_system_node.
type AbSystemNodeColumns struct {
	Id         string //
	Node       string // 节点代码
	Title      string // 节点标题
	Type       string // 节点类型（1：控制器，2：节点）
	IsAuth     string // 是否启动RBAC权限控制
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

//  abSystemNodeColumns holds the columns for table ab_system_node.
var abSystemNodeColumns = AbSystemNodeColumns{
	Id:         "id",
	Node:       "node",
	Title:      "title",
	Type:       "type",
	IsAuth:     "is_auth",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewAbSystemNodeDao creates and returns a new DAO object for table data access.
func NewAbSystemNodeDao() *AbSystemNodeDao {
	return &AbSystemNodeDao{
		Group:   "default",
		Table:   "ab_system_node",
		Columns: abSystemNodeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AbSystemNodeDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AbSystemNodeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AbSystemNodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
