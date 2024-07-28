package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo manage
type ConsumeRecordListReq struct {
	g.Meta `path:"/manage/pay/consumeRecord/list" tags:"资金记录" method:"get" summary:"资金记录表接口"`
	ml.BaseList

	UserId          uint   `json:"user_id"    ` // 用户编号
	RecordTimeStart uint64 `json:"record_time_start" field:"record_time"  type:"GE"        dc:"时间-开始" `
	RecordTimeEnd   uint64 `json:"record_time_end"  field:"record_time"  type:"LE"        dc:"时间-结束"`
}

type ConsumeRecordListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
