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

// DistributionOrderItem is the golang structure for table distribution_order_item.
type DistributionOrderItem struct {
	UoiId                     uint    `json:"uoi_id"                      ` // 订单收益编号
	UserId                    uint    `json:"user_id"                     ` // 用户编号
	OrderId                   string  `json:"order_id"                    ` // 订单编号
	ProductId                 uint64  `json:"product_id"                  ` // 商品编号
	ItemId                    uint64  `json:"item_id"                     ` // 商品SKU
	UoiBuyCommission          float64 `json:"uoi_buy_commission"          ` // 推广员佣金
	UoiDirectsellerCommission float64 `json:"uoi_directseller_commission" ` // 销售员佣金
	UoiDistributorCommission  float64 `json:"uoi_distributor_commission"  ` // 分销商收益-本店销售获取差价
	BuyerUserId               uint    `json:"buyer_user_id"               ` // 买家编号
	UoiLevel                  uint    `json:"uoi_level"                   ` // 等级
	UoiTime                   uint64  `json:"uoi_time"                    ` // 时间
	UoiActive                 bool    `json:"uoi_active"                  ` // 是否有效(BOOL):0-未生效;1-有效
	StoreId                   uint    `json:"store_id"                    ` // 店铺编号
	UoiIsPaid                 bool    `json:"uoi_is_paid"                 ` // 是否有效(BOOL):0-未支付;1-已支付
	UoiReceivetime            int64   `json:"uoi_receivetime"             ` // 收货时间
}
