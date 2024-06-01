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
	"golershop.cn/utility/array"
)

var (
	ProductSpec = cProductSpec{}
)

type cProductSpec struct{}

// =================== 管理端使用 =========================

// Tree 站点设置
func (c *cProductSpec) Tree(ctx context.Context, req *pt.ProductSpecTreeReq) (res []pt.ProductSpecTreeVoRes, err error) {
	// 获取配置列表
	var specData, error = service.ProductSpec().Find(ctx, &do.ProductSpecListInput{
		BaseList: ml.BaseList{
			Sidx: dao.ProductSpec.Columns().SpecId,
			Sort: "ASC"},
	})

	if error != nil {
		err = error
	}

	//过滤分类编号
	/*
		var categoryList []uint

		for _, v := range specData {
			categoryId := v.CategoryId
			// 加入数组
			categoryList = append(categoryList, categoryId)
		}
	*/

	//categoryList := utility.ColumnAny[*entity.ProductSpec, uint](specData, "CategoryId")
	categoryList := array.Column(specData, "CategoryId")

	var productCategory, errorType = service.ProductCategory().Gets(ctx, categoryList)
	if errorType != nil {
		err = errorType
	}

	for _, v := range productCategory {
		categoryId := v.CategoryId
		it := pt.ProductSpecTreeVoRes{}
		it.SpecId = v.CategoryId
		it.SpecName = v.CategoryName

		itemList := make([]map[string]interface{}, 0)
		for _, v := range specData {
			if categoryId == v.CategoryId {

				items := make(map[string]interface{})
				items["spec_id"] = v.SpecId
				items["spec_name"] = v.SpecName

				itemList = append(itemList, items)
			}
		}
		it.Children = itemList
		res = append(res, it)
	}
	return
}

// List
func (c *cProductSpec) List(ctx context.Context, req *pt.ProductSpecListReq) (res *pt.ProductSpecListRes, err error) {
	input := do.ProductSpecListInput{}
	gconv.Scan(req, &input)

	if !g.IsEmpty(req.SpecName) {
		var likes = []*ml.WhereExt{{
			Column: dao.ProductSpec.Columns().SpecName,
			Val:    "%" + req.SpecName + "%",
			Symbol: ml.LIKE,
		}}

		input.WhereExt = likes
	}

	var result, error = service.ProductSpec().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增菜单
func (c *cProductSpec) Add(ctx context.Context, req *pt.ProductSpecAddReq) (res *pt.ProductSpecEditRes, err error) {

	input := do.ProductSpec{}
	gconv.Scan(req, &input)

	var result, error = service.ProductSpec().Add(ctx, &input)
	//var result, error = service.ProductSpec().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &pt.ProductSpecEditRes{
		SpecId: result,
	}

	return
}

// Edit 编辑菜单
func (c *cProductSpec) Edit(ctx context.Context, req *pt.ProductSpecEditReq) (res *pt.ProductSpecEditRes, err error) {

	input := do.ProductSpec{}
	gconv.Scan(req, &input)

	var result, error = service.ProductSpec().Edit(ctx, &input)
	//var result, error = service.ProductSpec().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &pt.ProductSpecEditRes{
		SpecId: result,
	}

	return
}

// Remove 删除菜单
func (c *cProductSpec) Remove(ctx context.Context, req *pt.ProductSpecRemoveReq) (res *pt.ProductSpecRemoveRes, err error) {

	var _, error = service.ProductSpec().Remove(ctx, req.SpecId)

	/*
		input := do.ProductSpec{}
		input.ProductSpecTime = gtime.Now()
		input.SpecId = req.SpecId[0]
		input.ProductSpecSort = 0

		var _, error = service.ProductSpec().Edit(ctx, &input)
	*/

	if error != nil {
		err = error
	}

	res = &pt.ProductSpecRemoveRes{}

	return
}
