package shop

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type StoreTransportItemAdd struct {
	TransportItemId           uint    `json:"transport_item_id"            ` // 编号
	TransportTypeId           uint    `json:"transport_type_id"            ` // 自定义物流模板编号
	TransportItemDefaultNum   uint    `json:"transport_item_default_num"   ` // 默认数量
	TransportItemDefaultPrice float64 `json:"transport_item_default_price" ` // 默认运费
	TransportItemAddNum       uint    `json:"transport_item_add_num"       ` // 增加数量
	TransportItemAddPrice     float64 `json:"transport_item_add_price"     ` // 增加运费
	TransportItemCityIds      string  `json:"transport_item_city_ids"      ` // 区域城市id(DOT):分区域全国都可售卖使用 * 代替id 价格需要根据重量等等计算物流费用
	TransportCityIds          string  `json:"transport_city_ids"           ` // 区域城市id(JSON)
	TransportItemCityName     string  `json:"transport_item_city_name"     ` // 区域城市名称(JSON)
}
type StoreTransportItemEditReq struct {
	g.Meta `path:"/manage/shop/storeTransportItem/edit" tags:"物流工具" method:"post" summary:"物流工具编辑接口"`

	TransportItemId uint `json:"transport_item_id"   ` // 物流工具编号`
	StoreTransportItemAdd
}

type StoreTransportItemEditRes struct {
	TransportItemId interface{} `json:"transport_item_id"   dc:"物流工具信息"`
}

type StoreTransportItemAddReq struct {
	g.Meta `path:"/manage/shop/storeTransportItem/add" tags:"物流工具" method:"post" summary:"物流工具编辑接口"`

	StoreTransportItemAdd
}

type StoreTransportItemRemoveReq struct {
	g.Meta          `path:"/manage/shop/storeTransportItem/remove" tags:"物流工具" method:"post" summary:"物流工具删除接口"`
	TransportItemId uint `json:"transport_item_id"   ` // 物流工具编号
}

type StoreTransportItemRemoveRes struct {
}

type StoreTransportItemListReq struct {
	g.Meta `path:"/manage/shop/storeTransportItem/list" tags:"物流工具" method:"get" summary:"物流工具列表接口"`
	ml.BaseList

	TransportItemId uint `json:"transport_item_id"            ` // 编号
	TransportTypeId uint `json:"transport_type_id"            ` // 自定义物流模板编号
}

type StoreTransportItemListRes struct {
	Items interface{} `json:"items"    dc:"分页数据内容"`

	Page    int `json:"page"`    // 分页号码
	Total   int `json:"total"`   // 总页数
	Records int `json:"records"` // 数据总数
	Size    int `json:"size"`    // 单页数量
}
