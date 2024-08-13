package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"golershop.cn/internal/model/entity"
)

type StockBillVo struct {
	entity.StockBill
	Items []*entity.StockBillItem `json:"items"          dc:"出库商品信息"        ` // 出库商品信息
}

type OrderItemVo struct {
	entity.OrderItem

	CommentImage            []string                      `json:"comment_image"        `   // 评论上传的图片(DOT)
	CommentScores           uint                          `json:"comment_scores"        `  // 评论分数
	CommentContent          string                        `json:"comment_content"        ` // 评论内容
	ProductCommentReplyList []*entity.ProductCommentReply `json:"product_comment_reply_list"        `
}

type OrderVo struct {
	entity.OrderInfo
	entity.OrderData

	//order_base
	OrderId            string      `json:"order_id"                       ` // 订单编号
	OrderNumber        string      `json:"order_number"         `           // 订单编号
	OrderTime          *gtime.Time `json:"order_time"           `           // 下单时间
	OrderProductAmount float64     `json:"order_product_amount" `           // 商品原价总和:商品发布原价
	OrderPaymentAmount float64     `json:"order_payment_amount" `           // 应付金额:order_product_amount - order_discount_amount + order_shipping_fee - order_voucher_price - order_points_fee - order_adjust_fee
	CurrencyId         uint        `json:"currency_id"          `           // 货币编号
	CurrencySymbolLeft string      `json:"currency_symbol_left" `           // 左符号
	StoreName          string      `json:"store_name"           `           // 店铺名称
	UserNickname       string      `json:"user_nickname"        `           // 买家昵称

	//trade
	TradePaymentAmount float64 `json:"trade_payment_amount"        ` // 实付金额:在线支付金额,此为订单默认需要支付额度。

	Items          []*OrderItemVo               `json:"items"        `                          // 订单SKU
	Delivery       *entity.OrderDeliveryAddress `json:"delivery"     `                          //配送地址
	Logistics      []*entity.OrderLogistics     `json:"logistics"        `                      // 订单物流
	LogItems       []*entity.OrderStateLog      `json:"log_items"        `                      // 订单记录
	StockBill      []*StockBillVo               `json:"stock_bill"        `                     // 出库单
	ConsumeRecord  []*entity.ConsumeRecord      `json:"consume_record"        `                 // 支付记录
	ConsumeTrade   *entity.ConsumeTrade         `json:"consume_trade"                       `   // 交易数据
	WarehouseItems []*entity.WarehouseItem      `json:"warehouse_items"                       ` // 库存数据

	ReturnFlag     bool  `json:"return_flag"                       `      // 退款标记
	InvoiceIsApply bool  `json:"invoice_is_apply"                       ` // 发票标记
	RemainPayTime  int64 `json:"remain_pay_time"        `                 // 订单倒计时

	IfBuyerCancel bool `json:"if_buyer_cancel"                       ` // 是否可以取消

}

type OrderListOutput struct {
	Items   []*OrderVo // 列表
	Page    int        // 分页号码
	Total   int        // 总页数
	Records int        // 数据总数
	Size    int        // 单页数量
}

type LogisticsVo struct {
	StockBillId         string `json:"stock_bill_id"         dc:"出库单"       `
	OrderTrackingNumber string `json:"order_tracking_number"         dc:"物流单号"       `
	LogisticsId         uint   `json:"logistics_id"   dc:"商家物流编号"      ` // 商家物流编号
}

type CartItem struct {
	ItemId       uint64 `json:"item_id"         dc:"SKU编号"       ` // SKU编号
	ItemQuantity int    `json:"cart_quantity"   dc:"商品数量"      `   // 商品数量
}

type OrderAddInput struct {
	UserId      uint        `json:"user_id"        dc:"买家编号"           ` // 买家编号
	Items       []*CartItem `json:"items"          dc:"下单商品信息"        `  // 下单商品信息
	UdId        uint        `json:"ud_id"          dc:"地址编号"           ` // 地址编号
	FreightFlag bool        `json:"freight_flag"   dc:"是否需要运费"        `  // 是否需要运费
	OrderType   uint        `json:"order_type"     dc:"订单类型"           ` // 订单类型
	SrcOrderId  string      `json:"src_order_id"   dc:"供应商转单源订单"     `   // 供应商转单源订单
}

// CheckoutItemVo SKU信息
type CheckoutItemVo struct {
	CartId       uint64 `json:"cart_id" `       // 编号
	CartSelect   bool   `json:"cart_select" `   // 是否选中
	ItemId       uint64 `json:"item_id" `       // 商品编号
	CartQuantity uint   `json:"cart_quantity" ` // 购买商品数量
	CartType     uint   `json:"cart_type" `     // 购买类型(ENUM):1-购买; 2-积分兑换; 3-赠品; 4-活动促销
	ActivityId   uint   `json:"activity_id" `   // 活动Id-加价购等等加入购物的需要提示
	StoreId      uint   `json:"store_id" `      // 店铺编号
	ChainId      uint   `json:"chain_id" `      // 门店编号
}

// CheckoutInput 结算预览
type CheckoutInput struct {
	Items          []*CheckoutItemVo `json:"items"  `           // SKU信息
	UserId         uint              `json:"user_id" `          // 买家编号
	UserNickname   string            `json:"user_nickname" `    // 买家昵称
	UdId           uint              `json:"ud_id" `            // 地址编号 或者 地址数据Map
	ChainId        uint              `json:"chain_id" `         // 下单门店
	ActivityId     uint              `json:"activity_id" `      // 活动编号
	GbId           uint              `json:"gb_id" `            // 拼团编号
	DeliveryTypeId uint              `json:"delivery_type_id"`  // 配送方式:5-自提;10-物流配送
	PaymentTypeId  uint              `json:"payment_type_id"`   // 付款方式
	Message        map[uint]string   `json:"message" `          // 消息
	UserVoucherIds []uint            `json:"user_voucher_ids" ` // 优惠券
	UserInvoiceId  uint              `json:"user_invoice_id" `  // 发票

	// 生成参数
	OrderType  uint   `json:"order_type" `  // 订单类型
	CalFreight bool   `json:"cal_freight" ` // 是否需要计算运费
	SrcOrderId string `json:"src_order_id"` // 供应商转单源订单
}

// CheckoutOutput 结算输出结构体
type CheckoutOutput struct {
	Items               []StoreItemVo               `json:"items"`                 // 店铺信息
	OrderProductAmount  float64                     `json:"order_product_amount"`  // 商品原价总价
	OrderItemAmount     float64                     `json:"order_item_amount"`     // 单品优惠总价
	OrderFreightAmount  float64                     `json:"order_freight_amount"`  // 运费总价
	OrderDiscountAmount float64                     `json:"order_discount_amount"` // 优惠总额度
	OrderMoneyAmount    float64                     `json:"order_money_amount"`    // 商品最终总价
	OrderPointsAmount   float64                     `json:"order_points_amount"`   // 订单需要总积分
	OrderSpAmount       float64                     `json:"order_sp_amount"`       // 订单需要总积分2
	IsPaid              bool                        `json:"is_paid"`               // 是否支付完成
	UserId              uint                        `json:"user_id"`               // 买家编号
	UserDeliveryAddress *entity.UserDeliveryAddress `json:"user_delivery_address"` // 地址信息
	In                  *CheckoutInput              `json:"in"`                    // 输入参数
}

// CartAddInput 结构体
type CartAddInput struct {
	CheckoutItemVo
	UserId uint `json:"user_id"` // 买家编号
}

// OrderAddOutput 订单添加输出结构体
type OrderAddOutput struct {
	CheckoutOutput          // 继承自CheckoutOutput
	OrderIds       []string `json:"order_ids"`      // 订单编号
	MobileIsBind   bool     `json:"mobile_is_bind"` // 是否绑定手机
	GbId           uint     `json:"gb_id"`          // 拼团编号
}

// UserCartSelectInput struct
type UserCartSelectInput struct {
	Action     string `json:"action"`      // all:全部; store:店铺编号
	CartId     uint64 `json:"cart_id"`     // 编号
	StoreId    uint   `json:"store_id"`    // 店铺编号
	CartSelect bool   `json:"cart_select"` // 是否选中
	UserId     uint   `json:"user_id"`     // 用户编号
}

type PickingItem struct {
	ItemId           uint64  `json:"item_id"         dc:"SKU编号"       `    // SKU编号
	OrderItemId      uint64  `json:"order_item_id"   dc:"订单SKU编号"      `   // 商品数量
	BillItemQuantity uint    `json:"bill_item_quantity"   dc:"商品数量"      ` // 商品数量
	BillItemPrice    float64 `json:"bill_item_price"   dc:"商品单价"    `      // 商品单价

	ProductId uint64 `json:"product_id"         dc:"商品编号"       ` // 商品编号
}

type OrderPickingInput struct {
	OrderId         string         `json:"order_id"     dc:"订单编号"           `                   // 订单编号
	Items           []*PickingItem `json:"items"          dc:"出库商品信息"        `                  // 出库商品信息
	StockBillAmount float64        `json:"stock_bill_amount"    d:"0"     dc:"单据金额"           ` // 单据金额

	BillTypeId           uint `json:"bill_type_id"    d:"2700"     dc:"业务类型"           `    // 业务类型
	WarehouseId          uint `json:"warehouse_id"   d:"0"       dc:"地址编号"           `      // 地址编号
	StockTransportTypeId uint `json:"stock_transport_type_id"  d:"2751"  dc:"库存类型"        ` //

	PickingFlag bool `json:"picking_flag" d:"false"  dc:"出库标记:true-默认全出， false-指定出库"        ` //
}

type OrderShippingInput struct {
	OrderId             string `json:"order_id"     dc:"订单编号"           `                  // 订单编号
	StockBillId         uint   `json:"stock_bill_id"     dc:"出库单编号"           `            // 出库单编号
	SsId                uint   `json:"ss_id"        dc:"发货地址"           `                  // 发货地址
	LogisticsId         uint   `json:"logistics_id"        dc:"发货物流编号"           `         // 发货物流编号
	LogisticsTime       uint64 `json:"logistics_time"        dc:"发货时间"           `         // 发货时间
	OrderTrackingNumber uint   `json:"order_tracking_number"          dc:"运单号"           ` // 运单号
	LogisticsExplain    string `json:"logistics_explain"   dc:"备注"        `                //

	ShippingFlag bool `json:"shipping_flag" d:"false"  dc:"发货标记:true-默认全发， false-指定发货"        ` //
}

// OrderNumInput 订单信息表分页查询
type OrderNumInput struct {
	OrderTitle   string `json:"order_title"`    // 订单标题
	StoreId      uint   `json:"store_id"`       // 卖家店铺编号
	SubsiteId    uint   `json:"subsite_id"`     // 所属分站:0-总站
	UserId       uint   `json:"user_id"`        // 买家编号
	KindId       uint   `json:"kind_id"`        // 订单种类(ENUM): 1201-实物 ; 1202-教育类 ; 1203-电子卡券  ; 1204-其它
	OrderStateId uint   `json:"order_state_id"` // 订单状态(LIST):2011-待订单审核;2013-待财务审核;2020-待配货/待出库审核;2030-待发货;2040-已发货/待收货确认;2060-已完成/已签收;2070-已取消/已作废;
	UserNickname string `json:"user_nickname"`  // 买家昵称
	OrderStime   int64  `json:"order_stime"`    // 下单时间
	OrderEtime   int64  `json:"order_etime"`    // 下单时间
}

// OrderNumOutput 订单数量
type OrderNumOutput struct {
	FinNum                int64 `json:"fin_num"`                  // 完成订单数
	FinNumEntity          int64 `json:"fin_num_entity"`           // 完成订单数-实物
	FinNumV               int64 `json:"fin_num_v"`                // 完成订单数-虚拟
	CancelNum             int64 `json:"cancel_num"`               // 取消订单数
	CancelNumEntity       int64 `json:"cancel_num_entity"`        // 取消订单数-实物
	CancelNumV            int64 `json:"cancel_num_v"`             // 取消订单数-虚拟
	WaitShippingNum       int64 `json:"wait_shipping_num"`        // 待发货货订单数
	WaitShippingNumEntity int64 `json:"wait_shipping_num_entity"` // 待发货货订单数-实物
	WaitShippingNumV      int64 `json:"wait_shipping_num_v"`      // 待发货货订单数-虚拟
	ShipNum               int64 `json:"ship_num"`                 // 已发货订单数
	ShipNumEntity         int64 `json:"ship_num_entity"`          // 已发货订单数-实物
	ShipNumV              int64 `json:"ship_num_v"`               // 已发货订单数-虚拟
	WaitPayNum            int64 `json:"wait_pay_num"`             // 等待支付订单数
	WaitPayNumEntity      int64 `json:"wait_pay_num_entity"`      // 等待支付订单数-实物
	WaitPayNumV           int64 `json:"wait_pay_num_v"`           // 等待支付订单数-虚拟
	ReturningNum          int64 `json:"returning_num"`            // 售后订单数
}

// EvaluationVo 评价Vo
type EvaluationVo struct {
	UserId                    uint   `json:"user_id"                   dc:"用户编号"`    // 用户编号
	OrderItemEvaluationStatus []uint `json:"order_item_evaluation_status" dc:"评价状态"` // 评价状态
	OrderId                   string `json:"order_id"                 dc:"订单编号"`     // 订单编号
}

// EvaluationVo 评价Vo
type OrderCommentOutput struct {
	Items           []*OrderItemVo `json:"items"           dc:"评论项"`   // 评论项
	No              int            `json:"no"              dc:"否"`     // 否
	OrderEvaluation interface{}    `json:"order_evaluation" dc:"订单评价"` // 订单评价
	StoreInfo       StoreInfoVo    `json:"store_info"      dc:"店铺信息"`  // 店铺信息
	Yes             int            `json:"yes"             dc:"是"`     // 是
}

type OrderCommentVo struct {
	entity.OrderComment
	CommentImages []string `json:"comment_image"   ` // 评论上传的图片(DOT)
}

// OrderCommentItemReq 添加评论请求详细数据
type OrderCommentItemVo struct {
	OrderItemReturnNum           int      `json:"orderItemReturnNum"`           // 退货数量
	ItemPurchasePrice            int      `json:"itemPurchasePrice"`            // 购买价格
	BuyerId                      int      `json:"buyerId"`                      // 买家ID
	OrderItemReturnAgreeAmount   int      `json:"orderItemReturnAgreeAmount"`   // 同意退货金额
	OrderItemUnitPrice           int      `json:"orderItemUnitPrice"`           // 单价
	OrderItemInventoryLock       int      `json:"orderItemInventoryLock"`       // 库存锁定
	ItemSrcId                    int      `json:"itemSrcId"`                    // 来源ID
	CommentIsAnonymous           bool     `json:"commentIsAnonymous"`           // 是否匿名
	ProductId                    int      `json:"productId"`                    // 商品ID
	ActivityId                   int      `json:"activityId"`                   // 活动ID
	OrderItemImage               string   `json:"orderItemImage"`               // 订单商品图片
	ItemPurchaseRate             int      `json:"itemPurchaseRate"`             // 购买率
	OrderItemConfirmFile         string   `json:"orderItemConfirmFile"`         // 确认文件
	OrderItemSalerId             int      `json:"orderItemSalerId"`             // 卖家ID
	DesignFileImages             string   `json:"designFileImages"`             // 设计文件图片
	OrderItemReturnAgreeNum      int      `json:"orderItemReturnAgreeNum"`      // 同意退货数量
	OrderItemAmount              int      `json:"orderItemAmount"`              // 订单商品金额
	SrcOrderId                   string   `json:"srcOrderId"`                   // 源订单ID
	ItemPlatformPrice            int      `json:"itemPlatformPrice"`            // 平台价格
	ItemUnitPrice                int      `json:"itemUnitPrice"`                // 单位价格
	ActivityCode                 string   `json:"activityCode"`                 // 活动代码
	ItemChannelType              int      `json:"itemChannelType"`              // 渠道类型
	Version                      int      `json:"version"`                      // 版本
	OrderItemCommissionFeeRefund int      `json:"orderItemCommissionFeeRefund"` // 退款佣金费
	OrderItemId                  int      `json:"orderItemId"`                  // 订单商品ID
	OrderItemEvaluationStatus    int      `json:"orderItemEvaluationStatus"`    // 评价状态
	OrderItemConfirmStatus       int      `json:"orderItemConfirmStatus"`       // 确认状态
	ItemCostPrice                int      `json:"itemCostPrice"`                // 成本价
	OrderId                      string   `json:"orderId"`                      // 订单ID
	SpecInfo                     string   `json:"specInfo"`                     // 规格信息
	ItemSalesRate                int      `json:"itemSalesRate"`                // 销售率
	OrderItemCommissionFee       int      `json:"orderItemCommissionFee"`       // 佣金费
	OrderItemQuantity            int      `json:"orderItemQuantity"`            // 订单商品数量
	OrderItemReturnSubtotal      int      `json:"orderItemReturnSubtotal"`      // 退货小计
	OrderItemNote                string   `json:"orderItemNote"`                // 订单商品备注
	ItemVoucher                  int      `json:"itemVoucher"`                  // 代金券
	ItemUnitPoIntegers           int      `json:"itemUnitPoIntegers"`           // 单位积分
	CategoryId                   int      `json:"categoryId"`                   // 分类ID
	SpecId                       int      `json:"specId"`                       // 规格ID
	OrderItemDiscountAmount      int      `json:"orderItemDiscountAmount"`      // 订单商品折扣金额
	OrderItemAdjustFee           int      `json:"orderItemAdjustFee"`           // 调整费
	OrderGiveId                  int      `json:"orderGiveId"`                  // 赠品ID
	ItemUnitSp                   int      `json:"itemUnitSp"`                   // 单位SP
	StoreId                      int      `json:"storeId"`                      // 店铺ID
	OrderItemPoIntegersFee       int      `json:"orderItemPoIntegersFee"`       // 积分费用
	ItemId                       int      `json:"itemId"`                       // 商品ID
	OrderItemRedemptionVoucher   int      `json:"orderItemRedemptionVoucher"`   // 兑换券
	PolicyDiscountrate           int      `json:"policyDiscountrate"`           // 政策折扣率
	OrderItemFile                string   `json:"orderItemFile"`                // 订单文件
	ItemName                     string   `json:"itemName"`                     // 商品名称
	OrderItemSupplierSync        int      `json:"orderItemSupplierSync"`        // 供应商同步状态
	ActivityTypeId               int      `json:"activityTypeId"`               // 活动类型ID
	OrderItemCommissionRate      int      `json:"orderItemCommissionRate"`      // 佣金率
	OrderItemPaymentAmount       int      `json:"orderItemPaymentAmount"`       // 订单支付金额
	Content                      string   `json:"content"`                      // 内容
	CommentContent               string   `json:"commentContent"`               // 评论内容
	CommentImage                 []string `json:"commentImage"`                 // 评论图片
	CommentScores                uint     `json:"commentScores"`                // 评论分数
	Scores                       int      `json:"scores"`                       // 分数
	Comment                      string   `json:"comment"`                      // 评论
}
