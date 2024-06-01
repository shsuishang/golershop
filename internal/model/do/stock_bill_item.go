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

// StockBillItem is the golang structure of table invoicing_stock_bill_item for DAO operations like Where/Data.
type StockBillItem struct {
	g.Meta                `orm:"table:invoicing_stock_bill_item, do:true"`
	StockBillItemId       interface{} // 编号
	OrderId               interface{} // 源单号码
	OrderItemId           interface{} // 订单商品表编号
	StockBillId           interface{} // 订单编号
	ProductId             interface{} // 产品编号
	ProductName           interface{} // 商品名称
	ItemId                interface{} // 货品编号
	ItemName              interface{} // 商品名称
	BillItemQuantity      interface{} // 商品数量
	BillItemUnitPrice     interface{} // 单价
	BillItemSubtotal      interface{} // 小计
	UnitId                interface{} // 单位编号
	WarehouseItemQuantity interface{} // 库存量
	StoreId               interface{} // 店铺编号
	WarehouseId           interface{} // 所属仓库
	StockTransportTypeId  interface{} // 库存类型(ENUM)
	BillItemRemark        interface{} // 备注
	BillTypeId            interface{} // 业务类别(ENUM):2750-入库;2700-出库
}

type StockBillItemListInput struct {
	ml.BaseList
	Where StockBillItem // 查询条件
}

type StockBillItemListOutput struct {
	Items   []*entity.StockBillItem // 列表
	Page    int                     // 分页号码
	Total   int                     // 总页数
	Records int                     // 数据总数
	Size    int                     // 单页数量
}

type StockBillItemListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
