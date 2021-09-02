package middleware

import (
	v1 "blog-backend/app/api/v1"
	"blog-backend/app/model/system"
	"blog-backend/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

func AdminCtx(r *ghttp.Request) {
	// 实例化 admin 的 context 变量
	adminCtx := &service.AdminContext{
		Session: r.Session,
		Data:    g.Map{},
	}

	// 注入到 request 中
	service.AdminContextService.Init(r, adminCtx)
	// 获取登录的返回结果
	adminData := &system.AdminHasOneRole{}
	respData := v1.AdminGfToken.GetTokenData(r)
	if err := gconv.Struct(respData.Get("data"), &adminData); err != nil {
		g.Log().Info("设置admin信息到Context中失败")
	}

	if adminData != nil {
		adminCtx.User = &service.ContextAdmin{
			Id:        adminData.Admin.Id,
			UserName:  adminData.Admin.Username,
			NickName:  adminData.Admin.Nickname,
			Uuid:      adminData.Admin.Uuid,
			HeaderImg: adminData.Admin.HeaderImg,
			IsAdmin:   1,
			RoleId:    adminData.Role.RoleId,
		}
	}

	// 将admin上下文添加到request中
	r.Assigns(g.Map{
		"Context": adminCtx,
	})

	// 执行下一步的请求
	r.Middleware.Next()
}
