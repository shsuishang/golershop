package pt

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/model"
)

// start fo manage
type ProductCategoryAdd struct {
	CategoryParentId       uint    `json:"category_parent_id"       ` // 分类父编号
	CategoryName           string  `json:"category_name"         `    // 分类名称
	CategoryImage          string  `json:"category_image"           ` // 分类图片
	TypeId                 uint    `json:"type_id"                  ` // 类型编号
	CategoryCommissionRate float64 `json:"category_commission_rate" ` // 分佣比例
	CategorySort           uint    `json:"category_sort"            ` // 排序
	CategoryIsEnable       bool    `json:"category_is_enable"       ` // 是否启用(BOOL):0-不显示;1-显示
}
type ProductCategoryEditReq struct {
	g.Meta `path:"/manage/pt/productCategory/edit" tags:"商品分类" method:"post" summary:"分类编辑接口"`

	CategoryId uint `json:"category_id"              ` // 分类编号
	ProductCategoryAdd
}

type ProductCategoryEditRes struct {
	CategoryId interface{} `json:"category_id"   dc:"商品分类信息"`
}

type ProductCategoryEditStateReq struct {
	g.Meta `path:"/manage/pt/productCategory/editState" tags:"商品分类" method:"post" summary:"分类编辑接口"`

	CategoryId       uint `json:"category_id"              ` // 分类编号
	CategoryIsEnable bool `json:"category_is_enable"       ` // 是否启用(BOOL):0-不显示;1-显示
}

type ProductCategoryEditStateRes struct {
	CategoryId interface{} `json:"category_id"   dc:"商品分类信息"`
}

type ProductCategoryAddReq struct {
	g.Meta `path:"/manage/pt/productCategory/add" tags:"商品分类" method:"post" summary:"商品分类添加接口"`

	ProductCategoryAdd
}

type ProductCategoryRemoveReq struct {
	g.Meta     `path:"/manage/pt/productCategory/remove" tags:"商品分类" method:"post" summary:"商品分类删除接口"`
	CategoryId []uint `json:"category_id" v:"required#请输入商品分类编号"   dc:"商品分类信息"`
}

type ProductCategoryRemoveRes struct {
}

type ProductCategoryListReq struct {
	g.Meta `path:"/manage/pt/productCategory/list" tags:"商品分类" method:"get" summary:"商品分类列表接口"`
	Page   int `json:"page"  d:"1"  v:"min:0#分页号码错误"  dc:"分页号码"`
	Size   int `json:"size" d:"10" v:"max:500#分页数量最大500条"  dc:"分页数量"`

	CategoryParentId uint   `json:"category_parent_id"  d:"0"  v:"min:0#分页号码错误"  dc:"分类父编号"`
	CategoryName     string `json:"category_name"  dc:"搜索关键词"`
	CategoryIsEnable bool   `json:"category_is_enable"  dc:"是否启用"`
}

type ProductCategoryListRes struct {
	Items   interface{} `json:"items"    dc:"分类列表"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ProductCategoryTreeReq struct {
	g.Meta       `path:"/manage/pt/productCategory/tree" tags:"商品分类" method:"get" summary:"后台商品分类Tree"`
	CategoryName string `json:"category_name"  dc:"搜索关键词"`
}

//type ProductCategoryTreeRes model.TreeNode

//res []*v1.ProductCategoryTreeRes,

/*
type ProductCategoryTreeRes struct {
	model.ProductCategory
	Children []*model.TreeNode `json:"children"` // 子商品分类
}
*/

type ProductCategoryTreeRes []*model.CategoryTreeNode
