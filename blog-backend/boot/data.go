package boot

import (
	"blog-backend/app/model/db"
	"github.com/gogf/gf/frame/g"
)

func InitializeDataTableAndData() {
	var err error
	// 初始化表,如果出错则不会执行初始化数据
	err = db.TableApis()
	err = db.TableJwts()
	err = db.TableMenus()
	err = db.TableFiles()
	err = db.TableAdmins()
	err = db.TableCustomers()
	err = db.TableCasbinRule()
	err = db.TableOperations()
	err = db.TableParameters()
	err = db.TableAuthorities()
	//err = db.TableSimpleUpload()
	err = db.TableDictionaries()
	err = db.TableAuthorityMenu()
	err = db.TableBreakpointFiles()
	err = db.TableBreakpointChucks()
	err = db.TableDictionaryDetails()
	err = db.TableAuthorityResources()
	if err != nil {
		g.Log().Error(err)
		return
	}
	// 初始化数据,并且数据插入是10条10条这样插入的,每个表插入数据都有加事务
	err = db.DataApis()
	err = db.DataFiles()
	err = db.DataMenus()
	err = db.DataAdmins()
	err = db.DataCasbinRule()
	err = db.DataAuthorities()
	err = db.DataDictionaries()
	err = db.DataAuthorityMenus()
	err = db.DataDictionaryDetails()
	if err != nil {
		g.Log().Error(err)
		return
	}
	return
}
