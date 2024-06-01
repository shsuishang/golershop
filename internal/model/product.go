package model

import (
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
)

type ISpecItemVo struct {
	Id   uint   `json:"id"      `   // 规格值编号
	Name string `json:"name"      ` // 规格值名称
}

type ISpecVo struct {
	Id         uint        `json:"id"      `         // 规格编号
	Name       string      `json:"name"      `       // 规格名称
	SpecFormat string      `json:"spec_format"     ` // 商品主图
	Item       ISpecItemVo `json:"item"     `        // 规格值
}

type ProductListVo struct {
	entity.ProductIndex

	// base
	ProductTips           string  `json:"product_tips" description:"商品卖点:商品广告词"`              // 商品卖点
	ProductImage          string  `json:"product_image" description:"商品主图"`                   // 商品主图
	ProductVideo          string  `json:"product_video" description:"产品视频 "`                  // 产品视频
	TransportTypeId       uint    `json:"transport_type_id" description:"选择售卖区域:完成售卖区域及运费设置"` // 选择售卖区域
	ProductBuyLimit       uint    `json:"product_buy_limit" description:"每人限购"`               // 每人限购
	ProductCommissionRate float64 `json:"product_commission_rate" description:"平台佣金比率"`       // 平台佣金比率

	// info
	ProductSpec     string `json:"product_spec" description:"规格(JSON)-规格、规格值、goods_id  规格不需要全选就可以添加对应数据[{'id' : spec_id, 'name' : spec_name, 'item':[{'id' : spec_item_id, 'name' : spec_item_name}, {'id' : spec_item_id, 'name' : spec_item_name}]},{'id' : spec_id, 'name' : spec_name, 'item':[{'id' : spec_item_id, 'name' : spec_item_name}, {'id' : spec_item_id, 'name' : spec_item_name}]}]"` // 规格
	ProductUniqid   string `json:"product_uniqid" description:"商品SKU(JSON):{'uniq_id':[item_id, price, url]}"`                                                                                                                                                                                                                                                                                         // 商品SKU
	ProductItemName string `json:"product_item_name" description:"商品SKU全名"`                                                                                                                                                                                                                                                                                                                            // 商品SKU全名

	// others
	ItemId   uint64                `json:"item_id" description:"默认SKU"` // 默认SKU
	Items    []*entity.ProductItem `json:"items" description:"SKU信息"`   // SKU信息
	Activity interface{}           `json:"activity" description:"活动信息"` // 活动信息

}

type ProductListOutput struct {
	Items   []*ProductListVo `json:"items"    dc:"商品列表信息"` // 列表
	Page    int              `json:"page"`                 // 分页号码
	Total   int              `json:"total"`                // 总页数
	Records int              `json:"records"`              // 数据总数
	Size    int              `json:"size"`                 // 单页数量
}

type SaveProductInput struct {
	*do.ProductBase
	*do.ProductIndex
	*do.ProductInfo
	*do.ProductValidPeriod

	ProductItems  []*do.ProductItem
	ProductImages []*do.ProductImage
}

// ProductItemVo 商品下单使用结构体
type ProductItemVo struct {
	entity.ProductItem
	ActivityInfo *ActivityInfoVo `json:"activity_info"` // 活动信息

	ProductName            string        `json:"product_name"`             // SPU商品名称
	ProductTips            string        `json:"product_tips"`             // 商品卖点:商品广告词
	ProductImage           string        `json:"product_image"`            // 图片信息
	TransportTypeId        uint          `json:"transport_type_id"`        // 运费模板
	ProductTags            string        `json:"product_tags"`             // 商品标签
	ProductStateId         uint          `json:"product_state_id"`         // 商品状态
	ProductInventoryLock   uint          `json:"product_inventory_lock"`   // 冻结库存
	KindId                 uint          `json:"kind_id"`                  // 类型编号
	CartId                 uint64        `json:"cart_id"`                  // 购物车编号
	AvailableQuantity      uint          `json:"available_quantity"`       // 可用库存
	CartQuantity           uint          `json:"cart_quantity"`            // 购物数量
	CartSelect             bool          `json:"cart_select"`              // 是否选中
	IsOos                  bool          `json:"is_oos"`                   // 超出配送区域
	IsOnSale               bool          `json:"is_on_sale"`               // 销售中
	ProductDistEnable      bool          `json:"product_dist_enable"`      // 三级分销允许分销(BOOL):1-启用分销;0-禁用分销
	ProductCommissionRate  float64       `json:"product_commission_rate"`  // 平台佣金比率
	ItemPolicyDiscountrate float64       `json:"item_policy_discountrate"` // 折扣率
	ItemDiscountAmount     float64       `json:"item_discount_amount"`     // 优惠总额
	ItemSubtotal           float64       `json:"item_subtotal"`            // 金额小计
	ItemPointsSubtotal     float64       `json:"item_points_subtotal"`     // 应付积分小计
	ItemVoucher            float64       `json:"item_voucher"`             // 分配优惠券额度
	ItemReduction          float64       `json:"item_reduction"`           // 分配满减额度
	ItemRatePrice          float64       `json:"item_rate_price"`          // 底价
	PulseGiftCart          []interface{} `json:"pulse_gift_cart"`          // 礼品-废弃
	PulseReduction         []interface{} `json:"pulse_reduction"`
	PulseMultple           []interface{} `json:"pulse_multple"`
	PulseBargainsCart      []interface{} `json:"pulse_bargains_cart"`
	PulseBargains          []interface{} `json:"pulse_bargains"`
}

// 移动端装修Vo
type PageMobileVo struct {
	PageId         uint64 `json:"Id"          `      // 页面编号
	PageName       string `json:"PageTitle"        ` // 页面名称
	StoreId        uint   `json:"StoreId"         `  // 所属店铺
	UserId         uint   `json:"user_id"          ` // 所属用户
	SubsiteId      uint   `json:"subsite_id"       ` // 所属分站:0-总站
	PageBuildin    uint   `json:"page_buildin"     ` // 是否内置(BOOL):0-否;1-是
	PageType       uint   `json:"page_type"        ` // 类型(ENUM):1-WAP;2-PC;3-APP
	PageTpl        uint   `json:"page_tpl"         ` // 页面布局模板
	AppId          uint   `json:"AppId"           `  // 所属APP
	PageCode       string `json:"PageCode"        `  // 页面代码
	PageNav        string `json:"PageNav"         `  //
	PageConfig     string `json:"PageConfig"      `  //
	PageShareTitle string `json:"ShareTitle" `       //
	PageShareImage string `json:"ShareImg" `         //
	PageQrcode     string `json:"PageQRCode"      `  //
	PageIndex      bool   `json:"IsHome"       `     // 是否首页(BOOL):0-非首页;1-首页
	PageGb         bool   `json:"IsGb"          `    // 拼团首页(BOOL):0-非首页;1-首页
	PageActivity   bool   `json:"IsActivity"    `    // 活动首页(BOOL):0-非首页;1-首页
	PagePoint      bool   `json:"IsPoint"       `    // 积分首页(BOOL):0-非首页;1-首页
	PageGbs        bool   `json:"page_gbs"         ` // 团购首页(BOOL):0-非首页;1-首页
	PagePackage    bool   `json:"page_package"     ` // 组合套餐(BOOL):0-非首页;1-首页
	PagePfgb       bool   `json:"page_pfgb"        ` // 批发团购首页(BOOL):0-非首页;1-首页
	PageSns        bool   `json:"IsSns"         `    // 社区(BOOL):0-非首页;1-首页
	PageArticle    bool   `json:"IsArticle"     `    // 资讯(BOOL):0-非首页;1-首页
	PageZerobuy    bool   `json:"page_zerobuy"     ` // 零元购区(BOOL):0-否;1-是
	PageHigharea   bool   `json:"page_higharea"    ` // 高额返区(BOOL):0-否;1-是
	PageTaday      bool   `json:"page_taday"       ` // 今日爆款(BOOL):0-否;1-是
	PageEveryday   bool   `json:"page_everyday"    ` // 每日好店(BOOL):0-否;1-是
	PageSecondkill bool   `json:"SecondKill"  `      // 整点秒杀(BOOL):0-否;1-是
	PageSecondday  bool   `json:"page_secondday"   ` // 天天秒淘(BOOL):0-否;1-是
	PageRura       bool   `json:"page_rura"        ` // 设置土特产(BOOL):0-否;1-是
	PageLikeyou    bool   `json:"page_likeyou"     ` // 用户页banner(BOOL):0-否;1-是
	PageExchange   bool   `json:"page_exchange"    ` // 兑换专区(BOOL):0-否;1-是
	PageNew        bool   `json:"page_new"         ` // 新品首发(BOOL):0-否;1-是
	PageNewperson  bool   `json:"page_newperson"   ` // 新人优惠(BOOL):0-否;1-是
	PageUpgrade    bool   `json:"IsUpgrade"     `    // 升级VIP(BOOL):0-否;1-是
	PageMessage    bool   `json:"page_message"     ` // 信息发布(BOOL):0-否;1-是
	PageRelease    bool   `json:"IsRelease"       `  // 是否发布
}

type ProductDateOutput struct {
	ProductBase  *entity.ProductBase    `json:"product_base"    dc:"商品基础信息"`
	ProductIndex *entity.ProductIndex   `json:"product_index"    dc:"商品索引"`
	ProductInfo  *entity.ProductInfo    `json:"product_info"    dc:"商品信息"`
	ProductItem  []*entity.ProductItem  `json:"product_item"    dc:"商品SKU"`
	ProductImage []*entity.ProductImage `json:"product_image"    dc:"商品图片表"`
}
