package sys

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/sys"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	FeedbackBase = cFeedbackBase{}
)

type cFeedbackBase struct{}

// =================== 管理端使用 =========================
func (c *cFeedbackBase) List(ctx context.Context, req *sys.FeedbackBaseListReq) (res *sys.FeedbackBaseListRes, err error) {
	input := do.FeedbackBaseListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	input.Sidx = dao.FeedbackBase.Columns().FeedbackId
	input.Sort = ml.ORDER_BY_DESC

	var result, error = service.FeedbackBase().List(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.FeedbackBaseListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 新增反馈
func (c *cFeedbackBase) Add(ctx context.Context, req *sys.FeedbackBaseAddReq) (res *sys.FeedbackBaseEditRes, err error) {

	input := do.FeedbackBase{}
	gconv.Scan(req, &input)

	var result, error = service.FeedbackBase().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.FeedbackBaseEditRes{
		FeedbackId: uint(result),
	}

	return
}

// Edit 编辑反馈
func (c *cFeedbackBase) Edit(ctx context.Context, req *sys.FeedbackBaseEditReq) (res *sys.FeedbackBaseEditRes, err error) {

	input := do.FeedbackBase{}
	gconv.Scan(req, &input)

	var result, error = service.FeedbackBase().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.FeedbackBaseEditRes{
		FeedbackId: uint(result),
	}

	return
}

// Remove 删除反馈
func (c *cFeedbackBase) Remove(ctx context.Context, req *sys.FeedbackBaseRemoveReq) (res *sys.FeedbackBaseRemoveRes, err error) {

	idSlice := gstr.Split(req.FeedbackId, ",")
	for _, contractBaseId := range idSlice {
		var _, error = service.FeedbackBase().Remove(ctx, contractBaseId)

		if error != nil {
			err = error
		}
	}

	res = &sys.FeedbackBaseRemoveRes{}

	return
}

// EditState 编辑状态
func (c *cFeedbackBase) EditState(ctx context.Context, req *sys.FeedbackBaseEditStateReq) (res *sys.FeedbackBaseEditStateRes, err error) {
	input := do.FeedbackBase{}
	gconv.Scan(req, &input)

	_, err = service.FeedbackBase().Edit(ctx, &input)

	res = &sys.FeedbackBaseEditStateRes{
		FeedbackId: req.FeedbackId,
	}

	return
}
