// +----------------------------------------------------------------------
// | ShopSuite商城系统 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 随商信息技术（上海）有限公司
// +----------------------------------------------------------------------
// | 未获商业授权前，不得将本软件用于商业用途。禁止整体或任何部分基础上以发展任何派生版本、
// | 修改版本或第三方版本用于重新分发。
// +----------------------------------------------------------------------
// | 官方网站: https://www.shopsuite.cn  https://www.golershop.cn
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本公司对该软件产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细见https://www.golershop.cn/policy
// +----------------------------------------------------------------------

package entity

// ConsumeDeposit is the golang structure for table consume_deposit.
type ConsumeDeposit struct {
	DepositId               uint64  `json:"deposit_id"                  ` // 支付流水号
	DepositNo               string  `json:"deposit_no"                  ` // 商城支付编号
	DepositTradeNo          string  `json:"deposit_trade_no"            ` // 交易号:支付宝etc
	OrderId                 string  `json:"order_id"                    ` // 商户网站唯一订单号(DOT):合并支付则为多个订单号, 没有创建联合支付交易号
	PaymentChannelId        uint    `json:"payment_channel_id"          ` // 支付渠道
	DepositSubject          string  `json:"deposit_subject"             ` // 商品名称
	DepositPaymentType      uint    `json:"deposit_payment_type"        ` // 支付方式(ENUM):1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付;
	DepositTradeStatus      string  `json:"deposit_trade_status"        ` // 交易状态
	DepositSellerId         string  `json:"deposit_seller_id"           ` // 卖家户号:支付宝etc
	DepositSellerEmail      string  `json:"deposit_seller_email"        ` // 卖家支付账号
	DepositBuyerId          string  `json:"deposit_buyer_id"            ` // 买家支付用户号
	DepositBuyerEmail       string  `json:"deposit_buyer_email"         ` // 买家支付账号
	CurrencyId              uint    `json:"currency_id"                 ` // 货币编号
	CurrencySymbolLeft      string  `json:"currency_symbol_left"        ` // 左符号
	DepositTotalFee         float64 `json:"deposit_total_fee"           ` // 交易金额
	DepositQuantity         uint    `json:"deposit_quantity"            ` // 购买数量
	DepositPrice            float64 `json:"deposit_price"               ` // 商品单价
	DepositBody             string  `json:"deposit_body"                ` // 商品描述
	DepositIsTotalFeeAdjust uint    `json:"deposit_is_total_fee_adjust" ` // 是否调整总价
	DepositUseCoupon        uint    `json:"deposit_use_coupon"          ` // 是否使用红包买家
	DepositDiscount         float64 `json:"deposit_discount"            ` // 折扣
	DepositNotifyTime       string  `json:"deposit_notify_time"         ` // 通知时间
	DepositNotifyType       string  `json:"deposit_notify_type"         ` // 通知类型
	DepositNotifyId         string  `json:"deposit_notify_id"           ` // 通知校验编号
	DepositSignType         string  `json:"deposit_sign_type"           ` // 签名方式
	DepositSign             string  `json:"deposit_sign"                ` // 签名
	DepositExtraParam       string  `json:"deposit_extra_param"         ` // 额外参数
	DepositService          string  `json:"deposit_service"             ` // 支付
	DepositState            uint    `json:"deposit_state"               ` // 支付状态:0-默认; 1-接收正确数据处理完逻辑; 9-异常订单
	DepositAsync            bool    `json:"deposit_async"               ` // 是否同步(BOOL):0-同步; 1-异步回调使用
	DepositReview           bool    `json:"deposit_review"              ` // 收款确认(BOOL):0-未确认;1-已确认
	DepositEnable           bool    `json:"deposit_enable"              ` // 是否作废(BOOL):1-正常; 2-作废
	StoreId                 uint    `json:"store_id"                    ` // 所属店铺:直接交易起作用
	UserId                  uint    `json:"user_id"                     ` // 所属用户
	ChainId                 uint    `json:"chain_id"                    ` // 所属门店:直接交易起作用
	SubsiteId               uint    `json:"subsite_id"                  ` // 所属分站:直接交易起作用
	DepositTime             uint64  `json:"deposit_time"                ` // 支付时间
}
