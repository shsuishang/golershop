package pay

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/pay"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	Record = cRecord{}
)

type cRecord struct{}

// List 交易明细表
func (c *cRecord) List(ctx context.Context, req *pay.RecordListReq) (res *pay.RecordListRes, err error) {
	input := &do.ConsumeRecordListInput{}
	gconv.Struct(req, input)

	// 获取当前登录用户ID
	userId := service.BizCtx().GetUser(ctx).UserId
	input.Where.UserId = userId

	if req.ChangeType == 1 {
		input.Where.TradeTypeId = []uint{
			consts.TRADE_TYPE_SHOPPING,
			consts.TRADE_TYPE_TRANSFER,
			consts.TRADE_TYPE_WITHDRAW,
			consts.TRADE_TYPE_REFUND_PAY,
			consts.TRADE_TYPE_COMMISSION_TRANSFER,
		}
	}
	if req.ChangeType == 2 {
		input.Where.TradeTypeId = []uint{
			consts.TRADE_TYPE_DEPOSIT,
			consts.TRADE_TYPE_SALES,
			consts.TRADE_TYPE_COMMISSION,
			consts.TRADE_TYPE_REFUND_GATHERING,
			consts.TRADE_TYPE_TRANSFER_GATHERING,
		}
	}

	// 查询交易明细列表
	result, err := service.ConsumeRecord().List(ctx, input)
	if err != nil {
		return nil, err
	}

	res = &pay.RecordListRes{
		Items: result,
	}

	gconv.Scan(result, &res)

	return res, nil
}
