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

// ConsumeDeposit is the golang structure of table pay_consume_deposit for DAO operations like Where/Data.
type ConsumeDeposit struct {
	g.Meta                  `orm:"table:pay_consume_deposit, do:true"`
	DepositId               interface{} // 支付流水号
	DepositNo               interface{} // 商城支付编号
	DepositTradeNo          interface{} // 交易号:支付宝etc
	OrderId                 interface{} // 商户网站唯一订单号(DOT):合并支付则为多个订单号, 没有创建联合支付交易号
	PaymentChannelId        interface{} // 支付渠道
	DepositSubject          interface{} // 商品名称
	DepositPaymentType      interface{} // 支付方式(ENUM):1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付;
	DepositTradeStatus      interface{} // 交易状态
	DepositSellerId         interface{} // 卖家户号:支付宝etc
	DepositSellerEmail      interface{} // 卖家支付账号
	DepositBuyerId          interface{} // 买家支付用户号
	DepositBuyerEmail       interface{} // 买家支付账号
	CurrencyId              interface{} // 货币编号
	CurrencySymbolLeft      interface{} // 左符号
	DepositTotalFee         interface{} // 交易金额
	DepositQuantity         interface{} // 购买数量
	DepositPrice            interface{} // 商品单价
	DepositBody             interface{} // 商品描述
	DepositIsTotalFeeAdjust interface{} // 是否调整总价
	DepositUseCoupon        interface{} // 是否使用红包买家
	DepositDiscount         interface{} // 折扣
	DepositNotifyTime       interface{} // 通知时间
	DepositNotifyType       interface{} // 通知类型
	DepositNotifyId         interface{} // 通知校验编号
	DepositSignType         interface{} // 签名方式
	DepositSign             interface{} // 签名
	DepositExtraParam       interface{} // 额外参数
	DepositService          interface{} // 支付
	DepositState            interface{} // 支付状态:0-默认; 1-接收正确数据处理完逻辑; 9-异常订单
	DepositAsync            interface{} // 是否同步(BOOL):0-同步; 1-异步回调使用
	DepositReview           interface{} // 收款确认(BOOL):0-未确认;1-已确认
	DepositEnable           interface{} // 是否作废(BOOL):1-正常; 2-作废
	StoreId                 interface{} // 所属店铺:直接交易起作用
	UserId                  interface{} // 所属用户
	ChainId                 interface{} // 所属门店:直接交易起作用
	SubsiteId               interface{} // 所属分站:直接交易起作用
	DepositTime             interface{} // 支付时间
}

type ConsumeDepositListInput struct {
	ml.BaseList
	Where ConsumeDeposit // 查询条件
}

type ConsumeDepositListOutput struct {
	Items   []*entity.ConsumeDeposit // 列表
	Page    int                      // 分页号码
	Total   int                      // 总页数
	Records int                      // 数据总数
	Size    int                      // 单页数量
}

type ConsumeDepositListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
