package request

import (
	"github.com/gogf/gf/os/gtime"
)

type LoginLogSaveReq struct {
	Status   int    `p:"status"`
	UserName string `p:"username"`
	Browser  string `p:"browser"`
	Ip       string `p:"ip"`
	Msg      string `p:"msg"`
	Os       string `p:"os"`
	Location string `p:"location"`
}

// LoginLogQueryReq 登录日志查询请求
type LoginLogQueryReq struct {
	PageReq
	Id        int64       `p:"id"`
	UserName  string      `p:"username"`
	Ip        string      `p:"ip"`
	Browser   string      `p:"browser"`
	Os        string      `p:"os"`
	Status    int         `p:"status"`
	Msg       string      `p:"msg"`
	LoginTime *gtime.Time `p:"loginTime"`
	Location  string      `p:"location"`
}
