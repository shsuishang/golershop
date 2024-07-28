package cms

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/cms"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ArticleBase = cArticleBase{}
)

type cArticleBase struct{}

// =================== 管理端使用 =========================
func (c *cArticleBase) List(ctx context.Context, req *cms.ArticleBaseListReq) (res *cms.ArticleBaseListRes, err error) {
	input := do.ArticleBaseListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	input.Sidx = dao.ArticleBase.Columns().ArticleSort
	input.Sort = ml.ORDER_BY_ASC

	var result, error = service.ArticleBase().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增
func (c *cArticleBase) Add(ctx context.Context, req *cms.ArticleBaseAddReq) (res *cms.ArticleBaseEditRes, err error) {
	input := do.ArticleBase{}
	gconv.Scan(req, &input)
	input.ArticleAddTime = gtime.Now()

	var result, error = service.ArticleBase().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &cms.ArticleBaseEditRes{
		ArticleId: uint(result),
	}

	return
}

// Edit 编辑
func (c *cArticleBase) Edit(ctx context.Context, req *cms.ArticleBaseEditReq) (res *cms.ArticleBaseEditRes, err error) {

	input := do.ArticleBase{}
	gconv.Scan(req, &input)

	var result, error = service.ArticleBase().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &cms.ArticleBaseEditRes{
		ArticleId: uint(result),
	}

	return
}

// Remove 删除
func (c *cArticleBase) Remove(ctx context.Context, req *cms.ArticleBaseRemoveReq) (res *cms.ArticleBaseRemoveRes, err error) {
	var _, error = service.ArticleBase().Remove(ctx, req.ArticleId)

	if error != nil {
		err = error
	}

	res = &cms.ArticleBaseRemoveRes{}

	return
}

// RemoveBatch 删除
func (c *cArticleBase) RemoveBatch(ctx context.Context, req *cms.ArticleBaseRemoveBatchReq) (res *cms.ArticleBaseRemoveBatchRes, err error) {

	var _, error = service.ArticleBase().RemoveBatch(ctx, req.ArticleId)

	if error != nil {
		err = error
	}

	res = &cms.ArticleBaseRemoveBatchRes{}

	return
}

// Edit 编辑
func (c *cArticleBase) EditState(ctx context.Context, req *cms.ArticleBaseEditStateReq) (res *cms.ArticleBaseEditStateRes, err error) {

	input := do.ArticleBase{}
	gconv.Scan(req, &input)

	var result, error = service.ArticleBase().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &cms.ArticleBaseEditStateRes{
		ArticleId: result,
	}

	return
}
