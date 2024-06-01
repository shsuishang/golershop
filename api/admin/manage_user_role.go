package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type UserRoleAdd struct {
	UserRoleName string `json:"user_role_name"  ` // 角色名称
	UserRoleCode string `json:"user_role_code"  ` // 角色标识
	MenuIds      string `json:"menu_ids"        ` // 请求列表(DOT)
}

type UserRoleAddReq struct {
	g.Meta `path:"/manage/admin/userRole/add" tags:"角色" method:"post" summary:"角色编辑接口"`

	UserRoleAdd
}

type UserRoleEditReq struct {
	g.Meta `path:"/manage/admin/userRole/edit" tags:"角色" method:"post" summary:"角色编辑接口"`

	UserRoleId uint `json:"user_role_id" v:"required#请输入角色编号"    dc:"角色编号"   ` // 角色编号
	UserRoleAdd
}

type UserRoleEditRes struct {
	UserRoleId uint `json:"user_role_id" v:"required#请输入角色编号"    dc:"角色编号"   ` // 角色编号
}

type UserRoleRemoveReq struct {
	g.Meta     `path:"/manage/admin/userRole/remove" tags:"角色" method:"post" summary:"角色删除接口"`
	UserRoleId uint `json:"user_role_id" v:"required#请输入角色编号"    dc:"角色编号"   ` // 角色编号
}

type UserRoleRemoveRes struct {
}

type UserRoleListReq struct {
	g.Meta `path:"/manage/admin/userRole/list" tags:"角色" method:"get" summary:"角色列表接口"`
	ml.BaseList

	UserRoleId    uint        `json:"user_role_id"    dc:"角色编号"   ` // 角色编号
	UserRoleName  string      `json:"user_role_name"  `             // 角色名称
	UserRoleCode  string      `json:"user_role_code"  `             // 角色标识
	MenuIds       string      `json:"menu_ids"        `             // 请求列表(DOT)
	UserRoleCtime *gtime.Time `json:"user_role_ctime" `             // 创建时间
	UserRoleUtime *gtime.Time `json:"user_role_utime" `             // 更新时间
}

type UserRoleListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
