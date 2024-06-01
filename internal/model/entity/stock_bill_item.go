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

// StockBillItem is the golang structure for table stock_bill_item.
type StockBillItem struct {
	StockBillItemId       int64   `json:"stock_bill_item_id"      ` // 编号
	OrderId               string  `json:"order_id"                ` // 源单号码
	OrderItemId           uint64  `json:"order_item_id"           ` // 订单商品表编号
	StockBillId           string  `json:"stock_bill_id"           ` // 订单编号
	ProductId             uint64  `json:"product_id"              ` // 产品编号
	ProductName           string  `json:"product_name"            ` // 商品名称
	ItemId                uint64  `json:"item_id"                 ` // 货品编号
	ItemName              string  `json:"item_name"               ` // 商品名称
	BillItemQuantity      uint    `json:"bill_item_quantity"      ` // 商品数量
	BillItemUnitPrice     float64 `json:"bill_item_unit_price"    ` // 单价
	BillItemSubtotal      float64 `json:"bill_item_subtotal"      ` // 小计
	UnitId                uint    `json:"unit_id"                 ` // 单位编号
	WarehouseItemQuantity uint    `json:"warehouse_item_quantity" ` // 库存量
	StoreId               uint    `json:"store_id"                ` // 店铺编号
	WarehouseId           uint    `json:"warehouse_id"            ` // 所属仓库
	StockTransportTypeId  uint    `json:"stock_transport_type_id" ` // 库存类型(ENUM)
	BillItemRemark        string  `json:"bill_item_remark"        ` // 备注
	BillTypeId            uint    `json:"bill_type_id"            ` // 业务类别(ENUM):2750-入库;2700-出库
}
