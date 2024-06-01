package model

import "golershop.cn/internal/model/entity"

// UserDistributionVo 分销用户信息
type UserDistributionVo struct {
	entity.UserDistribution

	CommissionAmount     float64 `json:"commission_amount" dc:"佣金总额:历史总额度"`   // 佣金总额:历史总额度
	CommissionBuyAmount0 float64 `json:"commission_buy_amount_0" dc:"推广消费佣金"` // 推广消费佣金
	CommissionBuyAmount1 float64 `json:"commission_buy_amount_1" dc:"消费佣金"`   // 消费佣金
	CommissionBuyAmount2 float64 `json:"commission_buy_amount_2" dc:"消费佣金"`   // 消费佣金
	CommissionSettled    float64 `json:"commission_settled" dc:"已经结算佣金"`      // 已经结算佣金
	UserNickname         string  `json:"user_nickname" dc:"用户昵称"`             // 用户昵称
}

type UserDistributionOutput struct {
	Items   []*UserDistributionVo `json:"items"    dc:"分页数据内容"`
	Page    int                   `json:"page"`    // 分页号码
	Total   int                   `json:"total"`   // 总页数
	Records int                   `json:"records"` // 数据总数
	Size    int                   `json:"size"`    // 单页数量
}
