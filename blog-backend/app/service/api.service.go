package service

import (
	"blog-backend/app/api/request"
	"blog-backend/app/dao"
	"blog-backend/app/model"
	"blog-backend/app/model/database"
	"blog-backend/library/utils"
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
)

type apiService struct{}

var (
	ApiService = new(apiService)
)

func (s *apiService) CreateApi(ctx context.Context, req *request.CreateApiReq) (err error) {
	condition := g.Map{"path": req.Path, "method": req.Method}
	if s.CheckIfApiExist(ctx, condition) {
		return gerror.New("api已经存在，无法创建")
	}

	data := model.Apis{
		Path:        req.Path,
		Description: req.Description,
		ApiGroup:    req.ApiGroup,
		Method:      req.Method,
	}

	_, err = dao.Apis.Ctx(ctx).OmitEmpty().Insert(data)
	return
}

func (s *apiService) UpdateApi(ctx context.Context, req *request.UpdateApiReq) (err error) {
	condition := g.Map{"id": req.Id}
	if !s.CheckIfApiExist(ctx, condition) {
		return gerror.New("api不存在，无法修改")
	}

	oldApi := model.Apis{}
	err = dao.Apis.Ctx(ctx).Where(condition).Scan(&oldApi)
	if err != nil {
		return
	}

	data := model.Apis{
		Id:          uint(req.Id),
		Path:        req.Path,
		Method:      req.Method,
		Description: req.Description,
		ApiGroup:    req.ApiGroup,
	}

	// 更新apis表
	_, err = dao.Apis.Ctx(ctx).FieldsEx(dao.Apis.Columns.Id).Update(data, dao.Apis.Columns.Id)
	// 更新 casbin_rule 表
	err = CasbinService.UpdateCasbinApi(ctx, oldApi.Path, oldApi.Method, req.Path, req.Method)
	return
}

func (s *apiService) DeleteApi(ctx context.Context, req *request.DeleteApiReq) (err error) {

	err = g.DB(database.Db).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 先删除 apis 表中的数据
		_, e := tx.Ctx(ctx).Model(dao.Apis.Table).Delete(g.Map{dao.Apis.Columns.Id: req.Id})
		if e != nil {
			return e
		}
		// 再清除casbin中的数据
		if !CasbinService.ClearCasbin(ctx, 1, req.Path, req.Method) {
			return gerror.New("删除casbin规则失败")
		}

		return nil
	})
	return
}

func (s *apiService) CheckIfApiExist(ctx context.Context, condition interface{}) bool {
	c, err := dao.Apis.Ctx(ctx).Count(condition)
	if err != nil || c == 0 {
		return false
	}
	return true
}

func (s *apiService) GetApiById(ctx context.Context, req *request.GetApiByIdReq) (api *model.Apis, err error) {
	err = dao.Apis.Ctx(ctx).Scan(&api, dao.Apis.Columns.Id, req.Id)
	return
}

func (s *apiService) GetAllApis(ctx context.Context) (list []*model.Apis, err error) {
	err = dao.Apis.Ctx(ctx).WhereNot(dao.Apis.Columns.Id, "").Scan(&list)
	return
}

func (s *apiService) GetApiListInPageMode(ctx context.Context, req *request.GetApiListReq) (
	list []*model.Apis, total int, err error) {
	m := dao.Apis.Ctx(ctx).Safe()

	// 可选参数
	if req.Path != "" {
		m = m.WhereLike(dao.Apis.Columns.Path, utils.FuzzyAll(req.Path))
	}
	if req.Description != "" {
		m = m.WhereLike(dao.Apis.Columns.Description, utils.FuzzyAll(req.Description))
	}
	if req.Method != "" {
		m = m.WhereLike(dao.Apis.Columns.Method, utils.FuzzyAll(req.Method))
	}
	if req.ApiGroup != "" {
		m = m.WhereLike(dao.Apis.Columns.ApiGroup, utils.FuzzyAll(req.ApiGroup))
	}

	total, err = m.Count()
	// 字段排序配置
	if req.Desc {
		m.OrderDesc(req.OrderKey)
	} else {
		m.OrderAsc(req.OrderKey)
	}

	err = m.Page(req.PageNo, req.PageSize).Scan(&list)
	return
}
