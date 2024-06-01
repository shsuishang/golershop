package consts

// LevelCode 定义会员经验类型的常量
const (
	EXP_TYPE_REG              uint = 1 // 会员注册
	EXP_TYPE_LOGIN            uint = 2 // 会员登录
	EXP_TYPE_EVALUATE_PRODUCT uint = 3 // 商品评论
	EXP_TYPE_EVALUATE_STORE   uint = 6 // 店铺评论
	EXP_TYPE_CONSUME          uint = 4 // 购买商品
	EXP_TYPE_OTHER            uint = 5 // 管理员操作
	EXP_TYPE_EXCHANGE_PRODUCT uint = 7 // 积分换购商品
	EXP_TYPE_EXCHANGE_VOUCHER uint = 8 // 积分兑换优惠券
)
