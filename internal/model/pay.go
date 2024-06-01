package model

// UserPointsVo 用户积分
type UserPointsVo struct {
	UserId        uint    `json:"user_id,omitempty"`         // 所属用户
	Points        float64 `json:"points,omitempty"`          // 积分
	PointsTypeId  uint    `json:"points_type_id,omitempty"`  // 积分类型
	PointsLogDesc string  `json:"points_log_desc,omitempty"` // 描述
	UserIdOther   uint    `json:"user_id_other,omitempty"`   // 相关用户
	OrderId       string  `json:"order_id,omitempty"`        // 订单编号
	StoreId       uint    `json:"store_id,omitempty"`        // 卖家店铺编号
}

// ExperienceVo 操作用户经验对象
type ExperienceVo struct {
	UserId    uint    `json:"user_id"    dc:"用户编号"`  // 用户编号
	Exp       float64 `json:"exp"        dc:"经验值"`   // 经验值
	ExpTypeId uint    `json:"exp_type_id" dc:"等级编号"` // 等级编号
	Desc      string  `json:"desc"       dc:"描述"`    // 描述
}

type SignInfoOutput struct {
	ContinueSignDays int           `json:"continue_sign_days"` // 连续签到天数
	SignDayArr       []string      `json:"sign_day_arr"`       // 签到日期数组
	SignList         []PointStepVo `json:"sign_list"`          // 签到列表
	TodayIsSign      int           `json:"today_is_sign"`      // 今日是否签到
}

// PointStepVo 签到信息对象
type PointStepVo struct {
	Times     string `json:"times"`     // 时间
	Days      int    `json:"days"`      // 天数
	Multiples string `json:"multiples"` // 倍数
	ValueStr  string `json:"value_str"` // 前端映射 天数或倍数
}
