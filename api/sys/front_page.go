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
