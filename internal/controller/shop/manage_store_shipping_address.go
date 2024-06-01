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
	StoreShippingAddress = cStoreShippingAddress{}
)

type cStoreShippingAddress struct{}

// =================== 管理端使用 =========================

func (c *cStoreShippingAddress) List(ctx context.Context, req *shop.StoreShippingAddressListReq) (res *shop.StoreShippingAddressListRes, err error) {
	input := do.StoreShippingAddressListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.StoreShippingAddress().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增地址
func (c *cStoreShippingAddress) Add(ctx context.Context, req *shop.StoreShippingAddressAddReq) (res *shop.StoreShippingAddressEditRes, err error) {

	input := do.StoreShippingAddress{}
	gconv.Scan(req, &input)

	var result, error = service.StoreShippingAddress().Add(ctx, &input)
	//var result, error = service.StoreShippingAddress().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &shop.StoreShippingAddressEditRes{
		SsId: result,
	}

	return
}

// Edit 编辑地址
func (c *cStoreShippingAddress) Edit(ctx context.Context, req *shop.StoreShippingAddressEditReq) (res *shop.StoreShippingAddressEditRes, err error) {

	input := do.StoreShippingAddress{}
	gconv.Scan(req, &input)

	var result, error = service.StoreShippingAddress().Edit(ctx, &input)
	//var result, error = service.StoreShippingAddress().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &shop.StoreShippingAddressEditRes{
		SsId: result,
	}

	return
}

// Remove 删除地址
func (c *cStoreShippingAddress) Remove(ctx context.Context, req *shop.StoreShippingAddressRemoveReq) (res *shop.StoreShippingAddressRemoveRes, err error) {

	var _, error = service.StoreShippingAddress().Remove(ctx, req.SsId)

	/*
		input := do.StoreShippingAddress{}
		input.StoreShippingAddressTime = gtime.Now()
		input.StoreShippingAddressId = req.StoreShippingAddressId[0]
		input.StoreShippingAddressSort = 0

		var _, error = service.StoreShippingAddress().Edit(ctx, &input)
	*/

	if error != nil {
		err = error
	}

	res = &shop.StoreShippingAddressRemoveRes{}

	return
}
