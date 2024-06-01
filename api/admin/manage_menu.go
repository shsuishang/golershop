package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/entity"
)

// start fo front

// start fo manage
type MenuAdd struct {
	MenuParentId  uint   `json:"menu_parent_id" v:"required#请输入菜单父编号"  dc:"菜单父编号"`
	MenuTitle     string `json:"menu_title"   v:"required#请输入菜单名称"  dc:"菜单名称"    `
	MenuName      string `json:"menu_name"       `                                // 组件名称
	MenuPath      string `json:"menu_path"       `                                // 后端路由
	MenuComponent string `json:"menu_component" d:"Layout" v:"required#请输入组件路径" ` // 组件路径

	MenuClose      bool   `json:"menu_close"      ` // 允许关闭(BOOL):0-禁止;1-允许
	MenuHidden     bool   `json:"menu_hidden"     ` // 是否隐藏(BOOL):0-展示;1-隐藏
	MenuIcon       string `json:"menu_icon"       ` // 图标设置
	MenuDot        bool   `json:"menu_dot"        ` // 是否红点(BOOL):0-隐藏;1-显示
	MenuBubble     string `json:"menu_bubble"     ` // 菜单标签
	MenuSort       int    `json:"menu_sort"       ` // 菜单排序
	MenuType       int    `json:"menu_type"       ` // 菜单类型(LIST):0-按钮;1-菜单
	MenuNote       string `json:"menu_note"       ` // 备注
	MenuFunc       string `json:"menu_func"       ` // 功能开启:设置config_key
	MenuPermission string `json:"menu_permission" ` // 权限标识:后端地址
}
type MenuEditReq struct {
	g.Meta `path:"/manage/admin/menu/edit" tags:"菜单" method:"post" summary:"菜单详情接口"`

	MenuId uint `json:"menu_id"   v:"required#请输入菜单编号"    dc:"菜单编号"     `
	MenuAdd
}

type MenuEditRes struct {
	MenuId interface{} `json:"menu_id"   dc:"菜单信息"`
}

type MenuAddReq struct {
	g.Meta `path:"/manage/admin/menu/add" tags:"菜单" method:"post" summary:"菜单详情接口"`

	MenuAdd
}

type MenuEditStateReq struct {
	g.Meta `path:"/manage/admin/menu/editState" tags:"菜单" method:"post" summary:"菜单状态编辑接口"`

	MenuId     uint `json:"menu_id"   v:"required#请输入菜单编号"    dc:"菜单编号"     `
	MenuClose  bool `json:"menu_close"      ` // 允许关闭(BOOL):0-禁止;1-允许
	MenuHidden bool `json:"menu_hidden"     ` // 是否隐藏(BOOL):0-展示;1-隐藏
	MenuDot    bool `json:"menu_dot"        ` // 是否红点(BOOL):0-隐藏;1-显示
}

type MenuEditStateRes struct {
	MenuId interface{} `json:"menu_id"   dc:"菜单编号"`
}

type MenuDetailReq struct {
	g.Meta `path:"/manage/admin/menu/detail" tags:"菜单" method:"get" summary:"菜单详情接口"`
	MenuId uint `json:"menu_id" v:"required#请输入菜单编号"   dc:"菜单信息"`
}

type MenuDetailRes struct {
	*entity.MenuBase
}

type MenuRemoveReq struct {
	g.Meta `path:"/manage/admin/menu/remove" tags:"菜单" method:"post" summary:"菜单删除接口"`
	MenuId []uint `json:"menu_id" v:"required#请输入菜单编号"   dc:"菜单信息"`
}

type MenuRemoveRes struct {
}

type MenuListReq struct {
	g.Meta `path:"/manage/admin/menu/list" tags:"菜单" method:"get" summary:"菜单列表接口"`
	Page   int `json:"page"  d:"1"  v:"min:0#分页号码错误"  dc:"分页号码"`
	Size   int `json:"size" d:"10" v:"max:500#分页数量最大500条"  dc:"分页数量"`
}

type MenuListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type MenuTreeReq struct {
	g.Meta `path:"/manage/admin/menu/tree" tags:"菜单" method:"get" summary:"后台菜单Tree"`
	Type   int    `json:"type"  d:"1"  dc:"0-按钮;1-菜单;2-所有"`
	Title  string `json:"menu_title"  d:""  dc:"搜索关键词"`
}

//type MenuTreeRes model.TreeNode

//res []*v1.MenuTreeRes,

/*
type MenuTreeRes struct {
	model.Menu
	Children []*model.TreeNode `json:"children"` // 子菜单
}
*/

type MenuTreeRes []*model.TreeNode

//res v1.MenuTreeRes,
