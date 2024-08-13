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

// ConsumeWithdraw is the golang structure of table pay_consume_withdraw for DAO operations like Where/Data.
type ConsumeWithdraw struct {
	g.Meta              `orm:"table:pay_consume_withdraw, do:true"`
	WithdrawId          interface{} // 编号
	UserId              interface{} // 用户编号
	StoreId             interface{} // 所属店铺
	OrderId             interface{} // 所属订单(DOT)
	ReturnId            interface{} // 退款单号(DOT)
	WithdrawAmount      interface{} // 提现额度
	WithdrawState       interface{} // 提现状态(ENUM):0-申请中;1-提现通过;2-驳回;3-打款完成
	WithdrawDesc        interface{} // 描述
	WithdrawBank        interface{} // 银行
	WithdrawAccountNo   interface{} // 银行账户
	WithdrawAccountName interface{} // 开户名称
	WithdrawFee         interface{} // 提现手续费
	WithdrawTime        interface{} // 创建时间
	WithdrawBankflow    interface{} // 银行流水账号
	WithdrawUserId      interface{} // 操作管理员
	WithdrawOpertime    interface{} // 操作时间
	WithdrawMobile      interface{} // 联系手机
	WithdrawTransState  interface{} //
	WithdrawMode        interface{} // 提现方式(ENUM):0-余额提现;1-佣金提现
	WithdrawInvoiceNo   interface{} // 绑定对应的发票号
	SubsiteId           interface{} // 所属分站:0-总站
}

type ConsumeWithdrawListInput struct {
	ml.BaseList
	Where ConsumeWithdraw // 查询条件
}

type ConsumeWithdrawListOutput struct {
	Items   []*entity.ConsumeWithdraw // 列表
	Page    int                       // 分页号码
	Total   int                       // 总页数
	Records int                       // 数据总数
	Size    int                       // 单页数量
}

type ConsumeWithdrawListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
