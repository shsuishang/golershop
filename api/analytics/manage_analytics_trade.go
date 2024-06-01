package analytics

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/model"
)

type SalesAmountReq struct {
	g.Meta `path:"/manage/analytics/trade/getSalesAmount" tags:"交易统计" method:"get" summary:"交易销售额统计接口"`

	TradePaidStartTime int64 `json:"start_time"  type:"GE"     ` // 支付时间
	TradePaidEndTime   int64 `json:"end_time"  type:"LE"     `   // 支付时间
}

type SalesAmountRes struct {
	model.DashboardTopVo
}
