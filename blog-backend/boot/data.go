package boot

import (
	"blog-backend/app/model/database"
	"github.com/gogf/gf/frame/g"
)

func InitializeDataTableAndData() {
	var err error
	// 初始化表,如果出错则不会执行初始化数据
	err = database.TableApis()
	err = database.TableJwts()
	err = database.TableMenus()
	err = database.TableFiles()
	err = database.TableAdmins()
	err = database.TableCustomers()
	err = database.TableCasbinRule()
	err = database.TableOperations()
	err = database.TableParameters()
	err = database.TableAuthorities()
	//err = database.TableSimpleUpload()
	err = database.TableDictionaries()
	err = database.TableAuthorityMenu()
	err = database.TableBreakpointFiles()
	err = database.TableBreakpointChucks()
	err = database.TableDictionaryDetails()
	err = database.TableAuthorityResources()
	if err != nil {
		g.Log().Error(err)
		return
	}
	// 初始化数据,并且数据插入是10条10条这样插入的,每个表插入数据都有加事务
	err = database.DataApis()
	err = database.DataFiles()
	err = database.DataMenus()
	err = database.DataAdmins()
	err = database.DataCasbinRule()
	err = database.DataAuthorities()
	err = database.DataDictionaries()
	err = database.DataAuthorityMenus()
	err = database.DataDictionaryDetails()
	if err != nil {
		g.Log().Error(err)
		return
	}
	return
}
