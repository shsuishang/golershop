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

// PageCategoryNav is the golang structure for table page_category_nav.
type PageCategoryNav struct {
	CategoryNavId     uint   `json:"category_nav_id"     ` // 编号
	CategoryNavName   string `json:"category_nav_name"   ` // 分类名称
	CategoryNavImage  string `json:"category_nav_image"  ` // 分类图片
	CategoryIds       string `json:"category_ids"        ` // 推荐分类(DOT)
	ItemIds           string `json:"item_ids"            ` // 推荐商品(DOT)
	BrandIds          string `json:"brand_ids"           ` // 推荐品牌(DOT)
	CategoryNavAdv    string `json:"category_nav_adv"    ` // 广告数据(JSON)
	CategoryNavType   uint   `json:"category_nav_type"   ` // 模板分类(ENUM):1-分类模板1;2-商品模板
	CategoryNavOrder  uint   `json:"category_nav_order"  ` // 排序
	CategoryNavEnable bool   `json:"category_nav_enable" ` // 是否启用(BOOL):0-不显示;1-显示
}
