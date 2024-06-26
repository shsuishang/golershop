// +----------------------------------------------------------------------
// | ShopSuite商城系统 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 随商信息技术（上海）有限公司
// +----------------------------------------------------------------------
// | 未获商业授权前，不得将本软件用于商业用途。禁止整体或任何部分基础上以发展任何派生版本、
// | 修改版本或第三方版本用于重新分发。
// +----------------------------------------------------------------------
// | 官方网站: https://www.shopsuite.cn  https://www.golershop.cn
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本公司对该软件产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细见https://www.golershop.cn/policy
// +----------------------------------------------------------------------

package entity

// ProductIndex is the golang structure for table product_index.
type ProductIndex struct {
	ProductId                uint64  `json:"product_id"                  ` // 产品编号:定为SPU编号
	ProductNumber            string  `json:"product_number"              ` // SPU商家编码:货号
	ProductName              string  `json:"product_name"                ` // 产品名称
	ProductNameIndex         string  `json:"product_name_index"          ` // 名称索引关键字(DOT)
	StoreId                  uint    `json:"store_id"                    ` // 店铺编号
	StoreIsOpen              bool    `json:"store_is_open"               ` // 店铺状态(BOOL):0-关闭;1-运营中
	StoreType                uint    `json:"store_type"                  ` // 店铺类型(ENUM): 1-卖家店铺; 2-供应商店铺
	StoreCategoryIds         string  `json:"store_category_ids"          ` // 店铺分类(DOT)
	CategoryId               uint    `json:"category_id"                 ` // 商品分类
	CourseCategoryId         uint    `json:"course_category_id"          ` // 课程分类
	TypeId                   uint    `json:"type_id"                     ` // 类型编号:冗余检索
	ProductQuantity          uint    `json:"product_quantity"            ` // 商品库存:冗余计算
	ProductWarnQuantity      uint    `json:"product_warn_quantity"       ` // 预警数量
	BrandId                  uint    `json:"brand_id"                    ` // 品牌编号
	ProductServiceTypeIds    string  `json:"product_service_type_ids"    ` // 售后服务(DOT)
	ProductStateId           uint    `json:"product_state_id"            ` // 商品状态(ENUM):1001-正常;1002-下架仓库中;1003-待审核; 1000-违规禁售
	ProductSaleDistrictIds   string  `json:"product_sale_district_ids"   ` // 销售区域(DOT): district_id=1000全部区域
	ProductVerifyId          uint    `json:"product_verify_id"           ` // 商品审核(ENUM):3001-审核通过;3002-审核中;3000-审核未通过
	ProductIsInvoices        bool    `json:"product_is_invoices"         ` // 是否开票(BOOL): 1-是; 0-否
	ProductIsReturn          bool    `json:"product_is_return"           ` // 允许退换货(BOOL): 1-是; 0-否
	ProductIsRecommend       bool    `json:"product_is_recommend"        ` // 商品推荐(BOOL):1-是; 0-否
	ProductStockStatus       uint    `json:"product_stock_status"        ` // 缺货状态(ENUM):1-有现货;2-预售商品;3-缺货;4-2至3天
	KindId                   uint    `json:"kind_id"                     ` // 商品种类:1201-实物;1202-虚拟
	ActivityTypeIds          string  `json:"activity_type_ids"           ` // 参与活动(DOT)
	ContractTypeIds          string  `json:"contract_type_ids"           ` // 消费者保障(DOT):由店铺映射到商品
	ProductAssistData        string  `json:"product_assist_data"         ` // 辅助属性值列(DOT):assist_item_id每个都不用 , setFilter(tagid, array(2,3,4));是表示含有标签值2,3,4中的任意一个即符合筛选，这里是or关系。 setFilter(‘tagid’, array(2)); setFilter(‘tagid’, array(3)); 形成and关系| msyql where FIND_IN_SET('1', product_assist_data)
	ProductUnitPriceMin      float64 `json:"product_unit_price_min"      ` // 最低单价
	ProductUnitPriceMax      float64 `json:"product_unit_price_max"      ` // 最高单价
	ProductUnitPointsMin     float64 `json:"product_unit_points_min"     ` // 商品积分
	ProductUnitPointsMax     float64 `json:"product_unit_points_max"     ` // 商品积分
	ProductSaleNum           uint    `json:"product_sale_num"            ` // 销售数量
	ProductFavoriteNum       uint    `json:"product_favorite_num"        ` // 收藏数量
	ProductClick             uint    `json:"product_click"               ` // 点击数量
	ProductEvaluationNum     uint    `json:"product_evaluation_num"      ` // 评价次数
	ProductRegionDistrictIds string  `json:"product_region_district_ids" ` // 所属区域(DOT)
	ProductFreight           float64 `json:"product_freight"             ` // 运费:包邮为0
	ProductTags              string  `json:"product_tags"                ` // 商品标签(DOT)
	StoreIsSelfsupport       bool    `json:"store_is_selfsupport"        ` // 是否自营(BOOL):1-自营;0-非自营
	ProductSpEnable          bool    `json:"product_sp_enable"           ` // 允许分销(BOOL):1-启用分销;0-禁用分销
	ProductDistEnable        bool    `json:"product_dist_enable"         ` // 三级分销允许分销(BOOL):1-启用分销;0-禁用分销
	ProductAddTime           uint64  `json:"product_add_time"            ` // 添加时间
	ProductSaleTime          uint64  `json:"product_sale_time"           ` // 上架时间:预设上架时间,可以动态修正状态
	ProductOrder             uint    `json:"product_order"               ` // 排序:越小越靠前
	ProductSrcId             int64   `json:"product_src_id"              ` // 产品来源编号
	MarketCategoryId         string  `json:"market_category_id"          ` // 所属商圈(DOT)
	StoreLatitude            float64 `json:"store_latitude"              ` // 纬度
	StoreLongitude           float64 `json:"store_longitude"             ` // 经度
	ProductIsVideo           uint    `json:"product_is_video"            ` // 是否视频(BOOL):1-有视频;0-无视频
	ProductTransportId       string  `json:"product_transport_id"        ` // 配送服务(ENUM):1001-快递发货;1002-到店自提;1003-上门服务
	SubsiteId                uint    `json:"subsite_id"                  ` // 所属分站:0-总站
	ProductIsLock            bool    `json:"product_is_lock"             ` // 是否锁定(BOOL):0-未锁定; 1-锁定,参加团购的商品不予许修改
	ProductInventoryLock     uint    `json:"product_inventory_lock"      ` // 库存锁定(ENUM):1001-下单锁定;1002-支付锁定;
	ProductFrom              uint    `json:"product_from"                ` // 商品来源(ENUM):1000-发布;1001-天猫;1002-淘宝;1003-阿里巴巴;1004-京东;
	Version                  uint    `json:"version"                     ` // 乐观锁
}
