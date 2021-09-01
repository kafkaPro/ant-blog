// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

import (
	"github.com/gogf/gf/os/gtime"
)

// Admins is the golang structure for table admins.
type Admins struct {
	Id        uint        `orm:"id,primary" json:"id"`        // 自增ID
	CreateAt  *gtime.Time `orm:"create_at"  json:"createAt"`  // 创建时间
	UpdateAt  *gtime.Time `orm:"update_at"  json:"updateAt"`  // 更新时间
	DeleteAt  *gtime.Time `orm:"delete_at"  json:"deleteAt"`  // 删除时间
	Uuid      string      `orm:"uuid"       json:"uuid"`      // 用户唯一标识UUID
	Nickname  string      `orm:"nickname"   json:"nickname"`  // 用户昵称
	HeaderImg string      `orm:"header_img" json:"headerImg"` // 用户头像
	RoleId    uint        `orm:"role_id"    json:"roleId"`    // 用户角色ID
	Username  string      `orm:"username"   json:"username"`  // 用户名
	Password  string      `orm:"password"   json:"password"`  // 用户登录密码
}

// Apis is the golang structure for table apis.
type Apis struct {
	Id          uint        `orm:"id,primary"  json:"id"`          // 自增ID
	CreateAt    *gtime.Time `orm:"create_at"   json:"createAt"`    // 创建时间
	UpdateAt    *gtime.Time `orm:"update_at"   json:"updateAt"`    // 更新时间
	DeleteAt    *gtime.Time `orm:"delete_at"   json:"deleteAt"`    // 删除时间
	Path        string      `orm:"path"        json:"path"`        // api路径
	Description string      `orm:"description" json:"description"` // api中文描述
	ApiGroup    string      `orm:"api_group"   json:"apiGroup"`    // api组
	Method      string      `orm:"method"      json:"method"`      // 方法
}

// BreakpointChucks is the golang structure for table breakpoint_chucks.
type BreakpointChucks struct {
	Id              uint        `orm:"id,primary"        json:"id"`              // 自增ID
	CreateAt        *gtime.Time `orm:"create_at"         json:"createAt"`        // 更新时间
	UpdateAt        *gtime.Time `orm:"update_at"         json:"updateAt"`        // 更新时间
	DeleteAt        *gtime.Time `orm:"delete_at"         json:"deleteAt"`        // 删除时间
	ExaFileId       uint        `orm:"exa_file_id"       json:"exaFileId"`       // 文件id
	FileChunkPath   string      `orm:"file_chunk_path"   json:"fileChunkPath"`   // 切片路径
	FileChunkNumber int         `orm:"file_chunk_number" json:"fileChunkNumber"` // 切片标号
}

// BreakpointFiles is the golang structure for table breakpoint_files.
type BreakpointFiles struct {
	Id         uint        `orm:"id,primary"  json:"id"`         // 自增ID
	CreateAt   *gtime.Time `orm:"create_at"   json:"createAt"`   // 更新时间
	UpdateAt   *gtime.Time `orm:"update_at"   json:"updateAt"`   // 更新时间
	DeleteAt   *gtime.Time `orm:"delete_at"   json:"deleteAt"`   // 删除时间
	FileName   string      `orm:"file_name"   json:"fileName"`   // 文件名
	FileMd5    string      `orm:"file_md5"    json:"fileMd5"`    // 文件md5
	FilePath   string      `orm:"file_path"   json:"filePath"`   // 文件路径
	ChunkId    int         `orm:"chunk_id"    json:"chunkId"`    // 关联id
	ChunkTotal int         `orm:"chunk_total" json:"chunkTotal"` // 切片总数
	IsFinish   int         `orm:"is_finish"   json:"isFinish"`   // 是否完整
}

// CasbinRule is the golang structure for table casbin_rule.
type CasbinRule struct {
	Ptype string `orm:"ptype" json:"ptype"` //
	V0    string `orm:"v0"    json:"v0"`    //
	V1    string `orm:"v1"    json:"v1"`    //
	V2    string `orm:"v2"    json:"v2"`    //
	V3    string `orm:"v3"    json:"v3"`    //
	V4    string `orm:"v4"    json:"v4"`    //
	V5    string `orm:"v5"    json:"v5"`    //
}

// Customers is the golang structure for table customers.
type Customers struct {
	Id                uint        `orm:"id,primary"          json:"id"`                // 自增ID
	CreateAt          *gtime.Time `orm:"create_at"           json:"createAt"`          // 创建时间
	UpdateAt          *gtime.Time `orm:"update_at"           json:"updateAt"`          // 更新时间
	DeleteAt          *gtime.Time `orm:"delete_at"           json:"deleteAt"`          // 删除时间
	CustomerName      string      `orm:"customer_name"       json:"customerName"`      // 客户名
	CustomerPhoneData string      `orm:"customer_phone_data" json:"customerPhoneData"` // 客户电话
	SysUserId         uint        `orm:"sys_user_id"         json:"sysUserId"`         // 负责员工id
	SysUserRoleId     string      `orm:"sys_user_role_id"    json:"sysUserRoleId"`     // 负责员工角色
}

// Dictionaries is the golang structure for table dictionaries.
type Dictionaries struct {
	Id       uint        `orm:"id,primary" json:"id"`       // 自增ID
	CreateAt *gtime.Time `orm:"create_at"  json:"createAt"` // 创建时间
	UpdateAt *gtime.Time `orm:"update_at"  json:"updateAt"` // 更新时间
	DeleteAt *gtime.Time `orm:"delete_at"  json:"deleteAt"` // 删除时间
	Name     string      `orm:"name"       json:"name"`     // 字典名（中）
	Type     string      `orm:"type"       json:"type"`     // 字典名（英）
	Status   int         `orm:"status"     json:"status"`   // 状态
	Desc     string      `orm:"desc"       json:"desc"`     // 描述
}

// DictionaryDetails is the golang structure for table dictionary_details.
type DictionaryDetails struct {
	Id           uint        `orm:"id,primary"    json:"id"`           // 自增ID
	CreateAt     *gtime.Time `orm:"create_at"     json:"createAt"`     // 创建时间
	UpdateAt     *gtime.Time `orm:"update_at"     json:"updateAt"`     // 更新时间
	DeleteAt     *gtime.Time `orm:"delete_at"     json:"deleteAt"`     // 删除时间
	Label        string      `orm:"label"         json:"label"`        // 展示值
	Value        int         `orm:"value"         json:"value"`        // 字典值
	Status       int         `orm:"status"        json:"status"`       // 启用状态
	Sort         int         `orm:"sort"          json:"sort"`         // 排序标记
	DictionaryId int         `orm:"dictionary_id" json:"dictionaryId"` // 关联标记
}

// Files is the golang structure for table files.
type Files struct {
	Id       uint        `orm:"id,primary" json:"id"`       // 自增ID
	CreateAt *gtime.Time `orm:"create_at"  json:"createAt"` // 创建时间
	UpdateAt *gtime.Time `orm:"update_at"  json:"updateAt"` // 更新时间
	DeleteAt *gtime.Time `orm:"delete_at"  json:"deleteAt"` // 删除时间
	Name     string      `orm:"name"       json:"name"`     // 文件名
	Url      string      `orm:"url"        json:"url"`      // 文件地址
	Tag      string      `orm:"tag"        json:"tag"`      // 文件标签
	Key      string      `orm:"key"        json:"key"`      // 编号
}

// Jwts is the golang structure for table jwts.
type Jwts struct {
	Id       uint        `orm:"id,primary" json:"id"`       // 自增ID
	CreateAt *gtime.Time `orm:"create_at"  json:"createAt"` // 更新时间
	UpdateAt *gtime.Time `orm:"update_at"  json:"updateAt"` // 更新时间
	DeleteAt *gtime.Time `orm:"delete_at"  json:"deleteAt"` // 删除时间
	Jwt      string      `orm:"jwt"        json:"jwt"`      // jwt
}

// Menus is the golang structure for table menus.
type Menus struct {
	Id          uint        `orm:"id,primary"   json:"id"`          // 自增ID
	CreateAt    *gtime.Time `orm:"create_at"    json:"createAt"`    // 创建时间
	UpdateAt    *gtime.Time `orm:"update_at"    json:"updateAt"`    // 更新时间
	DeleteAt    *gtime.Time `orm:"delete_at"    json:"deleteAt"`    // 删除时间
	MenuLevel   uint        `orm:"menu_level"   json:"menuLevel"`   // 菜单等级(预留字段)
	ParentId    uint        `orm:"parent_id"    json:"parentId"`    // 父菜单ID
	Path        string      `orm:"path"         json:"path"`        // 路由path
	Name        string      `orm:"name"         json:"name"`        // 路由name
	Hidden      int         `orm:"hidden"       json:"hidden"`      // 是否在列表隐藏
	Component   string      `orm:"component"    json:"component"`   // 前端文件路径
	Title       string      `orm:"title"        json:"title"`       // 菜单名
	Icon        string      `orm:"icon"         json:"icon"`        // 菜单图标
	Sort        int         `orm:"sort"         json:"sort"`        // 排序标记
	KeepAlive   int         `orm:"keep_alive"   json:"keepAlive"`   // 是否缓存
	DefaultMenu int         `orm:"default_menu" json:"defaultMenu"` // 是否是基础路由(开发中)
}

// Operations is the golang structure for table operations.
type Operations struct {
	Id           uint        `orm:"id,primary"    json:"id"`           // 自增ID
	CreateAt     *gtime.Time `orm:"create_at"     json:"createAt"`     // 创建时间
	UpdateAt     *gtime.Time `orm:"update_at"     json:"updateAt"`     // 更新时间
	DeleteAt     *gtime.Time `orm:"delete_at"     json:"deleteAt"`     // 删除时间
	Ip           string      `orm:"ip"            json:"ip"`           // 请求ip
	Method       string      `orm:"method"        json:"method"`       // 请求方法
	Path         string      `orm:"path"          json:"path"`         // 请求路由
	Status       int         `orm:"status"        json:"status"`       // 状态
	Latency      int64       `orm:"latency"       json:"latency"`      // 延迟
	Agent        string      `orm:"agent"         json:"agent"`        // 代理
	ErrorMessage string      `orm:"error_message" json:"errorMessage"` // 报错信息
	Request      string      `orm:"request"       json:"request"`      // 请求Body
	UserId       int         `orm:"user_id"       json:"userId"`       // 用户id
	Response     string      `orm:"response"      json:"response"`     // 响应Body
}

// Parameters is the golang structure for table parameters.
type Parameters struct {
	Id         uint        `orm:"id,primary"   json:"id"`         // 自增ID
	CreateAt   *gtime.Time `orm:"create_at"    json:"createAt"`   // 创建时间
	UpdateAt   *gtime.Time `orm:"update_at"    json:"updateAt"`   // 更新时间
	DeleteAt   *gtime.Time `orm:"delete_at"    json:"deleteAt"`   // 删除时间
	BaseMenuId uint        `orm:"base_menu_id" json:"baseMenuId"` // BaseMenu的ID
	Type       string      `orm:"type"         json:"type"`       // 地址栏携带参数为params还是query
	Key        string      `orm:"key"          json:"key"`        // 地址栏携带参数的key
	Value      string      `orm:"value"        json:"value"`      // 地址栏携带参数的值
}

// Role is the golang structure for table role.
type Role struct {
	RoleId   uint        `orm:"role_id,primary" json:"roleId"`   // 角色ID
	RoleName string      `orm:"role_name"       json:"roleName"` // 角色名
	ParentId uint        `orm:"parent_id"       json:"parentId"` // 父角色ID
	CreateAt *gtime.Time `orm:"create_at"       json:"createAt"` // 创建时间
	UpdateAt *gtime.Time `orm:"update_at"       json:"updateAt"` // 更新时间
	DeleteAt *gtime.Time `orm:"delete_at"       json:"deleteAt"` // 删除时间
}

// RoleMenu is the golang structure for table role_menu.
type RoleMenu struct {
	RoleId uint `orm:"role_id" json:"roleId"` // 权限id
	MenuId uint `orm:"menu_id" json:"menuId"` // 菜单id
}

// RoleResources is the golang structure for table role_resources.
type RoleResources struct {
	RoleId      uint `orm:"role_id"      json:"roleId"`      // 权限id
	ResourcesId uint `orm:"resources_id" json:"resourcesId"` // 资源id
}
