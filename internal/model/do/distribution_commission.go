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

// DistributionCommission is the golang structure of table pay_distribution_commission for DAO operations like Where/Data.
type DistributionCommission struct {
	g.Meta                             `orm:"table:pay_distribution_commission, do:true"`
	UserId                             interface{} // 店铺编号
	CommissionAmount                   interface{} // 佣金总额:历史总额度
	CommissionDirectsellerAmount0      interface{} // 销售员佣金
	CommissionDirectsellerAmount1      interface{} // 二级销售员
	CommissionDirectsellerAmount2      interface{} // 三级销售员
	CommissionBuyAmount0               interface{} // 推广消费佣金
	CommissionBuyAmount1               interface{} // 消费佣金
	CommissionBuyAmount2               interface{} // 消费佣金
	CommissionClickAmount0             interface{} // 本店流量佣金
	CommissionClickAmount1             interface{} // 一级流量佣金
	CommissionClickAmount2             interface{} // 二级流量佣金
	CommissionRegAmount0               interface{} // 本店注册佣金
	CommissionRegAmount1               interface{} // 一级注册佣金
	CommissionRegAmount2               interface{} // 二级注册佣金
	CommissionSettled                  interface{} // 已经结算佣金
	CommissionDirectsellerSettled      interface{} // 销售员已经结算
	CommissionBuySettled               interface{} // 推广员已经结算
	CommissionBuyDa                    interface{} // 区代理收益
	CommissionBuyCa                    interface{} // 市代理收益
	CommissionDirectsellerDa           interface{} // 区代理收益
	CommissionDirectsellerCa           interface{} // 市代理收益
	CommissionBuyTrade0                interface{} // 交易总额
	CommissionBuyTrade1                interface{} // 交易总额
	CommissionBuyTrade2                interface{} // 交易总额
	CommissionBuyDaTrade               interface{} // 交易总额
	CommissionBuyCaTrade               interface{} // 交易总额
	CommissionDirectsellerTrade0       interface{} // 交易总额
	CommissionDirectsellerTrade1       interface{} // 交易总额
	CommissionDirectsellerTrade2       interface{} // 交易总额
	CommissionDirectsellerDaTrade      interface{} // 交易总额
	CommissionDirectsellerCaTrade      interface{} // 交易总额
	CommissionPartnerBuyTrade          interface{} // 合伙人交易总额
	CommissionPartnerDirectsellerTrade interface{} // 合伙人交易总额
	CommissionPartnerDepositTrade      interface{} // 合伙人充值总额
	CommissionDistributorAmount        interface{} // 分销商收益
	CommissionSalespersonAmount        interface{} // 销售员收益
	CommissionRefundAmount             interface{} // 退款总佣金
	Version                            interface{} // 版本
}

type DistributionCommissionListInput struct {
	ml.BaseList
	Where DistributionCommission // 查询条件
}

type DistributionCommissionListOutput struct {
	Items   []*entity.DistributionCommission // 列表
	Page    int                              // 分页号码
	Total   int                              // 总页数
	Records int                              // 数据总数
	Size    int                              // 单页数量
}

type DistributionCommissionListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
