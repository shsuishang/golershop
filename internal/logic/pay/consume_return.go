package pay

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"github.com/shopspring/decimal"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/log"
)

type sConsumeReturn struct{}

func init() {
	service.RegisterConsumeReturn(NewConsumeReturn())
}

func NewConsumeReturn() *sConsumeReturn {
	return &sConsumeReturn{}
}

// DoRefund 执行退款操作
func (s *sConsumeReturn) DoRefund(ctx context.Context, orderReturns []*entity.OrderReturn) bool {
	paidReturnIds := make([]string, 0)

	// 原理退回标记
	orderRefundFlag := service.ConfigBase().GetBool(ctx, "order_refund_flag", false)
	orderIds := make([]string, 0)
	userIds := make([]uint, 0)
	returnIds := make([]string, 0)

	for _, orderReturn := range orderReturns {
		orderIds = append(orderIds, orderReturn.OrderId)
		userIds = append(userIds, orderReturn.BuyerUserId)
		returnIds = append(returnIds, orderReturn.ReturnId)
	}

	orderDataList, _ := dao.OrderData.Gets(ctx, orderIds)
	userResourceList, _ := dao.UserResource.Gets(ctx, userIds)
	orderInfoList, _ := dao.OrderInfo.Gets(ctx, orderIds)

	orderReturnItems, _ := dao.OrderReturnItem.Find(ctx, &do.OrderReturnItemListInput{Where: do.OrderReturnItem{ReturnId: returnIds}})
	userInfoList, _ := dao.UserInfo.Gets(ctx, userIds)
	curDate := gtime.Now()
	ymdDate := gtime.Now()
	pointsVaueRate := service.ConfigBase().GetFloat(ctx, "points_vaue_rate", 0.0)

	for _, orderReturn := range orderReturns {
		userId := orderReturn.BuyerUserId
		if userId == 0 {
			panic("买家信息有误")
		}
		buyerStoreId := orderReturn.BuyerStoreId
		storeId := orderReturn.StoreId
		var userResource *entity.UserResource
		for _, item := range userResourceList {
			if item.UserId == userId {
				userResource = item
				break
			}
		}

		// 判断是否需要退佣金
		var returnCommisionFee decimal.Decimal

		// 不是退运费
		orderId := orderReturn.OrderId
		returnIsShippingFee := orderReturn.ReturnIsShippingFee
		if !returnIsShippingFee {
			withdrawReceivedDay := service.ConfigBase().GetFloat(ctx, "withdraw_received_day", 7.0)

			if withdrawReceivedDay == 0 {
				withdrawReceivedDay = 7.0
			}

			if withdrawReceivedDay >= 0 {
				var orderInfo *entity.OrderInfo
				for _, item := range orderInfoList {
					if item.OrderId == orderId {
						orderInfo = item

						orderStateId := orderInfo.OrderStateId
						orderIsPaid := orderInfo.OrderIsPaid

						// 未到可结算时间可退佣金
						if orderStateId != consts.ORDER_STATE_FINISH && orderIsPaid != consts.ORDER_PAID_STATE_YES {
							returnCommisionFee = decimal.NewFromFloat(orderReturn.ReturnCommisionFee)
						}

						break
					}
				}
			}
		}

		waitingRefundAmount := decimal.NewFromFloat(orderReturn.ReturnRefundAmount)
		if !waitingRefundAmount.IsZero() {
			var orderData *entity.OrderData
			for _, item := range orderDataList {
				if item.OrderId == orderId {
					orderData = item
					break
				}
			}

			orderPointsFee := orderData.OrderPointsFee
			orderRefundAgreePoints := orderData.OrderRefundAgreePoints

			buyerUserMoney := waitingRefundAmount
			buyerUserPoints := decimal.Zero

			//sellerUserMoney := waitingRefundAmount.Neg()
			returnId := orderReturn.ReturnId

			// 写入流水
			buyerConsumeRecord := &do.ConsumeRecord{
				OrderId:      returnId,
				UserId:       userId,
				StoreId:      buyerStoreId,
				UserNickname: userInfoList[0].UserNickname,
				RecordDate:   ymdDate,
				RecordYear:   curDate.Year(),
				RecordMonth:  int(curDate.Month()),
				RecordDay:    curDate.Day(),
				RecordTitle:  fmt.Sprintf("退款单:%s", returnId),
				RecordTime:   curDate.Unix(),
				PaymentMetId: consts.PAYMENT_MET_MONEY,
				RecordMoney:  waitingRefundAmount,
				TradeTypeId:  consts.TRADE_TYPE_REFUND_GATHERING,
			}

			// 卖家流水记录
			sellerConsumeRecord := &do.ConsumeRecord{
				OrderId:             returnId,
				UserId:              consts.ADMIN_PLANTFORM_USERID,
				StoreId:             storeId,
				RecordMoney:         waitingRefundAmount.Neg().Add(returnCommisionFee),
				RecordCommissionFee: returnCommisionFee.Neg(),
				TradeTypeId:         consts.TRADE_TYPE_REFUND_PAY,
			}

			orderDataDo := &do.OrderData{}
			orderDataDo.OrderRefundAgreeAmount, _ = waitingRefundAmount.Add(decimal.NewFromFloat(orderData.OrderRefundAgreeAmount)).Float64()

			// 读取退款单项目
			orderItems := make([]*do.OrderItem, 0)
			for _, orderReturnItem := range orderReturnItems {
				if orderReturnItem.ReturnId == returnId {
					var orderItem *do.OrderItem
					for _, item := range orderItems {
						if item.OrderItemId == orderReturnItem.OrderItemId {
							orderItem = item
							break
						}
					}

					returnItemSubtotal := orderReturnItem.ReturnItemSubtotal
					returnItemNum := orderReturnItem.ReturnItemNum
					orderItemReturnAgreeAmount := orderItem.OrderItemReturnAgreeAmount.(float64)
					orderItemReturnAgreeNum := orderItem.OrderItemReturnAgreeNum.(uint)

					orderItem.OrderItemReturnAgreeAmount = orderItemReturnAgreeAmount + returnItemSubtotal
					orderItem.OrderItemReturnAgreeNum = orderItemReturnAgreeNum + returnItemNum

					// 未结算才发放用金
					if !returnCommisionFee.IsZero() {
						returnItemCommisionFee := orderReturnItem.ReturnItemCommisionFee
						returnItemCommisionFeeRefund := orderItem.OrderItemCommissionFeeRefund
						orderItem.OrderItemCommissionFeeRefund = returnItemCommisionFeeRefund.(float64) + returnItemCommisionFee
					}
					orderItems = append(orderItems, orderItem)
				}
			}

			if len(orderItems) > 0 {
				if _, err := dao.OrderItem.Saves(ctx, orderItems); err != nil {
					panic("修改订单商品数据失败")
				}
			}

			// 买家数据
			if _, err := dao.ConsumeRecord.Save(ctx, buyerConsumeRecord); err != nil {
				panic("增加买家流水数据失败")
			}

			// 如果混合了积分，优先退积分
			if orderPointsFee > 0 && orderPointsFee > orderRefundAgreePoints {
				refundPoints := decimal.NewFromFloat(orderPointsFee - orderRefundAgreePoints)
				buyerUserMoney = buyerUserMoney.Sub(refundPoints)

				if pointsVaueRate > 0 {
					buyerUserPoints = refundPoints.Div(decimal.NewFromFloat(pointsVaueRate))
				}
				orderData.OrderRefundAgreePoints, _ = refundPoints.Add(decimal.NewFromFloat(orderData.OrderRefundAgreePoints)).Float64()
			}

			// 操作退款数据
			flag, err := s.DoRefundOrder(ctx, orderRefundFlag, userId, storeId, userResource, orderId, buyerUserMoney, buyerUserPoints, returnId)

			if err != nil {
				panic("退款失败")
			}

			if !flag {
				panic("退款失败")
			}

			// 卖家数据
			if _, err := dao.ConsumeRecord.Save(ctx, sellerConsumeRecord); err != nil {
				panic("增加卖家流水数据失败")
			}

			// 更新订单主表
			if _, err := dao.OrderData.Edit(ctx, orderId, orderDataDo); err != nil {
				panic("修改订单主表数据失败")
			}

			paidReturnIds = append(paidReturnIds, orderReturn.ReturnId)
		}
	}

	// 操作成功
	if len(paidReturnIds) == len(orderReturns) {
		// 标记退款完成
		if orderRefundFlag {
			// 标记退款完成
			if _, err := s.SetReturnPaidYes(ctx, paidReturnIds); err != nil {
				panic("标记退款状态失败")
			}
		}
		return true
	}
	return false
}

// SetReturnPaidYes 修改为退款已支付状态
func (s *sConsumeReturn) SetReturnPaidYes(ctx context.Context, returnIds []string) (bool, error) {
	if g.IsEmpty(returnIds) {
		return false, nil
	}

	orderReturn := &do.OrderReturn{
		ReturnIsPaid: true,
	}
	orderReturn.ReturnIsPaid = false

	// 构建查询条件
	returnQueryWrapper := &do.OrderReturnListInput{
		Where: do.OrderReturn{
			ReturnId: returnIds,
		},
	}

	// 更新订单退货信息
	_, err := dao.OrderReturn.Edit(ctx, returnQueryWrapper, orderReturn)
	if err != nil {
		return false, gerror.New("更新订单退货信息失败")
	}

	// 获取订单退货列表
	orderReturnList, err := dao.OrderReturn.Gets(ctx, returnIds)
	if err != nil {
		return false, gerror.New("获取订单退货列表失败")
	}

	// 提取订单ID
	var orderIds []string
	for _, orderReturn := range orderReturnList {
		orderIds = append(orderIds, orderReturn.OrderId)
	}

	//orderIds = array.ArrayUnique(orderIds)

	if g.IsEmpty(orderIds) {
		return false, nil
	}

	// 判断是否存在用金额退款
	orderQueryWrapper := &do.DistributionOrderListInput{
		Where: do.DistributionOrder{
			OrderId:  orderIds,
			UoActive: 1,
		},
	}
	distributionOrderList, err := dao.DistributionOrder.Find(ctx, orderQueryWrapper)
	if err != nil {
		return false, gerror.New("查询分销订单失败")
	}

	var userIds []uint
	for _, distributionOrder := range distributionOrderList {
		userIds = append(userIds, distributionOrder.UserId)
	}
	//userIds = g.SliceUnique(userIds)

	// 获取分销佣金列表
	commissionList, err := dao.DistributionCommission.Gets(ctx, userIds)
	if err != nil {
		return false, gerror.New("获取分销佣金列表失败")
	}

	commissionListDo := make([]*do.DistributionCommission, 0)
	if !g.IsEmpty(distributionOrderList) {
		for _, distributionOrder := range distributionOrderList {
			uoBuyCommission := decimal.NewFromFloat(distributionOrder.UoBuyCommission)
			uoDirectsellerCommission := distributionOrder.UoDirectsellerCommission
			addCommissionRefundAmount := uoBuyCommission.Add(decimal.NewFromFloat(uoDirectsellerCommission))

			userId := distributionOrder.UserId
			var commission *entity.DistributionCommission
			for _, c := range commissionList {
				if c.UserId == userId {
					commission = c
					break
				}
			}

			if commission != nil {
				commissionRefundAmount := gconv.Float64(commission.CommissionRefundAmount)
				commission.CommissionRefundAmount, _ = addCommissionRefundAmount.Add(decimal.NewFromFloat(commissionRefundAmount)).Float64()

				commissionListDo = append(commissionListDo, &do.DistributionCommission{
					CommissionRefundAmount: commission.CommissionRefundAmount,
				})
			}
		}

		if !g.IsEmpty(commissionList) {
			_, err = dao.DistributionCommission.Saves(ctx, commissionListDo)
			if err != nil {
				return false, gerror.New("更新分销佣金失败")
			}
		}

		uoIds := make([]uint, 0)
		for _, distributionOrder := range distributionOrderList {
			uoIds = append(uoIds, distributionOrder.UoId)
		}

		//uoIds = g.SliceUnique(uoIds)

		distributionOrder := &do.DistributionOrder{
			UoActive: true,
		}
		distributionOrderQueryWrapper := &do.DistributionOrderListInput{
			Where: do.DistributionOrder{
				UoId: uoIds,
			},
		}

		_, err = dao.DistributionOrder.EditWhere(ctx, distributionOrderQueryWrapper, distributionOrder)
		if err != nil {
			return false, gerror.New("更新分销订单失败")
		}

		distributionOrderItem := &do.DistributionOrderItem{
			UoiActive: true,
		}
		orderItemQueryWrapper := &do.DistributionOrderItemListInput{
			Where: do.DistributionOrderItem{
				OrderId: orderIds,
			},
		}
		_, err = dao.DistributionOrderItem.Edit(ctx, orderItemQueryWrapper, distributionOrderItem)
		if err != nil {
			return false, gerror.New("更新分销订单项失败")
		}
	}

	return true, nil
}

// DoRefundOrder 操作退款数据
func (s *sConsumeReturn) DoRefundOrder(ctx context.Context, orderRefundFlag bool, userId, storeId uint, userResource *entity.UserResource, orderId string, buyerUserMoney, buyerUserPoints decimal.Decimal, returnId string) (bool, error) {
	if orderRefundFlag {
		// 读取在线支付信息，如果无在线支付信息，则余额支付，否则在线支付【联合支付】判断
		depositQueryWrapper := &do.ConsumeDepositListInput{
			Where: do.ConsumeDeposit{},
		}
		if !g.IsEmpty(orderId) {
			var likes = []*ml.WhereExt{{
				Column: dao.ConsumeDeposit.Columns().OrderId,
				Val:    orderId,
				Symbol: ml.FIND_IN_SET_STR,
			}}

			depositQueryWrapper.WhereExt = likes
		}

		consumeDeposit, err := dao.ConsumeDeposit.FindOne(ctx, depositQueryWrapper)
		if err != nil {
			return false, err
		}
		if consumeDeposit != nil {
			paymentChannelId := consumeDeposit.PaymentChannelId
			depositTotalFee := decimal.NewFromFloat(consumeDeposit.DepositTotalFee)
			channelCode, err := service.ConfigBase().GetPaymentChannelCode(ctx, paymentChannelId)
			if err != nil {
				return false, err
			}

			// 微信，支付宝支付
			if gstr.InArray([]string{"alipay", "wxpay"}, channelCode) {
				dMoney := buyerUserMoney.Sub(depositTotalFee).Round(2)
				if dMoney.Cmp(decimal.Zero) > 0 {
					userResourceDo := &do.UserResource{}
					userResourceDo.UserMoney, _ = decimal.NewFromFloat(userResource.UserMoney).Add(dMoney).Float64()
					if _, err := dao.UserResource.Edit(ctx, userResource.UserId, userResourceDo); err != nil {
						return false, gerror.New("用户退款失败")
					}
				}

				if buyerUserPoints.Cmp(decimal.Zero) > 0 {
					points, _ := buyerUserPoints.Float64()
					if _, err := service.UserResource().Points(ctx, &model.UserPointsVo{
						UserId:       userId,
						Points:       points,
						PointsTypeId: consts.POINTS_TYPE_CONSUME_RETRUN,
						OrderId:      returnId,
						StoreId:      storeId,
					}); err != nil {
						return false, gerror.New("用户退积分失败")
					}
				}

				orderReturn := &do.OrderReturn{
					ReturnChannelCode:  channelCode,
					DepositTradeNo:     consumeDeposit.DepositTradeNo,
					PaymentChannelId:   paymentChannelId,
					TradePaymentAmount: depositTotalFee,
				}
				if _, err := dao.OrderReturn.Edit(ctx, returnId, orderReturn); err != nil {
					return false, gerror.New("修改退单信息失败")
				}
				// 执行第三方接口退款流程
				if err := s.DoOnlineRefund(ctx, returnId); err != nil {
					return false, err
				}
				return false, nil
			}
		}

		if buyerUserMoney.Cmp(decimal.Zero) > 0 {
			userResourceDo := &do.UserResource{}
			userResourceDo.UserMoney, _ = buyerUserMoney.Add(decimal.NewFromFloat(userResource.UserMoney)).Float64()
			if _, err := dao.UserResource.Edit(ctx, userResource.UserId, userResourceDo); err != nil {
				return false, gerror.New("用户退款失败")
			}
		}

		if buyerUserPoints.Cmp(decimal.Zero) > 0 {
			points, _ := buyerUserPoints.Float64()
			if _, err := service.UserResource().Points(ctx, &model.UserPointsVo{
				UserId:       userId,
				Points:       points,
				PointsTypeId: consts.POINTS_TYPE_CONSUME_RETRUN,
				OrderId:      returnId,
				StoreId:      storeId,
			}); err != nil {
				return false, gerror.New("用户退积分失败")
			}
		}

		orderReturn := &do.OrderReturn{
			ReturnChannelFlag: 1,
		}

		if _, err := dao.OrderReturn.Edit(ctx, returnId, orderReturn); err != nil {
			return false, gerror.New("修改退单信息失败")
		}
	}
	return true, nil
}

// doOnLineRefund 执行线上支付退款
func (s *sConsumeReturn) DoOnlineRefund(ctx context.Context, returnId string) error {
	// 获取OrderReturn信息
	orderReturn, err := dao.OrderReturn.Get(ctx, returnId)
	if err != nil {
		return err
	}

	returnChannelCode := orderReturn.ReturnChannelCode
	depositTradeNo := orderReturn.DepositTradeNo
	tradePaymentAmount := orderReturn.TradePaymentAmount

	shopOrderReturn := &entity.OrderReturn{
		ReturnChannelFlag: 1,
		OrderId:           orderReturn.OrderId,
	}

	// 进行退款处理
	switch returnChannelCode {
	case "alipay":
		err = s.doAliPayRefund(ctx, returnId, depositTradeNo, tradePaymentAmount, shopOrderReturn)
	case "wxpay":
		err = s.doWxPayRefund(ctx, depositTradeNo, tradePaymentAmount, shopOrderReturn)
	}

	if err != nil {
		log.Error(ctx, err)
		return err
	}

	// 更新退款状态
	return s.updateRefundOrderReturn(ctx, shopOrderReturn)
}

// doAliPayRefund 支付宝退款
func (s *sConsumeReturn) doAliPayRefund(ctx context.Context, returnId, depositTradeNo string, tradePaymentAmount float64, shopOrderReturn *entity.OrderReturn) error {
	// 实现支付宝退款逻辑
	// ...
	return nil
}

// doWxPayRefund 微信退款
func (s *sConsumeReturn) doWxPayRefund(ctx context.Context, depositTradeNo string, tradePaymentAmount float64, shopOrderReturn *entity.OrderReturn) error {
	// 实现微信退款逻辑
	// ...
	return nil
}

// updateRefundOrderReturn 更新退款订单
func (s *sConsumeReturn) updateRefundOrderReturn(ctx context.Context, shopOrderReturn *entity.OrderReturn) error {
	// 实现更新退款订单逻辑
	// ...
	return nil
}
