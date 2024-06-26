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

// ProductItem is the golang structure of table pt_product_item for DAO operations like Where/Data.
type ProductItem struct {
	g.Meta             `orm:"table:pt_product_item, do:true"`
	ItemId             interface{} // 商品编号-SKU编号
	ItemName           interface{} // 副标题(DOT):SKU名称
	ItemIndex          interface{} // 索引(DOT)
	ProductId          interface{} // 产品编号
	ColorId            interface{} // 颜色SKU，规格值
	ItemIsDefault      interface{} // 是否为默认展示的商品，必须为item_enable
	ItemNumber         interface{} // SKU商家编码:SKU商家编码为非必填项，若不填写，系统会自动生成一个SKU商家编码。
	ItemBarcode        interface{} // 条形码
	ItemCostPrice      interface{} // 成本价
	ItemUnitPrice      interface{} // 商品价格
	ItemMarketPrice    interface{} // 市场价
	ItemUnitPoints     interface{} // 积分价格
	ItemQuantity       interface{} // 商品库存
	ItemQuantityFrozen interface{} // 商品冻结库存
	ItemWarnQuantity   interface{} // 库存预警值
	ItemSpec           interface{} // 商品规格序列化(JSON):{spec_id:spec_item_id, spec_id:spec_item_id, spec_id:spec_item_id}
	SpecItemIds        interface{} // 商品规格值编号
	ItemEnable         interface{} // 是否启用(LIST):1001-正常;1002-下架仓库中;1000-违规禁售
	ItemIsChange       interface{} // 被改动(BOOL):0-未改动;1-已改动分销使用
	ItemWeight         interface{} // 商品重量:KG
	ItemVolume         interface{} // 商品体积:立方米
	ItemFxCommission   interface{} // 微小店分销佣金
	ItemRebate         interface{} // 返利额度
	ItemSrcId          interface{} // 供应商SKU编号
	CategoryId         interface{} // 商品分类
	CourseCategoryId   interface{} // 课程分类
	StoreId            interface{} // 所属店铺
	Version            interface{} // 版本
}

type ProductItemListInput struct {
	ml.BaseList
	Where ProductItem // 查询条件
	ProductStateId interface{}
}

type ProductItemListOutput struct {
	Items   []*entity.ProductItem // 列表
	Page    int                   // 分页号码
	Total   int                   // 总页数
	Records int                   // 数据总数
	Size    int                   // 单页数量
}

type ProductItemListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
