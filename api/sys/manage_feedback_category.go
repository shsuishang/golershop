package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type FeedbackCategoryAdd struct {
	FeedbackCategoryName   string `json:"feedback_category_name"   ` // 分类名称
	FeedbackTypeId         uint   `json:"feedback_type_id"         ` // 反馈分类
	FeedbackCategoryEnable bool   `json:"feedback_category_enable" ` // 是否启用(BOOL):0-禁用;1-启用
}

type FeedbackCategoryEditReq struct {
	g.Meta `path:"/manage/sys/feedbackCategory/edit" tags:"反馈分类" method:"post" summary:"反馈分类编辑接口"`

	FeedbackCategoryId uint `json:"feedback_category_id"     ` // 分类编号
	FeedbackCategoryAdd
}

type FeedbackCategoryEditRes struct {
	FeedbackCategoryId uint `json:"feedback_category_id"     ` // 类型编号
}

type FeedbackCategoryAddReq struct {
	g.Meta `path:"/manage/sys/feedbackCategory/add" tags:"反馈分类" method:"post" summary:"反馈分类编辑接口"`

	FeedbackCategoryAdd
}

type FeedbackCategoryRemoveReq struct {
	g.Meta `path:"/manage/sys/feedbackCategory/remove" tags:"反馈分类" method:"post" summary:"反馈分类删除接口"`

	FeedbackCategoryId string `json:"feedback_category_id"     ` // 分类编号
}

type FeedbackCategoryRemoveRes struct {
}

type FeedbackCategoryListReq struct {
	g.Meta `path:"/manage/sys/feedbackCategory/list" tags:"反馈分类" method:"get" summary:"反馈分类列表接口"`
	ml.BaseList

	FeedbackCategoryName string `json:"feedback_category_name"   ` // 反馈分类
}

type FeedbackCategoryListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type FeedbackCategoryEditStateReq struct {
	g.Meta `path:"/manage/sys/feedbackCategory/editState" tags:"反馈分类" method:"post" summary:"反馈分类状态编辑接口"`

	FeedbackCategoryId     uint `json:"feedback_category_id"     ` // 分类编号
	FeedbackCategoryEnable bool `json:"feedback_category_enable" ` // 是否启用(BOOL):0-禁用;1-启用
}

type FeedbackCategoryEditStateRes struct {
	FeedbackCategoryId uint `json:"feedback_category_id"     ` // 分类编号
}
