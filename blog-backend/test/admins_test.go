package test

import (
	"blog-backend/app/api/request"
	system2 "blog-backend/app/model/system"
	"blog-backend/app/service"
	"context"
	"fmt"
	"testing"
)

func TestGetAdminList(t *testing.T) {
	req := request.PageReq{
		PageNo:   1,
		PageSize: 10,
	}

	list, total, err := service.AdminService.GetAdminList(context.TODO(), &req)
	if err == nil {
		fmt.Println(total)
		for _, v := range list.([]*system2.AdminHasOneAuthority) {
			fmt.Println(*(v.Admin))
			fmt.Println(*(v.Authority))
		}
	}
}

func TestFindAdmin(t *testing.T) {
	uuid := "b4c54e5a-d015-4f8c-9f01-624c527a63ae"
	admin, err := service.AdminService.FindAdminByUuid(context.TODO(), uuid)
	if err == nil {
		fmt.Println(*admin)
	}
}
