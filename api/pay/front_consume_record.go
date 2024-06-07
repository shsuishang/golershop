package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/model/entity"
)

type RecordListReq struct {
	g.Meta `path:"/front/pay/consumeRecord/list" tags:"交易明细" method:"get" summary:"交易明细表-账户收支明细-资金流水表-账户金额变化流水分页查询"`

	*entity.ConsumeRecord
	ChangeType uint `json:"change_type" dc:"支出收入:1-支出;2-收入"`
}
type RecordListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
