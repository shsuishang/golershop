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

// OrderItem is the golang structure of table trade_order_item for DAO operations like Where/Data.
type OrderItem struct {
	g.Meta                       `orm:"table:trade_order_item, do:true"`
	OrderItemId                  interface{} // 编号
	OrderId                      interface{} // 订单编号
	UserId                       interface{} // 买家编号
	StoreId                      interface{} // 店铺编号
	ProductId                    interface{} // 产品编号
	ProductName                  interface{} // 商品名称
	ItemId                       interface{} // 货品编号
	ItemName                     interface{} // 商品名称
	CategoryId                   interface{} // 分类编号
	ItemCostPrice                interface{} // 成本价
	ItemUnitPrice                interface{} // 商品价格单价
	ItemUnitPoints               interface{} // 资源1单价
	ItemUnitSp                   interface{} // 资源2单价
	OrderItemSalePrice           interface{} // 商品实际成交价单价
	OrderItemQuantity            interface{} // 商品数量
	OrderItemInventoryLock       interface{} // 库存锁定(ENUM):1001-下单锁定;1002-支付锁定;
	OrderItemImage               interface{} // 商品图片
	OrderItemReturnNum           interface{} // 退货数量
	OrderItemReturnSubtotal      interface{} // 退款总额
	OrderItemReturnAgreeAmount   interface{} // 退款金额:同意额度
	OrderItemAmount              interface{} // 商品实际总金额: order_item_sale_price * order_item_quantity
	OrderItemDiscountAmount      interface{} // 优惠金额:只考虑单品的，订单及店铺总活动优惠不影响
	OrderItemAdjustFee           interface{} // 手工调整金额
	OrderItemPointsFee           interface{} // 积分费用
	OrderItemPointsAdd           interface{} // 赠送积分
	OrderItemPaymentAmount       interface{} // 实付金额: order_item_payment_amount =  order_item_amount - order_item_discount_amount - order_item_adjust_fee - order_item_point_fee
	OrderItemEvaluationStatus    interface{} // 评价状态(ENUM): 0-未评价;1-已评价;2-失效评价
	ActivityTypeId               interface{} // 活动类型(ENUM):0-默认;1101-加价购=搭配宝;1102-店铺满赠-小礼品;1103-限时折扣;1104-优惠套装;1105-店铺代金券coupon优惠券;1106-拼团;1107-满减送;1108-阶梯价;1109-积分换购
	ActivityId                   interface{} // 促销活动ID:与activity_type_id搭配使用, 团购ID/限时折扣ID/优惠套装ID/积分兑换编号
	ActivityCode                 interface{} // 礼包活动对应兑换码code
	OrderItemCommissionRate      interface{} // 分佣金比例百分比
	OrderItemCommissionFee       interface{} // 佣金
	OrderItemCommissionFeeRefund interface{} // 退款佣金
	PolicyDiscountrate           interface{} // 价格策略折扣率
	OrderItemVoucher             interface{} // 分配优惠券额度
	OrderItemReduce              interface{} // 分配满减额度
	OrderItemNote                interface{} // 备注
	OrderItemFile                interface{} // 订单附件
	OrderItemConfirmFile         interface{} // 商家附件
	OrderItemConfirmStatus       interface{} // 买家确认状态(BOOL):0-为确认;1-已确认
	OrderItemSalerId             interface{} // 单品分销者编号
	ItemSrcId                    interface{} // 分销商品编号
	OrderItemSupplierSync        interface{} // 拆单同步状态(BOOL):0-未同步;1-已同步
	SrcOrderId                   interface{} // 来源订单
	OrderItemReturnAgreeNum      interface{} // 同意退货数量
	OrderGiveId                  interface{} // 满返优惠券id
	Version                      interface{} // 版本
}

type OrderItemListInput struct {
	ml.BaseList
	Where OrderItem // 查询条件
}

type OrderItemListOutput struct {
	Items   []*entity.OrderItem // 列表
	Page    int                 // 分页号码
	Total   int                 // 总页数
	Records int                 // 数据总数
	Size    int                 // 单页数量
}

type OrderItemListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
