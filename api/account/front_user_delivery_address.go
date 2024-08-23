package account

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// 用户地址

type UserDeliveryAddressAdd struct {
	UdName       string `json:"ud_name"        ` // 联系人
	UdIntl       string `json:"ud_intl"        ` // 国家编码
	UdMobile     string `json:"ud_mobile"      ` // 手机号码
	UdTelephone  string `json:"ud_telephone"   ` // 联系电话
	UdProvinceId uint   `json:"ud_province_id" ` // 省编号
	UdProvince   string `json:"ud_province"    ` // 省份
	UdCityId     uint   `json:"ud_city_id"     ` // 市编号
	UdCity       string `json:"ud_city"        ` // 市
	UdCountyId   uint   `json:"ud_county_id"   ` // 县
	UdCounty     string `json:"ud_county"      ` // 县区
	UdAddress    string `json:"ud_address"     ` // 详细地址
	UdPostalcode string `json:"ud_postalcode"  ` // 邮政编码
	UdTagName    string `json:"ud_tag_name"    ` // 地址标签(ENUM):1001-家里;1002-公司
	UdIsDefault  bool   `json:"ud_is_default"  ` // 是否默认(BOOL):0-非默认;1-默认
}

type UserDeliveryAddressAddReq struct {
	g.Meta `path:"/front/account/userDeliveryAddress/add" tags:"user" method:"post" summary:"用户地址接口"`

	UserDeliveryAddressAdd
}

type UserDeliveryAddressAddRes entity.UserDeliveryAddress

type UserDeliveryAddressEditReq struct {
	g.Meta `path:"/front/account/userDeliveryAddress/save" tags:"user" method:"post" summary:"用户地址接口"`

	UdId        interface{} `json:"ud_id"   v:"required#请输入地址编号"    dc:"地址编号"          ` // 地址编号
	UdName      interface{} `json:"ud_name" type:"LIKE"        `                         // 联系人
	UdMobile    interface{} `json:"ud_mobile"      `                                     // 手机号码
	UdIntl      interface{} `json:"ud_intl"        `                                     // 国家编码
	UdAddress   interface{} `json:"ud_address"     `                                     // 详细地址
	UdIsDefault bool        `json:"ud_is_default"  `                                     // 是否默认(BOOL):0-非默认;1-默认
}

type UserDeliveryAddressEditRes struct {
	UdId interface{} `json:"ud_id"   dc:"地址编号"`
}

type UserDeliveryAddressGetReq struct {
	g.Meta `path:"/front/account/userDeliveryAddress/get" tags:"用户地址" method:"get" summary:"用户地址接口"`
	UdId   uint `json:"ud_id"   v:"required#请输入地址编号"    dc:"地址编号"          ` // 地址编号
}

type UserDeliveryAddressGetRes entity.UserDeliveryAddress

type UserDeliveryAddressRemoveReq struct {
	g.Meta `path:"/front/account/userDeliveryAddress/remove" tags:"用户地址" method:"post" summary:"用户地址接口"`
	UdId   uint `json:"ud_id"   v:"required#请输入地址编号"    dc:"地址编号"          ` // 地址编号
}

type UserDeliveryAddressRemoveRes struct {
}

type UserDeliveryAddressListReq struct {
	g.Meta `path:"/front/account/userDeliveryAddress/list" tags:"用户地址" method:"get" summary:"地址列表接口"`
	ml.BaseList

	UdName string `json:"ud_name" type:"LIKE"        ` // 联系人
}

type UserDeliveryAddressListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
