package model

type LoginInput struct {
	UserAccount string `json:"user_account"  v:"required#请输入用户账号"           `  // 用户账号
	Password    string `json:"password"      v:"required#请输入用户密码"           `  // 用户密码
	VerifyCode  string `json:"verify_code"   dc:"验证码"                        ` // 验证码
	VerifyKey   string `json:"verify_key"    dc:"验证码KEY"                    `  // 验证码KEY
	Encrypt     bool   `json:"encrypt"       dc:"密码是否加密"                    `  // 密码是否加密

	UserId uint `json:"user_id"`
}

type RegisterInput struct {
	LoginInput

	UserNickname  string `json:"user_nickname"    v:"required#请输入用户昵称"             `    // 用户昵称
	UserAvatar    string `json:"user_avatar"       v:"required#请输入用户头像"             `   // 用户头像
	UserIntl      string `json:"user_intl"         dc:"国家编码"                        `   // 国家编码
	UserMobile    uint64 `json:"user_mobile"       v:"required#请输入手机号"             `    // 手机号
	UserEmail     string `json:"user_email"        v:"required#请输入邮箱"               `   // 邮箱
	ActivityId    uint   `json:"activity_id"       v:"required#请输入活动编号"             `   // 活动编号
	UserParentId  uint   `json:"user_parent_id"    dc:"来源用户编号"                      `   // 来源用户编号
	SourceUccCode string `json:"source_ucc_code"   dc:"渠道码"                           ` // 渠道码
	BindType      uint   `json:"bind_type"         dc:"注册方式=>"         `                // 注册方式
	RoleId        uint   `json:"role_id"           dc:"角色编号:0-用户;2-商家;3-门店;9-平台;"`      // 角色编号
	StoreId       uint   `json:"store_id"         dc:"店铺编号"         `                   // 店铺编号
	ChainId       uint   `json:"chain_id"         dc:"门店编号"         `                   // 门店编号
}

type LoginOutput struct {
	Token  string `json:"token"`
	UserId uint   `json:"user_id"`
}
