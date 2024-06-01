package model

import "golershop.cn/internal/model/entity"

// 无限级分类Vo
type ArticleCategoryTreeNode struct {
	entity.ArticleCategory
	Children []*ArticleCategoryTreeNode `json:"children"` // 子菜单
}
