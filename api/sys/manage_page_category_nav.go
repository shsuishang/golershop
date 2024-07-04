package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type PageCategoryNavAdd struct {
	CategoryNavId     uint   `json:"category_nav_id"     ` // 编号
	CategoryNavName   string `json:"category_nav_name"   ` // 分类名称
	CategoryNavImage  string `json:"category_nav_image"  ` // 分类图片
	CategoryIds       string `json:"category_ids"        ` // 推荐分类(DOT)
	ItemIds           string `json:"item_ids"            ` // 推荐商品(DOT)
	BrandIds          string `json:"brand_ids"           ` // 推荐品牌(DOT)
	CategoryNavAdv    string `json:"category_nav_adv"    ` // 广告数据(JSON)
	CategoryNavType   uint   `json:"category_nav_type"   ` // 模板分类(ENUM):1-分类模板1;2-商品模板
	CategoryNavOrder  uint   `json:"category_nav_order"  ` // 排序
	CategoryNavEnable bool   `json:"category_nav_enable" ` // 是否启用(BOOL):0-不显示;1-显示
}

type PageCategoryNavEditReq struct {
	g.Meta `path:"/manage/sys/pageCategoryNav/edit" tags:"PC分类导航表" method:"post" summary:"PC分类导航表编辑接口"`

	CategoryNavId uint `json:"category_nav_id"     ` // 编号
	PageCategoryNavAdd
}

type PageCategoryNavEditRes struct {
	CategoryNavId uint `json:"category_nav_id"     ` // 编号
}

type PageCategoryNavAddReq struct {
	g.Meta `path:"/manage/sys/pageCategoryNav/add" tags:"PC分类导航表" method:"post" summary:"PC分类导航表编辑接口"`

	PageCategoryNavAdd
}

type PageCategoryNavRemoveReq struct {
	g.Meta `path:"/manage/sys/pageCategoryNav/remove" tags:"PC分类导航表" method:"post" summary:"PC分类导航表删除接口"`

	CategoryNavId uint `json:"category_nav_id"     ` // 编号
}

type PageCategoryNavRemoveRes struct {
}

type PageCategoryNavListReq struct {
	g.Meta `path:"/manage/sys/pageCategoryNav/list" tags:"PC分类导航表" method:"get" summary:"PC分类导航表列表接口"`
	ml.BaseList

	CategoryNavName string `json:"category_nav_name"   ` // 分类名称
}

type PageCategoryNavListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type PageCategoryNavEditStateReq struct {
	g.Meta `path:"/manage/sys/pageCategoryNav/editState" tags:"PC分类导航表" method:"post" summary:"PC分类导航表状态编辑接口"`

	CategoryNavId     uint `json:"category_nav_id"     ` // 编号
	CategoryNavEnable bool `json:"category_nav_enable" ` // 是否启用(BOOL):0-不显示;1-显示
}

type PageCategoryNavEditStateRes struct {
	CategoryNavId int64 `json:"category_nav_id"     ` // 编号
}
