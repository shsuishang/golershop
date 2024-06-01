package admin

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/admin"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	UserRole = cUserRole{}
)

type cUserRole struct{}

// =================== 管理端使用 =========================

func (c *cUserRole) List(ctx context.Context, req *admin.UserRoleListReq) (res *admin.UserRoleListRes, err error) {
	input := do.UserRoleListInput{}
	gconv.Scan(req, &input)

	if !g.IsEmpty(req.UserRoleName) {
		var likes = []*ml.WhereExt{{
			Column: dao.UserRole.Columns().UserRoleName,
			Val:    "%" + req.UserRoleName + "%",
			Symbol: ml.LIKE,
		}}

		input.WhereExt = likes
	}

	var result, error = service.UserRole().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增角色
func (c *cUserRole) Add(ctx context.Context, req *admin.UserRoleAddReq) (res *admin.UserRoleEditRes, err error) {
	input := do.UserRole{}
	gconv.Scan(req, &input)
	input.UserRoleCtime = gtime.Now()

	var result, error = service.UserRole().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &admin.UserRoleEditRes{
		UserRoleId: uint(result),
	}

	return
}

// Edit 编辑角色
func (c *cUserRole) Edit(ctx context.Context, req *admin.UserRoleEditReq) (res *admin.UserRoleEditRes, err error) {

	input := do.UserRole{}
	gconv.Scan(req, &input)
	input.UserRoleUtime = gtime.Now()

	var result, error = service.UserRole().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &admin.UserRoleEditRes{
		UserRoleId: uint(result),
	}

	return
}

// Remove 删除角色
func (c *cUserRole) Remove(ctx context.Context, req *admin.UserRoleRemoveReq) (res *admin.UserRoleRemoveRes, err error) {
	var _, error = service.UserRole().Remove(ctx, req.UserRoleId)

	if error != nil {
		err = error
	}

	res = &admin.UserRoleRemoveRes{}

	return
}
