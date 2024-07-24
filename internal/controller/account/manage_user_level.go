package account

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/account"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	UserLevel = cUserLevel{}
)

type cUserLevel struct{}

// =================== 管理端使用 =========================

func (c *cUserLevel) List(ctx context.Context, req *account.UserLevelListReq) (res *account.UserLevelListRes, err error) {
	input := do.UserLevelListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.UserLevel().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增用户等级
func (c *cUserLevel) Add(ctx context.Context, req *account.UserLevelAddReq) (res *account.UserLevelEditRes, err error) {

	input := do.UserLevel{}
	gconv.Scan(req, &input)
	input.UserLevelTime = gtime.Now().TimestampMilli()

	var result, error = service.UserLevel().Add(ctx, &input)
	//var result, error = service.UserLevel().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &account.UserLevelEditRes{
		UserLevelId: result,
	}

	return
}

// Edit 编辑用户等级
func (c *cUserLevel) Edit(ctx context.Context, req *account.UserLevelEditReq) (res *account.UserLevelEditRes, err error) {

	input := do.UserLevel{}
	gconv.Scan(req, &input)
	input.UserLevelTime = gtime.Now().TimestampMilli()

	var result, error = service.UserLevel().Edit(ctx, &input)
	//var result, error = service.UserLevel().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &account.UserLevelEditRes{
		UserLevelId: result,
	}

	return
}

// Remove 删除用户等级
func (c *cUserLevel) Remove(ctx context.Context, req *account.UserLevelRemoveReq) (res *account.UserLevelRemoveRes, err error) {

	var _, error = service.UserLevel().Remove(ctx, req.UserLevelId)

	/*
		input := do.UserLevel{}
		input.UserLevelTime = gtime.Now()
		input.UserLevelId = req.UserLevelId[0]
		input.UserLevelSort = 0

		var _, error = service.UserLevel().Edit(ctx, &input)
	*/

	if error != nil {
		err = error
	}

	res = &account.UserLevelRemoveRes{}

	return
}
