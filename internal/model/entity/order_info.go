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

// OrderInfo is the golang structure for table order_info.
type OrderInfo struct {
	OrderId                     string      `json:"order_id"                       ` // 订单编号
	OrderTitle                  string      `json:"order_title"                    ` // 订单标题
	StoreId                     uint        `json:"store_id"                       ` // 卖家店铺编号
	SubsiteId                   uint        `json:"subsite_id"                     ` // 所属分站:0-总站
	UserId                      uint        `json:"user_id"                        ` // 买家编号
	KindId                      uint        `json:"kind_id"                        ` // 订单种类(ENUM): 1201-实物 ; 1202-教育类 ; 1203-电子卡券  ; 1204-其它
	OrderLockStatus             bool        `json:"order_lock_status"              ` // 锁定状态(BOOL):0-正常;1-锁定,退款退货
	OrderIsSettlemented         bool        `json:"order_is_settlemented"          ` // 订单是否结算(BOOL):0-未结算; 1-已结算
	OrderSettlementTime         *gtime.Time `json:"order_settlement_time"          ` // 订单结算时间
	OrderBuyerEvaluationStatus  uint        `json:"order_buyer_evaluation_status"  ` // 买家针对订单对店铺评价(ENUM): 0-未评价;1-已评价;  2-已过期未评价
	OrderSellerEvaluationStatus uint        `json:"order_seller_evaluation_status" ` // 卖家评价状态(ENUM):0-未评价;1-已评价;  2-已过期未评价
	OrderBuyerHidden            bool        `json:"order_buyer_hidden"             ` // 买家删除(BOOL): 1-是; 0-否
	OrderShopHidden             bool        `json:"order_shop_hidden"              ` // 店铺删除(BOOL): 1-是; 0-否
	PaymentTypeId               uint        `json:"payment_type_id"                ` // 支付方式(ENUM): 1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付;
	PaymentTime                 *gtime.Time `json:"payment_time"                   ` // 付款时间
	OrderStateId                uint        `json:"order_state_id"                 ` // 订单状态(LIST):2011-待订单审核;2013-待财务审核;2020-待配货/待出库审核;2030-待发货;2040-已发货/待收货确认;2060-已完成/已签收;2070-已取消/已作废;
	OrderIsReview               bool        `json:"order_is_review"                ` // 订单审核(BOOL):0-未审核;1-已审核;
	OrderFinanceReview          bool        `json:"order_finance_review"           ` // 财务状态(BOOL):0-未审核;1-已审核
	OrderIsPaid                 uint        `json:"order_is_paid"                  ` // 付款状态(ENUM):3010-未付款;3011-付款待审核;3012-部分付款;3013-已付款
	OrderIsOut                  uint        `json:"order_is_out"                   ` // 出库状态(ENUM):3020-未出库;3021-部分出库通过拆单解决这种问题;3022-已出库
	OrderIsShipped              uint        `json:"order_is_shipped"               ` // 发货状态(ENUM):3030-未发货;3032-已发货;3031-部分发货
	OrderIsReceived             bool        `json:"order_is_received"              ` // 收货状态(BOOL):0-未收货;1-已收货
	OrderReceivedTime           *gtime.Time `json:"order_received_time"            ` // 订单签收时间
	ChainId                     uint        `json:"chain_id"                       ` // 门店编号
	DeliveryTypeId              uint        `json:"delivery_type_id"               ` // 配送方式
	OrderIsOffline              bool        `json:"order_is_offline"               ` // 线下订单(BOOL):0-线上;1-线下
	OrderExpressPrint           bool        `json:"order_express_print"            ` // 是否打印(BOOL):0-未打印;1-已打印
	ActivityId                  string      `json:"activity_id"                    ` // 活动编号(DOT)
	ActivityTypeId              string      `json:"activity_type_id"               ` // 活动类型(DOT)
	SalespersonId               uint        `json:"salesperson_id"                 ` // 销售员编号:用户编号
	OrderIsSync                 bool        `json:"order_is_sync"                  ` // 是否ERP同步(BOOL):0-未同步; 1-已同步
	StoreIsSelfsupport          bool        `json:"store_is_selfsupport"           ` // 是否自营(ENUM): 1-自营;0-非自营
	StoreType                   uint        `json:"store_type"                     ` // 店铺类型(ENUM): 1-卖家店铺; 2-供应商店铺
	OrderErpId                  string      `json:"order_erp_id"                   ` // ERP订单编号
	DistributorUserId           uint        `json:"distributor_user_id"            ` // 分销商编号:用户编号
	OrderIsCb                   bool        `json:"order_is_cb"                    ` // 跨境订单(BOOL):0-否; 1-是
	OrderIsCbSync               bool        `json:"order_is_cb_sync"               ` // 是否报关(BOOL):0-否; 1-是
	SrcOrderId                  string      `json:"src_order_id"                   ` // 来源订单
	OrderIsTransfer             bool        `json:"order_is_transfer"              ` // 是否代发(BOOL):0-否; 1-是
	OrderIsTransferNote         string      `json:"order_is_transfer_note"         ` // 转单执行结果
	OrderFxIsSettlemented       bool        `json:"order_fx_is_settlemented"       ` // 佣金是否发放(BOOL):0 -未发放;1 -已发放
	OrderFxSettlementTime       *gtime.Time `json:"order_fx_settlement_time"       ` // 佣金结算时间
	OrderType                   uint        `json:"order_type"                     ` // 订单类型(ENUM)
	OrderWithdrawConfirm        bool        `json:"order_withdraw_confirm"         ` // 提现审核(BOOL):0-未审核; 1-已审核
	PaymentFormId               uint        `json:"payment_form_id"                ` // 支付方式(BOOL):1-先预约后支付;0-先支付后预约
	CartTypeId                  uint        `json:"cart_type_id"                   ` // 购买类型(ENUM):1-购买; 2-积分兑换; 3-赠品; 4-活动促销
	CardKindId                  uint        `json:"card_kind_id"                   ` // 商品绑定卡片类型(ENUM): 1001-次卡商品; 1002-优惠券商品;1003-券码商品;
	CreateTime                  uint64      `json:"create_time"                    ` // 下单时间:检索使用
	UpdateTime                  uint64      `json:"update_time"                    ` // 当前状态的处理时间
	Version                     uint        `json:"version"                        ` // 乐观锁
}
