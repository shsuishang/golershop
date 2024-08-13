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
	ProductSpecItem = cProductSpecItem{}
)

type cProductSpecItem struct{}

// =================== 管理端使用 =========================

func (c *cProductSpecItem) List(ctx context.Context, req *pt.ProductSpecItemListReq) (res *pt.ProductSpecItemListRes, err error) {
	specId := req.SpecId
	item := do.ProductSpecItem{SpecId: specId}

	if specId == 0 {
		item.SpecItemId = nil
	}

	var likes []*ml.WhereExt

	if req.SpecItemName != "" {
		likes = []*ml.WhereExt{{
			Column: dao.ProductSpecItem.Columns().SpecItemName,
			Val:    "%" + req.SpecItemName + "%",
		}}
	}

	var result, error = service.ProductSpecItem().List(ctx, &do.ProductSpecItemListInput{
		BaseList: ml.BaseList{Page: req.Page,
			Size:      req.Size,
			WhereLike: likes,
			Sidx:      dao.ProductSpecItem.Columns().SpecItemSort,
			Sort:      "ASC"},
		Where: item,
	})

	if error != nil {
		err = error
	}

	res = &pt.ProductSpecItemListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 新增规格项
func (c *cProductSpecItem) Add(ctx context.Context, req *pt.ProductSpecItemAddReq) (res *pt.ProductSpecItemEditRes, err error) {

	input := do.ProductSpecItem{}
	gconv.Scan(req, &input)

	input.SpecItemEnable = true
	var result, error = service.ProductSpecItem().Add(ctx, &input)
	//var result, error = service.ProductSpecItem().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &pt.ProductSpecItemEditRes{
		SpecItemId: result,
	}

	return
}

// Edit 编辑规格项
func (c *cProductSpecItem) Edit(ctx context.Context, req *pt.ProductSpecItemEditReq) (res *pt.ProductSpecItemEditRes, err error) {

	input := do.ProductSpecItem{}
	gconv.Scan(req, &input)

	var result, error = service.ProductSpecItem().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &pt.ProductSpecItemEditRes{
		SpecItemId: result,
	}

	return
}

// Remove 删除规格项
func (c *cProductSpecItem) Remove(ctx context.Context, req *pt.ProductSpecItemRemoveReq) (res *pt.ProductSpecItemRemoveRes, err error) {

	var _, error = service.ProductSpecItem().Remove(ctx, req.SpecItemId)

	if error != nil {
		err = error
	}

	res = &pt.ProductSpecItemRemoveRes{}

	return
}

// EditState 编辑规格项
func (c *cProductSpecItem) EditState(ctx context.Context, req *pt.ProductSpecItemEditStateReq) (res *pt.ProductSpecItemEditStateRes, err error) {
	input := do.ProductSpecItem{}
	gconv.Scan(req, &input)

	_, err = service.ProductSpecItem().Edit(ctx, &input)

	res = &pt.ProductSpecItemEditStateRes{
		SpecItemId: req.SpecItemId,
	}

	return
}
