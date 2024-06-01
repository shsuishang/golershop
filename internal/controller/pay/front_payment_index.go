package pay

import (
	"context"
	"errors"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/go-pay/xlog"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
	"golershop.cn/api/pay"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/dao/global"
	"golershop.cn/internal/model"
	"golershop.cn/internal/service"
	"golershop.cn/utility"
	"time"
)

var (
	PaymentIndex = cPaymentIndex{}
)

type cPaymentIndex struct{}

// PaymentReq 余额支付
func (c *cPaymentIndex) MoneyPay(ctx context.Context, req *pay.PaymentReq) (res *pay.PaymentRes, err error) {
	// 获取登录用户ID
	userId := service.BizCtx().GetUserId(ctx)

	// 设置支付渠道和支付方式
	req.PaymentChannelId = consts.PAYMENT_MET_MONEY
	req.DepositPaymentType = consts.PAYMENT_TYPE_ONLINE

	// 处理订单支付结果
	payInfo := &model.PayMetVo{
		PaymentMetId:     consts.PAYMENT_MET_MONEY,
		PaymentChannelId: req.PaymentChannelId,
		PaymentTypeId:    req.DepositPaymentType,
		PmMoney:          req.PmMoney,
	}

	// 初始化支付结果
	res = &pay.PaymentRes{
		OrderId:    gstr.Join(req.OrderId, ","),
		StatusCode: 200,
		Paid:       false,
	}

	// 判断余额是否足够
	userResourceBuy, err := dao.UserResource.Get(ctx, userId)
	if userResourceBuy.UserMoney >= payInfo.PmMoney {
		// 检查支付密码
		_, err = c.checkPayPasswd(ctx, userId, req.Password)

		if err != nil {
			return nil, err
		}

		// 处理支付
		processPayOutput, err := service.ConsumeTrade().ProcessPay(ctx, gstr.Join(req.OrderId, ","), *payInfo)

		// 判断支付状态
		if err == nil && processPayOutput.Paid {
			res.Paid = processPayOutput.Paid
			res.StatusCode = 200
		} else {
			res.StatusCode = 250
		}

	} else {
		res.StatusCode = 250
	}

	return
}

/**
 * 验证支付密码
 *
 * @param userId
 * @param password
 * @return
 */
func (c *cPaymentIndex) checkPayPasswd(ctx context.Context, userId uint, password string) (bool, error) {
	// 获取用户支付信息
	userPay, _ := service.UserPay().Get(ctx, userId)

	// 如果用户支付信息不为空且支付密码不为空
	if userPay != nil && userPay.UserPayPasswd != "" {
		userPaySalt := userPay.UserPaySalt
		saltPassword := gmd5.MustEncryptString(userPaySalt + gmd5.MustEncryptString(password))

		// 如果密码不匹配，抛出异常
		if saltPassword != userPay.UserPayPasswd {
			return false, errors.New("支付密码错误！")
		}
	} else {
		return false, errors.New("未设置支付密码！")
	}

	return true, nil
}

// WechatAppletPay 微信小程序支付
func (c *cPaymentIndex) WechatAppletPay(ctx context.Context, req *pay.WechatAppletPayReq) (res *pay.WechatAppletPayRes, err error) {
	req.PaymentChannelId = consts.PAYMENT_CHANNEL_WECHAT

	input := &model.PaymentInput{}
	gconv.Struct(req, input)
	payInfo, err := service.ConsumeDeposit().GetPayResult(ctx, input)
	if err != nil {
		xlog.Error(err)
		return
	}

	//获得交易号
	tradeNo := payInfo.TradeNo

	appId := service.ConfigBase().GetStr(ctx, "wechat_app_id", "")
	client, err := service.PaymentWechat().GetClient(ctx)

	if err != nil {
		xlog.Error(err)
		return
	}

	//JSAPI下单 示例
	xlog.Debug("tradeNo:", tradeNo)
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	bm := make(gopay.BodyMap)
	bm.Set("appid", appId).
		Set("description", "在线购物").
		Set("out_trade_no", tradeNo).
		Set("time_expire", expire).
		Set("notify_url", global.BaseUrl+("/front/pay/callback/wechatNotify")).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", payInfo.Amount).
				Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) {
			bm.Set("openid", req.Openid)
		})

	wxRsp, err := client.V3TransactionJsapi(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == wechat.Success {
		xlog.Debugf("wxRsp: %+v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)

	//下单后，获取微信小程序支付、APP支付、JSAPI支付所需要的 pay sign
	// 小程序
	applet, err := client.PaySignOfApplet(appId, wxRsp.Response.PrepayId)
	gconv.Scan(applet, res.Data)

	return
}

// WechatAppletPay 微信App支付
func (c *cPaymentIndex) WechatAppPay(ctx context.Context, req *pay.WechatAppPayReq) (res *pay.WechatAppPayRes, err error) {
	req.PaymentChannelId = consts.PAYMENT_CHANNEL_WECHAT

	input := &model.PaymentInput{}
	gconv.Struct(req, input)
	payInfo, err := service.ConsumeDeposit().GetPayResult(ctx, input)
	if err != nil {
		xlog.Error(err)
		return
	}

	//获得交易号
	tradeNo := payInfo.TradeNo

	appId := service.ConfigBase().GetStr(ctx, "weixin_app_id", "")
	client, err := service.PaymentWechat().GetClient(ctx)

	if err != nil {
		xlog.Error(err)
		return
	}

	//JSAPI下单 示例
	xlog.Debug("tradeNo:", tradeNo)
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	bm := make(gopay.BodyMap)
	bm.Set("appid", appId).
		Set("description", "在线购物").
		Set("out_trade_no", tradeNo).
		Set("time_expire", expire).
		Set("notify_url", global.BaseUrl+("/front/pay/callback/wechatNotify")).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", payInfo.Amount).
				Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) {
			bm.Set("openid", req.Openid)
		})

	wxRsp, err := client.V3TransactionApp(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == wechat.Success {
		xlog.Debugf("wxRsp: %+v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)

	// App
	applet, err := client.PaySignOfApp(appId, wxRsp.Response.PrepayId)
	gconv.Scan(applet, res.Data)

	return

}

// WechatJSAPIPay 微信JSAPI支付
func (c *cPaymentIndex) WechatJSAPIPay(ctx context.Context, req *pay.WechatJSAPIPayReq) (res *pay.WechatJSAPIPayRes, err error) {
	req.PaymentChannelId = consts.PAYMENT_CHANNEL_WECHAT

	input := &model.PaymentInput{}
	gconv.Struct(req, input)
	payInfo, err := service.ConsumeDeposit().GetPayResult(ctx, input)
	if err != nil {
		xlog.Error(err)
		return
	}

	//获得交易号
	tradeNo := payInfo.TradeNo

	appId := service.ConfigBase().GetStr(ctx, "wechat_app_id", "")
	client, err := service.PaymentWechat().GetClient(ctx)

	if err != nil {
		xlog.Error(err)
		return
	}

	//JSAPI下单 示例
	xlog.Debug("tradeNo:", tradeNo)
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	bm := make(gopay.BodyMap)
	bm.Set("appid", appId).
		Set("description", "在线购物").
		Set("out_trade_no", tradeNo).
		Set("time_expire", expire).
		Set("notify_url", global.BaseUrl+("/front/pay/callback/wechatNotify")).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", payInfo.Amount).
				Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) {
			bm.Set("openid", req.Openid)
		})

	wxRsp, err := client.V3TransactionJsapi(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if wxRsp.Code == wechat.Success {
		xlog.Debugf("wxRsp: %+v", wxRsp.Response)
		return
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)

	//下单后，获取微信小程序支付、APP支付、JSAPI支付所需要的 pay sign
	// JSAPI
	applet, err := client.PaySignOfJSAPI(appId, wxRsp.Response.PrepayId)
	gconv.Scan(applet, res.Data)

	return
}

// WechatH5Pay 微信H5支付
func (c *cPaymentIndex) WechatH5Pay(ctx context.Context, req *pay.WechatH5PayReq) (res *pay.WechatH5PayRes, err error) {
	req.PaymentChannelId = consts.PAYMENT_CHANNEL_WECHAT

	input := &model.PaymentInput{}
	gconv.Struct(req, input)
	payInfo, err := service.ConsumeDeposit().GetPayResult(ctx, input)
	if err != nil {
		xlog.Error(err)
		return
	}

	//获得交易号
	tradeNo := payInfo.TradeNo

	appId := service.ConfigBase().GetStr(ctx, "wechat_app_id", "")
	client, err := service.PaymentWechat().GetClient(ctx)

	if err != nil {
		xlog.Error(err)
		return
	}

	//JSAPI下单 示例
	xlog.Debug("tradeNo:", tradeNo)
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	bm := make(gopay.BodyMap)
	bm.Set("appid", appId).
		Set("description", "在线购物").
		Set("out_trade_no", tradeNo).
		Set("time_expire", expire).
		Set("notify_url", global.BaseUrl+("/front/pay/callback/wechatNotify")).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", payInfo.Amount).
				Set("currency", "CNY")
		}).
		SetBodyMap("scene_info", func(b gopay.BodyMap) {
			b.Set("payer_client_ip", utility.GetClientIp(g.RequestFromCtx(ctx))).
				SetBodyMap("h5_info", func(b gopay.BodyMap) {
					b.Set("type", "Wap").
						Set("app_name", "随商商城").
						Set("app_url", global.BaseUrl) //.Set("bundle_id", "com.tencent.wzryiOS")
				})
		})

	wxRsp, err := client.V3TransactionH5(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}

	res = &pay.WechatH5PayRes{}
	res.OrderId = gstr.JoinAny(req.OrderId, ",")

	if wxRsp.Code == wechat.Success {
		xlog.Debugf("wxRsp: %+v", wxRsp.Response)
		res.MwebUrl = wxRsp.Response.H5Url

		res.StatusCode = 200
		return
	}

	res.StatusCode = 250
	xlog.Errorf("wxRsp:%s", wxRsp.Error)

	return
}

// AlipayH5Pay 支付宝H5支付
func (c *cPaymentIndex) AlipayH5Pay(ctx context.Context, req *pay.AlipayH5PayReq) (res *pay.AlipayH5PayRes, err error) {
	req.PaymentChannelId = consts.PAYMENT_CHANNEL_ALIPAY

	input := &model.PaymentInput{}
	gconv.Struct(req, input)
	payInfo, err := service.ConsumeDeposit().GetPayResult(ctx, input)
	if err != nil {
		xlog.Error(err)
		return
	}

	//获得交易号
	tradeNo := payInfo.TradeNo

	client, err := service.PaymentAlipay().GetClient(ctx, tradeNo)

	if err != nil {
		xlog.Error(err)
		return
	}

	/*
		//配置公共参数
		client.SetCharset("utf-8").
			SetSignType(alipay.RSA2).
			//SetReturnUrl("https://www.fmm.ink").
			SetNotifyUrl("https://www.fmm.ink")
	*/

	//请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "在线购物")
	bm.Set("out_trade_no", tradeNo)
	bm.Set("quit_url", client.ReturnUrl)
	bm.Set("total_amount", decimal.NewFromFloat(payInfo.Amount).Div(decimal.NewFromInt(100)))
	bm.Set("product_code", "QUICK_WAP_WAY")

	res = &pay.AlipayH5PayRes{}
	res.OrderId = gstr.JoinAny(req.OrderId, ",")

	//手机网站支付请求
	payUrl, err := client.TradeWapPay(ctx, bm)
	if err != nil {
		xlog.Error("err:", err)

		res.StatusCode = 250
		return
	} else {

		res.MwebUrl = payUrl
		res.StatusCode = 200
	}

	xlog.Debug("payUrl:", payUrl)

	return
}
