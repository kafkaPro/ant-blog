// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"blog-backend/app/dao/internal"
)

// breakpointChucksDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type breakpointChucksDao struct {
	*internal.BreakpointChucksDao
}

var (
	// BreakpointChucks is globally public accessible object for table breakpoint_chucks operations.
	BreakpointChucks = breakpointChucksDao{
		internal.NewBreakpointChucksDao(),
	}
)

// Fill with you ideas below.