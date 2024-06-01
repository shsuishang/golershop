package account

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/account"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
	"golershop.cn/utility/phone"
	"golershop.cn/utility/strings"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) User(ctx context.Context, req *account.UserInfoReq) (res *account.UserInfoRes, err error) {
	loginUser := service.BizCtx().GetUser(ctx)

	if loginUser == nil {
		return nil, gerror.NewCode(gcode.New(0, "尚未登录", nil))
	} else {
		var info *model.UserInfoOutput
		info, err = service.User().GetUserInfo(ctx)

		gconv.Scan(info, &res)
	}

	return res, err
}

// Edit 修改用户信息
func (c *cUser) EditUser(ctx context.Context, req *account.UserEditReq) (res *account.UserEditRes, err error) {
	// 将请求参数转换为数据对象
	userInfo := do.UserInfo{}
	gconv.Struct(req, &userInfo)

	// 调用编辑用户信息的服务方法
	success, err := service.User().UserEdit(ctx, &userInfo)
	if err != nil {
		return res, err
	}

	// 根据编辑结果返回对应的响应
	if success == 0 {
		return res, err
	} else {
		return res, gerror.New("操作失败")
	}
}

// BindMobile 绑定手机号
func (c *cUser) BindMobile(ctx context.Context, req *account.BindMobileReq) (res account.BindMobileRes, err error) {
	// 判断手机号
	phoneModelWithCountry := phone.GetPhoneModelWithCountry(req.VerifyKey)
	if phoneModelWithCountry == nil {
		return res, gerror.NewCode(gcode.CodeBusinessValidationFailed, "手机号码不正确")
	}

	// 验证码
	if !service.Captcha().VerifyAndClear(g.RequestFromCtx(ctx), req.VerifyKey, req.VerifyCode) {
		return res, gerror.NewCode(gcode.CodeBusinessValidationFailed, "验证码有误")
	}

	// 获取当前登录用户
	loginUser := service.BizCtx().GetUser(ctx)

	// 绑定手机号
	loginRes, err := service.User().BindMobile(ctx, loginUser, phoneModelWithCountry.CountryCodeStr, phoneModelWithCountry.NationalNumber, "")
	if err != nil {
		return res, gerror.NewCode(gcode.CodeBusinessValidationFailed, "绑定手机号失败")
	}

	gconv.Struct(loginRes, &res)

	return
}

// UnBindMobile 重新绑定手机号
func (c *cUser) UnBindMobile(ctx context.Context, req *account.UnBindMobileReq) (res account.UnBindMobileRes, err error) {
	// 判断手机号
	phoneModelWithCountry := phone.GetPhoneModelWithCountry(req.VerifyKey)
	if phoneModelWithCountry == nil {
		return res, gerror.NewCode(gcode.CodeBusinessValidationFailed, "手机号码不正确")
	}

	// 验证码
	if !service.Captcha().VerifyAndClear(g.RequestFromCtx(ctx), req.VerifyKey, req.VerifyCode) {
		return res, gerror.NewCode(gcode.CodeBusinessValidationFailed, "验证码有误")
	}

	// 获取当前登录用户
	loginUser := service.BizCtx().GetUser(ctx)

	// 解绑旧手机号并绑定新手机号
	success, err := service.User().UnBindMobile(ctx, loginUser, phoneModelWithCountry.CountryCodeStr, phoneModelWithCountry.NationalNumber)
	if err != nil || !success {
		return res, gerror.NewCode(gcode.CodeBusinessValidationFailed, "重新绑定手机号失败")
	}

	return res, nil
}

// SetNewPassword 重设密码接口
func (c *cUser) SetNewPassword(ctx context.Context, req *account.ResetPasswordReq) (res *account.ResetPasswordRes, err error) {
	// 验证码
	if !service.Captcha().VerifyAndClear(g.RequestFromCtx(ctx), req.VerifyKey, req.VerifyCode) {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "验证码有误")
	}

	if gstr.Trim(req.Password) == "" {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "密码有误")
	}

	var userId uint
	switch req.BindType {
	case consts.MOBILE:
		// 判断手机号
		phoneModelWithCountry := phone.GetPhoneModelWithCountry(req.VerifyKey)
		if phoneModelWithCountry == nil {
			return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "手机号码不正确")
		}
		bindId := fmt.Sprintf("%s%d", phoneModelWithCountry.CountryCodeStr, phoneModelWithCountry.NationalNumber)
		userBindConnect, err := dao.UserBindConnect.Get(ctx, bindId)
		if err != nil {
			return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "手机号码未绑定登录")
		}
		userId = userBindConnect.UserId

	case consts.EMAIL:
		if flag, _ := strings.IsEmail(ctx, req.VerifyKey); !flag {
			return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "该邮箱不存在")
		}

		userBindConnect, err := dao.UserBindConnect.Get(ctx, req.VerifyKey)
		if err != nil {
			return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "该邮箱绑定不存在")
		}
		userId = userBindConnect.UserId

	case consts.ACCOUNT:
		user := service.BizCtx().GetUser(ctx)
		if user == nil {
			return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户未登录")
		}
		userId = user.UserId
		userBase, err := dao.UserBase.Get(ctx, userId)
		if err != nil {
			return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户信息获取失败")
		}

		if userBase.UserPassword != service.Login().EncodePassword(ctx, req.OldPassword, userBase.UserSalt) {
			return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "旧密码不正确")
		}
	default:
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "绑定类型错误")
	}

	if _, err := service.Login().DoResetPasswd(ctx, userId, req.Password); err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "重设密码失败")
	}

	return &account.ResetPasswordRes{}, nil
}

// ChangePassword 修改密码
func (c *cUser) ChangePassword(ctx context.Context, req *account.ChangePasswordReq) (res *account.ChangePasswordRes, err error) {
	if gstr.Trim(req.Password) == "" {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "密码有误")
	}

	user := service.BizCtx().GetUser(ctx)
	if user == nil {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户未登录")
	}

	userBase, err := dao.UserBase.Get(ctx, user.UserId)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "用户信息获取失败")
	}

	if userBase.UserPassword != service.Login().EncodePassword(ctx, req.OldPassword, userBase.UserSalt) {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "旧密码不正确")
	}

	if _, err := service.Login().DoResetPasswd(ctx, user.UserId, req.Password); err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "修改密码失败")
	}

	return &account.ChangePasswordRes{}, nil
}

// SaveCertificate 实名认证
func (c *cUser) SaveCertificate(ctx context.Context, req *account.CertificateReq) (res *account.CertificateRes, err error) {
	res = &account.CertificateRes{}

	// 将请求参数转换为数据对象
	userInfo := &do.UserInfo{}
	gconv.Struct(req, userInfo)

	userInfo.UserId = service.BizCtx().GetUserId(ctx)
	userInfo.UserIsAuthentication = consts.USER_CERTIFICATION_VERIFY

	// 调用编辑用户信息的服务方法
	success, err := service.User().SaveCertificate(ctx, userInfo)
	if err != nil {
		return res, err
	}

	// 根据编辑结果返回对应的响应
	if success {
		return res, err
	} else {
		return res, gerror.New("操作失败")
	}
}
