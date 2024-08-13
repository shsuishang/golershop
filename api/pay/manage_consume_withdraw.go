package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo manage
type ConsumeWithdrawListReq struct {
	g.Meta `path:"/manage/pay/consumeWithdraw/list" tags:"提现申请表" method:"get" summary:"提现申请列表接口"`
	ml.BaseList

	UserId        uint `json:"user_id"     `
	WithdrawState uint `json:"withdraw_state"     ` // 提现状态(ENUM):0-申请中;1-提现通过;2-驳回;3-打款完成
}

type ConsumeWithdrawListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
