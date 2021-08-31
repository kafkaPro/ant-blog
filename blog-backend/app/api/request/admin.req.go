package request

type ChangePasswordReq struct {
	Uuid        string
	Username    string `p:"username" v:"required|length:1,30#请输入用户名称|您输入用户名称长度非法"`
	Password    string `p:"password" v:"required|length:6,30#请输入旧密码|旧密码长度为:min到:max位"`
	NewPassword string `p:"new_password" v:"required|length:6,30#请输入新密码|新密码长度为:min到:max位"`
}

type SetAdminAuthorityReq struct {
	Uuid        string `p:"uuid" v:"required|length:36,36#请输入管理员UUID|管理员UUID长度为:min到:max位"`
	AuthorityId string `p:"authority_id" v:"required|length:1, 100#请输入角色ID|角色ID长度为:min到:max位"`
}

type SetAdminInfoReq struct {
	Uuid      string
	HeaderImg string `p:"headerImg" v:"required|length:1, 100#请输入头像链接|头像链接长度为:min到:max位"`
}

type DeleteAdminReq struct {
	Id uint `p:"id" v:"required|min:1#请输入id|id必须为正数"`
}
