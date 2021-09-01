package casbin

type CasbinModel struct {
	ID     uint   `json:"id"`
	Ptype  string `json:"ptype"`
	RoleId uint   `json:"roleId"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
