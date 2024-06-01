package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CaptchaReq struct {
	g.Meta `path:"/front/sys/captcha/index" tags:"captcha" method:"get" summary:"获取默认的验证码" dc:"返回的是图片二进制内容"`
	IdKey  string `p:"verify_key" v:"required#验证码KEY不能为空"`
}

type CaptchaRes struct {
	g.Meta `mime:"png" dc:"验证码二进制内容" `
}

type SendMobileVerifyCodeReq struct {
	g.Meta `path:"/front/sys/captcha/mobile" tags:"captcha" method:"get" summary:"手机验证码" dc:"手机验证码"`
	Mobile string `p:"verify_key" v:"required#手机号码不能为空"`
}

type SendMobileVerifyCodeRes struct {
	VerifyCode string `p:"verify_code"`
}

type SendEmailVerifyCodeReq struct {
	g.Meta `path:"/front/sys/captcha/email" tags:"captcha" method:"get" summary:"Email验证码" dc:"手机验证码"`
	Email  string `p:"verify_key" v:"email"`
}

type SendEmailVerifyCodeRes struct {
	VerifyCode string `p:"verify_code"`
}
