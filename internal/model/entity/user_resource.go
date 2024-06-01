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

// UserResource is the golang structure for table user_resource.
type UserResource struct {
	UserId                 uint    `json:"user_id"                   ` // 用户编号
	CurrencyId             uint    `json:"currency_id"               ` // 货币编号
	CurrencySymbolLeft     string  `json:"currency_symbol_left"      ` // 左符号
	UserMoney              float64 `json:"user_money"                ` // 用户资金
	UserMoneyFrozen        float64 `json:"user_money_frozen"         ` // 冻结资金:待结算余额
	UserRechargeCard       float64 `json:"user_recharge_card"        ` // 充值卡余额
	UserRechargeCardFrozen float64 `json:"user_recharge_card_frozen" ` // 冻结充值卡:待结算
	UserPoints             float64 `json:"user_points"               ` // 积分
	UserPointsFrozen       float64 `json:"user_points_frozen"        ` // 冻结积分
	UserExp                uint64  `json:"user_exp"                  ` // 平台总经验
	UserCredit             float64 `json:"user_credit"               ` // 可用信用
	UserCreditFrozen       float64 `json:"user_credit_frozen"        ` // 冻结额度
	UserCreditUsed         float64 `json:"user_credit_used"          ` // 使用信用
	UserCreditTotal        float64 `json:"user_credit_total"         ` // 信用额度
	UserMargin             float64 `json:"user_margin"               ` // 保证金
	UserRedpack            float64 `json:"user_redpack"              ` // 红包额度
	UserRedpackFrozen      float64 `json:"user_redpack_frozen"       ` // 红包冻结额度
	UserSp                 float64 `json:"user_sp"                   ` // 积分2
	UserSpFrozen           float64 `json:"user_sp_frozen"            ` // 冻结积分2
	Version                uint    `json:"version"                   ` // 乐观锁
}
