package analytics

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/model"
)

type ReturnAmountReq struct {
	g.Meta `path:"/manage/analytics/return/getReturnAmount" tags:"售后" method:"get" summary:"退款总额接口"`

	ProductName   string `json:"product_name" dc:"产品名称"`                                                                               // 产品名称
	ProductId     int64  `json:"product_id" dc:"产品编号"`                                                                                 // 产品编号
	ReturnStateId []uint `json:"return_state_id" dc:"卖家处理状态(ENUM): 3100-【客户】提交退单;3105-退单审核;3110-收货确认;3115-退款确认;3120-【客户】收款确认;3125-完成"` // 卖家处理状态
	StoreId       uint   `json:"store_id" dc:"店铺编号"`                                                                                   // 店铺编号
	StoreType     uint   `json:"store_type" dc:"店铺类型"`                                                                                 // 店铺类型
}

type ReturnAmountRes struct {
	model.AnalyticsNumVo
}

type ReturnNumReq struct {
	g.Meta `path:"/manage/analytics/return/getReturnNum" tags:"售后" method:"get" summary:"退款数量接口"`

	ProductName   string `json:"product_name" dc:"产品名称"`                                                                               // 产品名称
	ProductId     uint64 `json:"product_id" dc:"产品编号"`                                                                                 // 产品编号
	ReturnStateId []uint `json:"return_state_id" dc:"卖家处理状态(ENUM): 3100-【客户】提交退单;3105-退单审核;3110-收货确认;3115-退款确认;3120-【客户】收款确认;3125-完成"` // 卖家处理状态
	StoreId       uint   `json:"store_id" dc:"店铺编号"`                                                                                   // 店铺编号
	StoreType     uint   `json:"store_type" dc:"店铺类型"`                                                                                 // 店铺类型
}
type ReturnNumRes struct {
	model.AnalyticsNumVo
}

type ReturnAmountTimelineReq struct {
	g.Meta `path:"/manage/analytics/return/getReturnAmountTimeline" tags:"售后" method:"get" summary:"退款数量接口"`

	TimelineVo
}
type ReturnAmountTimelineRes []*model.TimelineOutput

type ReturnNumTimelineReq struct {
	g.Meta `path:"/manage/analytics/return/getReturnNumTimeline" tags:"售后" method:"get" summary:"退款数量接口"`

	TimelineVo

	CategoryId  uint     `json:"category_id" dc:"分类编号"`  // 分类编号
	ProductName string   `json:"product_name" dc:"产品名称"` // 产品名称
	ProductId   uint64   `json:"product_id" dc:"产品编号"`   // 产品编号
	ItemId      []uint64 `json:"item_id" dc:"SKU编号"`     // SKU编号
	StoreId     uint     `json:"store_id" dc:"店铺编号"`     // 店铺编号
	StoreType   uint     `json:"store_type" dc:"店铺类型"`   // 店铺类型
	KindId      uint     `json:"kind_id" dc:"订单类型"`      // 订单类型
}
type ReturnNumTimelineRes []*model.TimelineOutput

type ReturnItemNumTimelineReq struct {
	g.Meta `path:"/manage/analytics/return/getReturnItemNumTimeLine" tags:"售后" method:"get" summary:"退单商品销量统计"`

	TimelineVo

	CategoryId  uint     `json:"category_id" dc:"分类编号"`  // 分类编号
	ProductName string   `json:"product_name" dc:"产品名称"` // 产品名称
	ProductId   uint64   `json:"product_id" dc:"产品编号"`   // 产品编号
	ItemId      []uint64 `json:"item_id" dc:"SKU编号"`     // SKU编号
	StoreId     uint     `json:"store_id" dc:"店铺编号"`     // 店铺编号
	StoreType   uint     `json:"store_type" dc:"店铺类型"`   // 店铺类型
	KindId      uint     `json:"kind_id" dc:"订单类型"`      // 订单类型
}
type ReturnItemNumTimelineRes []*model.TimelineOutput

type ReturnItemNumListReq struct {
	g.Meta `path:"/manage/analytics/return/listReturnItemNum" tags:"售后" method:"get" summary:"退单商品销量统计"`

	TimelineVo

	CategoryId  uint     `json:"category_id" dc:"分类编号"`  // 分类编号
	ProductName string   `json:"product_name" dc:"产品名称"` // 产品名称
	ProductId   uint64   `json:"product_id" dc:"产品编号"`   // 产品编号
	ItemId      []uint64 `json:"item_id" dc:"SKU编号"`     // SKU编号
	StoreId     uint     `json:"store_id" dc:"店铺编号"`     // 店铺编号
	StoreType   uint     `json:"store_type" dc:"店铺类型"`   // 店铺类型
	KindId      uint     `json:"kind_id" dc:"订单类型"`      // 订单类型
}
type ReturnItemNumListRes []*model.AnalyticsOrderItemNumOutput

type ReturnItemNumReq struct {
	g.Meta `path:"/manage/analytics/return/getReturnItemNum" tags:"售后" method:"get" summary:"商品访客数"`

	TimelineVo

	CategoryId  uint     `json:"category_id" dc:"分类编号"`  // 分类编号
	ProductName string   `json:"product_name" dc:"产品名称"` // 产品名称
	ProductId   uint64   `json:"product_id" dc:"产品编号"`   // 产品编号
	ItemId      []uint64 `json:"item_id" dc:"SKU编号"`     // SKU编号
	StoreId     uint     `json:"store_id" dc:"店铺编号"`     // 店铺编号
	StoreType   uint     `json:"store_type" dc:"店铺类型"`   // 店铺类型
	KindId      uint     `json:"kind_id" dc:"订单类型"`      // 订单类型
}
type ReturnItemNumRes *model.TimelineOutput
