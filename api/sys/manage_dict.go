package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

// =========================== 用户端使用 =============================

// =========================== 管理端使用 =============================
type DictAdd struct {
	DictId     string `json:"dict_id"   v:"required#请输入字典分类编号"    dc:"字典分类编号"     `
	DictName   string `json:"dict_name"    ` // 字典分类名称
	DictKey    string `json:"dict_key"     ` // 字典分类KEY
	DictSort   uint   `json:"dict_sort"    ` // 显示顺序:从小到大
	DictNote   string `json:"dict_note" `    // 字典分类备注
	DictEnable bool   `json:"dict_enable"  ` // 是否启用(BOOL):0-禁用;1-启用
}
type DictEditReq struct {
	g.Meta `path:"/manage/sys/dict/edit" tags:"字典分类" method:"post" summary:"字典分类详情接口"`

	DictAdd
}

type DictEditRes struct {
	DictId string `json:"dict_id"   dc:"字典分类信息"`
}

type DictAddReq struct {
	g.Meta `path:"/manage/sys/dict/add" tags:"字典分类" method:"post" summary:"字典分类详情接口"`

	DictAdd
}

type DictRemoveReq struct {
	g.Meta `path:"/manage/sys/dict/remove" tags:"字典分类" method:"post" summary:"字典分类删除接口"`
	DictId []string `json:"dict_id" v:"required#请输入字典分类编号"   dc:"字典分类信息"`
}

type DictRemoveRes struct {
}

type DictBaseListReq struct {
	g.Meta `path:"/manage/sys/dict/list" tags:"字典分类" method:"get" summary:"字典分类列表接口"`
	Page   int `json:"page"  d:"1"  v:"min:0#分页号码错误"  dc:"分页号码"`
	Size   int `json:"size" d:"10" v:"max:500#分页数量最大500条"  dc:"分页数量"`
}

type DictBaseListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

// ---------------------------- 字典项 -------------------------------

type DictItemAdd struct {
	DictItemId     string `json:"dict_item_id"   v:"required#请输入字典项目编号"    dc:"字典项目编号"`
	DictItemName   string `json:"dict_item_name"   `              // 字典项名称
	DictItemCode   string `json:"dict_item_code"   `              // 字典项值
	DictId         string `json:"dict_id" v:"required#请输入字典分类编号"` // 字典类型
	DictItemStatus bool   `json:"dict_item_status" `              // 是否使用(BOOL):0-未用;1-使用
	DictItemNote   string `json:"dict_item_note"   `              // 备注
	DictItemSort   uint   `json:"dict_item_sort"   `              // 显示顺序
	DictItemEnable bool   `json:"dict_item_enable" `              // 是否启用(BOOL):0-禁用;1-启用
}

type DictItemAddReq struct {
	g.Meta `path:"/manage/sys/dict/addItem" tags:"字典项目" method:"post" summary:"字典项目详情接口"`

	DictItemAdd
}

type DictItemEditReq struct {
	g.Meta `path:"/manage/sys/dict/editItem" tags:"字典项目" method:"post" summary:"字典项目详情接口"`

	DictItemAdd
}

type DictItemEditRes struct {
	DictItemId string `json:"dict_item_id"   dc:"字典项目信息"`
}

type DictItemRemoveReq struct {
	g.Meta     `path:"/manage/sys/dict/removeItem" tags:"字典项目" method:"post" summary:"字典项目删除接口"`
	DictItemId []string `json:"dict_item_id" v:"required#请输入字典项目编号"   dc:"字典项目信息"`
}

type DictItemRemoveRes struct {
}

type DictItemListReq struct {
	g.Meta `path:"/manage/sys/dict/listItem" tags:"字典项目" method:"get" summary:"字典项目列表接口"`
	Page   int    `json:"page"  d:"1"  v:"min:0#分页号码错误"  dc:"分页号码"`
	Size   int    `json:"size" d:"10" v:"max:500#分页数量最大500条"  dc:"分页数量"`
	DictId string `json:"dict_id"` // 字典类型
}

type DictItemListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
