package request

import "blog-backend/app/model/system"

type GrantMenuToRoleReq struct {
	RoleId uint               `r:"roleId"`
	Menus  []*system.BaseMenu `r:"menus"`
}
