package model

import (
	"golershop.cn/internal/model/entity"
)

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
	UserExp                float64 `json:"user_exp"                  ` // 经验值
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

	//其它信息
	UdAddress          string      `json:"ud_address"              ` //详细地址
	UserLevelName      string      `json:"user_level_name"         ` //等级名称
	UserRegTime        interface{} `json:"user_reg_time"           ` //注册时间
	UserLoginTime      interface{} `json:"user_login_time"         ` //登录时间
	TagTitles          string      `json:"tag_titles"              ` //标签标题(DOT)
	TagTitleList       []string    `json:"tag_title_list"          ` //标签标题(DOT)
	TagGroupNames      string      `json:"tag_group_names"         ` //分组名称(DOT)
	TagIds             string      `json:"tag_ids"                 ` //用户标签(DOT)
	MonthOrder         interface{} `json:"month_order"             ` //本月订单
	TotalOrder         interface{} `json:"total_order"             ` //总计订单
	MonthTrade         interface{} `json:"month_trade"             ` //本月消费金额
	TotalTrade         interface{} `json:"total_trade"             ` //总消费金额
	Voucher            int64       `json:"voucher"                 ` //优惠券数量
	WaitPayNum         int64       `json:"wait_pay_num"            ` //待付款数量
	FavoritesGoodsNum  int64       `json:"favorites_goods_num"     ` //收藏数量
	ConcernNum         int64       `json:"concern_num"             ` //关注数量
	UnreadNumber       int         `json:"unread_number"           ` //未读消息数量
	CommissionAmount   float64     `json:"commission_amount"       ` //佣金总额:历史总额度
	UserCommissionNow  float64     `json:"user_commission_now"     ` //累计佣金
	MonthCommissionBuy float64     `json:"month_commission_buy"    ` //本月预估收益
	UserParentId       interface{} `json:"user_parent_id"          ` //上级用户编号
}

// UserVoucherRes 优惠券列表结构体
type UserVoucherRes struct {
	entity.UserVoucher
	Id               uint           `json:"id"`
	StoreName        string         `json:"store_name"`
	ActivityRuleJson ActivityRuleVo `json:"activity_rule_json"`
	ItemIds          []string       `json:"item_ids"`
	ActivityState    uint           `json:"activity_state"`
	VoucherUserWay   uint           `json:"voucher_user_way"`
	VoucherEffect    bool           `json:"voucher_effect"`
}

type UserVoucherListOutput struct {
	Items []*UserVoucherRes

	Page    int // 分页号码
	Total   int // 总页数
	Records int // 数据总数
	Size    int // 单页数量
}

type UserInfo struct {
	entity.UserInfo
	UserParentId uint `json:"user_parent_id"   dc:"上级用户编号"`
	Puid         uint `json:"puid"   dc:"平台标识"`
	Suid         uint `json:"suid"   dc:"用户标识"`
}
type UserInfoListOutput struct {
	Items []*UserInfo

	Page    int // 分页号码
	Total   int // 总页数
	Records int // 数据总数
	Size    int // 单页数量
}
