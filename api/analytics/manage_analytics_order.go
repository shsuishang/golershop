package analytics

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/model"
)

type TimelineVo struct {
	Stime int64 `json:"stime"  type:"GE"     ` // 开始时间
	Etime int64 `json:"etime"  type:"LE"     ` // 截止时间
}

type OrderReq struct {
	TimelineVo

	CategoryId   uint   `json:"category_id"    dc:"分类编号"`               // 分类编号
	ProductName  string `json:"product_name"  type:"LIKE"    dc:"产品名称"` // 产品名称
	ProductId    uint64 `json:"product_id"     dc:"产品编号"`               // 产品编号
	OrderStateId []uint `json:"order_state_id" dc:"订单状态"`               // 订单状态
	OrderIsPaid  []uint `json:"order_is_paid"  dc:"支付状态"`               // 支付状态
	StoreId      uint   `json:"store_id"       dc:"店铺编号"`               // 店铺编号
	StoreType    uint   `json:"store_type"     dc:"店铺类型"`               // 店铺类型
	KindId       uint   `json:"kind_id"        dc:"订单类型"`               // 订单类型
}

type OrderAmountReq struct {
	g.Meta `path:"/manage/analytics/order/getOrderAmount" tags:"订单" method:"get" summary:"交易销售额统计接口"`

	OrderReq
}

type OrderAmountRes model.AnalyticsNumVo

type VisitorReq struct {
	g.Meta `path:"/manage/analytics/user/getVisitor" tags:"订单" method:"get" summary:"用户访问量"`
}

type VisitorRes model.DashboardTopVo

type OrderNumTodayReq struct {
	g.Meta `path:"/manage/analytics/order/getOrderNumToday" tags:"订单" method:"get" summary:"获取订单量"`
}

type OrderNumTodayRes model.DashboardTopVo

type OrderNumReq struct {
	g.Meta `path:"/manage/analytics/order/getOrderNum" tags:"订单" method:"get" summary:"获取订单量"`
	OrderReq
}

type OrderNumRes model.AnalyticsNumVo

type RegUserReq struct {
	g.Meta `path:"/manage/analytics/user/getRegUser" tags:"订单" method:"get" summary:"获取新增用户"`
}

type RegUserRes model.DashboardTopVo

type AnalyticsProductReq struct {
	g.Meta `path:"/manage/analytics/product/getProductNum" tags:"商品" method:"get" summary:"商品数量"`

	CategoryId     uint64 `json:"category_id" description:"分类编号"`      // 分类编号
	ProductName    string `json:"product_name" description:"产品名称"`     // 产品名称
	StoreId        uint   `json:"store_id" description:"店铺编号"`         // 店铺编号
	StoreType      uint   `json:"store_type" description:"店铺类型"`       // 店铺类型
	ProductStateId uint   `json:"product_state_id" description:"商品状态"` // 商品状态
}

type AnalyticsProductRes model.AnalyticsNumVo

type UserTimeLineReq struct {
	g.Meta `path:"/manage/analytics/user/getUserTimeLine" tags:"用户" method:"get" summary:"新增用户"`

	TimelineVo

	ItemId uint64 `json:"item_id" description:"SKU编号"` // SKU编号
}

type UserTimeLineRes []*model.TimelineOutput

type UserNumReq struct {
	g.Meta `path:"/manage/analytics/user/getUserNum" tags:"用户" method:"get" summary:"获取新增用户"`
	TimelineVo
}

type UserNumRes model.AnalyticsNumVo

type DashboardTimeLineReq struct {
	g.Meta `path:"/manage/analytics/order/getDashboardTimeLine" tags:"订单" method:"get" summary:"统计"`

	TimelineVo
}

// DashBoardTimelineRes 订单编辑接口
type DashboardTimeLineRes model.DashBoardTimelineOutput

type OrderCustomerNumTimelineReq struct {
	g.Meta `path:"/manage/analytics/order/getOrderCustomerNumTimeline" tags:"订单" method:"get" summary:"统计"`

	TimelineVo
}

type OrderCustomerNumTimelineRes []*model.TimelineOutput

type SaleOrderAmountReq struct {
	g.Meta `path:"/manage/analytics/order/getSaleOrderAmount" tags:"订单" method:"get" summary:"订单销售金额对比图"`

	TimelineVo
}

type SaleOrderAmountRes []*model.TimelineOutput

type OrderNumTimelineReq struct {
	g.Meta `path:"/manage/analytics/order/getOrderNumTimeline" tags:"订单" method:"get" summary:"订单数量统计"`

	TimelineVo
}

type OrderNumTimelineRes []*model.TimelineOutput

type OrderItemNumTimeLineReq struct {
	g.Meta `path:"/manage/analytics/order/getOrderItemNumTimeLine" tags:"订单" method:"get" summary:"订单数量统计"`

	OrderReq
}

type OrderItemNumTimeLineRes []*model.TimelineOutput

type OrderItemNumListReq struct {
	g.Meta `path:"/manage/analytics/order/listOrderItemNum" tags:"订单" method:"get" summary:"订单商品销量统计"`

	OrderReq
}

type OrderItemNumListRes []*model.AnalyticsOrderItemNumOutput

type AccessItemTimelineReq struct {
	g.Meta `path:"/manage/analytics/history/getAccessItemTimeLine" tags:"订单" method:"get" summary:"用户访问商品统计"`

	TimelineVo
	ItemId uint64 `json:"item_id" description:"SKU编号"` // SKU编号
}

type AccessItemTimelineRes []*model.TimelineOutput

type AccessItemUserTimeLineReq struct {
	g.Meta `path:"/manage/analytics/history/getAccessItemUserTimeLine" tags:"订单" method:"get" summary:"访客数"`

	TimelineVo
	ItemId uint64 `json:"item_id" description:"SKU编号"` // SKU编号
}

type AccessItemUserTimeLineRes []*model.TimelineOutput

type AccessVisitorTimeLineReq struct {
	g.Meta `path:"/manage/analytics/history/getAccessVisitorTimeLine" tags:"订单" method:"get" summary:"访客数"`

	TimelineVo
	ItemId uint64 `json:"item_id" description:"SKU编号"` // SKU编号
}

type AccessVisitorTimeLineRes []*model.TimelineOutput

type AccessVisitorNumReq struct {
	g.Meta `path:"/manage/analytics/history/getAccessVisitorNum" tags:"订单" method:"get" summary:"访客数"`

	TimelineVo
	ItemId uint64 `json:"item_id" description:"SKU编号"` // SKU编号
}

type AccessVisitorNumRes model.AnalyticsNumVo

type AccessNumReq struct {
	g.Meta `path:"/manage/analytics/history/getAccessNum" tags:"订单" method:"get" summary:"访客数"`

	TimelineVo
	ItemId uint64 `json:"item_id" description:"SKU编号"` // SKU编号
}

type AccessNumRes model.AnalyticsNumVo

type AccessItemNumReq struct {
	g.Meta `path:"/manage/analytics/history/getAccessItemNum" tags:"订单" method:"get" summary:"访客数"`

	TimelineVo
	ItemId uint64 `json:"item_id" description:"SKU编号"` // SKU编号
}

type AccessItemNumRes model.AnalyticsNumVo

type AccessItemListReq struct {
	g.Meta `path:"/manage/analytics/history/listAccessItem" tags:"订单" method:"get" summary:"访客数"`

	TimelineVo
	ItemId uint64 `json:"item_id" description:"SKU编号"` // SKU编号
}

type AccessItemListRes []*model.AnalyticsAccessItemOutput

type AccessItemUserNumReq struct {
	g.Meta `path:"/manage/analytics/history/getAccessItemUserNum" tags:"订单" method:"get" summary:"访客数"`

	TimelineVo
	ItemId uint64 `json:"item_id" description:"SKU编号"` // SKU编号
}

type AccessItemUserNumRes model.AnalyticsNumVo

type OrderItemNumReq struct {
	g.Meta `path:"/manage/analytics/order/getOrderItemNum" tags:"订单" method:"get" summary:"访客数"`

	TimelineVo
	ItemId uint64 `json:"item_id" description:"SKU编号"` // SKU编号
}

type OrderItemNumRes model.AnalyticsNumVo
