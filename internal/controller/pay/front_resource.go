package pay

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/pay"
	"golershop.cn/internal/service"
)

var (
	Resource = cResource{}
)

type cResource struct{}

// SignInfo 签到信息
func (c *cResource) SignInfo(ctx context.Context, req *pay.SignInfoReq) (res *pay.SignInfoRes, err error) {
	// 获取登录用户ID
	userId := service.BizCtx().GetUserId(ctx)

	info, err := service.UserResource().GetSignInfo(ctx, userId)

	res = &pay.SignInfoRes{}
	gconv.Struct(info, res)
	return
}

// SignIn 签到
func (c *cResource) SignIn(ctx context.Context, req *pay.SignInReq) (res *pay.SignInRes, err error) {
	// 获取登录用户ID
	userId := service.BizCtx().GetUserId(ctx)

	_, err = service.UserResource().SignIn(ctx, userId)

	return
}

// SignState 今天签到状态
func (c *cResource) SignState(ctx context.Context, req *pay.SignStateReq) (res *pay.SignStateRes, err error) {
	// 获取登录用户ID
	userId := service.BizCtx().GetUserId(ctx)

	flag, err := service.UserResource().GetSignState(ctx, userId)

	if flag {
		return nil, nil
	} else {
		return nil, errors.New("尚未签到")
	}
}

// GetCommissionInfo 账户余额
func (c *cResource) GetCommissionInfo(ctx context.Context, req *pay.DistributionCommissionReq) (res *pay.DistributionCommissionRes, err error) {

	return res, err
}
