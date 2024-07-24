package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type ExpressBaseAdd struct {
	ExpressId        uint   `json:"express_id"         ` // 快递编号
	ExpressName      string `json:"express_name"       ` // 快递名称
	ExpressPinyin    string `json:"express_pinyin"     ` // 快递编码
	ExpressPinyin100 string `json:"express_pinyin_100" ` // 快递公司100
	ExpressSite      string `json:"express_site"       ` // 快递官网
	ExpressIsFav     bool   `json:"express_is_fav"     ` // 是否常用
	ExpressOrder     uint   `json:"express_order"      ` // 快递排序:越小越靠前
	ExpressEnable    bool   `json:"express_enable"     ` // 启用状态(BOOL):0-禁用;1-启用
}
type ExpressBaseEditReq struct {
	g.Meta `path:"/manage/sys/expressBase/edit" tags:"快递公司" method:"post" summary:"快递公司编辑接口"`

	ExpressId uint `json:"express_id"   ` // 快递公司编号`
	ExpressBaseAdd
}

type ExpressBaseEditRes struct {
	ExpressId interface{} `json:"express_id"   dc:"快递公司信息"`
}

type ExpressBaseAddReq struct {
	g.Meta `path:"/manage/sys/expressBase/add" tags:"快递公司" method:"post" summary:"快递公司编辑接口"`

	ExpressBaseAdd
}

type ExpressBaseRemoveReq struct {
	g.Meta    `path:"/manage/sys/expressBase/remove" tags:"快递公司" method:"post" summary:"快递公司删除接口"`
	ExpressId uint `json:"express_id"   ` // 快递公司编号
}

type ExpressBaseRemoveRes struct {
}

type ExpressBaseListReq struct {
	g.Meta `path:"/manage/sys/expressBase/list" tags:"快递公司" method:"get" summary:"快递公司列表接口"`
	ml.BaseList

	ExpressId   uint   `json:"express_id"         `             // 快递编号
	ExpressName string `json:"express_name"  type:"LIKE"      ` // 快递名称
}

type ExpressBaseListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
type GetExpressListReq struct {
	g.Meta `path:"/manage/sys/expressBase/getExpressList" tags:"快递公司" method:"get" summary:"快递公司列表接口"`
	ml.BaseList
}

type GetExpressListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ExpressBaseEditStateReq struct {
	g.Meta `path:"/manage/sys/expressBase/editState" tags:"快递公司" method:"post" summary:"快递表-修改快递常用状态和启用状态"`

	ExpressId     uint `json:"express_id"   `       // 快递公司编号`
	ExpressIsFav  bool `json:"express_is_fav"     ` // 是否常用
	ExpressEnable bool `json:"express_enable"     ` // 启用状态(BOOL):0-禁用;1-启用
}

type ExpressBaseEditStateRes struct {
	ExpressId interface{} `json:"express_id"   dc:"快递公司信息"`
}
