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
	StoreTransportType = cStoreTransportType{}
)

type cStoreTransportType struct{}

// =================== 管理端使用 =========================

func (c *cStoreTransportType) List(ctx context.Context, req *shop.StoreTransportTypeListReq) (res *shop.StoreTransportTypeListRes, err error) {
	input := do.StoreTransportTypeListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.StoreTransportType().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增菜单
func (c *cStoreTransportType) Add(ctx context.Context, req *shop.StoreTransportTypeAddReq) (res *shop.StoreTransportTypeEditRes, err error) {

	input := do.StoreTransportType{}
	gconv.Scan(req, &input)

	var result, error = service.StoreTransportType().Add(ctx, &input)
	//var result, error = service.StoreTransportType().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &shop.StoreTransportTypeEditRes{
		TransportTypeId: result,
	}

	return
}

// Edit 编辑菜单
func (c *cStoreTransportType) Edit(ctx context.Context, req *shop.StoreTransportTypeEditReq) (res *shop.StoreTransportTypeEditRes, err error) {

	input := do.StoreTransportType{}
	gconv.Scan(req, &input)

	var result, error = service.StoreTransportType().Edit(ctx, &input)
	//var result, error = service.StoreTransportType().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &shop.StoreTransportTypeEditRes{
		TransportTypeId: result,
	}

	return
}

// Remove 删除菜单
func (c *cStoreTransportType) Remove(ctx context.Context, req *shop.StoreTransportTypeRemoveReq) (res *shop.StoreTransportTypeRemoveRes, err error) {

	var _, error = service.StoreTransportType().Remove(ctx, req.TransportTypeId)

	/*
		input := do.StoreTransportType{}
		input.StoreTransportTypeTime = gtime.Now()
		input.StoreTransportTypeId = req.StoreTransportTypeId[0]
		input.StoreTransportTypeSort = 0

		var _, error = service.StoreTransportType().Edit(ctx, &input)
	*/

	if error != nil {
		err = error
	}

	res = &shop.StoreTransportTypeRemoveRes{}

	return
}
