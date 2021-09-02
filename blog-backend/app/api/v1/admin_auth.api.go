package v1

import (
	"blog-backend/app/api/request"
	"blog-backend/app/model"
	c "blog-backend/app/model/common"
	"blog-backend/app/model/system"
	"blog-backend/app/service"
	"blog-backend/library/utils"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/mssola/user_agent"
	"strings"
)

type adminAuthApi struct{}

var (
	AdminAuthApi = new(adminAuthApi)

	AdminGfToken = &gtoken.GfToken{
		// 登录相关
		LoginBeforeFunc: AdminAuthApi.LoginBefore,
		LoginPath:       "/login",
		LoginAfterFunc:  AdminAuthApi.LoginAfter,
		// 登出相关
		LogoutBeforeFunc: AdminAuthApi.LogoutBefore,
		LogoutPath:       "/logout",
		LogoutAfterFunc:  AdminAuthApi.LogoutAfter,
		// 鉴权相关
		AuthExcludePaths: g.SliceStr{"/login"},
		AuthBeforeFunc:   AdminAuthApi.AuthBefore,
		AuthAfterFunc:    AdminAuthApi.AuthAfter,
		AuthFailMsg:      "该请求无权限",
		// token相关配置
		CacheMode:      int8(g.Cfg("gtoken").GetInt("gtoken.admin.CacheMode")),
		CacheKey:       g.Cfg("gtoken").GetString("gtoken.admin.CacheKey"),
		Timeout:        g.Cfg("gtoken").GetInt("gtoken.admin.Timeout"),
		MultiLogin:     g.Cfg("gtoken").GetBool("gtoken.admin.AllowMultiLogin"),
		MaxRefresh:     g.Cfg("gtoken").GetInt("gtoken.admin.MaxRefresh"),
		TokenDelimiter: g.Cfg("gtoken").GetString("gtoken.admin.TokenDelimiter"),
		EncryptKey:     g.Cfg("gtoken").GetBytes("gtoken.admin.EncryptKey"),
	}
)

func (api *adminAuthApi) LoginBefore(r *ghttp.Request) (string, interface{}) {
	ctx := r.GetCtx()
	var req *request.AdminLoginReq
	if err := r.Parse(&req); err != nil {
		c.FailWithMessage(r, "解析登录请求参数失败")
		return "", nil
	}

	clientIp := utils.GetClientIp(r)
	ua := r.Header.Get("User-Agent")
	userAgent := user_agent.New(ua)
	location := utils.GetCityByIp(clientIp)
	os := userAgent.OS()
	browser, _ := userAgent.Browser()
	if os == "" {
		os = "Windows10"
	}

	loginLogData := model.LoginLog{
		UserName:      req.Username,
		LoginLocation: location,
		LoginTime:     gtime.Now(),
		Ip:            clientIp,
		Browser:       browser,
		Os:            os,
	}

	// 先验证用户名和密码
	adminHasOneRole, err := service.LoginService.AdminLogin(ctx, req)
	if err != nil || adminHasOneRole == nil {
		c.FailWithMessage(r, "用户名或密码错误")
		loginLogData.Status = 0
		loginLogData.Msg = "登陆失败"
		_ = service.LoginLogService.AsyncSaveLoginLog(ctx, &loginLogData)
		return "", nil
	}

	// 再验证captcha验证码
	//if !service.LoginService.Verify(ctx, req.CaptchaId, req.Captcha) {
	//	c.FailWithMessage(r, "验证码输入错误")
	//	loginLogData.Status = 0
	//	loginLogData.Msg = "登陆失败"
	//	_ = service.LoginLogService.AsyncSaveLoginLog(ctx, &loginLogData)
	//	return "", nil
	//}

	loginLogData.Status = 1
	loginLogData.Msg = "登录成功"

	// 登录成功需要将用户的信息设置到request中继续传递下去
	r.SetParam("adminInfo", adminHasOneRole)

	// 异步保存登录日志
	err = service.LoginLogService.AsyncSaveLoginLog(ctx, &loginLogData)
	if err != nil {
		g.Log().Info("保存登录日志失败")
		// 这里登录日志保存失败，不需要管，不影响登录的进行
	}
	// 第二个返回的参数被保存在gtoken.RespData中的data字段
	return req.Username, adminHasOneRole
}

func (api *adminAuthApi) LoginAfter(r *ghttp.Request, respData gtoken.Resp) {
	// 登录之后会有登录成功或者失败的标识
	ctx := r.GetCtx()
	if !respData.Success() {
		// 登录失败，直接返回
		g.Log().Info("检测到登录失败")
		return
	}

	// 登录成功，获取uuid, token，以及admin的基本信息
	uuid, token := respData.GetString("uuid"), respData.GetString("token")
	var adminInfo *system.AdminHasOneRole
	if err := r.GetParamVar("adminInfo").Scan(&adminInfo); err != nil {
		g.Log().Info("解析admin信息失败")
		return
	}

	ua := r.Header.Get("User-Agent")
	userAgent := user_agent.New(ua)
	os := userAgent.OS()
	browser, _ := userAgent.Browser()
	onlineData := model.OnlineUser{
		Uuid:      uuid,
		Token:     token,
		Username:  adminInfo.Admin.Username,
		Os:        os,
		Explorer:  browser,
		LoginTime: gtime.Now(),
		IsAdmin:   1,
	}

	service.OnlineUserService.AsyncSaveOnlineUser(ctx, &onlineData)
}

func (api *adminAuthApi) LogoutBefore(r *ghttp.Request) bool {
	ctx := r.GetCtx()
	// 尝试从 Authorization 头部中去获取 token
	header := r.Header.Get("Authorization")
	if header != "" {
		h := strings.Split(header, "")
		token := h[1]
		if token != "" {
			service.OnlineUserService.DeleteOnlineUserByToken(ctx, token)
			return true
		}
	}

	// 如果 Authorization中没有获取到，就要从 request的请求参数param中去获取
	token := r.GetString("token")
	if token != "" {
		service.OnlineUserService.DeleteOnlineUserByToken(ctx, token)
		return true
	}
	return false
}

func (api *adminAuthApi) LogoutAfter(r *ghttp.Request, respData gtoken.Resp) {
	return
}

func (api *adminAuthApi) AuthBefore(r *ghttp.Request) bool {
	return true
}

func (api *adminAuthApi) AuthAfter(r *ghttp.Request, respData gtoken.Resp) {
	if respData.Success() {
		r.Middleware.Next()
	} else {
		c.FailWithMessage(r, "Token已过期或无效，请重新登录")
	}
}
