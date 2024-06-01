// +----------------------------------------------------------------------
// | ShopSuite商城系统 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 随商信息技术（上海）有限公司
// +----------------------------------------------------------------------
// | 未获商业授权前，不得将本软件用于商业用途。禁止整体或任何部分基础上以发展任何派生版本、
// | 修改版本或第三方版本用于重新分发。
// +----------------------------------------------------------------------
// | 官方网站: https://www.shopsuite.cn  https://www.golershop.cn
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本公司对该软件产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细见https://www.golershop.cn/policy
// +----------------------------------------------------------------------

package sys

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/sys"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/dao/global"
	"golershop.cn/internal/service"
	"golershop.cn/utility"
	"golershop.cn/utility/email"
	"golershop.cn/utility/phone"
)

// 用户控制器管理对象
var Captcha = new(cCaptcha)

type cCaptcha struct{}

// Captcha 验证码
func (c *cCaptcha) Captcha(ctx context.Context, req *sys.CaptchaReq) (out sys.CaptchaRes, err error) {
	captchaStoreKey := req.IdKey
	//captchaStoreKey := guid.S()
	err = service.Captcha().NewAndStore(ctx, captchaStoreKey)
	return
}

// 发送手机验证码
func (c *cCaptcha) SendMobileVerifyCode(ctx context.Context, req *sys.SendMobileVerifyCodeReq) (res sys.SendMobileVerifyCodeRes, err error) {
	configBaseUserId := service.ConfigBase().GetInt(ctx, "service_user_id", 0)
	configBaseAppKey := service.ConfigBase().GetStr(ctx, "service_app_key", "")
	verifyCode := gconv.String(utility.GenerateRandomNumbers(4))

	if !phone.IsValidNumber(req.Mobile) {
		err = gerror.New("手机号码不准确！")
		return
	}

	phoneModelWithCountry := phone.GetPhoneModelWithCountry(req.Mobile)

	smsDto := &ml.SmsDto{
		Mobile:        gconv.String(phoneModelWithCountry.NationalNumber),
		ServiceUserId: configBaseUserId,
		ServiceAppKey: configBaseAppKey,
		Content:       fmt.Sprintf("您的验证码: [%s] 5分钟内有效", verifyCode),
	}

	if err = service.Cloud().Send(smsDto); err != nil {
		err = gerror.New(err.Error())
		return
	}

	err = service.Captcha().Store(ctx, req.Mobile, verifyCode)
	if global.Debug {
		res.VerifyCode = verifyCode
	}

	return
}

// 发送邮件验证码
func (c *cCaptcha) SendEmailVerifyCode(ctx context.Context, req *sys.SendEmailVerifyCodeReq) (res sys.SendEmailVerifyCodeRes, err error) {
	verifyCode := gconv.String(utility.GenerateRandomNumbers(4))

	/*
		// 验证email格式
		if err := g.Validator().Rules("email").Data(req.Email).Run(ctx); err != nil {
			return res, gerror.New("Email不准确！")
		}

		if !gstr.IsEmail(req.Email) {
			err = gerror.New("Email不准确！")
			return
		}
	*/

	queryWrapper := dao.ConfigBase.Ctx(ctx).Where("config_type_id", "12")
	configs, err := queryWrapper.All()
	if err != nil {
		return
	}

	emailDto := &email.SmtpConf{}
	configMap := make(map[string]string)
	for _, config := range configs {
		if config["ConfigKey"].String() == "email_host" {
			emailDto.SmtpHost = gconv.String(config["ConfigValue"])
		}

		if config["ConfigKey"].String() == "email_port" {
			emailDto.SmtpPort = gconv.Int(config["ConfigValue"])
		}

		if config["ConfigKey"].String() == "email_id" {
			emailDto.SmtpUserName = gconv.String(config["ConfigValue"])
		}

		if config["ConfigKey"].String() == "email_pass" {
			emailDto.SmtpPassword = gconv.String(config["ConfigValue"])
		}

		if config["ConfigKey"].String() == "email_fromname" {
			emailDto.Fromname = gconv.String(config["ConfigValue"])
		}
	}

	gconv.Struct(configMap, emailDto)

	subject := fmt.Sprintf("%s 注册验证码", service.ConfigBase().GetStr(ctx, "site_name", "MallSuite"))
	body := fmt.Sprintf("您的验证码: [%s] 5分钟内有效", verifyCode)

	err = email.SendSMTPMail(emailDto, req.Email, subject, body)

	if err != nil {
		return res, gerror.New("邮件发送失败！")
	}

	err = service.Captcha().Store(ctx, req.Email, verifyCode)

	if global.Debug {
		res.VerifyCode = verifyCode
	}

	return
}
