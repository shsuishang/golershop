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

// OrderBase is the golang structure for table order_base.
type OrderBase struct {
	OrderId            string      `json:"order_id"             ` // 订单编号
	OrderNumber        string      `json:"order_number"         ` // 订单编号
	OrderTime          *gtime.Time `json:"order_time"           ` // 下单时间
	OrderProductAmount float64     `json:"order_product_amount" ` // 商品原价总和:商品发布原价
	OrderPaymentAmount float64     `json:"order_payment_amount" ` // 应付金额:order_product_amount - order_discount_amount + order_shipping_fee - order_voucher_price - order_points_fee - order_adjust_fee
	CurrencyId         uint        `json:"currency_id"          ` // 货币编号
	CurrencySymbolLeft string      `json:"currency_symbol_left" ` // 左符号
	StoreId            uint        `json:"store_id"             ` // 店铺编号
	StoreName          string      `json:"store_name"           ` // 店铺名称
	UserId             uint        `json:"user_id"              ` // 买家编号
	UserNickname       string      `json:"user_nickname"        ` // 买家昵称
	OrderStateId       uint        `json:"order_state_id"       ` // 订单状态(LIST):2011-待订单审核;2013-待财务审核;2020-待配货/待出库审核;2030-待发货;2040-已发货/待收货确认;2060-已完成/已签收;2070-已取消/已作废;
}
