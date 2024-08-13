package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type UserAdminAdd struct {
	UserId           uint        `json:"user_id"            ` // 用户编号
	UserRoleId       uint        `json:"user_role_id"       ` // 权限系统用户
	UserAdminCtime   *gtime.Time `json:"user_admin_ctime"   ` // 创建时间
	UserAdminUtime   *gtime.Time `json:"user_admin_utime"   ` // 更新时间
	UserIsSuperadmin bool        `json:"user_is_superadmin" ` // 是否超管
	RoleId           uint        `json:"role_id"            ` // 系统用户编号(ENUM):0-用户;2-商家;3-门店;9-平台;
	ChainId          uint        `json:"chain_id"           ` // 所属门店
	StoreId          uint        `json:"store_id"           ` // 店铺编号
}

type UserAdminAddReq struct {
	g.Meta `path:"/manage/admin/userAdmin/add" tags:"系统用户" method:"post" summary:"系统用户编辑接口"`

	UserAdminAdd
}

type UserAdminEditReq struct {
	g.Meta `path:"/manage/admin/userAdmin/edit" tags:"系统用户" method:"post" summary:"系统用户编辑接口"`

	UserId uint `json:"user_id"            ` // 用户编号
	UserAdminAdd
}

type UserAdminEditRes struct {
	UserId uint `json:"user_id"            ` // 用户编号
}

type UserAdminRemoveReq struct {
	g.Meta `path:"/manage/admin/userAdmin/remove" tags:"系统用户" method:"post" summary:"系统用户删除接口"`
	UserId uint `json:"user_id"            ` // 用户编号
}

type UserAdminRemoveRes struct {
}

type UserAdminListReq struct {
	g.Meta `path:"/manage/admin/userAdmin/list" tags:"系统用户" method:"get" summary:"系统用户列表接口"`
	ml.BaseList

	UserId     uint `json:"user_id"            ` // 用户编号
	UserRoleId uint `json:"user_role_id"       ` // 权限系统用户
}

type UserAdminListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
