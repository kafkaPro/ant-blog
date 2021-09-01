package request

type CreateDictionaryDetailReq struct {
	Label        string `p:"label" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`           //
	Value        int    `p:"value" v:"value@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`           // 字典值
	Status       bool   `p:"status" v:"boolean@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`        // 启用状态
	Sort         int    `p:"sort" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`            // 排序标记
	DictionaryId int    `p:"sysDictionaryID" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"` // 关联标记
}

type DeleteDictionaryDetailReq struct {
	Id uint `p:"id" v:"id@required|length:1,1000#请输入id|id长度为:min到:max位"` // 自增ID
}

type UpdateDictionaryDetailReq struct {
	Id           float64 `p:"id" v:"id@required|length:1,1000#请输入id|id长度为:min到:max位"`                   // 自增ID
	Label        string  `p:"label" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`           //
	Value        int     `p:"value" v:"value@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`           // 字典值
	Status       bool    `p:"status" v:"boolean@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`        // 启用状态
	Sort         int     `p:"sort" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`            // 排序标记
	DictionaryId int     `p:"sysDictionaryID" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"` // 关联标记
}

type FindDictionaryDetailReq struct {
	Id           float64 `p:"id" v:"id@required|length:1,1000#请输入id|id长度为:min到:max位"`                   // 自增ID
	Label        string  `p:"label" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`           //
	Value        int     `p:"value" v:"value@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`           // 字典值
	Status       bool    `p:"status" v:"boolean@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`        // 启用状态
	Sort         int     `p:"sort" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"`            // 排序标记
	DictionaryId int     `p:"sysDictionaryID" v:"label@required|length:1,1000#请输入展示值|展示值长度为:min到:max位"` // 关联标记
}

type GetDictionaryDetailListReq struct {
	Label        string `p:"label"`
	Value        int    `p:"value"`           // 字典值
	Status       bool   `p:"status"`          // 启用状态
	Sort         int    `p:"sort"`            // 排序标记
	DictionaryId int    `p:"sysDictionaryID"` // 关联标记
	PageReq
}
