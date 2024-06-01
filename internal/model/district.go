package model

// 无限级分类Vo
type DistrictTreeNode struct {
	DistrictId       uint   `json:"district_id"        ` // 地区编号
	DistrictName     string `json:"district_name"      ` // 地区名称
	DistrictParentId uint   `json:"district_parent_id" ` // 上级编号
	DistrictSort     uint   `json:"district_sort"      ` // 地区排序

	Children []*DistrictTreeNode `json:"children"` // 子菜单
}
