package model

import "golershop.cn/internal/model/entity"

type ProductAssistTreeVo struct {
	AssistId   uint                    `json:"assist_id" dc:"属性分类编号"` // 属性分类编号
	AssistName string                  `json:"assist_name" dc:"属性名称"` // 属性名称
	Children   []*entity.ProductAssist `json:"children" dc:"属性数据"`    // 属性数据
}

type AssistVo struct {
	entity.ProductAssist
	Items []*entity.ProductAssistItem `json:"items"    ` // 属性选项列表
}

type SpecVo struct {
	entity.ProductSpec
	Items []*entity.ProductSpecItem `json:"items"    ` // 规格选项列表
}

type ProductTypeInfoOutput struct {
	Brands  []*entity.ProductBrand `json:"brands"    `  // 品牌列表
	Assists []*AssistVo            `json:"assists"    ` // 属性列表
	Specs   []*SpecVo              `json:"specs"    `   // 规格列表
}
