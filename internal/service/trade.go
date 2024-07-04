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

	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
)

type (
	IDistributionOrderItem interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.DistributionOrderItemListInput) (out []*entity.DistributionOrderItem, err error)
		// List 分页读取
		List(ctx context.Context, in *do.DistributionOrderItemListInput) (out *do.DistributionOrderItemListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.DistributionOrderItem) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.DistributionOrderItem) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		SettleDistributionUserOrder(ctx context.Context, orderId string) (flag bool, err error)
	}
	IOrderData interface {
		// Get 读取订单
		Get(ctx context.Context, id any) (out *entity.OrderData, err error)
		// Gets 读取多条订单
		Gets(ctx context.Context, id any) (list []*entity.OrderData, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.OrderDataListInput) (out []*entity.OrderData, err error)
		// List 分页读取
		List(ctx context.Context, in *do.OrderDataListInput) (out *do.OrderDataListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.OrderData) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.OrderData) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IOrderLogistics interface {
		// Get 读取订单
		Get(ctx context.Context, id any) (out *entity.OrderLogistics, err error)
		// Gets 读取多条订单
		Gets(ctx context.Context, id any) (list []*entity.OrderLogistics, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.OrderLogisticsListInput) (out []*entity.OrderLogistics, err error)
		// List 分页读取
		List(ctx context.Context, in *do.OrderLogisticsListInput) (out *do.OrderLogisticsListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.OrderLogistics) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.OrderLogistics) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IOrderReturn interface {
		// Get 读取订单
		Get(ctx context.Context, id any) (out *entity.OrderReturn, err error)
		// Gets 读取多条订单
		Gets(ctx context.Context, id any) (list []*entity.OrderReturn, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.OrderReturnListInput) (out []*entity.OrderReturn, err error)
		// List 分页读取
		List(ctx context.Context, in *do.OrderReturnListInput) (out *do.OrderReturnListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.OrderReturn) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.OrderReturn) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// GetByReturnId 根据退货 ID 获取订单退货信息
		GetByReturnId(ctx context.Context, returnId string) (*model.OrderReturnVo, error)
		// Refused 拒绝退货退款
		Refused(ctx context.Context, orderReturn *do.OrderReturn) (bool, error)
		// Review 审核退货退款
		Review(ctx context.Context, orderReturn *do.OrderReturn, receivingAddress uint) (bool, error)
		// GetList 获取订单退货列表
		GetList(ctx context.Context, input *do.OrderReturnListInput) (res *model.OrderReturnOutput, err error)
		GetReturn(ctx context.Context, returnId any) (*model.OrderReturnVo, error)
		// EditReturn 编辑退货
		EditReturn(ctx context.Context, orderReturn *do.OrderReturn) (bool, error)
		// Cancel 取消订单
		Cancel(ctx context.Context, returnId string, userId uint) (bool, error)
		// ReturnItem 读取订单退货详情
		ReturnItem(ctx context.Context, orderId string, orderItemId string, userId uint) (*model.OrderReturnItemVo, error)
		// AddItem 添加订单退货
		AddItem(ctx context.Context, orderReturnInput *model.OrderReturnInput) (returnId string, err error)
		AddReturnByItem(ctx context.Context, orderReturn *do.OrderReturn, returnItems []*do.OrderReturnItem) (bool, error)
	}
	IOrderReturnItem interface {
		// Get 读取订单
		Get(ctx context.Context, id any) (out *entity.OrderReturnItem, err error)
		// Gets 读取多条订单
		Gets(ctx context.Context, id any) (list []*entity.OrderReturnItem, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.OrderReturnItemListInput) (out []*entity.OrderReturnItem, err error)
		// List 分页读取
		List(ctx context.Context, in *do.OrderReturnItemListInput) (out *do.OrderReturnItemListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.OrderReturnItem) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.OrderReturnItem) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IOrderStateLog interface {
		// Get 读取订单
		Get(ctx context.Context, id any) (out *entity.OrderStateLog, err error)
		// Gets 读取多条订单
		Gets(ctx context.Context, id any) (list []*entity.OrderStateLog, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.OrderStateLogListInput) (out []*entity.OrderStateLog, err error)
		// List 分页读取
		List(ctx context.Context, in *do.OrderStateLogListInput) (out *do.OrderStateLogListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.OrderStateLog) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.OrderStateLog) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IUserCart interface {
		// List 分页读取
		GetList(ctx context.Context, in *do.UserCartListInput) (out *model.CheckoutOutput, err error)
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
		Checkout(ctx context.Context, in *model.CheckoutInput) (out *model.CheckoutOutput, err error)
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
		FormatCartRows(ctx context.Context, in *model.CheckoutInput) (out *model.CheckoutOutput, err error)
		CalTransportFreight(ctx context.Context, storeItemVo *model.StoreItemVo, districtId uint) (out float64, err error)
		// List 分页读取
		List(ctx context.Context, in *do.UserCartListInput) (out *do.UserCartListOutput, err error)
		// AddCart 新增
		AddCart(ctx context.Context, in *model.CartAddInput) (res bool, err error)
		// Sel 选中状态
		Sel(ctx context.Context, input *model.UserCartSelectInput) (res bool, err error)
		// EditQuantity 编辑数量
		EditQuantity(ctx context.Context, userCart *do.UserCart, userId uint) (affected int64, err error)
		// Edit 编辑数量
		Edit(ctx context.Context, in *do.UserCart) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// Add 编辑数量
		Add(ctx context.Context, in *do.UserCart) (affected int64, err error)
	}
	IOrder interface {
		// Detail 读取订单
		Detail(ctx context.Context, orderId any) (detail *model.OrderVo, err error)
		// List 订单搜索查询列表
		List(ctx context.Context, in *do.OrderInfoListInput) (out *model.OrderListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *model.CheckoutInput) (out *model.OrderAddOutput, err error)
		/**
		 * 取消订单
		 *
		 * @param orderId       订单编号
		 * @param orderStateNote 订单状态备注
		 * @return bool 是否取消成功
		 */
		Cancel(ctx context.Context, orderId string, orderStateNote string) (flag bool, err error)
		/**
		 * 支付完成
		 *
		 * @param orderId 订单编号
		 * @return 是否支付成功
		 */
		SetPaidYes(ctx context.Context, orderId string) (flag bool, err error)
		// Review 审核订单
		Review(ctx context.Context, orderId string, orderStateNote string) (flag bool, err error)
		// Finance 财务审核
		Finance(ctx context.Context, orderId string, orderStateNote string) (flag bool, err error)
		// Picking 出库审核
		Picking(ctx context.Context, in *model.OrderPickingInput) (flag bool, err error)
		// CheckOrderReturnWaiting 判断是否有待审核售后订单条件限制
		CheckOrderReturnWaiting(ctx context.Context, orderId string) (flag bool, err error)
		// Shipping 发货
		Shipping(ctx context.Context, in *model.OrderShippingInput) (flag bool, err error)
		// CheckShippingComplete 检测是否发货完成
		CheckShippingComplete(ctx context.Context, orderId string) (isComplete bool, err error)
		// AddLogistics 添加订单日志
		AddLogistics(ctx context.Context, in *do.OrderLogistics) (flag bool, err error)
		// SaveLogistics 添加订单日志
		SaveLogistics(ctx context.Context, in *do.OrderLogistics) (flag bool, err error)
		// Receive 收货
		Receive(ctx context.Context, orderId string, orderStateNote string) (flag bool, err error)
		// EditNextState 修改订单为下一个待处理状态
		EditNextState(ctx context.Context, orderId string, orderStateId uint, nextOrderStateId uint, orderStateNote string) (flag bool, err error)
		// DoReviewPicking 出库审核 - 逻辑封装 - 涉及进销存
		DoReviewPicking(ctx context.Context, in *model.OrderPickingInput) (state uint, err error)
		// DoReviewShipping 发货审核  - 涉及快递单号处理
		DoReviewShipping(ctx context.Context, in *model.OrderShippingInput) (state uint, err error)
		// ReviewToState 审核订单到某个状态
		ReviewToState(ctx context.Context, orderId string, toOrderStateId uint) (flag bool, err error)
		CancelActivity(ctx context.Context, orderId string) (flag bool, err error)
		// 判断是否有活动条件限制
		IfActivity(ctx context.Context, orderId string) bool
		// GetOrderStatisticsInfo 根据用户id获取用户订单统计信息
		GetOrderStatisticsInfo(ctx context.Context, userId uint) (*model.OrderNumOutput, error)
		// GetOrderNum 订单数量
		GetOrderNum(ctx context.Context, in *model.OrderNumInput) int64
	}
	IOrderBase interface {
		// Get 读取订单
		Get(ctx context.Context, id any) (out *entity.OrderBase, err error)
		// Gets 读取多条订单
		Gets(ctx context.Context, id any) (list []*entity.OrderBase, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.OrderBaseListInput) (out []*entity.OrderBase, err error)
		// List 分页读取
		List(ctx context.Context, in *do.OrderBaseListInput) (out *do.OrderBaseListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.OrderBase) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.OrderBase) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IOrderInfo interface {
		// Get 读取订单
		Get(ctx context.Context, id any) (out *entity.OrderInfo, err error)
		// Gets 读取多条订单
		Gets(ctx context.Context, id any) (list []*entity.OrderInfo, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.OrderInfoListInput) (out []*entity.OrderInfo, err error)
		// GetList 订单搜索查询列表
		GetList(ctx context.Context, in *do.OrderInfoListInput) (out *model.OrderListOutput, err error)
		// List 分页读取
		List(ctx context.Context, in *do.OrderInfoListInput) (out *do.OrderInfoListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.OrderInfo) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.OrderInfo) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IDistributionOrder interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.DistributionOrderListInput) (out []*entity.DistributionOrder, err error)
		// List 分页读取
		List(ctx context.Context, in *do.DistributionOrderListInput) (out *do.DistributionOrderListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.DistributionOrder) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.DistributionOrder) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// InitDistributionUserOrder 删除多条记录模式
		InitDistributionUserOrder(ctx context.Context, distributionOrderVo *model.DistributionOrderVo, itemRows []*do.OrderItem) (distFlag bool, err error)
		SettleDistributionUserOrder(ctx context.Context, orderId string) (distFlag bool, err error)
	}
	IOrderItem interface {
		// Get 读取订单
		Get(ctx context.Context, id any) (out *entity.OrderItem, err error)
		// Gets 读取多条订单
		Gets(ctx context.Context, id any) (list []*entity.OrderItem, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.OrderItemListInput) (out []*entity.OrderItem, err error)
		// List 分页读取
		List(ctx context.Context, in *do.OrderItemListInput) (out *do.OrderItemListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.OrderItem) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.OrderItem) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// Save 保存
		Save(ctx context.Context, in *do.OrderItem) (affected int64, err error)
		// Saves 保存
		Saves(ctx context.Context, in []*do.OrderItem) (affected int64, err error)
	}
	IStockBill interface {
		// Get 读取订单
		Get(ctx context.Context, id any) (out *entity.StockBill, err error)
		// Gets 读取多条订单
		Gets(ctx context.Context, id any) (list []*entity.StockBill, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.StockBillListInput) (out []*entity.StockBill, err error)
		// FindDetail 查询数据
		FindDetail(ctx context.Context, in *do.StockBillListInput) (out []*model.StockBillVo, err error)
		// List 分页读取
		List(ctx context.Context, in *do.StockBillListInput) (out *do.StockBillListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.StockBill) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.StockBill) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IChainCode interface {
		// Get 读取兑换码
		Get(ctx context.Context, id any) (out *entity.ChainCode, err error)
		// Gets 读取多条兑换码
		Gets(ctx context.Context, id any) (list []*entity.ChainCode, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.ChainCodeListInput) (out []*entity.ChainCode, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ChainCodeListInput) (out *do.ChainCodeListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.ChainCode) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ChainCode) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IOrderDeliveryAddress interface {
		// Get 读取订单
		Get(ctx context.Context, id any) (out *entity.OrderDeliveryAddress, err error)
		// Gets 读取多条订单
		Gets(ctx context.Context, id any) (list []*entity.OrderDeliveryAddress, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.OrderDeliveryAddressListInput) (out []*entity.OrderDeliveryAddress, err error)
		// List 分页读取
		List(ctx context.Context, in *do.OrderDeliveryAddressListInput) (out *do.OrderDeliveryAddressListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.OrderDeliveryAddress) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.OrderDeliveryAddress) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// Save 报错
		Save(ctx context.Context, in *do.OrderDeliveryAddress) (affected int64, err error)
	}
	IOrderInvoice interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.OrderInvoiceListInput) (out []*entity.OrderInvoice, err error)
		// List 分页读取
		List(ctx context.Context, in *do.OrderInvoiceListInput) (out *do.OrderInvoiceListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.OrderInvoice) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.OrderInvoice) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// Remove 删除多条记录模式
		RemoveWhere(ctx context.Context, where *do.OrderInvoiceListInput) (affected int64, err error)
	}
	IOrderReturnReason interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.OrderReturnReasonListInput) (out []*entity.OrderReturnReason, err error)
		// List 分页读取
		List(ctx context.Context, in *do.OrderReturnReasonListInput) (out *do.OrderReturnReasonListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.OrderReturnReason) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.OrderReturnReason) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
)

var (
	localOrder                 IOrder
	localOrderBase             IOrderBase
	localOrderInfo             IOrderInfo
	localOrderLogistics        IOrderLogistics
	localOrderReturn           IOrderReturn
	localOrderReturnItem       IOrderReturnItem
	localOrderStateLog         IOrderStateLog
	localUserCart              IUserCart
	localDistributionOrder     IDistributionOrder
	localChainCode             IChainCode
	localOrderDeliveryAddress  IOrderDeliveryAddress
	localOrderInvoice          IOrderInvoice
	localOrderItem             IOrderItem
	localStockBill             IStockBill
	localDistributionOrderItem IDistributionOrderItem
	localOrderData             IOrderData
	localOrderReturnReason     IOrderReturnReason
)

func DistributionOrder() IDistributionOrder {
	if localDistributionOrder == nil {
		panic("implement not found for interface IDistributionOrder, forgot register?")
	}
	return localDistributionOrder
}

func RegisterDistributionOrder(i IDistributionOrder) {
	localDistributionOrder = i
}

func ChainCode() IChainCode {
	if localChainCode == nil {
		panic("implement not found for interface IChainCode, forgot register?")
	}
	return localChainCode
}

func RegisterChainCode(i IChainCode) {
	localChainCode = i
}

func OrderDeliveryAddress() IOrderDeliveryAddress {
	if localOrderDeliveryAddress == nil {
		panic("implement not found for interface IOrderDeliveryAddress, forgot register?")
	}
	return localOrderDeliveryAddress
}

func RegisterOrderDeliveryAddress(i IOrderDeliveryAddress) {
	localOrderDeliveryAddress = i
}

func OrderInvoice() IOrderInvoice {
	if localOrderInvoice == nil {
		panic("implement not found for interface IOrderInvoice, forgot register?")
	}
	return localOrderInvoice
}

func RegisterOrderInvoice(i IOrderInvoice) {
	localOrderInvoice = i
}

func OrderItem() IOrderItem {
	if localOrderItem == nil {
		panic("implement not found for interface IOrderItem, forgot register?")
	}
	return localOrderItem
}

func RegisterOrderItem(i IOrderItem) {
	localOrderItem = i
}

func StockBill() IStockBill {
	if localStockBill == nil {
		panic("implement not found for interface IStockBill, forgot register?")
	}
	return localStockBill
}

func RegisterStockBill(i IStockBill) {
	localStockBill = i
}

func DistributionOrderItem() IDistributionOrderItem {
	if localDistributionOrderItem == nil {
		panic("implement not found for interface IDistributionOrderItem, forgot register?")
	}
	return localDistributionOrderItem
}

func RegisterDistributionOrderItem(i IDistributionOrderItem) {
	localDistributionOrderItem = i
}

func OrderData() IOrderData {
	if localOrderData == nil {
		panic("implement not found for interface IOrderData, forgot register?")
	}
	return localOrderData
}

func RegisterOrderData(i IOrderData) {
	localOrderData = i
}

func UserCart() IUserCart {
	if localUserCart == nil {
		panic("implement not found for interface IUserCart, forgot register?")
	}
	return localUserCart
}

func RegisterUserCart(i IUserCart) {
	localUserCart = i
}

func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

func RegisterOrder(i IOrder) {
	localOrder = i
}

func OrderBase() IOrderBase {
	if localOrderBase == nil {
		panic("implement not found for interface IOrderBase, forgot register?")
	}
	return localOrderBase
}

func RegisterOrderBase(i IOrderBase) {
	localOrderBase = i
}

func OrderInfo() IOrderInfo {
	if localOrderInfo == nil {
		panic("implement not found for interface IOrderInfo, forgot register?")
	}
	return localOrderInfo
}

func RegisterOrderInfo(i IOrderInfo) {
	localOrderInfo = i
}

func OrderLogistics() IOrderLogistics {
	if localOrderLogistics == nil {
		panic("implement not found for interface IOrderLogistics, forgot register?")
	}
	return localOrderLogistics
}

func RegisterOrderLogistics(i IOrderLogistics) {
	localOrderLogistics = i
}

func OrderReturn() IOrderReturn {
	if localOrderReturn == nil {
		panic("implement not found for interface IOrderReturn, forgot register?")
	}
	return localOrderReturn
}

func RegisterOrderReturn(i IOrderReturn) {
	localOrderReturn = i
}

func OrderReturnItem() IOrderReturnItem {
	if localOrderReturnItem == nil {
		panic("implement not found for interface IOrderReturnItem, forgot register?")
	}
	return localOrderReturnItem
}

func RegisterOrderReturnItem(i IOrderReturnItem) {
	localOrderReturnItem = i
}

func OrderStateLog() IOrderStateLog {
	if localOrderStateLog == nil {
		panic("implement not found for interface IOrderStateLog, forgot register?")
	}
	return localOrderStateLog
}

func RegisterOrderStateLog(i IOrderStateLog) {
	localOrderStateLog = i
}
func OrderReturnReason() IOrderReturnReason {
	if localOrderReturnReason == nil {
		panic("implement not found for interface IOrderReturnReason, forgot register?")
	}
	return localOrderReturnReason
}

func RegisterOrderReturnReason(i IOrderReturnReason) {
	localOrderReturnReason = i
}
