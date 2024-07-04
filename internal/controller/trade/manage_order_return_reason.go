package trade

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/trade"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	OrderReturnReason = cOrderReturnReason{}
)

type cOrderReturnReason struct{}

func (c *cOrderReturnReason) List(ctx context.Context, req *trade.OrderReturnReasonListReq) (res *trade.OrderReturnReasonListRes, err error) {
	input := do.OrderReturnReasonListInput{
		BaseList: ml.BaseList{
			Sort: ml.ORDER_BY_ASC,
			Sidx: dao.OrderReturnReason.Columns().ReturnReasonSort,
		},
	}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.OrderReturnReason().List(ctx, &input)
	if error != nil {
		err = error
	}

	res = &trade.OrderReturnReasonListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 新增反馈
func (c *cOrderReturnReason) Add(ctx context.Context, req *trade.OrderReturnReasonAddReq) (res *trade.OrderReturnReasonEditRes, err error) {

	input := do.OrderReturnReason{}
	gconv.Scan(req, &input)

	var result, error = service.OrderReturnReason().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &trade.OrderReturnReasonEditRes{
		ReturnReasonId: uint(result),
	}

	return
}

// Edit 编辑反馈
func (c *cOrderReturnReason) Edit(ctx context.Context, req *trade.OrderReturnReasonEditReq) (res *trade.OrderReturnReasonEditRes, err error) {

	input := do.OrderReturnReason{}
	gconv.Scan(req, &input)

	var result, error = service.OrderReturnReason().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &trade.OrderReturnReasonEditRes{
		ReturnReasonId: uint(result),
	}

	return
}

// Remove 删除反馈
func (c *cOrderReturnReason) Remove(ctx context.Context, req *trade.OrderReturnReasonRemoveReq) (res *trade.OrderReturnReasonEditRes, err error) {

	var _, error = service.OrderReturnReason().Remove(ctx, req.ReturnReasonId)

	if error != nil {
		err = error
	}

	res = &trade.OrderReturnReasonEditRes{}

	return
}
