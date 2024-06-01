package shop

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type StoreExpressLogisticsAdd struct {
	LogisticsId        uint        `json:"logistics_id"         ` // 物流编号
	StoreId            uint        `json:"store_id"             ` // 店铺
	ChainId            uint        `json:"chain_id"             ` // store_express_logistics_cond
	LogisticsName      string      `json:"logistics_name"       ` // 物流名称
	LogisticsPinyin    string      `json:"logistics_pinyin"     ` // 物流
	LogisticsNumber    uint        `json:"logistics_number"     ` // 物流公司编号
	LogisticsState     uint        `json:"logistics_state"      ` // 电子面单状态
	ExpressId          uint        `json:"express_id"           ` // 对应快递公司
	ExpressName        string      `json:"express_name"         ` // 快递公司
	LogisticsIsDefault bool        `json:"logistics_is_default" ` // 是否为默认(BOOL):1-默认;0-非默认
	LogisticsMobile    string      `json:"logistics_mobile"     ` // 联系手机
	LogisticsContacter string      `json:"logistics_contacter"  ` // 联系人
	LogisticsAddress   string      `json:"logistics_address"    ` // 联系地址
	LogisticsFee       string      `json:"logistics_fee"        ` // 物流运费
	LogisticsIsEnable  bool        `json:"logistics_is_enable"  ` // 是否启用(BOOL):1-启用;0-禁用
	CreateTime         *gtime.Time `json:"create_time"          ` // 创建时间
}
type StoreExpressLogisticsEditReq struct {
	g.Meta `path:"/manage/shop/storeExpressLogistics/edit" tags:"物流管理" method:"post" summary:"物流管理编辑接口"`

	LogisticsId uint `json:"logistics_id"   ` // 物流管理编号`
	StoreExpressLogisticsAdd
}

type StoreExpressLogisticsEditRes struct {
	LogisticsId interface{} `json:"logistics_id"   dc:"物流管理信息"`
}

type StoreExpressLogisticsAddReq struct {
	g.Meta `path:"/manage/shop/storeExpressLogistics/add" tags:"物流管理" method:"post" summary:"物流管理编辑接口"`

	StoreExpressLogisticsAdd
}

type StoreExpressLogisticsRemoveReq struct {
	g.Meta      `path:"/manage/shop/storeExpressLogistics/remove" tags:"物流管理" method:"post" summary:"物流管理删除接口"`
	LogisticsId uint `json:"logistics_id"   ` // 物流管理编号

}

type StoreExpressLogisticsRemoveRes struct {
}

type StoreExpressLogisticsListReq struct {
	g.Meta `path:"/manage/shop/storeExpressLogistics/list" tags:"物流管理" method:"get" summary:"物流管理列表接口"`
	ml.BaseList

	LogisticsName string `json:"logistics_name"  type:"LIKE"     ` // 物流名称
}

type StoreExpressLogisticsListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
