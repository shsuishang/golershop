package admin

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/admin"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	UserAdmin = cUserAdmin{}
)

type cUserAdmin struct{}

// =================== 管理端使用 =========================

func (c *cUserAdmin) List(ctx context.Context, req *admin.UserAdminListReq) (res *admin.UserAdminListRes, err error) {
	input := do.UserAdminListInput{}
	gconv.Scan(req, &input)
	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.UserAdmin().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增角色
func (c *cUserAdmin) Add(ctx context.Context, req *admin.UserAdminAddReq) (res *admin.UserAdminEditRes, err error) {
	input := do.UserAdmin{}
	gconv.Scan(req, &input)
	input.UserAdminCtime = gtime.Now()

	var result, error = service.UserAdmin().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &admin.UserAdminEditRes{
		UserId: uint(result),
	}

	return
}

// Edit 编辑角色
func (c *cUserAdmin) Edit(ctx context.Context, req *admin.UserAdminEditReq) (res *admin.UserAdminEditRes, err error) {

	input := do.UserAdmin{}
	gconv.Scan(req, &input)
	input.UserAdminUtime = gtime.Now()

	var result, error = service.UserAdmin().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &admin.UserAdminEditRes{
		UserId: uint(result),
	}

	return
}

// Remove 删除角色
func (c *cUserAdmin) Remove(ctx context.Context, req *admin.UserAdminRemoveReq) (res *admin.UserAdminRemoveRes, err error) {
	var _, error = service.UserAdmin().Remove(ctx, req.UserId)

	if error != nil {
		err = error
	}

	res = &admin.UserAdminRemoveRes{}

	return
}
