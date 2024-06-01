package cms

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type ArticleBaseAdd struct {
	ArticleId             int         `json:"article_id"              ` // 编号
	ArticleTitle          string      `json:"article_title"           ` // 标题
	ArticleName           string      `json:"article_name"            ` // 文章别名slug
	ArticleExcerpt        string      `json:"article_excerpt"         ` // 文章摘要
	ArticleContent        string      `json:"article_content"         ` // 文章内容(HTML)
	ArticleUrl            string      `json:"article_url"             ` // 调用网址:默认为本页面构造的网址
	CategoryId            uint        `json:"category_id"             ` // 所属分类
	ArticleTemplate       string      `json:"article_template"        ` // 模板
	ArticleSeoTitle       string      `json:"article_seo_title"       ` // SEO标题
	ArticleSeoKeywords    string      `json:"article_seo_keywords"    ` // SEO关键字
	ArticleSeoDescription string      `json:"article_seo_description" ` // SEO描述
	ArticleReplyFlag      uint        `json:"article_reply_flag"      ` // 是否启用问答留言(BOOL):0-否;1-是
	ArticleLang           string      `json:"article_lang"            ` // 语言
	ArticleType           uint        `json:"article_type"            ` // 文章类型(ENUM):1-文章;2-公告
	ArticleSort           uint        `json:"article_sort"            ` // 排序
	ArticleStatus         uint        `json:"article_status"          ` // 状态(BOOL):0-关闭;1-启用
	ArticleAddTime        *gtime.Time `json:"article_add_time"        ` // 添加世间
	ArticleImage          string      `json:"article_image"           ` // 文章图片
	UserId                uint        `json:"user_id"                 ` // 文章作者
	ArticleTags           string      `json:"article_tags"            ` // 文章标签(DOT):文章标签
	ArticleIsPopular      uint        `json:"article_is_popular"      ` // 是否热门
}

type ArticleBaseAddReq struct {
	g.Meta `path:"/manage/cms/articleBase/add" tags:"文章管理" method:"post" summary:"文章管理编辑接口"`

	ArticleBaseAdd
}

type ArticleBaseEditReq struct {
	g.Meta `path:"/manage/cms/articleBase/edit" tags:"文章管理" method:"post" summary:"文章管理编辑接口"`

	ArticleId uint `json:"article_id"  dc:"文章管理编号"   ` // 文章管理编号
	ArticleBaseAdd
}

type ArticleBaseEditRes struct {
	ArticleId uint `json:"article_id" dc:"文章管理编号"   ` // 文章管理编号
}

type ArticleBaseRemoveReq struct {
	g.Meta    `path:"/manage/cms/articleBase/remove" tags:"文章管理" method:"post" summary:"文章管理删除接口"`
	ArticleId uint `json:"article_id" dc:"文章管理编号"   ` // 文章管理编号
}

type ArticleBaseRemoveRes struct {
}
type ArticleBaseRemoveBatchReq struct {
	g.Meta    `path:"/manage/cms/articleBase/removeBatch" tags:"文章管理" method:"post" summary:"文章管理删除接口"`
	ArticleId string `json:"article_id" dc:"文章管理编号"   ` // 文章管理编号
}

type ArticleBaseRemoveBatchRes struct {
}

type ArticleBaseListReq struct {
	g.Meta `path:"/manage/cms/articleBase/list" tags:"文章管理" method:"get" summary:"文章管理列表接口"`
	ml.BaseList

	ArticleTitle string `json:"article_title"   type:"LIKE"                 ` // 文章标题
}

type ArticleBaseListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
