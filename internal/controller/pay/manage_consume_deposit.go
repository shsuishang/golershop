package pay

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pay"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ConsumeDeposit = cConsumeDeposit{}
)

type cConsumeDeposit struct{}

// =================== 管理端使用 =========================
func (c *cConsumeDeposit) List(ctx context.Context, req *pay.ConsumeDepositListReq) (res *pay.ConsumeDepositListRes, err error) {
	input := do.ConsumeDepositListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.ConsumeDeposit().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// OfflinePay 线下支付
func (c *cConsumeDeposit) OfflinePay(ctx context.Context, req *pay.ConsumeDepositOfflinePayReq) (res *pay.ConsumeDepositOfflinePayRes, err error) {
	input := do.ConsumeDeposit{}
	gconv.Scan(req, &input)

	//交易号 == 流水号
	input.DepositNo = req.DepositTradeNo

	deposits, err := dao.ConsumeDeposit.Find(ctx, &do.ConsumeDepositListInput{Where: do.ConsumeDeposit{DepositTradeNo: input.DepositTradeNo}})

	if err != nil {
		return nil, err
	}

	if len(deposits) > 0 {
		return nil, errors.New("支付凭证号已经存在！")
	}

	_, err = service.ConsumeDeposit().OfflinePay(ctx, &input)

	res = &pay.ConsumeDepositOfflinePayRes{
		OrderId: req.OrderId,
	}

	return
}

func (c *cConsumeDeposit) EditReview(ctx context.Context, req *pay.ConsumeDepositEditReviewReq) (res *pay.ConsumeDepositEditReviewRes, err error) {
	input := do.ConsumeDeposit{}
	gconv.Scan(req, &input)

	_, err = service.ConsumeDeposit().Edit(ctx, &input)

	if err != nil {
		err = err
	}

	res = &pay.ConsumeDepositEditReviewRes{
		ConsumeDepositId: req.ConsumeDepositId,
	}

	return
}
