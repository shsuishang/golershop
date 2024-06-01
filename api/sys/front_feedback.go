package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/model/entity"
)

type FeedbackReq struct {
	g.Meta `path:"/front/sys/feedback/getCategory" tags:"平台反馈-举报" method:"get" summary:"平台反馈-举报"`
}

type FeedbackRes struct {
	*entity.FeedbackType
	Rows interface{} `json:"rows"`
}

type BaseListReq struct {
	g.Meta `path:"/front/sys/feedback/list" tags:"平台反馈-反馈列表" method:"get" summary:"平台反馈-反馈列表"`

	Page int `json:"page"  d:"1"   dc:"分页号码"`
	Size int `json:"size"  d:"500"  dc:"分页数量"`
}

type BaseListRes struct {
	Items   interface{} `json:"items"    dc:"分类列表"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type BaseAddReq struct {
	g.Meta `path:"/front/sys/feedback/add" tags:"添加平台反馈-举报" method:"post" summary:"添加平台反馈-举报"`

	*entity.FeedbackBase
}

type BaseEditRes struct {
	FeedbackId interface{} `json:"feedback_id"   dc:"反馈编号"`
}
