package service

import (
	"blog-backend/app/api/request"
	"blog-backend/app/dao"
	"blog-backend/app/model"
	"context"
	"github.com/gogf/gf/frame/g"
)

type roleMenuService struct{}

var RoleMenuService = new(roleMenuService)

// GetRoleMenuByMenuId
// 根据menuId获取
func (s *roleMenuService) GetRoleMenuByMenuId(ctx context.Context, menuId uint) (
	roleMenus []*model.RoleMenu, err error) {
	err = dao.RoleMenu.Ctx(ctx).Where(dao.RoleMenu.Columns.MenuId, menuId).Scan(&roleMenus)
	return
}

// GetRoleMenuByRoleId
// 根据roleId来获取
func (s *roleMenuService) GetRoleMenuByRoleId(ctx context.Context, roleId uint) (
	roleMenus []*model.RoleMenu, err error) {
	err = dao.RoleMenu.Ctx(ctx).Where(dao.RoleMenu.Columns.RoleId, roleId).Scan(&roleMenus)
	return
}

func (s *roleMenuService) GrantMenuToRole(ctx context.Context, req *request.GrantMenuToRoleReq) (err error) {
	m := dao.RoleMenu.Ctx(ctx)
	for _, menu := range req.Menus {
		_, err = m.Insert(g.Map{
			dao.RoleMenu.Columns.RoleId: req.RoleId,
			dao.RoleMenu.Columns.MenuId: menu.Id,
		})
	}
	return
}
