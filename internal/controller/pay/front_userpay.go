package pay

import (
	"context"
	"errors"
	"golershop.cn/api/pay"
	"golershop.cn/internal/service"
)

var (
	UserPay = cUserPay{}
)

type cUserPay struct{}

func (c *cUserPay) GetPayPasswd(ctx context.Context, req *pay.PayPasswdReq) (res *pay.PayPasswdRes, err error) {
	// 获取登录用户ID
	userId := service.BizCtx().GetUserId(ctx)

	result, err := service.UserPay().GetPayPasswd(ctx, userId)

	if result == nil {
		return nil, errors.New("未设置支付密码！")
	}

	if result.UserPayPasswd == "" {
		return nil, errors.New("未设置支付密码！")
	}

	return
}

func (c *cUserPay) ChangePayPasswd(ctx context.Context, req *pay.ChangeUserPayReq) (res *pay.ChangeUserPayRes, err error) {
	// 获取登录用户ID
	userId := service.BizCtx().GetUserId(ctx)

	_, err = service.UserPay().ChangePayPassword(ctx, req.OlePayPassword, req.NewPayPassword, req.PayPasssord, userId)

	if err != nil {
		return nil, errors.New("设置支付密码失败！")
	}

	return
}
