package consts

const (
	PAYMENT_MET_MONEY         uint = 1 //余额支付
	PAYMENT_MET_RECHARGE_CARD uint = 2 //充值卡支付
	PAYMENT_MET_POINTS        uint = 3 //积分支付
	PAYMENT_MET_CREDIT        uint = 4 //信用支付
	PAYMENT_MET_REDPACK       uint = 5 //红包支付
	PAYMENT_MET_OFFLINE       uint = 6 //线下支付
	PAYMENT_MET_SP            uint = 7 //众宝支付

	PAYMENT_CHANNEL_WECHAT  uint = 1403 //微信支付
	PAYMENT_CHANNEL_ALIPAY  uint = 1401 //支付宝支付
	PAYMENT_CHANNEL_OFFLINE uint = 1422 //线下支付
	PAYMENT_CHANNEL_MONEY   uint = 1406 //余额支付
	PAYMENT_CHANNEL_POINTS  uint = 1413 //积分支付
)
