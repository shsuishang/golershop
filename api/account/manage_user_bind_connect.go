package account

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

type UserBindConnectListReq struct {
	g.Meta `path:"/manage/account/userBindConnect/list" tags:"用户绑定" method:"get" summary:"用户绑定表分页查询"`
	ml.BaseList

	BindId       string      `json:"bind_id"            dc:"绑定标记"`                                          // 绑定标记
	BindType     int         `json:"bind_type"          dc:"绑定类型(EMUN):1-mobile;  2-email;   13-weixin公众号"` // 绑定类型
	UserId       int         `json:"user_id"            dc:"用户编号"`                                          // 用户编号
	BindTime     *gtime.Time `json:"bind_time"          dc:"绑定时间"`                                          // 绑定时间
	BindNickname string      `json:"bind_nickname"      dc:"用户名称"`                                          // 用户名称
	BindIcon     string      `json:"bind_icon"          dc:"用户图标"`                                          // 用户图标
	BindOpenid   string      `json:"bind_openid"        dc:"访问编号"`                                          // 访问编号
	BindUnionid  string      `json:"bind_unionid"       dc:"unionid"`                                       // unionid
	BindActive   bool        `json:"bind_active"        dc:"是否激活(BOOL):0-未激活;1-激活"`                         // 是否激活
	Sidx         string      `json:"sidx"               dc:"排序字段" default:"bind_id"`                        // 排序字段
}
type UserBindConnectListRes struct {
	Items   interface{} `json:"items"    dc:"用户标签列表页"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
