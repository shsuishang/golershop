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

// ConsumeTrade is the golang structure for table consume_trade.
type ConsumeTrade struct {
	ConsumeTradeId           uint    `json:"consume_trade_id"            ` // 交易订单编号
	TradeTitle               string  `json:"trade_title"                 ` // 标题
	OrderId                  string  `json:"order_id"                    ` // 商户订单编号
	BuyerId                  uint    `json:"buyer_id"                    ` // 买家编号
	BuyerStoreId             uint    `json:"buyer_store_id"              ` // 买家是否有店铺
	StoreId                  uint    `json:"store_id"                    ` // 店铺编号
	SubsiteId                uint    `json:"subsite_id"                  ` // 所属分站:0-总站
	SellerId                 uint    `json:"seller_id"                   ` // 卖家编号
	ChainId                  uint    `json:"chain_id"                    ` // 门店编号
	TradeIsPaid              uint    `json:"trade_is_paid"               ` // 支付状态
	TradeTypeId              uint    `json:"trade_type_id"               ` // 交易类型(ENUM):1201-购物; 1202-转账; 1203-充值; 1204-提现; 1205-销售; 1206-佣金;
	PaymentChannelId         uint    `json:"payment_channel_id"          ` // 支付渠道
	TradeModeId              uint    `json:"trade_mode_id"               ` // 交易模式(ENUM):1-担保交易;  2-直接交易
	RechargeLevelId          uint    `json:"recharge_level_id"           ` // 充值编号
	CurrencyId               uint    `json:"currency_id"                 ` // 货币编号
	CurrencySymbolLeft       string  `json:"currency_symbol_left"        ` // 左符号
	OrderPaymentAmount       float64 `json:"order_payment_amount"        ` // 总付款额度: trade_payment_amount + trade_payment_money + trade_payment_recharge_card + trade_payment_points
	OrderCommissionFee       float64 `json:"order_commission_fee"        ` // 平台交易佣金
	TradePaymentAmount       float64 `json:"trade_payment_amount"        ` // 实付金额:在线支付金额,此为订单默认需要支付额度。
	TradePaymentMoney        float64 `json:"trade_payment_money"         ` // 余额支付
	TradePaymentRechargeCard float64 `json:"trade_payment_recharge_card" ` // 充值卡余额支付
	TradePaymentPoints       float64 `json:"trade_payment_points"        ` // 积分支付
	TradePaymentSp           float64 `json:"trade_payment_sp"            ` // 众宝支付
	TradePaymentCredit       float64 `json:"trade_payment_credit"        ` // 信用支付
	TradePaymentRedpack      float64 `json:"trade_payment_redpack"       ` // 红包支付
	TradeDiscount            float64 `json:"trade_discount"              ` // 折扣优惠
	TradeAmount              float64 `json:"trade_amount"                ` // 总额虚拟:trade_order_amount + trade_discount
	TradeDesc                string  `json:"trade_desc"                  ` // 描述
	TradeRemark              string  `json:"trade_remark"                ` // 备注
	TradeCreateTime          uint64  `json:"trade_create_time"           ` // 创建时间
	TradePaidTime            int64   `json:"trade_paid_time"             ` // 付款时间
	TradeDelete              uint    `json:"trade_delete"                ` // 是否删除
	Version                  uint    `json:"version"                     ` // 版本
}
