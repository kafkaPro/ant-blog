package service

import (
	"blog-backend/app/api/request"
	"blog-backend/app/dao"
	"blog-backend/app/model"
	"blog-backend/app/model/database"
	"blog-backend/app/model/system"
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type menuService struct{}

var (
	MenuService = new(menuService)
)

// GetMenuTreeMap 根据角色id获取角色menu树图
func (s *menuService) GetMenuTreeMap(ctx context.Context, roleId uint) (
	menuMap map[uint][]*system.RoleMenuNode, err error) {
	roleMenus := ([]*system.RoleMenuNode)(nil)
	menuMap = make(map[uint][]*system.RoleMenuNode)

	err = g.
		DB(database.Db).
		Model("menus a").
		Safe().
		RightJoin("role_menu b", "a.id=b.menu_id").
		Fields("a.*,b.*").
		Order(dao.Menus.Columns.ParentId).
		Where("role_id", roleId).
		// 这里添加一个条件，去掉空值，where不起作用
		WhereGT("id", 0).
		Scan(&roleMenus)

	for _, roleMenuNode := range roleMenus {
		menuMap[roleMenuNode.ParentId] = append(menuMap[roleMenuNode.ParentId], roleMenuNode)
	}
	return menuMap, err
}

// GetMenuChildrenList 依次构建menu的子菜单
func (s *menuService) GetMenuChildrenList(ctx context.Context, parentMenu *system.RoleMenuNode,
	treeMap map[uint][]*system.RoleMenuNode) (err error) {
	parentMenu.Children = treeMap[parentMenu.MenuId]
	for i := 0; i < len(parentMenu.Children); i++ {
		err = s.GetMenuChildrenList(ctx, parentMenu.Children[i], treeMap)
	}
	return err
}

// GetMenuTree
// 获取对应的roleId的menu树状结构
func (s *menuService) GetMenuTree(ctx context.Context, roleId uint) (menus []*system.RoleMenuNode, err error) {
	menuMap, err := s.GetMenuTreeMap(ctx, roleId)
	// 获取该用户的根节点id，所有根节点的父节点是0
	menus = menuMap[0]
	for i := 0; i < len(menus); i++ {
		err = s.GetMenuChildrenList(ctx, menus[i], menuMap)
	}
	return menus, err
}

// 注意区分 RoleMenuNode 和 BaseMenu 之间的区别，前者包含了 role 的信息，也就是每一个角色role能够获取的菜单列表
// 而 BaseMenu 是所有列表之间的层级关系，没有包含任何的身份信息

func (s *menuService) GetBaseMenuTreeMap(ctx context.Context) (menuMap map[uint][]*system.BaseMenu, err error) {
	allMenuMap := ([]*system.BaseMenu)(nil)
	menuMap = make(map[uint][]*system.BaseMenu)
	// 直接从menu表中获取
	err = dao.Menus.Ctx(ctx).WhereGTE(dao.Menus.Columns.Id, 0).Scan(&allMenuMap)
	if err != nil {
		return
	}
	// 建立层级关系
	for i := 0; i < len(allMenuMap); i++ {
		menuMap[allMenuMap[i].ParentId] = append(menuMap[allMenuMap[i].ParentId], allMenuMap[i])
	}
	return
}

func (s *menuService) GetBaseMenuChildrenList(ctx context.Context, parentNode *system.BaseMenu,
	treeMap map[uint][]*system.BaseMenu) (err error) {
	parentNode.Children = treeMap[parentNode.Id]
	for i := 0; i < len(parentNode.Children); i++ {
		err = s.GetBaseMenuChildrenList(ctx, parentNode.Children[i], treeMap)
	}
	return
}

func (s *menuService) GetBaseMenuTree(ctx context.Context) (menus []*system.BaseMenu, err error) {
	treeMap, err := s.GetBaseMenuTreeMap(ctx)
	if err != nil {
		return
	}
	// 获取根节点
	menus = treeMap[0]
	for i := 0; i < len(menus); i++ {
		err = s.GetBaseMenuChildrenList(ctx, menus[i], treeMap)
	}
	return
}

func (s *menuService) GrantMenuToRole(ctx context.Context, req *request.GrantMenuToRoleReq) (err error) {
	return RoleMenuService.GrantMenuToRole(ctx, req)
}

// GetRoleOwnedMenu
// 获取该角色所用有的menu信息
func (s *menuService) GetRoleOwnedMenu(ctx context.Context, req *request.RoleIdParam) (
	roleMenus []*system.RoleMenuNode, err error) {
	roleMenus, err = s.GetMenuTree(ctx, req.RoleId)
	return
}

func (s *menuService) CreateBaseMenu(ctx context.Context, req *request.CreateBaseMenuReq) (err error) {
	if s.CheckIfMenuExistByName(ctx, req.Name) {
		return gerror.New("menu已经存在，无法添加")
	}

	data := model.Menus{
		MenuLevel: 0,
		ParentId:  req.ParentId,
		Path:      req.Path,
		Name:      req.Name,
		Hidden:    gconv.Int(req.Hidden),
		Component: req.Component,
		Sort:      req.Sort,
		Title:     req.Meta.Title,
		Icon:      req.Meta.Icon,
	}

	_, err = dao.Menus.Ctx(ctx).Insert(data)

	menu := model.Menus{}
	err = dao.Menus.Ctx(ctx).Scan(&menu, dao.Menus.Columns.Name, req.Name)
	if err != nil {
		return
	}
	// 添加菜单参数
	if len(req.Parameters) != 0 {
		params := g.List{}
		for _, v := range req.Parameters {
			params = append(params, g.Map{"base_menu_id": int(menu.Id), "key": v.Key, "type": v.Type})
		}
		_, err = dao.Parameters.Ctx(ctx).Data(params).Insert()
	}

	return
}

func (s *menuService) DeleteBaseMenu(ctx context.Context, req *request.DeleteBaseMenuReq) (err error) {
	if !s.CheckIfMenuExistById(ctx, req.MenuId) {
		return gerror.New("不存在该menu，无法删除")
	}

	if s.CheckIfMenuHasChildrenMenu(ctx, req.MenuId) {
		return gerror.New("该menu存在子menu，无法删除")
	}

	err = g.DB(database.Db).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 删除 menu 表
		_, e := tx.Model(dao.Menus.Table).Delete(dao.Menus.Columns.Id, req.MenuId)
		if e != nil {
			return e
		}

		// 删除 parameters表
		_, e = tx.Model(dao.Parameters.Table).Delete(dao.Parameters.Columns.BaseMenuId, req.MenuId)
		if e != nil {
			return e
		}

		// 删除 role_menu 表
		_, e = tx.Model(dao.RoleMenu.Table).Delete(dao.RoleMenu.Columns.MenuId, req.MenuId)
		if e != nil {
			return e
		}

		return nil
	})

	return
}

func (s *menuService) UpdateBaseMenu(ctx context.Context, req *request.UpdateBaseMenuReq) (err error) {
	condition := g.Map{dao.Menus.Columns.Id: req.Id}
	data := model.Menus{
		Id:        req.Id,
		ParentId:  req.ParentId,
		Name:      req.Name,
		Path:      req.Path,
		Component: req.Component,
		Hidden:    gconv.Int(req.Hidden),
		Sort:      req.Sort,
	}

	// 可选数据
	if req.Meta.DefaultMenu {
		data.DefaultMenu = 1
	}
	if req.Meta.Icon != "" {
		data.Icon = req.Meta.Icon
	}
	if req.Meta.Title != "" {
		data.Title = req.Meta.Title
	}
	if req.Meta.KeepAlive {
		data.KeepAlive = 1
	}

	_, err = dao.Menus.Ctx(ctx).Where(condition).OmitEmpty().Update()
	if err != nil {
		return
	}

	// 可选数据
	params := g.List{}
	if req.Parameters != nil {
		for _, p := range req.Parameters {
			params = append(params, g.Map{"base_menu_id": req.Id, "value": p.Value, "type": p.Type})
		}
		_, err = dao.Parameters.Ctx(ctx).Insert(params)
	}
	return
}

func (s *menuService) GetBaseMenuById(ctx context.Context, req *request.MenuIdParam) (
	menu *system.BaseMenu, err error) {
	if !s.CheckIfMenuExistById(ctx, req.MenuId) {
		return nil, gerror.New("该menu不存在")
	}

	err = dao.Menus.Ctx(ctx).Scan(&menu, dao.Menus.Columns.Id, req.MenuId)
	return
}

func (s *menuService) CheckIfMenuExistByName(ctx context.Context, menuName string) bool {
	if menuName == "" {
		return false
	}

	c, err := dao.Menus.Ctx(ctx).Where(dao.Menus.Columns.Name, menuName).Count()
	if err != nil || c == 0 {
		return false
	}
	return true
}

func (s *menuService) CheckIfMenuExistById(ctx context.Context, menuId uint) bool {
	c, err := dao.Menus.Ctx(ctx).Count(dao.Menus.Columns.Id, menuId)
	if err != nil || c == 0 {
		return false
	}
	return true
}

func (s *menuService) CheckIfMenuHasChildrenMenu(ctx context.Context, menuId uint) bool {
	c, err := dao.Menus.Ctx(ctx).Count(dao.Menus.Columns.ParentId, menuId)
	if err != nil || c == 0 {
		return false
	}
	return true
}
