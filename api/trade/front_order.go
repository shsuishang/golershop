package trade

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model"
)

// 用户订单表
type UserOrderListReq struct {
	g.Meta `path:"/front/trade/order/list" tags:"交易订单" method:"get" summary:"用户订单列表"`

	ml.BaseList

	OrderId                     string `json:"order_id"                       `                               // 订单编号
	OrderTitle                  string `json:"order_title"   type:"LIKE" field:"order_title"                ` // 订单标题
	StoreId                     uint   `json:"store_id"                       `                               // 卖家店铺编号
	SubsiteId                   uint   `json:"subsite_id"                     `                               // 所属分站:0-总站
	UserId                      uint   `json:"user_id"                        `                               // 买家编号
	KindId                      uint   `json:"kind_id"                        `                               // 订单种类(ENUM): 1201-实物 ; 1202-教育类 ; 1203-电子卡券  ; 1204-其它
	OrderIsSettlemented         uint   `json:"order_is_settlemented"          `                               // 订单是否结算(BOOL):0-未结算; 1-已结算
	OrderBuyerEvaluationStatus  uint   `json:"order_buyer_evaluation_status"  `                               // 买家针对订单对店铺评价(ENUM): 0-未评价;1-已评价;  2-已过期未评价
	OrderSellerEvaluationStatus uint   `json:"order_seller_evaluation_status" `                               // 卖家评价状态(ENUM):0-未评价;1-已评价;  2-已过期未评价
	OrderYear                   uint   `json:"order_year"                     `                               // 订单年份
	OrderMonth                  uint   `json:"order_month"                    `                               // 订单月份
	OrderDay                    uint   `json:"order_day"                      `                               // 订单日
	PaymentTypeId               uint   `json:"payment_type_id"                `                               // 支付方式(ENUM): 1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付;
	OrderStateId                uint   `json:"order_state_id"                 `                               // 订单状态(LIST):2011-待订单审核;2013-待财务审核;2020-待配货/待出库审核;2030-待发货;2040-已发货/待收货确认;2060-已完成/已签收;2070-已取消/已作废;
	OrderIsReview               bool   `json:"order_is_review"                `                               // 订单审核(BOOL):0-未审核;1-已审核;
	OrderFinanceReview          bool   `json:"order_finance_review"           `                               // 财务状态(BOOL):0-未审核;1-已审核
	OrderIsPaid                 uint   `json:"order_is_paid"                  `                               // 付款状态(ENUM):3010-未付款;3011-付款待审核;3012-部分付款;3013-已付款
	OrderIsOut                  uint   `json:"order_is_out"                   `                               // 出库状态(ENUM):3020-未出库;3021-部分出库通过拆单解决这种问题;3022-已出库
	OrderIsShipped              uint   `json:"order_is_shipped"               `                               // 发货状态(ENUM):3030-未发货;3032-已发货;3031-部分发货
	OrderIsReceived             int64  `json:"order_is_received"              `                               // 收货状态(BOOL):0-未收货;1-已收货
	ChainId                     uint   `json:"chain_id"                       `                               // 门店编号
	DeliveryTypeId              uint   `json:"delivery_type_id"               `                               // 配送方式
	OrderIsOffline              bool   `json:"order_is_offline"               `                               // 线下订单(BOOL):0-线上;1-线下
	OrderExpressPrint           bool   `json:"order_express_print"            `                               // 是否打印(BOOL):0-未打印;1-已打印
	OrderIsSync                 bool   `json:"order_is_sync"                  `                               // 是否ERP同步(BOOL):0-未同步; 1-已同步
	OrderFxIsSettlemented       int64  `json:"order_fx_is_settlemented"       `                               // 佣金是否发放(BOOL):0 -未发放;1 -已发放
	CreateSTime                 uint64 `json:"order_stime" field:"create_time" type:"GE"               `      // 下单时间:检索使用
	CreateETime                 uint64 `json:"order_etime" field:"create_time" type:"LE"                    ` // 下单时间:检索使用

}

type UserOrderListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

// UserOrderAddReq 订单添加请求结构体
type UserOrderAddReq struct {
	g.Meta         `path:"/front/trade/order/add" tags:"交易订单" method:"post" summary:"用户订单添加"`
	UdId           int    `json:"ud_id"`            // 收货地址编号
	CartId         string `json:"cart_id"`          // 下单商品数据:商品编号|数量,商品编号|数量...
	ChainId        int    `json:"chain_id"`         // 门店编号
	ActivityId     int    `json:"activity_id"`      // 活动编号
	GbId           int    `json:"gb_id"`            // 拼团编号
	DeliveryTypeId int    `json:"delivery_type_id"` // 配送方式
	ChannelType    int    `json:"channel_type"`     // 来源渠道 0:正常下单;1:直播渠道
	PaymentTypeId  int    `json:"payment_type_id"`  // 付款方式
	OrderMessage   string `json:"order_message"`    // 消息
	UserVoucherIds string `json:"user_voucher_ids"` // 优惠券
	UserInvoiceId  int    `json:"user_invoice_id"`  // 发票
	UserId         int    `json:"user_id"`          // 买家编号
	ProductItems   string `json:"product_items"`    // SKU信息
}

type UserOrderAddRes model.OrderAddOutput

type UserOrderDetailReq struct {
	g.Meta  `path:"/front/trade/order/detail" tags:"交易订单" method:"get" summary:"订单详情接口"`
	OrderId string `json:"order_id" v:"required#请输入订单编号"   dc:"订单编号"`
}

type UserOrderDetailRes struct {
	model.OrderVo
}

type UserOrderCancelReq struct {
	g.Meta            `path:"/front/trade/order/cancel" tags:"交易订单" method:"post" summary:"订单取消接口"`
	OrderId           []string `json:"order_id" v:"required#请输入订单编号"   dc:"订单编号"`
	OrderCancelReason string   `json:"order_cancel_reason"  dc:"取消原因"`
}

type UserOrderCancelRes struct {
	OrderId []string `json:"order_id"  dc:"订单编号"`
}

type UserOrderReviewReq struct {
	g.Meta            `path:"/front/trade/order/review" tags:"交易订单" method:"post" summary:"订单审核接口"`
	OrderId           []string `json:"order_id" v:"required#请输入订单编号"   dc:"订单编号"`
	OrderReviewReason string   `json:"order_review_reason"  dc:"原因"`
}

type UserOrderReviewRes struct {
	OrderId []string `json:"order_id"  dc:"订单编号"`
}

type UserOrderFinanceReq struct {
	g.Meta             `path:"/front/trade/order/finance" tags:"交易订单" method:"post" summary:"订单财务审核接口"`
	OrderId            []string `json:"order_id" v:"required#请输入订单编号"   dc:"订单编号"`
	OrderFinanceReason string   `json:"order_finance_reason"  dc:"原因"`
}

type UserOrderFinanceRes struct {
	OrderId []string `json:"order_id"  dc:"订单编号"`
}

type UserOrderPickingReq struct {
	g.Meta  `path:"/front/trade/order/picking" tags:"交易订单" method:"post" summary:"订单出库审核接口"`
	OrderId []string             `json:"order_id" v:"required#请输入订单编号"   dc:"订单编号"`
	Items   []*model.PickingItem `json:"items"          dc:"出库商品信息"        ` // 出库商品信息

	BillTypeId           uint `json:"bill_type_id"    d:"2700"     dc:"业务类型"           `    // 业务类型
	WarehouseId          uint `json:"warehouse_id"   d:"0"       dc:"地址编号"           `      // 地址编号
	StockTransportTypeId uint `json:"stock_transport_type_id"  d:"2751"  dc:"库存类型"        ` //

	PickingFlag bool `json:"picking_flag" d:"true"  dc:"出库标记:true-默认全出， false-指定出库"        ` //
}

type UserOrderPickingRes struct {
	OrderId []string `json:"order_id"  dc:"订单编号"`
}

type UserOrderShippingReq struct {
	g.Meta  `path:"/front/trade/order/shipping" tags:"交易订单" method:"post" summary:"订单发货审核接口"`
	OrderId []string `json:"order_id" v:"required#请输入订单编号"   dc:"订单编号"`

	StockBillId         string      `json:"stock_bill_id"     dc:"出库单编号"           `            // 出库单编号
	SsId                uint        `json:"ss_id"        dc:"发货地址"           `                  // 发货地址
	LogisticsId         uint        `json:"logistics_id"        dc:"发货物流编号"           `         // 发货物流编号
	LogisticsTime       *gtime.Time `json:"logistics_time"        dc:"发货时间"           `         // 发货时间
	OrderTrackingNumber string      `json:"order_tracking_number"          dc:"运单号"           ` // 运单号
	LogisticsExplain    string      `json:"logistics_explain"   dc:"备注"        `                //

	ShippingFlag bool `json:"shipping_flag" d:"false"  dc:"发货标记:true-默认全发， false-指定发货"        ` //
}

type UserOrderShippingRes struct {
	OrderId []string `json:"order_id"  dc:"订单编号"`
}

type UserOrderReceiveReq struct {
	g.Meta             `path:"/front/trade/order/receive" tags:"交易订单" method:"post" summary:"订单确认收货接口"`
	OrderId            []string `json:"order_id" v:"required#请输入订单编号"   dc:"订单编号"`
	OrderFinanceReason string   `json:"order_finance_reason"  dc:"原因"`
}

type UserOrderReceiveRes struct {
	OrderId []string `json:"order_id"  dc:"订单编号"`
}

type UserOrderStateLogListReq struct {
	g.Meta `path:"/front/trade/order/listStateLog" tags:"交易订单" method:"get" summary:"订单日志列表接口"`
	ml.BaseList

	OrderId string `json:"order_id"  v:"required#请输入订单编号"   ` // 订单编号
}

type UserOrderStateLogListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type UserOrderNumReq struct {
	g.Meta `path:"/front/trade/order/getOrderNum" tags:"交易订单" method:"get" summary:"用户中心订单数量"`
}

type UserOrderNumRes model.OrderNumOutput

type UserOrderInvoiceListReq struct {
	g.Meta `path:"/front/trade/order/listInvoice" tags:"发票管理" method:"get" summary:"发票管理列表接口"`
	ml.BaseList

	UserId        uint `json:"user_id"                   `        // 用户编号
	InvoiceStatus uint `json:"invoice_status"                   ` // 状态

	OrderId      string `json:"order_id"   type:"LIKE"       ` // 订单编号
	InvoiceTitle string `json:"invoice_title"  type:"LIKE"   ` // 发票抬头
}

type UserOrderInvoiceListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type UserOrderInvoiceAddReq struct {
	g.Meta `path:"/front/trade/order/addInvoice" tags:"发票管理" method:"post" summary:"发票管理编辑接口"`

	OrderInvoiceAdd
}

type UserOrderInvoiceAddRes struct {
	OrderInvoiceId uint `json:"order_invoice_id" dc:"发票管理编号"   ` // 发票管理编号
}

type UserOrderInvoiceEditReq struct {
	g.Meta `path:"/front/trade/order/editInvoice" tags:"发票管理" method:"post" summary:"发票管理编辑接口"`

	OrderInvoiceId uint `json:"order_invoice_id"  dc:"发票管理编号"   ` // 发票管理编号
	OrderInvoiceAdd
}

type UserOrderInvoiceEditRes struct {
	OrderInvoiceId uint `json:"order_invoice_id" dc:"发票管理编号"   ` // 发票管理编号
}
