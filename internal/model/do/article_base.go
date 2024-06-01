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
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// ArticleBase is the golang structure of table cms_article_base for DAO operations like Where/Data.
type ArticleBase struct {
	g.Meta                `orm:"table:cms_article_base, do:true"`
	ArticleId             interface{} // 文章编号
	ArticleTitle          interface{} // 文章标题
	ArticleName           interface{} // 文章别名
	ArticleExcerpt        interface{} // 文章摘要
	ArticleContent        interface{} // 文章内容(HTML)
	ArticleUrl            interface{} // 调用网址:默认为本页面构造的网址
	CategoryId            interface{} // 所属分类
	ArticleTemplate       interface{} // 文章模板
	ArticleSeoTitle       interface{} // SEO标题
	ArticleSeoKeywords    interface{} // SEO关键字
	ArticleSeoDescription interface{} // SEO描述
	ArticleReplyFlag      interface{} // 启用问答(BOOL):0-否;1-是
	ArticleLang           interface{} // 文章语言
	ArticleType           interface{} // 文章类型(ENUM):1-文章;2-公告
	ArticleSort           interface{} // 文章排序
	ArticleStatus         interface{} // 文章状态(BOOL):0-关闭;1-启用
	ArticleAddTime        *gtime.Time // 添加时间
	ArticleImage          interface{} // 文章图片
	UserId                interface{} // 文章作者
	ArticleTags           interface{} // 文章标签(DOT):文章标签
	ArticleIsPopular      interface{} // 是否热门(BOOL):0-否;1-是
}

type ArticleBaseListInput struct {
	ml.BaseList
	Where ArticleBase // 查询条件
}

type ArticleBaseListOutput struct {
	Items   []*entity.ArticleBase // 列表
	Page    int                   // 分页号码
	Total   int                   // 总页数
	Records int                   // 数据总数
	Size    int                   // 单页数量
}

type ArticleBaseListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
