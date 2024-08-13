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
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/account"
	"golershop.cn/internal/service"
)

// CallbackMp 公众号授权回调
func (c *cLogin) CallbackMp(ctx context.Context, req *account.CallbackMpReq) (out account.CallbackMpRes, err error) {
	res, err := service.Wechat().CallbackMp(ctx, req.Code, req.ActivityId)

	if err != nil {
		return out, err
	}

	gconv.Struct(res, &out)
	return
}

// Login 系统登录
func (c *cLogin) CheckAppLogin(ctx context.Context, req *account.CheckAppLoginReq) (out account.CheckAppLoginRes, err error) {
	res, err := service.Wechat().CheckAppLogin(ctx, req.Code)

	if err != nil {
		return out, err
	}

	gconv.Struct(res, &out)
	return
}

// Code2Session 授权注册
func (c *cLogin) JsCode2Session(ctx context.Context, req *account.JsCode2SessionReq) (out account.JsCode2SessionRes, err error) {
	res, err := service.Wechat().JsCode2Session(ctx, req.Code, req.EncryptedData, req.Iv, req.UserInfo, req.ActivityId, req.SourceUserId)

	if err != nil {
		return out, err
	}

	gconv.Struct(res, &out)
	return
}

// Code2Session 授权注册
func (c *cLogin) OpenIdByCode(ctx context.Context, req *account.OpenIdByCodeReq) (out account.OpenIdByCodeRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	res, err := service.Wechat().GetOpenIdByCode(ctx, req.Code, userId)

	if err != nil {
		return out, err
	}

	gconv.Struct(res, &out)

	return
}
