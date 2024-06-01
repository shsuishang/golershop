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

// OrderReturn is the golang structure of table trade_order_return for DAO operations like Where/Data.
type OrderReturn struct {
	g.Meta                     `orm:"table:trade_order_return, do:true"`
	ReturnId                   interface{} // 退单号
	ServiceTypeId              interface{} // 服务类型(ENUM):1-退款;2-退货;3-换货;4-维修
	OrderId                    interface{} // 订单编号
	ReturnRefundAmount         interface{} // 退款金额
	ReturnRefundPoint          interface{} // 积分部分
	StoreId                    interface{} // 店铺编号
	BuyerUserId                interface{} // 买家编号
	BuyerStoreId               interface{} // 买家是否有店铺
	ReturnAddTime              interface{} // 添加时间
	ReturnReasonId             interface{} // 退款理由编号
	ReturnBuyerMessage         interface{} // 买家退货备注
	ReturnAddrContacter        interface{} // 收货人
	ReturnTel                  interface{} // 联系电话
	ReturnAddr                 interface{} // 收货地址详情
	ReturnPostCode             interface{} // 邮编
	ExpressId                  interface{} // 物流公司编号
	ReturnTrackingName         interface{} // 物流名称
	ReturnTrackingNumber       interface{} // 物流单号
	PlantformReturnStateId     interface{} // 申请状态平台(ENUM):3180-未申请;3181-待处理;3182-为已完成
	ReturnStateId              interface{} // 卖家处理状态(ENUM): 3100-【客户】提交退单;3105-退单审核;3110-收货确认;3115-退款确认;3120-客户】收款确认;3125-完成
	ReturnIsPaid               interface{} // 退款完成
	ReturnIsShippingFee        interface{} // 退货类型(BOOL): 0-退款单;1-退运费单
	ReturnShippingFee          interface{} // 退运费额度
	ReturnFlag                 interface{} // 退货类型(ENUM): 0-不用退货;1-需要退货
	ReturnType                 interface{} // 申请类型(ENUM): 1-退款申请; 2-退货申请; 3-虚拟退款
	ReturnOrderLock            interface{} // 订单锁定类型(BOOL):1-不用锁定;2-需要锁定
	ReturnItemStateId          interface{} // 物流状态(LIST):2030-待发货;2040-已发货/待收货确认;2060-已完成/已签收;2070-已取消/已作废;
	ReturnStoreTime            *gtime.Time // 商家处理时间
	ReturnStoreMessage         interface{} // 商家备注
	ReturnCommisionFee         interface{} // 退还佣金
	ReturnFinishTime           *gtime.Time // 退款完成时间
	ReturnPlatformMessage      interface{} // 平台留言
	ReturnIsSettlemented       interface{} // 订单是否结算(BOOL): 0-未结算; 1-已结算
	ReturnSettlementTime       *gtime.Time // 订单结算时间
	ReturnChannelCode          interface{} // 退款渠道(ENUM):money-余额;alipay-支付宝;wx_native-微信
	ReturnChannelFlag          interface{} // 渠道是否退款(ENUM): 0-待退; 1-已退; 2-异常
	ReturnChannelTime          *gtime.Time // 渠道退款时间
	ReturnChannelTransId       interface{} // 渠道退款单号
	DepositTradeNo             interface{} // 交易号
	PaymentChannelId           interface{} // 支付渠道
	TradePaymentAmount         interface{} // 实付金额:在线支付金额
	ReturnContactName          interface{} // 联系人
	ReturnStoreUserId          interface{} // 审核人员id
	ReturnWithdrawConfirm      interface{} // 提现审核(BOOL):0-未审核; 1-已审核
	ReturnFinancialConfirm     interface{} // 退款财务确认(BOOL):0-未确认; 1-已确认
	ReturnFinancialConfirmTime *gtime.Time // 退款财务确认时间
	SubsiteId                  interface{} // 所属分站:0-总站
}

type OrderReturnListInput struct {
	ml.BaseList
	Where OrderReturn // 查询条件
}

type OrderReturnListOutput struct {
	Items   []*entity.OrderReturn // 列表
	Page    int                   // 分页号码
	Total   int                   // 总页数
	Records int                   // 数据总数
	Size    int                   // 单页数量
}

type OrderReturnListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
