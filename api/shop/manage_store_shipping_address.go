package shop

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type StoreShippingAddressAdd struct {
	SsId         uint        `json:"ss_id"          ` // 地址编号
	SsName       string      `json:"ss_name"        ` // 联系人
	SsIntl       string      `json:"ss_intl"        ` // 国家编码
	SsMobile     string      `json:"ss_mobile"      ` // 手机号码
	SsTelephone  string      `json:"ss_telephone"   ` // 联系电话
	SsContacter  string      `json:"ss_contacter"   ` // 联系人(未启用)
	SsPostalcode string      `json:"ss_postalcode"  ` // 邮编
	SsProvinceId uint        `json:"ss_province_id" ` // 省编号
	SsProvince   string      `json:"ss_province"    ` // 省份
	SsCityId     uint        `json:"ss_city_id"     ` // 市编号
	SsCity       string      `json:"ss_city"        ` // 市
	SsCountyId   uint        `json:"ss_county_id"   ` // 县
	SsCounty     string      `json:"ss_county"      ` // 县区
	SsAddress    string      `json:"ss_address"     ` // 详细地址:不必重复填写地区
	SsTime       *gtime.Time `json:"ss_time"        ` // 添加时间
	SsIsDefault  bool        `json:"ss_is_default"  ` // 默认地址(ENUM):0-否;1-是
	StoreId      uint        `json:"store_id"       ` // 所属店铺
}
type StoreShippingAddressEditReq struct {
	g.Meta `path:"/manage/shop/storeShippingAddress/edit" tags:"发货地址" method:"post" summary:"发货地址编辑接口"`

	SsId uint `json:"ss_id"   ` // 发货地址编号`
	StoreShippingAddressAdd
}

type StoreShippingAddressEditRes struct {
	SsId interface{} `json:"ss_id"   dc:"发货地址信息"`
}

type StoreShippingAddressAddReq struct {
	g.Meta `path:"/manage/shop/storeShippingAddress/add" tags:"发货地址" method:"post" summary:"发货地址编辑接口"`

	StoreShippingAddressAdd
}

type StoreShippingAddressRemoveReq struct {
	g.Meta `path:"/manage/shop/storeShippingAddress/remove" tags:"发货地址" method:"post" summary:"发货地址删除接口"`
	SsId   uint `json:"ss_id"   ` // 发货地址编号
}

type StoreShippingAddressRemoveRes struct {
}

type StoreShippingAddressListReq struct {
	g.Meta `path:"/manage/shop/storeShippingAddress/list" tags:"发货地址" method:"get" summary:"发货地址列表接口"`
	ml.BaseList

	SsName string `json:"ss_name"  type:"LIKE"      ` // 联系人
}

type StoreShippingAddressListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
