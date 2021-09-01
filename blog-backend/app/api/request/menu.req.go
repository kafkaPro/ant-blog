package request

import (
	"blog-backend/app/model/system"
	"github.com/gogf/gf/os/gtime"
)

type UpdateBaseMenuReq struct {
	Id         uint         `p:"Id"`
	CreateAt   *gtime.Time  `p:"CreateAt"`
	UpdateAt   *gtime.Time  `p:"UpdateAt"`
	DeleteAt   *gtime.Time  `p:"DeleteAt"`
	ParentId   uint         `p:"parentId" v:"required#请输入父id"`
	Path       string       `p:"path" v:"required|length:1,1000#请输入路由path|路由path长度为:min到:max位"`
	Name       string       `p:"name" v:"required|length:1,1000#请输入路由Name|路由Name长度为:min到:max位"`
	Hidden     bool         `p:"hidden" v:"required|length:1,1000#请输入是否在列表隐藏|是否在列表隐藏长度为:min到:max位"`
	Component  string       `p:"component" v:"required|length:1,1000#请输入前端文件路径|前端文件路径长度为:min到:max位"`
	Sort       int          `p:"sort" v:"required|length:1,1000#请输入排序标记|排序标记长度为:min到:max位"`
	Parameters []Parameters `p:"parameters"` // 地址栏携带参数
	Meta       Meta         `p:"meta"`       // 附加属性
}

type Meta struct {
	Title       string `p:"title" v:"required|length:1,1000#请输入菜单名|id长度为:min到:max位"`
	Icon        string `p:"icon" v:"required|length:1,1000#请输入菜单图标|id长度为:min到:max位"`
	KeepAlive   bool   `p:"keepAlive" v:"required|length:1,1000#请输入是否缓存|是否缓存长度为:min到:max位"`
	DefaultMenu bool   `p:"defaultMenu" v:"required|length:1,1000#请输入是否是基础路由(开发中)|是否是基础路由(开发中)长度为:min到:max位"`
}

type Parameters struct {
	Type  string `p:"type" ` // 地址栏携带参数为params还是query
	Key   string `p:"key"  ` // 地址栏携带参数的key
	Value string `p:"value"` // 地址栏携带参数的值
}

type CreateBaseMenuReq struct {
	ParentId   uint         `p:"parentId" v:"required|length:1,1000#请输入父菜单ID|父菜单ID长度为:min到:max位"`
	Path       string       `p:"path" v:"required|length:1,1000#请输入路由path|路由path长度为:min到:max位"`
	Name       string       `p:"name" v:"required|length:1,1000#请输入路由Name|路由Name长度为:min到:max位"`
	Hidden     bool         `p:"hidden" v:"required|length:1,1000#请输入是否在列表隐藏|是否在列表隐藏长度为:min到:max位"`
	Component  string       `p:"component" v:"required|length:1,1000#请输入前端文件路径|前端文件路径长度为:min到:max位"`
	Sort       int          `p:"sort" v:"required|length:1,1000#请输入排序标记|排序标记长度为:min到:max位"`
	Parameters []Parameters `p:"parameters"` // 地址栏携带参数
	Meta       Meta         `p:"meta"`       // 附加属性
}

type AddMenuRoleReq struct {
	Menus  []system.BaseMenu `p:"menus"`
	RoleId uint              `p:"roleId"`
}

type DeleteBaseMenuReq struct {
	MenuId uint `p:"menuId"   v:"required#请提供要删除的menuId"`
}

type MenuIdParam struct {
	MenuId uint `p:"menuId"   v:"required#请提供要查询的menuId"`
}
