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

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderData is the golang structure for table order_data.
type OrderData struct {
	OrderId                  string      `json:"order_id"                    ` // 订单编号
	OrderDesc                string      `json:"order_desc"                  ` // 订单描述
	OrderDelayTime           uint        `json:"order_delay_time"            ` // 延迟时间,默认为0 - 收货确认
	DeliveryTypeId           uint        `json:"delivery_type_id"            ` // 配送方式
	DeliveryTimeId           uint        `json:"delivery_time_id"            ` // 配送时间:要求，不限、周一~周五、周末等等
	DeliveryTime             int64       `json:"delivery_time"               ` // 配送日期
	DeliveryIstimer          bool        `json:"delivery_istimer"            ` // 是否定时配送(BOOL):0-不定时;1-定时
	OrderMessage             string      `json:"order_message"               ` // 买家订单留言
	OrderItemAmount          float64     `json:"order_item_amount"           ` // 商品总价格/商品金额, 不包含运费
	OrderDiscountAmount      float64     `json:"order_discount_amount"       ` // 折扣价格/优惠总金额
	OrderAdjustFee           float64     `json:"order_adjust_fee"            ` // 手工调整费用店铺优惠
	OrderPointsFee           float64     `json:"order_points_fee"            ` // 积分抵扣费用
	OrderShippingFeeAmount   float64     `json:"order_shipping_fee_amount"   ` // 运费价格/运费金额
	OrderShippingFee         float64     `json:"order_shipping_fee"          ` // 实际运费金额-卖家可修改
	VoucherId                uint        `json:"voucher_id"                  ` // 代金券id/优惠券/返现:发放选择使用
	VoucherNumber            string      `json:"voucher_number"              ` // 代金券编码
	VoucherPrice             float64     `json:"voucher_price"               ` // 代金券面额
	RedpacketId              uint        `json:"redpacket_id"                ` // 红包id-平台代金券
	RedpacketNumber          string      `json:"redpacket_number"            ` // 红包编码
	RedpacketPrice           float64     `json:"redpacket_price"             ` // 红包面额
	OrderRedpacketPrice      float64     `json:"order_redpacket_price"       ` // 红包抵扣订单金额
	OrderResourceExt1        float64     `json:"order_resource_ext_1"        ` // 第二需要支付资源例如积分
	OrderResourceExt2        float64     `json:"order_resource_ext_2"        ` // 众宝
	OrderResourceExt3        float64     `json:"order_resource_ext_3"        ` // 金宝
	TradePaymentMoney        float64     `json:"trade_payment_money"         ` // 余额支付
	TradePaymentRechargeCard float64     `json:"trade_payment_recharge_card" ` // 充值卡支付
	TradePaymentCredit       float64     `json:"trade_payment_credit"        ` // 信用支付
	OrderRefundStatus        uint        `json:"order_refund_status"         ` // 退款状态:0-是无退款;1-是部分退款;2-是全部退款
	OrderRefundAmount        float64     `json:"order_refund_amount"         ` // 退款金额:申请额度
	OrderRefundAgreeAmount   float64     `json:"order_refund_agree_amount"   ` // 退款金额:同意额度
	OrderRefundAgreeCash     float64     `json:"order_refund_agree_cash"     ` // 已同意退的现金
	OrderRefundAgreePoints   float64     `json:"order_refund_agree_points"   ` // 已退的积分额度
	OrderReturnStatus        uint        `json:"order_return_status"         ` // 退货状态(ENUM):0-是无退货;1-是部分退货;2-是全部退货
	OrderReturnNum           uint        `json:"order_return_num"            ` // 退货数量
	OrderReturnIds           string      `json:"order_return_ids"            ` // 退货单编号s(DOT):冗余
	OrderCommissionFee       float64     `json:"order_commission_fee"        ` // 平台交易佣金
	OrderCommissionFeeRefund float64     `json:"order_commission_fee_refund" ` // 交易佣金-退款
	OrderPointsAdd           float64     `json:"order_points_add"            ` // 订单赠送积分
	OrderActivityData        string      `json:"order_activity_data"         ` // 促销信息
	OrderCancelIdentity      uint        `json:"order_cancel_identity"       ` // 订单取消者身份(ENUM):1-买家; 2-卖家; 3-系统
	OrderCancelReason        string      `json:"order_cancel_reason"         ` // 订单取消原因
	OrderCancelTime          *gtime.Time `json:"order_cancel_time"           ` // 订单取消时间
	OrderBpAdd               uint        `json:"order_bp_add"                ` // 赠送资源2
	OrderRebate              float64     `json:"order_rebate"                ` // 订单返利
	BuyerMobile              uint        `json:"buyer_mobile"                ` // 手机号码
	BuyerContacter           string      `json:"buyer_contacter"             ` // 联系人
	ActivityManhuiId         string      `json:"activity_manhui_id"          ` // 满返优惠券活动id(DOT)
	ActivityDoublePointsId   uint        `json:"activity_double_points_id"   ` // 活动-多倍积分id
	OrderDoublePointsAdd     float64     `json:"order_double_points_add"     ` // 活动-多倍积分
	ActivityVoucherId        string      `json:"activity_voucher_id"         ` // 满返用户优惠券id(DOT)
	OrderActivityManhuiState uint        `json:"order_activity_manhui_state" ` // 满返优惠券发放状态(ENUM):1000-无需发放;1001-待发放; 1002-已发放; 1003-发放异常
	Version                  uint        `json:"version"                     ` // 版本
}
