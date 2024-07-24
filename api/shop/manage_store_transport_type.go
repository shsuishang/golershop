package shop

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type StoreTransportTypeAdd struct {
	TransportTypeId            uint        `json:"transport_type_id"             ` // 物流及售卖区域编号
	TransportTypeName          string      `json:"transport_type_name"           ` // 模板名称
	StoreId                    uint        `json:"store_id"                      ` // 所属店铺
	ChainId                    uint        `json:"chain_id"                      ` // 門店编号
	TransportTypePricingMethod uint        `json:"transport_type_pricing_method" ` // 计费规则(ENUM):1-按件数;2-按重量;3-按体积
	TransportTypeFreightFree   float64     `json:"transport_type_freight_free"   ` // 免运费额度
	TransportTypeBuildin       uint        `json:"transport_type_buildin"        ` // 系统内置(BOOL):0-非内置;1-内置
	TransportTypeFree          bool        `json:"transport_type_free"           ` // 全免运费(BOOL):0-不全免;1-全免（不限制地区且免运费）
	UpdateTime                 *gtime.Time `json:"update_time"                   ` // 编辑时间
}
type StoreTransportTypeEditReq struct {
	g.Meta `path:"/manage/shop/storeTransportType/edit" tags:"物流工具" method:"post" summary:"物流工具编辑接口"`

	TransportTypeId uint `json:"transport_type_id"   ` // 物流工具编号`
	StoreTransportTypeAdd
}

type StoreTransportTypeEditRes struct {
	TransportTypeId interface{} `json:"transport_type_id"   dc:"物流工具信息"`
}

type StoreTransportTypeAddReq struct {
	g.Meta `path:"/manage/shop/storeTransportType/add" tags:"物流工具" method:"post" summary:"物流工具编辑接口"`

	StoreTransportTypeAdd
}

type StoreTransportTypeRemoveReq struct {
	g.Meta          `path:"/manage/shop/storeTransportType/remove" tags:"物流工具" method:"post" summary:"物流工具删除接口"`
	TransportTypeId uint `json:"transport_type_id"   ` // 物流工具编号
}

type StoreTransportTypeRemoveRes struct {
}

type StoreTransportTypeListReq struct {
	g.Meta `path:"/manage/shop/storeTransportType/list" tags:"物流工具" method:"get" summary:"物流工具列表接口"`
	ml.BaseList

	TransportTypeName string `json:"transport_type_name"   type:"LIKE"        ` // 模板名称
}

type StoreTransportTypeListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
