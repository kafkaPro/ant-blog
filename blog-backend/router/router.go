package router

import (
	"blog-backend/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.CORS)
		//group.Middleware(middleware.AdminCtx)
		//v1.AdminGfToken.Middleware(group)

		group.GET("/as", func(r *ghttp.Request) {
			_ = r.Response.WriteJson("权限验证通过")
		})

		group.Middleware(middleware.AdminCasbin)
		group.POST("/base/login", func(r *ghttp.Request) {
			_ = r.Response.WriteJson("casbin权限认证通过")
		})
	})
}
