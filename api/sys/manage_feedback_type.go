package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type FeedbackTypeAdd struct {
	FeedbackTypeName   string `json:"feedback_type_name"   ` // 类型名称
	FeedbackTypeEnable bool   `json:"feedback_type_enable" ` // 是否启用(BOOL):0-禁用;1-启用
	FeedbackTypeGenus  uint   `json:"feedback_type_genus"  ` // 所属身份(ENUM):1-会员所属;2-经销商所属
}

type FeedbackTypeEditReq struct {
	g.Meta `path:"/manage/sys/feedbackType/edit" tags:"反馈类型" method:"post" summary:"反馈类型编辑接口"`

	FeedbackTypeId uint `json:"feedback_type_id"     ` // 类型编号
	FeedbackTypeAdd
}

type FeedbackTypeEditRes struct {
	FeedbackTypeId uint `json:"feedback_type_id"     ` // 类型编号
}

type FeedbackTypeAddReq struct {
	g.Meta `path:"/manage/sys/feedbackType/add" tags:"反馈类型" method:"post" summary:"反馈类型编辑接口"`

	FeedbackTypeAdd
}

type FeedbackTypeRemoveReq struct {
	g.Meta `path:"/manage/sys/feedbackType/remove" tags:"反馈类型" method:"post" summary:"反馈类型删除接口"`

	FeedbackTypeId string `json:"feedback_type_id"     ` // 类型编号
}

type FeedbackTypeRemoveRes struct {
}

type FeedbackTypeListReq struct {
	g.Meta `path:"/manage/sys/feedbackType/list" tags:"反馈类型" method:"get" summary:"反馈类型列表接口"`
	ml.BaseList

	FeedbackTypeName string `json:"feedback_type_name"   ` // 类型名称
}

type FeedbackTypeListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type FeedbackTypeEditStateReq struct {
	g.Meta `path:"/manage/sys/feedbackType/editState" tags:"反馈类型" method:"post" summary:"反馈类型状态编辑接口"`

	FeedbackTypeId     uint `json:"feedback_type_id"     ` // 类型编号
	FeedbackTypeEnable bool `json:"feedback_type_enable" ` // 是否启用(BOOL):0-禁用;1-启用
}

type FeedbackTypeEditStateRes struct {
	FeedbackTypeId uint `json:"feedback_type_id"     ` // 类型编号
}
