// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"blog-backend/app/dao/internal"
)

// casbinRuleDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type casbinRuleDao struct {
	*internal.CasbinRuleDao
}

var (
	// CasbinRule is globally public accessible object for table casbin_rule operations.
	CasbinRule = casbinRuleDao{
		internal.NewCasbinRuleDao(),
	}
)

// Fill with you ideas below.
