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

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// ConsumeTrade is the golang structure of table pay_consume_trade for DAO operations like Where/Data.
type ConsumeTrade struct {
	g.Meta                   `orm:"table:pay_consume_trade, do:true"`
	ConsumeTradeId           interface{} // 交易订单编号
	TradeTitle               interface{} // 标题
	OrderId                  interface{} // 商户订单编号
	BuyerId                  interface{} // 买家编号
	BuyerStoreId             interface{} // 买家是否有店铺
	StoreId                  interface{} // 店铺编号
	SubsiteId                interface{} // 所属分站:0-总站
	SellerId                 interface{} // 卖家编号
	ChainId                  interface{} // 门店编号
	TradeIsPaid              interface{} // 支付状态
	TradeTypeId              interface{} // 交易类型(ENUM):1201-购物; 1202-转账; 1203-充值; 1204-提现; 1205-销售; 1206-佣金;
	PaymentChannelId         interface{} // 支付渠道
	TradeModeId              interface{} // 交易模式(ENUM):1-担保交易;  2-直接交易
	RechargeLevelId          interface{} // 充值编号
	CurrencyId               interface{} // 货币编号
	CurrencySymbolLeft       interface{} // 左符号
	OrderPaymentAmount       interface{} // 总付款额度: trade_payment_amount + trade_payment_money + trade_payment_recharge_card + trade_payment_points
	OrderCommissionFee       interface{} // 平台交易佣金
	TradePaymentAmount       interface{} // 实付金额:在线支付金额,此为订单默认需要支付额度。
	TradePaymentMoney        interface{} // 余额支付
	TradePaymentRechargeCard interface{} // 充值卡余额支付
	TradePaymentPoints       interface{} // 积分支付
	TradePaymentSp           interface{} // 众宝支付
	TradePaymentCredit       interface{} // 信用支付
	TradePaymentRedpack      interface{} // 红包支付
	TradeDiscount            interface{} // 折扣优惠
	TradeAmount              interface{} // 总额虚拟:trade_order_amount + trade_discount
	TradeDesc                interface{} // 描述
	TradeRemark              interface{} // 备注
	TradeCreateTime          interface{} // 创建时间
	TradePaidTime            interface{} // 付款时间
	TradeDelete              interface{} // 是否删除
	Version                  interface{} // 版本
}

type ConsumeTradeListInput struct {
	ml.BaseList
	Where ConsumeTrade // 查询条件
}

type ConsumeTradeListOutput struct {
	Items   []*entity.ConsumeTrade // 列表
	Page    int                    // 分页号码
	Total   int                    // 总页数
	Records int                    // 数据总数
	Size    int                    // 单页数量
}

type ConsumeTradeListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
