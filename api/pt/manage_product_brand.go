package pt

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type ProductBrandAdd struct {
	BrandName      string `json:"brand_name" v:"required#请输入品牌名称"      ` // 品牌名称
	BrandDesc      string `json:"brand_desc" v:"required#请输入品牌描述"      ` // 品牌描述
	CategoryId     uint   `json:"category_id" v:"required#请选择所属分类"     ` // 所属分类:一级分类即可
	BrandShowType  uint   `json:"brand_show_type" `                      // 展示方式(ENUM):1-图片; 2-文字  | 在“全部品牌”页面的展示方式，如果设置为“图片”则显示该品牌的“品牌图片标识”，如果设置为“文字”则显示该品牌的“品牌名”
	BrandImage     string `json:"brand_image"     `                      // 品牌LOGO
	BrandRecommend bool   `json:"brand_recommend" `                      // 是否推荐(BOOL):1-是; 0-否
	BrandEnable    bool   `json:"brand_enable"    `                      // 是否启用(BOOL):1-启用; 0-禁用
	StoreId        uint   `json:"store_id"        `                      // 店铺编号
	BrandApply     uint   `json:"brand_apply"     `                      // 品牌申请(ENUM):0-申请中; 1-通过 | 申请功能是会员使用，系统后台默认为1
}
type ProductBrandEditReq struct {
	g.Meta `path:"/manage/pt/productBrand/edit" tags:"商品品牌" method:"post" summary:"品牌编辑接口"`

	BrandId uint `json:"brand_id"   v:"required#请输入品牌编号"    dc:"品牌编号"     `
	ProductBrandAdd
}

type ProductBrandEditRes struct {
	BrandId interface{} `json:"brand_id"   dc:"品牌信息"`
}

type ProductBrandEditStateReq struct {
	g.Meta `path:"/manage/pt/productBrand/editState" tags:"商品品牌" method:"post" summary:"商品品牌状态编辑接口"`

	BrandId        uint `json:"brand_id"   v:"required#请输入品牌编号"    dc:"品牌编号"     `
	BrandRecommend bool `json:"brand_recommend" `
	BrandEnable    bool `json:"brand_enable"    `
}

type ProductBrandEditStateRes struct {
	CategoryId interface{} `json:"category_id"   dc:"商品分类信息"`
}

type ProductBrandAddReq struct {
	g.Meta `path:"/manage/pt/productBrand/add" tags:"商品品牌" method:"post" summary:"品牌编辑接口"`

	ProductBrandAdd
}

type ProductBrandRemoveReq struct {
	g.Meta  `path:"/manage/pt/productBrand/remove" tags:"商品品牌" method:"post" summary:"品牌删除接口"`
	BrandId []uint `json:"brand_id" v:"required#请输入品牌编号"   dc:"品牌信息"`
}

type ProductBrandRemoveRes struct {
}

type ProductBrandListReq struct {
	g.Meta `path:"/manage/pt/productBrand/list" tags:"商品品牌" method:"get" summary:"品牌列表接口"`
	ml.BaseList

	BrandName    string `json:"brand_name"  type:"LIKE"    ` // 品牌名称
	BrandCode    string `json:"brand_code"      `            // 品牌拼音
	BrandInitial string `json:"brand_initial"   `            // 首字母
}

type ProductBrandListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ProductBrandTreeReq struct {
	g.Meta `path:"/manage/pt/productBrand/tree" tags:"商品品牌" method:"get" summary:"品牌配置tree"`
}

type ProductBrandTreeVoRes struct {
	BrandId   uint        `json:"brand_id"        dc:"主键编号"   `
	BrandName string      `json:"brand_name"      dc:"分组名称" `
	Children  interface{} `json:"children"    dc:"配置列表页"`
}
