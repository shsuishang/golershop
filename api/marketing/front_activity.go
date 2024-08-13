package marketing

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/entity"
)

type ActivityVoucherListReq struct {
	g.Meta `path:"/front/marketing/activityBase/listVoucher" tags:"活动管理" method:"get" summary:"活动列表接口"`
	ml.BaseList
	ActivityReqVo
}

type ActivityVoucherListRes model.ActivityListOutput

type ActivityListReq struct {
	g.Meta `path:"/front/marketing/activityBase/list" tags:"活动表" method:"get" summary:"活动表"`
	ml.BaseList

	ActivityName   string      `json:"activity_name" type:"LIKE"` // 活动名称      ` // 活动名称
	ActivityTypeId uint        `json:"activity_type_id"   `       // 活动类型
	ActivityState  interface{} `json:"activity_state"     `       // 活动状态(ENUM):0-未开启;1-正常;2-已结束;3-管理员关闭;4-商家关闭
}

type ActivityListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ProductItemListReq struct {
	g.Meta `path:"/front/marketing/activityBase/getActivityInfo" tags:"活动商品详情" method:"get" summary:"活动商品详情"`
	ml.BaseList

	entity.ProductItem
}
type ProductItemListRes struct {
	Activitybase entity.ActivityBase         `json:"activity_base" dc:"活动信息"`
	Assists      []model.ProductAssistOutput `json:"assists" dc:"分类辅助属性"`
	Items        interface{}                 `json:"items"    dc:"分页数据内容"`
	Page         int                         `json:"page"`    // 分页号码
	Total        int                         `json:"total"`   // 总页数
	Records      int                         `json:"records"` // 数据总数
	Size         int                         `json:"size"`    // 单页数量
}
