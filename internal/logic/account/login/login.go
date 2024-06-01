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

package login

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/google/uuid"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility"
	"golershop.cn/utility/phone"
	"strings"
	"time"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}

// 用户登录
func (s *sLogin) Login(ctx context.Context, user *entity.UserBase) (out model.LoginOutput, err error) {
	// 获取用户信息
	userInfo, err := dao.UserInfo.Get(ctx, user.UserId)

	// 判断当前用户状态
	if userInfo.UserState == 0 {
		return out, gerror.NewCode(gcode.CodeBusinessValidationFailed, "您的账号已被禁用,请联系管理员")
	}

	/*
		// 更新登录时间、登录IP
		dao.User.Data(g.Map{
			"login_time":  gtime.Now(),
			"login_ip":    utils.GetClientIp(r),
			"update_time": gtime.Now(),
		})
	*/

	// 生成Token
	token, _ := utility.GenerateToken(ctx, user.UserId, user.UserAccount, user.UserSalt)

	claim, err := utility.ParseToken(ctx, token)
	if err != nil {
		fmt.Println("解析token出现错误：", err)
	} else if time.Now().Unix() > claim.ExpiresAt {
		fmt.Println("时间超时")
	} else {
		fmt.Println("claim:", claim)
		//fmt.Println("username:", claim.UserId)
		//fmt.Println("username:", claim.Username)
		//fmt.Println("password:", claim.Password)
	}
	/*
		// 设置SESSION信息
		r.Session.Set("userId", user.Id)
		r.Session.Set("userInfo", user)
		sessionId := r.Session.Id()
		SessionList.Set(sessionId, r.Session)
	*/
	out.Token = token
	out.UserId = user.UserId

	return out, nil
}

// 用户登录
func (s *sLogin) LoginById(ctx context.Context, userId uint) (out model.LoginOutput, err error) {
	// 获取用户信息
	var user *entity.UserBase
	err = dao.UserBase.Ctx(ctx).WherePri(userId).Scan(&user)

	if err != nil {
		return out, err
	}

	return s.Login(ctx, user)
}

func (s *sLogin) DoLogin(ctx context.Context, in *model.LoginInput) (out model.LoginOutput, err error) {

	// 获取用户信息
	var user *entity.UserBase
	err = dao.UserBase.Ctx(ctx).Where(dao.UserBase.Columns().UserAccount, in.UserAccount).Scan(&user)

	if err != nil {
		return out, err
	}

	if user == nil {
		return out, gerror.NewCode(gcode.CodeBusinessValidationFailed, "用户名或者密码不正确")
	}

	// 密码校验
	//pwdOne, _ := gmd5.Encrypt(in.Password)
	pwd, _ := gmd5.Encrypt(in.Password + user.UserSalt)

	if user.UserPassword != pwd {
		return out, gerror.NewCode(gcode.CodeBusinessValidationFailed, "密码不正确")
	}

	return s.Login(ctx, user)
}

// Create creates user account.
func (s *sLogin) DoRegisterDep(ctx context.Context, in *model.RegisterInput) (token string, err error) {
	// If Nickname is not specified, it then uses Passport as its default Nickname.
	/*
		if in.Nickname == "" {
			in.Nickname = in.Passport
		}
		var (
			available bool
		)
		// Passport checks.
		available, err = s.IsPassportAvailable(ctx, in.Passport)
		if err != nil {
			return err
		}
		if !available {
			return gerror.Newf(`Passport "%s" is already token by others`, in.Passport)
		}
		// Nickname checks.
		available, err = s.IsNicknameAvailable(ctx, in.Nickname)
		if err != nil {
			return err
		}
		if !available {
			return gerror.Newf(`Nickname "%s" is already token by others`, in.Nickname)
		}

	*/

	//盐值
	salt := grand.S(12)

	// 密码校验
	//pwdOne, _ := gmd5.Encrypt(in.Password)
	pwd, _ := gmd5.Encrypt(in.Password + salt)

	// 互联登录校验扩展
	//user_token := guid.S()

	// 生成唯一用户编号

	dao.UserBase.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		userId, err2 := dao.UserBase.Add(ctx, &do.UserBase{
			UserAccount:  in.UserAccount,
			UserPassword: pwd,
			UserSalt:     salt,
		})

		if err2 != nil {
			return err2
		}

		_, err = dao.UserInfo.Add(ctx, &do.UserInfo{
			UserId:       userId,
			UserAccount:  in.UserAccount,
			UserNickname: in.UserNickname,
		})

		/*
			_, err = dao.UserBase.Ctx(ctx).Data(do.UserBase{
				UserAccount:  in.UserName,
				UserNickname: in.UserName,
				UserPassword: pwd,
				UserSalt:     salt,
			}).Insert()
		*/

		return err
	})

	return token, err
}

// Register 用户注册
func (s *sLogin) DoRegister(ctx context.Context, input *model.RegisterInput) (userId uint, err error) {
	userId = 0
	bindType := input.BindType

	// 账号，手机，邮箱注册方式走校验
	userAccount := input.UserAccount

	switch bindType {
	case consts.MOBILE:
		if !phone.IsValidNumber(userAccount) {
			err = errors.New("请输入正确的手机号！")
			return
		}

		phoneModelWithCountry := phone.GetPhoneModelWithCountry(userAccount)

		input.UserMobile = phoneModelWithCountry.NationalNumber
		input.UserIntl = fmt.Sprintf("+%d", phoneModelWithCountry.CountryCode)

		// 判断connect绑定操作
		bindConnect, err := dao.UserBindConnect.Get(ctx, userAccount)
		if err != nil {
			return userId, err
		}
		if bindConnect != nil && bindConnect.BindActive {
			err = errors.New("手机号已经绑定过，不可以使用此手机号注册")
			return userId, err
		}

	case consts.EMAIL:
		input.UserEmail = userAccount

		// 判断connect绑定操作
		bindConnect, err := dao.UserBindConnect.Get(ctx, userAccount)
		if err != nil {
			return userId, err
		}
		if bindConnect != nil && bindConnect.BindActive {
			err = errors.New("Email已经绑定过，不可以使用此Email注册")
			return userId, err
		}

	case consts.ACCOUNT:
		// userAccount 不可包含特殊字符 +
		if strings.Contains(userAccount, "+") {
			err = errors.New("用户账号不可以包含特殊字符串！")
			return userId, err
		}

	case consts.WEIXIN, consts.WEIXIN_XCX:

	}

	dbUserBase, err := dao.UserBase.FindOne(ctx, &do.UserBaseListInput{Where: do.UserBase{
		UserAccount: userAccount,
	}})
	if err != nil {
		return userId, err
	}
	if dbUserBase != nil {
		err = errors.New("用户已经存在,请更换用户名")
		return
	}

	now := gtime.Now()

	//基础表
	userSalt := uuid.New().String()

	userBase := &do.UserBase{}
	userBase.UserAccount = userAccount
	userBase.UserPassword, _ = gmd5.Encrypt(input.Password + userSalt)
	userBase.UserSalt = userSalt

	dao.UserBase.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if id, err := dao.UserBase.Add(ctx, userBase); err != nil {
			return err
		} else {
			userId = gconv.Uint(id)
			userBase.UserId = userId
		}

		//
		userLogin := &do.UserLogin{}
		userLogin.UserId = userBase.UserId
		userLogin.UserRegDate = now
		userLogin.UserRegTime = now.UnixMilli()
		userLogin.UserRegIp = utility.GetClientIp(g.RequestFromCtx(ctx))
		userLogin.UserLastloginTime = now.UnixMilli()
		userLogin.UserLastloginIp = utility.GetClientIp(g.RequestFromCtx(ctx))

		if _, err := dao.UserLogin.Save(ctx, userLogin); err != nil {
			return err
		}

		//
		if input.UserNickname == "" {
			input.UserNickname = input.UserAccount
		}

		userInfo := &do.UserInfo{}
		userInfo.UserId = userBase.UserId
		userInfo.UserAccount = input.UserAccount
		userInfo.UserNickname = input.UserNickname

		userInfo.UserMobile = gconv.String(input.UserMobile)
		userInfo.UserIntl = input.UserIntl
		userInfo.UserGender = 0
		userInfo.UserEmail = input.UserEmail
		userInfo.UserAvatar = input.UserAvatar
		userInfo.UserBirthday = gtime.NewFromStr("1971-01-01")
		userInfo.UserLevelId = 1001

		if _, err := dao.UserInfo.Save(ctx, userInfo); err != nil {
			return err
		}

		// 是否为手机号注册
		if bindType == consts.MOBILE {
			bindId := input.UserAccount

			// connect绑定操作
			bindConnect := &do.UserBindConnect{}
			bindConnect.BindOpenid = input.UserAccount
			bindConnect.BindActive = true

			if !s.CheckBind(ctx, bindId, consts.MOBILE, userId, bindConnect) {
				return err
			}
		}

		// 是否为邮箱注册
		if bindType == consts.EMAIL {
			// connect绑定操作
			bindConnect := &do.UserBindConnect{}
			bindConnect.BindOpenid = input.UserAccount
			bindConnect.BindActive = true

			bindId := input.UserAccount

			if !s.CheckBind(ctx, bindId, consts.EMAIL, userId, bindConnect) {
				return err
			}
		}

		// User_Resource初始化
		if err := service.UserResource().InitUserPoints(ctx, userId); err != nil {
			return err
		}

		// 初始化用户经验表
		if err := service.UserResource().InitUserExperience(ctx, userId); err != nil {
			return err
		}

		// 判断传递的活动id

		sourceUserId := s.GetSourceUserId(ctx)
		if sourceUserId != 0 {
			input.UserParentId = sourceUserId
		}

		// 分销用户来源 - 平台推广员功能，佣金平台出
		// 修改用户上级关系
		userParentId := input.UserParentId

		if userParentId != 0 {
			UserBase, err := dao.UserBase.Get(ctx, userParentId)
			if err != nil {
				return err
			}
			if UserBase != nil {
				s.AddSourceUserId(ctx, userId, userParentId, input.ActivityId, input.SourceUccCode)
			}

			// 分享券活动id
		}

		//记录渠道来源
		sourceUccCode := s.GetSourceUccCode(ctx)
		if sourceUccCode != "" {
			//s.AddChannelSourceUserId(ctx, userBase.UserId, sourceUccCode)
		}

		return err
	})
	/*
		//初次注册消息
		messageId := "registration-of-welcome-information"
		args := g.Map{
			"user_account":  input.UserNickname,
			"register_time": gtime.Now().Format("Y-m-d H:i:s"),
		}
		if err := messageService.SendNoticeMsg(ctx, userBase.UserId, messageId, args); err != nil {
			return userId, err
		}

	*/

	return userId, nil
}

// DoResetPasswd 修改密码
//
// @param userId   用户编号
// @param password 用户密码
// @return
func (s *sLogin) DoResetPasswd(ctx context.Context, userId uint, password string) (bool, error) {
	// 重置密码
	userSalt := guid.S()
	encodePassword := s.EncodePassword(ctx, password, userSalt)

	userBase := &do.UserBase{
		UserId:       userId,
		UserPassword: encodePassword,
		UserSalt:     userSalt,
	}

	_, err := dao.UserBase.Save(ctx, userBase)
	if err != nil {
		return false, gerror.New("修改密码失败")
	}

	return true, nil
}

// EncodePassword 加密密码
func (s *sLogin) EncodePassword(ctx context.Context, password, userSalt string) string {
	pwd, _ := gmd5.Encrypt(password + userSalt)
	return pwd
}

// CheckBind 验证绑定
func (s *sLogin) CheckBind(ctx context.Context, bindId string, bindType uint, userId uint, userInfoRow *do.UserBindConnect) bool {
	bindRow, err := dao.UserBindConnect.Get(ctx, bindId)
	if err != nil {
		// 错误处理
	}

	if bindRow != nil && !g.IsEmpty(bindRow.UserId) {
		// 验证通过, 登录成功.
		bindUserId := bindRow.UserId
		if userId != 0 && userId == bindUserId {
			panic("非法请求,已经登录用户不应该访问到此页面-重复绑定")
		} else if userId == 0 && userId == bindUserId {
			panic("非法请求,错误绑定数据")
		}
	} else if bindRow != nil && g.IsEmpty(bindRow.UserId) {
		userInfoRow.UserId = userId
		userInfoRow.BindId = bindId

		if _, err := dao.UserBindConnect.Save(ctx, userInfoRow); err != nil {
			panic("ResultCode.FAILED")
		}
	} else if bindRow == nil {
		// todo 头像会变动, 需要本地缓存, 生成新网址.
		userInfoRow.BindId = bindId
		userInfoRow.BindType = bindType
		userInfoRow.UserId = userId
		userInfoRow.BindActive = true

		if _, err := dao.UserBindConnect.Save(ctx, userInfoRow); err != nil {
			panic("ResultCode.FAILED")
		}
	}

	return true
}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *sLogin) IsSignedIn(ctx context.Context) bool {
	if v := service.BizCtx().Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// AddSourceUserId 添加分销来源用户 - 台推广员功能，佣金平台出
func (s *sLogin) AddSourceUserId(ctx context.Context, userId, userParentId, activityId uint, sourceUccCode string) bool {
	// 分销用户来源 - 平台推广员功能，佣金平台出
	if userParentId != 0 {
		// 初始化推广员记录
		userActive := service.ConfigBase().GetBool(ctx, "distribution_user_auto_active", false)
		service.UserDistribution().InitDistributionUser(ctx, userParentId, userActive)

		// 查找合伙人
		userParentRow, _ := dao.UserDistribution.Get(ctx, userParentId)
		rootUserId := uint(0)
		if userParentRow != nil && userParentRow.UserParentId != 0 {
			rootUserId = userParentRow.UserParentId
		}

		userDistribution := &do.UserDistribution{
			UserId:        userId,
			UserParentId:  userParentId,
			UserPartnerId: rootUserId,
			UserTime:      uint64(time.Now().UnixMilli()),
			ActivityId:    activityId,
		}

		// 新注册用户
		fxUserRow, _ := dao.UserDistribution.Get(ctx, userId)
		if fxUserRow == nil {
			// 添加用户关系
			if !service.UserDistribution().AddPlantformUser(ctx, userDistribution) {
				panic("添加用户关系失败")
				return false
			}

			// 带来粉丝，赠送积分
			// 赠送积分
			plantformFxGiftPoint := service.ConfigBase().GetFloat(ctx, "plantform_fx_gift_point", 0)
			if plantformFxGiftPoint > 0 {
				desc := fmt.Sprintf("推广粉丝 %d 送积分 %.2f", userId, plantformFxGiftPoint)

				userPointsVo := &model.UserPointsVo{
					UserId:        userParentId,
					Points:        plantformFxGiftPoint,
					PointsTypeId:  consts.POINTS_TYPE_FX_FANS,
					PointsLogDesc: desc,
				}

				if _, err := service.UserResource().Points(ctx, userPointsVo); err != nil {
					panic("积分操作失败")
					return false
				}
			}
		}
	}

	return true
}

// GetSourceUserId 用户来源
func (s *sLogin) GetSourceUserId(ctx context.Context) uint {
	request := g.RequestFromCtx(ctx)
	sourceUserId := gconv.Uint(request.Get("source_user_id"))

	if sourceUserId == 0 {
		sourceUserId = gconv.Uint(request.Get("user_parent_id"))
	}

	return sourceUserId
}

// GetSourceUccCode 获取用户 UCC Code
func (s *sLogin) GetSourceUccCode(ctx context.Context) string {
	request := g.RequestFromCtx(ctx)
	sourceUccCode := gconv.String(request.Get("source_ucc_code"))
	return sourceUccCode
}
