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

// ArticleBase is the golang structure for table article_base.
type ArticleBase struct {
	ArticleId             int         `json:"article_id"              ` // 文章编号
	ArticleTitle          string      `json:"article_title"           ` // 文章标题
	ArticleName           string      `json:"article_name"            ` // 文章别名
	ArticleExcerpt        string      `json:"article_excerpt"         ` // 文章摘要
	ArticleContent        string      `json:"article_content"         ` // 文章内容(HTML)
	ArticleUrl            string      `json:"article_url"             ` // 调用网址:默认为本页面构造的网址
	CategoryId            uint        `json:"category_id"             ` // 所属分类
	ArticleTemplate       string      `json:"article_template"        ` // 文章模板
	ArticleSeoTitle       string      `json:"article_seo_title"       ` // SEO标题
	ArticleSeoKeywords    string      `json:"article_seo_keywords"    ` // SEO关键字
	ArticleSeoDescription string      `json:"article_seo_description" ` // SEO描述
	ArticleReplyFlag      bool        `json:"article_reply_flag"      ` // 启用问答(BOOL):0-否;1-是
	ArticleLang           string      `json:"article_lang"            ` // 文章语言
	ArticleType           uint        `json:"article_type"            ` // 文章类型(ENUM):1-文章;2-公告
	ArticleSort           uint        `json:"article_sort"            ` // 文章排序
	ArticleStatus         bool        `json:"article_status"          ` // 文章状态(BOOL):0-关闭;1-启用
	ArticleAddTime        *gtime.Time `json:"article_add_time"        ` // 添加时间
	ArticleImage          string      `json:"article_image"           ` // 文章图片
	UserId                uint        `json:"user_id"                 ` // 文章作者
	ArticleTags           string      `json:"article_tags"            ` // 文章标签(DOT):文章标签
	ArticleIsPopular      bool        `json:"article_is_popular"      ` // 是否热门(BOOL):0-否;1-是
}
