package marketing

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/marketing"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ActivityBase = cActivityBase{}
)

type cActivityBase struct{}

// =================== 管理端使用 =========================

// List 活动分页列表
func (c *cActivityBase) List(ctx context.Context, req *marketing.ActivityBaseListReq) (res *marketing.ActivityBaseListRes, err error) {
	input := do.ActivityBaseListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.ActivityBase().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增活动
func (c *cActivityBase) Add(ctx context.Context, req *marketing.ActivityBaseAddReq) (res *marketing.ActivityBaseEditRes, err error) {

	input := do.ActivityBase{}
	gconv.Scan(req, &input)

	var result, error = service.ActivityBase().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &marketing.ActivityBaseEditRes{
		ActivityId: result,
	}

	return
}

// Edit 编辑活动
func (c *cActivityBase) Edit(ctx context.Context, req *marketing.ActivityBaseEditReq) (res *marketing.ActivityBaseEditRes, err error) {

	input := do.ActivityBase{}
	gconv.Scan(req, &input)

	var result, error = service.ActivityBase().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &marketing.ActivityBaseEditRes{
		ActivityId: result,
	}

	return
}

// Remove 删除活动
func (c *cActivityBase) Remove(ctx context.Context, req *marketing.ActivityBaseRemoveReq) (res *marketing.ActivityBaseRemoveRes, err error) {
	var _, error = service.ActivityBase().Remove(ctx, req.ActivityId)

	if error != nil {
		err = error
	}

	res = &marketing.ActivityBaseRemoveRes{}

	return
}
