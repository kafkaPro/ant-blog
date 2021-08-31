package request

import "blog-backend/app/model"

type CreateAuthorityReq struct {
	AuthorityId   string `p:"authority_id" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
	AuthorityName string `p:"authority_name" v:"required|length:1,1000#请输入角色名字|角色名字长度为:min到:max位"`
	ParentId      string `p:"authority_id" v:"required|length:1,1000#请输入角色父id|角色父id长度为:min到:max位"`
}

type UpdateAuthorityReq struct {
	AuthorityId   string `p:"authority_id" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
	AuthorityName string `p:"authority_name" v:"required|length:1,1000#请输入角色名字|角色名字长度为:min到:max位"`
	ParentId      string `p:"authority_id" v:"required|length:1,1000#请输入角色父id|角色父id长度为:min到:max位"`
}

type DeleteAuthorityReq struct {
	AuthorityId string `p:"authorityId" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
}

type AuthorityCopyReq struct {
	Authority      model.Authorities `json:"authority"`
	OldAuthorityId string            `r:"oldAuthorityId" json:"oldAuthorityId"`
}

type SetDataAuthorityReq struct {
	AuthorityId   string               `r:"authorityId" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
	DataAuthority []*model.Authorities `r:"dataAuthorityId" json:"dataAuthorityId"`
}

type AuthorityEntity struct {
	AuthorityId   string `r:"authorityId" v:"required|length:1,1000#请输入角色id|角色id长度为:min到:max位"`
	AuthorityName string `r:"authorityName" v:"required|length:1,1000#请输入角色名|角色名长度为:min到:max位"`
}
