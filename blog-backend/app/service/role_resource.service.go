package service

import (
	"blog-backend/app/api/request"
	"blog-backend/app/dao"
	"context"
	"github.com/gogf/gf/frame/g"
)

type roleResourceService struct{}

var (
	RoleResourceService = new(roleResourceService)
)

func (s *roleResourceService) GrantResourceToRole(ctx context.Context, req *request.GrantResourceToRoleReq) (err error) {
	_, err = dao.RoleResources.Ctx(ctx).Data(g.Map{"role_id": req.RoleId, "resource_id": req.ResourceId}).Insert()
	return
}
