package pt

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type ProductSpecAdd struct {
	SpecName   string `json:"spec_name"   v:"required#请输入规格名称"        `     // 规格名称
	SpecRemark string `json:"spec_remark"                                 ` // 规格注释
	SpecFormat string `json:"spec_format" v:"required#请选择显示类型"        `     // 显示类型(ENUM): text-文字; image-图片
	SpecSort   uint   `json:"spec_sort"                                   ` // 排序
	CategoryId uint   `json:"category_id" v:"required#请备注分类"           `    // 规格分类编号:不是商品类型编号，选择分类，可关联到任意级分类。（可以使用一级分类category_id，只在后台快捷定位中起作用）
}
type ProductSpecEditReq struct {
	g.Meta `path:"/manage/pt/productSpec/edit" tags:"规格管理" method:"post" summary:"规格编辑接口"`

	SpecId uint `json:"spec_id"   ` // 规格编号`
	ProductSpecAdd
}

type ProductSpecEditRes struct {
	SpecId interface{} `json:"spec_id"   dc:"规格信息"`
}

type ProductSpecAddReq struct {
	g.Meta `path:"/manage/pt/productSpec/add" tags:"规格管理" method:"post" summary:"规格编辑接口"`

	ProductSpecAdd
}

type ProductSpecRemoveReq struct {
	g.Meta `path:"/manage/pt/productSpec/remove" tags:"规格管理" method:"post" summary:"规格删除接口"`
	SpecId uint `json:"spec_id"   ` // 规格编号
}

type ProductSpecRemoveRes struct {
}

type ProductSpecListReq struct {
	g.Meta `path:"/manage/pt/productSpec/list" tags:"规格管理" method:"get" summary:"规格列表接口"`
	ml.BaseList

	SpecId      uint   `json:"spec_id"          ` // 规格编号
	SpecName    string `json:"spec_name"        ` // 规格名称
	SpecRemark  string `json:"spec_remark"      ` // 规格注释
	SpecFormat  string `json:"spec_format"      ` // 显示类型(ENUM): text-文字; image-图片
	SpecSort    uint   `json:"spec_sort"        ` // 排序
	CategoryId  uint   `json:"category_id" `      // 规格分类编号:不是商品类型编号，选择分类，可关联到任意级分类。（可以使用一级分类category_id，只在后台快捷定位中起作用）
	SpecBuildin uint   `json:"spec_buildin"     ` // 系统内置(BOOL):1-是; 0-否
}

type ProductSpecListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ProductSpecTreeReq struct {
	g.Meta `path:"/manage/pt/productSpec/tree" tags:"商品规格" method:"get" summary:"品牌配置tree"`
}

type ProductSpecTreeVoRes struct {
	SpecId   uint        `json:"spec_id"        dc:"主键编号"   `
	SpecName string      `json:"spec_name"      dc:"分组名称" `
	Children interface{} `json:"children"    dc:"配置列表页"`
}

//type ProductSpecTreeRes struct {
//	Items []ProductSpecTreeVo `json:"Items"    dc:"商品规格数据"`
//}
