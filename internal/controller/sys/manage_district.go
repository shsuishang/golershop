package sys

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/sys"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

// Tree
func (c *cConfig) DistrictBaseTree(ctx context.Context, req *sys.DistrictBaseTreeReq) (res sys.DistrictBaseTreeRes, err error) {
	input := do.DistrictBaseListInput{}
	gconv.Scan(req, &input)

	if !g.IsEmpty(req.DistrictName) {
		var likes = []*ml.WhereExt{{
			Column: dao.DistrictBase.Columns().DistrictName,
			Val:    "%" + req.DistrictName + "%",
			Symbol: ml.LIKE,
		}}

		input.WhereExt = likes
	}

	res, err = service.DistrictBase().Tree(ctx, &input)

	if err != nil {
	}

	return
}

// Add 新增菜单
func (c *cConfig) DistrictAdd(ctx context.Context, req *sys.DistrictBaseAddReq) (res *sys.DistrictBaseEditRes, err error) {

	input := do.DistrictBase{}
	gconv.Scan(req, &input)

	var result, error = service.DistrictBase().Add(ctx, &input)
	//var result, error = service.DistrictBase().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &sys.DistrictBaseEditRes{
		DistrictId: result,
	}

	return
}

// Edit 编辑菜单
func (c *cConfig) DistrictEdit(ctx context.Context, req *sys.DistrictBaseEditReq) (res *sys.DistrictBaseEditRes, err error) {

	input := do.DistrictBase{}
	gconv.Scan(req, &input)

	var result, error = service.DistrictBase().Edit(ctx, &input)
	//var result, error = service.DistrictBase().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &sys.DistrictBaseEditRes{
		DistrictId: result,
	}

	return
}

// Remove 删除菜单
func (c *cConfig) DistrictRemove(ctx context.Context, req *sys.DistrictBaseRemoveReq) (res *sys.DistrictBaseRemoveRes, err error) {

	var _, error = service.DistrictBase().Remove(ctx, req.DistrictId)

	/*
		input := do.DistrictBase{}
		input.DistrictBaseTime = gtime.Now()
		input.DistrictBaseId = req.DistrictBaseId[0]
		input.DistrictBaseSort = 0

		var _, error = service.DistrictBase().Edit(ctx, &input)
	*/

	if error != nil {
		err = error
	}

	res = &sys.DistrictBaseRemoveRes{}

	return
}
