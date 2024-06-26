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

// PagePcNav is the golang structure for table page_pc_nav.
type PagePcNav struct {
	NavId           uint   `json:"nav_id"            ` // 导航编号
	NavType         uint   `json:"nav_type"          ` // 类别(ENUM):0-自定义导航;1-商品分类;2-文章导航;3-活动导航
	NavItemId       uint   `json:"nav_item_id"       ` // 类别内容编号
	NavTitle        string `json:"nav_title"         ` // 导航标题
	NavUrl          string `json:"nav_url"           ` // 导航链接(HTML)
	NavPosition     uint   `json:"nav_position"      ` // 导航位置(ENUM):0-头部;1-中部;2-底部
	NavTargetBlank  bool   `json:"nav_target_blank"  ` // 是否以新窗口打开(BOOL):1-是; 0-否
	NavImage        string `json:"nav_image"         ` // 导航图片
	NavDropdownMenu string `json:"nav_dropdown_menu" ` // 导航下拉内容(HTML)
	NavOrder        uint   `json:"nav_order"         ` // 排序
	NavEnable       bool   `json:"nav_enable"        ` // 是否启用(BOOL):1-是; 0-否
	NavBuildin      uint   `json:"nav_buildin"       ` // 系统内置(ENUM):1-是; 0-否
}
