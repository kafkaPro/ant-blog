package v1

import (
	"blog-backend/app/api/request"
	c "blog-backend/app/model/common"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type adminApi struct{}

var (
	AdminApi = new(adminApi)
)

func (api *adminApi) ChangePassword(r *ghttp.Request) {
	var req *request.ChangePasswordReq
	if err := r.Parse(&req); err != nil {
		c.FailWithMessage(r, err.(gvalid.Error).Current().Error())
		r.Exit()
	}

}
