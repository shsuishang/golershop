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

package trade

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"github.com/shopspring/decimal"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/dao/global"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
	"golershop.cn/utility/phone"
	"time"
)

type sOrderReturn struct{}

func init() {
	service.RegisterOrderReturn(NewOrderReturn())
}

func NewOrderReturn() *sOrderReturn {
	return &sOrderReturn{}
}

// Get 读取订单
func (s *sOrderReturn) Get(ctx context.Context, id any) (out *entity.OrderReturn, err error) {
	var list []*entity.OrderReturn
	list, err = s.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// Gets 读取多条订单
func (s *sOrderReturn) Gets(ctx context.Context, id any) (list []*entity.OrderReturn, err error) {
	err = dao.OrderInfo.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sOrderReturn) Find(ctx context.Context, in *do.OrderReturnListInput) (out []*entity.OrderReturn, err error) {
	out, err = dao.OrderReturn.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sOrderReturn) List(ctx context.Context, in *do.OrderReturnListInput) (out *do.OrderReturnListOutput, err error) {
	out, err = dao.OrderReturn.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sOrderReturn) Add(ctx context.Context, in *do.OrderReturn) (lastInsertId int64, err error) {
	lastInsertId, err = dao.OrderReturn.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sOrderReturn) Edit(ctx context.Context, in *do.OrderReturn) (affected int64, err error) {
	_, err = dao.OrderReturn.Edit(ctx, in.ReturnId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sOrderReturn) Remove(ctx context.Context, id any) (affected int64, err error) {

	affected, err = dao.OrderReturn.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// GetByReturnId 根据退货 ID 获取订单退货信息
func (s *sOrderReturn) GetByReturnId(ctx context.Context, returnId string) (*model.OrderReturnVo, error) {
	orderReturnVo := &model.OrderReturnVo{}

	// 商品信息
	orderReturn, err := dao.OrderReturn.Get(ctx, returnId)
	if err != nil {
		return nil, err
	}
	if orderReturn == nil {
		return orderReturnVo, nil
	}
	gconv.Scan(orderReturn, orderReturnVo)

	// 客户名称
	userInfos, err := dao.UserInfo.Find(ctx, &do.UserInfoListInput{
		Where: do.UserInfo{
			UserId: orderReturn.BuyerUserId,
		},
	})
	if err != nil {
		return nil, err
	}
	if len(userInfos) > 0 {
		orderReturnVo.BuyerUserName = userInfos[0].UserNickname
	}

	orderReturnItemVos := []*model.OrderReturnItemVo{}

	// 查询订单退货详情表
	orderReturnItems, err := dao.OrderReturnItem.Find(ctx, &do.OrderReturnItemListInput{
		Where: do.OrderReturnItem{
			ReturnId: orderReturn.ReturnId,
		},
	})
	if err != nil {
		return nil, err
	}
	if len(orderReturnItems) > 0 {
		for _, orderReturnItem := range orderReturnItems {
			orderReturnItemVo := &model.OrderReturnItemVo{}
			gconv.Scan(orderReturnItem, orderReturnItemVo)
			if orderReturnItem.ReturnItemImage != "" {
				orderReturnItemVo.ReturnItemImageList = gconv.Strings(gstr.Split(orderReturnItem.ReturnItemImage, ","))
			}
			orderReturnItemVos = append(orderReturnItemVos, orderReturnItemVo)
		}
	}

	// 订单商品名称、单价
	if len(orderReturnItemVos) > 0 {
		orderItemIds := gset.New()
		for _, itemVo := range orderReturnItemVos {
			orderItemIds.Add(itemVo.OrderItemId)
		}

		orderItemList, err := dao.OrderItem.Gets(ctx, gconv.Int64s(orderItemIds.Slice()))
		if err != nil {
			return nil, err
		}
		if len(orderItemList) > 0 {
			orderItemMap := make(map[uint64]*entity.OrderItem)
			for _, item := range orderItemList {
				orderItemMap[item.OrderItemId] = item
			}
			for _, itemVo := range orderReturnItemVos {
				orderItem, ok := orderItemMap[itemVo.OrderItemId]
				if ok {
					itemVo.ProductItemName = fmt.Sprintf("%s %s", orderItem.ProductName, orderItem.ItemName)
					itemVo.ItemUnitPrice = orderItem.ItemUnitPrice
				}
			}
		}
	}
	orderReturnVo.Items = orderReturnItemVos

	// 收货信息
	orderDeliveryAddresses, err := dao.OrderDeliveryAddress.Find(ctx, &do.OrderDeliveryAddressListInput{
		Where: do.OrderDeliveryAddress{
			OrderId: orderReturn.OrderId,
		},
	})
	if err != nil {
		return nil, err
	}
	if len(orderDeliveryAddresses) > 0 {
		gconv.Scan(orderDeliveryAddresses[0], orderReturnVo)
	}

	return orderReturnVo, nil
}

// Refused 拒绝退货退款
func (s *sOrderReturn) Refused(ctx context.Context, orderReturn *do.OrderReturn) (bool, error) {
	// 获取storeId
	storeId := orderReturn.StoreId
	// 根据退货ID获取退货信息
	orderReturnData, err := dao.OrderReturn.Get(ctx, orderReturn.ReturnId)
	if err != nil {
		return false, err
	}

	if true || storeId == orderReturnData.StoreId {
		orderReturn.ReturnStateId = consts.RETURN_PROCESS_REFUSED

		if _, err := dao.OrderReturn.Edit(ctx, orderReturn.ReturnId, orderReturn); err != nil {
			return false, fmt.Errorf("拒绝退款/退货失败")
		}

		// 通知买家退货退款成功
		//messageId := "refunds-and-reminders"
		//args := map[string]interface{}{
		//	"order_id":             orderReturn.OrderId,
		//	"return_refund_amount": orderReturn.ReturnRefundAmount,
		//}
		//if err := service.Message().SendNoticeMsg(ctx, orderReturn.BuyerUserId, messageId, args); err != nil {
		//	return false, err
		//}
	} else {
		return false, fmt.Errorf("无权限")
	}

	return true, nil
}

// Review 审核退货退款
func (s *sOrderReturn) Review(ctx context.Context, orderReturn *do.OrderReturn, receivingAddress uint) (bool, error) {
	// 获取storeId
	storeId := orderReturn.StoreId.(uint)

	// 获取退货ID列表
	returnIds := gstr.Split(orderReturn.ReturnId.(string), ",")
	orderReturns, err := dao.OrderReturn.Gets(ctx, returnIds)
	if err != nil {
		return false, err
	}

	if true || storeId == orderReturns[0].StoreId {
		// 判断退货类型
		if orderReturn.ReturnFlag == consts.ORDER_NOT_NEED_RETURN_GOODS {
			// 修改退款退货表信息
			if _, err := dao.OrderReturn.Edit(ctx, orderReturn.ReturnId, orderReturn); err != nil {
				return false, fmt.Errorf("修改退款退货表信息失败")
			}
			// 不用退货
			if err := s.dealWithReturn(ctx, returnIds, storeId, consts.RETURN_PROCESS_CHECK, orderReturns, consts.RETURN_PROCESS_FINISH); err != nil {
				return false, err
			}
		} else {
			// 获取收货地址
			storeShippingAddress, err := dao.StoreShippingAddress.Get(ctx, receivingAddress)
			if err != nil {
				return false, fmt.Errorf("请选择收货人")
			}

			if storeShippingAddress != nil {
				returnAddr := fmt.Sprintf("%s%s%s%s", storeShippingAddress.SsProvince, storeShippingAddress.SsCity, storeShippingAddress.SsCounty, storeShippingAddress.SsAddress)
				returnMobile := storeShippingAddress.SsMobile
				returnContactName := storeShippingAddress.SsName
				orderReturn.ReturnAddr = returnAddr
				orderReturn.ReturnTel = returnMobile
				orderReturn.ReturnContactName = returnContactName

				// 修改退款退货表信息
				if _, err := dao.OrderReturn.Edit(ctx, orderReturn.ReturnId, orderReturn); err != nil {
					return false, fmt.Errorf("修改退款退货表信息失败")
				}

				// 需要退货
				if err := s.dealWithReturn(ctx, returnIds, storeId, consts.RETURN_PROCESS_CHECK, orderReturns, 0); err != nil {
					return false, err
				}
			} else {
				return false, fmt.Errorf("请选择收货人")
			}
		}
	} else {
		return false, fmt.Errorf("无权限")
	}

	// 通知买家退货退款成功
	//messageId := "refunds-and-reminders"
	//args := map[string]interface{}{
	//	"order_id":             orderReturn.ReturnId,
	//	"return_refund_amount": orderReturn.ReturnRefundAmount,
	//}
	//if err := service.Message().SendNoticeMsg(ctx, orderReturn.BuyerUserId, messageId, args); err != nil {
	//	return false, err
	//}

	return true, nil
}

func (s *sOrderReturn) getNextReturnProcess(ctx context.Context, stateId uint) (uint, error) {
	return service.ConfigBase().GetNextReturnStateId(ctx, stateId)
}

func (s *sOrderReturn) checkNeedRefund(ctx context.Context, stateId, nextStateId uint) bool {
	processRefund := global.ReturnProcessMap[consts.RETURN_PROCESS_REFUND]
	processReceiptConfirmation := global.ReturnProcessMap[consts.RETURN_PROCESS_RECEIPT_CONFIRMATION]
	processReturnStateId := global.ReturnProcessMap[stateId]
	processReturnNextStateId := global.ReturnProcessMap[nextStateId]

	return stateId != nextStateId &&
		processReturnStateId <= processRefund &&
		processReturnNextStateId >= processReceiptConfirmation
}

// DealWithReturn 处理退货
func (s *sOrderReturn) dealWithReturn(ctx context.Context, returnIds []string, storeId uint, stateId uint, orderReturns []*entity.OrderReturn, nextStateId uint) error {
	// 检查退货ID和订单退货是否为空
	if len(returnIds) == 0 || len(orderReturns) == 0 {
		return gerror.New("请选择需要审核的订单")
	}

	// 获取订单下一个状态
	if nextStateId == 0 {
		nextStateId, _ = s.getNextReturnProcess(ctx, stateId)
		if nextStateId == 0 {
			return gerror.New("读取退单状态有误")
		}
	}

	// 设置订单退货状态和时间
	orderReturn := &do.OrderReturn{
		ReturnStateId:   nextStateId,
		ReturnStoreTime: gtime.Now(),
	}

	// 商家收货确认，增加库存
	if stateId == consts.RETURN_PROCESS_RECEIVED {
		// 更改物流状态为已签收
		orderReturn.ReturnItemStateId = consts.ORDER_STATE_FINISH
	}

	// 执行真正退款逻辑
	if s.checkNeedRefund(ctx, stateId, nextStateId) {
		// 卖家账户扣款，买家账户增加
		orderReturn.ReturnIsPaid = true
		orderReturn.ReturnFinishTime = gtime.Now()
		// 执行退款操作
		//if !service.ConsumeReturn().DoRefund(ctx, orderReturns) {
		//	return gerror.New("退款失败")
		//}
	}

	// 修改退单状态
	if err := s.editReturnNextState(ctx, returnIds, stateId, orderReturn); err != nil {
		return err
	}

	// 当前状态为审核状态
	if stateId == consts.RETURN_PROCESS_CHECK {
		for _, item := range orderReturns {
			orderId := item.OrderId
			returnIsShippingFee := item.ReturnIsShippingFee

			if !returnIsShippingFee {
				// 查询订单商品列表
				orderItemList, err := dao.OrderItem.Find(ctx, &do.OrderItemListInput{
					Where: do.OrderItem{OrderId: orderId},
				})
				if err != nil {
					return err
				}
				if len(orderItemList) == 0 {
					return gerror.New("订单商品列表为空")
				}

				// 计算订单商品可退款金额
				var orderItemCanRefundAmount float64
				for _, item := range orderItemList {
					orderItemCanRefundAmount += item.OrderItemCanRefundAmount
				}

				// 查询订单退货详情列表
				orderReturnItemList, err := dao.OrderReturnItem.Find(ctx, &do.OrderReturnItemListInput{
					Where: do.OrderReturnItem{
						OrderId:       orderId,
						ReturnStateId: consts.RETURN_PROCESS_RECEIVED,
					},
				})
				if err != nil {
					return err
				}
				if len(orderReturnItemList) == 0 {
					return gerror.New("订单退货详情列表为空")
				}

				// 计算退货详情总金额
				var returnItemSubtotal float64
				for _, item := range orderReturnItemList {
					returnItemSubtotal += item.ReturnItemSubtotal
				}

				// 全部退款完成
				orderData := &do.OrderData{}
				if orderItemCanRefundAmount == returnItemSubtotal {
					orderData.OrderRefundStatus = 2
				} else {
					// 部分退款
					orderData.OrderRefundStatus = 1
				}
				if _, err := dao.OrderData.Edit(ctx, orderId, orderData); err != nil {
					return gerror.New("修改订单详细信息失败")
				}

				// 修改订单同步状态，重新同步
				orderInfo := &do.OrderInfo{
					OrderIsSync: false,
				}
				if _, err := dao.OrderInfo.Edit(ctx, orderId, orderInfo); err != nil {
					return gerror.New("修改订单同步状态失败")
				}
			}
		}
	}
	return nil
}

func (s *sOrderReturn) editReturnNextState(ctx context.Context, returnIds []string, stateId uint, orderReturn *do.OrderReturn) error {
	// 修改退款退货表 退单状态
	returnQueryWrapper := &do.OrderReturnListInput{
		Where: do.OrderReturn{
			ReturnId:      returnIds,
			ReturnStateId: stateId,
		},
	}

	if _, err := dao.OrderReturn.EditWhere(ctx, returnQueryWrapper, orderReturn); err != nil {
		return gerror.New("修改退单状态失败")
	}

	// 修改订单退货详情表 退单状态
	orderReturnItem := &do.OrderReturnItem{
		ReturnStateId: orderReturn.ReturnStateId,
	}
	returnItemQueryWrapper := &do.OrderReturnItemListInput{
		Where: do.OrderReturnItem{
			ReturnId: returnIds,
		},
	}

	if _, err := dao.OrderReturnItem.Edit(ctx, returnItemQueryWrapper, orderReturnItem); err != nil {
		return gerror.New("修改订单退货详情状态失败")
	}

	return nil
}

// GetList 获取订单退货列表
func (s *sOrderReturn) GetList(ctx context.Context, input *do.OrderReturnListInput) (res *model.OrderReturnOutput, err error) {
	res = &model.OrderReturnOutput{}
	input.Sidx = "return_add_time"
	input.Sort = "DESC"

	orderReturnIPage, err := s.List(ctx, input)
	if err != nil {
		return nil, err
	}

	if orderReturnIPage == nil || len(orderReturnIPage.Items) == 0 {
		return res, nil
	}

	// 复制分页信息
	res.Records = orderReturnIPage.Records
	res.Size = orderReturnIPage.Size
	res.Total = orderReturnIPage.Total
	res.Page = orderReturnIPage.Page
	gconv.Struct(orderReturnIPage.Items, &res.Items)

	// 退款退货表信息
	returnList := orderReturnIPage.Items
	// 通过 returnId 获得订单退货详情表信息
	returnIds := g.SliceStr{}
	for _, item := range returnList {
		returnIds = append(returnIds, item.ReturnId)
	}
	returnIds = gstr.Split(gstr.Join(returnIds, ","), ",")

	// 查询订单退货详情表信息
	orderReturnItems, err := dao.OrderReturnItem.Find(ctx, &do.OrderReturnItemListInput{
		Where: do.OrderReturnItem{ReturnId: returnIds},
	})
	if err != nil {
		return nil, err
	}

	// 订单退货详情 Map
	orderReturnItemMap := make(map[string][]*entity.OrderReturnItem)
	// 通过 orderItemId 获得订单商品表
	orderItemMap := make(map[string][]*entity.OrderItem)
	if len(orderReturnItems) > 0 {
		for _, item := range orderReturnItems {
			orderReturnItemMap[item.ReturnId] = append(orderReturnItemMap[item.ReturnId], item)
		}

		orderItemIds := make([]uint64, 0)
		for _, item := range orderReturnItems {
			orderItemIds = append(orderItemIds, item.OrderItemId)
		}
		orderItems, err := dao.OrderItem.Gets(ctx, orderItemIds)
		if err != nil {
			return nil, err
		}

		if len(orderItems) > 0 {
			for _, item := range orderItems {
				orderItemMap[item.OrderId] = append(orderItemMap[item.OrderId], item)
			}
		}
	}

	// 封装数据
	orderReturnResList := []*model.OrderReturnVo{}
	for _, orderReturn := range returnList {
		orderReturnRes := &model.OrderReturnVo{}
		gconv.Struct(orderReturn, orderReturnRes)

		// 退货商品总数量
		if len(orderReturnItemMap) > 0 {
			returnItems := orderReturnItemMap[orderReturn.ReturnId]
			if len(returnItems) > 0 {
				for _, item := range returnItems {
					orderReturnRes.ReturnNum += item.ReturnItemNum
				}
			}
		}

		orderReturnItemVos := []*model.OrderReturnItemVo{}
		if len(orderItemMap) > 0 {
			orderItems := orderItemMap[orderReturn.OrderId]
			if len(orderItems) > 0 {
				for _, orderItem := range orderItems {
					orderReturnItemVo := &model.OrderReturnItemVo{}
					gconv.Struct(orderItem, orderReturnItemVo)
					orderReturnItemVos = append(orderReturnItemVos, orderReturnItemVo)
				}
			}
		}
		orderReturnRes.Items = orderReturnItemVos
		orderReturnResList = append(orderReturnResList, orderReturnRes)
	}

	if len(orderReturnResList) > 0 {
		res.Items = orderReturnResList
	}

	return res, nil
}

func (s *sOrderReturn) GetReturn(ctx context.Context, returnId any) (*model.OrderReturnVo, error) {
	orderReturnRes := &model.OrderReturnVo{}
	orderReturn, err := dao.OrderReturn.Get(ctx, returnId)
	if err != nil {
		return nil, err
	}

	if orderReturn == nil {
		return nil, nil
	}

	gconv.Scan(orderReturn, orderReturnRes)

	//订单退货详情表
	orderReturnItemQueryWrapper := &do.OrderReturnItemListInput{Where: do.OrderReturnItem{ReturnId: returnId}}
	orderReturnItems, err := dao.OrderReturnItem.Find(ctx, orderReturnItemQueryWrapper)
	if err != nil {
		return nil, err
	}

	if len(orderReturnItems) > 0 {
		//订单退货详情 退款金额
		var reduce float64
		for _, item := range orderReturnItems {
			reduce += item.ReturnItemSubtotal
		}
		orderReturnRes.SubmitReturnRefundAmount = reduce
		//订单商品
		var orderItemIds []any
		orderItemMap := make(map[uint64]*entity.OrderItem)
		for _, item := range orderReturnItems {
			orderItemIds = append(orderItemIds, item.OrderItemId)
		}

		orderItems, err := dao.OrderItem.Gets(ctx, orderItemIds)
		if err != nil {
			return nil, err
		}

		for _, item := range orderItems {
			orderItemMap[item.OrderItemId] = item
		}

		//商品信息
		var orderReturnItemVos []*model.OrderReturnItemVo

		var returnReasonIds []any
		for _, item := range orderReturnItems {
			returnReasonIds = append(returnReasonIds, item.ReturnReasonId)
		}

		returnReasonList, err := dao.OrderReturnReason.Gets(ctx, returnReasonIds)
		if err != nil {
			return nil, err
		}

		//退款原因
		if !g.IsEmpty(orderReturn.ReturnReasonId) {
			for _, reason := range returnReasonList {
				if reason.ReturnReasonId == orderReturn.ReturnReasonId {
					orderReturnRes.ReturnReasonName = reason.ReturnReasonName
					break
				}
			}
		}

		for _, orderReturnItem := range orderReturnItems {
			orderReturnItemVo := &model.OrderReturnItemVo{}
			gconv.Scan(orderReturnItem, orderReturnItemVo)
			//订单商品
			orderItem, ok := orderItemMap[orderReturnItem.OrderItemId]
			if ok {
				gconv.Scan(orderItem, orderReturnItemVo)
			}
			//退款凭证
			if len(orderReturnItem.ReturnItemImage) > 0 {
				orderReturnItemVo.ReturnItemImageList = gconv.SliceStr(orderReturnItem.ReturnItemImage)
			}
			orderReturnItemVos = append(orderReturnItemVos, orderReturnItemVo)
		}
		orderReturnRes.Items = orderReturnItemVos
	}

	return orderReturnRes, nil
}

// EditReturn 编辑退货
func (s *sOrderReturn) EditReturn(ctx context.Context, orderReturn *do.OrderReturn) (bool, error) {
	// 获取快递公司ID
	expressId := orderReturn.ExpressId
	if expressId != 0 {
		// 获取快递公司信息
		expressBase, err := dao.ExpressBase.Get(ctx, expressId)
		if err != nil {
			return false, err
		}
		// 设置退货跟踪名称
		orderReturn.ReturnTrackingName = expressBase.ExpressName
	}

	// 编辑退货信息
	_, err := s.Edit(ctx, orderReturn)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Cancel 取消订单
func (s *sOrderReturn) Cancel(ctx context.Context, returnId string, userId uint) (bool, error) {
	orderReturn, err := dao.OrderReturn.Get(ctx, returnId)
	if err != nil {
		return false, err
	}

	if orderReturn == nil {
		return false, nil
	}

	orderReturnDo := &do.OrderReturn{}
	orderReturnDo.ReturnStateId = consts.RETURN_PROCESS_CANCEL

	// 修改退款退货表 处理状态
	if _, err := dao.OrderReturn.Edit(ctx, returnId, orderReturnDo); err != nil {
		return false, gerror.New("修改退款退货表处理状态失败")
	}

	// 修改订单退货详情表 处理状态
	orderReturnItem := &do.OrderReturnItem{
		ReturnStateId: consts.RETURN_PROCESS_CANCEL,
	}

	_, err = dao.OrderReturnItem.EditWhere(ctx, &do.OrderReturnItemListInput{
		Where: do.OrderReturnItem{ReturnId: returnId},
	}, orderReturnItem)

	if err != nil {
		return false, gerror.New("修改订单退货详情表处理状态失败")
	}

	orderReturnItemList, err := dao.OrderReturnItem.Find(ctx, &do.OrderReturnItemListInput{
		Where: do.OrderReturnItem{ReturnId: returnId},
	})

	var returnItemNumSum uint = 0

	if !g.IsEmpty(orderReturnItemList) {
		for _, returnItem := range orderReturnItemList {
			// 退货商品数量
			returnItemNum := returnItem.ReturnItemNum
			returnItemNumSum += returnItemNum

			// 退款总额
			returnItemSubtotal := returnItem.ReturnItemSubtotal
			// 订单商品信息
			orderItemId := returnItem.OrderItemId
			orderItem, err := dao.OrderItem.Get(ctx, orderItemId)
			if err != nil {
				return false, gerror.New("修改订单商品表退货信息失败")
			}

			if orderItem != nil {
				orderItemDo := &do.OrderItem{}
				orderItemDo.OrderItemReturnNum = orderItem.OrderItemReturnNum - returnItemNum
				orderItemDo.OrderItemReturnSubtotal = orderItem.OrderItemReturnSubtotal - returnItemSubtotal
				// 修改订单商品表退货信息
				if _, err := dao.OrderItem.Edit(ctx, orderItemId, orderItemDo); err != nil {
					return false, gerror.New("修改订单商品表退货信息失败")
				}
			}
		}
	}

	// 修正申请退款总额
	orderData, err := dao.OrderData.Get(ctx, orderReturn.OrderId)
	if err != nil {
		return false, gerror.New("修正申请退款总额失败")
	}

	if orderData != nil {
		orderDataDo := &do.OrderData{}
		orderDataDo.OrderRefundAmount = orderData.OrderRefundAmount - orderReturn.ReturnRefundAmount
		orderDataDo.OrderReturnNum = orderData.OrderReturnNum - returnItemNumSum

		if _, err := dao.OrderData.Edit(ctx, orderReturn.OrderId, orderDataDo); err != nil {
			return false, gerror.New("修正申请退款总额失败")
		}
	}

	return true, nil
}

// ReturnItem 读取订单退货详情
func (s *sOrderReturn) ReturnItem(ctx context.Context, orderId string, orderItemId string, userId uint) (*model.OrderReturnItemVo, error) {
	orderReturnItemVo := &model.OrderReturnItemVo{}

	// 订单信息
	orderInfo, err := dao.OrderInfo.Get(ctx, orderId)
	if err != nil {
		return nil, err
	}

	if orderInfo == nil {
		return nil, nil
	}

	if userId != orderInfo.UserId {
		return nil, gerror.New("无权访问该订单")
	}

	// 订单商品表
	orderItem, err := dao.OrderItem.Get(ctx, orderItemId)
	if err != nil {
		return nil, err
	}

	if orderItem == nil {
		return nil, nil
	}

	if orderItem.OrderId != orderInfo.OrderId {
		return nil, gerror.New("订单信息有误")
	}

	// 商品信息
	gconv.Struct(orderItem, orderReturnItemVo)

	itemRefundAmount := decimal.NewFromFloat(0)
	var itemReturnNum uint = 0

	// 退款退货信息
	returnStateIds := []uint{
		consts.RETURN_PROCESS_FINISH,
		consts.RETURN_PROCESS_CHECK,
		consts.RETURN_PROCESS_RECEIVED,
		consts.RETURN_PROCESS_REFUND,
		consts.RETURN_PROCESS_RECEIPT_CONFIRMATION,
		consts.RETURN_PROCESS_REFUSED,
		consts.RETURN_PROCESS_SUBMIT,
	}

	orderReturns, err := dao.OrderReturn.Find(ctx, &do.OrderReturnListInput{
		Where: do.OrderReturn{
			OrderId:       orderId,
			ReturnStateId: returnStateIds,
		},
	})
	if err != nil {
		return nil, err
	}

	if len(orderReturns) > 0 {
		returnIds := garray.NewStrArray()
		for _, orderReturn := range orderReturns {
			returnIds.Append(orderReturn.ReturnId)
		}

		// 订单退货详情表信息
		if returnIds.Len() > 0 {
			orderReturnItems, err := dao.OrderReturnItem.Find(ctx, &do.OrderReturnItemListInput{
				Where: do.OrderReturnItem{
					ReturnId:      returnIds.Slice(),
					OrderItemId:   orderItemId,
					ReturnStateId: returnStateIds,
				},
			})
			if err != nil {
				return nil, err
			}

			if len(orderReturnItems) > 0 {
				for _, orderReturnItem := range orderReturnItems {
					itemRefundAmount = itemRefundAmount.Add(decimal.NewFromFloat(orderReturnItem.ReturnItemSubtotal))
					itemReturnNum = itemReturnNum + orderReturnItem.ReturnItemNum
				}
			}
		}
	}

	orderItemCanRefundAmount := decimal.NewFromFloat(orderItem.OrderItemCanRefundAmount)
	orderReturnItemVo.CanRefundAmount, _ = orderItemCanRefundAmount.Sub(itemRefundAmount).Float64()
	orderReturnItemVo.CanRefundNum = orderItem.OrderItemQuantity - itemReturnNum

	// 退货原因集合
	orderReturnReasons, err := dao.OrderReturnReason.Find(ctx, &do.OrderReturnReasonListInput{
		BaseList: ml.BaseList{Sidx: dao.OrderReturnReason.Columns().ReturnReasonSort, Sort: ml.ORDER_BY_ASC},
	})

	if err != nil {
		return nil, err
	}

	if len(orderReturnReasons) > 0 {
		orderReturnItemVo.ReturnReasonList = orderReturnReasons
	}

	return orderReturnItemVo, nil
}

// AddItem 添加订单退货
func (s *sOrderReturn) AddItem(ctx context.Context, orderReturnInput *model.OrderReturnInput) (returnId string, err error) {
	userId := orderReturnInput.UserId

	// 是否有店铺
	buyerStoreId := 0
	//buyerStoreId, err := service.StoreBase().GetStoreId(ctx, userId)
	//if err != nil {
	//	return "", err
	//}

	orderItemIds := array.Column(orderReturnInput.ReturnItems, "OrderItemId")

	orderInfo, err := dao.OrderInfo.Get(ctx, orderReturnInput.OrderId)
	if err != nil {
		return "", err
	}
	if orderInfo == nil {
		return "", gerror.New("此订单信息有误！")
	}

	orderBase, err := dao.OrderBase.Get(ctx, orderReturnInput.OrderId)
	if err != nil {
		return "", err
	}
	if orderBase == nil {
		return "", gerror.New("订单详细信息有误！")
	}

	// 判断此订单商品是否有正在审核的退款单
	orderReturnQueryWrapper := &do.OrderReturnListInput{
		Where: do.OrderReturn{
			OrderId:       orderReturnInput.OrderId,
			ReturnStateId: consts.RETURN_PROCESS_CHECK,
		},
	}
	orderReturns, err := dao.OrderReturn.Find(ctx, orderReturnQueryWrapper)
	if err != nil {
		return "", err
	}

	if !g.IsEmpty(orderReturns) {
		returnIds := array.Column(orderReturns, "ReturnId")
		itemQueryWrapper := &do.OrderReturnItemListInput{
			Where: do.OrderReturnItem{
				OrderItemId: orderItemIds,
				ReturnId:    returnIds,
			},
		}
		orderReturnItems, err := dao.OrderReturnItem.Find(ctx, itemQueryWrapper)
		if err != nil {
			return "", err
		}
		if !g.IsEmpty(orderReturnItems) {
			return "", gerror.New("此订单有商品退货审核中！")
		}
	}

	if userId != orderBase.UserId {
		return "", gerror.New("无权操作")
	}

	storeId := orderBase.StoreId

	if orderReturnInput.ReturnTel != "" && !phone.IsValidNumber(orderReturnInput.ReturnTel) {
		return "", gerror.New("手机号输入有误！")
	}

	// 封装数据
	orderReturn := &do.OrderReturn{
		OrderId:            orderReturnInput.OrderId,
		BuyerUserId:        userId,
		BuyerStoreId:       buyerStoreId,
		ReturnReasonId:     orderReturnInput.ReturnReasonId,
		ReturnBuyerMessage: orderReturnInput.ReturnBuyerMessage,
		StoreId:            storeId,
		ReturnRefundAmount: 0,
		ReturnCommisionFee: 0,
		ReturnShippingFee:  0,
		ReturnStateId:      consts.RETURN_PROCESS_CHECK,
		ReturnTel:          orderReturnInput.ReturnTel,
		SubsiteId:          orderInfo.SubsiteId,
		ReturnFlag:         orderReturnInput.ReturnFlag,
	}

	orderData, err := dao.OrderData.Get(ctx, orderReturnInput.OrderId)
	if err != nil {
		return "", err
	}
	orderShippingFee := orderData.OrderShippingFee

	if orderReturnInput.ReviewFlag {
		orderReturn.ReturnShippingFee = orderShippingFee
	} else if orderInfo.OrderIsOut == consts.ORDER_PICKING_STATE_NO {
		orderReturn.ReturnShippingFee = orderShippingFee
	}

	var returnItems []*do.OrderReturnItem
	for _, itemInputVo := range orderReturnInput.ReturnItems {
		if itemInputVo.ReturnItemNum <= 0 {
			return "", gerror.New(fmt.Sprintf("%d: 退款数量不正确！", itemInputVo.OrderItemId))
		}

		returnItem := &do.OrderReturnItem{
			OrderItemId:        itemInputVo.OrderItemId,
			OrderId:            orderReturnInput.OrderId,
			ReturnItemNum:      itemInputVo.ReturnItemNum,
			ReturnItemSubtotal: itemInputVo.ReturnRefundAmount,
			ReturnReasonId:     orderReturnInput.ReturnReasonId,
			ReturnItemNote:     orderReturnInput.ReturnBuyerMessage,
			ReturnItemImage:    orderReturnInput.ReturnItemImage,
			ReturnStateId:      consts.RETURN_PROCESS_CHECK,
		}

		returnItems = append(returnItems, returnItem)
	}

	if _, err := s.AddReturnByItem(ctx, orderReturn, returnItems); err != nil {
		return "", gerror.New("退货添加失败")
	}

	returnId = orderReturn.ReturnId.(string)

	if orderReturnInput.ReviewFlag {
		if orderInfo.OrderIsShipped == consts.ORDER_SHIPPED_STATE_NO {
			orderReturn.ReturnStoreMessage = "取消订单，自动审核"
			if _, err := s.Review(ctx, orderReturn, 0); err != nil {
				return "", err
			}
		}
	} else if orderInfo.OrderIsOut == consts.ORDER_PICKING_STATE_NO {
		orderReturn.ReturnStoreMessage = "未出库，自动审核"
		if _, err := s.Review(ctx, orderReturn, 0); err != nil {
			return "", err
		}
	}

	return returnId, nil
}

func (s *sOrderReturn) AddReturnByItem(ctx context.Context, orderReturn *do.OrderReturn, returnItems []*do.OrderReturnItem) (bool, error) {
	// 获取 orderItemIds 列表
	orderItemIds := array.Column(returnItems, "OrderItemId")

	// 根据 orderItemIds 获取订单商品信息
	orderItems, err := dao.OrderItem.Gets(ctx, orderItemIds)
	if err != nil {
		return false, err
	}
	if g.IsEmpty(orderItems) {
		return false, errors.New("订单商品信息有误！")
	}

	orderId := orderReturn.OrderId

	// 获取订单数据
	orderData, err := dao.OrderData.Get(ctx, orderId)
	if err != nil {
		return false, err
	}

	for _, returnItem := range returnItems {
		orderItem := findOrderItem(orderItems, returnItem.OrderItemId.(uint64))
		if orderItem == nil {
			return false, errors.New("订单商品信息有误！")
		}

		returnItemSubtotal := returnItem.ReturnItemSubtotal.(float64)

		orderItemCanRefundAmount := orderItem.OrderItemCanRefundAmount
		orderItemReturnAgreeAmount := orderItem.OrderItemReturnAgreeAmount
		newAmount := orderItemCanRefundAmount - orderItemReturnAgreeAmount

		if returnItemSubtotal < 0 || returnItemSubtotal < newAmount {
			return false, errors.New("退货单金额错误！")
		}

		orderItemQuantity := orderItem.OrderItemQuantity
		returnNum := orderItem.OrderItemReturnAgreeNum
		returnItemNum := returnItem.ReturnItemNum

		if (orderItemQuantity - returnNum) < returnItemNum.(uint) {
			return false, errors.New("退货单商品数量错误！")
		}

		// 更新退货单金额
		returnRefundAmount := orderReturn.ReturnRefundAmount.(float64)
		orderReturn.ReturnRefundAmount = returnRefundAmount + returnItemSubtotal

		// 退还佣金计算
		orderItemCommissionRate := orderItem.OrderItemCommissionRate
		returnItem.ReturnItemCommisionFee = returnItemSubtotal * orderItemCommissionRate / 100

		returnCommisionFee := orderReturn.ReturnCommisionFee.(float64)
		orderReturn.ReturnCommisionFee = returnCommisionFee + returnItem.ReturnItemCommisionFee.(float64)

		// 修改订单数据
		orderRefundAmount := orderData.OrderRefundAmount
		orderData.OrderRefundAmount = orderRefundAmount + returnItemSubtotal

		orderReturnNum := orderData.OrderReturnNum
		orderData.OrderReturnNum = orderReturnNum + returnItemNum.(uint)

		// 修改订单SKU
		orderItemDo := &do.OrderItem{}
		orderItemDo.OrderItemReturnNum = orderItem.OrderItemReturnNum + returnItemNum.(uint)
		orderItemDo.OrderItemReturnSubtotal = orderItem.OrderItemReturnSubtotal + returnItemSubtotal

		if _, err = dao.OrderItem.Edit(ctx, orderItem.OrderItemId, orderItemDo); err != nil {
			return false, errors.New("修改订单信息失败！")
		}
	}

	// 退运费
	returnShippingFee := orderReturn.ReturnShippingFee.(float64)
	orderReturn.ReturnRefundAmount = returnShippingFee + orderReturn.ReturnRefundAmount.(float64)

	orderReturn.ReturnAddTime = time.Now().UnixMilli()
	orderReturn.ReturnType = consts.ORDER_RETURN

	returnId, err := service.NumberSeq().GetNextSeqString(ctx, "RT")
	if err != nil {
		return false, err
	}
	orderReturn.ReturnId = returnId

	for _, returnItem := range returnItems {
		returnItem.ReturnId = returnId
	}

	if _, err = s.Add(ctx, orderReturn); err != nil {
		return false, errors.New("保存退货基础单失败！")
	}

	if _, err = dao.OrderReturnItem.Saves(ctx, returnItems); err != nil {
		return false, errors.New("保存退货单失败！")
	}

	if orderData != nil {
		orderDataDo := &do.OrderData{}
		orderDataDo.OrderRefundAmount = orderData.OrderRefundAmount
		orderDataDo.OrderReturnNum = orderData.OrderReturnNum
		if _, err = dao.OrderData.Edit(ctx, orderData.OrderId, orderDataDo); err != nil {
			return false, errors.New("修改订单详细信息失败！")
		}
	}

	return true, nil
}

// findOrderItem 根据 orderItemId 从 orderItems 列表中找到对应的 OrderItem
func findOrderItem(orderItems []*entity.OrderItem, orderItemId uint64) *entity.OrderItem {
	for _, item := range orderItems {
		if item.OrderItemId == orderItemId {
			return item
		}
	}
	return nil
}
