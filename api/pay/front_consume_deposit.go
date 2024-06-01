package pay

import (
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/gogf/gf/v2/frame/g"
)

type ConsumeDepositOnlinePayReq struct {
	g.Meta `path:"/front/pay/consumeDeposit/online" tags:"支付" method:"post" summary:"在线支付接口"`

	OrderId          []string `json:"order_id" v:"required#请输入订单编号"  dc:"订单编号"` //订单编号
	PaymentChannelId uint     `json:"payment_channel_id"  d:"1403"`             //支付渠道

	DepositPaymentType uint `json:"deposit_payment_type"  d:"1302"` //支付方式(ENUM):1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付;

	PmMoney        float64 `json:"pm_money" `         //余额
	PmRechargeCard float64 `json:"pm_recharge_card" ` //充值卡
	PmPoints       float64 `json:"pm_points" `        //积分
	PmCredit       float64 `json:"pm_credit" `        //信用账户
	PmRedpack      float64 `json:"pm_redpack" `       //红包账户
}

type ConsumeDepositOnlinePayRes struct {
}

type PaymentVo struct {
	OrderId            []string `json:"order_id" v:"required#请输入订单编号"  dc:"订单编号"`                                 //订单编号
	PaymentChannelId   uint     `json:"payment_channel_id" d:"26"             `                                   // 支付渠道
	Openid             string   `json:"openid"       d:"ouUkf1ZjSzgdPa8zyYjpBncfiYEY"                           ` // openid
	DepositPaymentType uint     `json:"deposit_payment_type" d:"1302"         `                                   // 支付方式(ENUM):1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付;
	Password           string   `json:"password"                                `                                 // 支付密码
	PmMoney            float64  `json:"pm_money"                                `                                 // 余额支付
	PmRechargeCard     float64  `json:"pm_recharge_card"                       `                                  // 充值卡支付
	PmPoints           float64  `json:"pm_points"                              `                                  // 积分支付
	PmCredit           float64  `json:"pm_credit"                              `                                  // 积分支付
}
type PaymentReq struct {
	g.Meta `path:"/front/pay/consumeDeposit/moneyPay" tags:"支付" method:"post" summary:"支付接口"`
	PaymentVo
}

type MoneyPayRes struct {
	OrderId    string `json:"order_id"`                  // 订单编号(DOT)
	StatusCode int    `json:"status_code" default:"200"` // 状态码
	Paid       bool   `json:"paid" default:"false"`      // 订单已支付完成
	Code       int    `json:"code" default:"0"`          // 101：需要支付密码
}

type PaymentRes MoneyPayRes

type WechatAppletPayReq struct {
	g.Meta `path:"/front/pay/consumeDeposit/wechatAppletPay" tags:"支付" method:"post" summary:"微信小程序支付接口"`

	PaymentVo
}

type WechatAppletPayRes struct {
	MoneyPayRes
	Data *wechat.AppletParams `json:"data"` //微信V3支付结果
}

type WechatJSAPIPayReq struct {
	g.Meta `path:"/front/pay/consumeDeposit/wechatJSAPIPay" tags:"支付" method:"post" summary:"微信WechatJSAPI支付接口"`

	PaymentVo
}

type WechatJSAPIPayRes struct {
	MoneyPayRes
	Data *wechat.AppletParams `json:"data"` //微信V3支付结果
}

type WechatAppPayReq struct {
	g.Meta `path:"/front/pay/consumeDeposit/wechatAppPay" tags:"支付" method:"post" summary:"微信APP支付接口"`

	PaymentVo
}

type WechatAppPayRes struct {
	MoneyPayRes
	Data *wechat.AppletParams `json:"data"` //微信V3支付结果
}

type WechatH5PayReq struct {
	g.Meta `path:"/front/pay/consumeDeposit/wechatH5Pay" tags:"支付" method:"post" summary:"微信H5支付接口"`

	PaymentVo
}

type WechatH5PayRes struct {
	MoneyPayRes
	Data     *wechat.AppletParams `json:"data"`     //微信V3支付结果
	Response string               `json:"response"` //微信V3支付响应
	MwebUrl  string               `json:"mweb_url"` //支付跳转链接

}

// 支付宝
type AlipayH5PayReq struct {
	g.Meta `path:"/front/pay/consumeDeposit/alipayPay" tags:"支付" method:"post" summary:"支付宝H5支付接口"`

	PaymentVo
}

type AlipayH5PayRes struct {
	MoneyPayRes
	Data     *wechat.AppletParams `json:"data"`     //微信V3支付结果
	Response string               `json:"response"` //微信V3支付响应
	MwebUrl  string               `json:"mweb_url"` //支付跳转链接

}
