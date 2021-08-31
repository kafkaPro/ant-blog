package database

import "github.com/gogf/gf/frame/g"

var Db = g.Cfg("system").GetString("system.Db")
