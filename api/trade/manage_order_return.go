package trade

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model"
)

// start fo manage
type OrderReturnAdd struct {
	ServiceTypeId              uint        `json:"service_type_id"               ` // 服务类型(ENUM):1-退款;2-退货;3-换货;4-维修
	OrderId                    string      `json:"order_id"                      ` // 订单编号
	ReturnRefundAmount         float64     `json:"return_refund_amount"          ` // 退款金额
	ReturnRefundPoint          float64     `json:"return_refund_point"           ` // 积分部分
	StoreId                    uint        `json:"store_id"                      ` // 店铺编号
	BuyerUserId                uint        `json:"buyer_user_id"                 ` // 买家编号
	BuyerStoreId               uint        `json:"buyer_store_id"                ` // 买家是否有店铺
	ReturnAddTime              int64       `json:"return_add_time"               ` // 添加时间
	ReturnReasonId             uint        `json:"return_reason_id"              ` // 退款理由编号
	ReturnBuyerMessage         string      `json:"return_buyer_message"          ` // 买家退货备注
	ReturnAddrContacter        string      `json:"return_addr_contacter"         ` // 收货人
	ReturnTel                  string      `json:"return_tel"                    ` // 联系电话
	ReturnAddr                 string      `json:"return_addr"                   ` // 收货地址详情
	ReturnPostCode             int         `json:"return_post_code"              ` // 邮编
	ExpressId                  uint        `json:"express_id"                    ` // 物流公司编号
	ReturnTrackingName         string      `json:"return_tracking_name"          ` // 物流名称
	ReturnTrackingNumber       string      `json:"return_tracking_number"        ` // 物流单号
	PlantformReturnStateId     uint        `json:"plantform_return_state_id"     ` // 申请状态平台(ENUM):3180-处理中;3181-为待管理员处理卖家同意或者收货后;3182-为已完成
	ReturnStateId              uint        `json:"return_state_id"               ` // 卖家处理状态(ENUM): 3100-【客户】提交退单;3105-退单审核;3110-收货确认;3115-退款确认;3120-客户】收款确认;3125-完成
	ReturnIsPaid               uint        `json:"return_is_paid"                ` // 退款完成
	ReturnIsShippingFee        uint        `json:"return_is_shipping_fee"        ` // 退货类型(BOOL): 0-退款单;1-退运费单
	ReturnFlag                 uint        `json:"return_flag"                   ` // 退货类型(ENUM): 0-不用退货;1-需要退货
	ReturnType                 uint        `json:"return_type"                   ` // 申请类型(ENUM): 1-退款申请; 2-退货申请; 3-虚拟退款
	ReturnOrderLock            uint        `json:"return_order_lock"             ` // 订单锁定类型(BOOL):1-不用锁定;2-需要锁定
	ReturnItemStateId          uint        `json:"return_item_state_id"          ` // 物流状态(LIST):2030-待发货;2040-已发货/待收货确认;2060-已完成/已签收;2070-已取消/已作废;
	ReturnStoreTime            *gtime.Time `json:"return_store_time"             ` // 商家处理时间
	ReturnStoreMessage         string      `json:"return_store_message"          ` // 商家备注
	ReturnCommisionFee         float64     `json:"return_commision_fee"          ` // 退还佣金
	ReturnFinishTime           *gtime.Time `json:"return_finish_time"            ` // 退款完成时间
	ReturnPlatformMessage      string      `json:"return_platform_message"       ` // 平台留言
	ReturnIsSettlemented       uint        `json:"return_is_settlemented"        ` // 订单是否结算(BOOL): 0-未结算; 1-已结算
	ReturnSettlementTime       *gtime.Time `json:"return_settlement_time"        ` // 订单结算时间
	ReturnChannelCode          string      `json:"return_channel_code"           ` // 退款渠道(ENUM):money-余额;alipay-支付宝;wx_native-微信
	ReturnChannelFlag          uint        `json:"return_channel_flag"           ` // 渠道是否退款(ENUM): 0-待退; 1-已退; 2-异常
	ReturnChannelTime          *gtime.Time `json:"return_channel_time"           ` // 渠道退款时间
	ReturnChannelTransId       string      `json:"return_channel_trans_id"       ` // 渠道退款单号
	DepositTradeNo             string      `json:"deposit_trade_no"              ` // 交易号
	PaymentChannelId           uint        `json:"payment_channel_id"            ` // 支付渠道
	TradePaymentAmount         float64     `json:"trade_payment_amount"          ` // 实付金额:在线支付金额
	ReturnContactName          string      `json:"return_contact_name"           ` // 联系人
	ReturnStoreUserId          uint        `json:"return_store_user_id"          ` // 审核人员id
	ReturnMobile               uint64      `json:"return_mobile"                 ` // 手机号码
	ReturnTelephone            string      `json:"return_telephone"              ` // 卖家联系电话
	ReturnWithdrawConfirm      uint        `json:"return_withdraw_confirm"       ` // 提现审核(BOOL):0-未审核; 1-已审核
	ReturnFinancialConfirm     uint        `json:"return_financial_confirm"      ` // 退款财务确认(BOOL):0-未确认; 1-已确认
	ReturnFinancialConfirmTime *gtime.Time `json:"return_financial_confirm_time" ` // 退款财务确认时间
	SubsiteId                  uint        `json:"subsite_id"                    ` // 所属分站:0-总站
}

type OrderReturnEditReq struct {
	g.Meta `path:"/manage/trade/orderReturn/edit" tags:"售后订单" method:"post" summary:"退单编辑接口"`

	ReturnId string `json:"return_id"                     ` // 退单号
	OrderReturnAdd
}

type OrderReturnEditRes struct {
	ReturnId interface{} `json:"return_id"                     ` // 退单号
}

type OrderReturnEditStateReq struct {
	g.Meta `path:"/manage/trade/orderReturn/editState" tags:"售后订单" method:"post" summary:"售后订单状态编辑接口"`

	ReturnId string `json:"return_id"                     ` // 退单号
}

type OrderReturnEditStateRes struct {
	OrderId interface{} `json:"order_id"             ` // 订单编号
}

type OrderReturnAddReq struct {
	g.Meta `path:"/manage/trade/orderReturn/add" tags:"售后订单" method:"post" summary:"退单新增接口"`

	OrderReturnAdd
}

type OrderReturnDetailReq struct {
	g.Meta   `path:"/manage/trade/orderReturn/detail" tags:"售后订单" method:"get" summary:"退单详情接口"`
	ReturnId string `json:"return_id"   v:"required#请输入退单编号"                  ` // 退单号
}

type OrderReturnDetailRes struct {
	model.OrderVo
}

type OrderReturnListReq struct {
	g.Meta `path:"/manage/trade/orderReturn/list" tags:"售后订单" method:"get" summary:"退单列表接口"`
	ml.BaseList

	ReturnId          string `json:"return_id" dc:"退单号"`                                                                                   // 退单号
	OrderId           string `json:"order_id" dc:"订单编号"`                                                                                   // 订单编号
	BuyerUserId       uint   `json:"buyer_user_id" dc:"买家编号"`                                                                              // 买家编号
	ReturnStateId     uint   `json:"return_state_id" dc:"卖家处理状态(ENUM): 3100-【客户】提交退单;3105-退单审核;3110-收货确认;3115-退款确认;3120-【客户】收款确认;3125-完成"` // 卖家处理状态
	ReturnChannelCode string `json:"return_channel_code" dc:"退款渠道(ENUM):money-余额;alipay-支付宝;wx_native-微信"`                                 // 退款渠道
	ReturnAddStart    uint64 `json:"return_add_start"  type:"GE"   dc:"退款完成时间-开始"`                                                         // 退款完成时间-开始
	ReturnAddEnd      uint64 `json:"return_add_end" type:"LE"    dc:"退款完成时间-结束"`                                                           // 退款完成时间-结束
}

type OrderReturnListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type OrderReturnCancelReq struct {
	g.Meta            `path:"/manage/trade/orderReturn/cancel" tags:"售后订单" method:"post" summary:"退单取消接口"`
	ReturnId          string `json:"return_id"   v:"required#请输入退单编号"                  ` // 退单号
	OrderCancelReason string `json:"order_cancel_reason"  dc:"取消原因"`
}

type OrderReturnCancelRes struct {
	OrderId []string `json:"order_id"  dc:"订单编号"`
}

type OrderReturnReviewReq struct {
	g.Meta           `path:"/manage/trade/orderReturn/review" tags:"售后订单" method:"post" summary:"退单审核接口"`
	ReturnId         string `json:"return_id"   v:"required#请输入退单编号"                  ` // 退单号
	ReturnFlag       uint   `json:"return_flag"  dc:"退货类型(ENUM): 0-不用退货;1-需要退货"`
	ReceivingAddress uint   `json:"receiving_address"  dc:"收货地址"`
}

type OrderReturnReviewRes struct {
	ReturnId string `json:"return_id"   v:"required#请输入退单编号"                  ` // 退单号
}

type ReturnStateLogListReq struct {
	g.Meta `path:"/manage/trade/orderReturn/listStateLog" tags:"售后订单" method:"get" summary:"退单日志列表接口"`
	ml.BaseList

	ReturnId string `json:"return_id"   v:"required#请输入退单编号"                  ` // 退单号
}

type ReturnStateLogListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type GetByReturnIdReq struct {
	g.Meta `path:"/manage/trade/orderReturn/getByReturnId" tags:"退款退货" method:"get" summary:"退款退货详情接口"`
	ml.BaseList

	ReturnId string `json:"return_id"   v:"required#请输入退单编号"                  ` // 退单号
}

type GetByReturnIdRes model.OrderReturnVo

type OrderReturnReceiveReq struct {
	g.Meta `path:"/manage/trade/orderReturn/receive" tags:"退货单审核-确认收货" method:"post" summary:"退货单审核-确认收货接口"`
	ml.BaseList

	ReturnId string `json:"return_id"   v:"required#请输入退单编号"                  ` // 退单号
}

type OrderReturnReceiveRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type OrderReturnRefundReq struct {
	g.Meta `path:"/manage/trade/orderReturn/refund" tags:"退货单审核-确认收货" method:"post" summary:"退货单审核-确认收货接口"`
	ml.BaseList

	ReturnId string `json:"return_id"   v:"required#请输入退单编号"                  ` // 退单号
}

type OrderReturnRefundRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type OrderReturnRefusedReq struct {
	g.Meta `path:"/manage/trade/orderReturn/refused" tags:"退货单审核-确认收货" method:"post" summary:"退货单审核-确认收货接口"`
	ml.BaseList

	ReturnId string `json:"return_id"   v:"required#请输入退单编号"                  ` // 退单号
}

type OrderReturnRefusedRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
