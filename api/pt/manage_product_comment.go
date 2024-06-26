package pt

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type ProductCommentAdd struct {
	CommentId          uint64      `json:"comment_id"           ` // 评价编号
	OrderId            string      `json:"order_id"             ` // 订单编号
	ProductId          uint64      `json:"product_id"           ` // 产品编号
	ItemId             uint64      `json:"item_id"              ` // 商品编号
	ItemName           string      `json:"item_name"            ` // 商品规格
	StoreId            uint        `json:"store_id"             ` // 店铺编号
	StoreName          string      `json:"store_name"           ` // 店铺名称
	UserId             uint        `json:"user_id"              ` // 买家编号
	UserName           string      `json:"user_name"            ` // 买家姓名:user_nickname
	CommentPoints      float64     `json:"comment_points"       ` // 获得积分:冗余，独立表记录
	CommentScores      uint        `json:"comment_scores"       ` // 评价星级:1-5积分
	CommentContent     string      `json:"comment_content"      ` // 评价内容
	CommentImage       string      `json:"comment_image"        ` // 评论上传的图片(DOT)
	CommentHelpful     uint        `json:"comment_helpful"      ` // 有帮助
	CommentNohelpful   uint        `json:"comment_nohelpful"    ` // 无帮助
	CommentTime        *gtime.Time `json:"comment_time"         ` // 评价时间
	CommentIsAnonymous uint        `json:"comment_is_anonymous" ` // 匿名评价
	CommentEnable      bool        `json:"comment_enable"       ` // 评价信息的状态(BOOL): 1-正常显示; 0-禁止显示
	ChainId            uint        `json:"chain_id"             ` // 门店编号
	SubsiteId          uint        `json:"subsite_id"           ` // 所属分站:0-总站
}
type ProductCommentEditReq struct {
	g.Meta `path:"/manage/pt/productComment/edit" tags:"商品评价表" method:"post" summary:"商品评价表编辑接口"`

	ProductCommentAdd
}

type ProductCommentEditRes struct {
	CommentId int64 `json:"comment_id"           ` // 评价编号
}

type ProductCommentAddReq struct {
	g.Meta `path:"/manage/pt/productComment/add" tags:"商品评价表" method:"post" summary:"商品评价表编辑接口"`

	ProductCommentAdd
}

type ProductCommentRemoveReq struct {
	g.Meta `path:"/manage/pt/productComment/remove" tags:"商品评价表" method:"post" summary:"商品评价表删除接口"`

	CommentId int64 `json:"comment_id"           ` // 评价编号
}

type ProductCommentRemoveRes struct {
}

type ProductCommentListReq struct {
	g.Meta `path:"/manage/pt/productComment/list" tags:"商品评价表" method:"get" summary:"商品评价表列表接口"`
	ml.BaseList

	ProductId uint64 `json:"product_id"           ` // 产品编号
	ItemName  string `json:"item_name"            ` // 商品规格
}

type ProductCommentListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ProductCommentEditStateReq struct {
	g.Meta `path:"/manage/pt/productComment/editState" tags:"商品评价表" method:"post" summary:"商品评价表状态修改接口"`

	CommentId     int64 `json:"comment_id"           ` // 评价编号
	CommentEnable bool  `json:"comment_enable"       ` // 评价信息的状态(BOOL): 1-正常显示; 0-禁止显示
}

type ProductCommentEditStateRes struct {
}
