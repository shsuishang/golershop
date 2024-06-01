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

package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/phone"
	"net/http"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// 登录用户信息
func (s *sUser) GetUserInfo(ctx context.Context) (out *model.UserInfoOutput, err error) {
	loginUser := service.BizCtx().GetUser(ctx)

	if loginUser == nil {
		return nil, gerror.NewCode(gcode.New(http.StatusUnauthorized, "Token失效，请重新登录", nil))
	}

	// 获取用户基本信息
	var userBase = &entity.UserBase{}
	userBase, err = dao.UserBase.Get(ctx, loginUser.UserId)
	if err != nil {
		return nil, err
	}

	if userBase != nil && userBase.UserSalt != loginUser.UserSalt {
		return nil, gerror.NewCode(gcode.New(http.StatusUnauthorized, "Token失效，请重新登录", nil))
	}

	//用户信息
	var userInfo = &entity.UserInfo{}
	userInfo, err = dao.UserInfo.Get(ctx, loginUser.UserId)

	if err != nil {
		return nil, err
	}

	// 判断当前用户状态
	if userInfo.UserState == 0 {
		return nil, errors.New("您的账号已被禁用,请联系管理员")
	}

	gconv.Scan(userInfo, &out)
	if g.IsEmpty(out.UserIdcardImages) {
		out.UserIdcardImageList = []string{}
	} else {
		out.UserIdcardImageList = gstr.Split(out.UserIdcardImages, ",")
	}

	userResource, err := dao.UserResource.Get(ctx, loginUser.UserId)

	if err != nil {
		return nil, err
	}

	if userResource != nil {
		gconv.Scan(userResource, &out)
	}

	out.RoleId = loginUser.RoleId
	out.SiteId = loginUser.SiteId
	out.StoreId = loginUser.StoreId
	out.ChainId = loginUser.ChainId
	out.ClientId = loginUser.ClientId

	if len(loginUser.Roles) > 0 {
		gconv.Scan(loginUser.Roles[0], &out)
	}

	roles := make([]string, 0, len(loginUser.Roles))
	for _, val := range loginUser.Roles {
		if val.UserRoleName != "" && !gstr.InArray(roles, val.UserRoleName) {
			roles = append(roles, val.UserRoleName)
		}
	}

	out.Roles = roles

	//columns := []string{"/manage/account/userLevel/remove", "/manage/account/userLevel/removeBatch"}
	columns := make([]string, 0, len(loginUser.Authorities))
	for _, val := range loginUser.Authorities {
		if val.MenuPermission != "" && !gstr.InArray(columns, val.MenuPermission) {
			columns = append(columns, val.MenuPermission)
		}
	}

	out.Permissions = columns

	/*
		//user admin
		userAdmin, err := dao.UserAdmin.Get(ctx, loginUser.UserId)
		if err != nil {
			return nil, err
		}

		if userAdmin != nil {
			out.RoleId = consts.ROLE_ADMIN

			userRole, err := dao.UserRole.Get(ctx, userAdmin.UserRoleId)
			if err != nil {
				return nil, err
			}

			if userRole != nil {
				gconv.Scan(userRole, &out)

				entitys, err := dao.MenuBase.Gets(ctx, gstr.Split(userRole.MenuIds, ","))
				if err != nil {
					return nil, err
				}

				//columns := []string{"/manage/account/userLevel/remove", "/manage/account/userLevel/removeBatch"}
				columns := make([]string, 0, len(entitys))
				for _, val := range entitys {
					if val.MenuPermission != "" && !gstr.InArray(columns, val.MenuPermission) {
						columns = append(columns, val.MenuPermission)
					}
				}

				out.Permissions = columns
			}

		}

	*/

	return out, nil
}

func (s *sUser) UserEdit(ctx context.Context, userinfo *do.UserInfo) (affected int64, err error) {

	if userinfo.UserNickname != "" {
		// 查询是否存在相同昵称的用户
		existUser, err := dao.UserInfo.FindOne(ctx, &do.UserInfoListInput{
			Where: do.UserInfo{
				UserNickname: userinfo.UserNickname,
			},
		})
		if err != nil {
			return 0, nil
		}

		//如果存在且不是当前用户，则抛出异常
		if existUser != nil && existUser.UserId != userinfo.UserId {
			return 0, errors.New("昵称已存在")
		}
	}

	if _, err := dao.UserInfo.Edit(ctx, userinfo.UserId, userinfo); err != nil {
		return 0, err
	}

	return 0, err
}

// BindMobile 绑定手机号
func (s *sUser) BindMobile(ctx context.Context, user *model.ContextUser, userIntl string, mobile uint64, newPassword string) (*model.LoginOutput, error) {
	bindId := fmt.Sprintf("%s%d", userIntl, mobile)

	if !phone.IsValidNumber(bindId) {
		return nil, gerror.New("请输入正确的手机号！")
	}

	// 新用户编号，可能和老用户不一致的... 微信新账户，绑定手机号，手机绑定作为唯一账户有老手机号存在，绑定到老用户上
	userIdNew, err := s.doBindMobile(ctx, user, userIntl, mobile)
	if err != nil || userIdNew == 0 {
		return nil, gerror.New("绑定用户失败")
	}

	userBase, err := dao.UserBase.Get(ctx, userIdNew)
	if err != nil {
		return nil, err
	}

	// 对账号进行密码修改
	if gstr.Trim(newPassword) != "" {
		//userAccount := userBase.UserAccount
		// 若需要执行密码重置，请取消注释以下代码
		// if !s.doResetPasswd(ctx, userAccount, newPassword, "") {
		//     return nil, gerror.New("密码重置失败")
		// }
	}

	// 登录
	login, err := service.Login().Login(ctx, userBase)
	if err != nil {
		return nil, err
	}

	return &login, nil
}

// UnBindMobile 解绑手机号
func (s *sUser) UnBindMobile(ctx context.Context, user *model.ContextUser, userIntl string, mobile uint64) (bool, error) {
	bindId := fmt.Sprintf("%s%d", userIntl, mobile)

	if !phone.IsValidNumber(bindId) {
		return false, gerror.New("请输入正确的手机号！")
	}

	if _, err := dao.UserBindConnect.Remove(ctx, bindId); err != nil {
		return false, err
	}

	// 如果bind表中存在记录，判断info表中是否存在
	userInfo, err := dao.UserInfo.Get(ctx, user.UserId)
	if err != nil {
		return false, err
	}

	userData := &do.UserInfo{}
	userData.UserIntl = ""
	userData.UserMobile = ""

	_, err = dao.UserInfo.Edit(ctx, userInfo.UserId, userData)

	return true, err
}

// ReBindMobile 重新绑定手机号
func (s *sUser) ReBindMobile(ctx context.Context, user *model.ContextUser, userIntl string, mobile uint64, newPassword string) (*model.LoginOutput, error) {
	// 删除旧手机绑定
	if _, err := s.UnBindMobile(ctx, user, userIntl, mobile); err != nil {
		return nil, err
	}

	loginRes, err := s.BindMobile(ctx, user, userIntl, mobile, "")
	if err != nil {
		return nil, err
	}

	return loginRes, nil
}

func (s *sUser) doBindMobile(ctx context.Context, currentUser *model.ContextUser, userIntl string, mobile uint64) (uint, error) {
	bindId := fmt.Sprintf("%s%d", userIntl, mobile)

	var userId uint

	// 判断是否已经绑定
	queryWrapper := &do.UserBindConnectListInput{
		Where: do.UserBindConnect{BindId: bindId, BindType: consts.MOBILE},
	}
	userBindConnect, err := dao.UserBindConnect.FindOne(ctx, queryWrapper)

	if err != nil {
		return 0, err
	}

	// 取得当前手机是否绑定账号
	oldUserDto, err := dao.UserBase.FindOne(ctx, &do.UserBaseListInput{
		Where: do.UserBase{UserAccount: bindId},
	})

	if err != nil {
		fmt.Println(oldUserDto)
		return 0, err
	}

	// 根据手机号查询bind表中是否已存在记录
	if userBindConnect != nil {
		// 同一个用户,不做操作
		if userBindConnect.UserId == currentUser.UserId {
			// 如果bind表中存在记录，判断info表中是否存在
			userInfo, err := dao.UserInfo.Get(ctx, currentUser.UserId)
			if err != nil {
				return 0, err
			}

			if gstr.Trim(userInfo.UserIntl) == "" || gstr.Trim(userInfo.UserMobile) == "" {
				userInfoDo := &do.UserInfo{}
				userInfoDo.UserIntl = userIntl
				userInfoDo.UserMobile = gconv.String(mobile)
				if _, err := dao.UserInfo.Edit(ctx, userInfo.UserId, userInfoDo); err != nil {
					return 0, err
				}
			}

			userId = currentUser.UserId
		} else {
			return 0, gerror.New("手机号已经绑定用户！")
		}
	} else {
		// 若不存在则新增一个绑定信息
		accountUserBindConnectNew := &do.UserBindConnect{
			BindId:     bindId,
			BindType:   consts.MOBILE,
			UserId:     currentUser.UserId,
			BindActive: true,
		}

		userId, err = s.doUserBind(ctx, accountUserBindConnectNew, 0, false)
		if err != nil {
			return 0, err
		}

		// 如果bind表中存在记录，判断info表中是否存在
		userInfo, err := dao.UserInfo.Get(ctx, currentUser.UserId)
		if err != nil {
			return 0, err
		}
		if gstr.Trim(userInfo.UserIntl) == "" || gstr.Trim(userInfo.UserMobile) == "" {
			userInfoDo := &do.UserInfo{}
			userInfoDo.UserIntl = userIntl
			userInfoDo.UserMobile = gconv.String(mobile)

			if _, err := dao.UserInfo.Edit(ctx, userInfo.UserId, userInfoDo); err != nil {
				return 0, err
			}
		}

		if userId == 0 {
			return 0, gerror.New("保存用户绑定失败！")
		}
	}

	return userId, nil
}

func (s *sUser) doUserBind(ctx context.Context, userInfoRow *do.UserBindConnect, activityId uint, regFlag bool) (uint, error) {
	var userId uint

	bindId := userInfoRow.BindId.(string)

	userBase, err := dao.UserBase.FindOne(ctx, &do.UserBaseListInput{
		Where: do.UserBase{UserAccount: bindId},
	})

	if err != nil {
		return 0, err
	}

	// unionId 获取，如果存在
	if !g.IsEmpty(userInfoRow.BindUnionid) {
		// 检测unionId, 判断用户
		if userBase == nil {
			// 根据unionId，判断已经绑定的用户
			queryWrapper := &do.UserBindConnectListInput{
				Where: do.UserBindConnect{BindUnionid: userInfoRow.BindUnionid},
			}
			findBindRow, err := dao.UserBindConnect.FindOne(ctx, queryWrapper)
			if err != nil {
				return 0, err
			}

			if findBindRow != nil {
				userBase, err := dao.UserBase.Get(ctx, findBindRow.UserId)
				if err != nil {
					return 0, err
				}

				// 判断bind openid是否存在，不存在则需要添加
				openBindRow, err := dao.UserBindConnect.Get(ctx, bindId)
				if err != nil {
					return 0, err
				}

				if openBindRow == nil {
					userInfoRow.BindActive = true
					userInfoRow.BindType = userInfoRow.BindType
					if _, err := dao.UserBindConnect.Save(ctx, userInfoRow); err != nil {
						return 0, err
					}
				}

				userId = userBase.UserId
				if _, err := s.checkBind(ctx, bindId, userInfoRow.BindType.(uint), userId, userInfoRow); err != nil {
					return 0, err
				}

				return userId, nil
			}
		}
	}

	// 自动注册用户
	if userBase != nil {
		userId = userBase.UserId
	} else {
		// 需要注册用户
		if regFlag {
			regInput := &model.RegisterInput{}

			regInput.UserAccount = bindId
			regInput.Password = "1231Ss@123" + gtime.TimestampMicroStr()
			regInput.Encrypt = false
			regInput.UserNickname = userInfoRow.BindNickname.(string)

			if activityId != 0 {
				regInput.ActivityId = activityId
			}

			regInput.BindType = consts.ACCOUNT

			newUserId, err := service.Login().DoRegister(ctx, regInput)
			if err != nil || newUserId == 0 {
				return 0, gerror.New("注册失败！")
			}

			userBase, err := dao.UserBase.Get(ctx, newUserId)
			if err != nil {
				return 0, err
			}

			if userBase != nil {
				userId = userBase.UserId

				// 添加user info
				userInfo := &do.UserInfo{
					UserAccount:  userInfoRow.BindId,
					UserNickname: userInfoRow.BindNickname,
					UserAvatar:   userInfoRow.BindIcon,
				}

				if _, err := dao.UserInfo.Edit(ctx, userId, userInfo); err != nil {
					return 0, err
				}

				s.checkBind(ctx, bindId, userInfoRow.BindType.(uint), userId, userInfoRow)
			}
		} else {
			//读取用户信息
			userId = userInfoRow.UserId.(uint)
			if _, err := s.checkBind(ctx, bindId, userInfoRow.BindType.(uint), userId, userInfoRow); err != nil {
				return 0, err
			}

		}
	}

	if userId == 0 {
		return 0, gerror.New("保存用户绑定失败！")
	}

	return userId, nil
}

// CheckBind 检查绑定关系
func (s *sUser) checkBind(ctx context.Context, bindId string, bindType uint, userId uint, userInfoRow *do.UserBindConnect) (bool, error) {
	// 从数据库中获取绑定记录
	bindRow, err := dao.UserBindConnect.Get(ctx, bindId)
	if err != nil {
		return false, err
	}

	if bindRow != nil && !g.IsEmpty(bindRow.UserId) {
		// 验证通过, 登录成功
		bindUserId := bindRow.UserId
		if !g.IsEmpty(userId) && bindUserId == userId {
			return false, errors.New("非法请求, 已经登录用户不应该访问到此页面 - 重复绑定")
		} else if g.IsEmpty(userId) && bindUserId == userId {
			return false, errors.New("非法请求, 错误绑定数据")
		}
	} else if bindRow != nil && g.IsEmpty(bindRow.UserId) {
		userInfoRow.UserId = userId
		userInfoRow.BindId = bindId

		if _, err := dao.UserBindConnect.Save(ctx, userInfoRow); err != nil {
			return false, errors.New("保存绑定信息失败")
		}
	} else if bindRow == nil {
		// TODO: 头像会变动，需要本地缓存，生成新网址
		userInfoRow.BindId = bindId
		userInfoRow.BindType = bindType
		userInfoRow.UserId = userId
		userInfoRow.BindActive = true

		if _, err := dao.UserBindConnect.Save(ctx, userInfoRow); err != nil {
			return false, errors.New("保存绑定信息失败")
		}
	}

	return true, nil
}

// SaveCertificate 保存用户认证信息
//
// @param ctx      上下文
// @param userInfo 用户信息
// @return
func (s *sUser) SaveCertificate(ctx context.Context, userInfo *do.UserInfo) (bool, error) {
	// 获取用户信息
	info, err := dao.UserInfo.Get(ctx, userInfo.UserId)
	if err != nil {
		return false, gerror.New(g.I18n().Translate(ctx, "用户信息不存在！"))
	}

	if info == nil {
		return false, gerror.New(g.I18n().Translate(ctx, "请先登录！"))
	}

	// 检查用户认证状态
	if info.UserIsAuthentication != consts.USER_CERTIFICATION_VERIFY {
		// 更新用户信息
		_, err = dao.UserInfo.Edit(ctx, userInfo.UserId, userInfo)
		if err != nil {
			return false, gerror.New(g.I18n().Translate(ctx, "保存用户数据失败！"))
		}
	} else {
		return false, gerror.New(g.I18n().Translate(ctx, "已提交，请勿重复提交！"))
	}

	return true, nil
}
