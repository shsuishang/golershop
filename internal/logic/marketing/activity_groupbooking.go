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
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/marketing"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
)

type sActivityGroupbooking struct{}

func init() {
	service.RegisterActivityGroupbooking(NewActivityGroupbooking())
}

func NewActivityGroupbooking() *sActivityGroupbooking {
	return &sActivityGroupbooking{}
}

// Get 根据编号读取活动信息
func (s *sActivityGroupbooking) Get(ctx context.Context, id any) (out *entity.ActivityGroupbooking, err error) {
	var list []*entity.ActivityGroupbooking
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
func (s *sActivityGroupbooking) Gets(ctx context.Context, id any) (list []*entity.ActivityGroupbooking, err error) {
	err = dao.ActivityGroupbooking.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询活动数据
func (s *sActivityGroupbooking) Find(ctx context.Context, in *do.ActivityGroupbookingListInput) (out []*entity.ActivityGroupbooking, err error) {
	out, err = dao.ActivityGroupbooking.Find(ctx, in)

	return out, err
}

// List 分页读取活动
func (s *sActivityGroupbooking) List(ctx context.Context, in *do.ActivityGroupbookingListInput) (out *do.ActivityGroupbookingListOutput, err error) {
	out, err = dao.ActivityGroupbooking.List(ctx, in)

	return out, err
}

// Add 新增活动
func (s *sActivityGroupbooking) Add(ctx context.Context, in *do.ActivityGroupbooking) (lastInsertId int64, err error) {
	lastInsertId, err = dao.ActivityGroupbooking.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑活动
func (s *sActivityGroupbooking) Edit(ctx context.Context, in *do.ActivityGroupbooking) (affected int64, err error) {
	_, err = dao.ActivityGroupbooking.Edit(ctx, in.ActivityId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除活动记录
func (s *sActivityGroupbooking) Remove(ctx context.Context, id any) (affected int64, err error) {
	affected, err = dao.ActivityGroupbooking.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// DoGroupbooking 参加拼团活动
func (s *sActivityGroupbooking) DoGroupbooking(ctx context.Context, orderId string, userId uint, gbId uint, activityInfo *model.ActivityInfoVo) (id uint, err error) {
	//todo

	return 0, err
}

// DoGroupbooking 参加拼团活动
func (s *sActivityGroupbooking) SetPaidYes(ctx context.Context, orderId string, userId uint) (flag bool, err error) {
	//todo

	return
}

// CheckGroupbookingSuccess
func (s *sActivityGroupbooking) CheckGroupbookingSuccess(ctx context.Context, orderId string) (flag bool, err error) {
	//todo

	return
}

// ListsUserGroupbookingHistory 获取用户拼团历史记录列表
func (s *sActivityGroupbooking) ListsUserGroupbookingHistory(ctx context.Context, in *do.ActivityGroupbookingHistoryListInput) (res *marketing.ActivityGroupbookingHistoryRes, err error) {
	groupbookingHistoryRes := &marketing.ActivityGroupbookingHistoryRes{}
	historyPage, err := dao.ActivityGroupbookingHistory.List(ctx, in)
	if err != nil {
		return nil, err
	}

	if historyPage != nil && !g.IsEmpty(historyPage.Items) {
		gconv.Scan(historyPage, groupbookingHistoryRes)

		orderIds := array.Column(groupbookingHistoryRes.Items, "OrderId")
		gbIds := array.Column(groupbookingHistoryRes.Items, "GbId")

		// 获取拼团信息
		groupbookings, err := dao.ActivityGroupbooking.Gets(ctx, gbIds)
		if err != nil {
			return nil, err
		}

		gbAmountQuantityMap := make(map[uint]uint)
		if !g.IsEmpty(groupbookings) {
			for _, groupbooking := range groupbookings {
				gbAmountQuantityMap[groupbooking.GbId] = groupbooking.GbAmountQuantity
			}
		}

		// 获取应付金额
		orderBaseList, err := dao.OrderBase.Gets(ctx, orderIds)
		if err != nil {
			return nil, err
		}

		amountMap := make(map[string]float64)
		if !g.IsEmpty(orderBaseList) {
			for _, orderBase := range orderBaseList {
				amountMap[orderBase.OrderId] = orderBase.OrderPaymentAmount
			}
		}

		// 获取订单商品信息
		orderItemQueryWrapper := do.OrderItemListInput{
			Where: do.OrderItem{
				OrderId: orderIds,
				UserId:  in.Where.UserId,
			},
		}
		orderItemList, err := dao.OrderItem.Find(ctx, &orderItemQueryWrapper)
		if err != nil {
			return nil, err
		}

		orderItemMap := make(map[string]*entity.OrderItem)
		if !g.IsEmpty(orderItemList) {
			for _, orderItem := range orderItemList {
				orderItemMap[orderItem.OrderId] = orderItem
			}
		}

		// 构建返回结果
		groupbookingHistoryResList := make([]map[string]interface{}, 0)
		for _, groupbookingHistory := range historyPage.Items {

			activityGroupbookingHistoryRes := make(map[string]interface{})
			gconv.Scan(groupbookingHistory, &activityGroupbookingHistoryRes)

			if val, ok := gbAmountQuantityMap[groupbookingHistory.GbId]; ok {
				activityGroupbookingHistoryRes["gb_amount_quantity"] = val
			}

			if val, ok := amountMap[groupbookingHistory.OrderId]; ok {
				activityGroupbookingHistoryRes["order_payment_amount"] = val
			}

			if val, ok := orderItemMap[groupbookingHistory.OrderId]; ok {
				activityGroupbookingHistoryRes["item_id"] = val.ItemId
				activityGroupbookingHistoryRes["item_name"] = val.ItemName
				activityGroupbookingHistoryRes["item_unit_price"] = val.ItemUnitPrice
				activityGroupbookingHistoryRes["order_id"] = val.OrderId
				activityGroupbookingHistoryRes["order_item_quantity"] = val.OrderItemQuantity
				activityGroupbookingHistoryRes["order_item_image"] = val.OrderItemImage
				activityGroupbookingHistoryRes["order_item_sale_price"] = val.OrderItemSalePrice
				activityGroupbookingHistoryRes["order_item_payment_amount"] = val.OrderItemPaymentAmount
				activityGroupbookingHistoryRes["product_name"] = val.ProductName
			}

			groupbookingHistoryResList = append(groupbookingHistoryResList, activityGroupbookingHistoryRes)
		}

		groupbookingHistoryRes.Items = groupbookingHistoryResList
	}

	return groupbookingHistoryRes, err

}

// GetUserGroupbooking 获取用户拼团信息
func (s *sActivityGroupbooking) GetUserGroupbooking(ctx context.Context, activityGroupbookingReq *marketing.ActivityGroupbookingReq) (activityGroupbookingRes *marketing.ActivityGroupbookingRes, err error) {
	activityGroupbookingRes = &marketing.ActivityGroupbookingRes{}

	// 获取拼团信息
	activityGroupbooking, err := dao.ActivityGroupbooking.Get(ctx, activityGroupbookingReq.GbId)
	if err != nil {
		return nil, err
	}
	if activityGroupbooking == nil {
		return nil, gerror.New("拼团信息为空！")
	}
	activityGroupbookingRes.ActivityGroupbooking = activityGroupbooking

	// 获取订单商品信息
	orderItem, err := dao.OrderItem.FindOne(ctx, &do.OrderItemListInput{
		Where: do.OrderItem{
			OrderId: activityGroupbookingReq.OrderId,
		},
	})
	if err != nil {
		return nil, err
	}
	if orderItem == nil {
		return nil, gerror.New("订单商品信息为空！")
	}

	// 复制订单商品信息到拼团结果
	gconv.Scan(orderItem, activityGroupbookingRes)

	// 复制拼团信息到拼团结果
	gconv.Scan(activityGroupbooking, activityGroupbookingRes)

	// 获取活动信息
	activityBase, err := dao.ActivityBase.Get(ctx, activityGroupbooking.ActivityId)
	if err != nil {
		return nil, err
	}
	if activityBase == nil {
		return nil, gerror.New("该活动不存在！")
	}

	// 设置活动备注和结束时间
	activityGroupbookingRes.ActivityRemark = activityBase.ActivityRemark
	activityGroupbookingRes.ActivityEndtime = activityBase.ActivityEndtime

	// 获取用户信息
	activityGroupbookingHistories, err := dao.ActivityGroupbookingHistory.Find(ctx, &do.ActivityGroupbookingHistoryListInput{
		Where: do.ActivityGroupbookingHistory{
			GbId:    activityGroupbookingReq.GbId,
			GbhFlag: 1,
		},
	})
	if err != nil {
		return nil, err
	}
	if len(activityGroupbookingHistories) == 0 {
		return nil, gerror.New("拼团记录为空！")
	}

	// 提取用户ID列表
	userIds := array.Column(activityGroupbookingHistories, "UserId")

	// 获取用户信息
	gbUsers, err := dao.UserInfo.Gets(ctx, userIds)
	if err != nil {
		return nil, err
	}
	activityGroupbookingRes.GbUsers = gbUsers

	return activityGroupbookingRes, nil
}
