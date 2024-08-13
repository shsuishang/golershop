package pay

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pay"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ConsumeWithdraw = cConsumeWithdraw{}
)

type cConsumeWithdraw struct{}

func (c *cConsumeWithdraw) List(ctx context.Context, req *pay.ConsumeWithdrawListReq) (res *pay.ConsumeWithdrawListRes, err error) {
	input := do.ConsumeWithdrawListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.BaseList.WhereExt)

	result, err := service.ConsumeWithdraw().List(ctx, &input)

	if err != nil {
		err = err
	}

	gconv.Scan(result, &res)

	return
}
