package account

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"golershop.cn/internal/model"
)

type UserInfoReq struct {
	g.Meta `path:"/front/account/user/info" tags:"user" method:"get" summary:"登录用户信息接口"`
}

type UserInfoRes struct {
	model.UserInfoOutput
}
type UserEditReq struct {
	g.Meta `path:"/front/account/user/edit" tags:"user" method:"post" summary:"登录编辑信息接口"`

	UserId       uint        `json:"user_id"                ` // 用户编号
	UserNickname string      `json:"user_nickname"          ` // 用户昵称
	UserAvatar   string      `json:"user_avatar"            ` // 用户头像
	UserEmail    string      `json:"user_email"             ` // 用户邮箱(email)
	UserBirthday *gtime.Time `json:"user_birthday"          ` // 生日(DATE)
}

type UserEditRes struct {
	model.UserInfoOutput
}

type BindMobileReq struct {
	g.Meta `path:"/front/account/user/bindMobile" tags:"user" method:"post" summary:"绑定手机号"`

	VerifyCode string `json:"verify_code"` // 验证码
	VerifyKey  string `json:"verify_key"`  // 验证码KEY
}

type BindMobileRes model.LoginOutput

type UnBindMobileReq struct {
	g.Meta `path:"/front/account/user/unBindMobile" tags:"user" method:"post" summary:"重新绑定手机号"`

	VerifyCode string `json:"verify_code"` // 验证码
	VerifyKey  string `json:"verify_key"`  // 验证码KEY
}

type UnBindMobileRes model.LoginOutput

// ResetPasswordReq 登录参数
type ResetPasswordReq struct {
	g.Meta `path:"/front/account/login/setNewPassword" tags:"用户" method:"post" summary:"重设密码接口"`

	Password    string `json:"password"    v:"required#请输入用户密码"    dc:"用户密码"`                      // 用户密码
	OldPassword string `json:"old_password"                            dc:"用户密码"`                  // 用户密码
	VerifyCode  string `json:"verify_code"                             dc:"验证码"`                   // 验证码
	VerifyKey   string `json:"verify_key"                              dc:"验证码KEY"`                // 验证码KEY
	Encrypt     bool   `json:"encrypt"                                 dc:"密码是否加密"`                // 密码是否加密
	BindType    uint   `json:"bind_type"                               dc:"注册方式=>BindConnectCode"` // 注册方式=>BindConnectCode
}

type ResetPasswordRes struct {
}

type ChangePasswordReq struct {
	g.Meta `path:"/front/account/login/changePassword" tags:"用户" method:"post" summary:"修改密码"`

	Password    string `json:"password"    v:"required#请输入用户密码"    dc:"用户密码"`                      // 用户密码
	OldPassword string `json:"old_password"                            dc:"用户密码"`                  // 用户密码
	VerifyCode  string `json:"verify_code"                             dc:"验证码"`                   // 验证码
	VerifyKey   string `json:"verify_key"                              dc:"验证码KEY"`                // 验证码KEY
	Encrypt     bool   `json:"encrypt"                                 dc:"密码是否加密"`                // 密码是否加密
	BindType    uint   `json:"bind_type"                               dc:"注册方式=>BindConnectCode"` // 注册方式=>BindConnectCode
}

type ChangePasswordRes struct {
}

type CertificateReq struct {
	g.Meta `path:"/front/account/user/saveCertificate" tags:"用户" method:"post" summary:"实名认证"`

	UserId           uint   `json:"user_id"                ` // 用户编号
	UserRealname     string `json:"user_realname"          ` // 真实姓名
	UserIdcard       string `json:"user_idcard"            ` // 身份证
	UserIdcardImages string `json:"user_idcard_images"     ` // 身份证图片(DTO)
}

type CertificateRes struct {
}
