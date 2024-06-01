package pt

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type ProductAssistAdd struct {
	AssistId   uint   `json:"assist_id"   ` // 辅助属性编号
	AssistName string `json:"assist_name" ` // 辅助属性名称
	TypeId     uint   `json:"type_id"     ` // 所属类型编号
	AssistItem string `json:"assist_item" ` // 辅助属性值(DOT)
	AssistSort uint   `json:"assist_sort" ` // 排序
}
type ProductAssistEditReq struct {
	g.Meta `path:"/manage/pt/productAssist/edit" tags:"属性管理" method:"post" summary:"属性编辑接口"`

	AssistId uint `json:"assist_id" dc:"属性编号"     `
	ProductAssistAdd
}

type ProductAssistEditRes struct {
	AssistId interface{} `json:"assist_id"   dc:"属性信息"`
}

type ProductAssistAddReq struct {
	g.Meta `path:"/manage/pt/productAssist/add" tags:"属性管理" method:"post" summary:"属性编辑接口"`

	ProductAssistAdd
}

type ProductAssistRemoveReq struct {
	g.Meta   `path:"/manage/pt/productAssist/remove" tags:"属性管理" method:"post" summary:"属性删除接口"`
	AssistId uint `json:"assist_id" `
}

type ProductAssistRemoveRes struct {
}

type ProductAssistListReq struct {
	g.Meta     `path:"/manage/pt/productAssist/list" tags:"属性管理" method:"get" summary:"属性列表接口"`
	TypeId     uint   `json:"type_id"     ` // 所属类型编号
	AssistName string `json:"assist_name" ` // 辅助属性名称
	AssistId   uint   `json:"assist_id"   ` // 辅助属性编号
	ml.BaseList
}

type ProductAssistListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

// ---------------------------- 属性项 -------------------------------

type ProductAssistItemAdd struct {
	AssistItemId   uint   `json:"assist_item_id"   ` // 辅助属性编号
	AssistItemName string `json:"assist_item_name" ` // 辅助属性值名称
	AssistId       uint   `json:"assist_id"        ` // 所属属性编号
	AssistItemSort uint   `json:"assist_item_sort" ` // 辅助属性值排序
}
type ProductAssistItemEditReq struct {
	g.Meta `path:"/manage/pt/productAssistItem/edit" tags:"属性管理" method:"post" summary:"属性编辑接口"`

	AssistItemId uint `json:"assist_item_id"  dc:"属性编号"     `
	ProductAssistItemAdd
}

type ProductAssistItemEditRes struct {
	AssistItemId interface{} `json:"assist_item_id"   dc:"属性信息"`
}

type ProductAssistItemAddReq struct {
	g.Meta `path:"/manage/pt/productAssistItem/add" tags:"属性管理" method:"post" summary:"属性编辑接口"`

	ProductAssistItemAdd
}

type ProductAssistItemRemoveReq struct {
	g.Meta       `path:"/manage/pt/productAssistItem/remove" tags:"属性管理" method:"post" summary:"属性删除接口"`
	AssistItemId uint `json:"assist_item_id" v:"required#请输入属性编号"   dc:"属性信息"`
}

type ProductAssistItemRemoveRes struct {
}

type ProductAssistItemListReq struct {
	g.Meta         `path:"/manage/pt/productAssistItem/list" tags:"属性管理" method:"get" summary:"属性列表接口"`
	AssistId       uint   `json:"assist_id"        ` // 所属属性编号
	AssistItemId   uint   `json:"assist_item_id"   ` // 辅助属性编号
	AssistItemName string `json:"assist_item_name" ` // 辅助属性值名称
	ml.BaseList
}

type ProductAssistItemListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
