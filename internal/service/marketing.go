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

package service

import (
	"context"

	"golershop.cn/api/marketing"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
)

type (
	IActivityBase interface {
		// Get 根据编号读取活动信息
		Get(ctx context.Context, id any) (out *entity.ActivityBase, err error)
		// Gets 根据编号读取读取多条活动信息
		Gets(ctx context.Context, id any) (list []*entity.ActivityBase, err error)
		// Find 查询活动数据
		Find(ctx context.Context, in *do.ActivityBaseListInput) (out []*entity.ActivityBase, err error)
		// List 分页读取活动
		List(ctx context.Context, in *do.ActivityBaseListInput) (out *model.ActivityListOutput, err error)
		// Add 新增活动
		Add(ctx context.Context, in *do.ActivityBase) (lastInsertId int64, err error)
		// Edit 编辑活动
		Edit(ctx context.Context, in *do.ActivityBase) (affected int64, err error)
		// Remove 删除活动记录
		Remove(ctx context.Context, id any) (affected int64, err error)
		// ListVoucher 活动表-优惠券列表
		ListVoucher(ctx context.Context, input *do.ActivityBaseListInput) (res *model.ActivityListOutput, err error)
		// AddActivityBase 新增活动基础信息
		AddActivityBase(ctx context.Context, activityBase *marketing.ActivityBaseAddReq) (bool, error)
		// GetList 获取活动列表
		GetList(ctx context.Context, activityBaseListReq *do.ActivityBaseListInput) (activityBaseResPage *model.ActivityListOutput, err error)
		// UpdateActivityBase 更新活动基础信息
		UpdateActivityBase(ctx context.Context, activityBase *marketing.ActivityBaseEditReq) (bool, error)
		FixActivityData(ctx context.Context, activityBaseList []*entity.ActivityBase) ([]*entity.ActivityBase, error)
		// EditActivityBase 编辑活动基础信息
		EditActivityBase(ctx context.Context, activityId uint, data *do.ActivityBase) (bool, error)
	}
	IActivityItem interface {
		// Get 根据编号读取活动信息
		Get(ctx context.Context, id any) (out *entity.ActivityItem, err error)
		// Gets 根据编号读取读取多条活动信息
		Gets(ctx context.Context, id any) (list []*entity.ActivityItem, err error)
		// Find 查询活动数据
		Find(ctx context.Context, in *do.ActivityItemListInput) (out []*entity.ActivityItem, err error)
		// List 分页读取活动
		List(ctx context.Context, in *do.ActivityItemListInput) (out *do.ActivityItemListOutput, err error)
		// Add 新增活动
		Add(ctx context.Context, in *do.ActivityItem) (lastInsertId int64, err error)
		// Edit 编辑活动
		Edit(ctx context.Context, in *do.ActivityItem) (affected int64, err error)
		// Remove 删除活动记录
		Remove(ctx context.Context, id any) (affected int64, err error)
		// GetActivityInfo 获取活动信息
		GetActivityInfo(ctx context.Context, itemIds []uint64) (output []*model.ActivityInfoVo, err error)
	}
)

var (
	localActivityBase IActivityBase
	localActivityItem IActivityItem
)

func ActivityItem() IActivityItem {
	if localActivityItem == nil {
		panic("implement not found for interface IActivityItem, forgot register?")
	}
	return localActivityItem
}

func RegisterActivityItem(i IActivityItem) {
	localActivityItem = i
}

func ActivityBase() IActivityBase {
	if localActivityBase == nil {
		panic("implement not found for interface IActivityBase, forgot register?")
	}
	return localActivityBase
}

func RegisterActivityBase(i IActivityBase) {
	localActivityBase = i
}
