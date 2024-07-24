package shop

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/shop"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	StoreExpressLogistics = cStoreExpressLogistics{}
)

type cStoreExpressLogistics struct{}

// =================== 管理端使用 =========================

func (c *cStoreExpressLogistics) List(ctx context.Context, req *shop.StoreExpressLogisticsListReq) (res *shop.StoreExpressLogisticsListRes, err error) {
	input := do.StoreExpressLogisticsListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.StoreExpressLogistics().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增菜单
func (c *cStoreExpressLogistics) Add(ctx context.Context, req *shop.StoreExpressLogisticsAddReq) (res *shop.StoreExpressLogisticsEditRes, err error) {

	input := do.StoreExpressLogistics{}
	gconv.Scan(req, &input)

	var result, error = service.StoreExpressLogistics().Add(ctx, &input)
	//var result, error = service.StoreExpressLogistics().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &shop.StoreExpressLogisticsEditRes{
		LogisticsId: result,
	}

	return
}

// Edit 编辑菜单
func (c *cStoreExpressLogistics) Edit(ctx context.Context, req *shop.StoreExpressLogisticsEditReq) (res *shop.StoreExpressLogisticsEditRes, err error) {

	input := do.StoreExpressLogistics{}
	gconv.Scan(req, &input)

	var result, error = service.StoreExpressLogistics().Edit(ctx, &input)
	//var result, error = service.StoreExpressLogistics().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &shop.StoreExpressLogisticsEditRes{
		LogisticsId: result,
	}

	return
}

// Remove 删除菜单
func (c *cStoreExpressLogistics) Remove(ctx context.Context, req *shop.StoreExpressLogisticsRemoveReq) (res *shop.StoreExpressLogisticsRemoveRes, err error) {

	var _, error = service.StoreExpressLogistics().Remove(ctx, req.LogisticsId)

	/*
		input := do.StoreExpressLogistics{}
		input.StoreExpressLogisticsTime = gtime.Now()
		input.StoreExpressLogisticsId = req.StoreExpressLogisticsId[0]
		input.StoreExpressLogisticsSort = 0

		var _, error = service.StoreExpressLogistics().Edit(ctx, &input)
	*/

	if error != nil {
		err = error
	}

	res = &shop.StoreExpressLogisticsRemoveRes{}

	return
}

// EditState 修改状态
func (c *cStoreExpressLogistics) EditState(ctx context.Context, req *shop.StoreExpressLogisticsEditStateReq) (res *shop.StoreExpressLogisticsEditStateRes, err error) {

	input := do.StoreExpressLogistics{}
	gconv.Scan(req, &input)

	var result, error = service.StoreExpressLogistics().Edit(ctx, &input)
	//var result, error = service.StoreExpressLogistics().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &shop.StoreExpressLogisticsEditStateRes{
		LogisticsId: result,
	}

	return
}
