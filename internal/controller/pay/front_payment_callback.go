package pay

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/go-pay/xlog"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
	"golershop.cn/api/pay"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	PaymentCallback = cPaymentCallback{}
)

type cPaymentCallback struct{}

// WxNotify 微信支付回调
func (c *cPaymentCallback) WxNotify(ctx context.Context, req *pay.PaymentWechatNotifyReq) (res *pay.PaymentWechatNotifyRes, err error) {
	client, err := service.PaymentWechat().GetClient(ctx)
	now := gtime.Now()

	if err != nil {
		xlog.Error(err)
		return
	}

	notifyReq, err := wechat.V3ParseNotify(g.RequestFromCtx(ctx).Request)

	if err != nil {
		xlog.Error(err)
		return
	}

	// 获取微信平台证书
	certMap := client.WxPublicKeyMap()

	// 验证异步通知的签名
	err = notifyReq.VerifySignByPKMap(certMap)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 通用通知解密（推荐此方法）
	result, err := notifyReq.DecryptCipherText(string(client.ApiV3Key))

	//验证通过
	//交易号
	outTradeNo := result.OutTradeNo

	//处理订单
	//订单编号
	orderId, err := service.ConsumeDeposit().GetOrderId(ctx, outTradeNo)
	if err != nil {
		xlog.Error(err)
		return
	}

	deposit := do.ConsumeDeposit{}
	deposit.DepositNo = outTradeNo
	deposit.DepositTradeNo = result.TransactionId
	deposit.OrderId = orderId
	deposit.DepositSubject = result.Attach
	deposit.DepositQuantity = 1
	deposit.DepositNotifyTime = now
	deposit.DepositSellerId = result.Mchid
	//deposit.DepositIsTotalFeeAdjust = result.

	deposit.DepositTotalFee = decimal.NewFromInt(int64(result.Amount.Total)).Div(decimal.NewFromInt(100))
	deposit.DepositPrice = deposit.DepositTotalFee

	deposit.DepositBuyerId = result.Payer.Openid
	deposit.DepositTime = now.UnixMilli()
	deposit.DepositPaymentType = consts.PAYMENT_TYPE_ONLINE
	deposit.DepositService = result.TradeType
	//deposit.DepositSign = result.Si
	//deposit.DepositExtraParam(JSONUtil.toJsonStr(consumeDeposit));
	deposit.PaymentChannelId = consts.PAYMENT_CHANNEL_WECHAT
	deposit.DepositTradeStatus = result.TradeState

	//deposit.DepositState = 1

	_, err = service.ConsumeDeposit().ProcessDeposit(ctx, &deposit)

	if err != nil {
		return res, err
	}

	// ====↓↓↓====异步通知应答====↓↓↓====
	// 退款通知http应答码为200且返回状态码为SUCCESS才会当做商户接收成功，否则会重试。
	// 注意：重试过多会导致微信支付端积压过多通知而堵塞，影响其他正常通知。

	r := g.RequestFromCtx(ctx)
	r.Response.ClearBuffer()
	r.Response.WriteJsonExit(wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})

	return
}
