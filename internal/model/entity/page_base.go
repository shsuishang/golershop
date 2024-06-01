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

// PageBase is the golang structure for table page_base.
type PageBase struct {
	PageId         uint64 `json:"page_id"          ` // 页面编号
	PageName       string `json:"page_name"        ` // 页面名称
	StoreId        uint   `json:"store_id"         ` // 所属店铺
	UserId         uint   `json:"user_id"          ` // 所属用户
	SubsiteId      uint   `json:"subsite_id"       ` // 所属分站:0-总站
	PageBuildin    bool   `json:"page_buildin"     ` // 是否内置(BOOL):0-否;1-是
	PageType       uint   `json:"page_type"        ` // 类型(ENUM):1-WAP;2-PC;3-APP
	PageTpl        uint   `json:"page_tpl"         ` // 页面布局模板
	AppId          uint   `json:"app_id"           ` // 所属APP
	PageCode       string `json:"page_code"        ` // 页面代码
	PageNav        string `json:"page_nav"         ` // 导航数据
	PageConfig     string `json:"page_config"      ` // 页面配置
	PageShareTitle string `json:"page_share_title" ` // 分享标题
	PageShareImage string `json:"page_share_image" ` // 分享图片
	PageQrcode     string `json:"page_qrcode"      ` // 分享二维码
	PageIndex      bool   `json:"page_index"       ` // 是否首页(BOOL):0-非首页;1-首页
	PageGb         bool   `json:"page_gb"          ` // 拼团首页(BOOL):0-非首页;1-首页
	PageActivity   bool   `json:"page_activity"    ` // 活动首页(BOOL):0-非首页;1-首页
	PagePoint      bool   `json:"page_point"       ` // 积分首页(BOOL):0-非首页;1-首页
	PageGbs        bool   `json:"page_gbs"         ` // 团购首页(BOOL):0-非首页;1-首页
	PagePackage    bool   `json:"page_package"     ` // 组合套餐(BOOL):0-非首页;1-首页
	PagePfgb       bool   `json:"page_pfgb"        ` // 批发团购首页(BOOL):0-非首页;1-首页
	PageSns        bool   `json:"page_sns"         ` // 社区(BOOL):0-非首页;1-首页
	PageArticle    bool   `json:"page_article"     ` // 资讯(BOOL):0-非首页;1-首页
	PageZerobuy    bool   `json:"page_zerobuy"     ` // 零元购区(BOOL):0-否;1-是
	PageHigharea   bool   `json:"page_higharea"    ` // 高额返区(BOOL):0-否;1-是
	PageTaday      bool   `json:"page_taday"       ` // 今日爆款(BOOL):0-否;1-是
	PageEveryday   bool   `json:"page_everyday"    ` // 每日好店(BOOL):0-否;1-是
	PageSecondkill bool   `json:"page_secondkill"  ` // 整点秒杀(BOOL):0-否;1-是
	PageSecondday  bool   `json:"page_secondday"   ` // 天天秒淘(BOOL):0-否;1-是
	PageRura       bool   `json:"page_rura"        ` // 设置土特产(BOOL):0-否;1-是
	PageLikeyou    bool   `json:"page_likeyou"     ` // 用户页banner(BOOL):0-否;1-是
	PageExchange   bool   `json:"page_exchange"    ` // 兑换专区(BOOL):0-否;1-是
	PageNew        bool   `json:"page_new"         ` // 新品首发(BOOL):0-否;1-是
	PageNewperson  bool   `json:"page_newperson"   ` // 新人优惠(BOOL):0-否;1-是
	PageUpgrade    bool   `json:"page_upgrade"     ` // 升级VIP(BOOL):0-否;1-是
	PageMessage    bool   `json:"page_message"     ` // 信息发布(BOOL):0-否;1-是
	PageRelease    bool   `json:"page_release"     ` // 是否发布(BOOL):0-否;1-是
}
