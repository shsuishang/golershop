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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
	"sort"
	"strings"
	"time"
)

type sActivityBase struct{}

func init() {
	service.RegisterActivityBase(NewActivityBase())
}

func NewActivityBase() *sActivityBase {
	return &sActivityBase{}
}

// Get 根据编号读取活动信息
func (s *sActivityBase) Get(ctx context.Context, id any) (out *entity.ActivityBase, err error) {
	var list []*entity.ActivityBase
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
func (s *sActivityBase) Gets(ctx context.Context, id any) (list []*entity.ActivityBase, err error) {
	err = dao.ActivityBase.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询活动数据
func (s *sActivityBase) Find(ctx context.Context, in *do.ActivityBaseListInput) (out []*entity.ActivityBase, err error) {
	out, err = dao.ActivityBase.Find(ctx, in)

	return out, err
}

// List 分页读取活动
func (s *sActivityBase) List(ctx context.Context, in *do.ActivityBaseListInput) (out *model.ActivityListOutput, err error) {
	activityBaseList, err := dao.ActivityBase.List(ctx, in)

	gconv.Scan(activityBaseList, &out)

	return out, err
}

// Add 新增活动
func (s *sActivityBase) Add(ctx context.Context, in *do.ActivityBase) (lastInsertId int64, err error) {
	lastInsertId, err = dao.ActivityBase.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑活动
func (s *sActivityBase) Edit(ctx context.Context, in *do.ActivityBase) (affected int64, err error) {
	_, err = dao.ActivityBase.Edit(ctx, in.ActivityId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除活动记录
func (s *sActivityBase) Remove(ctx context.Context, id any) (affected int64, err error) {
	row, err := s.Get(ctx, id)

	if err != nil {
		return 0, err
	}

	if row.ActivityState == consts.ACTIVITY_STATE_NORMAL {
		return 0, errors.New(fmt.Sprintf("活动:%d 生效使用中，不可删除", row.ActivityId))
	}

	affected, err = dao.ActivityBase.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// getActivityItemNum 获取非排他活动商品及数量
func (s *sActivityBase) GetActivityItemNum(ctx context.Context, activityBase *entity.ActivityBase) (map[uint64]*model.ItemNumVo, error) {
	itemNumVoMap := make(map[uint64]*model.ItemNumVo)

	if activityBase != nil {
		activityRule := activityBase.ActivityRule

		if activityRule != "" {
			var activityRuleVo model.ActivityRuleVo
			err := json.Unmarshal([]byte(activityRule), &activityRuleVo)
			if err != nil {
				return nil, err
			}

			activityTypeId := activityBase.ActivityTypeId

			if activityTypeId == consts.ACTIVITY_TYPE_CUTPRICE {
				cutprice := activityRuleVo.Cutprice
				if cutprice.CutpriceQuantity < activityBase.ActivityEffectiveQuantity {
					return nil, errors.New("活动库存不足!")
				}

				if !g.IsEmpty(cutprice) {
					items := cutprice.Items

					for _, item := range items {
						itemNumVoMap[item.ItemId] = item
					}
				}
			} else if activityTypeId == consts.ACTIVITY_TYPE_GIFTBAG {
				giftbag := activityRuleVo.Giftbag
				if giftbag.GiftbagQuantity < activityBase.ActivityEffectiveQuantity {
					return nil, errors.New("活动库存不足!")
				}

				if !g.IsEmpty(giftbag) {
					items := giftbag.Items

					for _, item := range items {
						itemNumVoMap[item.ItemId] = item
					}
				}
			}
		}
	}

	return itemNumVoMap, nil
}

// ListVoucher 活动表-优惠券列表
func (s *sActivityBase) ListVoucher(ctx context.Context, input *do.ActivityBaseListInput) (res *model.ActivityListOutput, err error) {
	res = &model.ActivityListOutput{}

	input.Where.ActivityState = consts.ACTIVITY_STATE_NORMAL
	input.Where.ActivityTypeId = consts.ACTIVITY_TYPE_VOUCHER

	input.BaseList.Sidx = "activity_addtime"
	input.BaseList.Sort = ml.ORDER_BY_DESC

	userId := input.Where.UserId
	// 根据条件查询活动列表
	activityPage, err := service.ActivityBase().List(ctx, input)
	if err != nil {
		return nil, err
	}

	// 若查询结果不为空，则处理返回结果
	if len(activityPage.Items) > 0 {
		// 将活动列表转换为响应结构体
		res.Page = activityPage.Page
		res.Size = activityPage.Size
		res.Total = activityPage.Total

		// 获取活动列表中的店铺ID列表
		//storeIds := array.Column(activityPage.Items, "StoreId")
		// 根据店铺ID列表获取店铺名称
		storeMap := make(map[uint]string)

		// 获取用户领券数量列表
		userVoucherNums, err := dao.UserVoucherNum.Find(ctx, &do.UserVoucherNumListInput{Where: do.UserVoucherNum{UserId: userId}})
		if err != nil {
			return nil, err
		}

		// 将用户领券数量列表转换为以活动ID为键的映射
		userVoucherMap := make(map[uint]uint)
		for _, num := range userVoucherNums {
			userVoucherMap[num.ActivityId] = num.UvnNum
		}

		// 获取会员等级列表
		userLevels, err := dao.UserLevel.Find(ctx, &do.UserLevelListInput{})
		if err != nil {
			return nil, err
		}
		// 将会员等级列表转换为以等级ID为键的映射
		userLevelMap := make(map[uint]string)
		for _, level := range userLevels {
			userLevelMap[level.UserLevelId] = level.UserLevelName
		}

		// 遍历处理每个活动信息
		for _, activity := range activityPage.Items {
			activityRes := &model.ActivityOutput{}
			gconv.Struct(activity, activityRes)

			// 设置店铺名称
			activityRes.StoreName = storeMap[activity.StoreId]

			// 解析活动规则JSON
			var ruleVo model.ActivityRuleVo
			if activity.ActivityRule != "" {
				json.Unmarshal([]byte(activity.ActivityRule), &ruleVo)
				activityRes.ActivityRuleJson = &ruleVo
			}

			// 设置活动是否可领取
			if !g.IsEmpty(ruleVo.Voucher) {
				voucherPreQuantity := ruleVo.Voucher.VoucherPreQuantity
				if userVoucherNum, ok := userVoucherMap[activity.ActivityId]; ok {
					activityRes.IfGain = voucherPreQuantity > userVoucherNum
				} else {
					activityRes.IfGain = true
				}
			} else {
				activityRes.IfGain = false
			}

			// 设置会员等级名称
			var activityUseLevelName []string
			activityUseLevel := gconv.SliceUint(gstr.Split(activity.ActivityUseLevel, ","))

			//进行降序 i大于j 为降序
			sort.Slice(activityUseLevel, func(i, j int) bool {
				return activityUseLevel[i] > activityUseLevel[j]
			})

			for _, userLevel := range activityUseLevel {
				if levelName, ok := userLevelMap[userLevel]; ok {
					activityUseLevelName = append(activityUseLevelName, levelName)
				}
			}
			if len(activityUseLevelName) > 0 {
				activityRes.ActivityUseLevelName = strings.Join(activityUseLevelName, "、") + " 专属"
			}

			res.Items = append(res.Items, activityRes)
		}
	}

	return res, nil
}

// GetList 获取活动列表
func (s *sActivityBase) GetList(ctx context.Context, activityBaseListReq *do.ActivityBaseListInput) (activityBaseResPage *model.ActivityListOutput, err error) {
	activityBaseResPage = &model.ActivityListOutput{}
	activityBaseIPage, err := s.List(ctx, activityBaseListReq)

	if err != nil || g.IsEmpty(activityBaseIPage.Items) {
		return activityBaseResPage, nil
	}

	var activityBaseList []*entity.ActivityBase
	gconv.Scan(activityBaseIPage.Items, &activityBaseList)

	// 修正活动最新状态
	activityBases, err := s.FixActivityData(ctx, activityBaseList)
	if err != nil {
		return nil, err
	}

	var activityBaseRes []*model.ActivityOutput

	for _, activityBase := range activityBases {
		activityRes := &model.ActivityOutput{}
		gconv.Struct(activityBase, activityRes)

		// 会员等级处理
		activityUseLevel := activityBase.ActivityUseLevel
		if len(activityUseLevel) > 0 {
			levelIdsStr := gconv.Strings(gstr.Split(activityUseLevel, ","))
			levelIds := gconv.Ints(levelIdsStr)
			userLevelList, err := dao.UserLevel.Gets(ctx, levelIds)
			if err != nil {
				return nil, err
			}

			if !g.IsEmpty(userLevelList) {
				useLevelList := g.SliceStr{}
				for _, level := range userLevelList {
					useLevelList = append(useLevelList, level.UserLevelName)
				}
				useLevelStr := gstr.Join(useLevelList, ",")
				activityRes.UseLevel = useLevelStr
			}
		}

		activityTypeId := activityBase.ActivityTypeId
		activityRule := activityBase.ActivityRule
		activityRuleVo := &model.ActivityRuleVo{}

		switch activityTypeId {
		case consts.ACTIVITY_TYPE_REDUCTION,
			consts.ACTIVITY_TYPE_FULL_RETURN,
			consts.ACTIVITY_TYPE_PF_GROUPBUY_STORE,
			consts.ACTIVITY_TYPE_REDUCTION_AGAIN,
			consts.ACTIVITY_TYPE_MULTIPLEDISCOUNT,
			consts.ACTIVITY_TYPE_BATDISCOUNT,
			consts.ACTIVITY_TYPE_MULTIPLE_POINTS:
			//consts.ACTIVITY_TYPE_GIFTPACK:
			if len(activityRule) > 0 {
				gjson.DecodeTo(activityRule, activityRuleVo)

				if activityRuleVo != nil {
					requirement := &activityRuleVo.Requirement
					if requirement != nil {
						buy := requirement.Buy
						if buy != nil {
							items, err := service.ProductBase().GetItems(ctx, buy.Item, 0)
							if err != nil {
								return nil, err
							}
							activityRes.Item = items
						}
					}
				}
			}
		case consts.ACTIVITY_TYPE_CUTPRICE:
			if len(activityRule) > 0 {
				gjson.DecodeTo(activityRule, activityRuleVo)

				if activityRuleVo != nil {
					cutprice := &activityRuleVo.Cutprice
					if cutprice != nil {
						items := cutprice.Items
						if !g.IsEmpty(items) {
							// 假设 array.Column 返回的是 []interface{}
							itemIdsInterface := array.Column(items, "ItemId")

							// 将 itemIds 转换为 []uint64
							var itemIds []uint64
							for _, id := range itemIdsInterface {
								if uintID, ok := id.(uint64); ok {
									itemIds = append(itemIds, uintID)
								} else {
									return nil, fmt.Errorf("invalid itemId type: %T", id)
								}
							}

							productItemVos, err := service.ProductBase().GetItems(ctx, itemIds, 0)
							if err != nil {
								return nil, err
							}
							activityRes.Item = productItemVos

							totalPriceInterface := array.Column(productItemVos, "ItemUnitPrice")
							var totalPrice float64
							for _, price := range totalPriceInterface {
								if priceFloat, ok := price.(float64); ok {
									totalPrice += priceFloat
								} else {
									return nil, fmt.Errorf("invalid price type: %T", price)
								}
							}
							cutprice.TotalPrice = totalPrice
						}
					}
				}
			}
		case consts.ACTIVITY_TYPE_GIFTBAG:
			if len(activityRule) > 0 {
				gjson.DecodeTo(activityRule, activityRuleVo)
				if activityRuleVo != nil {
					giftbag := &activityRuleVo.Giftbag
					if giftbag != nil {
						items := giftbag.Items
						if !g.IsEmpty(items) {
							// 获取 itemIds 并转换为 []uint64
							itemIdsInterface := array.Column(items, "ItemId")
							var itemIds []uint64
							for _, id := range itemIdsInterface {
								if uintID, ok := id.(uint64); ok {
									itemIds = append(itemIds, uintID)
								} else {
									return nil, fmt.Errorf("invalid itemId type: %T", id)
								}
							}
							productItems, err := service.ProductBase().GetItems(ctx, itemIds, 0)
							if err != nil {
								return nil, err
							}
							activityRes.Item = productItems

							remain := giftbag.GiftbagQuantity - activityBase.ActivityEffectiveQuantity
							if remain < 0 {
								remain = 0
							}
							activityRes.RemainQuantity = int(remain)
						}
					}
				}
			}
		default:
			if len(activityRule) > 0 {
				gjson.DecodeTo(activityRule, &activityRuleVo)
			}
			itemIds, _ := s.GetActivityAllItemIds(ctx, activityBase)
			if !g.IsEmpty(itemIds) {
				items, err := service.ProductBase().GetItems(ctx, itemIds, 0)
				if err != nil {
					return nil, err
				}
				activityRes.Item = items
			}
		}

		activityRes.ActivityRuleJson = activityRuleVo
		activityBaseRes = append(activityBaseRes, activityRes)
	}

	activityBaseResPage.Items = activityBaseRes

	return activityBaseResPage, nil
}
func (s *sActivityBase) FixActivityData(ctx context.Context, activityBaseList []*entity.ActivityBase) ([]*entity.ActivityBase, error) {
	if len(activityBaseList) == 0 {
		return []*entity.ActivityBase{}, nil
	}

	now := time.Now()
	currentTime := now.UnixNano() / int64(time.Millisecond)

	var baserMap *do.ActivityBase
	for _, i := range activityBaseList {
		gconv.Scan(i, &baserMap)
		activityId := i.ActivityId
		// 判断活动店铺有效
		activityState := i.ActivityState
		// 活动开始时间
		activityStartTime := i.ActivityStarttime
		// 活动结束时间
		activityEndTime := i.ActivityEndtime
		// 活动类型
		activityType := i.ActivityType

		if !g.IsEmpty(activityType) && activityType == consts.GET_VOUCHER_BY_PURCHASE {
			// 根据售卖优惠券id获取对应商品信息
		}

		switch activityState {
		case consts.ACTIVITY_STATE_WAITING:
			if activityStartTime <= uint64(currentTime) {
				if activityEndTime <= uint64(currentTime) {
					baserMap.ActivityState = consts.ACTIVITY_STATE_FINISHED
				} else {
					i.ActivityState = consts.ACTIVITY_STATE_NORMAL
				}

				// 更新数据
				dao.ActivityBase.Edit(ctx, activityId, baserMap)
			}
		case consts.ACTIVITY_STATE_NORMAL:
			if activityEndTime < uint64(currentTime) {
				i.ActivityState = consts.ACTIVITY_STATE_FINISHED
				// 更新数据
				dao.ActivityBase.Edit(ctx, activityId, baserMap)
			}

			if activityStartTime > uint64(currentTime) {
				i.ActivityState = consts.ACTIVITY_STATE_WAITING

				// 更新数据
				dao.ActivityBase.Edit(ctx, activityId, baserMap)
			}
		case consts.ACTIVITY_STATE_FINISHED:
			if activityStartTime <= uint64(currentTime) {
				if activityEndTime <= uint64(currentTime) {
					i.ActivityState = consts.ACTIVITY_STATE_FINISHED
				} else {
					i.ActivityState = consts.ACTIVITY_STATE_NORMAL
				}

				// 更新数据
				dao.ActivityBase.Edit(ctx, activityId, baserMap)
			}
			if activityStartTime > uint64(currentTime) {
				i.ActivityState = consts.ACTIVITY_STATE_WAITING

				// 更新数据
				dao.ActivityBase.Edit(ctx, activityId, baserMap)
			}
		case consts.ACTIVITY_STATE_CLOSED:
			if activityStartTime <= uint64(currentTime) && activityEndTime <= uint64(currentTime) {
				i.ActivityState = consts.ACTIVITY_STATE_FINISHED
				// 更新数据
				dao.ActivityBase.Edit(ctx, activityId, baserMap)
			}
		}
	}

	return activityBaseList, nil
}

// GetActivityAllItemIds 获取所有活动商品的ID列表
func (s *sActivityBase) GetActivityAllItemIds(ctx context.Context, activityBase *entity.ActivityBase) ([]uint64, error) {
	var itemIds []uint64

	// 如果 activityBase 为空，返回空列表
	if activityBase == nil {
		return nil, nil
	}

	activityTypeId := activityBase.ActivityTypeId
	var activityRuleVo *model.ActivityRuleVo

	// 如果 activityBase 的活动规则不为空，将其解析为 ActivityRuleVo
	if !g.IsEmpty(activityBase.ActivityRule) {
		err := gconv.Struct(activityBase.ActivityRule, &activityRuleVo)
		if err != nil {
			return nil, err
		}
	}

	// 根据活动类型ID进行不同的处理
	switch activityTypeId {
	case consts.ACTIVITY_TYPE_BARGAIN,
		consts.ACTIVITY_TYPE_POINT_SHOPPING,
		consts.ACTIVITY_TYPE_GIFT,
		consts.ACTIVITY_TYPE_REDUCTION,
		consts.ACTIVITY_TYPE_PF_GROUPBUY_STORE:
		// 如果活动类型是特定类型，从活动规则中获取商品ID列表
		if activityRuleVo != nil && &activityRuleVo.Requirement != nil && activityRuleVo.Requirement.Buy != nil {
			itemIds = activityRuleVo.Requirement.Buy.Item
		}
	case consts.ACTIVITY_TYPE_LIMITED_DISCOUNT,
		consts.ACTIVITY_TYPE_GROUPBOOKING:
		// 如果活动类型是限时折扣或拼团，从活动商品表中获取商品ID列表
		activityId := activityBase.ActivityId
		activityItems, err := dao.ActivityItem.Find(ctx, &do.ActivityItemListInput{Where: do.ActivityItem{ActivityId: activityId}})
		if err != nil {
			return nil, err
		}

		if !g.IsEmpty(activityItems) {
			for _, item := range activityItems {
				itemIds = append(itemIds, item.ItemId)
			}
		}
	}

	return itemIds, nil
}

// EditActivityBase 编辑活动基础信息
func (s *sActivityBase) EditActivityBase(ctx context.Context, activityId uint, data *do.ActivityBase) (bool, error) {
	/*	activityTypeId := data.ActivityTypeId

		if g.Equal(activityTypeId, consts.ACTIVITY_TYPE_GROUPBOOKING) || gutil.Equal(activityTypeId, consts.ACTIVITY_TYPE_CUTPRICE) {
			// 判断商品是否更改
			activityItems, err := dao.ActivityItem.Find(ctx, &do.ActivityItemListInput{Where: do.ActivityItem{ActivityId: activityId}})
			if err != nil {
				return false, err
			}

			itemIdRow := array.Column(activityItems, "ItemId"

			activityRule := gjson.New(data.ActivityRule)
			itemId := gconv.Uint64(activityRule.Get("item_id"))

			for _, itemIdInRow := range itemIdRow {
				if garray.NotEqual(itemId, itemIdInRow) {
					if !s.Remove(ctx, activityId, itemIdInRow, data) {
						return false, gerror.New("修改活动商品信息失败！")
					}
				}
			}

			if !gutil.InArray(itemIdRow, itemId) {
				if !s.Add(ctx, activityId, []uint64{itemId}, data) {
					return false, gerror.New("修改活动商品信息失败！")
				}
			}
		}
	*/
	// 保存活动基础信息
	if _, err := dao.ActivityBase.Save(ctx, data); err != nil {
		return false, err
	}

	activityState := data.ActivityState

	if activityState != nil {
		// 获取活动项
		activityItems, err := dao.ActivityItem.Find(ctx, &do.ActivityItemListInput{Where: do.ActivityItem{ActivityId: activityId}})
		if err != nil {
			return false, err
		}

		// 修改store_activity_item
		if len(activityItems) > 0 {
			activityItemList := &do.ActivityItem{}
			activityItemIds := array.Column(activityItems, "ActivityItemId")
			activityItemList.ActivityItemId = activityItemIds
			activityItem := &do.ActivityItem{
				ActivityItemState: activityState,
			}
			if _, err := dao.ActivityItem.Edit(ctx, activityItem, activityItemList); err != nil {
				return false, err
			}
		}

		// 判断是否有状态的变化，如果有，则修改product_index
		if len(activityItems) > 0 {
			activityRow, err := s.Get(ctx, activityId)
			if err != nil {
				return false, err
			}
			activityTypeId := activityRow.ActivityTypeId

			productIdRow := array.Column(activityItems, "ProductId")
			productIndexRows, err := dao.ProductIndex.Gets(ctx, productIdRow)
			if err != nil {
				return false, err
			}

			for _, productIndexRow := range productIndexRows {
				uniqueActivityTypeIds := garray.NewArray()
				uniqueActivityTypeIds.Append(productIndexRow.ActivityTypeIds)
				if activityState == consts.ACTIVITY_STATE_NORMAL {
					// 不存在则添加
					if !uniqueActivityTypeIds.Contains(activityTypeId) {
						uniqueActivityTypeIds.Append(activityTypeId)
						activityTypeIds := uniqueActivityTypeIds.Slice()
						productIndexRow.ActivityTypeIds = gstr.Join(gconv.Strings(activityTypeIds), ",")
					}
				} else {
					// 存在则删除
					if uniqueActivityTypeIds.Contains(activityTypeId) {
						uniqueActivityTypeIds.Remove(int(activityTypeId))
						activityTypeIds := uniqueActivityTypeIds.Slice()
						productIndexRow.ActivityTypeIds = gstr.Join(gconv.Strings(activityTypeIds), ",")
					}
				}

				productIndex := &do.ProductIndex{}
				gconv.Scan(productIndexRows, productIndex)

				if _, err := dao.ProductIndex.Save(ctx, productIndex); err != nil {
					return false, err
				}
			}
		}
	}

	return true, nil
}
