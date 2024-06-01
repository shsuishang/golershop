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

// PageBase is the golang structure of table sys_page_base for DAO operations like Where/Data.
type PageBase struct {
	g.Meta         `orm:"table:sys_page_base, do:true"`
	PageId         interface{} // 页面编号
	PageName       interface{} // 页面名称
	StoreId        interface{} // 所属店铺
	UserId         interface{} // 所属用户
	SubsiteId      interface{} // 所属分站:0-总站
	PageBuildin    interface{} // 是否内置(BOOL):0-否;1-是
	PageType       interface{} // 类型(ENUM):1-WAP;2-PC;3-APP
	PageTpl        interface{} // 页面布局模板
	AppId          interface{} // 所属APP
	PageCode       interface{} // 页面代码
	PageNav        interface{} // 导航数据
	PageConfig     interface{} // 页面配置
	PageShareTitle interface{} // 分享标题
	PageShareImage interface{} // 分享图片
	PageQrcode     interface{} // 分享二维码
	PageIndex      interface{} // 是否首页(BOOL):0-非首页;1-首页
	PageGb         interface{} // 拼团首页(BOOL):0-非首页;1-首页
	PageActivity   interface{} // 活动首页(BOOL):0-非首页;1-首页
	PagePoint      interface{} // 积分首页(BOOL):0-非首页;1-首页
	PageGbs        interface{} // 团购首页(BOOL):0-非首页;1-首页
	PagePackage    interface{} // 组合套餐(BOOL):0-非首页;1-首页
	PagePfgb       interface{} // 批发团购首页(BOOL):0-非首页;1-首页
	PageSns        interface{} // 社区(BOOL):0-非首页;1-首页
	PageArticle    interface{} // 资讯(BOOL):0-非首页;1-首页
	PageZerobuy    interface{} // 零元购区(BOOL):0-否;1-是
	PageHigharea   interface{} // 高额返区(BOOL):0-否;1-是
	PageTaday      interface{} // 今日爆款(BOOL):0-否;1-是
	PageEveryday   interface{} // 每日好店(BOOL):0-否;1-是
	PageSecondkill interface{} // 整点秒杀(BOOL):0-否;1-是
	PageSecondday  interface{} // 天天秒淘(BOOL):0-否;1-是
	PageRura       interface{} // 设置土特产(BOOL):0-否;1-是
	PageLikeyou    interface{} // 用户页banner(BOOL):0-否;1-是
	PageExchange   interface{} // 兑换专区(BOOL):0-否;1-是
	PageNew        interface{} // 新品首发(BOOL):0-否;1-是
	PageNewperson  interface{} // 新人优惠(BOOL):0-否;1-是
	PageUpgrade    interface{} // 升级VIP(BOOL):0-否;1-是
	PageMessage    interface{} // 信息发布(BOOL):0-否;1-是
	PageRelease    interface{} // 是否发布(BOOL):0-否;1-是
}

type PageBaseListInput struct {
	ml.BaseList
	Where PageBase // 查询条件
}

type PageBaseListOutput struct {
	Items   []*entity.PageBase // 列表
	Page    int                // 分页号码
	Total   int                // 总页数
	Records int                // 数据总数
	Size    int                // 单页数量
}

type PageBaseListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
