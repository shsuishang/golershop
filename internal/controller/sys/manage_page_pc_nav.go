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
	PagePcNav = cPagePcNav{}
)

type cPagePcNav struct{}

func (c *cPagePcNav) List(ctx context.Context, req *sys.PagePcNavListReq) (res *sys.PagePcNavListRes, err error) {
	input := do.PagePcNavListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.PagePcNav().List(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.PagePcNavListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 新增反馈
func (c *cPagePcNav) Add(ctx context.Context, req *sys.PagePcNavAddReq) (res *sys.PagePcNavEditRes, err error) {

	input := do.PagePcNav{}
	gconv.Scan(req, &input)

	var result, error = service.PagePcNav().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.PagePcNavEditRes{
		NavId: uint(result),
	}

	return
}

// Edit 编辑反馈
func (c *cPagePcNav) Edit(ctx context.Context, req *sys.PagePcNavEditReq) (res *sys.PagePcNavEditRes, err error) {

	input := do.PagePcNav{}
	gconv.Scan(req, &input)

	var result, error = service.PagePcNav().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.PagePcNavEditRes{
		NavId: uint(result),
	}

	return
}

// Remove 删除反馈
func (c *cPagePcNav) Remove(ctx context.Context, req *sys.PagePcNavRemoveReq) (res *sys.PagePcNavEditRes, err error) {

	var _, error = service.PagePcNav().Remove(ctx, req.NavId)

	if error != nil {
		err = error
	}

	res = &sys.PagePcNavEditRes{}

	return
}

// EditState 编辑状态
func (c *cPagePcNav) EditState(ctx context.Context, req *sys.PagePcNavEditStateReq) (res *sys.PagePcNavEditStateRes, err error) {
	input := do.PagePcNav{}
	gconv.Scan(req, &input)

	result, err := service.PagePcNav().Edit(ctx, &input)
	if err != nil {
		return nil, err
	}

	res = &sys.PagePcNavEditStateRes{
		NavId: result,
	}

	return
}
