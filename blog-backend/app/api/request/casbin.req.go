package request

type CasbinInfo struct {
	Path   string `p:"path"   json:"path"`
	Method string `p:"method" json:"method"`
}

// CasbinQueryParameter
// casbin的输入请求参数
type CasbinQueryParameter struct {
	RoleId      string       `p:"roleId"        json:"roleId"`
	CasbinInfos []CasbinInfo `p:"casbinInfos"   json:"casbinInfos"`
}

type CasbinModel struct {
	ID     uint   `json:"id"`
	PType  string `json:"pType"`
	RoleId string `json:"roleId"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
