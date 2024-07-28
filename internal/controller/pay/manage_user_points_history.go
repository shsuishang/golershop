package pay

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pay"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	UserPointsHistory = cUserPointsHistory{}
)

type cUserPointsHistory struct{}

// =================== 管理端使用 =========================
func (c *cUserPointsHistory) List(ctx context.Context, req *pay.UserPointsHistoryListReq) (res *pay.UserPointsHistoryListRes, err error) {
	input := do.UserPointsHistoryListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.BaseList.WhereExt)

	var result, error = service.UserPointsHistory().GetList(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增
func (c *cUserPointsHistory) Add(ctx context.Context, req *pay.UserPointsHistoryAddReq) (res *pay.UserPointsHistoryEditRes, err error) {
	input := do.UserPointsHistory{}
	gconv.Scan(req, &input)

	var result, error = service.UserPointsHistory().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &pay.UserPointsHistoryEditRes{
		PointsLogId: uint(result),
	}

	return
}

// Edit 编辑
func (c *cUserPointsHistory) Edit(ctx context.Context, req *pay.UserPointsHistoryEditReq) (res *pay.UserPointsHistoryEditRes, err error) {

	input := do.UserPointsHistory{}
	gconv.Scan(req, &input)

	var result, error = service.UserPointsHistory().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &pay.UserPointsHistoryEditRes{
		PointsLogId: uint(result),
	}

	return
}

// Remove 删除
func (c *cUserPointsHistory) Remove(ctx context.Context, req *pay.UserPointsHistoryRemoveReq) (res *pay.UserPointsHistoryRemoveRes, err error) {
	var _, error = service.UserPointsHistory().Remove(ctx, req.PointsLogId)

	if error != nil {
		err = error
	}

	res = &pay.UserPointsHistoryRemoveRes{}

	return
}
