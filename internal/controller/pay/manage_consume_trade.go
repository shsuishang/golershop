package pay

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
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

	var result, error = service.ConsumeTrade().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}
