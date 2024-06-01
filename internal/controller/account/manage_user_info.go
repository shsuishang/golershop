package account

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/account"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	UserInfo = cUserInfo{}
)

type cUserInfo struct{}

// =================== 管理端使用 =========================

func (c *cUserInfo) List(ctx context.Context, req *account.UserInfoListReq) (res *account.UserInfoListRes, err error) {
	input := do.UserInfoListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.UserInfo().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增菜单
func (c *cUserInfo) Add(ctx context.Context, req *account.UserInfoAddReq) (res *account.UserInfoEditRes, err error) {

	input := &model.RegisterInput{}
	gconv.Struct(req, input)

	input.BindType = consts.ACCOUNT

	if g.IsEmpty(req.Password) {
		input.Password = "Shopsuite@2018" + uuid.New().String()
	}

	// 系统登录
	userId, err := service.Login().DoRegister(ctx, input)

	if err != nil {
		return res, err
	}

	res = &account.UserInfoEditRes{
		UserId: userId,
	}

	return
}

// Edit 编辑菜单
func (c *cUserInfo) Edit(ctx context.Context, req *account.UserInfoEditReq) (res *account.UserInfoEditRes, err error) {

	input := do.UserInfo{}
	gconv.Scan(req, &input)

	var result, error = service.UserInfo().Edit(ctx, &input)
	//var result, error = service.UserInfo().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &account.UserInfoEditRes{
		UserId: result,
	}

	return
}

// Remove 删除菜单
func (c *cUserInfo) Remove(ctx context.Context, req *account.UserInfoRemoveReq) (res *account.UserInfoRemoveRes, err error) {

	var _, error = service.UserInfo().Remove(ctx, req.UserId)

	/*
		input := do.UserInfo{}
		input.UserInfoTime = gtime.Now()
		input.UserId = req.UserId[0]
		input.UserInfoSort = 0

		var _, error = service.UserInfo().Edit(ctx, &input)
	*/

	if error != nil {
		err = error
	}

	res = &account.UserInfoRemoveRes{}

	return
}
