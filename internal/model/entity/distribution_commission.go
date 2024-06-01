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

// DistributionCommission is the golang structure for table distribution_commission.
type DistributionCommission struct {
	UserId                             uint    `json:"user_id"                               ` // 店铺编号
	CommissionAmount                   float64 `json:"commission_amount"                     ` // 佣金总额:历史总额度
	CommissionDirectsellerAmount0      float64 `json:"commission_directseller_amount_0"      ` // 销售员佣金
	CommissionDirectsellerAmount1      float64 `json:"commission_directseller_amount_1"      ` // 二级销售员
	CommissionDirectsellerAmount2      float64 `json:"commission_directseller_amount_2"      ` // 三级销售员
	CommissionBuyAmount0               float64 `json:"commission_buy_amount_0"               ` // 推广消费佣金
	CommissionBuyAmount1               float64 `json:"commission_buy_amount_1"               ` // 消费佣金
	CommissionBuyAmount2               float64 `json:"commission_buy_amount_2"               ` // 消费佣金
	CommissionClickAmount0             float64 `json:"commission_click_amount_0"             ` // 本店流量佣金
	CommissionClickAmount1             float64 `json:"commission_click_amount_1"             ` // 一级流量佣金
	CommissionClickAmount2             float64 `json:"commission_click_amount_2"             ` // 二级流量佣金
	CommissionRegAmount0               float64 `json:"commission_reg_amount_0"               ` // 本店注册佣金
	CommissionRegAmount1               float64 `json:"commission_reg_amount_1"               ` // 一级注册佣金
	CommissionRegAmount2               float64 `json:"commission_reg_amount_2"               ` // 二级注册佣金
	CommissionSettled                  float64 `json:"commission_settled"                    ` // 已经结算佣金
	CommissionDirectsellerSettled      float64 `json:"commission_directseller_settled"       ` // 销售员已经结算
	CommissionBuySettled               float64 `json:"commission_buy_settled"                ` // 推广员已经结算
	CommissionBuyDa                    float64 `json:"commission_buy_da"                     ` // 区代理收益
	CommissionBuyCa                    float64 `json:"commission_buy_ca"                     ` // 市代理收益
	CommissionDirectsellerDa           float64 `json:"commission_directseller_da"            ` // 区代理收益
	CommissionDirectsellerCa           float64 `json:"commission_directseller_ca"            ` // 市代理收益
	CommissionBuyTrade0                float64 `json:"commission_buy_trade_0"                ` // 交易总额
	CommissionBuyTrade1                float64 `json:"commission_buy_trade_1"                ` // 交易总额
	CommissionBuyTrade2                float64 `json:"commission_buy_trade_2"                ` // 交易总额
	CommissionBuyDaTrade               float64 `json:"commission_buy_da_trade"               ` // 交易总额
	CommissionBuyCaTrade               float64 `json:"commission_buy_ca_trade"               ` // 交易总额
	CommissionDirectsellerTrade0       float64 `json:"commission_directseller_trade_0"       ` // 交易总额
	CommissionDirectsellerTrade1       float64 `json:"commission_directseller_trade_1"       ` // 交易总额
	CommissionDirectsellerTrade2       float64 `json:"commission_directseller_trade_2"       ` // 交易总额
	CommissionDirectsellerDaTrade      float64 `json:"commission_directseller_da_trade"      ` // 交易总额
	CommissionDirectsellerCaTrade      float64 `json:"commission_directseller_ca_trade"      ` // 交易总额
	CommissionPartnerBuyTrade          float64 `json:"commission_partner_buy_trade"          ` // 合伙人交易总额
	CommissionPartnerDirectsellerTrade float64 `json:"commission_partner_directseller_trade" ` // 合伙人交易总额
	CommissionPartnerDepositTrade      float64 `json:"commission_partner_deposit_trade"      ` // 合伙人充值总额
	CommissionDistributorAmount        float64 `json:"commission_distributor_amount"         ` // 分销商收益
	CommissionSalespersonAmount        float64 `json:"commission_salesperson_amount"         ` // 销售员收益
	CommissionRefundAmount             float64 `json:"commission_refund_amount"              ` // 退款总佣金
	Version                            uint    `json:"version"                               ` // 版本
}
