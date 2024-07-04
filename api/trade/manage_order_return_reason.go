package trade

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type OrderReturnReasonAdd struct {
	ReturnReasonId   uint   `json:"return_reason_id"   ` // 售后编号
	ReturnReasonName string `json:"return_reason_name" ` // 售后理由
	ReturnReasonSort int    `json:"return_reason_sort" ` // 售后排序
}

type OrderReturnReasonEditReq struct {
	g.Meta `path:"/manage/trade/orderReturnReason/edit" tags:"退货原因表" method:"post" summary:"退货原因表编辑接口"`

	ReturnReasonId uint `json:"return_reason_id"   ` // 售后编号
	OrderReturnReasonAdd
}

type OrderReturnReasonEditRes struct {
	ReturnReasonId uint `json:"return_reason_id"   ` // 售后编号
}

type OrderReturnReasonAddReq struct {
	g.Meta `path:"/manage/trade/orderReturnReason/add" tags:"退货原因表" method:"post" summary:"退货原因表编辑接口"`

	OrderReturnReasonAdd
}

type OrderReturnReasonRemoveReq struct {
	g.Meta `path:"/manage/trade/orderReturnReason/remove" tags:"退货原因表" method:"post" summary:"退货原因表删除接口"`

	ReturnReasonId uint `json:"return_reason_id"   ` // 售后编号
}

type OrderReturnReasonRemoveRes struct {
}

type OrderReturnReasonListReq struct {
	g.Meta `path:"/manage/trade/orderReturnReason/list" tags:"退货原因表" method:"get" summary:"退货原因表列表接口"`
	ml.BaseList

	ReturnReasonName string `json:"return_reason_name" type:"LIKE" ` // 售后理由
}

type OrderReturnReasonListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
