package pay

import (
	"context"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/xlog"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mileusna/useragent"
	"golershop.cn/api/pay"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao/global"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
	"golershop.cn/utility"
	"os"
)

// WxReturn 微信支付回调
func (c *cPaymentCallback) AlipayRerun(ctx context.Context, req *pay.PaymentAlipayReturnReq) (res *pay.PaymentAlipayReturnRes, err error) {
	/*
			client, err := service.PaymentAlipay().GetClient(ctx)
			now := gtime.Now()

			// 初始化 BodyMap
			bm := make(gopay.BodyMap)
			bm.Set("subject", "条码支付").
				Set("scene", "bar_code").
				Set("auth_code", "286248566432274952").
				Set("out_trade_no", "GZ201909081743431443").
				Set("total_amount", "0.01").
				Set("timeout_express", "2m")

			aliRsp, err := client.TradePay(ctx, bm)
			if err != nil {
				xlog.Error("err:", err)
				return
			}

			/*
				// 公钥模式验签
				//    注意：APP支付，手机网站支付，电脑网站支付 不支持同步返回验签
				//    aliPayPublicKey：支付宝平台获取的支付宝公钥
				//    signData：待验签参数，aliRsp.SignData
				//    sign：待验签sign，aliRsp.Sign
				ok, err := alipay.VerifySyncSign(aliPayPublicKey, aliRsp.SignData, aliRsp.Sign)



		// 支付宝证书路径
		path := utility.UploadPath(ctx)
		alipayPublicCertPath := service.ConfigBase().GetStr(ctx, "alipay_cert_path", "")
		alipayPublicCertPath = path + alipayPublicCertPath
		alipayPublicCert, err := os.ReadFile(alipayPublicCertPath)
		if err != nil {
			xlog.Error(err)
			return
		}
		// 公钥证书模式验签
		//    aliPayPublicKeyCert：支付宝公钥证书存放路径 alipayPublicCert.crt 或文件内容[]byte
		//    signData：待验签参数，aliRsp.SignData
		//    sign：待验签sign，aliRsp.Sign
		ok, err := alipay.VerifySyncSignWithCert(alipayPublicCert, aliRsp.SignData, aliRsp.Sign)

		if err != nil {
			xlog.Error(err)
			return
		}

		if !ok {

		}
	*/

	r := g.RequestFromCtx(ctx)
	outTradeNo := req.OutTradeNo

	//订单编号
	orderId, err := service.ConsumeDeposit().GetOrderId(ctx, outTradeNo)
	if err != nil {
		xlog.Error(err)
		return
	}

	var redirectUrl string
	if useragent.Parse(g.RequestFromCtx(ctx).UserAgent()).Mobile {
		redirectUrl = global.UrlH5 + "/member/order/detail?init_pay_flag=1&on=" + orderId
	} else {
		redirectUrl = global.UrlPc + "/user/order/detail?init_pay_flag=1&order_id=" + orderId
	}

	r.Response.ClearBuffer()
	r.Response.RedirectTo(redirectUrl)

	res = &pay.PaymentAlipayReturnRes{}
	return
}

// WxNotify 微信支付回调
func (c *cPaymentCallback) AlipayNotify(ctx context.Context, req *pay.PaymentAlipayNotifyReq) (res *pay.PaymentAlipayNotifyRes, err error) {
	now := gtime.Now()

	// 解析异步通知的参数
	// req：*http.Request
	notifyReq, err := alipay.ParseNotifyToBodyMap(g.RequestFromCtx(ctx).Request)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 支付宝异步通知验签（公钥证书模式）
	// 支付宝证书路径
	path := utility.UploadPath(ctx)
	alipayPublicCertPath := service.ConfigBase().GetStr(ctx, "alipay_cert_path", "")
	alipayPublicCertPath = path + alipayPublicCertPath
	alipayPublicCert, err := os.ReadFile(alipayPublicCertPath)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 支付宝异步通知验签（公钥模式）
	//ok, err = alipay.VerifySign(aliPayPublicKey, notifyReq)

	// 支付宝异步通知验签（公钥证书模式）
	ok, err := alipay.VerifySignWithCert(alipayPublicCert, notifyReq)
	if err != nil {
		xlog.Error(err)
		return
	}

	if !ok {
		r := g.RequestFromCtx(ctx)
		r.Response.ClearBuffer()
		r.Response.WriteExit("failure")
	}

	// 如果需要，可将 BodyMap 内数据，Unmarshal 到指定结构体指针 ptr
	//err = notifyReq.Unmarshal(ptr)

	//验证通过
	//交易号
	outTradeNo := notifyReq.GetString("out_trade_no")

	//处理订单
	//订单编号
	orderId, err := service.ConsumeDeposit().GetOrderId(ctx, outTradeNo)
	if err != nil {
		xlog.Error(err)
		return
	}

	deposit := do.ConsumeDeposit{}
	deposit.DepositNo = outTradeNo
	deposit.DepositTradeNo = notifyReq.Get("trade_no")
	deposit.OrderId = orderId
	deposit.DepositSubject = notifyReq.Get("subject")
	deposit.DepositQuantity = notifyReq.Get("quantity")
	deposit.DepositNotifyTime = now
	deposit.DepositSellerId = notifyReq.Get("seller_id")
	deposit.DepositIsTotalFeeAdjust = notifyReq.Get("is_total_fee_adjust")

	deposit.DepositTotalFee = notifyReq.GetInterface("total_amount")
	deposit.DepositPrice = deposit.DepositTotalFee

	deposit.DepositBuyerId = notifyReq.Get("buyer_id")
	deposit.DepositTime = now.UnixMilli()
	deposit.DepositPaymentType = consts.PAYMENT_TYPE_ONLINE
	deposit.DepositService = notifyReq.Get("exterface")
	//deposit.DepositSign = result.Si
	//deposit.DepositExtraParam(JSONUtil.toJsonStr(consumeDeposit));
	deposit.PaymentChannelId = consts.PAYMENT_CHANNEL_WECHAT
	deposit.DepositTradeStatus = notifyReq.Get("trade_status")

	//deposit.DepositState = 1

	_, err = service.ConsumeDeposit().ProcessDeposit(ctx, &deposit)

	if err != nil {
		return res, err
	}

	// ====异步通知，返回支付宝平台的信息====
	// 文档：https://opendocs.alipay.com/open/203/105286
	// 程序执行完后必须打印输出“success”（不包含引号）。如果商户反馈给支付宝的字符不是success这7个字符，支付宝服务器会不断重发通知，直到超过24小时22分钟。一般情况下，25小时以内完成8次通知（通知的间隔频率一般是：4m,10m,10m,1h,2h,6h,15h）

	r := g.RequestFromCtx(ctx)
	r.Response.ClearBuffer()
	r.Response.WriteExit("success")

	return

}
