package trade

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo manage
type OrderLogisticsAdd struct {
	OrderId             string      `json:"order_id"              ` // 订单编号
	StockBillId         string      `json:"stock_bill_id"         ` // 出入库单据id=stock_bill_id
	OrderTrackingNumber string      `json:"order_tracking_number" ` // 订单物流单号AIRWAY BILL number
	LogisticsExplain    string      `json:"logistics_explain"     ` // 卖家备注发货备忘
	LogisticsTime       *gtime.Time `json:"logistics_time"        ` // 发货时间配送时间
	LogisticsId         uint        `json:"logistics_id"          ` // 对应快递公司
	SsId                uint        `json:"ss_id"          `        // 地址编号
}

type OrderLogisticsAddReq struct {
	g.Meta `path:"/manage/trade/orderLogistics/add" tags:"物流信息" method:"post" summary:"物流信息编辑接口"`

	OrderLogisticsAdd
}

type OrderLogisticsEditReq struct {
	g.Meta `path:"/manage/trade/orderLogistics/edit" tags:"物流信息" method:"post" summary:"物流信息编辑接口"`

	OrderLogisticsId uint64 `json:"order_logistics_id"    ` // 订单物流编号
	OrderLogisticsAdd
}

type OrderLogisticsEditRes struct {
	OrderLogisticsId uint64 `json:"order_logistics_id"    ` // 订单物流编号
}

type OrderLogisticsRemoveReq struct {
	g.Meta           `path:"/manage/trade/orderLogistics/remove" tags:"物流信息" method:"post" summary:"物流信息删除接口"`
	OrderLogisticsId uint64 `json:"order_logistics_id"    ` // 订单物流编号
}

type OrderLogisticsRemoveRes struct {
}

type OrderLogisticsListReq struct {
	g.Meta `path:"/manage/trade/orderLogistics/list" tags:"物流信息" method:"get" summary:"物流信息列表接口"`
	ml.BaseList

	OrderLogisticsId uint64 `json:"order_logistics_id"    ` // 订单物流编号
}

type OrderLogisticsListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
