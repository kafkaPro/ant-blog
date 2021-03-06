// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// CasbinRuleDao is the manager for logic model data accessing and custom defined data operations functions management.
type CasbinRuleDao struct {
	Table   string            // Table is the underlying table name of the DAO.
	Group   string            // Group is the database configuration group name of current DAO.
	Columns CasbinRuleColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// CasbinRuleColumns defines and stores column names for table casbin_rule.
type CasbinRuleColumns struct {
	Ptype string //
	V0    string //
	V1    string //
	V2    string //
	V3    string //
	V4    string //
	V5    string //
}

//  casbinRuleColumns holds the columns for table casbin_rule.
var casbinRuleColumns = CasbinRuleColumns{
	Ptype: "ptype",
	V0:    "v0",
	V1:    "v1",
	V2:    "v2",
	V3:    "v3",
	V4:    "v4",
	V5:    "v5",
}

// NewCasbinRuleDao creates and returns a new DAO object for table data access.
func NewCasbinRuleDao() *CasbinRuleDao {
	return &CasbinRuleDao{
		Group:   "default",
		Table:   "casbin_rule",
		Columns: casbinRuleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CasbinRuleDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CasbinRuleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CasbinRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
