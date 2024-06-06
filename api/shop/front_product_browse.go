package shop

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
	"time"
)

type UserProductBrowseListReq struct {
	g.Meta `path:"/front/shop/userProductBrowse/list" tags:"我的足迹" method:"get" summary:"我的足迹接口"`
	ml.BaseList

	ProductBrowseId uint64     `json:"product_browse_id"` // 浏览编号
	ItemId          uint64     `json:"item_id"`           // 商品编号
	UserId          uint       `json:"user_id"`           // 会员编号
	BrowseDate      *time.Time `json:"browse_date"`       // 浏览日期
	BrowseTime      uint64     `json:"browse_time"`       // 浏览时间
	CategoryId      uint       `json:"category_id"`       // 商品分类
}

// UserProductBrowseRes 我的足迹列表响应
type UserProductBrowseListRes struct {
	entity.UserProductBrowse
	ItemSalePrice    float64 `json:"item_sale_price"`    // 商品销售价
	ProductImage     string  `json:"product_image"`      // 图片信息
	ProductItemName  string  `json:"product_item_name"`  // Spec名称
	ProductName      string  `json:"product_name"`       // SPU商品名称
	ItemName         string  `json:"item_name"`          // 副标题(DOT):SKU名称
	ActivityTypeId   uint    `json:"activity_type_id"`   // 活动类型编号
	ActivityTypeName string  `json:"activity_type_name"` // 活动类型名称
}

type UserProductBrowseRemoveReq struct {
	g.Meta `path:"/front/shop/userProductBrowse/removeBrowser" tags:"商品浏览历史删除" method:"post" summary:"商品浏览历史删除"`
	UserId uint   `json:"user_id"`            // 会员编号
	ItemId uint64 `json:"item_id"           ` // 商品编号
}

type UserProductBrowseRemoveRes struct {
}
