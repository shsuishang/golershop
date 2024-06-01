package trade

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/trade"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	OrderLogistics = cOrderLogistics{}
)

type cOrderLogistics struct{}

// =================== 管理端使用 =========================
func (c *cOrderLogistics) List(ctx context.Context, req *trade.OrderLogisticsListReq) (res *trade.OrderLogisticsListRes, err error) {
	input := do.OrderLogisticsListInput{}
	gconv.Scan(req, &input)

	var result, error = service.OrderLogistics().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增
func (c *cOrderLogistics) Add(ctx context.Context, req *trade.OrderLogisticsAddReq) (res *trade.OrderLogisticsEditRes, err error) {
	input := do.OrderLogistics{}
	gconv.Scan(req, &input)

	var result, error = service.OrderLogistics().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &trade.OrderLogisticsEditRes{
		OrderLogisticsId: uint64(result),
	}

	return
}

// Edit 编辑
func (c *cOrderLogistics) Edit(ctx context.Context, req *trade.OrderLogisticsEditReq) (res *trade.OrderLogisticsEditRes, err error) {

	input := do.OrderLogistics{}
	gconv.Scan(req, &input)

	var result, error = service.OrderLogistics().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &trade.OrderLogisticsEditRes{
		OrderLogisticsId: uint64(result),
	}

	return
}

// Remove 删除
func (c *cOrderLogistics) Remove(ctx context.Context, req *trade.OrderLogisticsRemoveReq) (res *trade.OrderLogisticsRemoveRes, err error) {
	var _, error = service.OrderLogistics().Remove(ctx, req.OrderLogisticsId)

	if error != nil {
		err = error
	}

	res = &trade.OrderLogisticsRemoveRes{}

	return
}
