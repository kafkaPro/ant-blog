package middleware

import (
	"blog-backend/app/model/common"
	"blog-backend/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func AdminCasbin(r *ghttp.Request) {
	// 获取请求的URI
	obj := r.Request.URL.RequestURI()
	// 请求方法
	act := r.Request.Method
	// 哪个角色请求
	sub := r.GetParam("roleId")

	// casbin验证
	e := service.Casbin()
	success, _ := e.Enforce(sub, obj, act)
	// 如果是开发环境，跳过casbin验证
	if g.Cfg("system").GetString("system.Env") == "develop" || success {
		r.Middleware.Next()
	} else {
		common.FailWithMessage(r, "您没有当前权限")
		r.ExitAll()
	}
}
