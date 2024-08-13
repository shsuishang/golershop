package trade

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

type DistributionReq struct {
	g.Meta `path:"/front/trade/distribution/index" tags:"用户分销数据概况" method:"get" summary:"用户分销数据概况"`
}

type DistributionRes map[string]interface{}

type GetCommissionTimelineReq struct {
	g.Meta `path:"/front/trade/distribution/getCommissionTimeline" tags:"推广中心" method:"get" summary:"推广中心"`
}

type GetCommissionTimelineRes map[string]interface{}

type GetCommissionNumReq struct {
	g.Meta `path:"/front/trade/distribution/getCommissionNum" tags:"粉丝及佣金" method:"get" summary:"粉丝及佣金"`
}

type GetCommissionNumRes map[string]interface{}

type ListsCommissionReq struct {
	g.Meta `path:"/front/trade/distribution/listsCommission" tags:"粉丝及佣金" method:"get" summary:"粉丝及佣金"`
	ml.BaseList
}

type ListsCommissionRes struct {
}
