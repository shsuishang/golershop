package pt

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pt"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ProductTag = cProductTag{}
)

type cProductTag struct{}

// =================== 管理端使用 =========================

func (c *cProductTag) List(ctx context.Context, req *pt.ProductTagListReq) (res *pt.ProductTagListRes, err error) {
	input := do.ProductTagListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)
	input.Order = []*ml.BaseOrder{{Sidx: dao.ProductTag.Columns().ProductTagSort, Sort: ml.ORDER_BY_ASC}}
	var result, error = service.ProductTag().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增菜单
func (c *cProductTag) Add(ctx context.Context, req *pt.ProductTagAddReq) (res *pt.ProductTagEditRes, err error) {

	input := do.ProductTag{}
	gconv.Scan(req, &input)

	var result, error = service.ProductTag().Add(ctx, &input)
	//var result, error = service.ProductTag().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &pt.ProductTagEditRes{
		ProductTagId: result,
	}

	return
}

// Edit 编辑菜单
func (c *cProductTag) Edit(ctx context.Context, req *pt.ProductTagEditReq) (res *pt.ProductTagEditRes, err error) {

	input := do.ProductTag{}
	gconv.Scan(req, &input)

	var result, error = service.ProductTag().Edit(ctx, &input)
	//var result, error = service.ProductTag().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &pt.ProductTagEditRes{
		ProductTagId: result,
	}

	return
}

// Remove 删除菜单
func (c *cProductTag) Remove(ctx context.Context, req *pt.ProductTagRemoveReq) (res *pt.ProductTagRemoveRes, err error) {

	var _, error = service.ProductTag().Remove(ctx, req.ProductTagId)

	/*
		input := do.ProductTag{}
		input.ProductTagTime = gtime.Now()
		input.ProductTagId = req.ProductTagId[0]
		input.ProductTagSort = 0

		var _, error = service.ProductTag().Edit(ctx, &input)
	*/

	if error != nil {
		err = error
	}

	res = &pt.ProductTagRemoveRes{}

	return
}
