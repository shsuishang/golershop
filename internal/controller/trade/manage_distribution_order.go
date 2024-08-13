package trade

import (
	"context"
	"golershop.cn/api/trade"
)

var (
	DistributionOrder = cDistributionOrder{}
)

type cDistributionOrder struct{}

// List
func (c *cDistributionOrder) List(ctx context.Context, req *trade.DistributionOrderListReq) (res *trade.DistributionOrderListRes, err error) {

	return
}
