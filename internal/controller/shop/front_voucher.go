package shop

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/shop"
	"golershop.cn/internal/service"
)

var (
	UserVoucher = cUserVoucher{}
)

type cUserVoucher struct{}

// List 会员优惠券列表
func (c *cUserVoucher) List(ctx context.Context, req *shop.UserVoucherListReq) (res *shop.UserVoucherListRes, err error) {
	// 获取当前登录用户ID
	userId := service.BizCtx().GetUserId(ctx)
	req.UserId = userId

	var result, error = service.UserVoucher().GetLists(ctx, req)

	if error != nil {
		return nil, err
	}

	res = &shop.UserVoucherListRes{}
	gconv.Scan(result, &res)

	return res, nil
}

// GetEachVoucherNum 列举出不同优惠券的数量
func (c *cUserVoucher) GetEachVoucherNum(ctx context.Context, req *shop.GetVoucherNumReq) (res *shop.GetVoucherNumRes, err error) {
	// 获取当前登录用户ID
	user := service.BizCtx().GetUser(ctx)

	userId := user.UserId

	// 调用服务方法获取优惠券数量
	voucherCountRes, err := service.UserVoucher().GetEachVoucherNum(ctx, req.VoucherStateId, userId)
	if err != nil {
		return nil, err
	}

	// 转换结果
	res = &shop.GetVoucherNumRes{}
	gconv.Scan(voucherCountRes, res)

	return
}
