package cms

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type ArticleCommentAdd struct {
	CommentId         uint64      `json:"comment_id"          ` // 评论编号
	ArticleId         uint64      `json:"article_id"          ` // 文章编号
	CommentUserIp     string      `json:"comment_user_ip"     ` // 评论IP
	CommentTime       *gtime.Time `json:"comment_time"        ` // 评论时间
	CommentContent    string      `json:"comment_content"     ` // 评论内容
	CommentKarma      uint        `json:"comment_karma"       ` // 评论karma值
	CommentApproved   string      `json:"comment_approved"    ` // 评论许可(ENUM):0-0;1-1;spam-spam'
	CommentAgent      string      `json:"comment_agent"       ` // 评论代理:（浏览器，操作系统等）
	CommentType       string      `json:"comment_type"        ` // 评论类型:（pingback|trackback)
	CommentParentId   uint64      `json:"comment_parent_id"   ` // 上级编号
	UserId            uint        `json:"user_id"             ` // 评论者编号
	CommentHelpfulNum uint        `json:"comment_helpful_num" ` // 有帮助数
	CommentIsShow     bool        `json:"comment_is_show"     ` // 是否显示(BOOL):1-显示;0-不显示
}

type ArticleCommentAddReq struct {
	g.Meta `path:"/manage/cms/articleComment/add" tags:"文章评论" method:"post" summary:"文章评论编辑接口"`

	ArticleCommentAdd
}

type ArticleCommentEditReq struct {
	g.Meta `path:"/manage/cms/articleComment/edit" tags:"文章评论" method:"post" summary:"文章评论编辑接口"`

	CommentId uint `json:"article_id"  dc:"文章评论编号"   ` // 文章评论编号
	ArticleCommentAdd
}

type ArticleCommentEditRes struct {
	CommentId uint `json:"article_id" dc:"文章评论编号"   ` // 文章评论编号
}

type ArticleCommentRemoveReq struct {
	g.Meta    `path:"/manage/cms/articleComment/remove" tags:"文章评论" method:"post" summary:"文章评论删除接口"`
	CommentId uint `json:"article_id" dc:"文章评论编号"   ` // 文章评论编号
}

type ArticleCommentRemoveRes struct {
}

type ArticleCommentListReq struct {
	g.Meta `path:"/manage/cms/articleComment/list" tags:"文章评论" method:"get" summary:"文章评论列表接口"`
	ml.BaseList

	CommentId uint `json:"article_id"                   ` // 用户编号
}

type ArticleCommentListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
