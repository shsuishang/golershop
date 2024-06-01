package pt

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type ProductTypeAdd struct {
	TypeName    string `json:"type_name"  v:"required#请输入类型名称"      ` // 类型名称
	TypeRemark  string `json:"type_remark"      `                     // 备注
	CategoryId  uint   `json:"category_id"   `                        // 分类编号-快捷定位功能
	SpecIds     string `json:"spec_ids"         `                     // 规格编号(DOT)
	BrandIds    string `json:"brand_ids"        `                     // 品牌编号(DOT)
	AssistIds   string `json:"assist_ids"       `                     // 辅助属性(DOT)
	TypeIsDraft bool   `json:"type_is_draft"    `                     // 是否草稿(ENUM):1-草稿;0-发布
}
type ProductTypeEditReq struct {
	g.Meta `path:"/manage/pt/productType/edit" tags:"商品类型" method:"post" summary:"类型编辑接口"`

	TypeId uint `json:"type_id"   v:"required#请输入类型编号"    dc:"类型编号"     `
	ProductTypeAdd
}

type ProductTypeEditRes struct {
	TypeId interface{} `json:"type_id"   dc:"类型信息"`
}

type ProductTypeAddReq struct {
	g.Meta `path:"/manage/pt/productType/add" tags:"商品类型" method:"post" summary:"类型编辑接口"`

	ProductTypeAdd
}

type ProductTypeRemoveReq struct {
	g.Meta `path:"/manage/pt/productType/remove" tags:"商品类型" method:"post" summary:"类型删除接口"`
	TypeId []uint `json:"type_id" v:"required#请输入类型编号"   dc:"类型信息"`
}

type ProductTypeRemoveRes struct {
}

type ProductTypeListReq struct {
	g.Meta `path:"/manage/pt/productType/list" tags:"商品类型" method:"get" summary:"类型列表接口"`
	ml.BaseList

	TypeId     uint   `json:"type_id"          ` // 编号
	TypeName   string `json:"type_name"        ` // 类型名称
	TypeRemark string `json:"type_remark"      ` // 备注
}

type ProductTypeListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ProductTypeInfoReq struct {
	g.Meta `path:"/manage/pt/productType/info" tags:"商品类型" method:"get" summary:"类型信息接口"`
	TypeId uint `json:"type_id"          ` // 编号
}

type ProductTypeInfoRes struct {
	Brands  interface{} `json:"brands"    `  // 品牌列表
	Assists interface{} `json:"assists"    ` // 属性列表
	Specs   interface{} `json:"specs"    `   // 规格列表
}
