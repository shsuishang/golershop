package pt

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/entity"
)

type ItemListReq struct {
	g.Meta `path:"/front/pt/product/listItem" tags:"商品" method:"get" summary:"商品SKU列表接口"`
	ml.BaseList

	ItemId string `json:"item_id"    type:"IN_STR"           ` // 商品编号-SKU编号

	ItemName    string `json:"item_name" type:"LIKE"     ` // 副标题(DOT):SKU名称
	ItemNumber  string `json:"item_number"          `      // SKU商家编码:SKU商家编码为非必填项，若不填写，系统会自动生成一个SKU商家编码。
	ItemBarcode string `json:"item_barcode"         `      // 条形码
	ItemEnable  uint   `json:"item_enable"          `      // 是否启用(LIST):1001-正常;1002-下架仓库中;1000-违规禁售
	ActivityId  uint   `json:"activity_id"          `      // 活动编号

	ProductId              uint64  `json:"product_id"                  `                         // 产品编号:定为SPU编号
	ProductNumber          string  `json:"product_number"              `                         // SPU商家编码:货号
	ProductName            string  `json:"product_name" type:"LIKE"    `                         // 产品名称
	ProductNameIndex       string  `json:"keywords" type:"LIKE"     `                            // 名称索引关键字(DOT)
	StoreId                uint    `json:"store_id"                    `                         // 店铺编号
	StoreIsOpen            bool    `json:"store_is_open"               `                         // 店铺状态(BOOL):0-关闭;1-运营中
	StoreType              uint    `json:"store_type"                  `                         // 店铺类型(ENUM): 1-卖家店铺; 2-供应商店铺
	StoreCategoryIds       string  `json:"store_category_ids" type:"IN_STR"`                     // 店铺分类(DOT)
	CategoryId             string  `json:"category_id"              `                            // 商品分类
	CategoryIds            []uint  `json:"category_id"  type:"IN"               `                // 商品分类
	TypeId                 uint    `json:"type_id"                     `                         // 类型编号:冗余检索
	ProductQuantity        uint    `json:"product_quantity"            `                         // 商品库存:冗余计算
	ProductWarnQuantity    uint    `json:"product_warn_quantity"       `                         // 预警数量
	BrandId                string  `json:"brand_id"   type:"IN_STR"                  `           // 品牌编号
	ProductServiceTypeIds  string  `json:"product_service_type_ids" type:"FIND_IN_SET_STR"     ` // 售后服务(DOT)
	ProductStateId         uint    `json:"product_state_id"            `                         // 商品状态(ENUM):1001-正常;1002-下架仓库中;1003-待审核; 1000-违规禁售
	ProductSaleDistrictIds string  `json:"product_sale_district_ids"  type:"FIND_IN_SET_STR"  `  // 销售区域(DOT): district_id=1000全部区域
	ProductVerifyId        uint    `json:"product_verify_id"           `                         // 商品审核(ENUM):3001-审核通过;3002-审核中;3000-审核未通过
	ProductIsInvoices      bool    `json:"product_is_invoices"         `                         // 是否开票(BOOL): 1-是; 0-否
	ProductIsReturn        bool    `json:"product_is_return"           `                         // 允许退换货(BOOL): 1-是; 0-否
	ProductIsRecommend     bool    `json:"product_is_recommend"        `                         // 商品推荐(BOOL):1-是; 0-否
	ProductStockStatus     uint    `json:"product_stock_status"        `                         // 缺货状态(ENUM):1-有现货;2-预售商品;3-缺货;4-2至3天
	KindId                 uint    `json:"kind_id"                     `                         // 商品种类:1201-实物;1202-虚拟
	ActivityTypeIds        string  `json:"activity_type_ids"  type:"FIND_IN_SET_STR"           ` // 参与活动(DOT)
	ContractTypeIds        string  `json:"contract_type_ids"  type:"FIND_IN_SET_STR"         `   // 消费者保障(DOT):由店铺映射到商品
	ProductAssistData      string  `json:"product_assist_data"         `                         // 辅助属性值列(DOT):assist_item_id每个都不用 , setFilter(tagid, array(2,3,4));是表示含有标签值2,3,4中的任意一个即符合筛选，这里是or关系。 setFilter(‘tagid’, array(2)); setFilter(‘tagid’, array(3)); 形成and关系| msyql where FIND_IN_SET('1', product_assist_data)
	ProductUnitPriceMin    float64 `json:"product_unit_price_min" type:"GE"       `              // 最低单价
	ProductUnitPriceMax    float64 `json:"product_unit_price_max" type:"LE"       `              // 最高单价
	ProductUnitPointsMin   float64 `json:"product_unit_points_min" type:"GE"     `               // 商品积分
	ProductUnitPointsMax   float64 `json:"product_unit_points_max" type:"LE"     `               // 商品积分

	ProductRegionDistrictIds string  `json:"product_region_district_ids" `                          // 所属区域(DOT)
	ProductFreight           float64 `json:"product_freight"             `                          // 运费:包邮为0
	ProductTags              string  `json:"product_tags"   type:"FIND_IN_SET_STR"                ` // 商品标签(DOT)
	StoreIsSelfsupport       bool    `json:"store_is_selfsupport"        `                          // 是否自营(BOOL):1-自营;0-非自营
	ProductDistEnable        bool    `json:"product_dist_enable"         `                          // 三级分销允许分销(BOOL):1-启用分销;0-禁用分销

	ProductSrcId       int64   `json:"product_src_id"              ` // 产品来源编号
	MarketCategoryId   string  `json:"market_category_id"          ` // 所属商圈(DOT)
	StoreLatitude      float64 `json:"store_latitude"              ` // 纬度
	StoreLongitude     float64 `json:"store_longitude"             ` // 经度
	ProductIsVideo     uint    `json:"product_is_video"            ` // 是否视频(BOOL):1-有视频;0-无视频
	ProductTransportId string  `json:"product_transport_id"        ` // 配送服务(ENUM):1001-快递发货;1002-到店自提;1003-上门服务
	SubsiteId          uint    `json:"subsite_id"                  ` // 所属分站:0-总站
	ProductFrom        uint    `json:"product_from"                ` // 商品来源(ENUM):1000-发布;1001-天猫;1002-淘宝;1003-阿里巴巴;1004-京东;
}

type ItemListRes struct {
	model.ItemListOutput
}

type CategoryListReq struct {
	g.Meta `path:"/front/pt/product/listCategory" tags:"商品" method:"get" summary:"商品分类列表接口"`
	ml.BaseList

	CategoryParentId uint   `json:"category_parent_id" d:"0" description:"分类父编号"`
	CategoryName     string `json:"category_name"       description:"分类名称"`
	CategoryIsEnable bool   `json:"category_is_enable" d:"true"  description:"是否启用(BOOL):0-不显示;1-显示"`
}

type CategoryListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type CategoryTreeReq struct {
	g.Meta `path:"/front/pt/product/treeCategory" tags:"商品" method:"get" summary:"商品分类列表接口"`
	ml.BaseList

	CategoryParentId uint   `json:"category_parent_id" d:"0" description:"分类父编号"`
	CategoryName     string `json:"category_name"       description:"分类名称"`
	CategoryIsEnable bool   `json:"category_is_enable" d:"true"  description:"是否启用(BOOL):0-不显示;1-显示"`
}

type CategoryTreeRes []*model.CategoryTreeNode

type ListReq struct {
	g.Meta `path:"/front/pt/product/list" tags:"商品" method:"get" summary:"商品列表接口"`
	ml.BaseList

	ProductId                uint64  `json:"product_id"                  `                          // 产品编号:定为SPU编号
	ProductNumber            string  `json:"product_number"              `                          // SPU商家编码:货号
	ProductName              string  `json:"product_name" type:"LIKE"    `                          // 产品名称
	ProductNameIndex         string  `json:"keywords" type:"LIKE"     `                             // 名称索引关键字(DOT)
	StoreId                  uint    `json:"store_id"                    `                          // 店铺编号
	StoreIsOpen              bool    `json:"store_is_open"               `                          // 店铺状态(BOOL):0-关闭;1-运营中
	StoreType                uint    `json:"store_type"                  `                          // 店铺类型(ENUM): 1-卖家店铺; 2-供应商店铺
	StoreCategoryIds         string  `json:"store_category_ids" type:"IN_STR"`                      // 店铺分类(DOT)
	CategoryId               []uint  `json:"category_id" type:"IN"                 `                // 商品分类
	TypeId                   uint    `json:"type_id"                     `                          // 类型编号:冗余检索
	ProductQuantity          uint    `json:"product_quantity"            `                          // 商品库存:冗余计算
	ProductWarnQuantity      uint    `json:"product_warn_quantity"       `                          // 预警数量
	BrandId                  string  `json:"brand_id"   type:"IN_STR"                  `            // 品牌编号
	ProductServiceTypeIds    string  `json:"product_service_type_ids" type:"FIND_IN_SET_STR"     `  // 售后服务(DOT)
	ProductStateId           uint    `json:"product_state_id"            `                          // 商品状态(ENUM):1001-正常;1002-下架仓库中;1003-待审核; 1000-违规禁售
	ProductSaleDistrictIds   string  `json:"product_sale_district_ids"  type:"FIND_IN_SET_STR"  `   // 销售区域(DOT): district_id=1000全部区域
	ProductVerifyId          uint    `json:"product_verify_id"           `                          // 商品审核(ENUM):3001-审核通过;3002-审核中;3000-审核未通过
	ProductIsInvoices        bool    `json:"product_is_invoices"         `                          // 是否开票(BOOL): 1-是; 0-否
	ProductIsReturn          bool    `json:"product_is_return"           `                          // 允许退换货(BOOL): 1-是; 0-否
	ProductIsRecommend       bool    `json:"product_is_recommend"        `                          // 商品推荐(BOOL):1-是; 0-否
	ProductStockStatus       uint    `json:"product_stock_status"        `                          // 缺货状态(ENUM):1-有现货;2-预售商品;3-缺货;4-2至3天
	KindId                   uint    `json:"kind_id"                     `                          // 商品种类:1201-实物;1202-虚拟
	ActivityTypeIds          string  `json:"activity_type_ids"  type:"FIND_IN_SET_STR"           `  // 参与活动(DOT)
	ContractTypeIds          string  `json:"contract_type_ids"  type:"FIND_IN_SET_STR"         `    // 消费者保障(DOT):由店铺映射到商品
	ProductAssistData        string  `json:"product_assist_data"     type:"FIND_IN_SET_STR"      `  // 辅助属性值列(DOT):assist_item_id每个都不用 , setFilter(tagid, array(2,3,4));是表示含有标签值2,3,4中的任意一个即符合筛选，这里是or关系。 setFilter(‘tagid’, array(2)); setFilter(‘tagid’, array(3)); 形成and关系| msyql where FIND_IN_SET('1', product_assist_data)
	ProductUnitPriceMin      float64 `json:"product_unit_price_min" type:"GE"       `               // 最低单价
	ProductUnitPriceMax      float64 `json:"product_unit_price_max" type:"LE"       `               // 最高单价
	ProductUnitPointsMin     float64 `json:"product_unit_points_min" type:"GE"     `                // 商品积分
	ProductUnitPointsMax     float64 `json:"product_unit_points_max" type:"LE"     `                // 商品积分
	ProductSaleNum           uint    `json:"product_sale_num"            `                          // 销售数量
	ProductFavoriteNum       uint    `json:"product_favorite_num"        `                          // 收藏数量
	ProductClick             uint    `json:"product_click"               `                          // 点击数量
	ProductEvaluationNum     uint    `json:"product_evaluation_num"      `                          // 评价次数
	ProductRegionDistrictIds string  `json:"product_region_district_ids" `                          // 所属区域(DOT)
	ProductFreight           float64 `json:"product_freight"             `                          // 运费:包邮为0
	ProductTags              string  `json:"product_tags"   type:"FIND_IN_SET_STR"                ` // 商品标签(DOT)
	StoreIsSelfsupport       bool    `json:"store_is_selfsupport"        `                          // 是否自营(BOOL):1-自营;0-非自营
	ProductSpEnable          bool    `json:"product_sp_enable"           `                          // 允许分销(BOOL):1-启用分销;0-禁用分销
	ProductDistEnable        bool    `json:"product_dist_enable"         `                          // 三级分销允许分销(BOOL):1-启用分销;0-禁用分销
	ProductAddTime           uint64  `json:"product_add_time"            `                          // 添加时间
	ProductSaleTime          uint64  `json:"product_sale_time"           `                          // 上架时间:预设上架时间,可以动态修正状态
	ProductOrder             uint    `json:"product_order"               `                          // 排序:越小越靠前
	ProductSrcId             int64   `json:"product_src_id"              `                          // 产品来源编号
	MarketCategoryId         string  `json:"market_category_id"          `                          // 所属商圈(DOT)
	StoreLatitude            float64 `json:"store_latitude"              `                          // 纬度
	StoreLongitude           float64 `json:"store_longitude"             `                          // 经度
	ProductIsVideo           uint    `json:"product_is_video"            `                          // 是否视频(BOOL):1-有视频;0-无视频
	ProductTransportId       string  `json:"product_transport_id"        `                          // 配送服务(ENUM):1001-快递发货;1002-到店自提;1003-上门服务
	SubsiteId                uint    `json:"subsite_id"                  `                          // 所属分站:0-总站
	ProductIsLock            bool    `json:"product_is_lock"             `                          // 是否锁定(BOOL):0-未锁定; 1-锁定,参加团购的商品不予许修改
	ProductInventoryLock     uint    `json:"product_inventory_lock"      `                          // 库存锁定(ENUM):1001-下单锁定;1002-支付锁定;
	ProductFrom              uint    `json:"product_from"                `                          // 商品来源(ENUM):1000-发布;1001-天猫;1002-淘宝;1003-阿里巴巴;1004-京东;
}

type ListRes model.ProductListOutput

type SearchFilterReq struct {
	g.Meta `path:"/front/pt/product/getSearchFilter" tags:"搜索" method:"get" summary:"商品分类过滤属性"`

	CategoryId uint `json:"category_id"                 ` // 商品分类
}

type SearchFilterRes struct {
	Assists   []*model.ProductAssistOutput `json:"assists" description:"分类辅助属性"`
	Contracts []*entity.ContractType       `json:"contracts" description:"服务"`
	Markets   []interface{}                `json:"markets" description:"商圈"`
	Children  []*entity.ProductCategory    `json:"children" description:"下级分类"`
	Parent    []*entity.ProductCategory    `json:"parent" description:"上级分类"`
	Brands    []*entity.ProductBrand       `json:"brands" description:"品牌"`
	Info      *entity.ProductCategory      `json:"info" description:"信息"`
}

type SearchInfoReq struct {
	g.Meta `path:"/front/shop/mobile/getSearchInfo" tags:"搜索" method:"get" summary:"用户最新搜索记录及系统推荐搜索关键词"`
}

type SearchInfoRes struct {
	SearchHistoryWords []string           `json:"search_history_words"`
	SearchHotWords     []string           `json:"search_hot_words"`
	SuggestSearchWords model.SearchWordVo `json:"suggest_search_words"`
}

// ProductDetailReq 商品详情请求结构体
type ProductDetailReq struct {
	g.Meta     `path:"/front/pt/product/detail" tags:"商品" method:"get" summary:"商品SKU详情"`
	ItemId     int64 `json:"item_id" description:"SKU编号"`          // SKU编号
	DistrictId uint  `json:"district_id" description:"配送地区"`       // 配送地区
	GbId       uint  `json:"gb_id" description:"拼团活动编号"`           // 拼团活动编号
	GifgbagId  uint  `json:"gifgbag_id" description:"A+B组合套餐活动编号"` // A+B组合套餐活动编号
	CutpriceId uint  `json:"cutprice_id" description:"砍价活动编号"`     // 砍价活动编号
}

type ProductDetailRes model.ProductDetailOutput
