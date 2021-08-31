// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// MenusDao is the manager for logic model data accessing and custom defined data operations functions management.
type MenusDao struct {
	Table   string       // Table is the underlying table name of the DAO.
	Group   string       // Group is the database configuration group name of current DAO.
	Columns MenusColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// MenusColumns defines and stores column names for table menus.
type MenusColumns struct {
	Id          string // 自增ID
	CreateAt    string // 创建时间
	UpdateAt    string // 更新时间
	DeleteAt    string // 删除时间
	MenuLevel   string // 菜单等级(预留字段)
	ParentId    string // 父菜单ID
	Path        string // 路由path
	Name        string // 路由name
	Hidden      string // 是否在列表隐藏
	Component   string // 前端文件路径
	Title       string // 菜单名
	Icon        string // 菜单图标
	Sort        string // 排序标记
	KeepAlive   string // 是否缓存
	DefaultMenu string // 是否是基础路由(开发中)
}

//  menusColumns holds the columns for table menus.
var menusColumns = MenusColumns{
	Id:          "id",
	CreateAt:    "create_at",
	UpdateAt:    "update_at",
	DeleteAt:    "delete_at",
	MenuLevel:   "menu_level",
	ParentId:    "parent_id",
	Path:        "path",
	Name:        "name",
	Hidden:      "hidden",
	Component:   "component",
	Title:       "title",
	Icon:        "icon",
	Sort:        "sort",
	KeepAlive:   "keep_alive",
	DefaultMenu: "default_menu",
}

// NewMenusDao creates and returns a new DAO object for table data access.
func NewMenusDao() *MenusDao {
	return &MenusDao{
		Group:   "default",
		Table:   "menus",
		Columns: menusColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MenusDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MenusDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MenusDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}