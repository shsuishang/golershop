package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/sys"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	CrontabBase = cCrontabBase{}
)

type cCrontabBase struct{}

// =================== 管理端使用 =========================

// List 定时任务分页列表
func (c *cCrontabBase) List(ctx context.Context, req *sys.CrontabBaseListReq) (res *sys.CrontabBaseListRes, err error) {
	input := do.CrontabBaseListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.CrontabBase().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增定时任务
func (c *cCrontabBase) Add(ctx context.Context, req *sys.CrontabBaseAddReq) (res *sys.CrontabBaseEditRes, err error) {

	input := do.CrontabBase{}
	gconv.Scan(req, &input)

	var result, error = service.CrontabBase().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.CrontabBaseEditRes{
		CrontabId: result,
	}

	return
}

// Edit 编辑定时任务
func (c *cCrontabBase) Edit(ctx context.Context, req *sys.CrontabBaseEditReq) (res *sys.CrontabBaseEditRes, err error) {

	input := do.CrontabBase{}
	gconv.Scan(req, &input)

	var result, error = service.CrontabBase().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.CrontabBaseEditRes{
		CrontabId: result,
	}

	return
}

// Remove 删除定时任务
func (c *cCrontabBase) Remove(ctx context.Context, req *sys.CrontabBaseRemoveReq) (res *sys.CrontabBaseRemoveRes, err error) {
	var _, error = service.CrontabBase().Remove(ctx, req.CrontabId)

	if error != nil {
		err = error
	}

	res = &sys.CrontabBaseRemoveRes{}

	return
}

// EditState 编辑状态
func (c *cCrontabBase) EditState(ctx context.Context, req *sys.CrontabBaseEditStateReq) (res *sys.CrontabBaseEditStateRes, err error) {
	input := do.CrontabBase{}
	gconv.Scan(req, &input)

	_, err = service.CrontabBase().Edit(ctx, &input)

	res = &sys.CrontabBaseEditStateRes{
		CrontabId: req.CrontabId,
	}

	return
}
