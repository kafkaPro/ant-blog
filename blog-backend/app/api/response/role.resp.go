package response

import "blog-backend/app/model"

type RoleCopyResp struct {
	Role      *model.Role `json:"role"`
	OldRoleId string      `json:"oldRoleId"`
}
