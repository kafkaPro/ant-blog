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
	"github.com/gogf/guuid"
	"github.com/mojocn/base64Captcha"
)

type loginService struct {
	Store base64Captcha.Store
}

var (
	LoginService = &loginService{
		Store: base64Captcha.DefaultMemStore,
	}
)

// AdminLogin
// 管理员登录成功之后，会返回一个对应角色的信息体，关系到页面的前端初始化操作
// @TODO 添加登录日志，在线用户记录等功能
func (s *loginService) AdminLogin(ctx context.Context, req *request.AdminLoginReq) (
	data *system.AdminHasOneRole, err error) {
	// 检查admin是否存在
	condition := g.Map{"username": req.Username}
	if !AdminService.CheckIfAdminExist(ctx, condition) {
		return nil, gerror.New("admin账户不存在")
	}

	// 获取admin
	admin, err := AdminService.FindAdminByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	// 检查该username对应的密码是否正确
	if !utils.CompareHashAndPassword(admin.Admin.Password, req.Password) {
		return nil, gerror.New("账号或密码不正确")
	}
	return admin, nil
}

func (s *loginService) AdminRegister(ctx context.Context, req *request.AdminRegisterReq) (err error) {
	// 检查用户是否存在
	condition := g.Map{"username": req.Username}
	if AdminService.CheckIfAdminExist(ctx, condition) {
		return gerror.New("该用户已经存在，无法进行注册")
	}

	data := model.Admins{
		Uuid:      guuid.New().String(),
		Username:  req.Username,
		Password:  req.Password,
		Nickname:  req.Nickname,
		HeaderImg: req.HeaderImg,
		RoleId:    req.RoleId,
	}

	_, err = dao.Admins.Ctx(ctx).OmitEmpty().Insert(data)
	if err != nil {
		return gerror.New("注册失败")
	}
	return nil
}

func (s *loginService) Captcha(ctx context.Context) (id string, b64Str string, err error) {
	imgH := g.Cfg("captcha").GetInt("captcha.ImgHeight")
	imgW := g.Cfg("captcha").GetInt("captcha.ImgWidth")
	Key := g.Cfg("captcha").GetInt("captcha.KeyLong")
	driver := base64Captcha.NewDriverDigit(imgH, imgW, Key, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, s.Store)
	id, b64Str, err = cp.Generate()
	return
}
