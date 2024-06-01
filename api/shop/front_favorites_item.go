package shop

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type UserFavoritesItemAdd struct {
	ItemId uint64 `json:"item_id" v:"required#请输入商品编号"            ` // 商品编号
}
type UserFavoritesItemEditReq struct {
	g.Meta `path:"/front/shop/userFavoritesItem/edit" tags:"收藏" method:"post" summary:"收藏编辑接口"`

	UserFavoritesItemAdd
}

type UserFavoritesItemEditRes struct {
	FavoritesItemId uint `json:"favorites_item_id"   ` // 收藏编号
}

type UserFavoritesItemAddReq struct {
	g.Meta `path:"/front/shop/userFavoritesItem/add" tags:"收藏" method:"post" summary:"收藏编辑接口"`

	UserFavoritesItemAdd
}

type UserFavoritesItemAddRes struct {
	FavoritesItemId uint `json:"favorites_item_id"   ` // 收藏编号
}

type UserFavoritesItemRemoveReq struct {
	g.Meta `path:"/front/shop/userFavoritesItem/remove" tags:"收藏" method:"post" summary:"收藏删除接口"`
	ItemId uint64 `json:"item_id" v:"required#请输入商品编号"            ` // 商品编号
}

type UserFavoritesItemRemoveRes struct {
}

type UserFavoritesItemListReq struct {
	g.Meta `path:"/front/shop/userFavoritesItem/list" tags:"收藏" method:"get" summary:"收藏列表接口"`
	ml.BaseList

	UserId uint   `json:"user_id"             ` // 用户编号
	ItemId uint64 `json:"item_id"             ` // 商品编号
}

type UserFavoritesItemListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
