package service

import (
	"blog-backend/app/api/request"
	"blog-backend/app/dao"
	"blog-backend/library/gdbadapter"
	"context"
	"github.com/casbin/casbin/util"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

type casbinService struct{}

var (
	CasbinService = new(casbinService)
)

func (s *casbinService) AddCasbin(ctx context.Context, casbinModel request.CasbinModel) bool {
	e := Casbin()
	success, _ := e.AddPolicy(casbinModel.RoleId, casbinModel.Path, casbinModel.Method)
	return success
}

func (s *casbinService) ClearCasbin(ctx context.Context, v int, p ...string) bool {
	e := Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

// UpdateCasbinApi
// 主要是修改 casbin_rule 表中的api信息
func (s *casbinService) UpdateCasbinApi(ctx context.Context, oldPath, newPath, oldMethod, newMethod string) error {
	_, err := dao.CasbinRule.Ctx(ctx).Data(g.Map{"v1": newPath, "v2": newMethod}).Where(g.Map{"v1": oldPath}).Update()
	return err
}

// UpdateCasbin
// 修改casbin的规则，不涉及数据库表的操作
func (s *casbinService) UpdateCasbin(ctx context.Context, roleId uint, casbinInfos []request.CasbinInfo) error {
	s.ClearCasbin(ctx, 0, gconv.String(roleId))
	var rules [][]string
	for _, v := range casbinInfos {
		rules = append(rules, []string{gconv.String(roleId), v.Path, v.Method})
	}
	e := Casbin()
	success, _ := e.AddPolicy(rules)
	if !success {
		return gerror.New("存在相同的api, casbin规则添加失败，请联系管理员")
	}
	return nil
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

func (s *casbinService) GetCasbinPolicyByRoleId(ctx context.Context, roleId uint) (pathMaps []request.CasbinInfo) {
	e := Casbin()
	list := e.GetFilteredPolicy(0, gconv.String(roleId))
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}
