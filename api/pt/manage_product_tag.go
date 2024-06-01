package pt

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type ProductTagAdd struct {
	ProductTagName string      `json:"product_tag_name" v:"required#请输入标签名称" ` // 标签名称
	StoreId        uint        `json:"store_id"         `                      // 店铺编号
	CreateTime     *gtime.Time `json:"create_time"      `                      // 创建时间
}
type ProductTagEditReq struct {
	g.Meta `path:"/manage/pt/productTag/edit" tags:"标签管理" method:"post" summary:"标签编辑接口"`

	ProductTagId uint `json:"product_tag_id"   ` // 标签编号`
	ProductTagAdd
}

type ProductTagEditRes struct {
	ProductTagId interface{} `json:"product_tag_id"   dc:"标签信息"`
}

type ProductTagAddReq struct {
	g.Meta `path:"/manage/pt/productTag/add" tags:"标签管理" method:"post" summary:"标签编辑接口"`

	ProductTagAdd
}

type ProductTagRemoveReq struct {
	g.Meta       `path:"/manage/pt/productTag/remove" tags:"标签管理" method:"post" summary:"标签删除接口"`
	ProductTagId uint `json:"product_tag_id"   ` // 标签编号
}

type ProductTagRemoveRes struct {
}

type ProductTagListReq struct {
	g.Meta `path:"/manage/pt/productTag/list" tags:"标签管理" method:"get" summary:"标签列表接口"`
	ml.BaseList

	ProductTagId   uint        `json:"product_tag_id"   `             // 标签编号
	ProductTagName string      `json:"product_tag_name" type:"LIKE" ` // 标签名称
	StoreId        uint        `json:"store_id"         `             // 店铺编号
	CreateTime     *gtime.Time `json:"create_time"      `             // 创建时间
}

type ProductTagListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
