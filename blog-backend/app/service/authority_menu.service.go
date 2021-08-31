package service

import (
	"blog-backend/app/dao"
	"blog-backend/app/model"
	"context"
)

type authorityMenuService struct{}

var AuthorityMenuService = new(authorityMenuService)

// GetAuthorityMenuByMenuId
// 根据menuId获取
func (s *authorityMenuService) GetAuthorityMenuByMenuId(ctx context.Context, menuId string) (
	authorityMenus []*model.AuthorityMenu, err error) {
	err = dao.AuthorityMenu.Ctx(ctx).Where(dao.AuthorityMenu.Columns.MenuId, menuId).Scan(&authorityMenus)
	return
}

// GetAuthorityMenuByAuthorityId
// 根据authorityId来获取
func (s *authorityMenuService) GetAuthorityMenuByAuthorityId(ctx context.Context, authorityId string) (
	authorityMenus []*model.AuthorityMenu, err error) {
	err = dao.AuthorityMenu.Ctx(ctx).Where(dao.AuthorityMenu.Columns.AuthorityId, authorityId).Scan(&authorityMenus)
	return
}
