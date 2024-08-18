package account

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// 用户地址

type DeliveryAddressAdd struct {
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

type DeliveryAddressEditReq struct {
	g.Meta `path:"/manage/account/userDeliveryAddress/edit" tags:"用户地址表" method:"post" summary:"用户地址表编辑接口"`

	UdId int64 `json:"ud_id"   v:"required#请输入地址编号"    dc:"地址编号"     ` // 地址编号
	DeliveryAddressAdd
}

type DeliveryAddressAddRes struct {
	UdId int64 `json:"ud_id"   v:"required#请输入地址编号"    dc:"地址编号"     ` // 地址编号
}

type DeliveryAddressEditRes entity.UserDeliveryAddress

type DeliveryAddressAddReq struct {
	g.Meta `path:"/manage/account/userDeliveryAddress/add" tags:"用户地址表" method:"post" summary:"用户地址表编辑接口"`

	UserId uint `json:"user_id"                ` // 用户编号
	DeliveryAddressAdd
}

type DeliveryAddressRemoveReq struct {
	g.Meta `path:"/manage/account/userDeliveryAddress/remove" tags:"用户地址表" method:"post" summary:"用户地址表删除接口"`
	UdId   uint `json:"ud_id"   v:"required#请输入地址编号"    dc:"地址编号"          ` // 地址编号
}

type DeliveryAddressRemoveRes struct {
}

type DeliveryAddressListReq struct {
	g.Meta `path:"/manage/account/userDeliveryAddress/list" tags:"用户地址表" method:"get" summary:"用户地址表列表接口"`
	ml.BaseList

	UserId uint `json:"user_id"                ` // 用户编号
}

type DeliveryAddressListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
