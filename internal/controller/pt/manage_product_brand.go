package pt

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pt"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
)

var (
	ProductBrand = cProductBrand{}
)

type cProductBrand struct{}

// =================== 管理端使用 =========================

// Tree 站点设置
func (c *cProductBrand) Tree(ctx context.Context, req *pt.ProductBrandTreeReq) (res []pt.ProductBrandTreeVoRes, err error) {
	// 获取配置列表
	var brandData, error = service.ProductBrand().Find(ctx, &do.ProductBrandListInput{
		BaseList: ml.BaseList{
			Sidx: dao.ProductBrand.Columns().BrandId,
			Sort: "ASC"},
	})

	if error != nil {
		err = error
	}

	//过滤分类编号
	categoryList := array.Column(brandData, dao.ProductBrand.Columns().CategoryId)
	var productCategory, errorType = service.ProductCategory().Gets(ctx, categoryList)

	if errorType != nil {
		err = errorType
	}

	for _, v := range productCategory {
		categoryId := v.CategoryId
		it := pt.ProductBrandTreeVoRes{}
		it.BrandId = v.CategoryId
		it.BrandName = v.CategoryName

		// 查询配置项列表
		itemList := make([]map[string]interface{}, 0)
		for _, v := range brandData {
			if categoryId == v.CategoryId {

				item := make(map[string]interface{})
				item["brand_id"] = v.BrandId
				item["brand_name"] = v.BrandName

				itemList = append(itemList, item)
			}
		}

		it.Children = itemList

		if len(itemList) > 0 {
			res = append(res, it)
		}
	}

	return
}

// List
func (c *cProductBrand) List(ctx context.Context, req *pt.ProductBrandListReq) (res *pt.ProductBrandListRes, err error) {
	input := do.ProductBrandListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.ProductBrand().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增品牌
func (c *cProductBrand) Add(ctx context.Context, req *pt.ProductBrandAddReq) (res *pt.ProductBrandEditRes, err error) {

	input := do.ProductBrand{}
	gconv.Scan(req, &input)

	var result, error = service.ProductBrand().Add(ctx, &input)
	//var result, error = service.ProductBrand().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &pt.ProductBrandEditRes{
		BrandId: result,
	}

	return
}

// Edit 编辑品牌
func (c *cProductBrand) Edit(ctx context.Context, req *pt.ProductBrandEditReq) (res *pt.ProductBrandEditRes, err error) {

	input := do.ProductBrand{}
	gconv.Scan(req, &input)

	var result, error = service.ProductBrand().Edit(ctx, &input)
	//var result, error = service.ProductBrand().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &pt.ProductBrandEditRes{
		BrandId: result,
	}

	return
}

// EditState 编辑品牌
func (c *cProductBrand) EditState(ctx context.Context, req *pt.ProductBrandEditStateReq) (res *pt.ProductBrandEditStateRes, err error) {
	input := do.ProductBrand{}
	gconv.Scan(req, &input)

	var result, error = service.ProductBrand().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &pt.ProductBrandEditStateRes{
		CategoryId: result,
	}

	return
}

// Remove 删除品牌
func (c *cProductBrand) Remove(ctx context.Context, req *pt.ProductBrandRemoveReq) (res *pt.ProductBrandRemoveRes, err error) {

	var _, error = service.ProductBrand().Remove(ctx, req.BrandId)

	/*
		input := do.ProductBrand{}
		input.ProductBrandTime = gtime.Now()
		input.BrandId = req.BrandId[0]
		input.ProductBrandSort = 0

		var _, error = service.ProductBrand().Edit(ctx, &input)
	*/

	if error != nil {
		err = error
	}

	res = &pt.ProductBrandRemoveRes{}

	return
}
