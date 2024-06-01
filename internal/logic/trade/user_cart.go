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
	"encoding/json"
	"errors"
	"fmt"
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
	"golershop.cn/utility/array"
	"sort"
)

type sUserCart struct{}

func init() {
	service.RegisterUserCart(NewUserCart())
}

func NewUserCart() *sUserCart {
	return &sUserCart{}
}

// List 分页读取
func (s *sUserCart) GetList(ctx context.Context, in *do.UserCartListInput) (out *model.CheckoutOutput, err error) {
	out = &model.CheckoutOutput{}

	list, err := dao.UserCart.Find(ctx, in)
	gconv.Scan(list, &out)

	checkoutInput := &model.CheckoutInput{}
	checkoutInput.UserId = gconv.Uint(in.Where.UserId)
	gconv.Scan(list, &checkoutInput.Items)

	out, err = s.Checkout(ctx, checkoutInput)

	return out, err
}

// Checkout
/**
 * 生成订单数据，结算checkout预览及生成订单, 理论上属于订单模块
 * <p>
 * 1、购物车中商品基本信息读取
 * 2、格式化数据（阶梯价计算等等），店铺数据分组，商品仓库分组
 * 3、活动数据、优惠折扣团购等等
 * 4、活动商品数据满赠满减、加价购、换购等等
 * 5、根据选择，计算订单信息，将上一步权限验证并将结果计入数据中, 店铺商品总价 = 加价购商品总价 + 购物车非活动商品总价（按照限时折扣和团购价计算）
 * 6、
 * 7、结算使用部分
 * 7.1、可用店铺优惠券、店铺礼品卡、店铺红包
 * 7.2、可用平台红包、平台优惠券、礼品卡
 * 7.3、最终折扣计算(最终支付价格打折)
 * 7.4、计算每样商品的佣金
 * 7.5、计算最终运费，运费根据重量计费，同一店铺一次发货，商品独立计算快递费用不合理。
 * 9、计算总价
 *
 * @param in
 * @return
 */
func (s *sUserCart) Checkout(ctx context.Context, in *model.CheckoutInput) (out *model.CheckoutOutput, err error) {
	return s.FormatCartRows(ctx, in)
}

/**
 * formatCartRows 生成订单数据，结算checkout预览及生成订单
 * 1、购物车中商品基本信息读取
 * 2、店铺数据分组，商品仓库分组
 * 3、活动数据、优惠折扣、团购、（阶梯价计算）等等
 * 4、活动商品数据满赠、兑换等等
 * 5、根据选择，计算订单信息，将上一步权限验证并将结果计入数据中, 店铺商品总价 = 加价购商品总价 + 购物车非活动商品总价（按照限时折扣和团购价计算）
 * 6、
 *
 * @param in 购物车数据
 */
func (s *sUserCart) FormatCartRows(ctx context.Context, in *model.CheckoutInput) (out *model.CheckoutOutput, err error) {
	out = &model.CheckoutOutput{}
	out.UserId = in.UserId
	userInfo, err := dao.UserInfo.Get(ctx, in.UserId)

	if userInfo == nil {
		return nil, errors.New("用户信息不存在！")
	}

	// 处理店铺活动
	var activityBase *entity.ActivityBase

	if !g.IsEmpty(in.ActivityId) {
		// 参加活动 产品及数量 - 活动信息使用
		itemNumVoMap := make(map[uint64]*model.ItemNumVo)

		activityBase, err = dao.ActivityBase.Get(ctx, in.ActivityId)

		if activityBase != nil {
			if activityBase.ActivityState != consts.ACTIVITY_STATE_NORMAL {
				return nil, errors.New("活动尚未开启！")
			}

			// 会员等级判断
			_, err := s.checkoutLevel(ctx, activityBase.ActivityUseLevel, userInfo.UserLevelId)
			if err != nil {
				panic(err)
			}

			activityRule := activityBase.ActivityRule

			if activityRule != "" {
				itemNumVoMap, err = service.ActivityBase().GetActivityItemNum(ctx, activityBase)
			}

			// 礼包活动 必须套餐直接下单。
			if activityBase.ActivityTypeId == consts.ACTIVITY_TYPE_GIFTBAG {
				var items []*model.CheckoutItemVo

				for itemId, itemNumVo := range itemNumVoMap {
					item := &model.CheckoutItemVo{
						ItemId:       itemId,
						CartQuantity: itemNumVo.Num,
						CartSelect:   true,
						CartType:     1,
					}
					items = append(items, item)
				}

				in.Items = items
			}

			// 砍价活动
			if activityBase.ActivityTypeId == consts.ACTIVITY_TYPE_CUTPRICE {
				var items []*model.CheckoutItemVo

				for itemId, itemNumVo := range itemNumVoMap {
					item := &model.CheckoutItemVo{
						ItemId:       itemId,
						CartQuantity: itemNumVo.Num,
						CartSelect:   true,
						CartType:     1,
					}
					items = append(items, item)
				}

				in.Items = items
			}
		} else {
			return nil, errors.New("非法活动参数！")
		}
	}

	orderProductAmount := decimal.NewFromFloat(0) // 商品订单原价
	orderItemAmount := decimal.NewFromFloat(0)    // 单品优惠后价格累加
	orderFreightAmount := decimal.NewFromFloat(0)
	orderMoneyAmount := decimal.NewFromFloat(0)
	orderDiscountAmount := decimal.NewFromFloat(0)
	orderPointsAmount := decimal.NewFromFloat(0)
	orderSpAmount := decimal.NewFromFloat(0)

	itemIds := gconv.SliceUint64(array.Column(in.Items, "ItemId"))
	storeIds := gconv.SliceUint(array.Column(in.Items, "StoreId"))

	productItemList, err := service.ProductBase().GetItems(ctx, itemIds, in.UserId)

	// 活动商品数量 多件折 等使用
	activityItemQuantityTotalMap := make(map[uint]uint)

	// 店铺分组
	storeItemsMap := make(map[uint][]*model.ProductItemVo)

	for _, it := range productItemList {
		// 根据店铺分组数据
		storeItemsList := storeItemsMap[it.StoreId]
		if storeItemsList == nil {
			storeItemsMap[it.StoreId] = []*model.ProductItemVo{}
		}

		// 设置购物车商品数量
		var checkoutItemVo *model.CheckoutItemVo
		for _, s := range in.Items {
			if s.ItemId == it.ItemId {
				checkoutItemVo = s
				break
			}
		}

		// 处理商品购买信息
		it.CartId = checkoutItemVo.CartId
		it.CartQuantity = checkoutItemVo.CartQuantity
		it.CartSelect = checkoutItemVo.CartSelect

		// 判断可用库存
		it.AvailableQuantity = it.ItemQuantity - it.ItemQuantityFrozen
		it.IsOnSale = service.ProductItem().IfOnSale(ctx, it)

		// 商品是否可销售
		if checkoutItemVo.CartSelect {
			it.CartSelect = it.IsOnSale
		}

		// 判断是否在配送区域
		if checkoutItemVo.CartSelect {
			it.CartSelect = true
		}

		storeItemsList = append(storeItemsList, it)
		storeItemsMap[it.StoreId] = storeItemsList
	}

	deliveryAddress := &entity.UserDeliveryAddress{}

	// 配送地址 || 联系方式
	if g.IsEmpty(in.UdId) {
		// 默认配送地址
		defaultAddress, _ := dao.UserDeliveryAddress.FindOne(ctx, &do.UserDeliveryAddressListInput{
			Where: do.UserDeliveryAddress{
				UserId:      in.UserId,
				UdIsDefault: 1,
			},
		})

		// 配送地址为空，自动将第一个地址填入订单中
		if defaultAddress == nil {
			deliveryAddress, _ = dao.UserDeliveryAddress.FindOne(ctx, &do.UserDeliveryAddressListInput{
				Where: do.UserDeliveryAddress{
					UserId: in.UserId,
				},
			})
		} else {
			deliveryAddress = defaultAddress
		}
	} else {
		deliveryAddress, _ = dao.UserDeliveryAddress.Get(ctx, in.UdId)
	}

	out.UserDeliveryAddress = deliveryAddress

	//店铺信息
	//List<StoreBase> storeLists = storeBaseService.gets(storeIds);

	//可用代金券列表
	now := gtime.Now()
	var ext = []*ml.WhereExt{{
		Column: dao.UserVoucher.Columns().VoucherStartDate,
		Val:    now,
		Symbol: ml.GE,
	}, {
		Column: dao.UserVoucher.Columns().VoucherEndDate,
		Val:    now,
		Symbol: ml.LE,
	}}

	userVoucherListInput := &do.UserVoucherListInput{
		BaseList: ml.BaseList{Page: 1,
			Size:     consts.MAX_LIST_NUM,
			WhereExt: ext,
			Sidx:     dao.UserVoucher.Columns().UserVoucherTime,
			Sort:     "DESC"},
		Where: do.UserVoucher{
			UserId:         in.UserId,
			VoucherStateId: consts.VOUCHER_STATE_UNUSED,
		},
	}

	userVoucherPage, _ := service.UserVoucher().GetList(ctx, userVoucherListInput)
	voucherItems := userVoucherPage.Items

	// 店铺信息
	for _, storeId := range storeIds {
		// 是否虚拟商品
		isVirtual := false
		var kindId uint

		// 是否需要配送
		isDelivery := false

		if !g.IsEmpty(in.DeliveryTypeId) && in.DeliveryTypeId == consts.DELIVERY_TYPE_EXP {
			isDelivery = true
		}

		storeItemVo := &model.StoreItemVo{}
		storeItemVo.RedemptionItems = make([]*model.ActivitysVo, 0)
		storeItemVo.VoucherItems = make([]*model.UserVoucherRes, 0)
		storeItemVo.ActivityBase = activityBase
		items := storeItemsMap[storeId]

		if len(items) == 0 {
			continue
		}

		var productAmount, freightAmount, discountAmount, moneyAmount, itemAmount, pointsAmount, spAmount, itemSubtotal decimal.Decimal

		// 涉及商品个数
		itemSelectedSize := 0

		for _, item := range items {
			if !item.CartSelect {
				continue
			}

			itemSelectedSize++

			// todo 处理单品活动价格

			if item.ActivityInfo != nil {
				if item.ActivityInfo.ActivityTypeId == consts.ACTIVITY_TYPE_BATDISCOUNT {
					// 处理批发价
					percent := decimal.NewFromFloat(100)

					if item.ActivityInfo.ActivityBase.ActivityRule != "" {
						var activityRuleVo model.ActivityRuleVo
						err := json.Unmarshal([]byte(item.ActivityInfo.ActivityBase.ActivityRule), &activityRuleVo)
						if err != nil {
							return nil, err
						}

						rules := activityRuleVo.Rule

						// 使用 Comparator 进行排序
						sort.Slice(rules, func(i, j int) bool {
							return rules[i].Num < rules[j].Num
						})

						for _, r := range rules {
							if item.CartQuantity >= r.Num {
								percent = decimal.NewFromFloat(r.Percent)
							} else {
								break
							}
						}
					}

					// 是否触发阶梯价
					if percent.Equal(decimal.NewFromFloat(100)) {
						// 未触发，清除活动标记
						item.ActivityInfo = nil
						item.ActivityId = 0
					} else {
						salePrice := decimal.NewFromFloat(item.ItemUnitPrice).Mul(percent).Div(decimal.NewFromFloat(100)).Round(2)
						item.ItemSalePrice, _ = salePrice.Float64()

						savePrice := decimal.NewFromFloat(item.ItemUnitPrice).Sub(salePrice)
						item.ItemSavePrice, _ = savePrice.Float64()

						item.ItemDiscountAmount, _ = savePrice.Mul(decimal.NewFromFloat(float64(item.CartQuantity))).Float64()
					}
				} else if item.ActivityInfo.ActivityTypeId == consts.ACTIVITY_TYPE_MULTIPLEDISCOUNT {
					// 处理多件折
					// 需要根据活动提前计算好
					if g.IsEmpty(activityItemQuantityTotalMap[item.ActivityId]) {
						activityItemIds := gconv.SliceUint64(gstr.Split(item.ActivityInfo.ActivityBase.ActivityItemIds, ","))
						var activityItemQuantityTotal uint = 0

						for _, s := range items {
							if s.CartSelect && array.InArray(activityItemIds, s.ItemId) {
								activityItemQuantityTotal += s.CartQuantity
							}
						}

						activityItemQuantityTotalMap[item.ActivityId] = activityItemQuantityTotal
					}

					activityItemQuantityTotal := activityItemQuantityTotalMap[item.ActivityId]

					percent := decimal.NewFromFloat(100)

					if item.ActivityInfo.ActivityBase.ActivityRule != "" {
						var activityRuleVo model.ActivityRuleVo
						err := json.Unmarshal([]byte(item.ActivityInfo.ActivityBase.ActivityRule), &activityRuleVo)
						if err != nil {
							return nil, err
						}

						rules := activityRuleVo.Rule

						// 使用 Comparator 进行排序
						sort.Slice(rules, func(i, j int) bool {
							return rules[i].Num < rules[j].Num
						})

						for _, r := range rules {
							if activityItemQuantityTotal >= r.Num {
								percent = decimal.NewFromFloat(r.Percent)
							} else {
								break
							}
						}
					}

					// 是否触发阶梯价
					if percent.Equal(decimal.NewFromFloat(100)) {
						// 未触发，清除活动标记
						item.ActivityInfo = nil
						item.ActivityId = 0
					} else {
						salePrice := decimal.NewFromFloat(item.ItemUnitPrice).Mul(percent).Div(decimal.NewFromFloat(100)).Round(2)
						item.ItemSalePrice, _ = salePrice.Float64()

						savePrice := decimal.NewFromFloat(item.ItemUnitPrice).Sub(salePrice)
						item.ItemSavePrice, _ = savePrice.Float64()

						item.ItemDiscountAmount, _ = savePrice.Mul(decimal.NewFromFloat(float64(item.CartQuantity))).Float64()
					}
				}
				// 会员等级判断
				if item.ActivityInfo != nil {
					activity := item.ActivityInfo.ActivityBase

					if activity != nil {
						_, err := s.checkoutLevel(ctx, activity.ActivityUseLevel, userInfo.UserLevelId)
						if err != nil {
							panic(err)
						}
					}
				}
			}

			itemOriSubtotal := decimal.NewFromFloat(item.ItemUnitPrice).Mul(decimal.NewFromFloat(float64(item.CartQuantity)))
			itemSubtotal = decimal.NewFromFloat(item.ItemSalePrice).Mul(decimal.NewFromFloat(float64(item.CartQuantity)))
			itemPointsSubtotal := decimal.NewFromFloat(item.ItemUnitPoints).Mul(decimal.NewFromFloat(float64(item.CartQuantity)))

			// 汇总
			productAmount = productAmount.Add(itemOriSubtotal)
			itemAmount = itemAmount.Add(itemSubtotal)
			moneyAmount = moneyAmount.Add(itemSubtotal)
			discountAmount = productAmount.Sub(moneyAmount)
			pointsAmount = pointsAmount.Add(itemPointsSubtotal)
			// spAmount = spAmount.Add(item.ItemUnitSp)

			item.ItemPointsSubtotal, _ = itemPointsSubtotal.Float64()
			item.ItemSubtotal, _ = itemSubtotal.Float64()
			item.ItemDiscountAmount, _ = itemOriSubtotal.Sub(itemSubtotal).Float64()

			// 为虚拟商品，有一个就是虚拟的
			kinds := []uint{consts.PRODUCT_KIND_FUWU, consts.PRODUCT_KIND_CARD, consts.PRODUCT_KIND_EDU}
			if array.InArray(kinds, item.KindId) {
				isVirtual = true
			}

			kindId = item.KindId
		}

		// 过滤掉不可使用的优惠券，
		for i := len(voucherItems) - 1; i >= 0; i-- {
			voucherItem := voucherItems[i]
			// 指定优惠券
			itemIdStr := voucherItem.ItemId
			voucherSubtotal := decimal.NewFromFloat(voucherItem.VoucherSubtotal)

			if itemIdStr != "" {
				itemIdList := gconv.SliceUint64(gstr.Split(itemIdStr, ","))

				// 指定优惠券itemIdList不包含购买的任何商品
				if array.Disjoint(itemIdList, itemIds) {
					// 删除不符合条件的优惠券
					voucherItems = append(voucherItems[:i], voucherItems[i+1:]...)
					continue
				}

				// 使用优惠券的订单金额 （指定）
				var assignProductAmount decimal.Decimal
				for _, item := range items {
					if array.InArray(itemIdList, item.ItemId) {
						itemSalSubtotal := decimal.NewFromFloat(item.ItemSalePrice).Mul(decimal.NewFromFloat(float64(item.CartQuantity)))
						assignProductAmount = assignProductAmount.Add(itemSalSubtotal)
					}
				}

				if assignProductAmount.Cmp(voucherSubtotal) < 0 {
					voucherItems = append(voucherItems[:i], voucherItems[i+1:]...)
					continue
				}
			}

			// 使用优惠券的订单金额 （全部）
			if itemAmount.Cmp(voucherSubtotal) < 0 {
				voucherItems = append(voucherItems[:i], voucherItems[i+1:]...)
			}
		}

		storeItemVo.IsVirtual = isVirtual
		storeItemVo.KindId = kindId
		storeItemVo.Items = items
		storeItemVo.MoneyItemAmount, _ = itemAmount.Float64()

		// 店铺选中的使用代金券
		var voucherItemSelected *model.UserVoucherRes

		for _, voucherItem := range voucherItems {
			if voucherItem.StoreId == storeId && voucherItem.UserId == in.UserId {
				storeItemVo.VoucherItems = append(storeItemVo.VoucherItems, voucherItem)

				// 判断是否为选中的
				if array.InArray(in.UserVoucherIds, voucherItem.UserVoucherId) {
					voucherItemSelected = voucherItem
				}
			}
		}

		// 处理店铺活动
		if !g.IsEmpty(in.ActivityId) {
			if activityBase != nil {
				var activityRuleVo model.ActivityRuleVo
				err := json.Unmarshal([]byte(activityBase.ActivityRule), &activityRuleVo)
				if err != nil {
					return nil, err
				}

				if !g.IsNil(activityRuleVo) {
					// 礼包活动 必须套餐直接下单。
					if activityBase.ActivityTypeId == consts.ACTIVITY_TYPE_GIFTBAG || activityBase.ActivityTypeId == consts.ACTIVITY_TYPE_CUTPRICE {
						if activityBase.ActivityTypeId == consts.ACTIVITY_TYPE_GIFTBAG {
							// 组合套餐活动
							giftbag := activityRuleVo.Giftbag

							if g.IsNil(giftbag) {
								panic(errors.New("组合套餐信息不存在！"))
							}
							moneyAmount = decimal.NewFromFloat(giftbag.GiftbagAmount)
						} else if activityBase.ActivityTypeId == consts.ACTIVITY_TYPE_CUTPRICE {
							// 砍价活动
							activityCutprice, err := dao.ActivityCutprice.FindOne(ctx, &do.ActivityCutpriceListInput{Where: do.ActivityCutprice{
								ActivityId: in.ActivityId,
								UserId:     in.UserId,
							}})

							if activityCutprice == nil || err != nil {
								panic(errors.New("砍价记录信息不存在！"))
							}

							moneyAmount = decimal.NewFromFloat(activityCutprice.AcSalePrice)
						}

						discountAmount = productAmount.Sub(moneyAmount)

						// 套餐优惠比例
						discountRate := moneyAmount.Div(productAmount).Round(6)
						newAmount := decimal.NewFromFloat(0) // 已分配额度

						var i int
						// 取出参与的产品的总值
						for _, item := range items {
							if !item.CartSelect {
								continue
							}

							i++

							itemOriSubtotal := decimal.NewFromFloat(item.ItemUnitPrice).Mul(decimal.NewFromFloat(float64(item.CartQuantity)))

							// 最后一个商品
							if i == itemSelectedSize {
								itemNewSubtotal := moneyAmount.Sub(newAmount)
								item.ItemSalePrice, _ = itemNewSubtotal.Div(decimal.NewFromFloat(float64(item.CartQuantity))).Round(6).Float64()
								item.ItemSavePrice = item.ItemUnitPrice - item.ItemSalePrice
								item.ItemSubtotal, _ = itemNewSubtotal.Float64()
								item.ItemDiscountAmount, _ = itemOriSubtotal.Sub(itemNewSubtotal).Float64()
							} else {
								itemNewSubtotal := itemOriSubtotal.Mul(discountRate)
								// item.ItemPointsSubtotal = itemPointsSubtotal
								item.ItemSalePrice, _ = itemNewSubtotal.Div(decimal.NewFromFloat(float64(item.CartQuantity))).Round(6).Float64()
								item.ItemSavePrice = item.ItemUnitPrice - item.ItemSalePrice
								item.ItemSubtotal, _ = itemNewSubtotal.Float64()
								item.ItemDiscountAmount, _ = itemOriSubtotal.Sub(itemNewSubtotal).Float64()

								newAmount = newAmount.Add(itemNewSubtotal)
							}
						}
					}
				}
			}
		}

		// 是否需要配送地址
		if isVirtual {
			// todo 判断是否有需要上门服务，需要计算配送费
			isDelivery = false
		}

		// 是否需要配送地址， 自提及虚拟到店服务，不需要配送地址。
		if isDelivery {
			// 计算运费
			// 如果没有配送地址，则忽略地址选择问题。
			if deliveryAddress != nil && !g.IsEmpty(deliveryAddress.UdCityId) {
				districtId := deliveryAddress.UdCityId
				freightAmountFloat, err := s.CalTransportFreight(ctx, storeItemVo, districtId) // 配送检测

				if err != nil {
					return nil, err
				}

				freightAmount = decimal.NewFromFloat(freightAmountFloat)
			} else {
				// throw new BusinessException(__("请选择正确的收货地址！"))
			}
			// end 运费计算
		}

		// 优惠券 voucherItemSelected 修正最终付款价格： moneyAmount， 不修正itemAmount原价
		if voucherItemSelected != nil {
			storeItemVo.UserVoucherId = voucherItemSelected.UserVoucherId
			storeItemVo.VoucherAmount = voucherItemSelected.VoucherPrice
			// moneyAmount = moneyAmount.Sub(voucherItemSelected.VoucherPrice)
		}

		moneyAmount = moneyAmount.Add(freightAmount).Sub(decimal.NewFromFloat(storeItemVo.VoucherAmount))

		storeItemVo.DiscountAmount, _ = discountAmount.Float64()
		storeItemVo.ProductAmount, _ = productAmount.Float64()
		storeItemVo.MoneyAmount, _ = moneyAmount.Float64() // 下单时候运费及活动会修正此值
		storeItemVo.MoneyItemAmount, _ = itemAmount.Float64()
		storeItemVo.PointsAmount, _ = pointsAmount.Float64()
		storeItemVo.FreightAmount, _ = freightAmount.Float64()

		out.Items = append(out.Items, *storeItemVo)

		orderProductAmount = orderProductAmount.Add(productAmount)
		orderItemAmount = orderItemAmount.Add(itemAmount)
		orderMoneyAmount = orderMoneyAmount.Add(moneyAmount)
		orderFreightAmount = orderFreightAmount.Add(freightAmount)
		orderDiscountAmount = orderDiscountAmount.Add(discountAmount)
		orderPointsAmount = orderPointsAmount.Add(pointsAmount)
		orderSpAmount = orderSpAmount.Add(spAmount)
	}

	out.OrderProductAmount, _ = orderProductAmount.Float64()
	out.OrderItemAmount, _ = orderProductAmount.Float64()
	out.OrderFreightAmount, _ = orderFreightAmount.Float64()
	out.OrderMoneyAmount, _ = orderMoneyAmount.Float64()
	out.OrderDiscountAmount, _ = orderDiscountAmount.Float64()
	out.OrderPointsAmount, _ = orderPointsAmount.Float64()
	out.OrderSpAmount, _ = orderSpAmount.Float64()

	return out, err
}

func (s *sUserCart) checkoutLevel(ctx context.Context, activityUseLevel string, userLevelId uint) (out bool, err error) {
	if activityUseLevel != "" {
		userLevels := gconv.SliceInt(gstr.Split(activityUseLevel, ","))
		if len(userLevels) > 0 {
			sort.Ints(userLevels)
			userLevelList, err := dao.UserLevel.List(ctx, &do.UserLevelListInput{Where: do.UserLevel{UserLevelId: userLevels}})
			if err != nil {
				return false, errors.New("等级信息不存在！")
			}

			if !array.InArray(userLevels, userLevelId) {
				names := array.Column(userLevelList, "UserLevelName")
				return false, errors.New(fmt.Sprintf("活动商品会员等级为 %s ，用户等级未达到", gstr.JoinAny(names, "、")))
			}
		}
	}

	return
}

func (s *sUserCart) CalTransportFreight(ctx context.Context, storeItemVo *model.StoreItemVo, districtId uint) (out float64, err error) {
	// 运费模板商品数量
	ttIdsMap := make(map[uint]uint)
	transportTypeNoneIds := make([]uint, 0)
	deliveryItemNoneRow := make([]*model.ProductItemVo, 0)

	// 运费
	freight := decimal.NewFromFloat(0)
	freightFreeAmountMax := decimal.NewFromFloat(0)
	var ifEnabledFreeFreight bool

	items := storeItemVo.Items

	// 处理 ttIdsMap
	for _, pi := range items {
		quantity := ttIdsMap[pi.TransportTypeId]
		quantity += pi.CartQuantity
		ttIdsMap[pi.TransportTypeId] = quantity
	}

	if ifEnabledFreeFreight {

	} else {
		ttIds := gconv.SliceUint(array.Column(items, "TransportTypeId"))
		var storeTransportTypes []*entity.StoreTransportType

		// 判断运费方式，如果发现同一个订单计费模式不一致，报错，禁止下单。
		if len(ttIds) > 0 {
			storeTransportTypes, err = dao.StoreTransportType.Gets(ctx, ttIds)
			transportTypePricingMethod := make(map[uint]bool)

			for _, tt := range storeTransportTypes {
				transportTypePricingMethod[tt.TransportTypePricingMethod] = true
			}

			if len(transportTypePricingMethod) > 1 {
				return 0, errors.New("所选商品运费模式不统一，请拆分下单！")
			}
		} else {
			return 0, errors.New("商品运费设置有误！请联系商家检查商品设置！")
		}

		// 按件计费
		for _, ttId := range ttIds {
			quantity := ttIdsMap[ttId]
			ttTransportTypeId := ttId
			var transportType *entity.StoreTransportType

			// 此处免运费，传递的是 transportType 中的设置。
			for _, tt := range storeTransportTypes {
				if tt.TransportTypeId == ttId {
					transportType = tt
					continue
				}
			}

			var freightFreeAmount float64
			if transportType != nil {
				freightFreeAmount = transportType.TransportTypeFreightFree
			}

			freightFreeAmountMax = decimal.Max(freightFreeAmountMax, decimal.NewFromFloat(freightFreeAmount))

			data, err := service.StoreTransportType().CalFreight(ctx, ttTransportTypeId, districtId, quantity, storeItemVo.MoneyItemAmount, freightFreeAmount)

			if err != nil {
				return 0, err
			}

			typeFreight := data.CanDelivery
			_freight := data.Freight

			if !typeFreight {
				// 配送不到这个区域，提示删除商品
				transportTypeNoneIds = append(transportTypeNoneIds, ttTransportTypeId)
			} else {
				freight = freight.Add(decimal.NewFromFloat(_freight))
			}
		}

		// 配送区域无货设置；
		if len(transportTypeNoneIds) > 0 {
			for _, itemRow := range items {
				// 配送区域库存问题。
				showOOS := array.InArray(transportTypeNoneIds, itemRow.TransportTypeId)
				itemRow.IsOos = showOOS

				deliveryItemNoneRow = append(deliveryItemNoneRow, itemRow)
			}
		}
	}

	if decimal.NewFromFloat(storeItemVo.MoneyItemAmount).Cmp(freightFreeAmountMax) < 0 {
		storeItemVo.FreightFreeBalance, _ = freightFreeAmountMax.Sub(decimal.NewFromFloat(storeItemVo.MoneyItemAmount)).Abs().Float64()
	} else {
		storeItemVo.FreightFreeBalance = 0
	}

	out, _ = freight.Float64()

	return out, err
}

// List 分页读取
func (s *sUserCart) List(ctx context.Context, in *do.UserCartListInput) (out *do.UserCartListOutput, err error) {
	out, err = dao.UserCart.List(ctx, in)

	return out, err
}

// AddCart 新增
func (s *sUserCart) AddCart(ctx context.Context, in *model.CartAddInput) (res bool, err error) {
	// 流程可配置 节点触发api调用

	// todo 判断库存，提示加入购物车
	if in.CartQuantity <= 0 {
		panic(fmt.Sprintf("最低备货数量 1 件，请确认！"))
	}

	// 获取商品信息
	productItem, err := dao.ProductItem.Get(ctx, in.ItemId)
	if err != nil {
		panic(fmt.Sprintf("获取商品信息失败：%v", err))
	}

	if productItem.ItemEnable != consts.PRODUCT_STATE_NORMAL {
		panic("商品未上架，不可加入购物车！")
	}

	// 判断可用库存
	availableQuantity := productItem.ItemQuantity - productItem.ItemQuantityFrozen

	// 查找购物车中是否已经有该商品
	cart, err := dao.UserCart.FindOne(ctx, &do.UserCartListInput{Where: do.UserCart{UserId: in.UserId, ItemId: in.ItemId}})
	if err != nil {
		panic(fmt.Sprintf("查找购物车信息失败：%v", err))
	}

	if cart == nil {
		if in.CartQuantity > availableQuantity {
			panic(fmt.Sprintf("库存可用数量 %d 件，请确认！", availableQuantity))
		}

		cart = &entity.UserCart{
			UserId:       in.UserId,
			ItemId:       in.ItemId,
			ProductId:    productItem.ProductId,
			CartQuantity: in.CartQuantity,
			CartSelect:   in.CartSelect,
			ActivityId:   in.ActivityId,
			CartType:     in.CartType,
			StoreId:      in.StoreId,
			ChainId:      in.ChainId,
		}
	} else {
		if in.CartQuantity+cart.CartQuantity > availableQuantity {
			panic(fmt.Sprintf("库存可用数量 %d 件，请确认！", availableQuantity))
		}

		cart.CartQuantity = cart.CartQuantity + in.CartQuantity
	}

	// 购物车当前商品
	userCartQueryInput := &do.UserCartListInput{Where: do.UserCart{UserId: in.UserId, CartSelect: true}}
	userCarts, err := dao.UserCart.Find(ctx, userCartQueryInput)
	if err != nil {
		panic(fmt.Sprintf("查找用户购物车信息失败：%v", err))
	}

	if len(userCarts) > 0 {
		// 将当前购物车项加入购物车列表
		if cart.CartSelect {
			userCarts = append(userCarts, cart)
		}

		checkoutInput := &model.CheckoutInput{
			UserId: in.UserId,
			Items:  make([]*model.CheckoutItemVo, len(userCarts)),
		}
		for i, c := range userCarts {
			checkoutInput.Items[i] = &model.CheckoutItemVo{
				ItemId:       c.ItemId,
				CartSelect:   c.CartSelect,
				CartQuantity: c.CartQuantity,
			}
		}
		s.Checkout(ctx, checkoutInput)
	}

	// 保存购物车项
	cartDo := &do.UserCart{}
	gconv.Scan(cart, cartDo)

	_, err = dao.UserCart.Save(ctx, cartDo)

	return err != nil, err
}

// Sel 选中状态
func (s *sUserCart) Sel(ctx context.Context, input *model.UserCartSelectInput) (res bool, err error) {
	userId := input.UserId
	cartId := input.CartId
	storeId := input.StoreId
	action := input.Action

	var cartIds []uint64

	switch action {
	case "all":
		// 查询所有购物车项
		cartList, _ := dao.UserCart.Find(ctx, &do.UserCartListInput{Where: do.UserCart{UserId: userId}})
		for _, cart := range cartList {
			cartIds = append(cartIds, cart.CartId)
		}
	case "store":
		// 查询指定店铺的购物车项
		cartList, _ := dao.UserCart.Find(ctx, &do.UserCartListInput{Where: do.UserCart{UserId: userId, StoreId: storeId}})
		for _, cart := range cartList {
			cartIds = append(cartIds, cart.CartId)
		}
	default:
		cartIds = append(cartIds, cartId)
	}

	// 更新购物车项的选择状态
	updateInput := &do.UserCartListInput{
		Where: do.UserCart{CartId: cartIds},
	}
	data := &do.UserCart{CartSelect: input.CartSelect}
	if _, err := dao.UserCart.EditWhere(ctx, updateInput, data); err != nil {
		panic(fmt.Sprintf("更改购物车选中状态失败：%s", err))
	}

	// 活动商品 达到会员等级判断
	if input.CartSelect {
		s.checkoutCart(ctx, userId, input.CartSelect)
	}

	return
}

// checkoutCart 实现了结账购物车的逻辑
func (s *sUserCart) checkoutCart(ctx context.Context, userId uint, cartSelect bool) {
	if cartSelect {
		// 查询选择的购物车项
		cartList, _ := dao.UserCart.Find(ctx, &do.UserCartListInput{Where: do.UserCart{UserId: userId, CartSelect: true}})
		if len(cartList) == 0 {
			return
		}

		// 构造结账输入
		checkoutInput := &model.CheckoutInput{
			UserId: userId,
			Items:  make([]*model.CheckoutItemVo, len(cartList)),
		}
		for i, cart := range cartList {
			checkoutInput.Items[i] = &model.CheckoutItemVo{
				CartId:       cart.CartId,
				StoreId:      cart.StoreId,
				CartSelect:   cart.CartSelect,
				ItemId:       cart.ItemId,
				CartQuantity: cart.CartQuantity,
			}
		}

		// 执行结账操作
		s.Checkout(ctx, checkoutInput)
	}
}

// EditQuantity 编辑数量
func (s *sUserCart) EditQuantity(ctx context.Context, userCart *do.UserCart, userId uint) (affected int64, err error) {
	// 根据购物车Id获取购物车信息
	cart, err := dao.UserCart.Get(ctx, userCart.CartId)
	if err != nil {
		return 0, err
	}

	// 检查购物车是否存在，且用户ID匹配
	if cart != nil && cart.UserId == userId {
		// 减少购物车数量
		if userCart.CartQuantity.(uint) == 0 {
			// 删除购物车项
			return dao.UserCart.Remove(ctx, cart.CartId)
		} else {
			// 增加购物车数量
			// 获取商品信息
			productItem, err := dao.ProductItem.Get(ctx, cart.ItemId)
			if err != nil {
				return 0, err
			}
			if productItem == nil {
				return 0, errors.New("该商品不存在！")
			}
			// 判断可用库存
			availableQuantity := productItem.ItemQuantity - productItem.ItemQuantityFrozen
			if userCart.CartQuantity.(uint) > availableQuantity {
				return 0, errors.New(fmt.Sprintf("库存可用数量 %d 件，请确认！", availableQuantity))
			}
			// 更新购物车数量
			cart.CartQuantity = userCart.CartQuantity.(uint)
			cartDo := &do.UserCart{}
			gconv.Struct(cart, cartDo)
			if affected, err = dao.UserCart.Save(ctx, cartDo); err != nil {
				return 0, err
			}

			// 检查购物车中活动商品是否达到会员等级
			s.checkoutCart(ctx, userId, cart.CartSelect)
			return affected, nil
		}
	}

	return
}

// Edit 编辑数量
func (s *sUserCart) Edit(ctx context.Context, in *do.UserCart) (affected int64, err error) {
	_, err = dao.UserCart.Edit(ctx, in.CartId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sUserCart) Remove(ctx context.Context, id any) (affected int64, err error) {

	affected, err = dao.UserCart.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// Add 编辑数量
func (s *sUserCart) Add(ctx context.Context, in *do.UserCart) (affected int64, err error) {
	_, err = dao.UserCart.Add(ctx, in)

	if err != nil {
		return 0, err
	}
	return
}
