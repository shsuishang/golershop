package pt

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pt"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	Product = cProduct{}
)

type cProduct struct{}

// Detail 商品详情
func (c *cProduct) Detail(ctx context.Context, req *pt.ProductDetailReq) (res *pt.ProductDetailRes, err error) {
	input := model.ProductDetailInput{}
	gconv.Scan(req, &input)

	// 获取当前登录用户
	user := service.BizCtx().GetUser(ctx)

	if user != nil {
		input.UserId = user.UserId
	}

	var result, error = service.ProductIndex().Detail(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)
	return
}

// List 商品列表
func (c *cProduct) List(ctx context.Context, req *pt.ListReq) (res *pt.ListRes, err error) {
	if !g.IsEmpty(req.CategoryId) {
		categoryLeafs, _ := service.ProductCategory().GetCategoryLeafs(ctx, req.CategoryId[0])
		if len(*categoryLeafs) > 0 {
			req.CategoryId = *categoryLeafs
		} else {
		}
	}

	req.ProductStateId = consts.PRODUCT_STATE_NORMAL

	input := do.ProductIndexListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInput(req, &input.Where, &input.WhereExt)

	var result, error = service.ProductIndex().GetList(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// ListItem
func (c *cProductItem) ListItem(ctx context.Context, req *pt.ItemListReq) (res *pt.ProductItemListRes, err error) {
	// 如果分类ID不为空，则处理分类ID
	if !g.IsEmpty(req.CategoryId) {
		categoryLeafs, _ := service.ProductCategory().GetCategoryLeafs(ctx, gconv.Uint(req.CategoryId))
		if len(*categoryLeafs) > 0 {
			req.CategoryIds = *categoryLeafs
		} else {
		}
	}

	var result, error = service.ProductIndex().ListItem(ctx, req)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// ListCategory
func (c *cProductCategory) ListCategory(ctx context.Context, req *pt.CategoryListReq) (res *pt.CategoryListRes, err error) {
	input := do.ProductCategoryListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)
	input.Where.CategoryParentId = req.CategoryParentId

	var result, error = service.ProductCategory().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// CategoryTree
func (c *cProductCategory) TreeCategory(ctx context.Context, req *pt.CategoryTreeReq) (res pt.CategoryTreeRes, err error) {
	input := do.ProductCategoryListInput{}
	gconv.Scan(req, &input)

	if !g.IsEmpty(req.CategoryName) {
		input.Where.CategoryName = req.CategoryName
	}

	input.Where.CategoryIsEnable = true

	var result, error = service.ProductCategory().GetTree(ctx, &input, req.CategoryParentId)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	if g.IsEmpty(res) {
		res = make([]*model.CategoryTreeNode, 0)
	}

	return
}

func (c *cProductCategory) GetSearchFilter(ctx context.Context, req *pt.SearchFilterReq) (res *pt.SearchFilterRes, err error) {

	res, err = service.ProductCategory().GetSearchFilter(ctx, req.CategoryId)

	return
}
