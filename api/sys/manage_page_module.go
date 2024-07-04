package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type PageModuleAdd struct {
	PmId       uint        `json:"pm_id"       ` //
	PmName     string      `json:"pm_name"     ` // 模块名称
	PageId     int64       `json:"page_id"     ` //
	UserId     uint        `json:"user_id"     ` // 所属用户
	PmColor    string      `json:"pm_color"    ` // 颜色
	PmType     string      `json:"pm_type"     ` // 所在页面
	ModuleId   string      `json:"module_id"   ` // 模版
	PmUtime    *gtime.Time `json:"pm_utime"    ` // 更新时间
	PmOrder    uint        `json:"pm_order"    ` // 排序
	PmEnable   uint        `json:"pm_enable"   ` // 是否显示
	PmHtml     string      `json:"pm_html"     ` // 模块html代码
	PmJson     string      `json:"pm_json"     ` // 模块JSON代码(JSON)
	SubsiteId  uint        `json:"subsite_id"  ` // 所属分站Id:0-总站
	PmPosition string      `json:"pm_position" ` // column_left:content_top
}

type PageModuleEditReq struct {
	g.Meta `path:"/manage/sys/pageModule/edit" tags:"页面模块" method:"post" summary:"页面模块编辑接口"`

	PmId uint `json:"pm_id"       ` //
	PageModuleAdd
}

type PageModuleAddRes struct {
	PmId       interface{} `json:"pm_id"       ` //
	PmName     interface{} `json:"pm_name"     ` // 模块名称
	PageId     interface{} `json:"page_id"     ` //
	UserId     interface{} `json:"user_id"     ` // 所属用户
	PmColor    interface{} `json:"pm_color"    ` // 颜色
	PmType     interface{} `json:"pm_type"     ` // 所在页面
	ModuleId   interface{} `json:"module_id"   ` // 模版
	PmUtime    interface{} `json:"pm_utime"    ` // 更新时间
	PmOrder    interface{} `json:"pm_order"    ` // 排序
	PmEnable   interface{} `json:"pm_enable"   ` // 是否显示
	PmHtml     interface{} `json:"pm_html"     ` // 模块html代码
	PmJson     interface{} `json:"pm_json"     ` // 模块JSON代码(JSON)
	SubsiteId  interface{} `json:"subsite_id"  ` // 所属分站Id:0-总站
	PmPosition interface{} `json:"pm_position" ` // column_left:content_top
}
type PageModuleEditRes struct{}

type PageModuleAddReq struct {
	g.Meta `path:"/manage/sys/pageModule/add" tags:"页面模块" method:"post" summary:"页面模块编辑接口"`

	PageModuleAdd
}

type PageModuleRemoveReq struct {
	g.Meta `path:"/manage/sys/pageModule/remove" tags:"页面模块" method:"post" summary:"页面模块删除接口"`

	PmId uint `json:"pm_id"       ` //
}

type PageModuleRemoveRes struct{}

type PageModuleListReq struct {
	g.Meta `path:"/manage/sys/pageModule/list" tags:"页面模块" method:"get" summary:"页面模块列表接口"`
	ml.BaseList

	PageId int64 `json:"page_id"     ` //
}

type PageModuleListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type PageModuleReq struct {
	g.Meta `path:"/manage/sys/pageModule/listTpl" tags:"PC楼层模板" method:"get" summary:"PC楼层模板接口"`
}
type PageModuleRes struct {
	Items   interface{} `json:"items"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type PageModuleEnableReq struct {
	g.Meta `path:"/manage/sys/pageModule/enable" tags:"页面模块" method:"post" summary:"页面模块启用接口"`

	PmId     uint   `json:"pm_id"       ` //
	PmEnable string `json:"usable"   `    // 是否显示
	PageId   int64  `json:"page_id"     ` //
}
type PageModuleEnableRes struct{}

type PageModuleSortReq struct {
	g.Meta `path:"/manage/sys/pageModule/sort" tags:"页面模块" method:"post" summary:"页面模块启用接口"`

	PmId       uint   `json:"pm_id"       `        //
	PmIdString string `json:"pm_id_string"       ` //
}
type PageModuleSortRes struct{}
