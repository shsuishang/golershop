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
	ConsumeTrade = cConsumeTrade{}
)

type cConsumeTrade struct{}

// =================== 管理端使用 =========================
func (c *cConsumeTrade) List(ctx context.Context, req *pay.ConsumeTradeListReq) (res *pay.ConsumeTradeListRes, err error) {
	input := do.ConsumeTradeListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.BaseList.WhereExt)

	result, err := service.ConsumeTrade().List(ctx, &input)

	if err != nil {
		err = err
	}

	gconv.Scan(result, &res)

	return
}
