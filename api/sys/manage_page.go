package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model"
)

// start fo manage
type PageBaseAdd struct {
}

type PageBaseEditReq struct {
	g.Meta `path:"/manage/sys/pageBase/edit" tags:"页面" method:"post" summary:"页面编辑接口"`

	PageId uint `json:"page_id"   v:"required#请输入页面编号"    dc:"页面编号"     `
	PageBaseAdd
}

type PageBaseEditRes struct {
	PageId interface{} `json:"page_id"   dc:"页面信息"`
}

type PageBaseListReq struct {
	g.Meta `path:"/manage/sys/pageBase/list" tags:"页面" method:"get" summary:"页面列表接口"`
	ml.BaseList

	PageName    string `json:"page_name"        `       // 页面名称
	StoreId     uint   `json:"store_id"         `       // 所属店铺
	SubsiteId   uint   `json:"subsite_id"       `       // 所属分站:0-总站
	PageBuildin uint   `json:"page_buildin"     `       // 是否内置(BOOL):0-否;1-是
	PageType    uint   `json:"page_type"   d:"3"      ` // 类型(ENUM):1-WAP;2-PC;3-APP
}

type PageBaseListRes struct {
	Items   interface{} `json:"items"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type PageMobileListReq struct {
	g.Meta `path:"/manage/sys/pageBase/listMobile" tags:"移动页面" method:"get" summary:"页面列表接口"`
	ml.BaseList

	PageName    string `json:"page_name"        `       // 页面名称
	StoreId     uint   `json:"store_id"         `       // 所属店铺
	SubsiteId   uint   `json:"subsite_id"       `       // 所属分站:0-总站
	PageBuildin uint   `json:"page_buildin"     `       // 是否内置(BOOL):0-否;1-是
	PageType    uint   `json:"page_type"   d:"3"      ` // 类型(ENUM):1-WAP;2-PC;3-APP
}

type PageMobileListRes struct {
	Items   []model.PageMobileVo `json:"items"`
	Center  string               `json:"center"`  // 分页号码
	Page    int                  `json:"page"`    // 分页号码
	Total   int                  `json:"total"`   // 总页数
	Records int                  `json:"records"` // 数据总数
	Size    int                  `json:"size"`    // 单页数量
}

type MobileEditReq struct {
	g.Meta `path:"/manage/sys/pageBase/saveMobile" tags:"移动页面" method:"post" summary:"移动端页面装修保存"`

	AppPageList     string `json:"app_page_list"     `     // JSON字符串
	AppMemberCenter string `json:"app_member_center"     ` // JSON字符串
}

type MobileEditRes struct {
}

type PageBaseEditStateReq struct {
	g.Meta `path:"/manage/sys/pageBase/editState" tags:"页面" method:"post" summary:"页面编辑接口"`

	PageId       uint `json:"page_id"   v:"required#请输入页面编号"    dc:"页面编号"     `
	PageIndex    bool `json:"page_index"   v:"required#请输入页面编号"    dc:"页面编号"     `
	PageGb       bool `json:"page_gb"   v:"required#请输入页面编号"    dc:"页面编号"     `
	PageActivity bool `json:"page_activity"   v:"required#请输入页面编号"    dc:"页面编号"     `
	PagePoint    bool `json:"page_point"   v:"required#请输入页面编号"    dc:"页面编号"     `
	PageGbs      bool `json:"page_gbs"   v:"required#请输入页面编号"    dc:"页面编号"     `
}

type PageBaseEditStateRes struct {
	PageId interface{} `json:"page_id"   dc:"页面信息"`
}

type PageBaseGetDataInfoReq struct {
	g.Meta `path:"/manage/sys/pageBase/getDataInfo" tags:"装修数据选择" method:"get" summary:"装修数据选择接口"`
	ml.BaseList

	Name    string `json:"name" dc:"页面名称"`                                                                                                // 页面名称
	Type    int    `json:"type" dc:"类型(ENUM):1-商品;2-店铺分类;3-APP;4-快捷入口;5-资讯分类;6-资讯;8-自定义页面;10-社区版块;11-帖子;12-拼团;14-秒杀;17-代金券;" default:"1"` // 类型(ENUM)
	StoreId int    `json:"store_id" dc:"店铺编号" default:"0"`                                                                                // 店铺
}

type PageBaseGetDataInfoRes struct {
	Items   interface{} `json:"items"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
