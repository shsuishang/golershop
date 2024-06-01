package trade

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/trade"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	OrderReturn = cOrderReturn{}
)

type cOrderReturn struct{}

// =================== 管理端使用 =========================
// Detail 订单详情
func (c *cOrderReturn) Detail(ctx context.Context, req *trade.OrderReturnDetailReq) (res *trade.OrderReturnDetailRes, err error) {

	var result, error = service.Order().Detail(ctx, req.ReturnId)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// List
func (c *cOrderReturn) List(ctx context.Context, req *trade.OrderReturnListReq) (res *trade.OrderReturnListRes, err error) {
	input := do.OrderReturnListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.OrderReturn().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// EditState 编辑订单
func (c *cOrderReturn) EditState(ctx context.Context, req *trade.OrderReturnEditStateReq) (res *trade.OrderReturnEditStateRes, err error) {
	input := do.OrderReturn{}
	gconv.Scan(req, &input)

	var result, error = service.OrderReturn().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &trade.OrderReturnEditStateRes{
		OrderId: result,
	}

	return
}

// Cancel 取消订单
func (c *cOrderReturn) Cancel(ctx context.Context, req *trade.OrderReturnCancelReq) (res *trade.OrderReturnCancelRes, err error) {
	res = &trade.OrderReturnCancelRes{}

	return
}

// Review 审核订单
func (c *cOrderReturn) Review(ctx context.Context, req *trade.OrderReturnReviewReq) (res *trade.OrderReturnReviewRes, err error) {
	res = &trade.OrderReturnReviewRes{}

	if len(res.ReturnId) == 0 {
		err = errors.New("未更改到符合条件的订单！")
	} else {

	}

	return
}
