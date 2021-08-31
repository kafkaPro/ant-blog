package service

import (
	"blog-backend/app/api/request"
	"blog-backend/app/dao"
	"blog-backend/app/model"
	"context"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type authoritiesService struct{}

var (
	AuthoritiesService = new(authoritiesService)
)

// CreateAuthority
// 创建一个橘色
func (s *authoritiesService) CreateAuthority(ctx context.Context, req *request.CreateAuthorityReq) (
	authority *model.Authorities, err error) {
	authority = &model.Authorities{
		AuthorityId:   req.AuthorityId,
		AuthorityName: req.AuthorityName,
		ParentId:      req.ParentId,
	}

	// 检查要创建的角色id是否存在
	if s.CheckIfAuthorityExist(ctx, req.AuthorityId) {
		return nil, gerror.New("角色id已经存在，无法创建")
	}

	// 插入记录
	_, err = dao.Authorities.Ctx(ctx).OmitEmpty().Insert(authority)
	if err != nil {
		return nil, gerror.New("插入角色失败")
	}
	return
}

// CheckIfAuthorityExist
// 检查一个角色是否存在
func (s *authoritiesService) CheckIfAuthorityExist(ctx context.Context, id string) bool {
	c, err := dao.Authorities.Ctx(ctx).Count(dao.Authorities.Columns.AuthorityId, id)
	if err != nil || c == 0 {
		return false
	}
	return true
}

func (s *authoritiesService) CopyAuthority(ctx context.Context, req *request.AuthorityCopyReq) (
	authority *model.Authorities, err error) {
	if s.CheckIfAuthorityExist(ctx, req.Authority.AuthorityId) {
		return nil, gerror.New("角色已经存在，无法继续创建")
	}

	// 将复制的角色信息插入到数据库中
	_, err = dao.Authorities.Ctx(ctx).Insert(g.Map{
		dao.Authorities.Columns.AuthorityId:   req.Authority.AuthorityId,
		dao.Authorities.Columns.AuthorityName: req.Authority.AuthorityName,
		dao.Authorities.Columns.ParentId:      req.Authority.ParentId,
	})

	if err != nil {
		return nil, gerror.New("复制角色信息失败")
	}

	// 复制原来角色的api权限信息
	err = s.CopyCasbins(ctx, req)
	// 复制原来角色的动态菜单信息
	err = s.CopyMenus(ctx, req)

	err = dao.Authorities.Ctx(ctx).Where(dao.Authorities.Columns.AuthorityId, req.Authority.AuthorityId).
		Scan(&authority)
	return
}

// CopyCasbins
// 拷贝角色的api权限
func (s *authoritiesService) CopyCasbins(ctx context.Context, req *request.AuthorityCopyReq) (err error) {
	paths := s.GetPolicyPathByAuthorityId(ctx, req.OldAuthorityId)
	if err := s.UpdateCasbin(ctx, req.Authority.AuthorityId, paths); err != nil {
		_ = s.DeleteAuthority(ctx, &request.DeleteAuthorityReq{AuthorityId: req.Authority.AuthorityId})
	}
	return
}

// CopyMenus
// 将原来角色所拥有的菜单复制到新的角色中去
func (s *authoritiesService) CopyMenus(ctx context.Context, req *request.AuthorityCopyReq) (err error) {
	// 根据authorityId从authority_menu表中去获取menu_id
	// 也就是当前角色对应于哪些menu
	authorityMenus, err := AuthorityMenuService.GetAuthorityMenuByAuthorityId(ctx, req.OldAuthorityId)
	for _, authorityMenu := range authorityMenus {
		// 当前的authorityMenu包含了原来的authorityId和menuId
		// 要想复制这些信息，需要将原来的menuId和现在的authorityId进行组合
		// 最后插入数据库当中
		newAuthorityId := req.Authority.AuthorityId
		menuIdToCopy := authorityMenu.MenuId
		data := model.AuthorityMenu{
			AuthorityId: newAuthorityId,
			MenuId:      menuIdToCopy,
		}
		// 执行插入
		_, err = dao.AuthorityMenu.Ctx(ctx).Insert(data)
	}
	return
}
