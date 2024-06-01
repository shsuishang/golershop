package pay

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/pay"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ConsumeRecord = cConsumeRecord{}
)

type cConsumeRecord struct{}

// =================== 管理端使用 =========================
func (c *cConsumeRecord) List(ctx context.Context, req *pay.ConsumeRecordListReq) (res *pay.ConsumeRecordListRes, err error) {
	input := do.ConsumeRecordListInput{}
	gconv.Scan(req, &input)

	var result, error = service.ConsumeRecord().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}
