package cms

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type ArticleTagAdd struct {
	TagId    uint   `json:"tag_id"    ` // 标签编号
	TagName  string `json:"tag_name"  ` // 标签名称
	TagCount uint   `json:"tag_count" ` // 内容数量
	UserId   uint   `json:"user_id"   ` // 用户编号
	Version  uint   `json:"version"   ` // 版本
}

type ArticleTagAddReq struct {
	g.Meta `path:"/manage/cms/articleTag/add" tags:"标签管理" method:"post" summary:"标签管理编辑接口"`

	ArticleTagAdd
}

type ArticleTagEditReq struct {
	g.Meta `path:"/manage/cms/articleTag/edit" tags:"标签管理" method:"post" summary:"标签管理编辑接口"`

	TagId uint `json:"tag_id"  dc:"标签管理编号"   ` // 标签管理编号
	ArticleTagAdd
}

type ArticleTagEditRes struct {
	TagId uint `json:"tag_id" dc:"标签管理编号"   ` // 标签管理编号
}

type ArticleTagRemoveReq struct {
	g.Meta `path:"/manage/cms/articleTag/remove" tags:"标签管理" method:"post" summary:"标签管理删除接口"`
	TagId  uint `json:"tag_id" dc:"标签管理编号"   ` // 标签管理编号
}

type ArticleTagRemoveRes struct {
}

type ArticleTagListReq struct {
	g.Meta `path:"/manage/cms/articleTag/list" tags:"标签管理" method:"get" summary:"标签管理列表接口"`
	ml.BaseList

	TagId uint `json:"tag_id"                   ` // 用户编号
}

type ArticleTagListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
