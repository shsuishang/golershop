package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo manage
type LogActionListReq struct {
	g.Meta `path:"/manage/sys/logAction/list" tags:"日志管理" method:"get" summary:"日志列表接口"`
	ml.BaseList

	UserId   uint   `json:"user_id"        ` // 玩家编号
	LogName  string `json:"log_name"       ` // 请求名称
	LogUrl   string `json:"log_url"        ` // 请求接口
	LogParam string `json:"log_param"      ` // 请求的参数
}

type LogActionListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
