package pt

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type ProductSpecItemAdd struct {
	SpecItemId     uint   `json:"spec_item_id"     ` // 商品规格值编号
	StoreId        uint   `json:"store_id"         ` // 店铺编号
	CategoryId     uint   `json:"category_id"      ` // 商品分类编号,通过type决定规格，但是分类下的规格值都不同
	SpecId         uint   `json:"spec_id"          ` // 规格分类编号
	SpecItemName   string `json:"spec_item_name"   ` // 规格值名称
	SpecItemSort   uint   `json:"spec_item_sort"   ` // 排序:越小越靠前
	SpecItemEnable bool   `json:"spec_item_enable" ` // 是否启用(BOOL):0-不显示;1-显示
}
type ProductSpecItemEditReq struct {
	g.Meta `path:"/manage/pt/productSpecItem/edit" tags:"规格项目管理" method:"post" summary:"规格项目编辑接口"`

	SpecItemId uint `json:"product_tag_id"   ` // 规格项目编号`
	ProductSpecItemAdd
}

type ProductSpecItemEditRes struct {
	SpecItemId interface{} `json:"product_tag_id"   dc:"规格项目信息"`
}

type ProductSpecItemAddReq struct {
	g.Meta `path:"/manage/pt/productSpecItem/add" tags:"规格项目管理" method:"post" summary:"规格项目编辑接口"`

	ProductSpecItemAdd
}

type ProductSpecItemRemoveReq struct {
	g.Meta     `path:"/manage/pt/productSpecItem/remove" tags:"规格项目管理" method:"post" summary:"规格项目删除接口"`
	SpecItemId uint `json:"product_tag_id"   ` // 规格项目编号
}

type ProductSpecItemRemoveRes struct {
}

type ProductSpecItemListReq struct {
	g.Meta `path:"/manage/pt/productSpecItem/list" tags:"规格项目管理" method:"get" summary:"规格项目列表接口"`
	ml.BaseList

	SpecItemId     uint   `json:"spec_item_id"     ` // 商品规格值编号
	StoreId        uint   `json:"store_id"         ` // 店铺编号
	CategoryId     uint   `json:"category_id"      ` // 商品分类编号,通过type决定规格，但是分类下的规格值都不同
	SpecId         uint   `json:"spec_id"          ` // 规格分类编号
	SpecItemName   string `json:"spec_item_name"   ` // 规格值名称
	SpecItemSort   uint   `json:"spec_item_sort"   ` // 排序:越小越靠前
	SpecItemEnable uint   `json:"spec_item_enable" ` // 是否启用(BOOL):0-不显示;1-显示
}

type ProductSpecItemListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ProductSpecItemEditStateReq struct {
	g.Meta `path:"/manage/pt/productSpecItem/editState" tags:"规格项目管理" method:"post" summary:"规格项目编辑接口"`

	SpecItemId     uint `json:"spec_item_id"     ` // 商品规格值编号
	SpecItemEnable bool `json:"spec_item_enable" ` // 是否启用(BOOL):0-不显示;1-显示
}

type ProductSpecItemEditStateRes struct {
	SpecItemId uint `json:"spec_item_id"     ` // 商品规格值编号
}
