package marketing

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
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

	ActivityId uint `json:"activity_id"        ` // 活动编号
	ActivityBaseAdd
}

type ActivityBaseEditRes struct {
	ActivityId interface{} `json:"activity_id"        ` // 活动编号
}

type ActivityBaseAddReq struct {
	g.Meta `path:"/manage/marketing/activityBase/add" tags:"活动管理" method:"post" summary:"活动新增接口"`

	ActivityBaseAdd
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
	ActivityReqVo
}

type ActivityBaseListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

//front/marketing/activityBase/listVoucher
