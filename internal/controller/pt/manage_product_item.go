package pt

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pt"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ProductItem = cProductItem{}
)

type cProductItem struct{}

// =================== 管理端使用 =========================

// List
func (c *cProductItem) List(ctx context.Context, req *pt.ProductItemListReq) (res *pt.ProductItemListRes, err error) {
	input := do.ProductItemListInput{}
	gconv.Scan(req, &input)

	input.Sidx = dao.ProductItem.Columns().ItemId
	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.ProductItem().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Edit 编辑SKU
func (c *cProductItem) Edit(ctx context.Context, req *pt.ProductItemEditReq) (res *pt.ProductItemEditRes, err error) {

	input := do.ProductItem{}
	gconv.Scan(req, &input)

	var result, error = service.ProductItem().Edit(ctx, &input)
	//var result, error = service.ProductItem().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &pt.ProductItemEditRes{
		ItemId: result,
	}

	return
}

// EditState 编辑SKU
func (c *cProductItem) EditState(ctx context.Context, req *pt.ProductItemEditStateReq) (res *pt.ProductItemEditStateRes, err error) {
	input := do.ProductItem{}
	gconv.Scan(req, &input)

	var result, error = service.ProductItem().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &pt.ProductItemEditStateRes{
		ItemId: result,
	}

	return
}

// EditStock 更改库存
func (c *cProductItem) EditStock(ctx context.Context, req *pt.ProductEditStockReq) (res *pt.ProductEditStockRes, err error) {
	// 将请求参数转换为输入类型
	input := []*model.ProductEditStockInput{}
	gconv.Structs(req, &input)

	// 批量更新库存
	err = service.ProductItem().BatchEditStock(ctx, input)

	// 如果更新成功，返回成功响应
	if err != nil {
		return nil, err
	}

	// 否则返回错误信息
	return nil, err
}

// List
func (c *cProductItem) GetStockBillItems(ctx context.Context, req *pt.StockBillItemListReq) (res *pt.StockBillItemListRes, err error) {
	input := do.StockBillItemListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)
	var result, error = service.StockBillItem().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}
