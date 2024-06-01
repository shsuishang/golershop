package model

import "golershop.cn/internal/model/entity"

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
