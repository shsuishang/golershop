package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type PagePcNavAdd struct {
	NavId           uint   `json:"nav_id"            ` // 导航编号
	NavType         uint   `json:"nav_type"          ` // 类别(ENUM):0-自定义导航;1-商品分类;2-文章导航;3-活动导航
	NavItemId       uint   `json:"nav_item_id"       ` // 类别内容编号
	NavTitle        string `json:"nav_title"         ` // 导航标题
	NavUrl          string `json:"nav_url"           ` // 导航链接(HTML)
	NavPosition     uint   `json:"nav_position"      ` // 导航位置(ENUM):0-头部;1-中部;2-底部
	NavTargetBlank  bool   `json:"nav_target_blank"  ` // 是否以新窗口打开(BOOL):1-是; 0-否
	NavImage        string `json:"nav_image"         ` // 导航图片
	NavDropdownMenu string `json:"nav_dropdown_menu" ` // 导航下拉内容(HTML)
	NavOrder        uint   `json:"nav_order"         ` // 排序
	NavEnable       bool   `json:"nav_enable"        ` // 是否启用(BOOL):1-是; 0-否
	NavBuildin      uint   `json:"nav_buildin"       ` // 系统内置(ENUM):1-是; 0-否
}

type PagePcNavEditReq struct {
	g.Meta `path:"/manage/sys/pagePcNav/edit" tags:"页面导航表" method:"post" summary:"页面导航表编辑接口"`

	NavId uint `json:"nav_id"            ` // 导航编号
	PagePcNavAdd
}

type PagePcNavEditRes struct {
	NavId uint `json:"nav_id"            ` // 导航编号
}

type PagePcNavAddReq struct {
	g.Meta `path:"/manage/sys/pagePcNav/add" tags:"页面导航表" method:"post" summary:"页面导航表编辑接口"`

	PagePcNavAdd
}

type PagePcNavRemoveReq struct {
	g.Meta `path:"/manage/sys/pagePcNav/remove" tags:"页面导航表" method:"post" summary:"页面导航表删除接口"`

	NavId uint `json:"nav_id"            ` // 导航编号
}

type PagePcNavRemoveRes struct {
}

type PagePcNavListReq struct {
	g.Meta `path:"/manage/sys/pagePcNav/list" tags:"页面导航表" method:"get" summary:"页面导航表列表接口"`
	ml.BaseList

	NavTitle string `json:"nav_title"   type:"LIKE"       ` // 导航标题
}

type PagePcNavListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type PagePcNavEditStateReq struct {
	g.Meta `path:"/manage/sys/pagePcNav/editState" tags:"页面导航表" method:"post" summary:"页面导航表状态编辑接口"`

	NavId     uint `json:"nav_id"            ` // 导航编号
	NavEnable bool `json:"nav_enable"        ` // 是否启用(BOOL):1-是; 0-否
}

type PagePcNavEditStateRes struct {
	NavId int64 `json:"nav_id"            ` // 导航编号
}
