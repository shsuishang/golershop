package cms

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/cms"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ArticleTag = cArticleTag{}
)

type cArticleTag struct{}

// =================== 管理端使用 =========================
func (c *cArticleTag) List(ctx context.Context, req *cms.ArticleTagListReq) (res *cms.ArticleTagListRes, err error) {
	input := do.ArticleTagListInput{}
	gconv.Scan(req, &input)

	var result, error = service.ArticleTag().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增
func (c *cArticleTag) Add(ctx context.Context, req *cms.ArticleTagAddReq) (res *cms.ArticleTagEditRes, err error) {
	input := do.ArticleTag{}
	gconv.Scan(req, &input)

	var result, error = service.ArticleTag().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &cms.ArticleTagEditRes{
		TagId: uint(result),
	}

	return
}

// Edit 编辑
func (c *cArticleTag) Edit(ctx context.Context, req *cms.ArticleTagEditReq) (res *cms.ArticleTagEditRes, err error) {

	input := do.ArticleTag{}
	gconv.Scan(req, &input)

	var result, error = service.ArticleTag().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &cms.ArticleTagEditRes{
		TagId: uint(result),
	}

	return
}

// Remove 删除
func (c *cArticleTag) Remove(ctx context.Context, req *cms.ArticleTagRemoveReq) (res *cms.ArticleTagRemoveRes, err error) {
	var _, error = service.ArticleTag().Remove(ctx, req.TagId)

	if error != nil {
		err = error
	}

	res = &cms.ArticleTagRemoveRes{}

	return
}
