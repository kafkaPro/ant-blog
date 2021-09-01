package service

import (
	"blog-backend/app/api/request"
	"blog-backend/app/dao"
	"blog-backend/app/model"
	"blog-backend/app/model/database"
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type roleService struct{}

var (
	RoleService = new(roleService)
)

// CreateRole
// 创建一个橘色
func (s *roleService) CreateRole(ctx context.Context, req *request.CreateRoleReq) (
	authority *model.Role, err error) {
	authority = &model.Role{
		RoleId:   req.RoleId,
		RoleName: req.RoleName,
		ParentId: req.ParentId,
	}

	// 检查要创建的角色id是否存在
	if s.CheckIfRoleExist(ctx, req.RoleId) {
		return nil, gerror.New("角色id已经存在，无法创建")
	}

	// 插入记录
	_, err = dao.Role.Ctx(ctx).OmitEmpty().Insert(authority)
	if err != nil {
		return nil, gerror.New("插入角色失败")
	}
	return
}

// CheckIfRoleExist
// 检查一个角色是否存在
func (s *roleService) CheckIfRoleExist(ctx context.Context, id uint) bool {
	c, err := dao.Role.Ctx(ctx).Count(dao.Role.Columns.RoleId, id)
	if err != nil || c == 0 {
		return false
	}
	return true
}

func (s *roleService) CopyRole(ctx context.Context, req *request.RoleCopyReq) (
	authority *model.Role, err error) {
	if s.CheckIfRoleExist(ctx, req.Role.RoleId) {
		return nil, gerror.New("角色已经存在，无法继续创建")
	}

	// 将复制的角色信息插入到数据库中
	_, err = dao.Role.Ctx(ctx).Insert(g.Map{
		dao.Role.Columns.RoleId:   req.Role.RoleId,
		dao.Role.Columns.RoleName: req.Role.RoleName,
		dao.Role.Columns.ParentId: req.Role.ParentId,
	})

	if err != nil {
		return nil, gerror.New("复制角色信息失败")
	}

	// 复制原来角色的api权限信息
	err = s.CopyCasbins(ctx, req)
	// 复制原来角色的动态菜单信息
	err = s.CopyMenus(ctx, req)

	err = dao.Role.Ctx(ctx).Where(dao.Role.Columns.RoleId, req.Role.RoleId).
		Scan(&authority)
	return
}

// DeleteRole
// 删除角色，同时删除 role 表和 role_menu 表
func (s *roleService) DeleteRole(ctx context.Context, req *request.DeleteRoleReq) error {
	// 删除之前首先检查role是否存在以及role是否含有子role，防止开启事务会产生新能损耗
	if !s.CheckIfRoleExist(ctx, req.RoleId) && s.CheckIfRoleHasChildRole(ctx, req.RoleId) {
		return gerror.New("角色不存在或角色存在子角色，无法删除")
	}

	// 检查该角色是否含有子角色，存在子角色的情况下不允许进行删除操作

	err := g.DB(database.Db).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 删除 role 表中的信息
		_, e := tx.Ctx(ctx).Model(dao.Role.Table).Delete(dao.Role.Columns.RoleId, req.RoleId)
		if e != nil {
			return e
		}

		// 删除 role_menu 表中的信息
		_, e = tx.Ctx(ctx).Model(dao.RoleMenu.Table).Delete(dao.RoleMenu.Columns.RoleId, req.RoleId)
		if e != nil {
			return e
		}
		return nil
	})

	return err
}

// CopyCasbins
// 拷贝角色的api权限
func (s *roleService) CopyCasbins(ctx context.Context, req *request.RoleCopyReq) (err error) {
	//paths := s.GetPolicyPathByRoleId(ctx, req.OldRoleId)
	paths := CasbinService.GetCasbinPolicyByRoleId(ctx, req.OldRoleId)
	if err := CasbinService.UpdateCasbin(ctx, req.Role.RoleId, paths); err != nil {
		_ = s.DeleteRole(ctx, &request.DeleteRoleReq{RoleId: req.Role.RoleId})
	}
	return
}

// CopyMenus
// 将原来角色所拥有的菜单复制到新的角色中去
func (s *roleService) CopyMenus(ctx context.Context, req *request.RoleCopyReq) (err error) {
	// 根据roleId从role_menu表中去获取menu_id
	// 也就是当前角色对应于哪些menu
	authorityMenus, err := RoleMenuService.GetRoleMenuByRoleId(ctx, req.OldRoleId)
	for _, authorityMenu := range authorityMenus {
		// 当前的roleMenu包含了原来的roleId和menuId
		// 要想复制这些信息，需要将原来的menuId和现在的roleId进行组合
		// 最后插入数据库当中
		newRoleId := req.Role.RoleId
		menuIdToCopy := authorityMenu.MenuId
		data := model.RoleMenu{
			RoleId: newRoleId,
			MenuId: menuIdToCopy,
		}
		// 执行插入
		_, err = dao.RoleMenu.Ctx(ctx).Insert(data)
	}
	return
}

// CheckIfRoleHasChildRole
// 检查给定parentId的角色是否含有子角色
func (s *roleService) CheckIfRoleHasChildRole(ctx context.Context, parentId uint) bool {
	c, err := dao.Role.Ctx(ctx).Where(dao.Role.Columns.ParentId, parentId).Count()
	if err != nil || c == 0 {
		return false
	}
	return true
}

func (s *roleService) UpdateRole(ctx context.Context, req *request.UpdateRoleReq) (err error) {
	data := model.Role{
		RoleId:   req.RoleId,
		RoleName: req.RoleName,
		ParentId: req.ParentId,
	}

	_, err = dao.Role.Ctx(ctx).FieldsEx(dao.Role.Columns.RoleId).Update(data, dao.Role.Columns.RoleId, req.RoleId)
	return
}

func (s *roleService) GetRoleListInPageMode(ctx context.Context, req *request.PageReq) (
	list []*model.Role, total int, err error) {
	err = dao.Role.Ctx(ctx).Safe().Page(req.PageNo, req.PageSize).Scan(&list)
	return list, len(list), err
}

func (s *roleService) GetRoleById(ctx context.Context, roleId uint) (role *model.Role, err error) {
	err = dao.Role.Ctx(ctx).Where(dao.Role.Columns.RoleId, roleId).Scan(&role)
	return
}

func (s *roleService) GetRoleTreeMap(ctx context.Context) (treeMap map[uint][]*model.Role, total int, err error) {
	treeMap = make(map[uint][]*model.Role)
	roleList := ([]*model.Role)(nil)

	err = dao.Role.Ctx(ctx).WhereGTE(dao.Role.Columns.RoleId, 0).Scan(&roleList)
	if err != nil {
		return
	}

	for _, v := range roleList {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, len(roleList), err
}

func (s *roleService) GrantResourceToRole(ctx context.Context, req *request.GrantResourceToRoleReq) error {
	return RoleResourceService.GrantResourceToRole(ctx, req)
}
