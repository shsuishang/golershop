package model

import "golershop.cn/internal/model/entity"

// 无限级分类Vo
type CategoryTreeNode struct {
	entity.ProductCategory
	Children []*CategoryTreeNode `json:"children"` // 子菜单
}
