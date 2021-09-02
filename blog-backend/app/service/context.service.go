package service

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

const (
	// ContextKey 常量表示存储在context.Context上下文变量中的键名，
	// 该键名用于从传递的context.Context变量中存储/获取业务自定义的共享变量
	ContextKey = "AdminCtxKey"
)

// AdminContext 请求上下文结构
type AdminContext struct {
	// Context结构体中的Session表示当前请求的Session对象，在GoFrame框架中每个 HTTP请求对象
	// 中都会有一个空的Session对象，该对象采用了懒初始化设计，只有在真正执行读写操作时才会初始化。
	Session *ghttp.Session
	// User表示当前登录的用户基本信息，只有在用户登录后才有数据，否则是nil
	User *ContextAdmin
	// Data属性用于存储自定义的KV变量，因此一般来说开发者无需再往context.Context上下文变量
	// 中增加自定义的键值对，而是直接使用model.Context对象的这个Data属性即可
	Data g.Map
}

// ContextAdmin 请求上下文中的用户信息
type ContextAdmin struct {
	Id        uint   // 用户id
	NickName  string // 用户昵称
	UserName  string // 用户名
	HeaderImg string // 用户头像
	Status    int    // 用户状态
	IsAdmin   int    // 是否是管理员
	RoleId    uint   // 角色id
	Uuid      string // uuid
}

// 由于该上下文是业务逻辑相关的，因此需要通过service对象将上下文变量封装起来

type adminContextService struct{}

var AdminContextService = new(adminContextService)

// Init 初始上下文对象指针到上下文对象中，以便后续的请求可以修改
func (s *adminContextService) Init(r *ghttp.Request, customCtx *AdminContext) {
	r.SetCtxVar(ContextKey, customCtx)
}

// Get 从上下文中获取
func (s *adminContextService) Get(ctx context.Context) *AdminContext {
	v := ctx.Value(ContextKey)
	if v == nil {
		return nil
	}
	// 类型转化
	if localCtx, ok := v.(*AdminContext); ok {
		return localCtx
	}
	return nil
}

// SetUser 将用户上下文设置到系统上下文中
func (s *adminContextService) SetUser(ctx context.Context, ctxUser *ContextAdmin) {
	s.Get(ctx).User = ctxUser
}
