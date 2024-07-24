package marketing

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// start fo manage
type ActivityBaseAdd struct {
	ActivityName      string `json:"activity_name"      ` // 活动名称
	ActivityTitle     string `json:"activity_title"     ` // 活动标题
	ActivityRemark    string `json:"activity_remark"    ` // 活动说明
	ActivityTypeId    uint   `json:"activity_type_id"   ` // 活动类型
	ActivityStarttime uint64 `json:"activity_starttime" ` // 活动开始时间
	ActivityEndtime   uint64 `json:"activity_endtime"   ` // 活动结束时间
	ActivityState     uint   `json:"activity_state"     ` // 活动状态(ENUM):0-未开启;1-正常;2-已结束;3-管理员关闭;4-商家关闭
	ActivityRule      string `json:"activity_rule"      ` // 活动规则(JSON):不检索{rule_id:{}, rule_id:{}},统一解析规则{"requirement":{"buy":{"item":[1,2,3],"subtotal":"通过计算修正满足的条件"}},"rule":[{"total":100,"max_num":1,"item":{"1":1,"1200":3}},{"total":200,"max_num":1,"item":{"1":1,"1200":3}}]}
	ActivityType      uint   `json:"activity_type"      ` // 参与类型(ENUM):1-免费参与;2-积分参与;3-购买参与;4-分享参与
	ActivitySort      uint   `json:"activity_sort"      ` // 活动排序
	ActivityUseLevel  string `json:"activity_use_level" ` // 使用等级(DOT)
}

type ActivityBaseEditReq struct {
	g.Meta `path:"/manage/marketing/activityBase/edit" tags:"活动管理" method:"post" summary:"活动编辑接口"`

	entity.ActivityBase
	ActivityId uint   `json:"activity_id"        ` // 活动编号
	ItemIds    string `json:"item_ids"        `    // 活动SKU(DOT)
}

type ActivityBaseEditRes struct {
	ActivityId interface{} `json:"activity_id"        ` // 活动编号
}

type ActivityBaseAddReq struct {
	g.Meta `path:"/manage/marketing/activityBase/add" tags:"活动管理" method:"post" summary:"活动新增接口"`

	entity.ActivityBase
}

type ActivityBaseRemoveReq struct {
	g.Meta     `path:"/manage/marketing/activityBase/remove" tags:"活动管理" method:"post" summary:"活动删除接口"`
	ActivityId uint `json:"activity_id"        ` // 活动编号
}

type ActivityBaseRemoveRes struct {
}

type ActivityReqVo struct {
	ActivityName      string `json:"activity_name"      ` // 活动名称
	ActivityTitle     string `json:"activity_title"     ` // 活动标题
	ActivityTypeId    uint   `json:"activity_type_id"   ` // 活动类型
	ActivityStarttime uint64 `json:"activity_starttime" ` // 活动开始时间
	ActivityEndtime   uint64 `json:"activity_endtime"   ` // 活动结束时间
	ActivityState     uint   `json:"activity_state"     ` // 活动状态(ENUM):0-未开启;1-正常;2-已结束;3-管理员关闭;4-商家关闭
	ActivityType      uint   `json:"activity_type"      ` // 参与类型(ENUM):1-免费参与;2-积分参与;3-购买参与;4-分享参与
	ActivityUseLevel  string `json:"activity_use_level" ` // 使用等级(DOT)

	ActivityTypeIds string `json:"activity_type_ids"  type:"FIND_IN_SET_STR"` // 活动SKU(DOT)
	Met             string `json:"met" `                                      // 活动SKU(DOT)
}

type ActivityBaseListReq struct {
	g.Meta `path:"/manage/marketing/activityBase/list" tags:"活动管理" method:"get" summary:"活动列表接口"`
	ml.BaseList

	ActivityName   string      `json:"activity_name" type:"LIKE"` // 活动名称      ` // 活动名称
	ActivityTypeId uint        `json:"activity_type_id"   `       // 活动类型
	ActivityState  interface{} `json:"activity_state"     `       // 活动状态(ENUM):0-未开启;1-正常;2-已结束;3-管理员关闭;4-商家关闭
}

type ActivityBaseListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ActivityBaseEditStateReq struct {
	g.Meta `path:"/manage/marketing/activityBase/editState" tags:"活动管理" method:"post" summary:"活动编辑接口"`

	ActivityId uint   `json:"activity_id"        ` // 活动编号
	ItemIds    string `json:"item_ids"        `    // 活动SKU(DOT)
	ActivityBaseAdd
}

type ActivityBaseEditStateRes struct {
	ActivityId interface{} `json:"activity_id"        ` // 活动编号
}

type ActivityItemReq struct {
	g.Meta `path:"/manage/marketing/activityBase/getActivityBuyItems" tags:"活动商品列表" method:"get" summary:"活动商品列表"`

	ActivityId uint `json:"activity_id"        ` // 活动编号
}

type ActivityItemRes struct {
	entity.ActivityItem
	ProductName       string  `json:"product_name" dc:"产品名称"`                                   // 产品名称
	ItemUnitPrice     float64 `json:"item_unit_price" dc:"商品价格"`                                // 商品价格
	ProductImage      string  `json:"product_image" dc:"商品主图"`                                  // 商品主图
	IsOnSale          bool    `json:"is_on_sale" dc:"销售中"`                                      // 销售中
	ItemName          string  `json:"item_name" dc:"副标题(DOT):SKU名称"`                            // 副标题(DOT):SKU名称
	ItemEnable        int     `json:"item_enable" dc:"是否启用(LIST):1001-正常;1002-下架仓库中;1000-违规禁售"` // 是否启用
	AvailableQuantity int     `json:"available_quantity" dc:"可用库存"`                             // 可用库存
}

type ActivityGroupbookingListReq struct {
	g.Meta `path:"/manage/marketing/activityBase/getGroupbookingList" tags:"形成的拼团-分页列表查询" method:"get" summary:"形成的拼团-分页列表查询"`
	ml.BaseList

	ActivityId interface{} `json:"activity_id"        ` // 活动编号
	GbEnable   interface{} `json:"gb_enable"   `        // 拼团状态(ENUM):1-成功;2-进程中;0-失败;3-未生效
	UserId     interface{} `json:"user_id"  `
}
type ActivityGroupbookingListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ActivityGroupbookingHistoryListReq struct {
	g.Meta `path:"/manage/marketing/activityBase/getGroupbookingHistoryList" tags:"拼团历史记录-分页列表查询" method:"get" summary:"拼团历史记录-分页列表查询"`
	ml.BaseList

	GbId     interface{} `json:"gb_id"       ` // 参团的编号
	GbhFlag  interface{} `json:"gbh_flag"    ` // 是否支付(BOOL):0-未支付;1-已支付
	GbEnable interface{} `json:"gb_enable"   ` // 拼团状态(ENUM):1-成功;2-进程中;0-失败;3-未生效
	UserId   interface{} `json:"user_id"     ` // 用户编号
	OrderId  interface{} `json:"order_id"    ` // 拼团订单
}
type ActivityGroupbookingHistoryListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ActivityItemEditReq struct {
	g.Meta `path:"/manage/marketing/activityBase/editActivityItem" tags:"活动-修改商品活动价格" method:"post" summary:"活动-修改商品活动价格	"`

	entity.ActivityItem
}
type ActivityItemEditRes struct {
}

type ActivityItemBatchPriceEditReq struct {
	g.Meta `path:"/manage/marketing/activityBase/editBatchPrice" tags:"活动-统一折扣修改价格" method:"post" summary:"活动-统一折扣修改价格	"`

	ActivityId uint64  `json:"activity_id"`
	Discount   float64 `json:"discount"`
}
type ActivityItemBatchPriceEditRes struct {
}

type AddActivityBuyItemsReq struct {
	g.Meta `path:"/manage/marketing/activityBase/addActivityBuyItems" tags:"活动-添加商品" method:"post" summary:"活动-添加商品"`

	entity.ActivityBase
	ActivityId uint   `json:"activity_id"        ` // 活动编号
	ItemIds    string `json:"item_ids"        `    // 活动SKU(DOT)

}
type AddActivityBuyItemsRes struct {
}

type RemoveActivityBuyItemsReq struct {
	g.Meta `path:"/manage/marketing/activityBase/removeActivityBuyItems" tags:"活动-删除商品" method:"post" summary:"活动-删除商品"`

	entity.ActivityBase
	ActivityId uint   `json:"activity_id"        ` // 活动编号
	ItemIds    string `json:"item_ids"        `    // 活动SKU(DOT)
}
type RemoveActivityBuyItemsRes struct{}

type ActivityCutpriceListReq struct {
	g.Meta `path:"/manage/marketing/activityBase/getCutpriceList" tags:"砍价记录-分页列表查询" method:"get" summary:"砍价记录-分页列表查询"`

	ActivityId uint `json:"activity_id"        ` // 活动编号
}
type ActivityCutpriceListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

//front/marketing/activityBase/listVoucher
