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
	PageCategoryNav = cPageCategoryNav{}
)

type cPageCategoryNav struct{}

func (c *cPageCategoryNav) List(ctx context.Context, req *sys.PageCategoryNavListReq) (res *sys.PageCategoryNavListRes, err error) {
	input := do.PageCategoryNavListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.PageCategoryNav().List(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.PageCategoryNavListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 新增反馈
func (c *cPageCategoryNav) Add(ctx context.Context, req *sys.PageCategoryNavAddReq) (res *sys.PageCategoryNavEditRes, err error) {

	input := do.PageCategoryNav{}
	gconv.Scan(req, &input)

	var result, error = service.PageCategoryNav().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.PageCategoryNavEditRes{
		CategoryNavId: uint(result),
	}

	return
}

// Edit 编辑反馈
func (c *cPageCategoryNav) Edit(ctx context.Context, req *sys.PageCategoryNavEditReq) (res *sys.PageCategoryNavEditRes, err error) {

	input := do.PageCategoryNav{}
	gconv.Scan(req, &input)

	var result, error = service.PageCategoryNav().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.PageCategoryNavEditRes{
		CategoryNavId: uint(result),
	}

	return
}

// Remove 删除反馈
func (c *cPageCategoryNav) Remove(ctx context.Context, req *sys.PageCategoryNavRemoveReq) (res *sys.PageCategoryNavEditRes, err error) {

	var _, error = service.PageCategoryNav().Remove(ctx, req.CategoryNavId)

	if error != nil {
		err = error
	}

	res = &sys.PageCategoryNavEditRes{}

	return
}

// EditState 编辑状态
func (c *cPageCategoryNav) EditState(ctx context.Context, req *sys.PageCategoryNavEditStateReq) (res *sys.PageCategoryNavEditStateRes, err error) {
	input := do.PageCategoryNav{}
	gconv.Scan(req, &input)

	result, err := service.PageCategoryNav().Edit(ctx, &input)
	if err != nil {
		return nil, err
	}

	res = &sys.PageCategoryNavEditStateRes{
		CategoryNavId: result,
	}

	return
}
