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

/**
 * 登录-控制器
 * @author Xinze
 * @since 2021/11/18
 * @File : login
 */
package account

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/account"
	"golershop.cn/internal/model"
	"golershop.cn/internal/service"
	utility "golershop.cn/utility/rsa"
)

// 用户控制器管理对象
var Login = new(cLogin)

type cLogin struct{}

// Login 系统登录
func (c *cLogin) Login(ctx context.Context, req *account.LoginReq) (out account.LoginRes, err error) {
	//密码是否加密
	if req.Encrypt {
		decrypt, err := utility.RsaDecrypt(req.Password, utility.PrivateKey(ctx))

		if err != nil {
			return out, err
		}

		req.Password = decrypt
	}

	// 系统登录
	res, err := service.Login().DoLogin(ctx, &model.LoginInput{UserAccount: req.UserName, Password: req.Password})

	if err != nil {
		return out, err
	}

	if !g.IsEmpty(req.IdKey) && !service.Captcha().VerifyAndClear(g.RequestFromCtx(ctx), req.IdKey, req.Captcha) {
		return out, gerror.NewCode(gcode.CodeBusinessValidationFailed, "请输入正确的验证码")
	}

	gconv.Struct(res, &out)

	return
}

// Register 系统注册
func (c *cLogin) Register(ctx context.Context, req *account.RegisterReq) (out account.LoginRes, err error) {

	/*
		// 校验验证码
		verifyRes := base64Captcha.VerifyCaptcha(req.IdKey, req.Captcha)
		if !verifyRes {
			r.Response.WriteJsonExit(common.JsonResult{
				Code: -1,
				Msg:  "验证码不正确",
			})
		}

	*/
	input := &model.RegisterInput{}
	gconv.Struct(req, input)

	// 系统登录
	userId, err := service.Login().DoRegister(ctx, input)

	if err != nil {
		return out, err
	}

	res, err := service.Login().LoginById(ctx, userId)
	gconv.Struct(res, &out)

	return
}

// Logout 退出登录
func (c *cLogin) Logout(ctx context.Context, req *account.LogoutReq) (out account.LogoutRes, err error) {

	return
}

// Protocol 隐私政策
func (c *cLogin) Protocol(ctx context.Context, req *account.LoginPolicyReq) (res account.LoginPolicyRes, err error) {
	res = account.LoginPolicyRes{}

	if req.DocumentType == "store" {
		storeDescription, _ := service.ConfigBase().Get(ctx, "open_store_description")
		if storeDescription != nil {
			res.Document = storeDescription.ConfigValue
		}
	} else {
		configBase, _ := service.ConfigBase().Get(ctx, req.ProtocolsKey)
		if configBase != nil {
			res.Document = configBase.ConfigValue
		}
	}

	return res, nil
}
