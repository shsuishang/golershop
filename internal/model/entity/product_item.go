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

// ProductItem is the golang structure for table product_item.
type ProductItem struct {
	ItemId             uint64  `json:"item_id"              ` // 商品编号-SKU编号
	ItemName           string  `json:"item_name"            ` // 副标题(DOT):SKU名称
	ItemIndex          string  `json:"item_index"           ` // 索引(DOT)
	ProductId          uint64  `json:"product_id"           ` // 产品编号
	ColorId            int64   `json:"color_id"             ` // 颜色SKU，规格值
	ItemIsDefault      bool    `json:"item_is_default"      ` // 是否为默认展示的商品，必须为item_enable
	ItemNumber         string  `json:"item_number"          ` // SKU商家编码:SKU商家编码为非必填项，若不填写，系统会自动生成一个SKU商家编码。
	ItemBarcode        string  `json:"item_barcode"         ` // 条形码
	ItemCostPrice      float64 `json:"item_cost_price"      ` // 成本价
	ItemUnitPrice      float64 `json:"item_unit_price"      ` // 商品价格
	ItemMarketPrice    float64 `json:"item_market_price"    ` // 市场价
	ItemUnitPoints     float64 `json:"item_unit_points"     ` // 积分价格
	ItemQuantity       uint    `json:"item_quantity"        ` // 商品库存
	ItemQuantityFrozen uint    `json:"item_quantity_frozen" ` // 商品冻结库存
	ItemWarnQuantity   uint    `json:"item_warn_quantity"   ` // 库存预警值
	ItemSpec           string  `json:"item_spec"            ` // 商品规格序列化(JSON):{spec_id:spec_item_id, spec_id:spec_item_id, spec_id:spec_item_id}
	SpecItemIds        string  `json:"spec_item_ids"        ` // 商品规格值编号
	ItemEnable         uint    `json:"item_enable"          ` // 是否启用(LIST):1001-正常;1002-下架仓库中;1000-违规禁售
	ItemIsChange       bool    `json:"item_is_change"       ` // 被改动(BOOL):0-未改动;1-已改动分销使用
	ItemWeight         float64 `json:"item_weight"          ` // 商品重量:KG
	ItemVolume         float64 `json:"item_volume"          ` // 商品体积:立方米
	ItemFxCommission   float64 `json:"item_fx_commission"   ` // 微小店分销佣金
	ItemRebate         float64 `json:"item_rebate"          ` // 返利额度
	ItemSrcId          int64   `json:"item_src_id"          ` // 供应商SKU编号
	CategoryId         uint    `json:"category_id"          ` // 商品分类
	CourseCategoryId   uint    `json:"course_category_id"   ` // 课程分类
	StoreId            uint    `json:"store_id"             ` // 所属店铺
	Version            uint    `json:"version"              ` // 版本

	// 表中不存在的字段
	AvailableQuantity uint        `json:"available_quantity" gorm:"-"` // 可用库存
	ProductItemName   string      `json:"product_item_name" gorm:"-"`  // Spec名称
	ActivityId        uint        `json:"activity_id" gorm:"-"`        //当前使用活动编号
	ItemSalePrice     float64     `json:"item_sale_price" gorm:"-"`    // 商品销售价
	ItemSavePrice     float64     `json:"item_save_price" gorm:"-"`    // 节省单价
	ActivityInfo      interface{} `json:"activity_info" gorm:"-"`      // 活动信息
}

// GetAvailableQuantity 获取有效
func (m *ProductItem) GetAvailableQuantity() uint {
	return m.ItemQuantity - m.ItemQuantityFrozen
}
