package test

import (
	"blog-backend/app/service"
	"context"
	"fmt"
	"github.com/gogf/gf/encoding/gparser"
	"testing"
)

func TestCasbin(t *testing.T) {
	casbinInfos := service.CasbinService.GetCasbinPolicyByRoleId(context.TODO(), 0)
	fmt.Println(gparser.VarToJsonString(casbinInfos))
}
