package request

import "blog-backend/app/model"

type CreateRoleReq struct {
	RoleId   uint   `p:"roleId"   v:"required#请输入角色id"`
	RoleName string `p:"roleName" v:"required#请输入角色名字"`
	ParentId uint   `p:"parentId"  v:"required#请输入父角色id"`
}

type UpdateRoleReq struct {
	RoleId   uint   `p:"roleId" v:"required#请输入角色id"`
	RoleName string `p:"roleName" v:"required|length:1,1000#请输入角色名字|角色名字长度为:min到:max位"`
	ParentId uint   `p:"parentId" v:"required#请输入角色父id"`
}

type DeleteRoleReq struct {
	RoleId uint `p:"roleId" v:"required#请输入角色id"`
}

type RoleCopyReq struct {
	Role      model.Role `json:"role"`
	OldRoleId uint       `r:"oldRoleId" json:"oldRoleId"`
}

type RoleIdParam struct {
	RoleId uint `p:"roleId"  v:"required#角色id参数"`
}

//type SetRoleBatchDataReq struct {
//	RoleId   string               `r:"roleId" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
//	DataRole []*model.Role 		  `r:"dataRoleId" json:"dataRoleId"`
//}

//type RoleEntity struct {
//	RoleId   string `r:"roleId" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
//	RoleName string `r:"roleName" v:"required|length:1,1000#请输入角色名|角色名长度为:min到:max位"`
//}
