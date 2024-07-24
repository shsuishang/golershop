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
type ActivityCutpriceHistoryListReq struct {
	g.Meta     `path:"/front/marketing/activityBase/listsCutpriceHistory" tags:"砍价历史记录" method:"get" summary:"砍价历史记录"`
	ActivityId int `json:"activity_id" dc:"活动编号"` // 活动编号
	UserId     int `json:"user_id" dc:"砍价用户"`     // 砍价用户
	AcId       int `json:"ac_id" dc:"砍价编号"`       // 砍价编号

	Sidx string `json:"sidx" dc:"排序字段" default:"ach_id"` // 排序字段
	Sort string `json:"sort" dc:"排序方式" default:"desc"`   // 排序方式
}

type ActivityCutpriceHistoryListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ActivityCutpriceReq struct {
	g.Meta `path:"/front/marketing/activityBase/getCutpriceActivity" tags:"砍价详情" method:"get" summary:"砍价详情"`

	ActivityId    uint `json:"activity_id"      dc:"活动编号"`   // 活动编号
	UserId        uint `json:"user_id"          dc:"用户编号"`   // 用户编号
	ParticipantId uint `json:"participant_id"   dc:"当前用户编号"` // 当前用户编号
}

type ActivityCutpriceRes struct {
	entity.ActivityCutprice
	ActivityTitle     string                          `json:"activity_title" dc:"活动名称"`                                                                                                                                                                                                                                                            // 活动名称
	ActivityState     uint                            `json:"activity_state" dc:"活动状态(ENUM):0-未开启;1-正常;2-已结束;3-关闭;"`                                                                                                                                                                                                                               // 活动状态
	ActivityStarttime uint64                          `json:"activity_starttime" dc:"活动开始时间"`                                                                                                                                                                                                                                                      // 活动开始时间
	ActivityEndtime   uint64                          `json:"activity_endtime" dc:"活动结束时间"`                                                                                                                                                                                                                                                        // 活动结束时间
	ActivityRuleJson  model.ActivityRuleVo            `json:"activity_rule_json" dc:"活动规则(JSON):不检索{rule_id:{}, rule_id:{}},统一解析规则{\"requirement\":{\"buy\":{\"item\":[1,2,3],\"subtotal\":\"通过计算修正满足的条件\"}},\"rule\":[{\"total\":100,\"max_num\":1,\"item\":{\"1\":1,\"1200\":3}},{\"total\":200,\"max_num\":1,\"item\":{\"1\":1,\"1200\":3}}]}"` // 活动规则
	Items             []*model.ProductItemVo          `json:"items" dc:"商品信息集合"`                                                                                                                                                                                                                                                                   // 商品信息集合
	CutRow            *entity.ActivityCutpriceHistory `json:"cut_row" dc:"砍价记录历史"`                                                                                                                                                                                                                                                                 // 砍价记录历史
	IsCut             bool                            `json:"is_cut" dc:"是否已砍"`                                                                                                                                                                                                                                                                    // 是否已砍
	UserAvatar        string                          `json:"user_avatar" dc:"用户头像"`
	UserNickname      string                          `json:"nickname" dc:"用户昵称"`
}

type ActivityCutpriceDoCutpriceReq struct {
	g.Meta `path:"/front/marketing/activityBase/doCutprice" tags:"砍价-参加活动，并报名" method:"post" summary:"砍价-参加活动，并报名"`

	AcId uint `json:"ac_id" dc:"砍价编号"` // 砍价编号
}
type ActivityCutpriceDoCutpriceRes struct {
	entity.ActivityCutpriceHistory
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
