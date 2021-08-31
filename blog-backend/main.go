package main

import (
	"blog-backend/boot"
	_ "blog-backend/boot"
	_ "blog-backend/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	// 初始化数据库内容，如果初始化成功之后，注释掉这一段代码
	boot.InitializeDataTableAndData()
	g.Server().Run()
}
