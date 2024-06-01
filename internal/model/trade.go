package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type OfflinePayVo struct {
	OrderId string `json:"order_id"    dc:"订单编号"`

	PaymentChannelId  uint        `json:"payment_channel_id"  d:"26"`   //支付渠道
	DepositTradeNo    string      `json:"deposit_trade_no"            ` // 交易号:支付宝etc
	DepositNotifyTime *gtime.Time `json:"deposit_notify_time" `         //时间
	DepositTotalFee   float64     `json:"deposit_total_fee"           ` // 交易金额
}

type DashboardTopVo struct {
	Today    interface{} `json:"today"          `      //今日
	Yestoday interface{} `json:"yestoday"            ` // 昨日
	Daym2m   interface{} `json:"daym2m" `              //日环比
	Month    interface{} `json:"month"           `     // 本月
}

type PaymentInput struct {
	OrderId            []string `json:"order_id" v:"required#请输入订单编号"    `        // 订单编号(DOT)
	PaymentChannelId   uint     `json:"payment_channel_id" d:"26"             `   // 支付渠道
	Openid             string   `json:"openid"                                 `  // openid
	DepositPaymentType uint     `json:"deposit_payment_type" d:"1302"         `   // 支付方式(ENUM):1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付;
	Password           string   `json:"password"                                ` // 支付密码
	PmMoney            float64  `json:"pm_money"                                ` // 余额支付
	PmRechargeCard     float64  `json:"pm_recharge_card"                       `  // 充值卡支付
	PmPoints           float64  `json:"pm_points"                              `  // 积分支付
	PmCredit           float64  `json:"pm_credit"                              `  // 信用账户
	PmRedpack          float64  `json:"pm_redpack" `                              //红包账户
}

type PaymentOutput struct {
	TradeNo string  `json:"trade_no"     `                           // 交易订单号)
	Title   string  `json:"title"           `                        // 订单标题
	Amount  float64 `json:"amount"                                 ` // 支付金额
}

type PayMetVo struct {
	PaymentMetId uint `json:"payment_met_id"    dc:"付款账户"`

	PmMoney        float64 `json:"pm_money" `         //余额
	PmRechargeCard float64 `json:"pm_recharge_card" ` //充值卡
	PmPoints       float64 `json:"pm_points" `        //积分
	PmCredit       float64 `json:"pm_credit" `        //信用账户
	PmRedpack      float64 `json:"pm_redpack" `       //红包账户

	PaymentChannelId uint `json:"payment_channel_id"          ` // 支付渠道
	PaymentTypeId    uint `json:"payment_type_id"        `      // 支付方式(ENUM):1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付;
}

type ProcessPayOutput struct {
	OrderId string `json:"order_id"    dc:"订单编号"`
	Paid    bool   `json:"paid"    dc:"是否支付成功"`
}

type DistributionOrderVo struct {
	OrderId            string  `json:"order_id" description:"订单编号"`                  // 订单编号
	StoreId            uint    `json:"store_id" description:"卖家店铺编号"`                // 卖家店铺编号
	UserId             uint    `json:"user_id" description:"买家编号"`                   // 买家编号
	OrderCommissionFee float64 `json:"order_commission_fee" description:"平台交易佣金"`    // 平台交易佣金
	SalespersonId      uint    `json:"salesperson_id" description:"销售员编号:用户编号"`      // 销售员编号:用户编号
	DistributorUserId  uint    `json:"distributor_user_id" description:"分销商编号:用户编号"` // 分销商编号:用户编号
}

type OrderFreightVo struct {
	CanDelivery    bool    `json:"can_delivery" swagger:"description:是否可配送"`
	FreightFreeMin float64 `json:"freight_free_min" swagger:"description:免运费额度"`
	Freight        float64 `json:"freight" swagger:"description:运费"`
}
