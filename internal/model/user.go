package model

import "golershop.cn/internal/model/entity"

type UserInfoOutput struct {
	entity.UserInfo
	UserIdcardImageList []string `json:"user_idcard_image_list"               ` // 身份证图片(DOT)

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

	Permissions []string `json:"permissions"    ` // 权限列表
	Roles       []string `json:"roles"    `       // 角色列表

	RoleId   uint `json:"role_id"    `   // 角色编号:0-用户;2-商家;3-门店;9-平台;
	SiteId   uint `json:"site_id"    `   // 分站编号:0-总站
	StoreId  uint `json:"store_id"    `  // 店铺编号
	ChainId  uint `json:"chain_id"    `  // 门店编号
	ClientId uint `json:"client_id"    ` // 后台管理:admin=1;移动端front=0
}

// UserVoucherRes 优惠券列表结构体
type UserVoucherRes struct {
	entity.UserVoucher
	Id               uint           `json:"id"`
	StoreName        string         `json:"store_name"`
	ActivityRuleJson ActivityRuleVo `json:"activity_rule_json"`
	ItemIds          []string       `json:"item_ids"`
	ActivityState    uint           `json:"activity_state"`
	VoucherEffect    bool           `json:"voucher_effect"`
}

type UserVoucherListOutput struct {
	Items []*UserVoucherRes

	Page    int // 分页号码
	Total   int // 总页数
	Records int // 数据总数
	Size    int // 单页数量
}
