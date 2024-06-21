package account

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// =========================== 管理端使用 =============================

type UserTagGroupAdd struct {
	TagGroupId      uint        `json:"tag_group_id"      ` // 分组编号
	TagGroupName    string      `json:"tag_group_name"    ` // 分组名称
	TagGroupSort    uint        `json:"tag_group_sort"    ` // 分组排序:从小到大
	TagGroupEnable  bool        `json:"tag_group_enable"  ` // 是否有效(BOOL):0-禁用;1-启用
	TagGroupBuildin bool        `json:"tag_group_buildin" ` // 系统内置(BOOL):1-是; 0-否
	UpdateTime      *gtime.Time `json:"update_time"       ` // 更新时间
	CreateTime      *gtime.Time `json:"create_time"       ` // 添加时间
}

type UserTagGroupAddReq struct {
	g.Meta `path:"/manage/account/userTagGroup/add" tags:"用户标签" method:"post" summary:"用户标签详情接口"`

	UserTagGroupAdd
}

type UserTagGroupEditReq struct {
	g.Meta     `path:"/manage/account/userTagGroup/edit" tags:"用户标签" method:"post" summary:"用户标签详情接口"`
	TagGroupId uint `json:"tag_group_id" v:"required#请输入用户标签编号"   dc:"用户标签信息"`
	UserTagGroupAdd
}

type UserTagGroupAddRes struct {
	TagGroupId interface{} `json:"tag_group_id"   dc:"主键编号"   `
}

type UserTagGroupEditRes struct {
	TagGroupId interface{} `json:"tag_group_id"   dc:"主键编号"   `
}

type UserTagGroupRemoveReq struct {
	g.Meta     `path:"/manage/account/userTagGroup/remove" tags:"用户标签" method:"post" summary:"用户标签删除接口"`
	TagGroupId []uint `json:"tag_group_id" v:"required#请输入用户标签编号"   dc:"用户标签信息"`
}

type UserTagGroupRemoveRes struct {
}

type UserTagGroupListReq struct {
	g.Meta `path:"/manage/account/userTagGroup/list" tags:"用户标签" method:"get" summary:"用户标签列表接口"`

	Page int `json:"page"  d:"1"  v:"min:0#分页号码错误"  dc:"分页号码"`
	Size int `json:"size" d:"10" v:"max:500#分页数量最大500条"  dc:"分页数量"`
}

type UserTagGroupListRes struct {
	Items   interface{} `json:"items"    dc:"用户标签列表页"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type UserTagGroupTreeReq struct {
	g.Meta `path:"/manage/account/userTagGroup/tree" tags:"标签分组表-树形集合" method:"get" summary:"标签分组表-树形集合接口"`
	ml.BaseList

	TagGroupName   string `json:"tag_group_name"  dc:"分组名称"`                  // 分组名称
	TagGroupEnable bool   `json:"tag_group_enable" dc:"是否有效(BOOL):0-禁用;1-启用"` // 是否有效(BOOL):0-禁用;1-启用
}

type UserTagGroupTreeRes struct {
	TagTitle string                `json:"tag_title" dc:"分组名称"` // 分组名称
	Children []*entity.UserTagBase `json:"children"  dc:"子集标签"` // 子集标签
}

// ---------------------------- 标签项 -------------------------------

type UserTagBaseAdd struct {
	TagId      string `json:"tag_id"       ` // 标签编码
	TagTitle   string `json:"tag_title"    ` // 标签标题
	TagGroupId uint   `json:"tag_group_id" ` // 所属分类
	TagSort    uint   `json:"tag_sort"     ` // 标签排序:从小到大
	TagEnable  bool   `json:"tag_enable"   ` // 是否启用(BOOL):0-禁用;1-启用
	TagBuildin bool   `json:"tag_buildin"  ` // 系统内置(BOOL):1-是; 0-否
}

type UserTagBaseAddReq struct {
	g.Meta `path:"/manage/account/userTagBase/add" tags:"标签" method:"post" summary:"标签详情接口"`

	UserTagBaseAdd
}

type UserTagBaseEditReq struct {
	g.Meta `path:"/manage/account/userTagBase/edit" tags:"标签" method:"post" summary:"标签详情接口"`

	TagId int64 `json:"tag_id"       ` // 标签编码
	UserTagBaseAdd
}

type UserTagBaseAddRes struct {
	TagId int64 `json:"tag_id"       ` // 标签编码
}

type UserTagBaseEditRes struct {
	TagId int64 `json:"tag_id"       ` // 标签编码
}

type UserTagBaseRemoveReq struct {
	g.Meta `path:"/manage/account/userTagBase/remove" tags:"标签" method:"post" summary:"标签删除接口"`
	TagId  []int64 `json:"tag_id"       ` // 标签编码
}

type UserTagBaseRemoveRes struct {
}

type UserTagBaseListReq struct {
	g.Meta     `path:"/manage/account/userTagBase/list" tags:"标签" method:"get" summary:"标签列表接口"`
	TagGroupId uint   `json:"tag_group_id" ` // 所属分类
	TagTitle   string `json:"tag_title"    ` // 标签标题

	Page int `json:"page"  d:"1"  v:"min:0#分页号码错误"  dc:"分页号码"`
	Size int `json:"size" d:"10" v:"max:500#分页数量最大500条"  dc:"分页数量"`
}

type UserTagBaseListRes struct {
	Items   interface{} `json:"items"    dc:"素材列表页"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
