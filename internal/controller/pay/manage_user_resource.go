package pay

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pay"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	UserResource = cUserResource{}
)

type cUserResource struct{}

// =================== 管理端使用 =========================
func (c *cUserResource) List(ctx context.Context, req *pay.UserResourceListReq) (res *pay.UserResourceListRes, err error) {
	input := do.UserResourceListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.BaseList.WhereExt)

	var result, error = service.UserResource().GetList(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增
func (c *cUserResource) Add(ctx context.Context, req *pay.UserResourceAddReq) (res *pay.UserResourceEditRes, err error) {
	input := do.UserResource{}
	gconv.Scan(req, &input)

	var result, error = service.UserResource().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &pay.UserResourceEditRes{
		UserId: uint(result),
	}

	return
}

// Edit 编辑
func (c *cUserResource) Edit(ctx context.Context, req *pay.UserResourceEditReq) (res *pay.UserResourceEditRes, err error) {

	input := do.UserResource{}
	gconv.Scan(req, &input)

	var result, error = service.UserResource().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &pay.UserResourceEditRes{
		UserId: uint(result),
	}

	return
}

// Remove 删除
func (c *cUserResource) Remove(ctx context.Context, req *pay.UserResourceRemoveReq) (res *pay.UserResourceRemoveRes, err error) {
	var _, error = service.UserResource().Remove(ctx, req.UserId)

	if error != nil {
		err = error
	}

	res = &pay.UserResourceRemoveRes{}

	return
}

// UpdateUserMoney 修改资金
func (c *cUserResource) UpdateUserMoney(ctx context.Context, req *pay.UpdateUserMoneyReq) (res *pay.UpdateUserMoneyRes, err error) {
	// 将请求参数复制到 MoneyVo 对象
	input := &model.MoneyVo{}
	gconv.Scan(req, input)

	// 调用服务更新用户资金
	success, err := service.UserResource().UpdateUserMoney(ctx, input)
	if err != nil {
		return nil, err
	}

	// 返回结果
	if success {
		return res, err
	}

	return nil, err
}

// UpdatePoints 修改积分
func (c *cUserResource) UpdatePoints(ctx context.Context, req *pay.UpdatePointsReq) (res *pay.UpdatePointsRes, err error) {
	// 将请求参数复制到 MoneyVo 对象
	input := &model.UserPointsVo{}
	gconv.Scan(req, input)

	// 调用服务更新用户资金
	success, err := service.UserResource().UpdatePoints(ctx, input)
	if err != nil {
		return nil, err
	}

	// 返回结果
	if success {
		return res, err
	}

	return nil, err
}
