package request

type CasbinInfo struct {
	Path   string `p:"path"  json:"path"`
	Method string `p:"method" json:"method"`
}

// CasbinQueryParameter
// casbin的输入请求参数
type CasbinQueryParameter struct {
	AuthorityId string       `p:"authorityId"   json:"authorityId"`
	CasbinInfos []CasbinInfo `p:"casbinInfos"   json:"casbinInfos"`
}

type CasbinModel struct {
	ID          uint   `json:"id"`
	Ptype       string `json:"ptype"`
	AuthorityId string `json:"rolename"`
	Path        string `json:"path"`
	Method      string `json:"method"`
}
