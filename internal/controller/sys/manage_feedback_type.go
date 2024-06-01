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
	FeedbackType = cFeedbackType{}
)

type cFeedbackType struct{}

// =================== 管理端使用 =========================
func (c *cFeedbackType) List(ctx context.Context, req *sys.FeedbackTypeListReq) (res *sys.FeedbackTypeListRes, err error) {
	input := do.FeedbackTypeListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.FeedbackType().List(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.FeedbackTypeListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 新增类型
func (c *cFeedbackType) Add(ctx context.Context, req *sys.FeedbackTypeAddReq) (res *sys.FeedbackTypeEditRes, err error) {

	input := do.FeedbackType{}
	gconv.Scan(req, &input)

	var result, error = service.FeedbackType().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.FeedbackTypeEditRes{
		FeedbackTypeId: uint(result),
	}

	return
}

// Edit 编辑类型
func (c *cFeedbackType) Edit(ctx context.Context, req *sys.FeedbackTypeEditReq) (res *sys.FeedbackTypeEditRes, err error) {

	input := do.FeedbackType{}
	gconv.Scan(req, &input)

	var result, error = service.FeedbackType().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.FeedbackTypeEditRes{
		FeedbackTypeId: uint(result),
	}

	return
}

// Remove 删除类型
func (c *cFeedbackType) Remove(ctx context.Context, req *sys.FeedbackTypeRemoveReq) (res *sys.FeedbackTypeRemoveRes, err error) {

	idSlice := gstr.Split(req.FeedbackTypeId, ",")
	for _, contractTypeId := range idSlice {
		var _, error = service.FeedbackType().Remove(ctx, contractTypeId)

		if error != nil {
			err = error
		}
	}

	res = &sys.FeedbackTypeRemoveRes{}

	return
}

// EditState 编辑状态
func (c *cFeedbackType) EditState(ctx context.Context, req *sys.FeedbackTypeEditStateReq) (res *sys.FeedbackTypeEditStateRes, err error) {
	input := do.FeedbackType{}
	gconv.Scan(req, &input)

	_, err = service.FeedbackType().Edit(ctx, &input)

	res = &sys.FeedbackTypeEditStateRes{
		FeedbackTypeId: req.FeedbackTypeId,
	}

	return
}
