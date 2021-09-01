package test

import (
	"blog-backend/app/model/system"
	"blog-backend/app/service"
	"context"
	"fmt"
	"github.com/gogf/gf/encoding/gparser"
	"testing"
)

func TestGetRoleMenuMap(t *testing.T) {
	roleId := uint(888)
	treeMap, err := service.MenuService.GetMenuTreeMap(context.TODO(), roleId)
	if err == nil {
		for k, v := range treeMap {
			fmt.Println("parentId:", k)
			for _, m := range v {
				fmt.Println(m)
			}
		}
	}
}

func TestBuildMenuChildren(t *testing.T) {
	treeMap, err := service.MenuService.GetMenuTreeMap(context.TODO(), uint(888))
	fmt.Println(gparser.VarToJsonString(treeMap))
	node := system.RoleMenuNode{
		RoleId: uint(888),
		MenuId: uint(1),
	}
	err = service.MenuService.GetMenuChildrenList(context.TODO(), &node, treeMap)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(gparser.VarToJsonString(node))
	}
}

func TestGetMenuTree(t *testing.T) {
	tree, err := service.MenuService.GetMenuTree(context.TODO(), 888)
	if err == nil {
		fmt.Println(gparser.VarToJsonString(tree))
	}
}

func TestGetBaseMenuTree(t *testing.T) {
	tree, err := service.MenuService.GetBaseMenuTree(context.TODO())
	if err == nil {
		fmt.Println(gparser.VarToJsonString(tree))
	}
}
