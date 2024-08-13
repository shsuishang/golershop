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

package marketing

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
	"strings"
)

type sActivityItem struct{}

func init() {
	service.RegisterActivityItem(NewActivityItem())
}

func NewActivityItem() *sActivityItem {
	return &sActivityItem{}
}

// Get 根据编号读取活动信息
func (s *sActivityItem) Get(ctx context.Context, id any) (out *entity.ActivityItem, err error) {
	var list []*entity.ActivityItem
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
func (s *sActivityItem) Gets(ctx context.Context, id any) (list []*entity.ActivityItem, err error) {
	err = dao.ActivityItem.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询活动数据
func (s *sActivityItem) Find(ctx context.Context, in *do.ActivityItemListInput) (out []*entity.ActivityItem, err error) {
	out, err = dao.ActivityItem.Find(ctx, in)

	return out, err
}

// List 分页读取活动
func (s *sActivityItem) List(ctx context.Context, in *do.ActivityItemListInput) (out *do.ActivityItemListOutput, err error) {
	out, err = dao.ActivityItem.List(ctx, in)

	return out, err
}

// Add 新增活动
func (s *sActivityItem) Add(ctx context.Context, in *do.ActivityItem) (lastInsertId int64, err error) {
	lastInsertId, err = dao.ActivityItem.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑活动
func (s *sActivityItem) Edit(ctx context.Context, in *do.ActivityItem) (affected int64, err error) {
	_, err = dao.ActivityItem.Edit(ctx, in.ActivityId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除活动记录
func (s *sActivityItem) Remove(ctx context.Context, id any) (affected int64, err error) {
	row, err := s.Get(ctx, id)

	if err != nil {
		return 0, err
	}

	if row.ActivityItemState == consts.ACTIVITY_STATE_NORMAL {
		return 0, errors.New(fmt.Sprintf("活动:%d 生效使用中，不可删除", row.ActivityId))
	}

	affected, err = dao.ActivityItem.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// GetActivityInfo 获取活动信息
func (s *sActivityItem) GetActivityInfo(ctx context.Context, itemIds []uint64) (output []*model.ActivityInfoVo, err error) {
	if len(itemIds) == 0 {
		return []*model.ActivityInfoVo{}, nil
	}

	// 排他性活动信息
	activityItemList, err := dao.ActivityItem.Find(ctx, &do.ActivityItemListInput{
		Where: do.ActivityItem{ItemId: itemIds, ActivityItemState: consts.ACTIVITY_STATE_NORMAL},
	})
	if err != nil {
		return nil, err
	}

	output = make([]*model.ActivityInfoVo, len(activityItemList))

	for i, item := range activityItemList {
		output[i] = &model.ActivityInfoVo{
			ActivityItem: *item,
			ActivityBase: &entity.ActivityBase{},
		}
		gconv.Struct(item, output[i])
	}

	activityItemIds := array.Column(output, "ItemId")

	// 获取活动基础信息
	activityIds := array.Column(output, "ActivityId")
	activityBaseList, err := dao.ActivityBase.Gets(ctx, activityIds)
	if err != nil {
		return nil, err
	}

	for _, activityBase := range activityBaseList {
		for _, item := range output {
			if item.ActivityId == activityBase.ActivityId {
				item.ActivityBase = activityBase
				break
			}
		}
	}

	// 非排他性活动
	for _, activityBase := range activityBaseList {
		ids := gconv.SliceUint64(strings.Split(activityBase.ActivityItemIds, ","))

		for _, itemId := range itemIds {
			if !array.InArray(activityItemIds, itemId) && array.InArray(ids, itemId) {
				activityInfoVo := &model.ActivityInfoVo{
					ActivityBase: activityBase,
					ActivityItem: entity.ActivityItem{
						ActivityTypeId: activityBase.ActivityTypeId,
						ActivityId:     activityBase.ActivityId,
						ItemId:         itemId,
					},
				}
				output = append(output, activityInfoVo)
			}
		}
	}

	return output, nil
}
