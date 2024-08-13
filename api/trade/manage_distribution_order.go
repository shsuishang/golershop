package trade

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

type DistributionOrderListReq struct {
	g.Meta `path:"/manage/trade/distributionOrder/list" tags:"推广订单收益详情表分页查询" method:"get" summary:"推广订单收益详情表分页查询"`
	ml.BaseList
}

type DistributionOrderListRes struct {
}
