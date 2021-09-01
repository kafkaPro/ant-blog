package service

import (
	"blog-backend/app/api/request"
	"blog-backend/app/dao"
	"blog-backend/app/model"
	"blog-backend/app/model/database"
	"blog-backend/app/model/system"
	"blog-backend/library/utils"
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type dictionaryService struct{}

var (
	DictionaryService = new(dictionaryService)
)

func (s *dictionaryService) CreateDictionary(ctx context.Context, req *request.CreateDictionaryReq) (err error) {
	// 检查dict记录是否存在
	condition := g.Map{"name": req.Name, "type": req.Type}
	if s.CheckIfDictionaryExist(ctx, condition) {
		return gerror.New("字典已经存在，无法添加")
	}

	data := model.Dictionaries{
		Name: req.Name,
		Type: req.Type,
		Desc: req.Desc,
	}

	if req.Status {
		data.Status = 1
	}

	_, err = dao.Dictionaries.Ctx(ctx).OmitEmpty().Insert(data)
	return
}

func (s *dictionaryService) DeleteDictionary(ctx context.Context, req *request.DeleteDictionaryReq) (err error) {
	condition := g.Map{"id": req.Id}
	if !s.CheckIfDictionaryExist(ctx, condition) {
		return gerror.New("字典不存在，无法删除")
	}

	err = g.DB(database.Db).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 删除 dictionaries 表信息
		_, e := tx.Ctx(ctx).Model(dao.Dictionaries.Table).Delete(dao.Dictionaries.Columns.Id, req.Id)
		if e != nil {
			return e
		}

		// 删除 dictionary_details 表中的信息
		_, e = tx.Ctx(ctx).Model(dao.DictionaryDetails.Table).Delete(dao.DictionaryDetails.Columns.DictionaryId, req.Id)
		if e != nil {
			return e
		}

		return nil
	})
	return
}

func (s *dictionaryService) UpdateDictionary(ctx context.Context, req *request.UpdateDictionaryReq) (err error) {
	condition := g.Map{"id": req.Id}
	if !s.CheckIfDictionaryExist(ctx, condition) {
		return gerror.New("字典不存在，无法更新")
	}

	data := model.Dictionaries{
		Name:   req.Name,
		Type:   req.Type,
		Status: req.Status,
		Desc:   req.Desc,
	}

	_, err = dao.Dictionaries.Ctx(ctx).
		FieldsEx(dao.Dictionaries.Columns.Id).
		OmitEmpty().
		Update(data, dao.DictionaryDetails.Columns.Id, req.Id)
	return
}

func (s *dictionaryService) FindDictionary(ctx context.Context, req *request.FindDictionaryReq) (
	dict *system.DictionaryHasManyDetails, err error) {
	dict = new(system.DictionaryHasManyDetails)
	// 检查dictionary是否存在
	condition := g.Map{"id": req.Id}
	if !s.CheckIfDictionaryExist(ctx, condition) {
		return nil, gerror.New("字典不存在")
	}

	err = g.DB(database.Db).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 查找 dictionaries 表
		e := tx.Ctx(ctx).Model(dao.Dictionaries.Table).
			Scan(&dict.Dictionaries, dao.Dictionaries.Columns.Id, req.Id)
		if e != nil {
			return e
		}

		// 查找 dictionary_details 表
		e = tx.Ctx(ctx).Model(dao.DictionaryDetails.Table).
			Scan(&dict.DictionaryDetails, dao.DictionaryDetails.Columns.DictionaryId, req.Id)
		if e != nil {
			return e
		}

		return nil
	})
	return
}

func (s *dictionaryService) GetDictionaryListInPageMode(ctx context.Context, req *request.GetDictionaryListReq) (
	list interface{}, total int, err error) {
	m := dao.Dictionaries.Ctx(ctx).Safe()
	if req.Name != "" {
		m = m.WhereLike(dao.Dictionaries.Columns.Name, utils.FuzzyAll(req.Name))
	}
	if req.Type != "" {
		m = m.WhereLike(dao.Dictionaries.Columns.Type, utils.FuzzyAll(req.Type))
	}
	if req.Desc != "" {
		m = m.WhereLike(dao.Dictionaries.Columns.Desc, utils.FuzzyAll(req.Desc))
	}
	if req.Status {
		m = m.Where(dao.Dictionaries.Columns.Status, gconv.String(req.Status))
	}

	total, err = m.Count()
	err = m.Page(req.PageNo, req.PageSize).Order(dao.Dictionaries.Columns.Id).Scan(&list)
	return
}

func (s *dictionaryService) CheckIfDictionaryExist(ctx context.Context, condition interface{}) bool {
	c, err := dao.Dictionaries.Ctx(ctx).Count(condition)
	if err != nil || c == 0 {
		return false
	}
	return true
}
