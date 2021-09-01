package test

import (
	"blog-backend/app/api/request"
	"blog-backend/app/service"
	"context"
	"fmt"
	"github.com/gogf/gf/encoding/gparser"
	"testing"
)

func TestDictionary(t *testing.T) {
	dict, err := service.DictionaryService.FindDictionary(context.TODO(), &request.FindDictionaryReq{Id: 2})
	if err == nil {
		fmt.Println(gparser.VarToJsonString(dict))
	} else {
		fmt.Println("获取字典信息失败")
	}
}
