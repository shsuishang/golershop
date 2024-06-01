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

// OrderItem is the golang structure for table order_item.
type OrderItem struct {
	OrderItemId                  uint64  `json:"order_item_id"                    `      // 编号
	OrderId                      string  `json:"order_id"                         `      // 订单编号
	UserId                       uint    `json:"user_id"                          `      // 买家编号
	StoreId                      uint    `json:"store_id"                         `      // 店铺编号
	ProductId                    uint64  `json:"product_id"                       `      // 产品编号
	ProductName                  string  `json:"product_name"                     `      // 商品名称
	ItemId                       uint64  `json:"item_id"                          `      // 货品编号
	ItemName                     string  `json:"item_name"                        `      // 商品名称
	CategoryId                   uint    `json:"category_id"                      `      // 分类编号
	ItemCostPrice                float64 `json:"item_cost_price"                  `      // 成本价
	ItemUnitPrice                float64 `json:"item_unit_price"                  `      // 商品价格单价
	ItemUnitPoints               float64 `json:"item_unit_points"                 `      // 资源1单价
	ItemUnitSp                   float64 `json:"item_unit_sp"                     `      // 资源2单价
	OrderItemSalePrice           float64 `json:"order_item_sale_price"            `      // 商品实际成交价单价
	OrderItemQuantity            uint    `json:"order_item_quantity"              `      // 商品数量
	OrderItemInventoryLock       uint    `json:"order_item_inventory_lock"        `      // 库存锁定(ENUM):1001-下单锁定;1002-支付锁定;
	OrderItemImage               string  `json:"order_item_image"                 `      // 商品图片
	OrderItemReturnNum           uint    `json:"order_item_return_num"            `      // 退货数量
	OrderItemReturnSubtotal      float64 `json:"order_item_return_subtotal"       `      // 退款总额
	OrderItemReturnAgreeAmount   float64 `json:"order_item_return_agree_amount"   `      // 退款金额:同意额度
	OrderItemAmount              float64 `json:"order_item_amount"                `      // 商品实际总金额: order_item_sale_price * order_item_quantity
	OrderItemDiscountAmount      float64 `json:"order_item_discount_amount"       `      // 优惠金额:只考虑单品的，订单及店铺总活动优惠不影响
	OrderItemAdjustFee           float64 `json:"order_item_adjust_fee"            `      // 手工调整金额
	OrderItemPointsFee           float64 `json:"order_item_points_fee"            `      // 积分费用
	OrderItemPointsAdd           float64 `json:"order_item_points_add"            `      // 赠送积分
	OrderItemPaymentAmount       float64 `json:"order_item_payment_amount"        `      // 实付金额: order_item_payment_amount =  order_item_amount - order_item_discount_amount - order_item_adjust_fee - order_item_point_fee
	OrderItemCanRefundAmount     float64 `json:"order_item_can_refund_amount" gorm:"-" ` // 最终在线支付金额(退款金额):order_item_payment_amount - order_item_voucher - order_item_reduce
	OrderItemEvaluationStatus    uint    `json:"order_item_evaluation_status"     `      // 评价状态(ENUM): 0-未评价;1-已评价;2-失效评价
	ActivityTypeId               uint    `json:"activity_type_id"                 `      // 活动类型(ENUM):0-默认;1101-加价购=搭配宝;1102-店铺满赠-小礼品;1103-限时折扣;1104-优惠套装;1105-店铺代金券coupon优惠券;1106-拼团;1107-满减送;1108-阶梯价;1109-积分换购
	ActivityId                   uint    `json:"activity_id"                      `      // 促销活动ID:与activity_type_id搭配使用, 团购ID/限时折扣ID/优惠套装ID/积分兑换编号
	ActivityCode                 string  `json:"activity_code"                    `      // 礼包活动对应兑换码code
	OrderItemCommissionRate      float64 `json:"order_item_commission_rate"       `      // 分佣金比例百分比
	OrderItemCommissionFee       float64 `json:"order_item_commission_fee"        `      // 佣金
	OrderItemCommissionFeeRefund float64 `json:"order_item_commission_fee_refund" `      // 退款佣金
	PolicyDiscountrate           float64 `json:"policy_discountrate"              `      // 价格策略折扣率
	OrderItemVoucher             float64 `json:"order_item_voucher"               `      // 分配优惠券额度
	OrderItemReduce              float64 `json:"order_item_reduce"                `      // 分配满减额度
	OrderItemNote                string  `json:"order_item_note"                  `      // 备注
	OrderItemFile                string  `json:"order_item_file"                  `      // 订单附件
	OrderItemConfirmFile         string  `json:"order_item_confirm_file"          `      // 商家附件
	OrderItemConfirmStatus       bool    `json:"order_item_confirm_status"        `      // 买家确认状态(BOOL):0-为确认;1-已确认
	OrderItemSalerId             uint    `json:"order_item_saler_id"              `      // 单品分销者编号
	ItemSrcId                    int64   `json:"item_src_id"                      `      // 分销商品编号
	OrderItemSupplierSync        bool    `json:"order_item_supplier_sync"         `      // 拆单同步状态(BOOL):0-未同步;1-已同步
	SrcOrderId                   string  `json:"src_order_id"                     `      // 来源订单
	OrderItemReturnAgreeNum      uint    `json:"order_item_return_agree_num"      `      // 同意退货数量
	OrderGiveId                  uint    `json:"order_give_id"                    `      // 满返优惠券id
	Version                      uint    `json:"version"                          `      // 版本

	ItemSavePrice   float64  `json:"item_save_price" gorm:"-"`   // 节省单价
	IfReturn        bool     `json:"if_return" gorm:"-"`         // 是否可退货
	ReturnIds       []string `json:"return_ids" gorm:"-"`        // 服务编号
	ProductItemName string   `json:"product_item_name" gorm:"-"` // Spec名称
}
