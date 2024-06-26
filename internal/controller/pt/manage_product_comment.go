package pt

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pt"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ProductComment = cProductComment{}
)

type cProductComment struct{}

func (c *cProductComment) List(ctx context.Context, req *pt.ProductCommentListReq) (res *pt.ProductCommentListRes, err error) {
	input := do.ProductCommentListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	result, err := service.ProductComment().GetList(ctx, &input)

	if err != nil {
		err = err
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增菜单
func (c *cProductComment) Add(ctx context.Context, req *pt.ProductCommentAddReq) (res *pt.ProductCommentEditRes, err error) {

	input := do.ProductComment{}
	gconv.Scan(req, &input)

	var result, error = service.ProductComment().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &pt.ProductCommentEditRes{
		CommentId: result,
	}

	return
}

// Edit 编辑菜单
func (c *cProductComment) Edit(ctx context.Context, req *pt.ProductCommentEditReq) (res *pt.ProductCommentEditRes, err error) {

	input := do.ProductComment{}
	gconv.Scan(req, &input)

	var result, error = service.ProductComment().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &pt.ProductCommentEditRes{
		CommentId: result,
	}

	return
}

// Remove 删除菜单
func (c *cProductComment) Remove(ctx context.Context, req *pt.ProductCommentRemoveReq) (res *pt.ProductCommentRemoveRes, err error) {

	var _, error = service.ProductComment().Remove(ctx, req.CommentId)

	if error != nil {
		err = error
	}

	res = &pt.ProductCommentRemoveRes{}

	return
}

// EditState 状态修改
func (c *cProductComment) EditState(ctx context.Context, req *pt.ProductCommentEditStateReq) (res *pt.ProductCommentEditStateRes, err error) {

	input := do.ProductComment{}
	gconv.Scan(req, &input)

	input.CommentId = req.CommentId

	_, err = service.ProductComment().Edit(ctx, &input)

	if err != nil {
		err = err
	}

	return
}
