package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/sys"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ExpressBase = cExpressBase{}
)

type cExpressBase struct{}

// =================== 管理端使用 =========================

func (c *cExpressBase) GetExpressList(ctx context.Context, req *sys.GetExpressListReq) (res *sys.GetExpressListRes, err error) {
	input := do.ExpressBaseListInput{}
	gconv.Scan(req, &input)

	var result, error = service.ExpressBase().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

func (c *cExpressBase) List(ctx context.Context, req *sys.ExpressBaseListReq) (res *sys.ExpressBaseListRes, err error) {
	input := do.ExpressBaseListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	input.BaseList = ml.BaseList{
		Sidx: dao.ExpressBase.Columns().ExpressOrder,
		Sort: ml.ORDER_BY_ASC,
	}
	var result, error = service.ExpressBase().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增菜单
func (c *cExpressBase) Add(ctx context.Context, req *sys.ExpressBaseAddReq) (res *sys.ExpressBaseEditRes, err error) {

	input := do.ExpressBase{}
	gconv.Scan(req, &input)

	var result, error = service.ExpressBase().Add(ctx, &input)
	//var result, error = service.ExpressBase().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &sys.ExpressBaseEditRes{
		ExpressId: result,
	}

	return
}

// Edit 编辑菜单
func (c *cExpressBase) Edit(ctx context.Context, req *sys.ExpressBaseEditReq) (res *sys.ExpressBaseEditRes, err error) {

	input := do.ExpressBase{}
	gconv.Scan(req, &input)

	var result, error = service.ExpressBase().Edit(ctx, &input)
	//var result, error = service.ExpressBase().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &sys.ExpressBaseEditRes{
		ExpressId: result,
	}

	return
}

// Remove 删除菜单
func (c *cExpressBase) Remove(ctx context.Context, req *sys.ExpressBaseRemoveReq) (res *sys.ExpressBaseRemoveRes, err error) {

	var _, error = service.ExpressBase().Remove(ctx, req.ExpressId)

	/*
		input := do.ExpressBase{}
		input.ExpressBaseTime = gtime.Now()
		input.ExpressBaseId = req.ExpressBaseId[0]
		input.ExpressBaseSort = 0

		var _, error = service.ExpressBase().Edit(ctx, &input)
	*/

	if error != nil {
		err = error
	}

	res = &sys.ExpressBaseRemoveRes{}

	return
}

func (c *cExpressBase) EditState(ctx context.Context, req *sys.ExpressBaseEditStateReq) (res *sys.ExpressBaseEditStateRes, err error) {

	input := do.ExpressBase{}
	gconv.Scan(req, &input)

	var result, error = service.ExpressBase().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.ExpressBaseEditStateRes{
		ExpressId: result,
	}

	return
}
