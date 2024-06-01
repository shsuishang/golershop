package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type UserPointsHistoryAdd struct {
	PointsLogId     uint        `json:"points_log_id"     ` //
	PointsKindId    uint        `json:"points_kind_id"    ` // 类型(ENUM):1-获取积分;2-消费积分;
	PointsTypeId    uint        `json:"points_type_id"    ` // 积分类型(ENUM):1-会员注册;2-会员登录;3-商品评论;4-购买商品;5-管理员操作;7-积分换购商品;8-积分兑换代金券
	UserId          uint        `json:"user_id"           ` // 会员编号
	PointsLogPoints float64     `json:"points_log_points" ` // 可用积分
	UserPoints      float64     `json:"user_points"       ` // 当前积分
	PointsLogTime   *gtime.Time `json:"points_log_time"   ` // 创建时间
	PointsLogDesc   string      `json:"points_log_desc"   ` // 描述
	StoreId         uint        `json:"store_id"          ` // 所属店铺
	PointsLogDate   *gtime.Time `json:"points_log_date"   ` // 积分日期
	UserIdOther     uint        `json:"user_id_other"     ` // 交互会员
	PointsLogState  uint        `json:"points_log_state"  ` // 领取状态(BOOL):0-未领取;1-已领取
	ExtId           string      `json:"ext_id"            ` // 关联单号
}

type UserPointsHistoryAddReq struct {
	g.Meta `path:"/manage/pay/userPointsHistory/add" tags:"余额记录" method:"post" summary:"余额记录编辑接口"`

	UserPointsHistoryAdd
}

type UserPointsHistoryEditReq struct {
	g.Meta `path:"/manage/pay/userPointsHistory/edit" tags:"余额记录" method:"post" summary:"余额记录编辑接口"`

	PointsLogId uint `json:"points_log_id"  dc:"余额记录编号"   ` // 余额记录编号
	UserPointsHistoryAdd
}

type UserPointsHistoryEditRes struct {
	PointsLogId uint `json:"points_log_id" dc:"余额记录编号"   ` // 余额记录编号
}

type UserPointsHistoryRemoveReq struct {
	g.Meta      `path:"/manage/pay/userPointsHistory/remove" tags:"余额记录" method:"post" summary:"余额记录删除接口"`
	PointsLogId uint `json:"points_log_id" dc:"余额记录编号"   ` // 余额记录编号
}

type UserPointsHistoryRemoveRes struct {
}

type UserPointsHistoryListReq struct {
	g.Meta `path:"/manage/pay/userPointsHistory/list" tags:"余额记录" method:"get" summary:"余额记录列表接口"`
	ml.BaseList

	PointsLogId uint `json:"points_log_id"                   ` // 用户编号
}

type UserPointsHistoryListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
