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

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// ProductIndex is the golang structure of table pt_product_index for DAO operations like Where/Data.
type ProductIndex struct {
	g.Meta                   `orm:"table:pt_product_index, do:true"`
	ProductId                interface{} // 产品编号:定为SPU编号
	ProductNumber            interface{} // SPU商家编码:货号
	ProductName              interface{} // 产品名称
	ProductNameIndex         interface{} // 名称索引关键字(DOT)
	StoreId                  interface{} // 店铺编号
	StoreIsOpen              interface{} // 店铺状态(BOOL):0-关闭;1-运营中
	StoreType                interface{} // 店铺类型(ENUM): 1-卖家店铺; 2-供应商店铺
	StoreCategoryIds         interface{} // 店铺分类(DOT)
	CategoryId               interface{} // 商品分类
	TypeId                   interface{} // 类型编号:冗余检索
	ProductQuantity          interface{} // 商品库存:冗余计算
	ProductWarnQuantity      interface{} // 预警数量
	BrandId                  interface{} // 品牌编号
	ProductServiceTypeIds    interface{} // 售后服务(DOT)
	ProductStateId           interface{} // 商品状态(ENUM):1001-正常;1002-下架仓库中;1003-待审核; 1000-违规禁售
	ProductSaleDistrictIds   interface{} // 销售区域(DOT): district_id=1000全部区域
	ProductVerifyId          interface{} // 商品审核(ENUM):3001-审核通过;3002-审核中;3000-审核未通过
	ProductIsInvoices        interface{} // 是否开票(BOOL): 1-是; 0-否
	ProductIsReturn          interface{} // 允许退换货(BOOL): 1-是; 0-否
	ProductIsRecommend       interface{} // 商品推荐(BOOL):1-是; 0-否
	ProductStockStatus       interface{} // 缺货状态(ENUM):1-有现货;2-预售商品;3-缺货;4-2至3天
	KindId                   interface{} // 商品种类:1201-实物;1202-虚拟
	ActivityTypeIds          interface{} // 参与活动(DOT)
	ContractTypeIds          interface{} // 消费者保障(DOT):由店铺映射到商品
	ProductAssistData        interface{} // 辅助属性值列(DOT):assist_item_id每个都不用 , setFilter(tagid, array(2,3,4));是表示含有标签值2,3,4中的任意一个即符合筛选，这里是or关系。 setFilter(‘tagid’, array(2)); setFilter(‘tagid’, array(3)); 形成and关系| msyql where FIND_IN_SET('1', product_assist_data)
	ProductUnitPriceMin      interface{} // 最低单价
	ProductUnitPriceMax      interface{} // 最高单价
	ProductUnitPointsMin     interface{} // 商品积分
	ProductUnitPointsMax     interface{} // 商品积分
	ProductSaleNum           interface{} // 销售数量
	ProductFavoriteNum       interface{} // 收藏数量
	ProductClick             interface{} // 点击数量
	ProductEvaluationNum     interface{} // 评价次数
	ProductRegionDistrictIds interface{} // 所属区域(DOT)
	ProductFreight           interface{} // 运费:包邮为0
	ProductTags              interface{} // 商品标签(DOT)
	StoreIsSelfsupport       interface{} // 是否自营(BOOL):1-自营;0-非自营
	ProductSpEnable          interface{} // 允许分销(BOOL):1-启用分销;0-禁用分销
	ProductDistEnable        interface{} // 三级分销允许分销(BOOL):1-启用分销;0-禁用分销
	ProductAddTime           interface{} // 添加时间
	ProductSaleTime          interface{} // 上架时间:预设上架时间,可以动态修正状态
	ProductOrder             interface{} // 排序:越小越靠前
	ProductSrcId             interface{} // 产品来源编号
	MarketCategoryId         interface{} // 所属商圈(DOT)
	StoreLatitude            interface{} // 纬度
	StoreLongitude           interface{} // 经度
	ProductIsVideo           interface{} // 是否视频(BOOL):1-有视频;0-无视频
	ProductTransportId       interface{} // 配送服务(ENUM):1001-快递发货;1002-到店自提;1003-上门服务
	SubsiteId                interface{} // 所属分站:0-总站
	ProductIsLock            interface{} // 是否锁定(BOOL):0-未锁定; 1-锁定,参加团购的商品不予许修改
	ProductInventoryLock     interface{} // 库存锁定(ENUM):1001-下单锁定;1002-支付锁定;
	ProductFrom              interface{} // 商品来源(ENUM):1000-发布;1001-天猫;1002-淘宝;1003-阿里巴巴;1004-京东;
	Version                  interface{} // 乐观锁
}

type ProductIndexListInput struct {
	ml.BaseList
	Where ProductIndex // 查询条件
}

type ProductIndexListOutput struct {
	Items   []*entity.ProductIndex // 列表
	Page    int                    // 分页号码
	Total   int                    // 总页数
	Records int                    // 数据总数
	Size    int                    // 单页数量
}

type ProductIndexListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
