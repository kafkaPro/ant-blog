package request

type CreateApiReq struct {
	Path        string `p:"path" v:"path@required|length:1,100#请输入api路径|api路径长度为:min到:max位"`
	Description string `p:"description" v:"description@required|length:1,100#请输入api中文描述|api中文描述长度为:min到:max位"`
	ApiGroup    string `p:"apiGroup" v:"apiGroup@required|length:1,100#请输入api组|api组描述长度为:min到:max位"`
	Method      string `p:"method" v:"method@required|length:1,100#请输入api请求方法|api请求方法长度为:min到:max位"`
}

type DeleteApiReq struct {
	Id     int    `p:"id" v:"method@required|length:1,1000#请输入ID|ID长度为:min到:max位"`
	Path   string `p:"path" v:"path@required|length:1,100#请输入api路径|api路径长度为:min到:max位"`
	Method string `p:"method" v:"method@required|length:1,100#请输入api请求方法|api请求方法长度为:min到:max位"`
}

type UpdateApiReq struct {
	Id          int    `p:"id" v:"method@required|length:1,1000#请输入ID|ID长度为:min到:max位"`
	Path        string `p:"path" v:"path@required|length:1,100#请输入api路径|api路径长度为:min到:max位"`
	Description string `p:"description" v:"description@required|length:1,100#请输入api中文描述|api中文描述长度为:min到:max位"`
	ApiGroup    string `p:"apiGroup" v:"apiGroup@required|length:1,100#请输入api组|api组描述长度为:min到:max位"`
	Method      string `p:"method" v:"method@required|length:1,100#请输入api请求方法|api请求方法长度为:min到:max位"`
}

type GetApiByIdReq struct {
	Id int `p:"id" v:"method@required|length:1,1000#请输入ID|ID长度为:min到:max位"`
}

type GetApiListReq struct {
	Id          int    `p:"id"`
	Path        string `p:"path"`
	Description string `p:"description"`
	ApiGroup    string `p:"apiGroup"`
	Method      string `p:"method"`
	OrderKey    string `p:"orderKey"`
	Desc        bool   `p:"desc"`
	PageReq
}
