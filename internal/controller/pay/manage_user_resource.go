package pay

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pay"
	"golershop.cn/internal/dao"
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

	if !g.IsEmpty(req.UserId) {
		var likes = []*ml.WhereExt{{
			Column: dao.UserResource.Columns().UserId,
			Val:    req.UserId,
			Symbol: ml.LIKE,
		}}

		input.WhereExt = likes
	}

	var result, error = service.UserResource().List(ctx, &input)

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
