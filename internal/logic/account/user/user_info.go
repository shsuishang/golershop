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
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"time"
)

type sUserInfo struct{}

func init() {
	service.RegisterUserInfo(NewuserInfo())
}

func NewuserInfo() *sUserInfo {
	return &sUserInfo{}
}

// Get 根据编号读取活动信息
func (s *sUserInfo) Get(ctx context.Context, id any) (out *entity.UserInfo, err error) {
	var list []*entity.UserInfo
	list, err = s.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// Gets 根据编号读取读取多条活动信息
func (s *sUserInfo) Gets(ctx context.Context, id any) (list []*entity.UserInfo, err error) {
	err = dao.UserInfo.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sUserInfo) Find(ctx context.Context, in *do.UserInfoListInput) (out []*entity.UserInfo, err error) {
	out, err = dao.UserInfo.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sUserInfo) List(ctx context.Context, in *do.UserInfoListInput) (out *do.UserInfoListOutput, err error) {
	out, err = dao.UserInfo.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sUserInfo) Add(ctx context.Context, in *do.UserInfo) (lastInsertId int64, err error) {
	lastInsertId, err = dao.UserInfo.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sUserInfo) Edit(ctx context.Context, in *do.UserInfo) (affected int64, err error) {
	_, err = dao.UserInfo.Edit(ctx, in.UserId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sUserInfo) Remove(ctx context.Context, id any) (affected int64, err error) {
	affected, err = dao.UserInfo.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// PassWordEdit 修改密码
func (s *sUserInfo) PassWordEdit(ctx context.Context, userId uint, userPassword string) (bool, error) {
	if g.IsEmpty(userId) {
		return false, gerror.New("用户Id为空")
	}

	if g.IsEmpty(userPassword) {
		return false, gerror.New("密码为空")
	}

	// 修改密码
	userSalt := grand.S(32) // 生成随机盐值
	resetPassWord := gmd5.MustEncryptString(userPassword + userSalt)

	userBase := &do.UserBase{
		UserId:       userId,
		UserSalt:     userSalt,
		UserPassword: resetPassWord,
	}

	_, err := dao.UserBase.Edit(ctx, userId, userBase)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetUserData 获取用户详细信息
func (s *sUserInfo) GetUserData(ctx context.Context, userId uint) (userInfoOutput *model.UserInfoOutput, err error) {
	userInfoOutput = &model.UserInfoOutput{}

	// 用户基本信息
	userBase, err := dao.UserBase.Get(ctx, userId)
	if err != nil {
		return nil, err
	}

	gconv.Scan(userBase, userInfoOutput)

	// 用户详情信息
	userInfo, err := dao.UserInfo.Get(ctx, userId)
	if err != nil {
		return nil, err
	}

	gconv.Scan(userInfo, userInfoOutput)

	// 身份证图片
	if !g.IsEmpty(userInfo.UserIdcardImages) {
		userInfoOutput.UserIdcardImageList = gstr.Split(userInfo.UserIdcardImages, ",")
	}

	// 用户等级
	userLevel, err := dao.UserLevel.Get(ctx, userInfo.UserLevelId)
	if err == nil && userLevel != nil {
		userInfoOutput.UserLevelName = userLevel.UserLevelName
	}

	// 用户标签、分组
	if !g.IsEmpty(userInfo.TagIds) {
		tagIds := gconv.SliceInt(gstr.Split(userInfo.TagIds, ","))
		userTagBases, err := dao.UserTagBase.Gets(ctx, tagIds)
		if err == nil && len(userTagBases) > 0 {
			tagNames := make([]string, len(userTagBases))
			for i, tag := range userTagBases {
				tagNames[i] = tag.TagTitle
			}
			userInfoOutput.TagTitleList = tagNames
			userInfoOutput.TagTitles = gstr.Join(tagNames, "、")

			tagGroupIds := make([]uint, len(userTagBases))
			for i, tag := range userTagBases {
				tagGroupIds[i] = tag.TagGroupId
			}

			tagGroupGroups, err := dao.UserTagGroup.Gets(ctx, tagGroupIds)
			if err == nil && len(tagGroupGroups) > 0 {
				groupNames := make([]string, len(tagGroupGroups))
				for i, group := range tagGroupGroups {
					groupNames[i] = group.TagGroupName
				}
				userInfoOutput.TagGroupNames = gstr.Join(groupNames, "、")
			}
		}
	}

	// 统计没有取消的订单
	orderStates := []uint{consts.ORDER_STATE_WAIT_PAY, consts.ORDER_STATE_WAIT_PAID, consts.ORDER_STATE_WAIT_REVIEW, consts.ORDER_STATE_WAIT_FINANCE_REVIEW, consts.ORDER_STATE_PICKING, consts.ORDER_STATE_WAIT_SHIPPING, consts.ORDER_STATE_SHIPPED, consts.ORDER_STATE_RECEIVED, consts.ORDER_STATE_FINISH, consts.ORDER_STATE_SELF_PICKUP}
	month := GetMonthTimeRange()
	// 本月订单
	orderNum, err := dao.AnalyticsOrder.GetOrderNum(ctx, month.Start, month.End, orderStates, nil, userId, 0)
	if err == nil && orderNum != nil {
		userInfoOutput.MonthOrder = orderNum
	}

	// 总计订单
	totalNum, err := dao.AnalyticsOrder.GetOrderNum(ctx, 0, 0, orderStates, nil, userId, 0)
	if err == nil && totalNum != nil {
		userInfoOutput.TotalOrder = totalNum
	}

	// 本月消费金额
	tradeAmount, err := dao.AnalyticsTrade.SalesAmount(ctx, month.Start, month.End, int64(userId))
	if err == nil && tradeAmount != nil {
		userInfoOutput.MonthTrade = tradeAmount
	}

	// 总消费金额
	totalAmount, err := dao.AnalyticsTrade.SalesAmount(ctx, 0, 0, int64(userId))
	if err == nil && totalAmount != nil {
		userInfoOutput.TotalTrade = totalAmount
	}

	// 用户地址
	input := &do.UserDeliveryAddressListInput{
		Where: do.UserDeliveryAddress{
			UserId: userId,
		},
	}
	address, err := dao.UserDeliveryAddress.FindOne(ctx, input)
	if err == nil && address != nil {
		userInfoOutput.UdAddress = address.UdProvince + address.UdCity + address.UdCounty + address.UdAddress
	}

	// 用户资源
	userResource, err := dao.UserResource.Get(ctx, userId)
	if err == nil && userResource != nil {
		gconv.Scan(userResource, userInfoOutput)
	}

	// 用户登录信息
	userLogin, err := dao.UserLogin.Get(ctx, userId)
	if err == nil && userLogin != nil {
		userInfoOutput.UserRegTime = userLogin.UserRegTime
	}

	// 推广员信息
	userDistribution, err := dao.UserDistribution.Get(ctx, userId)
	if err == nil && userDistribution != nil {
		userInfoOutput.UserParentId = userDistribution.UserParentId
	}

	// 登录时间
	loginHistoryPage, err := dao.UserLoginHistory.List(ctx, &do.UserLoginHistoryListInput{
		Where: do.UserLoginHistory{
			UserId: userId,
		},
		BaseList: ml.BaseList{
			Page: 1,
			Size: 1,
			Sidx: dao.UserLoginHistory.Columns().UserLoginTime,
			Sort: "DESC",
		},
	})
	if err == nil && loginHistoryPage != nil && len(loginHistoryPage.Items) > 0 {
		userInfoOutput.UserLoginTime = loginHistoryPage.Items[0].UserLoginTime
	}

	// 累计佣金
	distributionCommission, err := dao.DistributionCommission.Get(ctx, userId)
	if err == nil && distributionCommission != nil {
		userInfoOutput.UserCommissionNow = distributionCommission.CommissionAmount - distributionCommission.CommissionSettled
	}

	// 本月佣金
	monthCommission, _ := dao.DistributionCommission.CalCommission(ctx, userId, 0, month.Start, month.End, 0, 0)
	userInfoOutput.MonthCommissionBuy = gconv.Float64(monthCommission)

	return userInfoOutput, nil
}

// TimeRange 结构体表示时间范围
type TimeRange struct {
	Start int64 // 起始时间，Unix时间戳（毫秒）
	End   int64 // 结束时间，Unix时间戳（毫秒）
}

// GetMonthTimeRange 获取当前月份的时间范围
func GetMonthTimeRange() TimeRange {
	now := time.Now()
	year, month, _ := now.Date()
	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Millisecond) // 当前月的最后一毫秒

	// 转换为Unix时间戳（毫秒）
	startUnixMilli := startOfMonth.UnixNano() / int64(time.Millisecond)
	endUnixMilli := endOfMonth.UnixNano() / int64(time.Millisecond)

	return TimeRange{
		Start: startUnixMilli,
		End:   endUnixMilli,
	}
}
