package service

import (
	"blog-backend/app/dao"
	"blog-backend/app/model"
	"context"
	"github.com/gogf/gf/os/grpool"
)

type onlineUserService struct {
	Pool *grpool.Pool
}

var (
	OnlineUserService = &onlineUserService{
		Pool: grpool.New(200),
	}
)

func (s *onlineUserService) AsyncSaveOnlineUser(ctx context.Context, data *model.OnlineUser) {
	_ = s.Pool.Add(func() {
		s.DoSaveOnlineUser(ctx, data)
	})
}

func (s *onlineUserService) DoSaveOnlineUser(ctx context.Context, data *model.OnlineUser) {
	_, _ = dao.OnlineUser.Ctx(ctx).Insert(data)
}

func (s *onlineUserService) CheckIfOnlineUserExist(ctx context.Context, condition interface{}) bool {
	c, err := dao.OnlineUser.Ctx(ctx).Count(condition)
	if err != nil || c == 0 {
		return false
	}
	return true
}

func (s *onlineUserService) DeleteOnlineUserByToken(ctx context.Context, token string) {
	_, _ = dao.OnlineUser.Ctx(ctx).Delete(dao.OnlineUser.Columns.Token, token)
}
