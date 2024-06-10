package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DownloadReq struct {
	g.Meta `path:"/manage/sys/release/download" tags:"发布" method:"post" summary:"下载代码"`

	ReleaseType  string `json:"release_type"   v:"required#请输入类型"    dc:"类型"     `
	UrlApi       string `json:"url_api"   v:"required#请输入页面Api"    dc:"Url"     `
	PrimaryColor string `json:"primary_color"   v:"required#请输入颜色"    dc:"颜色"     `
}

type DownloadRes string

type LicenceReq struct {
	g.Meta `path:"/manage/sys/release/getLicence" tags:"发布" method:"get" summary:"授权信息"`
}

type LicenceRes struct {
	LicenceStr   string `json:"licence_str"     dc:"类型"     `
	IsAuthorized bool   `json:"is_authorized"    dc:"类型"     `
}
