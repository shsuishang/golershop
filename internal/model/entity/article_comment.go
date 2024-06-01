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

// ArticleComment is the golang structure for table article_comment.
type ArticleComment struct {
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
