package model

import (
	"golershop.cn/internal/model/entity"
)

// ActivityInfoVo 活动信息结构体
type ActivityInfoVo struct {
	entity.ActivityItem
	ActivityBase *entity.ActivityBase `json:"activity_base"`
}

// ActivitysVo 店铺及商品信息结构体
type ActivitysVo struct {
	Gift      []interface{} `json:"gift"`
	Reduction []interface{} `json:"reduction"`
	Multple   []interface{} `json:"multple"`
	Bargains  []interface{} `json:"bargains"`
}

// ActivityRuleVo 活动规则结构体
type ActivityRuleVo struct {
	Rule          []RuleVo       `json:"rule"`
	Requirement   RequirementVo  `json:"requirement"`
	Voucher       VoucherVo      `json:"voucher"`
	Groupbooking  GroupbookingVo `json:"groupbooking"`
	Giftbag       GiftbagVo      `json:"giftbag"`
	GroupBuyStore GroupbuyVo     `json:"group_buy_store"`
	Marketing     MarketingVo    `json:"marketing"`
	Lottery       LotteryVo      `json:"lottery"`
	Cutprice      CutpriceVo     `json:"cutprice"`
	Popup         PopupVo        `json:"popup"`
}

// RuleVo 规则结构体
type RuleVo struct {
	Amount         float64      `json:"amount"`     // 总额
	MaxAmount      float64      `json:"max_amount"` // 最大总额
	Percent        float64      `json:"percent"`    // 百分比
	Num            uint         `json:"num"`        // 总数量
	MaxNum         uint         `json:"max_num"`    // 最多数量
	PointsStandard uint         `json:"points_standard"`
	PointsDouble   uint         `json:"points_double"`
	Item           []*ItemNumVo `json:"item"` // 产品及数量
}

// ItemNumVo 结构体
type ItemNumVo struct {
	ItemId uint64  `json:"item_id"` // SKU编号
	Num    uint    `json:"num"`     // 数量
	Price  float64 `json:"price"`   // 价格
}

// RequirementVo 结构体
type RequirementVo struct {
	Buy    *BuyVo    `json:"buy"`
	Points *PointsVo `json:"points"` // 积分
}

// VoucherVo 结构体
type VoucherVo struct {
	VoucherStartDate    int64   `json:"voucher_start_date"`    // 优惠券开始时间
	VoucherEndDate      int64   `json:"voucher_end_date"`      // 优惠券失效日期
	VoucherPrice        float64 `json:"voucher_price"`         // 优惠券价格
	VoucherQuantityUse  uint    `json:"voucher_quantity_use"`  // 已领取张数
	VoucherImage        string  `json:"voucher_image"`         // 优惠券图片
	VoucherQuantity     uint    `json:"voucher_quantity"`      // 优惠券数量
	VoucherQuantityFree int     `json:"voucher_quantity_free"` // 可领数量
	VoucherPreQuantity  uint    `json:"voucher_pre_quantity"`  // 优惠券限制
	VoucherProductLimit uint    `json:"voucher_product_limit"` // 适用商品(ENUM):1-全部商品可用;2-指定商品可用
}

// GroupbookingVo 结构体
type GroupbookingVo struct {
	GroupBuyLimit       int `json:"group_buy_limit"`       // 拼团次数限制
	GroupRemainQuantity int `json:"group_remain_quantity"` // 剩余团数
	GroupDaysLimit      int `json:"group_days_limit"`      // 成团天数限制
	GroupQuantity       int `json:"group_quantity"`        // 拼团人数
}

// GiftbagVo 结构体
type GiftbagVo struct {
	ActivityBagCategory string       `json:"activity_bag_category"` // 礼包分类
	GiftbagImage        string       `json:"giftbag_image"`         // 活动海报
	GiftbagZuImage      string       `json:"giftbag_zu_image"`      // 活动主图
	ActivityVideo       string       `json:"activity_video"`        // 商品介绍
	GiftbagQuantity     uint         `json:"giftbag_quantity"`      // 库存数量
	GiftbagAmount       float64      `json:"giftbag_amount"`        // 售卖金额
	TransportTypeId     uint         `json:"transport_type_id"`     // 运费设置
	Items               []*ItemNumVo `json:"items"`                 // 产品及数量
}

// GroupbuyVo 团购结构体
type GroupbuyVo struct {
	// Define the structure based on your actual requirements.
	// You may need to replace it with the correct types.

	GroupBuyLimit     int     `json:"group_buy_limit"    dc:"每人限购"` // 每人限购
	GroupQuantity     int     `json:"group_quantity"    dc:""`      // 团购数量
	GroupSaleQuantity int     `json:"group_sale_quantity" dc:""`    // 团购销售数量
	GroupSalePrice    float64 `json:"group_sale_price"   dc:""`     // 团购销售价格
	ProductId         uint64  `json:"product_id"         dc:"产品编号"` // 产品编号
	ProductUnitPrice  float64 `json:"product_unit_price" dc:""`     // 产品单价
	ProductImage      string  `json:"product_image"      dc:"商品主图"` // 商品主图

}

// MarketingVo 结构体
type MarketingVo struct {
	StartJoinTime       int64  `json:"start_join_time"`       // 参加开始时间
	EndJoinTime         int64  `json:"end_join_time"`         // 参加截止时间
	ActivityAddress     string `json:"activity_address"`      // 举办地址
	ActivitySponsor     string `json:"activity_sponsor"`      // 主办方
	ActivityCoSponsor   string `json:"activity_co_sponsor"`   // 主办方
	ContactOrganizer    string `json:"contact_organizer"`     // 联系人
	ContactPhone        int    `json:"contact_phone"`         // 联系人电话
	ActivityImage       string `json:"activity_image"`        // 活动主图
	GuestImage          string `json:"guest_image"`           // 嘉宾介绍
	ActivityProcess     string `json:"activity_process"`      // 活动流程
	ActivityDetailIntro string `json:"activity_detail_intro"` // 活动详细规则
}

// LotteryVo 结构体
type LotteryVo struct {
	LotterySubtitle        string `json:"lottery_subtitle"`          // 抽奖副标题
	LotteryType            int    `json:"lottery_type"`              // 2砸金蛋，1大转盘
	LotteryImage           string `json:"lottery_image"`             // 抽奖主题图片
	LotteryNotAwardsRemark string `json:"lottery_not_awards_remark"` // ""
	LotteryProbability     string `json:"lottery_probability"`       // ""
	LotteryDayTimes        int    `json:"lottery_day_times"`         // ""
	LotteryShareAddTimes   int    `json:"lottery_share_add_times"`   // ""
	LotteryMaxAwardsTimes  int    `json:"lottery_max_awards_times"`  // ""
	LotteryUsedTimes       int    `json:"lottery_used_times"`        // 抽奖次数
	LotteryAwardsTimes     int    `json:"lottery_awards_times"`      // 抽奖中奖次数
	LotteryAward           string `json:"lottery_award"`             // 奖品信息
}

// CutpriceVo 结构体
type CutpriceVo struct {
	CutDownFixedPrice    float64      `json:"cut_down_fixed_price"`     // 固定砍价价格
	CutDownMaxPrice      float64      `json:"cut_down_max_price"`       // 砍价最高范围
	CutDownMinPrice      float64      `json:"cut_down_min_price"`       // 砍价最低范围
	CutDownType          uint         `json:"cut_down_type"`            // 砍价方式1.固定砍价价格2.范围
	CutDownMinLimitPrice float64      `json:"cut_down_min_limit_price"` // 砍价底价
	CutDownUserNum       uint         `json:"cut_down_user_num"`        // 砍价人数
	CutpriceImage        string       `json:"cutprice_image"`           // 活动主图
	CutpriceQuantity     uint         `json:"cutprice_quantity"`        // 库存数量
	CutDaysLimit         uint         `json:"cut_days_limit"`           // 砍价天数限制
	Items                []*ItemNumVo `json:"items"`                    // 产品及数量
	TotalPrice           float64      `json:"total_price"`              // 砍价商品总价
}

// PopupVo 结构体
type PopupVo struct {
	PopUpImage string `json:"pop_up_image"` // 弹窗图片
	PopUpUrl   string `json:"pop_up_url"`   // 弹窗网址
	PopUpType  int    `json:"pop_up_type"`  // 弹窗活动类型(ENUM):0-新人礼包;1-其他活动
}

// BuyVo 结构体
type BuyVo struct {
	Item     []uint64 `json:"item"`     // 商品
	Subtotal float64  `json:"subtotal"` // 总额
	Num      int      `json:"num"`      // 总数量
}

// PointsVo 结构体
type PointsVo struct {
	Needed int `json:"needed"` // 所需积分
}

type ActivityOutput struct {
	entity.ActivityBase
	StoreName            string           `json:"store_name"`              //店铺名称
	ActivityRuleJson     *ActivityRuleVo  `json:"activity_rule_json"`      //活动规则
	IfGain               bool             `json:"if_gain"`                 //是否领取
	ItemNumber           int              `json:"item_number"`             //优惠套装数量
	Item                 []*ProductItemVo `json:"item"`                    //商品信息集合
	UseLevel             string           `json:"use_level"`               //会员等级
	ProductItemName      string           `json:"product_item_name"`       //折扣商品
	ActivityUseLevelName string           `json:"activity_use_level_name"` //使用等级名称(DOT)

	RemainQuantity int `json:"remain_quantity"` //剩余库存
}

type ActivityBaseRes ActivityOutput

//type ActivityListInput struct {
//	do.ActivityBaseListInput
//}

type ActivityListOutput struct {
	Items   []*ActivityOutput `json:"items"    dc:"分页数据内容"`
	Page    int               `json:"page"`    // 分页号码
	Total   int               `json:"total"`   // 总页数
	Records int               `json:"records"` // 数据总数
	Size    int               `json:"size"`    // 单页数量
}

type ActivityGroupookingVo struct {
	entity.ActivityGroupbooking
	UserNickname string `json:"user_nickname"` // 买家昵称
	UserAvatar   string `json:"user_avatar"`   // 用户头像
}

type ActivityGroupbookingHistoryVo struct {
	entity.ActivityGroupbookingHistory
	UserNickname string `json:"user_nickname"` // 买家昵称
}
type ActivityCutpriceVo struct {
	entity.ActivityCutprice
	UserNickname string `json:"user_nickname"` // 买家昵称
}
