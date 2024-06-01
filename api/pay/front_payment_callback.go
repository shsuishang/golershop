package pay

import (
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/gogf/gf/v2/frame/g"
)

// start fo front
type PaymentWechatNotifyReq struct {
	g.Meta `path:"/front/pay/callback/wechatNotify" tags:"支付回调" method:"post" summary:"微信支付回调接口"`
}

type PaymentWechatNotifyRes struct {
	wechat.V3NotifyRsp
}

// start fo aliay
type PaymentAlipayNotifyReq struct {
	g.Meta `path:"/front/pay/callback/alipayNotify" tags:"支付回调" method:"post" summary:"支付宝支付回调接口"`
}

type PaymentAlipayNotifyRes struct {
}

type PaymentAlipayReturnReq struct {
	g.Meta `path:"/front/pay/callback/returnUrl" tags:"支付回调" method:"get" summary:"支付宝支付返回接口"`

	OutTradeNo string `json:"out_trade_no"` // 交易编号
}

type PaymentAlipayReturnRes struct {
}
