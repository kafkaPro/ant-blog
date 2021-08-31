package service

import (
	"blog-backend/app/api/request"
	"blog-backend/app/dao"
	"blog-backend/app/model"
	"blog-backend/app/model/system"
	"blog-backend/library/utils"
	"context"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type adminService struct{}

var (
	AdminService = new(adminService)
)

func (s *adminService) ChangePassword(ctx context.Context, req *request.ChangePasswordReq) (err error) {
	// 首先检查req对应的用户是否存在，不存在的用户无法进行密码的修改
	c, err := dao.Admins.Ctx(ctx).Count(dao.Admins.Columns.Username, req.Username)
	if err != nil || c == 0 {
		return gerror.New("该用户不存在，无法修改密码")
	}

	// 该用户存在，获取用户信息
	var admin *model.Admins
	err = dao.Admins.Ctx(ctx).Where(dao.Admins.Columns.Username, req.Username).Scan(&admin)
	if err != nil {
		return gerror.New("获取admin信息失败")
	}

	// 尝试匹配密码
	if utils.CompareHashAndPassword(admin.Password, req.Password) {
		newPwd, err := utils.EncryptPassword(req.NewPassword)
		if err != nil {
			return gerror.New("加密密码失败")
		}
		// 将新的信息更新到数据库中
		admin.Password = newPwd
		_, err = dao.Admins.Ctx(ctx).Where(dao.Admins.Columns.Username, req.Username).Update(admin)
		if err != nil {
			return gerror.New("更新密码失败")
		}
		return nil
	}
	return gerror.New("输入的旧密码不正确")
}

// GetAdminList
// 根据分页参数获取admin列表，同时获取每一个admin对应的权限信息
func (s *adminService) GetAdminList(ctx context.Context, req *request.PageReq) (list interface{}, total int, err error) {
	adminList := ([]*system.AdminHasOneAuthority)(nil)
	adminDao := dao.Admins.Ctx(ctx).Safe()
	authorityDao := dao.Authorities.Ctx(ctx).Safe()
	err = adminDao.Page(req.PageNo, req.PageSize).ScanList(&adminList, "Admin")
	if err != nil {
		return nil, 0, gerror.New("获取admin列表失败")
	}

	// 获取了admin列表之后，需要将权限设置进去
	for _, admin := range adminList {
		err = authorityDao.Where(dao.Authorities.Columns.AuthorityId, admin.Admin.AuthorityId).Scan(&admin.Authority)
	}
	return adminList, total, err
}

// FindAdminByUuid
// 根据uuid来获取admin信息
func (s *adminService) FindAdminByUuid(ctx context.Context, uuid string) (admin *system.AdminHasOneAuthority, err error) {
	admin = &system.AdminHasOneAuthority{}
	adminDao := dao.Admins.Ctx(ctx).Safe()
	authorityDao := dao.Authorities.Ctx(ctx).Safe()
	err = adminDao.Where(dao.Admins.Columns.Uuid, uuid).Scan(&admin.Admin)
	if err != nil {
		return
	}
	err = authorityDao.Where(dao.Authorities.Columns.AuthorityId, admin.Admin.AuthorityId).
		Scan(&admin.Authority)
	return
}

// FindAdminById
// 根据id来寻找admin
func (s *adminService) FindAdminById(ctx context.Context, id int) (admin *system.AdminHasOneAuthority, err error) {
	admin = &system.AdminHasOneAuthority{}
	adminDao := dao.Admins.Ctx(ctx).Safe()
	authorityDao := dao.Authorities.Ctx(ctx).Safe()
	err = adminDao.Where(dao.Admins.Columns.Id, id).Scan(&admin.Admin)
	if err != nil {
		return
	}
	err = authorityDao.Where(dao.Authorities.Columns.AuthorityId, admin.Admin.AuthorityId).
		Scan(&admin.Authority)
	return
}

// SetUserAuthority
// 设置用户的权限
func (s *adminService) SetUserAuthority(ctx context.Context, req *request.SetAdminAuthorityReq) error {
	_, err := dao.Admins.Ctx(ctx).
		Where(dao.Admins.Columns.Uuid, req.Uuid).
		Update(g.Map{dao.Admins.Columns.AuthorityId: req.AuthorityId})
	if err != nil {
		return err
	}
	return nil
}

// DeleteAdmin
// 根据id删除admin信息
func (s *adminService) DeleteAdmin(ctx context.Context, req *request.DeleteAdminReq) (err error) {
	_, err = dao.Admins.Ctx(ctx).Delete(dao.Admins.Columns.Id, req.Id)
	return
}

// SetAdminInfo
// 设置admin的信息
func (s *adminService) SetAdminInfo(ctx context.Context, req *request.SetAdminInfoReq) (
	admin *system.AdminHasOneAuthority, err error) {
	_, err = dao.Admins.Ctx(ctx).Update(dao.Admins.Columns.HeaderImg, req.HeaderImg, dao.Admins.Columns.Uuid, req.Uuid)
	return
}
