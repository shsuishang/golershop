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

type ActivityGroupbookingHistoryReq struct {
	g.Meta `path:"/front/marketing/activityBase/listsUserGroupbookingHistory" tags:"拼团记录表" method:"get" summary:"平团订单接口"`

	UserId   uint   `json:"user_id" dc:"用户编号"`                               // 用户编号
	OrderId  string `json:"order_id" dc:"拼团订单"`                              // 拼团订单
	GbId     int    `json:"gb_id" dc:"参团的编号"`                                // 参团的编号
	GbhFlag  bool   `json:"gbh_flag" dc:"是否支付(BOOL):0-未支付;1-已支付"`            // 是否支付(BOOL):0-未支付;1-已支付
	GbEnable string `json:"gb_enable" dc:"拼团状态(ENUM):1-成功;2-进程中;0-失败;3-未生效"` // 拼团状态(ENUM):1-成功;2-进程中;0-失败;3-未生效

	Page int `json:"page"  d:"1"   dc:"分页号码"`
	Size int `json:"size"  d:"500"  dc:"分页数量"`
}

type ActivityGroupbookingHistoryRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量

}

// ActivityGroupbookingReq 形成的拼团参数

type ActivityGroupbookingReq struct {
	g.Meta `path:"/front/marketing/activityBase/getUserGroupbooking" tags:"形成的拼团参数" method:"get" summary:"形成的拼团参数"`

	GbId    int    `json:"gb_id"     dc:"拼团编号"` // 拼团编号
	UserId  uint   `json:"user_id"   dc:"用户编号"` // 用户编号
	OrderId string `json:"order_id"  dc:"订单编号"` // 订单编号
}

type ActivityGroupbookingRes struct {
	*entity.ActivityGroupbooking
	OrderItemImage     string             `json:"order_item_image"    dc:"商品图片"`        // 商品图片
	ItemUnitPrice      float64            `json:"item_unit_price"     dc:"商品价格单价"`      // 商品价格单价
	OrderItemSalePrice float64            `json:"order_item_sale_price" dc:"商品实际成交价单价"` // 商品实际成交价单价
	ItemId             int64              `json:"item_id"             dc:"SKU编号"`       // SKU编号
	ProductName        string             `json:"product_name"        dc:"商品名称"`        // 商品名称
	ItemName           string             `json:"item_name"           dc:"规格名称"`        // 规格名称
	ActivityRemark     string             `json:"activity_remark"     dc:"活动说明"`        // 活动说明
	ActivityEndtime    uint64             `json:"activity_endtime"    dc:"活动结束时间"`      // 活动结束时间
	GbUsers            []*entity.UserInfo `json:"gb_users"            dc:"参团用户信息"`      // 参团用户信息
}
