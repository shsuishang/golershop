package sys

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/sys"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	FeedbackCategory = cFeedbackCategory{}
)

type cFeedbackCategory struct{}

// =================== 管理端使用 =========================
func (c *cFeedbackCategory) List(ctx context.Context, req *sys.FeedbackCategoryListReq) (res *sys.FeedbackCategoryListRes, err error) {
	input := do.FeedbackCategoryListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.FeedbackCategory().List(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.FeedbackCategoryListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 新增分类
func (c *cFeedbackCategory) Add(ctx context.Context, req *sys.FeedbackCategoryAddReq) (res *sys.FeedbackCategoryEditRes, err error) {

	input := do.FeedbackCategory{}
	gconv.Scan(req, &input)

	var result, error = service.FeedbackCategory().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.FeedbackCategoryEditRes{
		FeedbackCategoryId: uint(result),
	}

	return
}

// Edit 编辑分类
func (c *cFeedbackCategory) Edit(ctx context.Context, req *sys.FeedbackCategoryEditReq) (res *sys.FeedbackCategoryEditRes, err error) {

	input := do.FeedbackCategory{}
	gconv.Scan(req, &input)

	var result, error = service.FeedbackCategory().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.FeedbackCategoryEditRes{
		FeedbackCategoryId: uint(result),
	}

	return
}

// Remove 删除分类
func (c *cFeedbackCategory) Remove(ctx context.Context, req *sys.FeedbackCategoryRemoveReq) (res *sys.FeedbackCategoryRemoveRes, err error) {

	idSlice := gstr.Split(req.FeedbackCategoryId, ",")
	for _, contractCategoryId := range idSlice {
		var _, error = service.FeedbackCategory().Remove(ctx, contractCategoryId)

		if error != nil {
			err = error
		}
	}

	res = &sys.FeedbackCategoryRemoveRes{}

	return
}

// EditState 编辑状态
func (c *cFeedbackCategory) EditState(ctx context.Context, req *sys.FeedbackCategoryEditStateReq) (res *sys.FeedbackCategoryEditStateRes, err error) {
	input := do.FeedbackCategory{}
	gconv.Scan(req, &input)

	_, err = service.FeedbackCategory().Edit(ctx, &input)

	res = &sys.FeedbackCategoryEditStateRes{
		FeedbackCategoryId: req.FeedbackCategoryId,
	}

	return
}
