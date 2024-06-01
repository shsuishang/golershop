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
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// OrderData is the golang structure of table trade_order_data for DAO operations like Where/Data.
type OrderData struct {
	g.Meta                   `orm:"table:trade_order_data, do:true"`
	OrderId                  interface{} // 订单编号
	OrderDesc                interface{} // 订单描述
	OrderDelayTime           interface{} // 延迟时间,默认为0 - 收货确认
	DeliveryTypeId           interface{} // 配送方式
	DeliveryTimeId           interface{} // 配送时间:要求，不限、周一~周五、周末等等
	DeliveryTime             interface{} // 配送日期
	DeliveryIstimer          interface{} // 是否定时配送(BOOL):0-不定时;1-定时
	OrderMessage             interface{} // 买家订单留言
	OrderItemAmount          interface{} // 商品总价格/商品金额, 不包含运费
	OrderDiscountAmount      interface{} // 折扣价格/优惠总金额
	OrderAdjustFee           interface{} // 手工调整费用店铺优惠
	OrderPointsFee           interface{} // 积分抵扣费用
	OrderShippingFeeAmount   interface{} // 运费价格/运费金额
	OrderShippingFee         interface{} // 实际运费金额-卖家可修改
	VoucherId                interface{} // 代金券id/优惠券/返现:发放选择使用
	VoucherNumber            interface{} // 代金券编码
	VoucherPrice             interface{} // 代金券面额
	RedpacketId              interface{} // 红包id-平台代金券
	RedpacketNumber          interface{} // 红包编码
	RedpacketPrice           interface{} // 红包面额
	OrderRedpacketPrice      interface{} // 红包抵扣订单金额
	OrderResourceExt1        interface{} // 第二需要支付资源例如积分
	OrderResourceExt2        interface{} // 众宝
	OrderResourceExt3        interface{} // 金宝
	TradePaymentMoney        interface{} // 余额支付
	TradePaymentRechargeCard interface{} // 充值卡支付
	TradePaymentCredit       interface{} // 信用支付
	OrderRefundStatus        interface{} // 退款状态:0-是无退款;1-是部分退款;2-是全部退款
	OrderRefundAmount        interface{} // 退款金额:申请额度
	OrderRefundAgreeAmount   interface{} // 退款金额:同意额度
	OrderRefundAgreeCash     interface{} // 已同意退的现金
	OrderRefundAgreePoints   interface{} // 已退的积分额度
	OrderReturnStatus        interface{} // 退货状态(ENUM):0-是无退货;1-是部分退货;2-是全部退货
	OrderReturnNum           interface{} // 退货数量
	OrderReturnIds           interface{} // 退货单编号s(DOT):冗余
	OrderCommissionFee       interface{} // 平台交易佣金
	OrderCommissionFeeRefund interface{} // 交易佣金-退款
	OrderPointsAdd           interface{} // 订单赠送积分
	OrderActivityData        interface{} // 促销信息
	OrderCancelIdentity      interface{} // 订单取消者身份(ENUM):1-买家; 2-卖家; 3-系统
	OrderCancelReason        interface{} // 订单取消原因
	OrderCancelTime          *gtime.Time // 订单取消时间
	OrderBpAdd               interface{} // 赠送资源2
	OrderRebate              interface{} // 订单返利
	BuyerMobile              interface{} // 手机号码
	BuyerContacter           interface{} // 联系人
	ActivityManhuiId         interface{} // 满返优惠券活动id(DOT)
	ActivityDoublePointsId   interface{} // 活动-多倍积分id
	OrderDoublePointsAdd     interface{} // 活动-多倍积分
	ActivityVoucherId        interface{} // 满返用户优惠券id(DOT)
	OrderActivityManhuiState interface{} // 满返优惠券发放状态(ENUM):1000-无需发放;1001-待发放; 1002-已发放; 1003-发放异常
	Version                  interface{} // 版本
}

type OrderDataListInput struct {
	ml.BaseList
	Where OrderData // 查询条件
}

type OrderDataListOutput struct {
	Items   []*entity.OrderData // 列表
	Page    int                 // 分页号码
	Total   int                 // 总页数
	Records int                 // 数据总数
	Size    int                 // 单页数量
}

type OrderDataListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
