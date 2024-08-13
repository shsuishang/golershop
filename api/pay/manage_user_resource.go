package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type UserResourceAdd struct {
	UserId                 uint    `json:"user_id"                   ` // 用户编号
	CurrencyId             uint    `json:"currency_id"               ` // 货币编号
	CurrencySymbolLeft     string  `json:"currency_symbol_left"      ` // 左符号
	UserMoney              float64 `json:"user_money"                ` // 用户资金
	UserMoneyFrozen        float64 `json:"user_money_frozen"         ` // 冻结资金:待结算余额
	UserRechargeCard       float64 `json:"user_recharge_card"        ` // 充值卡余额
	UserRechargeCardFrozen float64 `json:"user_recharge_card_frozen" ` // 冻结充值卡:待结算
	UserPoints             float64 `json:"user_points"               ` // 积分
	UserPointsFrozen       float64 `json:"user_points_frozen"        ` // 冻结积分
	UserCredit             float64 `json:"user_credit"               ` // 可用信用
	UserCreditFrozen       float64 `json:"user_credit_frozen"        ` // 冻结额度
	UserCreditUsed         float64 `json:"user_credit_used"          ` // 使用信用
	UserCreditTotal        float64 `json:"user_credit_total"         ` // 信用额度
	UserMargin             float64 `json:"user_margin"               ` // 保证金
	UserRedpack            float64 `json:"user_redpack"              ` // 红包额度
	UserRedpackFrozen      float64 `json:"user_redpack_frozen"       ` // 红包冻结额度
	Version                uint    `json:"version"                   ` // 乐观锁
}

type UserResourceAddReq struct {
	g.Meta `path:"/manage/pay/userResource/add" tags:"会员支付" method:"post" summary:"会员支付编辑接口"`

	UserResourceAdd
}

type UserResourceEditReq struct {
	g.Meta `path:"/manage/pay/userResource/edit" tags:"会员支付" method:"post" summary:"会员支付编辑接口"`

	UserId uint `json:"user_id"  dc:"会员支付编号"   ` // 会员支付编号
	UserResourceAdd
}

type UserResourceEditRes struct {
	UserId uint `json:"user_id" dc:"会员支付编号"   ` // 会员支付编号
}

type UserResourceRemoveReq struct {
	g.Meta `path:"/manage/pay/userResource/remove" tags:"会员支付" method:"post" summary:"会员支付删除接口"`
	UserId uint `json:"user_id" dc:"会员支付编号"   ` // 会员支付编号
}

type UserResourceRemoveRes struct {
}

type UserResourceListReq struct {
	g.Meta `path:"/manage/pay/userResource/list" tags:"会员支付" method:"get" summary:"会员支付列表接口"`
	ml.BaseList

	UserId uint `json:"user_id"                   ` // 用户编号
}

type UserResourceListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type UpdateUserMoneyReq struct {
	g.Meta `path:"/manage/pay/userResource/updateUserMoney" tags:"修改资金" method:"post" summary:"修改资金接口"`

	UserId      uint    `json:"user_id"                   `             // 用户编号
	RecordTotal float64 `json:"record_total"  v:"required#用户资金不能为空"   ` // 用户资金
}

type UpdateUserMoneyRes struct {
	UserId uint `json:"user_id"                   ` // 用户编号
}

type UpdatePointsReq struct {
	g.Meta `path:"/manage/pay/userResource/updatePoints" tags:"修改积分" method:"post" summary:"修改积分接口"`

	UserId     uint    `json:"user_id"                   ` // 用户编号
	UserPoints float64 `json:"points"               `      // 积分
}

type UpdatePointsRes struct {
	UserId uint `json:"user_id"                   ` // 用户编号
}
