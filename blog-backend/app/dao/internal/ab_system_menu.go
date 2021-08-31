// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// AbSystemMenuDao is the manager for logic model data accessing and custom defined data operations functions management.
type AbSystemMenuDao struct {
	Table   string              // Table is the underlying table name of the DAO.
	Group   string              // Group is the database configuration group name of current DAO.
	Columns AbSystemMenuColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// AbSystemMenuColumns defines and stores column names for table ab_system_menu.
type AbSystemMenuColumns struct {
	Id         string //
	Pid        string // 父id
	Title      string // 名称
	Icon       string // 菜单图标
	Href       string // 链接
	Params     string // 链接参数
	Target     string // 链接打开方式
	Sort       string // 菜单排序
	Status     string // 状态(0:禁用,1:启用)
	Remark     string //
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeleteTime string // 删除时间
}

//  abSystemMenuColumns holds the columns for table ab_system_menu.
var abSystemMenuColumns = AbSystemMenuColumns{
	Id:         "id",
	Pid:        "pid",
	Title:      "title",
	Icon:       "icon",
	Href:       "href",
	Params:     "params",
	Target:     "target",
	Sort:       "sort",
	Status:     "status",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
}

// NewAbSystemMenuDao creates and returns a new DAO object for table data access.
func NewAbSystemMenuDao() *AbSystemMenuDao {
	return &AbSystemMenuDao{
		Group:   "default",
		Table:   "ab_system_menu",
		Columns: abSystemMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AbSystemMenuDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AbSystemMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AbSystemMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}