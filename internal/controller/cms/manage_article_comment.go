package cms

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/cms"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ArticleComment = cArticleComment{}
)

type cArticleComment struct{}

// =================== 管理端使用 =========================
func (c *cArticleComment) List(ctx context.Context, req *cms.ArticleCommentListReq) (res *cms.ArticleCommentListRes, err error) {
	input := do.ArticleCommentListInput{}
	gconv.Scan(req, &input)

	var result, error = service.ArticleComment().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增
func (c *cArticleComment) Add(ctx context.Context, req *cms.ArticleCommentAddReq) (res *cms.ArticleCommentEditRes, err error) {
	input := do.ArticleComment{}
	gconv.Scan(req, &input)
	input.CommentTime = gtime.Now()

	var result, error = service.ArticleComment().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &cms.ArticleCommentEditRes{
		CommentId: uint(result),
	}

	return
}

// Edit 编辑
func (c *cArticleComment) Edit(ctx context.Context, req *cms.ArticleCommentEditReq) (res *cms.ArticleCommentEditRes, err error) {

	input := do.ArticleComment{}
	gconv.Scan(req, &input)

	var result, error = service.ArticleComment().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &cms.ArticleCommentEditRes{
		CommentId: uint(result),
	}

	return
}

// Remove 删除
func (c *cArticleComment) Remove(ctx context.Context, req *cms.ArticleCommentRemoveReq) (res *cms.ArticleCommentRemoveRes, err error) {
	var _, error = service.ArticleComment().Remove(ctx, req.CommentId)

	if error != nil {
		err = error
	}

	res = &cms.ArticleCommentRemoveRes{}

	return
}

// Edit 编辑
func (c *cArticleComment) EditState(ctx context.Context, req *cms.ArticleCommentEditStateReq) (res *cms.ArticleCommentEditStateRes, err error) {

	input := do.ArticleComment{}
	gconv.Scan(req, &input)

	var result, error = service.ArticleComment().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &cms.ArticleCommentEditStateRes{
		CommentId: uint(result),
	}

	return
}
