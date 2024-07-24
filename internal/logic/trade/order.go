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
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	gconv "github.com/gogf/gf/v2/util/gconv"
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
	"math"
	"sync"
	"time"
)

type sOrder struct{}

func init() {
	service.RegisterOrder(NewOrder())
}

func NewOrder() *sOrder {
	return &sOrder{}
}

// Detail 读取订单
func (s *sOrder) Detail(ctx context.Context, orderId any) (detail *model.OrderVo, err error) {
	detail = &model.OrderVo{}

	//信息表
	orderInfo, err := dao.OrderInfo.Get(ctx, orderId)

	if err != nil {
		return nil, err
	}

	if err := gconv.Scan(orderInfo, detail); err != nil {
		return nil, err
	}

	//详情表
	orderData, err := dao.OrderData.Get(ctx, orderId)
	if err != nil {
		return nil, err
	}

	if err := gconv.Scan(orderData, &detail.OrderData); err != nil {
		return nil, err
	}

	//基础表
	orderBase, err := dao.OrderBase.Get(ctx, orderId)

	if err != nil {
		return nil, err
	}

	detail.OrderNumber = orderBase.OrderNumber
	detail.OrderTime = orderBase.OrderTime
	detail.OrderProductAmount = orderBase.OrderProductAmount
	detail.OrderPaymentAmount = orderBase.OrderPaymentAmount
	detail.CurrencyId = orderBase.CurrencyId
	detail.CurrencySymbolLeft = orderBase.CurrencySymbolLeft
	detail.StoreName = orderBase.StoreName
	detail.UserNickname = orderBase.UserNickname

	//读取订单商品
	orderItems, err := dao.OrderItem.Find(ctx, &do.OrderItemListInput{Where: do.OrderItem{OrderId: orderId}})

	if err != nil {
		return nil, err
	}

	gconv.Scan(orderItems, &detail.Items)

	if !g.IsEmpty(detail.Items) {
		for i := range detail.Items {
			detail.Items[i].ProductItemName = fmt.Sprintf("%s %s",
				detail.Items[i].ProductName,
				gstr.Replace(detail.Items[i].ItemName, ",", " "),
			)
		}
	}

	//售后服务
	returnItemList, err := dao.OrderReturnItem.Find(ctx, &do.OrderReturnItemListInput{Where: do.OrderReturnItem{OrderId: orderId, ReturnStateId: consts.RETURN_PROCESS_CANCEL}})

	if err != nil {
		return nil, err
	}

	// 处理为map
	orderReturnItemMap := make(map[uint64][]string)

	for _, item := range returnItemList {
		if _, ok := orderReturnItemMap[item.OrderItemId]; !ok {
			orderReturnItemMap[item.OrderItemId] = []string{}
		}

		orderReturnItemMap[item.OrderItemId] = append(orderReturnItemMap[item.OrderItemId], item.ReturnId)
	}

	for i := range detail.Items {
		// 是否可以退货
		detail.Items[i].IfReturn = detail.Items[i].OrderItemQuantity > detail.Items[i].OrderItemReturnNum && s.ifReturn(orderInfo.OrderStateId, orderInfo.OrderIsPaid)

		returnIds := orderReturnItemMap[detail.Items[i].OrderItemId]

		if g.IsEmpty(returnIds) {
			returnIds = []string{}
		}

		detail.Items[i].ReturnIds = returnIds
	}

	//商品库存
	itemIds := array.Column(orderItems, "item_id")
	productItems, err := service.ProductItem().Gets(ctx, itemIds)

	if err != nil {
		return nil, err
	}

	for _, it := range productItems {
		i := &entity.WarehouseItem{}

		gconv.Scan(it, i)
		i.WarehouseId = 0
		i.ItemId = it.ItemId
		i.WarehouseItemQuantity = it.ItemQuantity

		detail.WarehouseItems = append(detail.WarehouseItems, i)
	}

	//配送地址
	orderDeliveryAddress, err := dao.OrderDeliveryAddress.Get(ctx, orderId)
	if err != nil {
		return nil, err
	}

	detail.Delivery = orderDeliveryAddress

	//订单日志

	//orderStateLogItems, err := dao.OrderStateLog.Find(ctx, &do.OrderStateLogListInput{Where: do.OrderStateLog{OrderId: orderId}})
	//if err != nil {
	//	return nil, err
	//}
	//
	//detail.LogItems = orderStateLogItems

	//物流记录
	orderLogistics, err := dao.OrderLogistics.Find(ctx, &do.OrderLogisticsListInput{Where: do.OrderLogistics{OrderId: orderId}})
	if err != nil {
		return nil, err
	}

	detail.Logistics = orderLogistics

	//StockBill
	stockBill, err := service.StockBill().FindDetail(ctx, &do.StockBillListInput{Where: do.StockBill{OrderId: orderId}})
	if err != nil {
		return nil, err
	}

	detail.StockBill = stockBill

	//consumeRecord
	consumeRecord, err := dao.ConsumeRecord.Find(ctx, &do.ConsumeRecordListInput{Where: do.ConsumeRecord{OrderId: orderId, TradeTypeId: consts.TRADE_TYPE_SALES}})
	if err != nil {
		return nil, err
	}

	detail.ConsumeRecord = consumeRecord

	//ConsumeTrade
	consumeTrades, err := dao.ConsumeTrade.Find(ctx, &do.ConsumeTradeListInput{Where: do.ConsumeTrade{OrderId: orderId}})
	if err != nil {
		return nil, err
	}

	if len(consumeTrades) > 0 {
		detail.ConsumeTrade = consumeTrades[0]
	}

	// 订单倒计时
	showCancelTime := service.ConfigBase().GetBool(ctx, "show_cancel_time", false)
	orderAutocancelTime := service.ConfigBase().GetFloat(ctx, "order_autocancel_time", 0.0)
	if showCancelTime && orderAutocancelTime > 0 && detail.OrderStateId == consts.ORDER_STATE_WAIT_PAY {
		orderTime := detail.OrderTime
		orderTime = orderTime.Add(time.Duration(int(orderAutocancelTime*60*60)) * time.Second)
		remainPayTime := (orderTime.Unix() - time.Now().Unix())
		detail.RemainPayTime = remainPayTime
	}

	// 是否拼团成功订单
	if len(orderInfo.ActivityTypeId) > 0 {
		activityTypeId := gconv.SliceUint(gstr.Split(orderInfo.ActivityTypeId, ","))

		if array.InArray(activityTypeId, consts.ACTIVITY_TYPE_GROUPBOOKING) {
			// 拼团是否可以审核
			groupbookingHistorys, _ := dao.ActivityGroupbookingHistory.Find(ctx, &do.ActivityGroupbookingHistoryListInput{Where: do.ActivityGroupbookingHistory{OrderId: orderId}})
			if len(groupbookingHistorys) > 0 {
				history := groupbookingHistorys[0]
				detail.ActivityGroupbookingHistory = history
				detail.OrderIsGroupbookingSuccess = history.GbEnable == consts.ACTIVITY_GROUPBOOKING_SUCCESS
			}
		}
	}

	return detail, nil
}

// List 订单搜索查询列表
func (s *sOrder) List(ctx context.Context, in *do.OrderInfoListInput) (out *model.OrderListOutput, err error) {
	//修正 待配货和代发货 为同一种状态查询
	if len(in.BaseList.WhereExt) > 0 {
		if !g.IsEmpty(in.Where.OrderStateId) {
			if in.Where.OrderStateId == consts.ORDER_STATE_WAIT_SHIPPING || in.Where.OrderStateId == consts.ORDER_STATE_PICKING {
				in.Where.OrderStateId = []uint{consts.ORDER_STATE_PICKING, consts.ORDER_STATE_WAIT_SHIPPING}
			}
		}

		for _, c := range in.BaseList.WhereExt {
			if c.Column == "OrderStateId" && c.Symbol == ml.EQ {
				if gconv.Uint(c.Val) == consts.ORDER_STATE_WAIT_SHIPPING || gconv.Uint(c.Val) == consts.ORDER_STATE_PICKING {
					c.Symbol = ml.IN
					c.Val = []uint{consts.ORDER_STATE_PICKING, consts.ORDER_STATE_WAIT_SHIPPING}
				}
			}
		}
	}

	orderInfoList, err := dao.OrderInfo.List(ctx, in)
	gconv.Scan(orderInfoList, &out)

	// 是否可以取消
	if len(out.Items) > 0 {
		for _, item := range out.Items {
			item.IfBuyerCancel = s.ifCancel(item.OrderStateId, item.OrderIsPaid)
		}
	}

	// 查询订单Id列表
	orderIdRow := array.Column(out.Items, "OrderId")

	// 判断是否有退款退货
	var orderReturnIdList []interface{}
	if len(orderIdRow) > 0 {
		// 查询退款信息
		orderReturnList, err := dao.OrderReturn.Find(ctx, &do.OrderReturnListInput{
			Where: do.OrderReturn{OrderId: orderIdRow},
		})
		if err == nil {
			// 获取退款Id列表
			orderReturnIdList = array.Column(orderReturnList, "OrderId")
		}
	}

	var orderInvoiceIdList []interface{}
	// 判断是否有订单Id
	if len(orderIdRow) > 0 {
		// 查询发票信息
		orderInvoiceList, err := dao.OrderInvoice.Find(ctx, &do.OrderInvoiceListInput{
			Where: do.OrderInvoice{OrderId: orderIdRow},
		})
		if err == nil {
			// 获取发票Id列表
			orderInvoiceIdList = array.Column(orderInvoiceList, "OrderId")
		}
	}

	//补全商品基础表信息
	ids := array.Column(orderInfoList.Items, dao.OrderInfo.Columns().OrderId)

	if len(ids) > 0 {
		orderDataList, err := dao.OrderData.Gets(ctx, ids)
		orderBaseList, err := dao.OrderBase.Gets(ctx, ids)

		//读取订单商品
		orderItem, err := dao.OrderItem.Find(ctx, &do.OrderItemListInput{Where: do.OrderItem{OrderId: ids}})

		//处理为map
		orderItemMap := make(map[string][]*entity.OrderItem)

		for _, item := range orderItem {
			orderItemMap[item.OrderId] = append(orderItemMap[item.OrderId], item)
		}

		for _, item := range orderBaseList {
			for _, vo := range out.Items {
				if item.OrderId == vo.OrderId {
					vo.OrderNumber = item.OrderNumber
					vo.OrderTime = item.OrderTime
					vo.OrderProductAmount = item.OrderProductAmount
					vo.OrderPaymentAmount = item.OrderPaymentAmount
					vo.CurrencyId = item.CurrencyId
					vo.CurrencySymbolLeft = item.CurrencySymbolLeft
					vo.StoreName = item.StoreName
					vo.UserNickname = item.UserNickname
				}
			}
		}

		for _, item := range orderDataList {
			for _, vo := range out.Items {
				if item.OrderId == vo.OrderId {
					vo.OrderReturnStatus = item.OrderReturnStatus
					vo.OrderRefundStatus = item.OrderRefundStatus
				}
			}
		}

		//ConsumeTrade
		consumeTrades, err := dao.ConsumeTrade.Find(ctx, &do.ConsumeTradeListInput{Where: do.ConsumeTrade{OrderId: ids}})
		if err != nil {
			return nil, err
		}

		for _, item := range consumeTrades {
			for _, vo := range out.Items {
				if item.OrderId == vo.OrderId {
					vo.TradePaymentAmount = item.TradePaymentAmount
				}
			}
		}

		for _, vo := range out.Items {
			vo.ReturnFlag = array.InArray(orderReturnIdList, vo.OrderId)

			//是否开了发票
			vo.InvoiceIsApply = array.InArray(orderInvoiceIdList, vo.OrderId)
		}

		// 是否拼团成功订单
		activityGroupbookingHistories, err := dao.ActivityGroupbookingHistory.Find(ctx, &do.ActivityGroupbookingHistoryListInput{
			Where: do.ActivityGroupbookingHistory{OrderId: ids},
		})

		if err != nil {
			// 处理错误
		}

		if len(activityGroupbookingHistories) > 0 {
			userGroupbookingMap := make(map[string]*entity.ActivityGroupbookingHistory)
			for _, history := range activityGroupbookingHistories {
				userGroupbookingMap[history.OrderId] = history
			}

			// 设置 ActivityGroupbookingHistory 和 OrderIsGroupbookingSuccess
			for _, vo := range out.Items {
				if history, ok := userGroupbookingMap[vo.OrderId]; ok {
					vo.ActivityGroupbookingHistory = history
					vo.OrderIsGroupbookingSuccess = history.GbEnable == consts.ACTIVITY_GROUPBOOKING_SUCCESS
				}
			}
		}

		for _, vo := range out.Items {
			gconv.Scan(orderItemMap[vo.OrderId], &vo.Items)
		}
	}

	return out, err
}

// Add 新增
func (s *sOrder) Add(ctx context.Context, in *model.CheckoutInput) (out *model.OrderAddOutput, err error) {
	checkout, _ := service.UserCart().Checkout(ctx, in)
	checkout.In = in

	//添加订单
	out, err = s.addOrder(ctx, checkout)

	return
}

// Add 新增
func (s *sOrder) addOrder(ctx context.Context, cartData *model.CheckoutOutput) (out *model.OrderAddOutput, err error) {
	now := gtime.Now()

	userId := cartData.UserId
	buyerUserNickname := cartData.In.UserNickname
	gbId := cartData.In.GbId
	//activityId := cartData.In.ActivityId

	orderSelMoneyAmount := decimal.NewFromInt(0)
	orderSelPointsAmount := decimal.NewFromInt(0)
	orderSelSpAmount := decimal.NewFromInt(0)

	orderNeedPayPointsAmount := decimal.NewFromFloat(cartData.OrderPointsAmount)
	orderResourceExt1 := decimal.NewFromInt(0)
	orderResourceExt1Use := decimal.NewFromInt(0)

	userResource, _ := service.UserResource().Get(ctx, cartData.UserId)
	var orderIdRow []string

	usePoint := false
	if !g.IsEmpty(orderNeedPayPointsAmount) {
		orderResourceExt1 = orderNeedPayPointsAmount

		if userResource != nil {
			userPoints := decimal.NewFromFloat(userResource.UserPoints)

			if userPoints.Cmp(orderNeedPayPointsAmount) >= 0 {
				orderResourceExt1Use = orderResourceExt1
			} else {
				pointsVaueRate := decimal.NewFromFloat(service.ConfigBase().GetFloat(ctx, "points_vaue_rate", 0))

				if pointsVaueRate.Cmp(decimal.NewFromFloat(0)) <= 0 {
					orderResourceExt1Use = decimal.NewFromInt(0)
					panic(errors.New("积分价值配置有误，无法下单！"))
				} else {
					orderResourceExt1Use = userPoints
				}
			}
		}

		if orderResourceExt1Use.Cmp(decimal.NewFromInt(0)) >= 0 {
			desc := fmt.Sprintf("%s 积分兑换", orderResourceExt1Use)

			pts, _ := orderResourceExt1Use.Neg().Float64()

			pointsVo := &model.UserPointsVo{
				UserId:        cartData.UserId,
				Points:        pts,
				PointsTypeId:  consts.POINTS_TYPE_EXCHANGE_PRODUCT,
				PointsLogDesc: desc,
			}

			_, err = service.UserResource().Points(ctx, pointsVo)

			if err != nil {
				panic(errors.New("积分操作失败！"))
			}

			usePoint = true
		}
	}

	orderNeedSpAmount := decimal.NewFromFloat(cartData.OrderSpAmount)
	orderResourceExt2 := decimal.NewFromInt(0)
	orderResourceExt2Use := decimal.NewFromInt(0)

	if !g.IsEmpty(orderNeedSpAmount) {
		orderResourceExt2 = orderNeedSpAmount

		userSpTotal := decimal.NewFromFloat(userResource.UserSp)
		if userSpTotal.Cmp(orderNeedSpAmount) >= 0 {
			orderResourceExt2Use = orderResourceExt2
		} else {
			spVaueRate := decimal.NewFromFloat(service.ConfigBase().GetFloat(ctx, "sp_vaue_rate", 0))

			if spVaueRate.Cmp(decimal.NewFromFloat(0)) <= 0 {
				orderResourceExt2Use = decimal.NewFromInt(0)
				panic(errors.New("积分2不足，无法下单！"))
			} else {
				orderResourceExt2Use = userSpTotal
			}

			// 扣除众宝
			if !orderResourceExt2Use.IsZero() {
				//desc := fmt.Sprintf("%s 众宝兑换", orderResourceExt2Use);
				// todo User_ResourceModel::sp
			}
		}
	}

	var cartIds []uint64
	orderIdRow = make([]string, 0)

	checkoutRow := cartData.In
	categoryRateRow := make(map[uint]float64)

	chainId := checkoutRow.ChainId

	// 判断库存是否可以下单
	for _, storeItem := range cartData.Items {
		for _, itemTmpRow := range storeItem.Items {
			if itemTmpRow.IsOos {
				panic(errors.New(fmt.Sprintf("商品: %s 不在可售区域", itemTmpRow.ProductName)))
			}

			// 0元主商品，不可以下单
			if itemTmpRow.ItemSalePrice <= 0 && itemTmpRow.ItemUnitPoints <= 0 {
				panic(errors.New(fmt.Sprintf("商品: %s 总价为0，不可以下单，请联系商家！", itemTmpRow.ProductName)))
			}

			cartQuantity := itemTmpRow.CartQuantity

			// 直接判断库存
			itemQuantity := itemTmpRow.ItemQuantity
			kindId := itemTmpRow.KindId

			if itemQuantity < cartQuantity && kindId != consts.PRODUCT_KIND_EDU {
				panic(errors.New(fmt.Sprintf("商品: %s 库存不足！当前可购买：%d", itemTmpRow.ProductName, itemTmpRow.ItemQuantity)))
			}

			// 判断下单数量必须大于0
			if cartQuantity <= 0 {
				panic(errors.New(fmt.Sprintf("商品: %s 购买数量必须大于： %d", itemTmpRow.ProductName, 0)))
			}
		}
	}

	//订单编号每次新取
	orderId, err := service.NumberSeq().GetNextSeqString(ctx, fmt.Sprintf("JD-%s-", now.Format("Ymd")))
	if err != nil {
		return nil, err
	}

	//开启事务
	err = dao.OrderBase.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, storeItemVo := range cartData.Items {
			// 判断库存是否可以下单
			itemsList := storeItemVo.Items

			// 类型判断
			kindIds := array.Column(itemsList, "KindId")

			if len(kindIds) > 1 {
				panic(errors.New("可以购买同一种类商品"))
			}

			kinds := []uint{consts.PRODUCT_KIND_FUWU, consts.PRODUCT_KIND_CARD, consts.PRODUCT_KIND_EDU}

			if array.InArray(kinds, kindIds[0]) {
				if len(itemsList) > 1 {
					panic(errors.New("服务类商品每单只可以购买一种商品"))
				}
			}

			var orderData entity.OrderData
			var itemRows []*do.OrderItem
			var orderInfo *do.OrderInfo

			if !g.IsEmpty(orderId) {
				address := cartData.UserDeliveryAddress

				if address != nil {
					delivery := &do.OrderDeliveryAddress{
						OrderId:      orderId,
						DaName:       address.UdName,
						DaIntl:       address.UdIntl,
						DaMobile:     address.UdMobile,
						DaTelephone:  address.UdTelephone,
						DaProvinceId: address.UdProvinceId,
						DaProvince:   address.UdProvince,
						DaCityId:     address.UdCityId,
						DaCity:       address.UdCity,
						DaCountyId:   address.UdCountyId,
						DaCounty:     address.UdCounty,
						DaAddress:    address.UdAddress,
						DaPostalcode: address.UdPostalcode,
						DaTagName:    address.UdTagName,
						DaLatitude:   address.UdLatitude,
						DaLongitude:  address.UdLongitude,
						DaTime:       now,
					}

					_, err = service.OrderDeliveryAddress().Save(ctx, delivery)
					if err != nil {
						panic(errors.New("保存订单配送地址数据失败!"))
					}
				}
			} else {
				panic(errors.New("生成订单编号异常!"))
			}

			storeId := storeItemVo.StoreId
			storeName := storeItemVo.StoreName

			// 优惠券使用标记更新
			if !g.IsEmpty(storeItemVo.UserVoucherId) {
				userVoucher := &do.UserVoucher{
					UserVoucherId:         storeItemVo.UserVoucherId,
					VoucherStateId:        consts.VOUCHER_STATE_USED,
					OrderId:               orderId,
					UserVoucherActivetime: now,
				}

				_, err = service.UserVoucher().Edit(ctx, userVoucher)
				if err != nil {
					panic(errors.New("订单优惠券信息失败"))
				}
			}
			// end 优惠券使用标记更新

			// 使用掉的积分额度
			var orderResourceExt1UseCurrent, orderResourceExt1NeedMoney decimal.Decimal

			// 需要ext1，但是ext1不足， 将ext1变为money使用
			if orderResourceExt1.Cmp(decimal.Zero) > 0 {
				productPointsSel := decimal.NewFromFloat(storeItemVo.PointsAmount)

				if orderResourceExt1.Cmp(orderResourceExt1Use) > 0 {
					orderResourceExt1UseCurrent = orderResourceExt1Use.Mul(productPointsSel).Div(orderResourceExt1)

					// 将积分变成钱去支付
					orderResourceExt1Need := productPointsSel.Sub(orderResourceExt1UseCurrent)

					// 将积分换成钱
					pointsVaueRate := decimal.NewFromFloat(service.ConfigBase().GetFloat(ctx, "points_vaue_rate", 0))
					orderResourceExt1NeedMoney = pointsVaueRate.Mul(orderResourceExt1Need)
					orderMoneySelectItems := decimal.NewFromFloat(storeItemVo.MoneyAmount)

					storeItemVo.MoneyAmount, _ = orderMoneySelectItems.Add(decimal.Max(decimal.Zero, orderResourceExt1NeedMoney)).Float64()
				} else {
					orderResourceExt1UseCurrent = productPointsSel
				}
			}

			storeItemVo.PointsAmount, _ = orderResourceExt1UseCurrent.Float64()

			// 积分不足，使用钱支付的部分均分到对应商品上。
			//orderMoneySelectItems := storeItemVo.MoneyItemAmount
			//freight := storeItemVo.FreightAmount

			voucherAmount := decimal.NewFromFloat(storeItemVo.VoucherAmount)
			orderPaymentAmount := decimal.NewFromFloat(storeItemVo.MoneyAmount)
			orderPaymentAmount = decimal.Max(orderPaymentAmount, decimal.Zero)

			// begain 均分积分
			//order_resource_ext1_use_current;

			var itemSharePoints, nowTmpPointsTotal, tmpItemPointsTotal decimal.Decimal

			if orderResourceExt1UseCurrent.Cmp(decimal.NewFromInt(0)) > 0 {
				//var pointsItemIds []int64

				//涉及商品个数
				var size int

				//取出参与的产品的总值
				for idx := 0; idx < len(itemsList); idx++ {
					item := itemsList[idx]

					if item.ItemUnitPoints > 0 {
						tmpItemPointsTotal = tmpItemPointsTotal.Add(decimal.NewFromFloat(item.ItemPointsSubtotal))

						size++
					}
				}

				var i int

				for idx := 0; idx < len(itemsList); idx++ {
					item := itemsList[idx]

					if item.ItemUnitPoints > 0 {
						i++

						//最后一个商品
						if i == size {
							itemSharePoints = orderResourceExt1UseCurrent.Sub(nowTmpPointsTotal)
						} else {
							itemSharePoints = decimal.NewFromFloat(item.ItemSubtotal).Div(tmpItemPointsTotal).Mul(orderResourceExt1UseCurrent).Round(2)
							nowTmpPointsTotal = nowTmpPointsTotal.Add(itemSharePoints)
						}
					}

					item.ItemPointsSubtotal, _ = itemSharePoints.Float64()
				}

				usePoint = true
			} else {
				usePoint = false
			}

			// begain 均分积分
			// begain 均分优惠券
			var itemShareVoucher, nowTmpVoucherTotal, tmpItemTotal decimal.Decimal

			if voucherAmount.Cmp(decimal.Zero) > 0 {
				var voucherItemIds []uint64
				userVoucher, _ := service.UserVoucher().Get(ctx, storeItemVo.UserVoucherId)

				if userVoucher != nil && !g.IsEmpty(userVoucher.ItemId) {
					voucherItemIds = gconv.SliceUint64(gstr.Split(userVoucher.ItemId, ","))
				}

				//涉及商品个数
				var size int

				//取出参与的产品的总值
				for idx := 0; idx < len(itemsList); idx++ {
					item := itemsList[idx]

					if (len(voucherItemIds) > 0 && array.InArray(voucherItemIds, item.ItemId)) || len(voucherItemIds) == 0 {
						subtotalTmp := item.ItemSubtotal
						tmpItemTotal.Add(decimal.NewFromFloat(subtotalTmp))

						size++
					}
				}

				var i int

				for idx := 0; idx < len(itemsList); idx++ {
					item := itemsList[idx]

					if (len(voucherItemIds) > 0 && array.InArray(voucherItemIds, item.ItemId)) || len(voucherItemIds) == 0 {
						i++

						//最后一个商品
						if i == size {
							itemShareVoucher = voucherAmount.Sub(nowTmpVoucherTotal)
						} else {
							itemShareVoucher = decimal.NewFromFloat(item.ItemSubtotal).Div(tmpItemTotal).Mul(voucherAmount).Round(2)
							nowTmpVoucherTotal = nowTmpVoucherTotal.Add(itemShareVoucher)
						}

						item.ItemVoucher, _ = itemShareVoucher.Float64()
					}

				}
			}
			// end 优惠券均分

			//todo 处理店铺活动优惠

			// 1、订单基础表
			orderBase := &do.OrderBase{}

			// 订单默认状态
			stateIdList, _ := service.ConfigBase().GetStateIdList(ctx)
			orderStateId := stateIdList[0]

			orderBase.OrderId = orderId
			orderBase.OrderStateId = orderStateId
			orderBase.OrderProductAmount = storeItemVo.ProductAmount
			orderBase.OrderPaymentAmount = orderPaymentAmount

			// 修改最终下单数据
			orderSelMoneyAmount = orderSelMoneyAmount.Add(orderPaymentAmount)
			orderSelPointsAmount = orderSelPointsAmount.Add(orderResourceExt1UseCurrent)
			orderSelSpAmount = orderSelSpAmount.Add(orderResourceExt2)

			// 应付金额/应支付金额:order_goods_amount - order_discount_amount + order_shipping_fee - order_voucher_price - order_points_fee - order_adjust_fee
			// 手工调整默认为0，order_points_fee积分折扣暂未开启

			orderBase.OrderTime = now
			orderBase.StoreId = storeId
			orderBase.StoreName = storeName
			orderBase.UserId = userId
			orderBase.UserNickname = buyerUserNickname

			// 订单基础信息保存
			_, err = service.OrderBase().Add(ctx, orderBase)

			deliveryTypeId := checkoutRow.DeliveryTypeId

			// 2、订单信息保存处理
			if err == nil {
				orderTitle := gstr.JoinAny(array.Column(itemsList, "ProductName"), "|")
				orderTitle = orderTitle[:int32(math.Min(190, float64(len(orderTitle)-1)))]

				subsiteId := uint(0)
				storeIsSelfsupport := true
				paymentTypeId := checkoutRow.PaymentTypeId
				orderIsOffline := false
				var salespersonId uint = 0
				var distributorUserId uint = 0
				storeType := uint(1)
				orderType := checkoutRow.OrderType
				srcOrderId := checkoutRow.SrcOrderId
				paymentFormId := uint(1)
				cartTypeId := uint(1)

				orderInfo = &do.OrderInfo{
					OrderId:                     orderId,
					OrderTitle:                  orderTitle,
					StoreId:                     storeId,
					SubsiteId:                   subsiteId,
					UserId:                      userId,
					KindId:                      storeItemVo.KindId,
					OrderLockStatus:             false,
					OrderIsSettlemented:         false,
					OrderBuyerEvaluationStatus:  0,
					OrderSellerEvaluationStatus: 0,
					CreateTime:                  now.UnixNano() / int64(time.Millisecond),
					OrderBuyerHidden:            false,
					OrderShopHidden:             false,
					PaymentTypeId:               paymentTypeId,
					OrderStateId:                orderStateId,
					UpdateTime:                  now.UnixNano() / int64(time.Millisecond),
					OrderIsReceived:             false,
					OrderFinanceReview:          false,
					OrderIsPaid:                 consts.ORDER_PAID_STATE_NO,
					OrderIsOut:                  consts.ORDER_PICKING_STATE_NO,
					OrderIsShipped:              consts.ORDER_SHIPPED_STATE_NO,
					ChainId:                     chainId,
					DeliveryTypeId:              deliveryTypeId,
					OrderIsOffline:              orderIsOffline,
					CartTypeId:                  cartTypeId,
					SalespersonId:               salespersonId,
					DistributorUserId:           distributorUserId,
					StoreIsSelfsupport:          storeIsSelfsupport,
					StoreType:                   storeType,
					SrcOrderId:                  srcOrderId,
					OrderType:                   orderType,
					PaymentFormId:               paymentFormId,
				}

				// 店铺活动 - 非排他
				if storeItemVo.ActivityBase != nil {
					orderInfo.ActivityId = storeItemVo.ActivityBase.ActivityId
					orderInfo.ActivityTypeId = storeItemVo.ActivityBase.ActivityTypeId
				} else {
					var activityIds []uint
					var activityTypeIds []uint
					for _, s := range itemsList {
						if !g.IsEmpty(s.ActivityId) {
							activityIds = append(activityIds, s.ActivityId)
						}

						if s.ActivityInfo != nil {
							activityTypeIds = append(activityTypeIds, s.ActivityInfo.ActivityTypeId)
						}
					}

					orderInfo.ActivityId = gstr.JoinAny(activityIds, ",")
					orderInfo.ActivityTypeId = gstr.JoinAny(activityTypeIds, ",")
				}

				// 订单基本info信息保存
				if _, err = service.OrderInfo().Add(ctx, orderInfo); err != nil {
					panic("保存订单基础数据失败!")
				}
			} else {
				panic(err)
				panic(errors.New("保存订单基础数据失败！"))
			}
			// end 订单信息保存处理

			if err == nil {
				// 服务类虚拟订单数据
				if storeItemVo.KindId != consts.PRODUCT_KIND_ENTITY {
					if storeItemVo.KindId == consts.PRODUCT_KIND_EDU {
						// todo 具体业务逻辑，在支付完成处理。
					} else {
						if storeItemVo.KindId == consts.PRODUCT_KIND_CARD {
							// 卡券类，发送code， 则代表交易完成，类似充话费、虚拟卡号等等
							// todo 具体业务逻辑，在支付完成处理。
						}

						// 发送虚拟码，并记录客户递交的数据。
						virtualRow := &do.ChainCode{}
						itemId := itemsList[0].ItemId

						//virtualServiceTime := ConvertToDate(getParameter("virtual_service_time"))
						//virtualServiceDate := ConvertToDate(getParameter("virtual_service_date"))

						virtualRow.ChainCode = ""
						virtualRow.OrderId = orderId
						virtualRow.ItemId = itemId
						virtualRow.ChainCodeStatus = 0
						//virtualRow.VirtualServiceDate = virtualServiceDate
						//virtualRow.VirtualServiceTime = virtualServiceTime
						virtualRow.ChainId = chainId
						virtualRow.UserId = userId
						virtualRow.StoreId = storeId

						if _, err := service.ChainCode().Add(ctx, virtualRow); err != nil {
							panic("保存服务订单数据失败!")
						}
					}
				}

				// 自提码数据
				if storeItemVo.KindId == consts.PRODUCT_KIND_ENTITY && deliveryTypeId == consts.DELIVERY_TYPE_SELF_PICK_UP {
					// 发送虚拟码，并记录客户递交的数据。
					chainCodeRow := &do.ChainCode{}

					//chainCodeStatus := ConvertToInt(getParameter("chain_code_status"))
					//  virtualServiceTime := ConvertToDate(getParameter("virtual_service_time"))
					//  virtualServiceDate := ConvertToDate(getParameter("virtual_service_date"))
					//

					chainCodeRow.ChainCode = ""
					chainCodeRow.OrderId = orderId
					//chainCodeRow.ItemId = itemId
					chainCodeRow.ChainCodeStatus = 0
					//chainCodeRow.VirtualServiceDate = virtualServiceDate
					//chainCodeRow.VirtualServiceTime = virtualServiceTime
					chainCodeRow.ChainId = chainId
					chainCodeRow.UserId = userId
					chainCodeRow.StoreId = storeId

					if _, err := service.ChainCode().Add(ctx, chainCodeRow); err != nil {
						panic("保存自提码数据失败!")
					}
				}

				// 店铺商品信息item_rows
				for _, item := range itemsList {
					cartId := item.CartId

					if !g.IsEmpty(cartId) {
						cartIds = append(cartIds, cartId)
					}

					isOnSale := item.IsOnSale

					if isOnSale {
						itemRow := &do.OrderItem{}

						itemRow.OrderId = orderId
						itemRow.UserId = userId            // 买家user_id  冗余
						itemRow.StoreId = storeId          // 店铺Id
						itemRow.ProductId = item.ProductId // 产品id
						itemRow.ProductName = item.ProductName
						itemRow.ItemId = item.ItemId         // 货品id
						itemRow.ItemSrcId = item.ItemSrcId   // 货品id
						itemRow.ItemName = item.ItemName     // 商品名称
						itemRow.OrderItemFile = ""           // 订单附件
						itemRow.CategoryId = item.CategoryId // 商品对应的类目Id

						itemRow.ItemUnitPrice = item.ItemUnitPrice      // 商品价格单价
						itemRow.ItemUnitPoints = item.ItemUnitPoints    // 商品价格单价
						itemRow.ItemUnitSp = 0                          // 商品价格单价
						itemRow.OrderItemSalePrice = item.ItemSalePrice // 商品实际成交价单价

						itemRow.ItemCostPrice = item.ItemCostPrice                 // 成本价
						itemRow.OrderItemQuantity = item.CartQuantity              // 商品数量
						itemRow.OrderItemInventoryLock = item.ProductInventoryLock // 锁库存时机
						itemRow.OrderItemImage = item.ProductImage                 // 商品图片
						itemRow.OrderItemReturnNum = 0                             // 退货数量
						itemRow.OrderItemAmount = item.ItemSubtotal                // 商品实际总金额 =  item_sale_price * goods_quantity
						itemRow.OrderItemDiscountAmount = item.ItemDiscountAmount  // 优惠金额  负数
						itemRow.PolicyDiscountrate = item.ItemPolicyDiscountrate   // 价格策略折扣率%
						itemRow.OrderItemVoucher = item.ItemVoucher                // 均分优惠券
						//itemRow.OrderItemRedemptionVoucher = itemShareRedemption // 均分提货券
						itemRow.OrderItemConfirmStatus = true // 默认用户审核

						//itemRow.ItemPurchasePrice = itemPurchasePrice
						//itemRow.ItemPurchaseRate = itemPurchaseRate
						//itemRow.ItemSalesRate = itemSalesRate

						// todo 积分费用
						itemRow.OrderItemAdjustFee = 0 // 手工调整金额 负数

						if usePoint {
							itemRow.OrderItemPointsFee = item.ItemPointsSubtotal
						} else {
							itemRow.OrderItemPointsFee = 0
						}

						itemRow.OrderItemPaymentAmount = item.ItemSubtotal // // 实付金额 : goods_payment_amount =  goods_amount + goods_discount_amount + goods_adjust_fee + goods_point_fee

						itemRow.OrderItemEvaluationStatus = false // 评价状态: 0-未评价;1-已评价

						// 活动订单处理
						if !g.IsEmpty(item.ActivityId) {
							itemRow.ActivityTypeId = item.ActivityInfo.ActivityTypeId // 活动类型:0-默认;1101-加价购=搭配宝;1102-店铺满赠-小礼品;1103-限时折扣;1104-优惠套装;1105-店铺优惠券coupon优惠券;1106-拼团;1107-满减送;1108-阶梯价
							itemRow.ActivityId = item.ActivityId                      // 促销活动Id:与activity_type_id搭配使用, 团购Id/限时折扣Id/优惠套装Id
						}

						itemRow.OrderItemDiscountAmount = item.ItemDiscountAmount // 优惠金额  负数

						// 根据分类获取category_commission_rate
						if item.KindId != consts.PRODUCT_KIND_EDU {
							if _, exists := categoryRateRow[item.CategoryId]; exists {
							} else {
								categoryRow, _ := service.ProductCategory().Get(ctx, item.CategoryId)
								categoryRateRow[item.CategoryId] = categoryRow.CategoryCommissionRate
							}
						}

						// 允许分销
						productDistEnable := item.ProductDistEnable

						if productDistEnable {
							// 存在单品平台佣金行为
							productCommissionRate := item.ProductCommissionRate
							orderItemPaymentAmount := decimal.NewFromFloat(itemRow.OrderItemPaymentAmount.(float64))

							if productCommissionRate > 0 {
								itemRow.OrderItemCommissionRate = productCommissionRate                                                                               // 分佣金比例
								itemRow.OrderItemCommissionFee = orderItemPaymentAmount.Mul(decimal.NewFromFloat(productCommissionRate)).Div(decimal.NewFromInt(100)) // 分佣金
							} else {
								orderItemCommissionRate := categoryRateRow[item.CategoryId]
								itemRow.OrderItemCommissionRate = orderItemCommissionRate
								itemRow.OrderItemCommissionFee = orderItemPaymentAmount.Mul(decimal.NewFromFloat(orderItemCommissionRate)).Div(decimal.NewFromInt(100)) // 分佣金
							}
						} else {
							itemRow.OrderItemCommissionRate = 0
							itemRow.OrderItemCommissionFee = decimal.Zero // 分佣金
						}

						// start 判断增加冻结库存
						if item.ProductInventoryLock == 1001 {
							if num, _ := service.ProductItem().LockSkuStock(ctx, item.ItemId, item.CartQuantity); num == 0 {
								panic(fmt.Sprintf("更改: %d 冻结库存失败!", item.ItemId))
							}
						}
						// end

						itemRows = append(itemRows, itemRow)
					} else {
						panic(fmt.Sprintf("商品：%s 已经下架，不可下单", item.ItemName))
					}
				}
				// 检查活动价格是否设置正确
				for _, itemRow := range itemRows {
					if itemRow.OrderItemPaymentAmount.(float64) < 0 {
						panic("活动价格设置有误")
					}
				}

				// 保存订单信息数据
				if _, err = service.OrderItem().Saves(ctx, itemRows); err != nil {
					panic("保存订单信息数据失败!")
				}

				orderData := &do.OrderData{}

				orderData.OrderId = orderId               // 订单编号
				orderData.OrderDesc = ""                  // 订单描述
				orderData.OrderDelayTime = 0              // 延迟时间,默认为0 - 收货确认
				orderData.DeliveryTypeId = deliveryTypeId // 配送方式

				// 买家订单留言
				userMessage := ""
				if checkoutRow.Message != nil && len(checkoutRow.Message[storeId]) > 0 {
					userMessage = checkoutRow.Message[storeId]
				}
				orderData.OrderMessage = userMessage

				orderData.OrderItemAmount = storeItemVo.MoneyItemAmount      // 商品总价格/商品金额, 不包含运费
				orderData.OrderAdjustFee = 0                                 // 手工调整费用店铺优惠
				orderData.OrderPointsFee = storeItemVo.PointsAmount          // 积分费用
				orderData.OrderDiscountAmount = storeItemVo.DiscountAmount   // 折扣价格/优惠总金额
				orderData.OrderShippingFeeAmount = storeItemVo.FreightAmount // 运费价格/运费金额
				orderData.OrderShippingFee = storeItemVo.FreightAmount       // 实际运费金额-卖家可修改
				orderData.VoucherId = storeItemVo.UserVoucherId              // 优惠券id/优惠券/返现:发放选择使用
				orderData.VoucherPrice = storeItemVo.VoucherAmount           // 优惠券面额
				orderData.OrderResourceExt1 = orderResourceExt1UseCurrent

				// 平台交易佣金
				var orderCommissionFee decimal.Decimal
				for _, itemRow := range itemRows {
					orderCommissionFee = orderCommissionFee.Add(itemRow.OrderItemCommissionFee.(decimal.Decimal))
				}

				orderData.OrderCommissionFee = orderCommissionFee

				// 保存订单数据
				if _, err = service.OrderData().Add(ctx, orderData); err != nil {
					panic("保存订单数据失败!")
				}

				if !g.IsEmpty(checkoutRow.UserInvoiceId) {
					userInvoice, err := service.UserInvoice().Get(ctx, checkoutRow.UserInvoiceId)

					if userInvoice != nil && err != nil {
						orderInvoice := &do.OrderInvoice{}

						orderInvoice.UserId = orderBase.UserId
						orderInvoice.StoreId = orderBase.StoreId
						orderInvoice.OrderId = orderId

						orderInvoice.InvoiceTitle = userInvoice.InvoiceTitle
						orderInvoice.InvoiceCompanyCode = userInvoice.InvoiceCompanyCode
						orderInvoice.InvoiceIsCompany = userInvoice.InvoiceIsCompany
						orderInvoice.InvoiceAddress = userInvoice.InvoiceAddress
						orderInvoice.InvoicePhone = userInvoice.InvoicePhone
						orderInvoice.InvoiceBankname = userInvoice.InvoiceBankname
						orderInvoice.InvoiceBankaccount = userInvoice.InvoiceBankaccount
						orderInvoice.InvoiceContent = orderInfo.OrderTitle

						orderInvoice.InvoiceType = userInvoice.InvoiceType
						orderInvoice.InvoiceAmount = orderBase.OrderPaymentAmount
						orderInvoice.OrderIsPaid = false
						orderInvoice.InvoiceTime = time.Now()

						orderInvoice.InvoiceContactName = userInvoice.InvoiceContactName

						if _, err := service.OrderInvoice().Add(ctx, orderInvoice); err != nil {
							panic("保存订单发票信息数据失败!")
						}
					}
				}
			}

			orderIdRow = append(orderIdRow, orderId)

			// 获取店铺主账号
			// shopBaseRow := shopStoreBaseService.Get(orderBase.StoreId)
			// sellerId := shopBaseRow.UserId
			// invoicingCustomerBaseService.DoStoreAddCustomer(orderBase.StoreId, orderBase.BuyerUserId)

			// 暂时注释掉 Java 中的逻辑，Go 中未提供具体的函数调用和对象定义

			// 或者通过API
			subsiteId := 0

			//买家是否开店
			buyerStoreId := 0

			//获取超管用户编号
			userAdminQueryWrapper := &do.UserAdminListInput{}
			userAdminQueryWrapper.Where.UserIsSuperadmin = true
			userAdmin, err := service.UserAdmin().FindOne(ctx, userAdminQueryWrapper)

			if err != nil {
				return err
			}

			consumeTradeRow := &do.ConsumeTrade{
				OrderId:                  orderId,
				BuyerId:                  userId,
				BuyerStoreId:             buyerStoreId,
				SellerId:                 userAdmin.UserId,
				StoreId:                  storeId,
				SubsiteId:                subsiteId,
				ChainId:                  chainId,
				TradeIsPaid:              consts.ORDER_PAID_STATE_NO,
				TradeTypeId:              consts.TRADE_TYPE_SHOPPING,
				PaymentChannelId:         0,
				TradeModeId:              1,
				OrderPaymentAmount:       orderBase.OrderPaymentAmount,
				OrderCommissionFee:       orderData.OrderCommissionFee,
				TradePaymentAmount:       orderBase.OrderPaymentAmount,
				TradePaymentMoney:        0,
				TradePaymentRechargeCard: 0,
				TradePaymentPoints:       0,
				TradePaymentCredit:       0,
				TradePaymentRedpack:      0,
				TradeDiscount:            orderData.OrderDiscountAmount,
				TradeAmount:              orderData.OrderItemAmount,
				TradeTitle:               orderInfo.OrderTitle,
				TradeCreateTime:          time.Now().UnixNano() / int64(time.Millisecond),
			}

			if _, err := service.ConsumeTrade().Add(ctx, consumeTradeRow); err != nil {
				panic("订单支付信息失败!")
			} else {
				orderPaymentAmount := orderBase.OrderPaymentAmount

				if orderPaymentAmount.(decimal.Decimal).Cmp(decimal.Zero) <= 0 {
					// 订单付款状态处理，
					// 不需要添加收款记录，直接修改订单状态
					if _, err := s.SetPaidYes(ctx, orderId); err != nil {
						panic("订单支付状态修改失败!")
					} else {
						cartData.IsPaid = true
					}
				}

				// 分销用户来源 - 平台推广员功能，佣金平台出
				ifPlantformFx := service.ConfigBase().IfPlantformFx(ctx)

				if ifPlantformFx {
					distributionOrderVo := &model.DistributionOrderVo{
						UserId:             userId,
						StoreId:            storeId,
						OrderId:            orderId,
						OrderCommissionFee: orderData.OrderCommissionFee,
						SalespersonId:      orderInfo.SalespersonId.(uint),
						DistributorUserId:  orderInfo.DistributorUserId.(uint),
					}

					productIdRow := make([]uint64, len(itemRows))
					for i, item := range itemRows {
						productIdRow[i] = item.ProductId.(uint64)
					}

					userLevelIdFlag := false

					// 暂时注释掉 Java 中的逻辑，Go 中未提供具体的函数调用和对象定义

					if !userLevelIdFlag {
						// 分销，初始化订单信息
						service.DistributionOrder().InitDistributionUserOrder(ctx, distributionOrderVo, itemRows)
					}
				}

				if len(cartIds) > 0 {
					if _, err := service.UserCart().Remove(ctx, cartIds); err != nil {
						panic("删除购物车失败")
					}
				}
				// 添加订单事件
			}
		}

		return nil
	})

	cartData.OrderMoneyAmount, _ = orderSelMoneyAmount.Float64()
	cartData.OrderPointsAmount, _ = orderSelPointsAmount.Float64()
	cartData.OrderSpAmount, _ = orderSelSpAmount.Float64()

	orderAddOutput := &model.OrderAddOutput{
		CheckoutOutput: *cartData,
		OrderIds:       orderIdRow,
		GbId:           gbId,
		// 其他属性赋值...
	}

	return orderAddOutput, err
}

// ifCancel 是否可以取消 支付后也可以取消
func (s *sOrder) ifCancel(orderStateId uint, orderIsPaid uint) bool {
	orderStates := []uint{consts.ORDER_STATE_WAIT_PAY, consts.ORDER_STATE_WAIT_REVIEW, consts.ORDER_STATE_WAIT_FINANCE_REVIEW, consts.ORDER_STATE_PICKING, consts.ORDER_STATE_WAIT_SHIPPING}
	//return array.InArray(orderStates, orderStateId) && orderIsPaid == consts.ORDER_PAID_STATE_NO
	return array.InArray(orderStates, orderStateId)
}

// ifReturn 是否可以退货
func (s *sOrder) ifReturn(orderStateId uint, orderIsPaid uint) bool {
	orderStates := []uint{consts.ORDER_STATE_SHIPPED, consts.ORDER_STATE_RECEIVED, consts.ORDER_STATE_FINISH}
	return array.InArray(orderStates, orderStateId) && orderIsPaid == consts.ORDER_PAID_STATE_YES
}

/**
 * 取消订单
 *
 * @param orderId       订单编号
 * @param orderStateNote 订单状态备注
 * @return bool 是否取消成功
 */
func (s *sOrder) Cancel(ctx context.Context, orderId string, orderStateNote string) (flag bool, err error) {
	flag = false

	orderBase, err := dao.OrderBase.Get(ctx, orderId)

	if err != nil || g.IsEmpty(orderBase) {
		return false, err
	}

	orderInfo, err := dao.OrderInfo.Get(ctx, orderId)

	// 拼团支付，不可取消
	//if (checkPaidFlag && !ifCancel(orderInfo.getOrderStateId(), orderInfo.getOrderIsPaid())) {
	if !s.ifCancel(orderInfo.OrderStateId, orderInfo.OrderIsPaid) {
		panic("无符合取消条件的订单")
	}

	//开启事务
	err = dao.OrderBase.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		if orderBase.OrderStateId != consts.ORDER_STATE_CANCEL {
			flag, err = s.EditNextState(ctx, orderId, orderBase.OrderStateId, consts.ORDER_STATE_CANCEL, orderStateNote)

			if err != nil {
				return err
			}
		} else {
			return errors.New("未更改到符合条件的订单！")
		}

		orderItemRows, err := dao.OrderItem.Find(ctx, &do.OrderItemListInput{
			Where: do.OrderItem{
				OrderId: orderId,
			},
		})

		// 判断是否出库，出库会释放冻结库存，取消出库后的订单，不释放冻结库存
		orderStates := []uint{consts.ORDER_STATE_WAIT_PAY, consts.ORDER_STATE_WAIT_REVIEW, consts.ORDER_STATE_WAIT_FINANCE_REVIEW, consts.ORDER_STATE_PICKING}
		if array.InArray(orderStates, orderInfo.OrderStateId) {
			// 部分出库商品数量
			billItems, _ := dao.StockBillItem.Find(ctx, &do.StockBillItemListInput{
				Where: do.StockBillItem{
					OrderId: orderId,
				},
			})

			billItemQuantityAll := make(map[uint64]uint)

			for _, billItem := range billItems {
				if !g.IsEmpty(billItem.OrderItemId) {
					quantity := billItemQuantityAll[billItem.ItemId]
					quantity += billItem.BillItemQuantity
					billItemQuantityAll[billItem.ItemId] = quantity
				}
			}

			// 释放冻结库存
			for _, orderItem := range orderItemRows {
				// start 释放冻结库存
				orderItemInventoryLock := orderItem.OrderItemInventoryLock
				if orderItemInventoryLock == 1001 || (orderInfo.OrderIsPaid == consts.ORDER_PAID_STATE_YES && orderItemInventoryLock == 1002) {
					releaseQuantity := orderItem.OrderItemQuantity
					quantity := billItemQuantityAll[orderItem.ItemId]

					// 去掉部分出库商品数量
					if quantity != 0 {
						releaseQuantity -= quantity

						// 出库未发货，可以注释掉后， 商品数量需要手工入库。
						input := &model.ProductEditStockInput{
							ItemId:       orderItem.ItemId,
							ItemQuantity: quantity,
							BillTypeId:   consts.BILL_TYPE_IN,
						}

						_, err = service.ProductItem().EditStock(ctx, input)
						if err != nil {
							panic("编辑库存失败！")
						}
					}

					if releaseQuantity > 0 {
						num, _ := service.ProductItem().ReleaseSkuStock(ctx, orderItem.ItemId, releaseQuantity)

						if num == 0 {
							panic(fmt.Sprintf("释放: %d 冻结库存失败!", orderItem.ItemId))
						}
					}
				}
			}
		}

		if orderInfo.OrderIsShipped != consts.ORDER_SHIPPED_STATE_NO {
			// 部分发货
			panic("订单部分发货，不可取消，请联系商家！")
		}

		// todo 判断是否需要退款
		if orderInfo.OrderIsPaid != consts.ORDER_PAID_STATE_NO {
			orderReturnInput := &model.OrderReturnInput{
				ReturnAllFlag:      true,
				OrderId:            orderId,
				ReturnBuyerMessage: orderStateNote,
				UserId:             orderBase.UserId,
				ReturnFlag:         consts.ORDER_NOT_NEED_RETURN_GOODS,
				ReviewFlag:         true,
			}

			for _, orderItem := range orderItemRows {
				returnItemInputVo := &model.OrderReturnItemInputVo{
					OrderItemId:        orderItem.OrderItemId,
					ReturnRefundAmount: orderItem.OrderItemCanRefundAmount - orderItem.OrderItemReturnSubtotal,
					ReturnItemNum:      orderItem.OrderItemQuantity - orderItem.OrderItemReturnNum,
				}

				orderReturnInput.ReturnItems = append(orderReturnInput.ReturnItems, returnItemInputVo)
			}

			_, err = service.OrderReturn().AddItem(ctx, orderReturnInput)
			if err != nil {
				panic("退款失败！")
			}
		}

		// todo 积分退还
		// 积分退还 order_resource_ext1 默认为积分。
		orderData, err := service.OrderData().Get(ctx, orderId)
		if (err != nil && orderData != nil) && orderData.OrderResourceExt1 > 0 && orderBase.UserId != 0 {
			desc := fmt.Sprintf("%s 积分退还，订单号 %s", orderData.OrderResourceExt1, orderData.OrderId)

			userPointsVo := &model.UserPointsVo{
				UserId:        orderBase.UserId,
				Points:        orderData.OrderResourceExt1,
				PointsTypeId:  consts.POINTS_TYPE_CONSUME_RETRUN,
				PointsLogDesc: desc,
				OrderId:       orderData.OrderId,
			}

			if _, err := service.UserResource().Points(ctx, userPointsVo); err != nil {
				panic("积分操作失败！")
			}
		}

		// 订单取消优惠券退还
		_, err = service.UserVoucher().EditWhere(ctx, &do.UserVoucherListInput{
			Where: do.UserVoucher{
				OrderId: orderId,
			},
		}, &do.UserVoucher{OrderId: "", VoucherStateId: consts.VOUCHER_STATE_UNUSED})

		if err != nil {
			panic("取消优惠券失败！")
		}

		//取消拼团, 未支付的取消活动数据， 如果支付，则必须通过计划任务取消
		_, err = s.CancelActivity(ctx, orderId)
		if err != nil {
			panic("取消拼团失败！")
		}

		// 取消发票
		_, err = service.OrderInvoice().RemoveWhere(ctx, &do.OrderInvoiceListInput{
			Where: do.OrderInvoice{
				OrderId:       orderId,
				InvoiceStatus: 0,
			},
		})

		if err != nil {
			panic("取消发票失败！")
		}

		return nil
	})

	return flag, err
}

/**
 * 支付完成
 *
 * @param orderId 订单编号
 * @return 是否支付成功
 */
func (s *sOrder) SetPaidYes(ctx context.Context, orderId string) (flag bool, err error) {
	flag = false
	orderInfo, err := service.OrderInfo().Get(ctx, orderId)

	if err != nil {
		return false, err
	}

	if orderInfo == nil {
		panic(fmt.Sprintf("订单信息 %s 不存在!", orderId))
	}

	orderBase, err := service.OrderBase().Get(ctx, orderId)

	if err != nil {
		return false, err
	}

	if orderBase == nil {
		panic(fmt.Sprintf("订单基础 %s 不存在!", orderId))
	}

	if orderInfo.OrderStateId == consts.ORDER_STATE_WAIT_PAY && orderInfo.OrderIsPaid != consts.ORDER_PAID_STATE_YES {
		//库存是否足够
		orderItemList, err := service.OrderItem().Find(ctx, &do.OrderItemListInput{Where: do.OrderItem{OrderId: orderId}})

		if err != nil {
			return false, err
		}

		for _, item := range orderItemList {
			// start 判断增加冻结库存
			if item.OrderItemInventoryLock == 1002 {
				if affected, err := service.ProductItem().LockSkuStock(ctx, item.ItemId, item.OrderItemQuantity); affected == 0 || err == nil {
					format := fmt.Sprintf("更改: %d 冻结库存失败!", item.ItemId)
					service.LogError().Error(ctx, format, consts.ERR_NOT_DEFINITION)

					//不报错，允许执行，日志记录
					//throw new BusinessException(String.format(__("更改: %d 冻结库存失败!"), item.getItemId()))

					//库存不足，走自动退款流程
					s.Cancel(ctx, orderId, fmt.Sprintf("%d 库存不足，取消订单!", item.ItemId))

					return false, err
				}
			}
			// end
		}

		//获取订单的下一条状态
		nextOrderStateId, err := s.getNextOrderStateId(ctx, orderInfo.OrderStateId)

		flag, err = s.EditNextState(ctx, orderId, orderInfo.OrderStateId, nextOrderStateId, "")

		if err != nil {
			return false, err
		}

		//更新支付状态
		_, err = dao.OrderInfo.Edit(ctx, do.OrderInfo{OrderId: orderId}, &do.OrderInfo{OrderIsPaid: consts.ORDER_PAID_STATE_YES})

		//更新发票订单支付状态
		queryWrapper := &do.OrderInvoiceListInput{
			Where: do.OrderInvoice{OrderId: orderId},
		}

		orderInvoice := &do.OrderInvoice{OrderIsPaid: true}
		dao.OrderInvoice.EditWhere(ctx, queryWrapper, orderInvoice)

		// 活动判断

		// 推广

		// 读取订单商品，更新销量
		for _, orderItem := range orderItemList {
			productId := orderItem.ProductId
			orderItemQuantity := orderItem.OrderItemQuantity
			dao.ProductIndex.Increment(ctx, productId, dao.ProductIndex.Columns().ProductSaleNum, orderItemQuantity)
		}

		//// 付款成功，对用户进行提醒
		//messageId := "payment-success-reminding"
		//args := map[string]interface{}{
		//	"order_id":             orderId,
		//	"product_name":         orderInfo.OrderTitle,
		//	"order_payment_amount": orderBase.OrderPaymentAmount,
		//	"order_add_time":       time.Now().Format("2006-01-02 15:04:05"),
		//}
		//messageService.SendNoticeMsg(orderBase.UserId, messageId, args)
		//
		//// 提醒商家发货
		//messageId = "notice-of-delivery"
		//args = map[string]interface{}{
		//	"order_id": orderId,
		//	"date":     time.Now().Format("2006-01-02 15:04:05"),
		//}
		//
		//adminUserId := service.ConfigBase().GetInt(ctx, "message_notice_user_id", 10001)
		//
		//messageService.SendNoticeMsg(adminUserId, messageId, args)

	} else {
		panic("未更改到符合条件的订单！")
	}

	// Todo: 处理支付成功逻辑

	return flag, err
}

// Review 审核订单
func (s *sOrder) Review(ctx context.Context, orderId string, orderStateNote string) (flag bool, err error) {
	//判断活动前置条件
	s.IfActivity(ctx, orderId)

	orderBase, err := dao.OrderBase.Get(ctx, orderId)

	if err != nil {
		return false, err
	}

	if orderBase.OrderStateId == consts.ORDER_STATE_WAIT_REVIEW {
		//获取订单的下一条状态
		nextOrderStateId, err := s.getNextOrderStateId(ctx, orderBase.OrderStateId)

		flag, err = s.EditNextState(ctx, orderId, orderBase.OrderStateId, nextOrderStateId, orderStateNote)

		if err != nil {
			return false, err
		}
	} else {
		return false, errors.New("未更改到符合条件的订单！")
	}

	return
}

// Finance 财务审核
func (s *sOrder) Finance(ctx context.Context, orderId string, orderStateNote string) (flag bool, err error) {
	//判断活动前置条件
	s.IfActivity(ctx, orderId)

	orderBase, err := dao.OrderBase.Get(ctx, orderId)

	if err != nil {
		return false, err
	}

	if orderBase.OrderStateId == consts.ORDER_STATE_WAIT_FINANCE_REVIEW {
		//获取订单的下一条状态
		nextOrderStateId, err := s.getNextOrderStateId(ctx, orderBase.OrderStateId)

		flag, err = s.EditNextState(ctx, orderId, orderBase.OrderStateId, nextOrderStateId, orderStateNote)

		if err != nil {
			return false, err
		}
	} else {
		return false, errors.New("未更改到符合条件的订单！")
	}

	return
}

// Picking 出库审核
func (s *sOrder) Picking(ctx context.Context, in *model.OrderPickingInput) (flag bool, err error) {
	//判断活动前置条件
	s.IfActivity(ctx, in.OrderId)

	//判断前置条件
	_, err = s.CheckOrderReturnWaiting(ctx, in.OrderId)

	if err != nil {
		return false, err
	}

	if len(in.Items) > 0 {
		in.PickingFlag = false
	}

	orderBase, err := dao.OrderBase.Get(ctx, in.OrderId)

	if err != nil {
		return false, err
	}

	if orderBase.OrderStateId == consts.ORDER_STATE_PICKING || orderBase.OrderStateId == consts.ORDER_STATE_WAIT_SHIPPING {

		//开启事务
		err = dao.OrderBase.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

			state, err := s.DoReviewPicking(ctx, in)

			if err != nil {
				return err
			}

			if state == consts.ORDER_PICKING_STATE_YES && orderBase.OrderStateId == consts.ORDER_STATE_PICKING {
				//获取订单的下一条状态
				nextOrderStateId, err := s.getNextOrderStateId(ctx, orderBase.OrderStateId)

				flag, err = s.EditNextState(ctx, in.OrderId, orderBase.OrderStateId, nextOrderStateId, "")

				if err != nil {
					return err
				}
			}

			return nil
		})

	} else {
		return false, errors.New("未更改到符合条件的订单！")
	}

	return
}

// CheckOrderReturnWaiting 判断是否有待审核售后订单条件限制
func (s *sOrder) CheckOrderReturnWaiting(ctx context.Context, orderId string) (flag bool, err error) {
	out, err := dao.OrderReturn.FindKey(ctx, &do.OrderReturnListInput{Where: do.OrderReturn{OrderId: orderId, ReturnStateId: consts.RETURN_PROCESS_CHECK}})

	if err != nil {
		return false, err
	}

	if len(out) > 0 {
		return false, errors.New(fmt.Sprintf("有待处理的退款或者退货单: %s，请先处理！", gstr.JoinAny(out, ",")))
	}

	return true, err
}

// ifShipping 是否可以发货
func (s *sOrder) ifShipping(ctx context.Context, orderStateId uint) (flag bool) {
	return orderStateId == consts.ORDER_STATE_WAIT_SHIPPING || orderStateId == consts.ORDER_STATE_PICKING
}

// Shipping 发货
func (s *sOrder) Shipping(ctx context.Context, in *model.OrderShippingInput) (flag bool, err error) {
	orderId := in.OrderId

	s.IfActivity(ctx, orderId)

	//判断前置条件
	_, err = s.CheckOrderReturnWaiting(ctx, in.OrderId)

	if err != nil {
		return false, err
	}

	orderBase, err := dao.OrderBase.Get(ctx, orderId)

	if err != nil {
		return false, err
	}

	if s.ifShipping(ctx, orderBase.OrderStateId) {
		state, err := s.DoReviewShipping(ctx, in)

		if err != nil {
			return false, err
		}

		if state == consts.ORDER_SHIPPED_STATE_YES {
			//获取订单的下一条状态
			nextOrderStateId, err := s.getNextOrderStateId(ctx, orderBase.OrderStateId)

			//当前状态可能为待配货，下一个状态为代发货，则直接更改为已发货。
			if nextOrderStateId == consts.ORDER_STATE_WAIT_SHIPPING {
				nextOrderStateId = consts.ORDER_STATE_SHIPPED
			}

			flag, err = s.EditNextState(ctx, orderId, orderBase.OrderStateId, nextOrderStateId, "")

			if err != nil {
				return false, err
			}
		}
	} else {
		return false, errors.New("未更改到符合条件的订单！")
	}

	return
}

// CheckShippingComplete 检测是否发货完成
func (s *sOrder) CheckShippingComplete(ctx context.Context, orderId string) (isComplete bool, err error) {
	orderInfo, err := dao.OrderInfo.Get(ctx, orderId)
	if err != nil {
		return false, err
	}

	if g.IsEmpty(orderInfo) {
		return false, errors.New(fmt.Sprintf("订单 %s 不存在！", orderId))
	}

	//检测是否发货完成
	isComplete = true

	//出库单无对应发货信息的，完成发货操作
	//物流记录
	orderLogistics, err := dao.OrderLogistics.Find(ctx, &do.OrderLogisticsListInput{Where: do.OrderLogistics{OrderId: orderId}})
	if err != nil {
		return false, err
	}

	//StockBill
	stockBills, err := dao.StockBill.Find(ctx, &do.StockBillListInput{Where: do.StockBill{OrderId: orderId}})

	if err != nil {
		return false, err
	}

	ids := array.Column(orderLogistics, dao.OrderLogistics.Columns().StockBillId)

	for _, bill := range stockBills {
		if !array.InArray(ids, bill.StockBillId) {
			//完成发货信息
			isComplete = false
			break
		}
	}

	if isComplete {
		//判断商品是否全部出库

		//订单商品
		orderItems, err := dao.OrderItem.Find(ctx, &do.OrderItemListInput{Where: do.OrderItem{OrderId: orderId}})
		if err != nil {
			return false, err
		}

		//已出库商品
		billItems, err := dao.StockBillItem.Find(ctx, &do.StockBillItemListInput{Where: do.StockBillItem{OrderId: orderId}})
		if err != nil {
			return false, err
		}

		//差量商品
		billItemQuantityAll := make(map[uint64]*model.PickingItem)

		for _, orderItem := range orderItems {
			//todo 扣除同意退货数量
			billItemQuantityAll[orderItem.OrderItemId] = &model.PickingItem{OrderItemId: orderItem.OrderItemId, ItemId: orderItem.ItemId, BillItemQuantity: orderItem.OrderItemQuantity, BillItemPrice: orderItem.OrderItemSalePrice, ProductId: orderItem.ProductId}
		}

		for _, billItem := range billItems {
			if !g.IsEmpty(billItem.OrderItemId) {
				if value, ok := billItemQuantityAll[billItem.OrderItemId]; ok {
					value.BillItemQuantity -= billItem.BillItemQuantity
				} else {
					return false, errors.New(fmt.Sprintf("出库数据有误 '%s'", billItem.OrderItemId))
				}
			}
		}

		if !g.IsEmpty(billItemQuantityAll) {
			for _, pickingItem := range billItemQuantityAll {
				if pickingItem.BillItemQuantity > 0 {
					isComplete = false
					break
				}
			}

			//商品已经全部发货了
			if isComplete {
				//所有发货状态商品不是发货完成状态，
				if orderInfo.OrderIsShipped != consts.ORDER_SHIPPED_STATE_YES {
					state := consts.ORDER_SHIPPED_STATE_YES
					_, err = dao.OrderInfo.Edit(ctx, orderId, &do.OrderInfo{OrderIsShipped: state})

					//获取订单的下一条状态
					nextOrderStateId, err := service.ConfigBase().GetNextOrderStateId(ctx, orderInfo.OrderStateId)
					if err != nil {
						return false, err
					}

					_, err = service.Order().EditNextState(ctx, orderId, orderInfo.OrderStateId, nextOrderStateId, "")

					if err != nil {
						return false, err
					}
				}
			}
		}
	}

	return isComplete, err
}

// AddLogistics 添加订单日志
func (s *sOrder) AddLogistics(ctx context.Context, in *do.OrderLogistics) (flag bool, err error) {
	_, err = s.SaveLogistics(ctx, in)
	s.CheckShippingComplete(ctx, in.OrderId.(string))
	return
}

// SaveLogistics 添加订单日志
func (s *sOrder) SaveLogistics(ctx context.Context, in *do.OrderLogistics) (flag bool, err error) {
	logisticsId := in.LogisticsId
	ssId := in.SsId

	storeExpressLogistics, _ := dao.StoreExpressLogistics.Get(ctx, logisticsId)
	storeShippingAddress, _ := dao.StoreShippingAddress.Get(ctx, ssId)

	in.ExpressName = storeExpressLogistics.ExpressName
	in.ExpressId = storeExpressLogistics.ExpressId

	in.LogisticsPhone = storeShippingAddress.SsIntl + storeShippingAddress.SsMobile
	in.LogisticsMobile = storeShippingAddress.SsIntl + storeShippingAddress.SsMobile
	in.LogisticsContacter = storeShippingAddress.SsContacter
	in.LogisticsAddress = storeShippingAddress.SsAddress
	in.LogisticsPostcode = storeShippingAddress.SsPostalcode

	_, err = dao.OrderLogistics.Save(ctx, in)

	return flag, err
}

// Receive 收货
func (s *sOrder) Receive(ctx context.Context, orderId string, orderStateNote string) (flag bool, err error) {
	orderBase, err := dao.OrderBase.Get(ctx, orderId)

	if err != nil {
		return false, err
	}

	if orderBase.OrderStateId == consts.ORDER_STATE_SHIPPED {
		//获取订单的下一条状态
		nextOrderStateId, err := s.getNextOrderStateId(ctx, orderBase.OrderStateId)

		flag, err = s.EditNextState(ctx, orderId, orderBase.OrderStateId, nextOrderStateId, orderStateNote)

		if err != nil {
			return false, err
		}

		if flag {
			// 分销功能。目前付款成功发放佣金，此处为收货后发放
			fxSettleType := service.ConfigBase().GetStr(ctx, "fx_settle_type", "receive")

			if fxSettleType == "receive" {
				service.DistributionOrder().SettleDistributionUserOrder(ctx, orderId)
			}
		}

		return flag, err

	} else {
		return false, errors.New("未更改到符合条件的订单！")
	}

	return
}

// getNextOrderStateId 读取启用配置，根据当前orderStateId获得下一状态 sc_order_process
func (s *sOrder) getNextOrderStateId(ctx context.Context, orderStateId uint) (nextOrderStateId uint, err error) {
	return service.ConfigBase().GetNextOrderStateId(ctx, orderStateId)
}

// getAllNextOrderStateId 读取所有配置，根据当前orderStateId获得下一状态 sc_order_process
func (s *sOrder) getAllNextOrderStateId(ctx context.Context, orderStateId uint) (nextOrderStateId uint, err error) {
	return service.ConfigBase().GetAllNextOrderStateId(ctx, orderStateId)
}

// addOrderStateLog 添加订单日志
func (s *sOrder) addOrderStateLog(ctx context.Context, data *do.OrderStateLog) (flag bool, err error) {
	_, err = dao.OrderStateLog.Add(ctx, data)

	return
}

// EditNextState 修改订单为下一个待处理状态
func (s *sOrder) EditNextState(ctx context.Context, orderId string, orderStateId uint, nextOrderStateId uint, orderStateNote string) (flag bool, err error) {
	//下一个状态存在
	if consts.ORDER_STATE_CANCEL != nextOrderStateId {
		ss := global.StateIdRow
		if !array.InArray(ss, nextOrderStateId) {
			return false, errors.New("订单下个状态不符合配置要求！")
		}
	}

	//必须更新到记录
	num, err := dao.OrderBase.Edit(ctx, do.OrderBase{OrderId: orderId, OrderStateId: orderStateId}, &do.OrderBase{OrderStateId: nextOrderStateId})
	if err != nil {
		return false, err
	}

	if num <= 0 {
		return false, errors.New("未更改到符合条件的订单！")
	}

	//订单信息更改
	oldInfo := do.OrderInfo{OrderId: orderId, OrderStateId: orderStateId}
	newInfo := &do.OrderInfo{OrderStateId: nextOrderStateId}

	if consts.ORDER_STATE_CANCEL != nextOrderStateId {
		switch orderStateId {
		case consts.ORDER_STATE_WAIT_PAY:
			//newInfo.OrderIsPaid = true //放入支付回调更改
		case consts.ORDER_STATE_WAIT_REVIEW:
			oldInfo.OrderIsReview = false
			newInfo.OrderIsReview = true
		case consts.ORDER_STATE_WAIT_FINANCE_REVIEW:
			oldInfo.OrderFinanceReview = false
			newInfo.OrderFinanceReview = true
		case consts.ORDER_STATE_PICKING:
			//oldInfo.OrderIsOut = []int{consts.ORDER_PICKING_STATE_NO, consts.ORDER_PICKING_STATE_PART}
			newInfo.OrderIsOut = consts.ORDER_PICKING_STATE_YES
		case consts.ORDER_STATE_WAIT_SHIPPING:
			//发货完成状态已经修改，
			//oldInfo.OrderIsShipped = []int{consts.ORDER_SHIPPED_STATE_NO, consts.ORDER_SHIPPED_STATE_PART}
			newInfo.OrderIsShipped = consts.ORDER_SHIPPED_STATE_YES
		case consts.ORDER_STATE_SHIPPED:
			//oldInfo.OrderIsReceived = false
			newInfo.OrderIsReceived = true
			newInfo.OrderReceivedTime = gtime.Now()
		default:

		}
	}

	//必须更新到记录
	num, err = dao.OrderInfo.Edit(ctx, oldInfo, newInfo)
	if err != nil {
		return false, err
	}

	if num <= 0 {
		return false, errors.New("未更改到符合条件的订单！")
	}
	user := service.BizCtx().GetUser(ctx)

	var userId uint
	var userAccount string

	if user != nil {
		userId = service.BizCtx().GetUser(ctx).UserId
		userAccount = service.BizCtx().GetUser(ctx).UserAccount
	}

	//添加日志
	_, err = s.addOrderStateLog(ctx, &do.OrderStateLog{
		OrderId:         orderId,
		OrderStateId:    nextOrderStateId,
		OrderStatePreId: orderStateId,
		UserId:          userId,
		UserAccount:     userAccount,
		OrderStateNote:  orderStateNote,
		OrderStateTime:  gtime.Now(),
	})

	return
}

// DoReviewPicking 出库审核 - 逻辑封装 - 涉及进销存
func (s *sOrder) DoReviewPicking(ctx context.Context, in *model.OrderPickingInput) (state uint, err error) {
	// 清理数据
	if len(in.Items) > 0 {
		for i := 0; i < len(in.Items); i++ {
			pickingItem := in.Items[i]

			if pickingItem.BillItemQuantity <= 0 {
				in.Items = append(in.Items[:i], in.Items[i+1:]...)
				i-- // 减少索引以避免跳过下一个元素
			}
		}
	}

	//订单商品
	orderItems, err := dao.OrderItem.Find(ctx, &do.OrderItemListInput{Where: do.OrderItem{OrderId: in.OrderId}})

	//已出库商品
	billItems, err := dao.StockBillItem.Find(ctx, &do.StockBillItemListInput{Where: do.StockBillItem{OrderId: in.OrderId}})

	//差量商品
	billItemQuantityAll := make(map[uint64]*model.PickingItem)
	billItemWaiting := make(map[uint64]*model.PickingItem)

	for _, orderItem := range orderItems {
		//todo 扣除同意退货数量
		billItemQuantityAll[orderItem.OrderItemId] = &model.PickingItem{OrderItemId: orderItem.OrderItemId, ItemId: orderItem.ItemId, BillItemQuantity: orderItem.OrderItemQuantity, BillItemPrice: orderItem.OrderItemSalePrice, ProductId: orderItem.ProductId}
	}

	for _, billItem := range billItems {
		if !g.IsEmpty(billItem.OrderItemId) {
			if value, ok := billItemQuantityAll[billItem.OrderItemId]; ok {
				value.BillItemQuantity -= billItem.BillItemQuantity
			} else {
				return state, errors.New(fmt.Sprintf("出库数据有误 '%s'", billItem.OrderItemId))
			}
		}
	}

	//全部出库
	if in.PickingFlag {
		billItemWaiting = billItemQuantityAll
	} else {
		//todo 指定的出库商品及数量，需要判断是否符合billItemQuantityAll中的要求。
		for _, item := range in.Items {
			if item.BillItemQuantity > 0 {
				billItemWaiting[item.OrderItemId] = &model.PickingItem{OrderItemId: item.OrderItemId, ItemId: item.ItemId, BillItemQuantity: item.BillItemQuantity, BillItemPrice: item.BillItemPrice, ProductId: item.ProductId}
			}
		}
	}

	if g.IsEmpty(billItemWaiting) {
		return state, errors.New(fmt.Sprintf("无待出库出库数据 '%s'", in.OrderId))
	}

	//商品库存

	//出库单
	now := gtime.Now()
	user := service.BizCtx().GetUser(ctx)

	stockBillId, err := service.NumberSeq().GetNextSeqString(ctx, fmt.Sprintf("OUT-%s-", now.Format("Ymd")))
	if err != nil {
		return 0, err
	}

	stockBill := do.StockBill{
		StockBillId:          stockBillId,
		StockBillChecked:     1,
		StockBillDate:        now,
		StockBillModifyTime:  now,
		StockBillTime:        now.UnixMilli(),
		BillTypeId:           in.BillTypeId,
		StockTransportTypeId: in.StockTransportTypeId,
		StoreId:              0,
		WarehouseId:          0,
		OrderId:              in.OrderId,
		StockBillRemark:      "",
		EmployeeId:           user.UserId,
		AdminId:              user.UserId,
		StockBillOtherMoney:  0,
		StockBillAmount:      0,  // 订单金额
		StockBillEnable:      1,  // 是否有效(BOOL):1-有效; 0-无效
		StockBillSrcId:       "", // 关联编号
	}

	//单据金额
	stockBillAmount := 0.0

	for orderItemId, pickingItem := range billItemWaiting {
		//单据商品小计
		billItemSubtotal := pickingItem.BillItemPrice * gconv.Float64(pickingItem.BillItemQuantity)

		stockBillItem := do.StockBillItem{
			StockBillId:       stockBillId,
			OrderId:           in.OrderId,
			OrderItemId:       orderItemId,
			ItemId:            pickingItem.ItemId,
			BillItemQuantity:  pickingItem.BillItemQuantity,
			BillItemUnitPrice: pickingItem.BillItemPrice,
			BillItemSubtotal:  billItemSubtotal,
			ProductId:         pickingItem.ProductId,

			//ProductName           interface{} // 商品名称
			//ItemName              interface{} // 商品名称
			//UnitId                interface{} // 单位编号
			//WarehouseItemQuantity interface{} // 库存量
			//StoreId               interface{} // 店铺编号
			WarehouseId:          stockBill.WarehouseId,
			StockTransportTypeId: stockBill.StockTransportTypeId,
			//BillItemRemark        interface{} // 备注
			BillTypeId: stockBill.BillTypeId,
		}

		// 获取订单商品信息
		orderItem, err := dao.OrderItem.Get(ctx, stockBillItem.OrderItemId)
		if err == nil {
			stockBillItem.ProductName = orderItem.ProductName
			stockBillItem.ItemName = orderItem.ItemName
		}

		_, err = dao.StockBillItem.Add(ctx, &stockBillItem)

		if err != nil {
			return state, err
		}

		billItemQuantityAll[orderItemId].BillItemQuantity -= pickingItem.BillItemQuantity

		stockBillAmount += billItemSubtotal
	}

	stockBill.StockBillAmount = stockBillAmount
	_, err = dao.StockBill.Add(ctx, &stockBill)

	if err != nil {
		return state, err
	}

	//判断是否已经全部出库， 需要修改订单状态

	state = consts.ORDER_PICKING_STATE_YES

	for _, pickingItem := range billItemQuantityAll {
		if pickingItem.BillItemQuantity > 0 {
			state = consts.ORDER_PICKING_STATE_PART
			break
		}
	}

	_, err = dao.OrderInfo.Edit(ctx, in.OrderId, &do.OrderInfo{OrderIsOut: state})

	if err != nil {
		return state, err
	}

	return
}

// DoReviewShipping 发货审核  - 涉及快递单号处理
func (s *sOrder) DoReviewShipping(ctx context.Context, in *model.OrderShippingInput) (state uint, err error) {
	//如果为出库状态
	orderInfo, err := dao.OrderInfo.Get(ctx, in.OrderId)

	if err != nil {
		return state, err
	}

	if orderInfo.OrderIsOut != consts.ORDER_PICKING_STATE_YES {
		_, err = s.DoReviewPicking(ctx, &model.OrderPickingInput{OrderId: in.OrderId, PickingFlag: true})
	}

	//发货
	//出库单无对应发货信息的，完成发货操作
	//物流记录
	orderLogistics, err := dao.OrderLogistics.Find(ctx, &do.OrderLogisticsListInput{Where: do.OrderLogistics{OrderId: in.OrderId}})
	if err != nil {
		return 0, err
	}

	//StockBill
	stockBills, err := dao.StockBill.Find(ctx, &do.StockBillListInput{Where: do.StockBill{OrderId: in.OrderId}})

	if err != nil {
		return 0, err
	}

	ids := array.Column(orderLogistics, dao.OrderLogistics.Columns().StockBillId)

	input := do.OrderLogistics{}
	gconv.Scan(in, &input)

	for _, bill := range stockBills {
		if !array.InArray(ids, bill.StockBillId) {
			//完成发货信息
			input.StockBillId = bill.StockBillId
			_, err = s.SaveLogistics(ctx, &input)

			if err != nil {
				return 0, err
			}

			//
			//// 获取快递信息
			//expressLogistics, err := dao.StoreExpressLogistics.Get(ctx, in.LogisticsId)
			//if err != nil {
			//	// 处理错误情况
			//}
			//
			//if orderLogistics == nil {
			//	// 发货通知
			//	messageID := "order_complete_shipping"
			//	args := map[string]interface{}{
			//		"order_id":              in.OrderId,
			//		"logistics_name":        expressLogistics.ExpressName,
			//		"order_tracking_number": in.OrderTrackingNumber, // 这里的 orderLogistics 似乎应该是 expressLogistics
			//	}
			//	err := messageService.SendNoticeMsg(ctx, orderInfo.UserId, messageID, args)
			//	if err != nil {
			//		// 处理发送通知消息失败的情况
			//	}
			//}

		}
	}

	state = consts.ORDER_SHIPPED_STATE_YES
	_, err = dao.OrderInfo.Edit(ctx, in.OrderId, &do.OrderInfo{OrderIsShipped: state})

	return
}

// ReviewToState 审核订单到某个状态
func (s *sOrder) ReviewToState(ctx context.Context, orderId string, toOrderStateId uint) (flag bool, err error) {
	//判断前置条件
	_, err = s.CheckOrderReturnWaiting(ctx, orderId)

	if err != nil {
		return false, err
	}

	//开启事务
	err = dao.OrderBase.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		if consts.ORDER_STATE_CANCEL == toOrderStateId {
			flag, err = s.Cancel(ctx, orderId, "")

			if err != nil {
				return err
			}
		} else {
			tryCount := 0
			for tryCount < 10 {
				//读取订单
				orderBase, err := dao.OrderBase.Get(ctx, orderId)

				if err != nil {
					return err
				}

				//订单已经为目标状态
				if orderBase.OrderStateId == toOrderStateId {
					flag = true

					return nil
					//break
				}

				//获取订单的下一条状态
				nextOrderStateId, err := s.getNextOrderStateId(ctx, orderBase.OrderStateId)

				if nextOrderStateId == consts.ORDER_STATE_WAIT_SHIPPING {
					_, err = s.Picking(ctx, &model.OrderPickingInput{OrderId: orderId, PickingFlag: true})
				} else if nextOrderStateId == consts.ORDER_STATE_SHIPPED {
					flag, err = s.Shipping(ctx, &model.OrderShippingInput{OrderId: orderId})
				} else {
					flag, err = s.EditNextState(ctx, orderId, orderBase.OrderStateId, nextOrderStateId, "")
				}

				if err != nil {
					return err
				}

				tryCount++
			}
		}

		return nil
	})

	return
}

func (s *sOrder) CancelActivity(ctx context.Context, orderId string) (flag bool, err error) {
	// 如果此订单是是拼团生成的订单，取消后要转让团长或改变状态为拼团失败
	// 查找订单关联的团
	activityGroupQueryWrapper := &do.ActivityGroupbookingHistoryListInput{
		Where: do.ActivityGroupbookingHistory{OrderId: orderId, GbEnable: g.Slice{consts.ACTIVITY_GROUPBOOKING_INEFFECTIVE, consts.ACTIVITY_GROUPBOOKING_UNDERWAY}},
	}

	bookingHistory, err := dao.ActivityGroupbookingHistory.FindOne(ctx, activityGroupQueryWrapper)
	if err != nil {
		return false, err
	}

	if bookingHistory != nil {
		gbId := bookingHistory.GbId
		gbhId := bookingHistory.GbhId

		// 查找团的所有关联的订单
		historyQueryWrapper := &do.ActivityGroupbookingHistoryListInput{
			Where: do.ActivityGroupbookingHistory{GbId: gbId},
		}

		histories, err := dao.ActivityGroupbookingHistory.Find(ctx, historyQueryWrapper)
		if err != nil {
			return false, err
		}

		// 如果团目前只关联了一个订单，订单取消同时团也作废，如果已经关联一个以上订单，取消订单同时将团转让给其他参团者
		if len(histories) == 1 && histories[0].OrderId == orderId {
			// 团只有自己（有唯一历史记录，且订单号是自己的订单号），取消订单则拼团失败
			activityGroupbooking := &do.ActivityGroupbooking{GbId: gbId, GbEnable: consts.ACTIVITY_GROUPBOOKING_FAIL}
			// 取消团
			_, err = dao.ActivityGroupbooking.Edit(ctx, gbId, activityGroupbooking)
			if err != nil {
				return false, err
			}
		}

		// 取消参团记录
		activityGroupbookingHistory := &do.ActivityGroupbookingHistory{GbhId: gbhId, GbEnable: consts.ACTIVITY_GROUPBOOKING_FAIL}
		_, err = dao.ActivityGroupbookingHistory.Edit(ctx, gbhId, activityGroupbookingHistory)
		if err != nil {
			return false, err
		}
	}

	return true, err
}

// 判断是否有活动条件限制
func (s *sOrder) IfActivity(ctx context.Context, orderId string) bool {
	orderInfo, _ := dao.OrderInfo.Get(ctx, orderId)

	if len(orderInfo.ActivityTypeId) > 0 {

	}

	return true
}

// GetOrderStatisticsInfo 根据用户id获取用户订单统计信息
func (s *sOrder) GetOrderStatisticsInfo(ctx context.Context, userId uint) (*model.OrderNumOutput, error) {
	orderNumOutput := &model.OrderNumOutput{}

	wg := sync.WaitGroup{}
	wg.Add(3)

	// 创建一个协程池，设置最大并发数为10
	//pool := grpool.New(10)

	// 提交一个任务到协程池中
	grpool.Add(ctx, func(ctx context.Context) {
		//已完成
		orderNumInput := &model.OrderNumInput{
			UserId:       userId,
			OrderStateId: consts.ORDER_STATE_FINISH,
			KindId:       consts.PRODUCT_KIND_ENTITY,
		}
		orderNumOutput.FinNumEntity = s.GetOrderNum(ctx, orderNumInput)

		orderNumInput.KindId = consts.PRODUCT_KIND_FUWU
		orderNumOutput.FinNumV = s.GetOrderNum(ctx, orderNumInput)

		// 取消订单数
		orderNumInput.OrderStateId = consts.ORDER_STATE_CANCEL
		orderNumInput.KindId = consts.PRODUCT_KIND_ENTITY
		orderNumOutput.CancelNumEntity = s.GetOrderNum(ctx, orderNumInput)

		orderNumInput.KindId = consts.PRODUCT_KIND_FUWU
		orderNumOutput.CancelNumV = s.GetOrderNum(ctx, orderNumInput)

		wg.Done()
	})

	grpool.Add(ctx, func(ctx context.Context) {
		// 任务函数
		orderNumInput := &model.OrderNumInput{
			UserId:       userId,
			OrderStateId: consts.ORDER_STATE_PICKING,
			KindId:       consts.PRODUCT_KIND_ENTITY,
		}
		orderPickingNum := s.GetOrderNum(ctx, orderNumInput)

		orderNumInput.OrderStateId = consts.ORDER_STATE_WAIT_SHIPPING
		orderShippingNum := s.GetOrderNum(ctx, orderNumInput)
		orderNumOutput.WaitShippingNumEntity = orderPickingNum + orderShippingNum

		orderNumInput.KindId = consts.PRODUCT_KIND_FUWU
		orderPickingNum = s.GetOrderNum(ctx, orderNumInput)

		orderNumInput.OrderStateId = consts.ORDER_STATE_WAIT_SHIPPING
		orderShippingNum = s.GetOrderNum(ctx, orderNumInput)
		orderNumOutput.WaitShippingNumV = orderPickingNum + orderShippingNum

		// 已发货货订单数
		orderNumInput.OrderStateId = consts.ORDER_STATE_SHIPPED
		orderNumInput.KindId = consts.PRODUCT_KIND_ENTITY
		orderNumOutput.ShipNumEntity = s.GetOrderNum(ctx, orderNumInput)

		orderNumInput.KindId = consts.PRODUCT_KIND_FUWU
		orderNumOutput.ShipNumV = s.GetOrderNum(ctx, orderNumInput)

		wg.Done()
	})

	grpool.Add(ctx, func(ctx context.Context) {
		// 任务函数

		orderNumInput := &model.OrderNumInput{
			UserId:       userId,
			OrderStateId: consts.ORDER_STATE_WAIT_PAY,
			KindId:       consts.PRODUCT_KIND_ENTITY,
		}
		orderNumOutput.WaitPayNumEntity = s.GetOrderNum(ctx, orderNumInput)

		orderNumInput.KindId = consts.PRODUCT_KIND_FUWU
		orderNumOutput.WaitPayNumV = s.GetOrderNum(ctx, orderNumInput)

		// 售后订单数
		returningNum, _ := dao.OrderReturn.Count(ctx, &do.OrderReturnListInput{
			Where: do.OrderReturn{
				BuyerUserId: userId,
				ReturnStateId: []uint{
					consts.RETURN_PROCESS_SUBMIT,
					consts.RETURN_PROCESS_CHECK,
					consts.RETURN_PROCESS_RECEIVED,
					consts.RETURN_PROCESS_REFUND,
				}},
		})

		orderNumOutput.ReturningNum = int64(returningNum)

		wg.Done()
	})

	// 等待所有任务完成
	wg.Wait()

	return orderNumOutput, nil
}

// GetOrderNum 订单数量
func (s *sOrder) GetOrderNum(ctx context.Context, in *model.OrderNumInput) int64 {
	input := &do.OrderInfoListInput{Where: do.OrderInfo{}}

	if !g.IsEmpty(in.OrderStateId) {
		input.Where.OrderStateId = in.OrderStateId
	}
	if !g.IsEmpty(in.UserId) {
		input.Where.UserId = in.UserId
	}

	if !g.IsEmpty(in.KindId) {
		input.Where.KindId = in.KindId
	}

	if !g.IsEmpty(in.OrderStime) {
		ext := &ml.WhereExt{Column: dao.OrderInfo.Columns().CreateTime, Val: in.OrderStime, Symbol: ml.GE}
		input.BaseList.WhereExt = append(input.BaseList.WhereExt, ext)
	}

	if !g.IsEmpty(in.OrderEtime) {
		ext := &ml.WhereExt{Column: dao.OrderInfo.Columns().CreateTime, Val: in.OrderEtime, Symbol: ml.LE}
		input.BaseList.WhereExt = append(input.BaseList.WhereExt, ext)
	}

	num, _ := dao.OrderInfo.Count(ctx, input)

	return int64(num)
}
