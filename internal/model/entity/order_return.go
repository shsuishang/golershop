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

// OrderReturn is the golang structure for table order_return.
type OrderReturn struct {
	ReturnId                   string      `json:"return_id"                     ` // 退单号
	ServiceTypeId              uint        `json:"service_type_id"               ` // 服务类型(ENUM):1-退款;2-退货;3-换货;4-维修
	OrderId                    string      `json:"order_id"                      ` // 订单编号
	ReturnRefundAmount         float64     `json:"return_refund_amount"          ` // 退款金额
	ReturnRefundPoint          float64     `json:"return_refund_point"           ` // 积分部分
	StoreId                    uint        `json:"store_id"                      ` // 店铺编号
	BuyerUserId                uint        `json:"buyer_user_id"                 ` // 买家编号
	BuyerStoreId               uint        `json:"buyer_store_id"                ` // 买家是否有店铺
	ReturnAddTime              int64       `json:"return_add_time"               ` // 添加时间
	ReturnReasonId             uint        `json:"return_reason_id"              ` // 退款理由编号
	ReturnBuyerMessage         string      `json:"return_buyer_message"          ` // 买家退货备注
	ReturnAddrContacter        string      `json:"return_addr_contacter"         ` // 收货人
	ReturnTel                  string      `json:"return_tel"                    ` // 联系电话
	ReturnAddr                 string      `json:"return_addr"                   ` // 收货地址详情
	ReturnPostCode             int         `json:"return_post_code"              ` // 邮编
	ExpressId                  uint        `json:"express_id"                    ` // 物流公司编号
	ReturnTrackingName         string      `json:"return_tracking_name"          ` // 物流名称
	ReturnTrackingNumber       string      `json:"return_tracking_number"        ` // 物流单号
	PlantformReturnStateId     uint        `json:"plantform_return_state_id"     ` // 申请状态平台(ENUM):3180-未申请;3181-待处理;3182-为已完成
	ReturnStateId              uint        `json:"return_state_id"               ` // 卖家处理状态(ENUM): 3100-【客户】提交退单;3105-退单审核;3110-收货确认;3115-退款确认;3120-客户】收款确认;3125-完成
	ReturnIsPaid               bool        `json:"return_is_paid"                ` // 退款完成
	ReturnIsShippingFee        bool        `json:"return_is_shipping_fee"        ` // 退货类型(BOOL): 0-退款单;1-退运费单
	ReturnShippingFee          float64     `json:"return_shipping_fee"           ` // 退运费额度
	ReturnFlag                 uint        `json:"return_flag"                   ` // 退货类型(ENUM): 0-不用退货;1-需要退货
	ReturnType                 uint        `json:"return_type"                   ` // 申请类型(ENUM): 1-退款申请; 2-退货申请; 3-虚拟退款
	ReturnOrderLock            uint        `json:"return_order_lock"             ` // 订单锁定类型(BOOL):1-不用锁定;2-需要锁定
	ReturnItemStateId          uint        `json:"return_item_state_id"          ` // 物流状态(LIST):2030-待发货;2040-已发货/待收货确认;2060-已完成/已签收;2070-已取消/已作废;
	ReturnStoreTime            *gtime.Time `json:"return_store_time"             ` // 商家处理时间
	ReturnStoreMessage         string      `json:"return_store_message"          ` // 商家备注
	ReturnCommisionFee         float64     `json:"return_commision_fee"          ` // 退还佣金
	ReturnFinishTime           *gtime.Time `json:"return_finish_time"            ` // 退款完成时间
	ReturnPlatformMessage      string      `json:"return_platform_message"       ` // 平台留言
	ReturnIsSettlemented       uint        `json:"return_is_settlemented"        ` // 订单是否结算(BOOL): 0-未结算; 1-已结算
	ReturnSettlementTime       *gtime.Time `json:"return_settlement_time"        ` // 订单结算时间
	ReturnChannelCode          string      `json:"return_channel_code"           ` // 退款渠道(ENUM):money-余额;alipay-支付宝;wx_native-微信
	ReturnChannelFlag          uint        `json:"return_channel_flag"           ` // 渠道是否退款(ENUM): 0-待退; 1-已退; 2-异常
	ReturnChannelTime          *gtime.Time `json:"return_channel_time"           ` // 渠道退款时间
	ReturnChannelTransId       string      `json:"return_channel_trans_id"       ` // 渠道退款单号
	DepositTradeNo             string      `json:"deposit_trade_no"              ` // 交易号
	PaymentChannelId           uint        `json:"payment_channel_id"            ` // 支付渠道
	TradePaymentAmount         float64     `json:"trade_payment_amount"          ` // 实付金额:在线支付金额
	ReturnContactName          string      `json:"return_contact_name"           ` // 联系人
	ReturnStoreUserId          uint        `json:"return_store_user_id"          ` // 审核人员id
	ReturnWithdrawConfirm      uint        `json:"return_withdraw_confirm"       ` // 提现审核(BOOL):0-未审核; 1-已审核
	ReturnFinancialConfirm     uint        `json:"return_financial_confirm"      ` // 退款财务确认(BOOL):0-未确认; 1-已确认
	ReturnFinancialConfirmTime *gtime.Time `json:"return_financial_confirm_time" ` // 退款财务确认时间
	SubsiteId                  uint        `json:"subsite_id"                    ` // 所属分站:0-总站
}
