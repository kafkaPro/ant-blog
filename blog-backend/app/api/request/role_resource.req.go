package request

type GrantResourceToRoleReq struct {
	RoleId     uint `p:"roleId"     v:"required#请输入角色id"`
	ResourceId uint `p:"resourceId" v:"required#请输入资源id"`
}
