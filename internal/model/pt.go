package model

import "golershop.cn/internal/model/entity"

// 商品SKU表
type ProductEditStockInput struct {
	ItemId       uint64 `json:"item_id,omitempty"`       // SKU编号
	ItemQuantity uint   `json:"item_quantity,omitempty"` // 商品库存
	BillTypeId   uint   `json:"bill_type_id,omitempty"`  // 业务类别(ENUM):2750-入库;2700-出库
}

type ItemListOutput struct {
	Items   []*ItemOutput
	Page    int // 分页号码
	Total   int // 总页数
	Records int // 数据总数
	Size    int // 单页数量

	Assists      []*ProductAssistOutput `json:"assists"`
	ActivityBase *entity.ActivityBase   `json:"activity_base"`
}

// ItemOutput 商品展示使用
type ItemOutput struct {
	entity.ProductItem

	ProductStateId  uint   `json:"product_state_id"`  // 商品状态(ENUM):1001-正常;1002-下架仓库中;1003-待审核; 1000-违规禁售
	ProductId       uint64 `json:"product_id"`        // 产品编号
	ProductNumber   string `json:"product_number"`    // SPU货号:货号
	ProductName     string `json:"product_name"`      // 产品名称
	ItemSpecName    string `json:"item_spec_name"`    // Spec名称
	ProductTips     string `json:"product_tips"`      // 商品卖点:商品广告词
	StoreId         uint   `json:"store_id"`          // 店铺编号
	ProductImage    string `json:"product_image"`     // 商品主图
	ProductVideo    string `json:"product_video"`     // 产品视频
	TransportTypeId uint   `json:"transport_type_id"` // 选择售卖区域:完成售卖区域及运费设置
	ProductBuyLimit uint   `json:"product_buy_limit"` // 每人限购
	ActivityItemNum uint   `json:"activity_item_num"` // 活动产品数量
}

type ProductAssistOutput struct {
	entity.ProductAssist
	Items []*entity.ProductAssistItem `json:"items"`
}

type SearchWordVo struct {
	DefaultSearchLabel string `json:"default_search_label"`
	DefaultSearchWords string `json:"default_search_words"`
}

// ProductOutput 商品输出结构体
type ProductOutput struct {
	entity.ProductIndex

	// Base
	ProductTips           string  `json:"product_tips" description:"商品卖点:商品广告词"`              // 商品卖点
	ProductImage          string  `json:"product_image" description:"商品主图"`                   // 商品主图
	ProductVideo          string  `json:"product_video" description:"产品视频 "`                  // 产品视频
	TransportTypeId       int     `json:"transport_type_id" description:"选择售卖区域:完成售卖区域及运费设置"` // 选择售卖区域
	ProductBuyLimit       int     `json:"product_buy_limit" description:"每人限购"`               // 每人限购
	ProductCommissionRate float64 `json:"product_commission_rate" description:"平台佣金比率"`       // 平台佣金比率

	// Info
	ProductSpec     string `json:"product_spec" description:"规格(JSON)-规格、规格值、goods_id  规格不需要全选就可以添加对应数据"`      // 规格
	ProductUniqId   string `json:"product_uniqid" description:"商品SKU(JSON):{'uniq_id':[item_id, price, url]}"` // 商品SKU
	ProductItemName string `json:"product_item_name" description:"商品SKU全名"`                                    // 商品SKU全名

	// Default SKU
	ItemId uint64                `json:"item_id" description:"默认SKU"` // 默认SKU
	Items  []*entity.ProductItem `json:"items" description:"SKU信息"`   // SKU信息

	// Activity information
	Activity interface{} `json:"activity" description:"活动信息"` // 活动信息
}

// ProductDetailInput 商品详情输入对象
type ProductDetailInput struct {
	UserId     uint   `json:"user_id"`     // 用户编号
	ItemId     uint64 `json:"item_id"`     // SKU编号
	DistrictId uint   `json:"district_id"` // 配送地区
	GbId       uint   `json:"gb_id"`       // 拼团活动编号
	GifgbagId  uint   `json:"gifgbag_id"`  // A+B组合套餐活动编号
	CutpriceId uint   `json:"cutprice_id"` // 砍价活动编号
}

// ProductDetailRes 商品详情响应结构体
type ProductDetailOutput struct {
	// Inherited fields from ProductOutput
	ProductOutput

	// Additional fields
	ProductDetail    string                    `json:"product_detail" description:"商品描述"`    // 商品描述
	ItemRow          *entity.ProductItem       `json:"item_row" description:"SKU"`           // SKU
	Image            *entity.ProductImage      `json:"image" description:"Image"`            // Image
	Freight          float64                   `json:"freight" description:"默认运费"`           // 默认运费
	DistrictList     []int                     `json:"district_list" description:"默认区域"`     // 默认区域
	IfStore          bool                      `json:"if_store" description:"是否可销售"`         // 是否可销售
	IsFavorite       bool                      `json:"is_favorite" description:"是否收藏"`       // 是否收藏
	LastComments     []*entity.ProductComment  `json:"last_comments" description:"最后几条评论"`   // 最后几条评论
	LastComment      *entity.ProductComment    `json:"last_comment" description:"最后一条评论"`    // 最后一条评论
	Assists          []*ProductAssistOutput    `json:"assists" description:"分类辅助属性"`         // 分类辅助属性
	Contracts        []*entity.ContractType    `json:"contracts" description:"服务"`           // 服务
	Markets          []interface{}             `json:"markets" description:"商圈"`             // 商圈
	ProductCategorys []*entity.ProductCategory `json:"product_categorys" description:"商品分类"` // 商品分类
}
