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
	IActivityCutprice interface {
		// Get 根据编号读取活动信息
		Get(ctx context.Context, id any) (out *entity.ActivityCutprice, err error)
		// Gets 根据编号读取读取多条活动信息
		Gets(ctx context.Context, id any) (list []*entity.ActivityCutprice, err error)
		// Find 查询活动数据
		Find(ctx context.Context, in *do.ActivityCutpriceListInput) (out []*entity.ActivityCutprice, err error)
		// FindOne 查询活动数据
		FindOne(ctx context.Context, in *do.ActivityCutpriceListInput) (out *entity.ActivityCutprice, err error)
		// List 分页读取活动
		List(ctx context.Context, in *do.ActivityCutpriceListInput) (out *do.ActivityCutpriceListOutput, err error)
		// Add 新增活动
		Add(ctx context.Context, in *do.ActivityCutprice) (lastInsertId int64, err error)
		// Edit 编辑活动
		Edit(ctx context.Context, in *do.ActivityCutprice) (affected int64, err error)
		// Remove 删除活动记录
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IActivityCutpriceHistory interface {
		// Get 根据编号读取活动信息
		Get(ctx context.Context, id any) (out *entity.ActivityCutpriceHistory, err error)
		// Gets 根据编号读取读取多条活动信息
		Gets(ctx context.Context, id any) (list []*entity.ActivityCutpriceHistory, err error)
		// Find 查询活动数据
		Find(ctx context.Context, in *do.ActivityCutpriceHistoryListInput) (out []*entity.ActivityCutpriceHistory, err error)
		// List 分页读取活动
		List(ctx context.Context, in *do.ActivityCutpriceHistoryListInput) (out *do.ActivityCutpriceHistoryListOutput, err error)
		// Add 新增活动
		Add(ctx context.Context, in *do.ActivityCutpriceHistory) (lastInsertId int64, err error)
		// Edit 编辑活动
		Edit(ctx context.Context, in *do.ActivityCutpriceHistory) (affected int64, err error)
		// Remove 删除活动记录
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IActivityGroupbooking interface {
		// Get 根据编号读取活动信息
		Get(ctx context.Context, id any) (out *entity.ActivityGroupbooking, err error)
		// Gets 根据编号读取读取多条活动信息
		Gets(ctx context.Context, id any) (list []*entity.ActivityGroupbooking, err error)
		// Find 查询活动数据
		Find(ctx context.Context, in *do.ActivityGroupbookingListInput) (out []*entity.ActivityGroupbooking, err error)
		// List 分页读取活动
		List(ctx context.Context, in *do.ActivityGroupbookingListInput) (out *do.ActivityGroupbookingListOutput, err error)
		// Add 新增活动
		Add(ctx context.Context, in *do.ActivityGroupbooking) (lastInsertId int64, err error)
		// Edit 编辑活动
		Edit(ctx context.Context, in *do.ActivityGroupbooking) (affected int64, err error)
		// Remove 删除活动记录
		Remove(ctx context.Context, id any) (affected int64, err error)
		// DoGroupbooking 参加拼团活动
		DoGroupbooking(ctx context.Context, orderId string, userId uint, gbId uint, activityInfo *model.ActivityInfoVo) (id uint, err error)
		// DoGroupbooking 参加拼团活动
		SetPaidYes(ctx context.Context, orderId string, userId uint) (flag bool, err error)
		// CheckGroupbookingSuccess
		CheckGroupbookingSuccess(ctx context.Context, orderId string) (flag bool, err error)
		// ListsUserGroupbookingHistory 拼团列表
		ListsUserGroupbookingHistory(ctx context.Context, in *do.ActivityGroupbookingHistoryListInput) (res *marketing.ActivityGroupbookingHistoryRes, err error)
		// GetUserGroupbooking 拼团详情
		GetUserGroupbooking(ctx context.Context, activityGroupbookingReq *marketing.ActivityGroupbookingReq) (activityGroupbookingRes *marketing.ActivityGroupbookingRes, err error)
	}
	IActivityGroupbookingHistory interface {
		// Get 根据编号读取活动信息
		Get(ctx context.Context, id any) (out *entity.ActivityGroupbookingHistory, err error)
		// Gets 根据编号读取读取多条活动信息
		Gets(ctx context.Context, id any) (list []*entity.ActivityGroupbookingHistory, err error)
		// Find 查询活动数据
		Find(ctx context.Context, in *do.ActivityGroupbookingHistoryListInput) (out []*entity.ActivityGroupbookingHistory, err error)
		// List 分页读取活动
		List(ctx context.Context, in *do.ActivityGroupbookingHistoryListInput) (out *do.ActivityGroupbookingHistoryListOutput, err error)
		// Add 新增活动
		Add(ctx context.Context, in *do.ActivityGroupbookingHistory) (lastInsertId int64, err error)
		// Edit 编辑活动
		Edit(ctx context.Context, in *do.ActivityGroupbookingHistory) (affected int64, err error)
		// Remove 删除活动记录
		Remove(ctx context.Context, id any) (affected int64, err error)
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
		// getActivityItemNum 获取非排他活动商品及数量
		GetActivityItemNum(ctx context.Context, activityBase *entity.ActivityBase) (map[uint64]*model.ItemNumVo, error)
		// ListVoucher 活动表-优惠券列表
		ListVoucher(ctx context.Context, input *do.ActivityBaseListInput) (res *model.ActivityListOutput, err error)
	}
)

var (
	localActivityBase                IActivityBase
	localActivityCutprice            IActivityCutprice
	localActivityCutpriceHistory     IActivityCutpriceHistory
	localActivityGroupbooking        IActivityGroupbooking
	localActivityGroupbookingHistory IActivityGroupbookingHistory
	localActivityItem                IActivityItem
)

func ActivityBase() IActivityBase {
	if localActivityBase == nil {
		panic("implement not found for interface IActivityBase, forgot register?")
	}
	return localActivityBase
}

func RegisterActivityBase(i IActivityBase) {
	localActivityBase = i
}

func ActivityCutprice() IActivityCutprice {
	if localActivityCutprice == nil {
		panic("implement not found for interface IActivityCutprice, forgot register?")
	}
	return localActivityCutprice
}

func RegisterActivityCutprice(i IActivityCutprice) {
	localActivityCutprice = i
}

func ActivityCutpriceHistory() IActivityCutpriceHistory {
	if localActivityCutpriceHistory == nil {
		panic("implement not found for interface IActivityCutpriceHistory, forgot register?")
	}
	return localActivityCutpriceHistory
}

func RegisterActivityCutpriceHistory(i IActivityCutpriceHistory) {
	localActivityCutpriceHistory = i
}

func ActivityGroupbooking() IActivityGroupbooking {
	if localActivityGroupbooking == nil {
		panic("implement not found for interface IActivityGroupbooking, forgot register?")
	}
	return localActivityGroupbooking
}

func RegisterActivityGroupbooking(i IActivityGroupbooking) {
	localActivityGroupbooking = i
}

func ActivityGroupbookingHistory() IActivityGroupbookingHistory {
	if localActivityGroupbookingHistory == nil {
		panic("implement not found for interface IActivityGroupbookingHistory, forgot register?")
	}
	return localActivityGroupbookingHistory
}

func RegisterActivityGroupbookingHistory(i IActivityGroupbookingHistory) {
	localActivityGroupbookingHistory = i
}

func ActivityItem() IActivityItem {
	if localActivityItem == nil {
		panic("implement not found for interface IActivityItem, forgot register?")
	}
	return localActivityItem
}

func RegisterActivityItem(i IActivityItem) {
	localActivityItem = i
}
