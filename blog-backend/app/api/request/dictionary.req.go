package request

type CreateDictionaryReq struct {
	Name   string `p:"name" v:"required|length:1,1000#请输入字典中文名|字典中文名长度为:min到:max位"`
	Type   string `p:"type" v:"required|length:1,1000#请输入字典英文名|字典英文名长度为:min到:max位"`
	Status bool   `p:"status" v:"required|length:1,1000#请输入状态|状态长度为:min到:max位"`
	Desc   string `p:"desc" v:"required|length:1,1000#请输入描述|描述长度为:min到:max位"`
}

type DeleteDictionaryReq struct {
	Id float64 `p:"id" v:"required|length:1,1000#请输入id|id长度为:min到:max位" orm:"id,primary"` // 自增ID
}

type UpdateDictionaryReq struct {
	Id     float64 `p:"id" v:"required|length:1,1000#请输入id|id长度为:min到:max位"` // 自增ID
	Name   string  `p:"name" v:"required|length:1,1000#请输入字典中文名|字典中文名长度为:min到:max位"`
	Type   string  `p:"type" v:"required|length:1,1000#请输入字典英文名|字典英文名长度为:min到:max位"`
	Status int     `p:"status" v:"required|length:1,1000#请输入状态|状态长度为:min到:max位"`
	Desc   string  `p:"desc" v:"required|length:1,1000#请输入描述|描述长度为:min到:max位" `
}

type GetDictionaryReq struct {
	ID   float64 `p:"id" v:"required|length:1,1000#请输入id|id长度为:min到:max位"`
	Type string  `p:"type" v:"required|length:1,1000#请输入字典英文名|字典英文名长度为:min到:max位"`
}

type FindDictionaryReq struct {
	Id   float64 `p:"id" v:"required|length:1,1000#请输入id|id长度为:min到:max位"` // 自增ID
	Type string  `p:"type" `
}

type GetDictionaryListReq struct {
	Name   string `p:"name"`
	Type   string `p:"type"`
	Status bool   `p:"status"`
	Desc   string `p:"desc"`
	PageReq
}
