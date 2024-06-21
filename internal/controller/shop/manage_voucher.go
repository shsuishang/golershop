package shop

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/shop"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	Voucher = cVoucher{}
)

type cVoucher struct{}

// List 用户优惠券表-分页列表查询
func (c *cVoucher) List(ctx context.Context, req *shop.UserVoucherReq) (res *shop.UserVoucherRes, err error) {
	input := do.UserVoucherListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.BaseList.WhereExt)

	result, err := service.UserVoucher().List(ctx, &input)
	if err != nil {
		return nil, err
	}

	gconv.Scan(result, &res)
	return
}
