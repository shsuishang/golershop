package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/model"
)

type MobileIndexNavListReq struct {
	g.Meta `path:"/front/sys/page/getMobileIndexNav" tags:"page" method:"get" summary:"移动导航" dc:"移动导航"`
}

type MobileIndexNavListRes struct {
}

type PageDetailReq struct {
	g.Meta     `path:"/front/sys/page/getMobilePage" tags:"page" method:"get" summary:"移动首页" dc:"移动首页"`
	PageId     int64  `json:"page_id,omitempty" description:"页面编号"`
	PageIndex  string `json:"page_index,omitempty" description:"主页类型"`
	CategoryId int    `json:"category_id,omitempty" description:"分类编号"`
}

type PageDetailRes struct {
	model.PageDetail
}

type GetPcPageReq struct {
	g.Meta `path:"/front/sys/page/getPcPage" tags:"page" method:"get" summary:"读取PC页面"`

	PageId     int64  `json:"page_id,omitempty"  description:"页面编号"`
	PageIndex  string `json:"page_index,omitempty" description:"主页类型"`
	CategoryId int    `json:"category_id,omitempty" description:"分类编号"`
}

type GetPcPageRes struct{}

type PcLayoutReq struct {
	g.Meta `path:"/front/sys/page/pcLayout" tags:"page" method:"get" summary:"PC头尾数据"`
}

type PcLayoutRes struct {
	CategoryNav   interface{} `json:"category_nav"`
	FooterArticle interface{} `json:"footer_article"`
	PagePcNav     interface{} `json:"page_pc_nav"`
	UserAvatar    interface{} `json:"user_avatar"`
	UserNickname  interface{} `json:"user_nickname"`
}
