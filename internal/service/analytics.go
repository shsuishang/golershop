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
)

type (
	IAnalyticsOrder interface {
		// 获取后台仪表盘中部柱形数据，包含[订单数据，用户数据，商品数据，销售额数据]
		GetDashboardTimeLine(ctx context.Context, input *model.TimelineInput) (model.DashBoardTimelineOutput, error)
		GetSaleOrderAmount(ctx context.Context, input *model.TimelineInput) (out []*model.TimelineOutput, err error)
		GetOrderCustomerNumTimeline(ctx context.Context, input *model.TimelineInput) (out []*model.TimelineOutput, err error)
		// 计算本月订单量，今日订单量，昨日订单量，以及日环比
		GetOrderNumToday(ctx context.Context) (topRes model.DashboardTopVo, err error)
		GetOrderNum(ctx context.Context, input *model.AnalyticsOrderInput) (topRes model.AnalyticsNumOutput, err error)
		// 获取订单金额
		GetOrderAmount(ctx context.Context, in *model.AnalyticsOrderInput) (*model.AnalyticsNumOutput, error)
		// 获取订单时间线
		GetOrderNumTimeline(ctx context.Context, in *model.TimelineInput) (out []*model.TimelineOutput, err error)
		// 获取订单商品数量时间线
		GetOrderItemNumTimeLine(ctx context.Context, in *model.OrderItemNumTimelineInput) (out []*model.TimelineOutput, err error)
		// 获取订单商品数量
		GetOrderItemNum(ctx context.Context, in *model.OrderItemNumTimelineInput) (*model.AnalyticsNumOutput, error)
		// 列出订单商品数量
		ListOrderItemNum(ctx context.Context, in *model.OrderItemNumTimelineInput) (out []*model.AnalyticsOrderItemNumOutput, err error)
	}
	IAnalyticsProduct interface {
		// 获取商品数量
		GetProductNum(ctx context.Context, input *model.AnalyticsProductInput) (*model.AnalyticsNumOutput, error)
	}
	IAnalyticsReturn interface {
		// GetReturnAmountTimeline 获取退单金额时间线
		GetReturnAmountTimeline(ctx context.Context, input *model.TimelineInput) (out []*model.TimelineOutput, err error)
		// GetReturnNum 获取退单数量
		GetReturnNum(ctx context.Context, input *model.AnalyticsReturnInput) (*model.AnalyticsNumOutput, error)
		// GetReturnAmount 获取退单金额
		GetReturnAmount(ctx context.Context, input *model.AnalyticsReturnInput) (*model.AnalyticsNumOutput, error)
		// GetReturnNumTimeline 获取退单数量时间线
		GetReturnNumTimeline(ctx context.Context, input *model.TimelineInput) (out []*model.TimelineOutput, err error)
		// GetReturnItemNumTimeLine 获取退单商品数量时间线
		GetReturnItemNumTimeLine(ctx context.Context, input *model.OrderItemNumTimelineInput) (out []*model.TimelineOutput, err error)
		// GetReturnItemNum 获取退单商品数量
		GetReturnItemNum(ctx context.Context, input *model.OrderItemNumTimelineInput) (*model.AnalyticsNumOutput, error)
		// ListReturnItemNum 获取退单商品数量列表
		ListReturnItemNum(ctx context.Context, input *model.OrderItemNumTimelineInput) (out []*model.AnalyticsOrderItemNumOutput, err error)
		// GetReturnCustomerNumTimeline 获取退单客户统计
		GetReturnCustomerNumTimeline(ctx context.Context, input *model.TimelineInput) (out []*model.TimelineOutput, err error)
	}
	IAnalyticsSys interface {
		// 计算本月访问量，今日访问量，昨日访问量，以及日环比
		GetVisitor(ctx context.Context) *model.DashboardTopVo
		// 获取访问访问者数量
		GetAccessVisitorNum(ctx context.Context, input *model.AccessItemTimelineInput) *model.AnalyticsNumOutput
		// 获取访问数量
		GetAccessNum(ctx context.Context, input *model.AccessItemTimelineInput) *model.AnalyticsNumOutput
		// 获取访问项目时间线
		GetAccessItemTimeLine(ctx context.Context, input *model.AccessItemTimelineInput) (out []*model.TimelineOutput, err error)
		// 获取访问项目数量
		GetAccessItemNum(ctx context.Context, input *model.AccessItemTimelineInput) *model.AnalyticsNumOutput
		// 获取访问项目用户时间线
		GetAccessItemUserTimeLine(ctx context.Context, input *model.AccessItemTimelineInput) (out []*model.TimelineOutput, err error)
		// 获取访问物品用户数量
		GetAccessItemUserNum(ctx context.Context, input *model.AccessItemTimelineInput) *model.AnalyticsNumOutput
		// 获取访问者时间线
		GetAccessVisitorTimeLine(ctx context.Context, input *model.TimelineInput) (out []*model.TimelineOutput, err error)
		// 列出访问物品
		ListAccessItem(ctx context.Context, timelineInput *model.AccessItemTimelineInput) (out []*model.AnalyticsAccessItemOutput, err error)
	}
	IAnalyticsTrade interface {
		// TradeAmount 交易总额
		SalesAmount(ctx context.Context, start int64, end int64, buyerId int64) (res interface{}, err error)
	}
	IAnalyticsUser interface {
		// 获取注册用户数量
		GetRegUser(ctx context.Context) *model.DashboardTopVo
		// 获取用户时间线
		GetUserTimeLine(ctx context.Context, input *model.TimelineInput) (out []*model.TimelineOutput, err error)
		// 获取用户数量
		GetUserNum(ctx context.Context, input *model.TimelineInput) (out *model.AnalyticsNumOutput, err error)
	}
)

var (
	localAnalyticsUser    IAnalyticsUser
	localAnalyticsOrder   IAnalyticsOrder
	localAnalyticsProduct IAnalyticsProduct
	localAnalyticsReturn  IAnalyticsReturn
	localAnalyticsSys     IAnalyticsSys
	localAnalyticsTrade   IAnalyticsTrade
)

func AnalyticsTrade() IAnalyticsTrade {
	if localAnalyticsTrade == nil {
		panic("implement not found for interface IAnalyticsTrade, forgot register?")
	}
	return localAnalyticsTrade
}

func RegisterAnalyticsTrade(i IAnalyticsTrade) {
	localAnalyticsTrade = i
}

func AnalyticsUser() IAnalyticsUser {
	if localAnalyticsUser == nil {
		panic("implement not found for interface IAnalyticsUser, forgot register?")
	}
	return localAnalyticsUser
}

func RegisterAnalyticsUser(i IAnalyticsUser) {
	localAnalyticsUser = i
}

func AnalyticsOrder() IAnalyticsOrder {
	if localAnalyticsOrder == nil {
		panic("implement not found for interface IAnalyticsOrder, forgot register?")
	}
	return localAnalyticsOrder
}

func RegisterAnalyticsOrder(i IAnalyticsOrder) {
	localAnalyticsOrder = i
}

func AnalyticsProduct() IAnalyticsProduct {
	if localAnalyticsProduct == nil {
		panic("implement not found for interface IAnalyticsProduct, forgot register?")
	}
	return localAnalyticsProduct
}

func RegisterAnalyticsProduct(i IAnalyticsProduct) {
	localAnalyticsProduct = i
}

func AnalyticsReturn() IAnalyticsReturn {
	if localAnalyticsReturn == nil {
		panic("implement not found for interface IAnalyticsReturn, forgot register?")
	}
	return localAnalyticsReturn
}

func RegisterAnalyticsReturn(i IAnalyticsReturn) {
	localAnalyticsReturn = i
}

func AnalyticsSys() IAnalyticsSys {
	if localAnalyticsSys == nil {
		panic("implement not found for interface IAnalyticsSys, forgot register?")
	}
	return localAnalyticsSys
}

func RegisterAnalyticsSys(i IAnalyticsSys) {
	localAnalyticsSys = i
}
