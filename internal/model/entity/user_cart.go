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

// UserCart is the golang structure for table user_cart.
type UserCart struct {
	CartId         uint64 `json:"cart_id"          ` // 编号
	UserId         uint   `json:"user_id"          ` // 买家编号
	StoreId        uint   `json:"store_id"         ` // 店铺编号
	ChainId        uint   `json:"chain_id"         ` // 门店编号
	ProductId      uint64 `json:"product_id"       ` // 产品编号
	ItemId         uint64 `json:"item_id"          ` // 商品编号
	CartQuantity   uint   `json:"cart_quantity"    ` // 购买商品数量
	CartType       uint   `json:"cart_type"        ` // 购买类型(ENUM):1-购买; 2-积分兑换; 3-赠品; 4-活动促销
	ActivityId     uint   `json:"activity_id"      ` // 活动Id-加价购等等加入购物的需要提示
	ActivityItemId uint64 `json:"activity_item_id" ` // 加入购物车所属活动Item编号
	CartSelect     bool   `json:"cart_select"      ` // 选中状态(BOOL):0-未选;1-已选
	CartTtl        uint64 `json:"cart_ttl"         ` // 有效时间戳
	CartTime       uint64 `json:"cart_time"        ` // 添加时间戳
	CartFile       string `json:"cart_file"        ` // 文件
	Version        uint   `json:"version"          ` // 版本
}
