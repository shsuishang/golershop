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

// ActivityItem is the golang structure for table activity_item.
type ActivityItem struct {
	ActivityItemId          uint64  `json:"activity_item_id"           ` // 商品表编号
	StoreId                 uint    `json:"store_id"                   ` // 店铺编号
	ActivityTypeId          uint    `json:"activity_type_id"           ` // 活动类型编号
	ActivityId              uint    `json:"activity_id"                ` // 活动编号
	ProductId               uint64  `json:"product_id"                 ` // 产品编号
	ItemId                  uint64  `json:"item_id"                    ` // 商品编号
	CategoryId              uint    `json:"category_id"                ` // 商品分类
	ActivityItemStarttime   uint64  `json:"activity_item_starttime"    ` // 开始时间
	ActivityItemEndtime     uint64  `json:"activity_item_endtime"      ` // 结束时间
	ActivityItemPrice       float64 `json:"activity_item_price"        ` // 活动价格
	ActivityItemMinQuantity uint    `json:"activity_item_min_quantity" ` // 购买下限
	ActivityItemState       uint    `json:"activity_item_state"        ` // 活动状态(ENUM):0-未开启;1-正常;2-已结束;3-管理员关闭
	ActivityItemRecommend   bool    `json:"activity_item_recommend"    ` // 推荐标志(BOOL):0-未推荐;1-已推荐
}
