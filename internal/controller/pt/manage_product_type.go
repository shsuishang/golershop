package pt

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pt"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ProductType = cProductType{}
)

type cProductType struct{}

// =================== 管理端使用 =========================

func (c *cProductType) List(ctx context.Context, req *pt.ProductTypeListReq) (res *pt.ProductTypeListRes, err error) {
	input := do.ProductTypeListInput{}
	gconv.Scan(req, &input)

	if !g.IsEmpty(req.TypeName) {
		var likes = []*ml.WhereExt{{
			Column: dao.ProductType.Columns().TypeName,
			Val:    "%" + req.TypeName + "%",
			Symbol: ml.LIKE,
		}}

		input.WhereExt = likes
	}

	var result, error = service.ProductType().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增菜单
func (c *cProductType) Add(ctx context.Context, req *pt.ProductTypeAddReq) (res *pt.ProductTypeEditRes, err error) {

	input := do.ProductType{}
	gconv.Scan(req, &input)
	var result, error = service.ProductType().Add(ctx, &input)
	//var result, error = service.ProductType().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &pt.ProductTypeEditRes{
		TypeId: result,
	}

	return
}

// Edit 编辑菜单
func (c *cProductType) Edit(ctx context.Context, req *pt.ProductTypeEditReq) (res *pt.ProductTypeEditRes, err error) {

	input := do.ProductType{}
	gconv.Scan(req, &input)

	var result, error = service.ProductType().Edit(ctx, &input)
	//var result, error = service.ProductType().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &pt.ProductTypeEditRes{
		TypeId: result,
	}

	return
}

// Remove 删除菜单
func (c *cProductType) Remove(ctx context.Context, req *pt.ProductTypeRemoveReq) (res *pt.ProductTypeRemoveRes, err error) {

	var _, error = service.ProductType().Remove(ctx, req.TypeId)

	/*
		input := do.ProductType{}
		input.ProductTypeTime = gtime.Now()
		input.TypeId = req.TypeId[0]
		input.ProductTypeSort = 0

		var _, error = service.ProductType().Edit(ctx, &input)
	*/

	if error != nil {
		err = error
	}

	res = &pt.ProductTypeRemoveRes{}

	return
}

func (c *cProductType) Info(ctx context.Context, req *pt.ProductTypeInfoReq) (res *pt.ProductTypeInfoRes, err error) {
	var result, error = service.ProductType().Info(ctx, req.TypeId)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}
