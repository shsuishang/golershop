package pt

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo manage
type ProductItemAdd struct {
	ItemId             uint64  `json:"item_id"              ` // 商品编号-SKU编号
	ItemName           string  `json:"item_name"            ` // 副标题(DOT):SKU名称
	ProductId          uint64  `json:"product_id"           ` // 产品编号
	ColorId            int64   `json:"color_id"             ` // 颜色SKU，规格值
	ItemIsDefault      bool    `json:"item_is_default"      ` // 是否为默认展示的商品，必须为item_enable
	ItemNumber         string  `json:"item_number"          ` // SKU商家编码:SKU商家编码为非必填项，若不填写，系统会自动生成一个SKU商家编码。
	ItemBarcode        string  `json:"item_barcode"         ` // 条形码
	ItemCostPrice      float64 `json:"item_cost_price"      ` // 成本价
	ItemUnitPrice      float64 `json:"item_unit_price"      ` // 商品价格
	ItemMarketPrice    float64 `json:"item_market_price"    ` // 市场价
	ItemUnitPoints     float64 `json:"item_unit_points"     ` // 积分价格
	ItemQuantity       int     `json:"item_quantity"        ` // 商品库存
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
	StoreId            uint    `json:"store_id"             ` // 所属店铺
	Version            uint    `json:"version"              ` // 版本
}

type ProductItemEditReq struct {
	g.Meta `path:"/manage/pt/productItem/edit" tags:"商品SKU" method:"post" summary:"商品SKU编辑接口"`

	ItemId uint64 `json:"item_id"              ` // 商品编号-SKU编号
	ProductItemAdd
}

type ProductItemEditRes struct {
	ItemId interface{} `json:"item_id"   dc:"商品SKU信息"`
}

type ProductItemEditStateReq struct {
	g.Meta `path:"/manage/pt/productItem/editState" tags:"商品SKU" method:"post" summary:"商品SKU状态编辑接口"`

	ItemId     uint64 `json:"item_id"              ` // 商品编号-SKU编号
	ItemEnable uint   `json:"item_enable"          ` // 是否启用(LIST):1001-正常;1002-下架仓库中;1000-违规禁售
}

type ProductItemEditStateRes struct {
	ItemId interface{} `json:"item_id"   dc:"商品SKU"`
}

type ProductItemListReq struct {
	g.Meta `path:"/manage/pt/productItem/list" tags:"商品SKU" method:"get" summary:"商品SKU列表接口"`
	ml.BaseList

	ItemName    string `json:"item_name"            ` // 副标题(DOT):SKU名称
	ProductId   uint64 `json:"product_id"           ` // 产品编号
	ItemNumber  string `json:"item_number"          ` // SKU商家编码:SKU商家编码为非必填项，若不填写，系统会自动生成一个SKU商家编码。
	ItemBarcode string `json:"item_barcode"         ` // 条形码
	ItemEnable  uint   `json:"item_enable"          ` // 是否启用(LIST):1001-正常;1002-下架仓库中;1000-违规禁售
	CategoryId  uint   `json:"category_id"          ` // 商品分类
}

type ProductItemListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
