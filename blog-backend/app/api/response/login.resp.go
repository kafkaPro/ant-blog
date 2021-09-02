package response

import (
	"blog-backend/app/model/system"
)

type AdminLoginResp struct {
	User   *system.AdminHasOneRole `json:"detail"`
	Id     uint                    `json:"id"`
	RoleId uint                    `json:"roleId"`
	Uuid   string                  `json:"uuid"`
}
