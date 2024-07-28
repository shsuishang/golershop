package model

import "golershop.cn/internal/model/entity"

type TimelineInput struct {
	Stime int64 `json:"stime"  type:"GE"     ` // 开始时间
	Etime int64 `json:"etime"  type:"LE"     ` // 截止时间
}

type TimelineOutput struct {
	Time string `json:"time"     `
	Num  int    `json:"num"    `
}

// AnalyticsOrderInput 后台DashBoard
type AnalyticsOrderInput struct {
	TimelineInput

	CategoryId   int64  `json:"category_id" dc:"分类编号"`    // 分类编号
	ProductName  string `json:"product_name" dc:"产品名称"`   // 产品名称
	ProductId    int64  `json:"product_id" dc:"产品编号"`     // 产品编号
	OrderStateId []uint `json:"order_state_id" dc:"订单状态"` // 订单状态
	OrderIsPaid  []uint `json:"order_is_paid" dc:"支付状态"`  // 支付状态
	StoreId      uint   `json:"store_id" dc:"店铺编号"`       // 店铺编号
	StoreType    uint   `json:"store_type" dc:"店铺类型"`     // 店铺类型
	KindId       uint   `json:"kind_id" dc:"订单类型"`        // 订单类型
	UserId       uint   `json:"user_id" dc:"用户编号"`        // 用户编号
}

// AnalyticsOrderItemNumOutput 商品销售统计
type AnalyticsOrderItemNumOutput struct {
	ProductId                    int64   `json:"product_id"                v:"required#产品编号"`           // 产品编号
	ProductName                  string  `json:"product_name"              v:"required#商品名称"`           // 商品名称
	ItemId                       int64   `json:"item_id"                   v:"required#货品编号"`           // 货品编号
	ItemName                     string  `json:"item_name"                 v:"required#商品名称"`           // 商品名称
	CategoryId                   int     `json:"category_id"               v:"required#分类编号"`           // 分类编号
	ItemCostPrice                float64 `json:"item_cost_price"           v:"required#成本价"`            // 成本价
	ItemUnitPrice                float64 `json:"item_unit_price"           v:"required#商品价格单价"`         // 商品价格单价
	ItemUnitPoints               float64 `json:"item_unit_points"          v:"required#资源1单价"`          // 资源1单价
	ItemUnitSp                   float64 `json:"item_unit_sp"              v:"required#资源2单价"`          // 资源2单价
	OrderItemSalePrice           float64 `json:"order_item_sale_price"     v:"required#商品实际成交价单价"`      // 商品实际成交价单价
	OrderItemQuantity            int     `json:"order_item_quantity"       v:"required#商品数量"`           // 商品数量
	OrderItemImage               string  `json:"order_item_image"          v:"required#商品图片"`           // 商品图片
	OrderItemReturnNum           int     `json:"order_item_return_num"      v:"required#退货数量"`          // 退货数量
	OrderItemReturnSubtotal      float64 `json:"order_item_return_subtotal" v:"required#退款总额"`          // 退款总额
	OrderItemReturnAgreeAmount   float64 `json:"order_item_return_agree_amount" v:"required#退款金额:同意额度"` // 退款金额:同意额度
	OrderItemAmount              float64 `json:"order_item_amount"         v:"required#商品实际总金额"`        // 商品实际总金额
	OrderItemDiscountAmount      float64 `json:"order_item_discount_amount" v:"required#优惠金额"`          // 优惠金额
	OrderItemAdjustFee           float64 `json:"order_item_adjust_fee"     v:"required#手工调整金额"`         // 手工调整金额
	OrderItemPointsFee           float64 `json:"order_item_points_fee"     v:"required#积分费用"`           // 积分费用
	OrderItemPointsAdd           float64 `json:"order_item_points_add"     v:"required#赠送积分"`           // 赠送积分
	OrderItemPaymentAmount       float64 `json:"order_item_payment_amount" v:"required#实付金额"`           // 实付金额
	OrderItemEvaluationStatus    bool    `json:"order_item_evaluation_status" v:"required#评价状态(ENUM)"`  // 评价状态(ENUM)
	ActivityTypeId               int     `json:"activity_type_id"          v:"required#活动类型(ENUM)"`     // 活动类型(ENUM)
	ActivityId                   int     `json:"activity_id"               v:"required#促销活动ID"`         // 促销活动ID
	ActivityCode                 string  `json:"activity_code"             v:"required#礼包活动对应兑换码code"`  // 礼包活动对应兑换码code
	OrderItemCommissionRate      float64 `json:"order_item_commission_rate" v:"required#分佣金比例百分比"`      // 分佣金比例百分比
	OrderItemCommissionFee       float64 `json:"order_item_commission_fee"  v:"required#佣金"`            // 佣金
	OrderItemCommissionFeeRefund float64 `json:"order_item_commission_fee_refund" v:"required#退款佣金"`    // 退款佣金
	PolicyDiscountrate           float64 `json:"policy_discountrate"       v:"required#价格策略折扣率"`        // 价格策略折扣率
	OrderItemVoucher             float64 `json:"order_item_voucher"        v:"required#分配优惠券额度"`        // 分配优惠券额度
	OrderItemReduce              float64 `json:"order_item_reduce"         v:"required#分配满减额度"`         // 分配满减额度
	OrderItemReturnAgreeNum      int     `json:"order_item_return_agree_num" v:"required#同意退货数量"`       // 同意退货数量
	OrderItemAmountSum           float64 `json:"order_item_amount_sum"     v:"required#统计数量"`           // 统计数量
	Num                          int64   `json:"num"                       v:"required#统计数量"`           // 统计数量
}

type DashBoardTimelineOutput struct {
	OrderTimeLine []*TimelineOutput `json:"order_time_line"` // 最近一周订单增长数据
	UserTimeLine  []*TimelineOutput `json:"user_time_line"`  // 最近一周订单用户数据
	PtTimeLine    []*TimelineOutput `json:"pt_time_line" `   // 最近一周订单商品数据
	PayTimeLine   []*TimelineOutput `json:"pay_time_line"`   // 最近一周销售额增长数据
}

type AnalyticsNumVo struct {
	Current interface{} `json:"current"     ` //本周期
	Pre     interface{} `json:"pre"    `      //上个周期
	Daym2m  interface{} `json:"daym2m"    `   //周期环比
}

type AnalyticsNumOutput AnalyticsNumVo

type AmountVo struct {
	Time   string      `json:"time"`   // 时间
	Amount interface{} `json:"amount"` // 金额
}

// AdminDashBoardRes 仪表板看板对象
type AdminDashBoardVo struct {
	TradeAmount                      int64   `json:"trade_amount" description:"总交易额"`
	TradeAmountIncreaseRate          float64 `json:"trade_amount_increase_rate" description:"总交易额增长率"`
	OrderFinishNum                   int64   `json:"order_finish_num" description:"总成交"`
	OrderFinishNumIncreaseRate       float64 `json:"order_finish_num_increase_rate" description:"总成交增长率"`
	UserCertificationNum             int64   `json:"user_certification_num" description:"会员总数"`
	UserCertificationNumIncreaseRate float64 `json:"user_certification_num_increase_rate" description:"会员总数增长率"`
	OrderNum                         int64   `json:"order_num" description:"订单总量"`
	OrderNumIncreaseRate             float64 `json:"order_num_increase_rate" description:"订单总量增长率"`
}

// OrderItemNumTimelineInput 订单商品统计
type OrderItemNumTimelineInput struct {
	TimelineInput

	CategoryId  int64   `json:"category_id" dc:"分类编号"`  // 分类编号
	ProductName string  `json:"product_name" dc:"产品名称"` // 产品名称
	ProductId   int64   `json:"product_id" dc:"产品编号"`   // 产品编号
	ItemId      []int64 `json:"item_id" dc:"SKU编号"`     // SKU编号
	StoreId     uint    `json:"store_id" dc:"店铺编号"`     // 店铺编号
	StoreType   uint    `json:"store_type" dc:"店铺类型"`   // 店铺类型
	KindId      uint    `json:"kind_id" dc:"订单类型"`      // 订单类型
}

// AnalyticsProductInput 后台DashBoard
type AnalyticsProductInput struct {
	TimelineInput

	CategoryId     int64  `json:"category_id" dc:"分类编号"`      // 分类编号
	ProductName    string `json:"product_name" dc:"产品名称"`     // 产品名称
	StoreId        uint   `json:"store_id" dc:"店铺编号"`         // 店铺编号
	StoreType      uint   `json:"store_type" dc:"店铺类型"`       // 店铺类型
	ProductStateId uint   `json:"product_state_id" dc:"商品状态"` // 商品状态
}

type AccessItemTimelineInput struct {
	TimelineInput

	ItemId int64 `json:"item_id" dc:"SKU编号"` // SKU编号
}

// AnalyticsAccessItemOutput 商品浏览统计
type AnalyticsAccessItemOutput struct {
	ProductId     int64   `json:"product_id" dc:"产品编号"`      // 产品编号
	ProductName   string  `json:"product_name" dc:"商品名称"`    // 商品名称
	ItemId        int64   `json:"item_id" dc:"货品编号"`         // 货品编号
	ItemName      string  `json:"item_name" dc:"商品名称"`       // 商品名称
	ItemUnitPrice float64 `json:"item_unit_price" dc:"商品价格"` // 商品价格
	Num           int64   `json:"num" dc:"浏览量"`              // 浏览量
}

// AnalyticsReturnInput 后台DashBoard
type AnalyticsReturnInput struct {
	TimelineInput

	ProductName   string `json:"product_name" dc:"产品名称"`                                                                               // 产品名称
	ProductId     int64  `json:"product_id" dc:"产品编号"`                                                                                 // 产品编号
	ReturnStateId []uint `json:"return_state_id" dc:"卖家处理状态(ENUM): 3100-【客户】提交退单;3105-退单审核;3110-收货确认;3115-退款确认;3120-【客户】收款确认;3125-完成"` // 卖家处理状态
	StoreId       uint   `json:"store_id" dc:"店铺编号"`                                                                                   // 店铺编号
	StoreType     uint   `json:"store_type" dc:"店铺类型"`                                                                                 // 店铺类型
}

type ArticleBase struct {
	entity.ArticleBase
	ArticleTagList []*entity.ArticleTag `json:"article_tag_list" dc:"文章标签集合"` // 文章标签集合
	UserNickname   string               `json:"user_nickname" `               // 用户昵称
}
type ArticleBaseOutput struct {
	Items   []*ArticleBase // 列表
	Page    int            // 分页号码
	Total   int            // 总页数
	Records int            // 数据总数
	Size    int            // 单页数量
}
