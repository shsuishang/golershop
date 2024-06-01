package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

type LogErrorListReq struct {
	g.Meta `path:"/manage/sys/logError/list" tags:"日志管理" method:"get" summary:"日志列表接口"`
	ml.BaseList

	LogErrorType string `json:"log_error_type" ` // 日志类型
	LogErrorName string `json:"log_error_name" ` // 日志名称
	LogErrorInfo string `json:"log_error_info" ` // 日志内容
}

type LogErrorListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
