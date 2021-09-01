package service

import (
	"blog-backend/app/api/request"
	"blog-backend/app/dao"
	"blog-backend/app/model"
	"blog-backend/library/utils"
	"context"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type dictionaryDetailService struct{}

var (
	DictionaryDetailService = new(dictionaryDetailService)
)

func (s *dictionaryDetailService) CreateDictionaryDetails(ctx context.Context, req *request.CreateDictionaryDetailReq) (
	err error) {
	condition := g.Map{"label": req.Label}
	if s.CheckIfDictionaryDetailExist(ctx, condition) {
		return gerror.New("字典数据已经存在，无法创建")
	}

	data := model.DictionaryDetails{
		Label:        req.Label,
		Value:        req.Value,
		Status:       gconv.Int(req.Status),
		DictionaryId: req.DictionaryId,
	}

	_, err = dao.DictionaryDetails.Ctx(ctx).OmitEmpty().Insert(data)
	return
}

func (s *dictionaryDetailService) DeleteDictionaryDetails(ctx context.Context, req *request.DeleteDictionaryDetailReq) (
	err error) {
	condition := g.Map{"id": req.Id}
	if !s.CheckIfDictionaryDetailExist(ctx, condition) {
		return gerror.New("字典不存在，无法删除")
	}

	_, err = dao.DictionaryDetails.Ctx(ctx).Delete(dao.DictionaryDetails.Columns.Id, req.Id)
	return
}

func (s *dictionaryDetailService) UpdateDictionaryDetails(ctx context.Context, req *request.UpdateDictionaryDetailReq) (
	err error) {
	condition := g.Map{"id": req.Id}
	if !s.CheckIfDictionaryDetailExist(ctx, condition) {
		return gerror.New("字典数据不存在，无法进行更新")
	}

	data := model.DictionaryDetails{
		Label:        req.Label,
		Value:        req.Value,
		Status:       gconv.Int(req.Status),
		Sort:         req.Sort,
		DictionaryId: req.DictionaryId,
	}

	_, err = dao.DictionaryDetails.Ctx(ctx).OmitEmpty().Safe().
		FieldsEx(dao.DictionaryDetails.Columns.Id).Update(data, dao.DictionaryDetails.Columns.Id, req.Id)
	return
}

func (s *dictionaryDetailService) FindDictionaryDetails(ctx context.Context, req *request.FindDictionaryDetailReq) (
	dictionaryDetails *model.DictionaryDetails, err error) {
	condition := g.Map{"id": req.Id}
	if !s.CheckIfDictionaryDetailExist(ctx, condition) {
		return nil, gerror.New("字典数据不存在")
	}

	err = dao.DictionaryDetails.Ctx(ctx).Where(dao.DictionaryDetails.Columns.Id, req.Id).Scan(&dictionaryDetails)
	return
}

func (s *dictionaryDetailService) GetDictionaryDetailsListInPageMode(ctx context.Context,
	req *request.GetDictionaryDetailListReq) (list []*model.DictionaryDetails, total int, err error) {
	list = ([]*model.DictionaryDetails)(nil)
	m := dao.DictionaryDetails.Ctx(ctx).Safe()

	if req.Label != "" {
		m = m.WhereLike(dao.DictionaryDetails.Columns.Label, utils.FuzzyAll(req.Label))
	}
	if req.Value >= 0 {
		m = m.Where(dao.DictionaryDetails.Columns.Value, string(rune(req.Value)))
	}
	if req.Sort >= 0 {
		m = m.Where(dao.DictionaryDetails.Columns.Sort, req.Sort)
	}
	if req.DictionaryId > 0 {
		m = m.Where(dao.DictionaryDetails.Columns.DictionaryId, req.DictionaryId)
	}
	if req.Status {
		m = m.Where(dao.DictionaryDetails.Columns.Status, 1)
	}

	err = m.Page(req.PageNo, req.PageSize).Order(dao.DictionaryDetails.Columns.Id).Scan(&list)
	return
}

func (s *dictionaryDetailService) CheckIfDictionaryDetailExist(ctx context.Context, condition interface{}) bool {
	c, err := dao.DictionaryDetails.Ctx(ctx).Count(condition)
	if err != nil || c == 0 {
		return false
	}
	return true
}
