// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// CustomersDao is the manager for logic model data accessing and custom defined data operations functions management.
type CustomersDao struct {
	Table   string           // Table is the underlying table name of the DAO.
	Group   string           // Group is the database configuration group name of current DAO.
	Columns CustomersColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// CustomersColumns defines and stores column names for table customers.
type CustomersColumns struct {
	Id                 string // 自增ID
	CreateAt           string // 创建时间
	UpdateAt           string // 更新时间
	DeleteAt           string // 删除时间
	CustomerName       string // 客户名
	CustomerPhoneData  string // 客户电话
	SysUserId          string // 负责员工id
	SysUserAuthorityId string // 负责员工角色
}

//  customersColumns holds the columns for table customers.
var customersColumns = CustomersColumns{
	Id:                 "id",
	CreateAt:           "create_at",
	UpdateAt:           "update_at",
	DeleteAt:           "delete_at",
	CustomerName:       "customer_name",
	CustomerPhoneData:  "customer_phone_data",
	SysUserId:          "sys_user_id",
	SysUserAuthorityId: "sys_user_authority_id",
}

// NewCustomersDao creates and returns a new DAO object for table data access.
func NewCustomersDao() *CustomersDao {
	return &CustomersDao{
		Group:   "default",
		Table:   "customers",
		Columns: customersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CustomersDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CustomersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CustomersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}