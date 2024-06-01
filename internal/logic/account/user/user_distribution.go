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
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
	"time"
)

type sUserDistribution struct{}

func init() {
	service.RegisterUserDistribution(NewUserDistribution())
}

func NewUserDistribution() *sUserDistribution {
	return &sUserDistribution{}
}

// Find 查询数据
func (s *sUserDistribution) Find(ctx context.Context, in *do.UserDistributionListInput) (out []*entity.UserDistribution, err error) {
	out, err = dao.UserDistribution.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sUserDistribution) List(ctx context.Context, in *do.UserDistributionListInput) (out *do.UserDistributionListOutput, err error) {
	out, err = dao.UserDistribution.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sUserDistribution) Add(ctx context.Context, in *do.UserDistribution) (lastInsertId int64, err error) {
	lastInsertId, err = dao.UserDistribution.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sUserDistribution) Edit(ctx context.Context, in *do.UserDistribution) (affected int64, err error) {
	_, err = dao.UserDistribution.Edit(ctx, in.UserId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sUserDistribution) Remove(ctx context.Context, id any) (affected int64, err error) {
	affected, err = dao.UserDistribution.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// GetList 获取用户分销列表
func (s *sUserDistribution) GetList(ctx context.Context, input *do.UserDistributionListInput) (*model.UserDistributionOutput, error) {
	var output *model.UserDistributionOutput
	lists, err := s.List(ctx, input)
	if err != nil {
		return nil, err
	}

	if lists != nil && !g.IsEmpty(lists.Items) {
		if err := gconv.Scan(lists, &output); err != nil {
			return nil, err
		}

		err = gconv.Struct(lists.Items, &output.Items)
		userIds := array.Column(output.Items, "UserId")

		if !g.IsEmpty(userIds) {
			// 增加用户昵称
			userInfos, err := dao.UserInfo.Gets(ctx, userIds)
			if err != nil {
				return nil, err
			}

			nickNameMap := make(map[uint]string)
			if !g.IsEmpty(userInfos) {
				for _, userInfo := range userInfos {
					nickNameMap[userInfo.UserId] = userInfo.UserNickname
				}
			}

			distributionCommissionList, err := dao.DistributionCommission.Gets(ctx, userIds)
			if err != nil {
				return nil, err
			}

			for i := range output.Items {
				// 增加用户昵称
				if len(nickNameMap) > 0 {
					output.Items[i].UserNickname = nickNameMap[output.Items[i].UserId]
				}

				for _, distributionCommission := range distributionCommissionList {
					if output.Items[i].UserId == distributionCommission.UserId {
						output.Items[i].CommissionAmount = distributionCommission.CommissionAmount
						output.Items[i].CommissionBuyAmount0 = distributionCommission.CommissionBuyAmount0
						output.Items[i].CommissionBuyAmount1 = distributionCommission.CommissionBuyAmount1
						output.Items[i].CommissionBuyAmount2 = distributionCommission.CommissionBuyAmount2
						output.Items[i].CommissionSettled = distributionCommission.CommissionSettled
					}
				}
			}
		}
	}

	return output, nil
}

// GetChildNum 获取子用户数量
func (s *sUserDistribution) GetChildNum(ctx context.Context, userId uint, startTime, endTime int64) (int, error) {
	queryWrapper := do.UserDistribution{
		UserParentId: userId,
	}

	m := dao.UserDistribution.Ctx(ctx).Where(queryWrapper)

	if startTime > 0 || endTime > 0 {
		m = m.WhereBetween("UserTime", startTime, endTime)
	}

	count, err := m.Count()

	if err != nil {
		return 0, err
	}

	return count, nil
}

// InitDistributionUser 添加
// 添加分销用户记录 - 推广员记录
func (s *sUserDistribution) InitDistributionUser(ctx context.Context, userParentId uint, userActive bool) bool {
	if userParentId != 0 {
		// 判断用户是否存在
		userBase, err := dao.UserBase.Get(ctx, userParentId)
		if err != nil || userBase == nil {
			return false
		}

		// 添加父收益表，判断
		userCommission, err := dao.DistributionCommission.Get(ctx, userParentId)
		if err != nil {
			return false
		}
		if userCommission == nil {
			distributionCommission := &do.DistributionCommission{
				UserId: userParentId,
			}

			if _, err := dao.DistributionCommission.Save(ctx, distributionCommission); err != nil {
				return false
			}
		}

		// 初始化推销员记录
		distributionUserRow, err := dao.UserDistribution.Get(ctx, userParentId)
		if err != nil {
			return false
		}

		if distributionUserRow == nil {
			userDistribution := &do.UserDistribution{
				UserId:      userParentId,
				UserActive:  userActive,
				UserTime:    time.Now().UnixMilli(),
				UserFansNum: 1,
			}
			if _, err := dao.UserDistribution.Save(ctx, userDistribution); err != nil {
				return false
			}
		} else {
			userDistribution := &do.UserDistribution{
				UserId:      userParentId,
				UserFansNum: distributionUserRow.UserFansNum + 1,
			}
			if _, err := dao.UserDistribution.Save(ctx, userDistribution); err != nil {
				return false
			}
		}
	}

	return true
}

// 添加用户关系
func (s *sUserDistribution) AddPlantformUser(ctx context.Context, plantformUser *do.UserDistribution) bool {
	if _, err := dao.UserDistribution.Save(ctx, plantformUser); err != nil {
		return false
	}

	// 初始化推广员记录
	userActive := service.ConfigBase().GetBool(ctx, "distribution_user_auto_active", false)
	userParentId := plantformUser.UserParentId
	s.InitDistributionUser(ctx, userParentId.(uint), userActive)

	// 一级
	userRow := &do.DistributionGeneratedCommission{
		UgcId:        fmt.Sprintf("%d-%d-%d", plantformUser.UserId, plantformUser.UserParentId, 1),
		UserId:       plantformUser.UserId,
		UserParentId: plantformUser.UserParentId,
		UgcLevel:     1,
		UserTime:     plantformUser.UserTime,
	}
	if _, err := dao.DistributionGeneratedCommission.Save(ctx, userRow); err != nil {
		return false
	}

	// 二级
	plantformRow, err := dao.UserDistribution.Get(ctx, plantformUser.UserParentId)
	if err == nil && plantformRow.UserParentId != 0 {
		userRow.UgcId = fmt.Sprintf("%d-%d-%d", plantformUser.UserId, plantformRow.UserParentId, 2)
		userRow.UserParentId = plantformRow.UserParentId
		userRow.UgcLevel = 2
		if _, err := dao.DistributionGeneratedCommission.Save(ctx, userRow); err != nil {
			return false
		}

		// 三级
		plantUserParent, err := dao.UserDistribution.Get(ctx, plantformRow.UserParentId)
		if err == nil && plantUserParent.UserParentId != 0 {
			userRow.UgcId = fmt.Sprintf("%d-%d-%d", plantformUser.UserId, plantUserParent.UserParentId, 3)
			userRow.UserParentId = plantUserParent.UserParentId
			userRow.UgcLevel = 3
			if _, err := dao.DistributionGeneratedCommission.Save(ctx, userRow); err != nil {
				return false
			}
		}
	}

	/*
	   // 统计数据 检测升级
	   userSns := &entity.AccountUserSns{
	       UserId:   userParentId,
	       UserFans: 1,
	   }

	   if !accountService.SaveAndCheckUpdateUserLevel(userSns) {
	       return false
	   }
	*/

	return true
}

// 添加分销用户
func (s *sUserDistribution) AddDistribution(ctx context.Context, userDistribution *do.UserDistribution) bool {
	userId := userDistribution.UserId
	userInfo, err := dao.UserInfo.Get(ctx, userId)
	if err != nil {
		return false
	}

	if userInfo == nil {
		return false
	}

	userIsPa := userDistribution.UserIsPa
	if !g.IsEmpty(userIsPa) {
		_, err := dao.UserDistribution.FindOne(ctx, &do.UserDistributionListInput{Where: do.UserDistribution{UserIsPa: userIsPa}})
		if err == nil {
			//判断该省是否设置代理
			return false
		}
	}

	userIsCa := userDistribution.UserIsCa
	if !g.IsEmpty(userIsCa) {
		_, err := dao.UserDistribution.List(ctx, &do.UserDistributionListInput{Where: do.UserDistribution{UserIsCa: userIsCa}})
		if err == nil {
			//判断该市是否设置代理
			return false
		}
	}

	userIsDa := userDistribution.UserIsDa
	if !g.IsEmpty(userIsDa) {
		_, err := dao.UserDistribution.List(ctx, &do.UserDistributionListInput{Where: do.UserDistribution{UserIsDa: userIsDa}})
		if err == nil {
			//判断该区是否设置代理
			return false
		}
	}

	distributionCommission, err := dao.DistributionCommission.Get(ctx, userId)
	if distributionCommission != nil {
		commission := &do.DistributionCommission{
			UserId: userId,
		}
		if _, err := dao.DistributionCommission.Save(ctx, commission); err != nil {
			return false
		}
	}

	distribution, err := dao.UserDistribution.Get(ctx, userId)
	if distribution != nil {
		userDistribution.UserTime = uint64(time.Now().UnixMilli())
	}

	_, err = dao.UserDistribution.Save(ctx, userDistribution)

	return err == nil
}
