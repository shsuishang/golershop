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

// ConsumeWithdraw is the golang structure for table consume_withdraw.
type ConsumeWithdraw struct {
	WithdrawId          uint    `json:"withdraw_id"           ` // 编号
	UserId              uint    `json:"user_id"               ` // 用户编号
	StoreId             uint    `json:"store_id"              ` // 所属店铺
	OrderId             string  `json:"order_id"              ` // 所属订单(DOT)
	ReturnId            string  `json:"return_id"             ` // 退款单号(DOT)
	WithdrawAmount      float64 `json:"withdraw_amount"       ` // 提现额度
	WithdrawState       uint    `json:"withdraw_state"        ` // 提现状态(ENUM):0-申请中;1-提现通过;2-驳回;3-打款完成
	WithdrawDesc        string  `json:"withdraw_desc"         ` // 描述
	WithdrawBank        string  `json:"withdraw_bank"         ` // 银行
	WithdrawAccountNo   string  `json:"withdraw_account_no"   ` // 银行账户
	WithdrawAccountName string  `json:"withdraw_account_name" ` // 开户名称
	WithdrawFee         float64 `json:"withdraw_fee"          ` // 提现手续费
	WithdrawTime        uint64  `json:"withdraw_time"         ` // 创建时间
	WithdrawBankflow    string  `json:"withdraw_bankflow"     ` // 银行流水账号
	WithdrawUserId      uint    `json:"withdraw_user_id"      ` // 操作管理员
	WithdrawOpertime    uint64  `json:"withdraw_opertime"     ` // 操作时间
	WithdrawMobile      string  `json:"withdraw_mobile"       ` // 联系手机
	WithdrawTransState  string  `json:"withdraw_trans_state"  ` //
	WithdrawMode        uint    `json:"withdraw_mode"         ` // 提现方式(ENUM):0-余额提现;1-佣金提现
	WithdrawInvoiceNo   string  `json:"withdraw_invoice_no"   ` // 绑定对应的发票号
	SubsiteId           uint    `json:"subsite_id"            ` // 所属分站:0-总站
}
