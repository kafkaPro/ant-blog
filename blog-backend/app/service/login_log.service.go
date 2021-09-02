package service

import (
	"blog-backend/app/api/request"
	"blog-backend/app/dao"
	"blog-backend/app/model"
	"context"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/os/grpool"
)

type loginLogService struct {
	Pool *grpool.Pool
}

var (
	LoginLogService = &loginLogService{
		Pool: grpool.New(200),
	}
)

// AsyncSaveLoginLog 异步保存登录日志
func (s *loginLogService) AsyncSaveLoginLog(ctx context.Context, data *model.LoginLog) error {
	err := s.Pool.Add(func() {
		// s.DoSaveLoginLog(ctx, req) 异步插入时会出现"context canceled"错误
		_ = s.DoSaveLoginLog(context.TODO(), data)
	})
	return err
}

// DoSaveLoginLog 实际执行写入登录日志的操作
func (s *loginLogService) DoSaveLoginLog(ctx context.Context, data *model.LoginLog) error {
	// dao层写入
	_, err := dao.LoginLog.Ctx(ctx).OmitEmpty().Insert(data)
	return err
}

func (s *loginLogService) GetLoginLogList(ctx context.Context, req *request.LoginLogQueryReq) (
	total, page int, list []*model.LoginLog, err error) {

	m := dao.LoginLog.Ctx(ctx).Unscoped()
	if req.Status == 0 || req.Status == 1 {
		m = m.Where(dao.LoginLog.Columns.Status, req.Status)
	}
	if req.Id != 0 {
		m = m.Where(dao.LoginLog.Columns.Id, req.Id)
	}
	if req.UserName != "" {
		m = m.Where(dao.LoginLog.Columns.UserName, req.UserName)
	}
	if req.Ip == "" {
		m = m.Where(dao.LoginLog.Columns.Ip, req.Ip)
	}
	if req.Location != "" {
		m = m.Where(dao.LoginLog.Columns.LoginLocation, req.Location)
	}
	if req.Browser != "" {
		m = m.Where(dao.LoginLog.Columns.Browser, req.Browser)
	}
	if req.Os != "" {
		m = m.Where(dao.LoginLog.Columns.Os, req.Os)
	}
	if req.Msg != "" {
		m = m.Where(dao.LoginLog.Columns.Msg, req.Msg)
	}
	if req.LoginTime != nil {
		m = m.Where(dao.LoginLog.Columns.LoginTime, req.LoginTime)
	}

	// 总记录条数
	total, err = m.Count()
	if err != nil {
		return 0, 0, nil, err
	}

	// 执行分页查询
	err = m.Page(req.PageNo, req.PageSize).Order(dao.LoginLog.Columns.Id).Scan(&list)
	if err != nil {
		return 0, 0, nil, err
	}

	// 执行成功
	return
}

func (s *loginLogService) DeleteLoginLogById(ctx context.Context, id int64) error {
	if id == 0 {
		return gerror.New("无效的日志Id，无法删除")
	}

	_, err := dao.LoginLog.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (s *loginLogService) ClearLoginLog(ctx context.Context) error {
	_, err := dao.LoginLog.Ctx(ctx).WhereGT(dao.LoginLog.Columns.Id, 0).Delete()
	if err != nil {
		return err
	}
	return nil
}
