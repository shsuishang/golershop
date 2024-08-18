package pt

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/entity"
)

// start fo front

// start fo manage
type ProductItemVo struct {
	ItemId uint64 `json:"item_id"              ` // 商品编号-SKU编号
	//ItemName           string  `json:"item_name"            ` // 副标题(DOT):SKU名称
	//ProductId          uint64  `json:"product_id"           ` // 产品编号
	ColorId         int64   `json:"color_id"             ` // 颜色SKU，规格值
	ItemIsDefault   bool    `json:"item_is_default"      ` // 是否为默认展示的商品，必须为item_enable
	ItemNumber      string  `json:"item_number"          ` // SKU商家编码:SKU商家编码为非必填项，若不填写，系统会自动生成一个SKU商家编码。
	ItemBarcode     string  `json:"item_barcode"         ` // 条形码
	ItemCostPrice   float64 `json:"item_cost_price"      ` // 成本价
	ItemUnitPrice   float64 `json:"item_unit_price"      ` // 商品价格
	ItemMarketPrice float64 `json:"item_market_price"    ` // 市场价
	ItemUnitPoints  float64 `json:"item_unit_points"     ` // 积分价格
	ItemQuantity    uint    `json:"item_quantity"        ` // 商品库存
	//ItemQuantityFrozen uint    `json:"item_quantity_frozen" ` // 商品冻结库存
	ItemWarnQuantity uint   `json:"item_warn_quantity"   ` // 库存预警值
	ItemSpec         string `json:"item_spec"            ` // 商品规格序列化(JSON):{spec_id:spec_item_id, spec_id:spec_item_id, spec_id:spec_item_id}
	//SpecItemIds        string  `json:"spec_item_ids"        ` // 商品规格值编号
	ItemEnable uint `json:"item_enable"          ` // 是否启用(LIST):1001-正常;1002-下架仓库中;1000-违规禁售
	//ItemIsChange       bool    `json:"item_is_change"       ` // 被改动(BOOL):0-未改动;1-已改动分销使用
	ItemWeight       float64 `json:"item_weight"          ` // 商品重量:KG
	ItemVolume       float64 `json:"item_volume"          ` // 商品体积:立方米
	ItemFxCommission float64 `json:"item_fx_commission"   ` // 微小店分销佣金
	ItemRebate       float64 `json:"item_rebate"          ` // 返利额度
	//ItemSrcId          int64   `json:"item_src_id"          ` // 供应商SKU编号
	//CategoryId         uint    `json:"category_id"          ` // 商品分类
	//StoreId            uint    `json:"store_id"             ` // 所属店铺
	//Version            uint    `json:"version"              ` // 版本
}

type ProductImageVo struct {
	ProductId uint64 `json:"product_id"         ` // 产品编号:product_id-color_id
	//StoreId          uint   `json:"store_id"           ` // 店铺编号
	ColorId          int64  `json:"color_id"           ` // 系统默认颜色规格Id/spec_item_id, 如果没有则一条记录为0
	ColorName        string `json:"color_name"         ` // 规格值
	ItemImageDefault string `json:"item_image_default" ` // 商品主图
	ItemImageOther   string `json:"item_image_other"   ` // 副图(DOT)
}

type ProductAdd struct {
	//base
	ProductNumber string `json:"product_number"         ` // SPU商家编码:货号
	ProductName   string `json:"product_name"           ` // 产品名称
	ProductTips   string `json:"product_tips"           ` // 商品卖点:商品广告词
	//StoreId         uint   `json:"store_id"               ` // 店铺编号
	ProductImage    string `json:"product_image"          ` // 商品主图
	ProductVideo    string `json:"product_video"          ` // 产品视频
	TransportTypeId uint   `json:"transport_type_id"      ` // 选择售卖区域:完成售卖区域及运费设置
	ProductBuyLimit uint   `json:"product_buy_limit"      ` // 每人限购
	//ProductSrcId    int64  `json:"product_src_id"         ` // 产品来源编号

	//index
	ProductId uint64 `json:"product_id"                  ` // 产品编号:定为SPU编号
	//ProductNumber            string  `json:"product_number"              ` // SPU商家编码:货号
	//ProductName              string  `json:"product_name"                ` // 产品名称:店铺平台先在对用表中检索后通过id检索,检索使用
	//ProductNameIndex string `json:"product_name_index"          ` // 名称索引关键字(DOT)
	//StoreId                  uint    `json:"store_id"                    ` // 店铺编号
	//StoreIsOpen              bool    `json:"store_is_open"               `     // 店铺状态(BOOL):0-关闭;1-运营中
	//StoreType                uint    `json:"store_type"                  `     // 店铺类型(ENUM): 1-卖家店铺; 2-供应商店铺
	//StoreCategoryIds         string  `json:"store_category_ids"          `     // 店铺分类编号(DOT)
	CategoryId uint `json:"category_id"                 ` // 商品分类
	//TypeId                   uint    `json:"type_id"                     `     // 类型编号:冗余检索
	//ProductQuantity          uint    `json:"product_quantity"            `     // 商品库存:冗余计算
	//ProductWarnQuantity      uint    `json:"product_warn_quantity"       `     // 预警数量
	BrandId                uint   `json:"brand_id"                    ` // 品牌编号
	ProductServiceTypeIds  string `json:"product_service_type_ids"    ` // 售后服务(DOT)
	ProductStateId         uint   `json:"product_state_id"            ` // 商品状态:1001-正常;1002-下架仓库中;1003-待审核; 1000-违规禁售
	ProductSaleDistrictIds string `json:"product_sale_district_ids"   ` // 销售区域(DOT): district_id=1000全部区域
	//ProductVerifyId          uint    `json:"product_verify_id"           `     // 商品审核(ENUM):3001-审核通过;3002-审核中;3000-审核未通过
	//ProductIsInvoices        bool    `json:"product_is_invoices"         `     // 是否开票(BOOL): 1-是; 0-否
	//ProductIsReturn          bool    `json:"product_is_return"           `     // 是否允许退换货(BOOL): 1-是; 0-否
	//ProductIsRecommend       bool    `json:"product_is_recommend"        `     // 商品推荐(BOOL):1-是; 0-否
	//ProductStockStatus       uint    `json:"product_stock_status"        `     // 缺货状态(ENUM):1-有现货;2-预售商品;3-缺货;4-2至3天
	KindId uint `json:"kind_id"                     ` // 商品种类:1201-实物;1202-虚拟
	//ActivityTypeIds          string  `json:"activity_type_ids"           `     // 参与活动(DOT)
	ContractTypeIds string `json:"contract_type_ids"           ` // 消费者保障(DOT):由店铺映射到商品
	//ProductAssistData        string  `json:"product_assist_data"         `     // 辅助属性值列(DOT):assist_item_id每个都不用 , setFilter(tagid, array(2,3,4));是表示含有标签值2,3,4中的任意一个即符合筛选，这里是or关系。 setFilter(‘tagid’, array(2)); setFilter(‘tagid’, array(3)); 形成and关系| msyql where FIND_IN_SET('1', product_assist_data)
	//ProductUnitPriceMin      float64 `json:"product_unit_price_min"          ` // 商品单价
	//ProductUnitPriceMax      float64 `json:"product_unit_price_max"      `     // 商品最高单价
	//ProductUnitPointsMin     float64 `json:"product_unit_points_min"         ` // 商品积分
	//ProductUnitPointsMax     float64 `json:"product_unit_points_max"     `     // 商品积分
	//ProductSaleNum           uint    `json:"product_sale_num"            `     // 销售量
	//ProductFavoriteNum       uint    `json:"product_favorite_num"        `     // 收藏数量人气
	//ProductClick             uint    `json:"product_click"               `     // 商品点击数量
	//ProductEvaluationNum     uint    `json:"product_evaluation_num"      `     // 评价次数
	ProductRegionDistrictIds string `json:"product_region_district_ids" ` // 所属区域(DOT)
	//ProductFreight           float64 `json:"product_freight"             `     // 运费:包邮为0，检索使用
	ProductTags string `json:"product_tags"                ` // 商品标签(DOT)
	//StoreIsSelfsupport       uint    `json:"store_is_selfsupport"        `     // 是否自营(BOOL):1-自营;0-非自营
	//ProductSpEnable          uint    `json:"product_sp_enable"           ` // 允许分销(BOOL):1-启用分销;0-禁用分销
	//ProductDistEnable        uint    `json:"product_dist_enable"         ` // 三级分销允许分销(BOOL):1-启用分销;0-禁用分销
	//ProductAddTime           uint64  `json:"product_add_time"            ` // 添加时间
	ProductSaleTime uint64 `json:"product_sale_time"           ` // 上架时间:预设上架时间,可以动态修正状态
	//ProductOrder    uint   `json:"product_order"               ` // 排序:越小越靠前
	//ProductSrcId             int64   `json:"product_src_id"              ` // 产品来源编号
	MarketCategoryId string `json:"market_category_id"          ` // 所属商圈(DOT)
	//StoreLatitude        float64 `json:"store_latitude"              ` // 纬度
	//StoreLongitude       float64 `json:"store_longitude"             ` // 经度
	//ProductIsVideo       uint    `json:"product_is_video"            ` // 是否视频(BOOL):1-有视频;0-无视频
	ProductTransportId string `json:"product_transport_id"        ` // 配送服务(ENUM):1001-快递发货;1002-到店自提;1003-上门服务
	//SubsiteId            uint    `json:"subsite_id"                  ` // 所属分站:0-总站
	//ProductIsLock        bool    `json:"product_is_lock"        `      // 是否锁定(BOOL):0-未锁定; 1-锁定,参加团购的商品不予许修改
	ProductInventoryLock uint `json:"product_inventory_lock" ` // 库存锁定(ENUM):1001-下单锁定;1002-支付锁定;
	//ProductFrom          uint    `json:"product_from"           `      // 商品来源(ENUM):1000-发布;1001-天猫;1002-淘宝;1003-阿里巴巴;1004-京东;
	//Version              uint    `json:"version"                     ` // 乐观锁

	//info
	ProductAssist string `json:"product_assist"           ` // 属性(JSON) - 辅助属性及VAL
	ProductSpec   string `json:"product_spec"             ` // 规格(JSON)-规格、规格值、goods_id  规格不需要全选就可以添加对应数据[{'id' : spec_id, 'name' : spec_name, 'item':[{'id' : spec_item_id, 'name' : spec_item_name}, {'id' : spec_item_id, 'name' : spec_item_name}]},{'id' : spec_id, 'name' : spec_name, 'item':[{'id' : spec_item_id, 'name' : spec_item_name}, {'id' : spec_item_id, 'name' : spec_item_name}]}]
	//ProductUniqid          string `json:"product_uniqid"           ` // 商品SKU(JSON):{'uniq_id':[item_id, price, url]}
	ProductDetail string `json:"product_detail"           ` // 商品描述
	//ProductMetaTitle       string `json:"product_meta_title"       ` // Meta Tag 标题
	//ProductMetaDescription string `json:"product_meta_description" ` // Meta Tag 描述
	//ProductMetaKeyword     string `json:"product_meta_keyword"     ` // Meta Tag 关键字

	//虚拟
	ProductValidPeriod          uint  `json:"product_valid_period"           ` // 有效期:1001-长期有效;1002-自定义有效期;1003-购买起有效时长年单位
	ProductValidityStart        int64 `json:"product_validity_start"         ` // 开始时间
	ProductValidityEnd          int64 `json:"product_validity_end"           ` // 失效时间
	ProductValidType            uint  `json:"product_valid_type"             ` // 服务类型(ENUM):1001-到店服务;1002-上门服务
	ProductServiceDateFlag      bool  `json:"product_service_date_flag"      ` // 填写预约日期(BOOL):0-否;1-是
	ProductServiceContactorFlag bool  `json:"product_service_contactor_flag" ` // 填写联系人(BOOL):0-否;1-是
	ProductValidRefundFlag      bool  `json:"product_valid_refund_flag"      ` // 支持过期退款(BOOL):0-否;1-是

	//product_items
	ProductItems []*ProductItemVo `json:"product_items"     ` // 商品SKU信息 JSON字符串

	//product_images
	ProductImages []*ProductImageVo `json:"product_images"     ` // 图片信息 JSON字符串

}
type ProductSaveReq struct {
	g.Meta `path:"/manage/pt/productBase/save" tags:"商品管理" method:"post" summary:"商品编辑接口"`

	//ProductId uint `json:"product_id"   v:"required#请输入商品编号"    dc:"商品编号"     `
	ProductAdd
}

type ProductSaveRes struct {
	ProductId interface{} `json:"product_id"   dc:"商品信息"`
}

type ProductRemoveReq struct {
	g.Meta    `path:"/manage/pt/productBase/remove" tags:"商品管理" method:"post" summary:"商品删除接口"`
	ProductId uint `json:"product_id" v:"required#请输入商品编号"   dc:"商品信息"`
}

type ProductRemoveRes struct {
}

type ProductListReq struct {
	g.Meta `path:"/manage/pt/productBase/list" tags:"商品管理" method:"get" summary:"商品列表接口"`
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
	ProductAssistData        string  `json:"product_assist_data"         `                          // 辅助属性值列(DOT):assist_item_id每个都不用 , setFilter(tagid, array(2,3,4));是表示含有标签值2,3,4中的任意一个即符合筛选，这里是or关系。 setFilter(‘tagid’, array(2)); setFilter(‘tagid’, array(3)); 形成and关系| msyql where FIND_IN_SET('1', product_assist_data)
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

type ProductListRes struct {
	Items   interface{} `json:"items"    dc:"商品列表信息"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ProductDateReq struct {
	g.Meta `path:"/manage/pt/productBase/getProductDate" tags:"商品管理" method:"get" summary:"商品信息接口"`

	ProductId uint64 `json:"product_id"             ` // 产品编号

}

type ProductDateRes struct {
	model.ProductDateOutput
}

type ProductEditStateRes struct {
	ProductId interface{} `json:"product_id"             ` // 产品编号
}

type ProductEditStateReq struct {
	g.Meta `path:"/manage/pt/productBase/editState" tags:"商品管理" method:"post" summary:"商品管理状态编辑接口"`

	ProductId      interface{} `json:"product_id"   v:"required#请输入商品编号"           ` // 产品编号`
	ProductStateId uint        `json:"product_state_id"            `                 // 商品状态:1001-正常;1002-下架仓库中;1003-待审核; 1000-违规禁售
}

type ProductBaseItemListReq struct {
	g.Meta `path:"/manage/pt/productBase/listItem" tags:"商品SKU列表" method:"get" summary:"商品SKU列表接口"`
	ml.BaseList

	BrandId     uint   `json:"brand_id"                    ` // 品牌编号
	ProductId   uint64 `json:"product_id"                  ` // 产品编号
	ProductName string `json:"product_name" type:"LIKE"    ` // 产品名称
	CategoryId  uint   `json:"category_id"                 ` // 商品分类
	ItemId      string `json:"item_id"                  `    // 商品编号-SKU编号
}

type ProductBaseItemListRes struct {
	Assists      []model.ProductAssistOutput `json:"assists"       dc:"分类辅助属性"` // 分类辅助属性
	ActivityBase entity.ActivityBase         `json:"activity_base" dc:"活动信息"`   // 活动信息
	Items        interface{}                 `json:"items"    dc:"商品列表信息"`
	Page         int                         `json:"page"`    // 分页号码
	Total        int                         `json:"total"`   // 总页数
	Records      int                         `json:"records"` // 数据总数
	Size         int                         `json:"size"`    // 单页数量
}

type BatchEditStateReq struct {
	g.Meta `path:"/manage/pt/productBase/batchEditState" tags:"批量修改商品状态" method:"post" summary:"批量修改商品状态"`

	ProductIds     string `json:"product_ids"                  ` // 产品编号
	ProductStateId uint   `json:"product_state_id"            `  // 商品状态:1001-正常;1002-下架仓库中;1003-待审核; 1000-违规禁售
}
type BatchEditStateRes struct{}

type GetStockWarningItemsReq struct {
	g.Meta `path:"//manage/pt/productItem/getStockWarningItems" tags:"库存警告商品" method:"get" summary:"库存警告商品item-分页列表查询"`
	ml.BaseList

	ProductName string `json:"product_name" `  // 产品名称
	ProductId   uint64 `json:"product_id"    ` // 产品编号
	ItemId      string `json:"item_id"     `   // 商品编号-SKU编号
}
type GetStockWarningItemsRes struct {
	Items   interface{} `json:"items"    dc:"商品列表信息"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
