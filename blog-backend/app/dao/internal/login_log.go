// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// LoginLogDao is the manager for logic model data accessing and custom defined data operations functions management.
type LoginLogDao struct {
	Table   string          // Table is the underlying table name of the DAO.
	Group   string          // Group is the database configuration group name of current DAO.
	Columns LoginLogColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// LoginLogColumns defines and stores column names for table login_log.
type LoginLogColumns struct {
	Id            string //
	UserName      string // 登录用户名
	Ip            string // 登录ip地址
	Browser       string // 登录的浏览器
	Os            string // 登录的操作系统类型，默认Windows
	Status        string // 登录结果(0->失败,1->成功)
	Msg           string // 登录结果信息(例如:登录成功)
	LoginLocation string //
	LoginTime     string // 登录时间
}

//  loginLogColumns holds the columns for table login_log.
var loginLogColumns = LoginLogColumns{
	Id:            "id",
	UserName:      "user_name",
	Ip:            "ip",
	Browser:       "browser",
	Os:            "os",
	Status:        "status",
	Msg:           "msg",
	LoginLocation: "login_location",
	LoginTime:     "login_time",
}

// NewLoginLogDao creates and returns a new DAO object for table data access.
func NewLoginLogDao() *LoginLogDao {
	return &LoginLogDao{
		Group:   "default",
		Table:   "login_log",
		Columns: loginLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LoginLogDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LoginLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
