package trade

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model"
)

type ReturnListReq struct {
	g.Meta `path:"/front/trade/orderReturn/list" tags:"售后订单" method:"get" summary:"退单列表接口"`
	ml.BaseList

	ReturnId          string `json:"return_id" dc:"退单号"`                                                                                                                // 退单号
	OrderId           string `json:"order_id" dc:"订单编号"`                                                                                                               // 订单编号
	BuyerUserId       uint   `json:"buyer_user_id" dc:"买家编号"`                                                                                                          // 买家编号
	ReturnStateId     uint   `json:"return_state_id" dc:"卖家处理状态(ENUM): 3100-【客户】提交退单;3105-退单审核;3110-收货确认;3115-退款确认;3120-【客户】收款确认;3125-完成"` // 卖家处理状态
	ReturnChannelCode string `json:"return_channel_code" dc:"退款渠道(ENUM):money-余额;alipay-支付宝;wx_native-微信"`                                                      // 退款渠道
	ReturnAddStart    uint64 `json:"return_add_start"  type:"GE"   dc:"退款完成时间-开始"`                                                                                 // 退款完成时间-开始
	ReturnAddEnd      uint64 `json:"return_add_end" type:"LE"    dc:"退款完成时间-结束"`
}

type ReturnListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ReturnDetailReq struct {
	g.Meta   `path:"/front/trade/orderReturn/get" tags:"售后订单" method:"get" summary:"退单详情接口"`
	ReturnId string `json:"return_id"   v:"required#请输入退单编号"                  ` // 退单号
}

type ReturnDetailRes model.OrderReturnVo

type ReturnEditReq struct {
	g.Meta `path:"/front/trade/orderReturn/edit" tags:"售后订单" method:"post" summary:"退单编辑接口"`

	ReturnId string `json:"return_id"                     ` // 退单号
	OrderReturnAdd
}

type ReturnEditRes struct {
	ReturnId interface{} `json:"return_id"                     ` // 退单号
}

type ReturnCancelReq struct {
	g.Meta            `path:"/front/trade/orderReturn/cancel" tags:"售后订单" method:"post" summary:"退单取消接口"`
	ReturnId          string `json:"return_id"   v:"required#请输入退单编号"                  ` // 退单号
	OrderCancelReason string `json:"order_cancel_reason"  dc:"取消原因"`
}

type ReturnCancelRes struct {
	OrderId []string `json:"return_id"  dc:"订单编号"`
}

type ReturnItemReq struct {
	g.Meta `path:"/front/trade/orderReturn/returnItem" tags:"售后订单" method:"post" summary:"退单新增接口"`

	// 退单号
	OrderId     string `json:"order_id" dc:"订单编号"`
	OrderItemId string `json:"order_item_id" dc:"订单商品编号"`
}

type ReturnItemRes model.OrderReturnItemVo

type ReturnAddReq struct {
	g.Meta `path:"/front/trade/orderReturn/add" tags:"售后订单" method:"post" summary:"退单新增接口"`

	ServiceTypeId      uint    `json:"service_type_id"     dc:"服务类型(ENUM):1-退款;2-退货;3-换货;4-维修"` // 服务类型(ENUM):1-退款;2-退货;3-换货;4-维修
	OrderId            string  `json:"order_id"            dc:"订单编号"`                                   // 订单编号
	OrderItemId        uint64  `json:"order_item_id"       dc:"订单商品表编号"`                             // 订单商品表编号
	ReturnRefundAmount float64 `json:"return_refund_amount" dc:"退款金额"`                                  // 退款金额
	ReturnItemNum      uint    `json:"return_item_num"     dc:"退货商品数量"`                               // 退货商品数量
	ReturnBuyerMessage string  `json:"return_buyer_message" dc:"买家退货备注"`                              // 买家退货备注
	ReturnTel          string  `json:"return_tel"          dc:"联系电话"`                                   // 联系电话
	ReturnReasonId     uint    `json:"return_reason_id"    dc:"退款理由编号"`                               // 退款理由编号
	ReturnItemImage    string  `json:"return_item_image"   dc:"退款凭据(DOT)"`                              // 退款凭据(DOT)
	ReturnAllFlag      bool    `json:"return_all_flag"     dc:"退货标记(BOOL):true-全退， false-单品退"`     // 退货标记(BOOL):true-全退， false-单品退
}

type ReturnAddRes struct {
	ReturnId string `json:"return_id"                     ` // 退单号
}
