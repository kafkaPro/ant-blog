package request

// 分页相关请求参数

type PageReq struct {
	PageNo   int `p:"page" v:"required|length:1,1000#请输入页数|页数长度为:min到:max位" json:"page" form:"page"`
	PageSize int `p:"pageSize" v:"required|length:1,1000#请输入每页大小|每页大小为:min到:max位" json:"pageSize" form:"pageSize"`
}
