package service

import (
	"blog-backend/app/api/request"
	"blog-backend/library/gdbadapter"
	"context"
	"github.com/casbin/casbin/util"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/frame/g"
	"strings"
)

type casbinService struct{}

var (
	CasbinService = new(casbinService)
)

func (s *casbinService) AddCasbin(ctx context.Context, casbinModel request.CasbinModel) bool {
	e := Casbin()
	success, _ := e.AddPolicy(casbinModel.AuthorityId, casbinModel.Path, casbinModel.Method)
	return success
}

func (s *casbinService) ClearCasbin(ctx context.Context, v int, p ...string) bool {
	e := Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

func Casbin() *casbin.Enforcer {
	adapter, err := gdbadapter.NewAdapterByConfig()
	if err != nil {
		panic(err)
	}

	enforcer, err := casbin.NewEnforcer(g.Cfg("casbin").GetString("casbin.ModelPath"), adapter)
	if err != nil {
		panic(err)
	}

	enforcer.AddFunction("ParamsMatch", ParamMatchFunc)
	_ = enforcer.LoadPolicy()
	return enforcer
}

func ParamMatchFunc(args ...interface{}) (interface{}, error) {
	k1 := args[0].(string)
	k2 := args[1].(string)
	return ParamMatch(k1, k2), nil
}

func ParamMatch(k1, k2 string) bool {
	key1 := strings.Split(k1, "?")[0] // 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, k2)
}
