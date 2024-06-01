package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type FeedbackBaseAdd struct {
	FeedbackCategoryId         uint        `json:"feedback_category_id"          ` // 分类编号
	UserId                     uint        `json:"user_id"                       ` // 用户编号
	UserNickname               string      `json:"user_nickname"                 ` // 用户昵称
	FeedbackQuestion           string      `json:"feedback_question"             ` // 反馈问题:在这里描述您遇到的问题
	FeedbackQuestionUrl        string      `json:"feedback_question_url"         ` // 页面链接
	FeedbackQuestionAnswer     string      `json:"feedback_question_answer"      ` // 反馈问题
	FeedbackQuestionTime       *gtime.Time `json:"feedback_question_time"        ` // 反馈时间
	FeedbackQuestionAnswerTime *gtime.Time `json:"feedback_question_answer_time" ` // 回复时间
	FeedbackQuestionStatus     bool        `json:"feedback_question_status"      ` // 举报状态(BOOL):0-未处理;1-已处理
	FeedbackQuestionResult     uint        `json:"feedback_question_result"      ` // 举报处理结果(ENUM):1-无效举报;2-恶意举报;3-有效举报
	ItemId                     uint64      `json:"item_id"                       ` // 产品编号
	AdminId                    uint        `json:"admin_id"                      ` // 回复人员
}

type FeedbackBaseEditReq struct {
	g.Meta `path:"/manage/sys/feedbackBase/edit" tags:"反馈" method:"post" summary:"反馈编辑接口"`

	FeedbackId uint `json:"feedback_id"                   ` // 反馈编号
	FeedbackBaseAdd
}

type FeedbackBaseEditRes struct {
	FeedbackId interface{} `json:"feedback_id"   dc:"反馈编号"`
}

type FeedbackBaseAddReq struct {
	g.Meta `path:"/manage/sys/feedbackBase/add" tags:"反馈" method:"post" summary:"反馈编辑接口"`

	FeedbackBaseAdd
}

type FeedbackBaseRemoveReq struct {
	g.Meta `path:"/manage/sys/feedbackBase/remove" tags:"反馈" method:"post" summary:"反馈删除接口"`

	FeedbackId string `json:"feedback_id"                   ` // 反馈编号
}

type FeedbackBaseRemoveRes struct {
}

type FeedbackBaseListReq struct {
	g.Meta `path:"/manage/sys/feedbackBase/list" tags:"反馈" method:"get" summary:"反馈列表接口"`
	ml.BaseList

	FeedbackBaseName string `json:"feedback_type_name"   ` // 反馈
}

type FeedbackBaseListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type FeedbackBaseEditStateReq struct {
	g.Meta `path:"/manage/sys/feedbackBase/editState" tags:"反馈" method:"post" summary:"反馈状态编辑接口"`

	FeedbackId         uint `json:"feedback_id"                   ` // 反馈编号
	FeedbackBaseEnable bool `json:"feedback_category_enable" `      // 是否启用(BOOL):0-禁用;1-启用
}

type FeedbackBaseEditStateRes struct {
	FeedbackId interface{} `json:"feedback_id"                   ` // 反馈编号
}
