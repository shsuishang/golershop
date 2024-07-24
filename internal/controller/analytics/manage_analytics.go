package analytics

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/analytics"
	"golershop.cn/internal/model"
	"golershop.cn/internal/service"
	"golershop.cn/utility/mtime"
)

var (
	Analytics = cAnalytics{}
)

type cAnalytics struct{}

func (c *cAnalytics) GetSalesAmount(ctx context.Context, req *analytics.SalesAmountReq) (res analytics.SalesAmountRes, err error) {

	//当天
	startTime, endTime := mtime.Today()
	amount, err := service.AnalyticsTrade().SalesAmount(ctx, startTime, endTime, 0)

	if err != nil {
		return
	}

	res.Today = amount

	//昨天
	startTime, endTime = mtime.Yestoday()
	amount, err = service.AnalyticsTrade().SalesAmount(ctx, startTime, endTime, 0)

	if err != nil {
		return
	}

	res.Yestoday = amount

	//日环比
	if res.Yestoday != 0 {
		res.Daym2m = res.Today.(int64)/res.Yestoday.(int64) - 1
	} else {
		res.Daym2m = nil
	}

	//本月
	startTime, endTime = mtime.Month()
	amount, err = service.AnalyticsTrade().SalesAmount(ctx, startTime, endTime, 0)

	if err != nil {
		return
	}

	res.Month = amount

	return
}

// 获取订单量
func (c *cAnalytics) GetOrderAmount(ctx context.Context, req *analytics.OrderAmountReq) (res analytics.OrderAmountRes, err error) {
	timelineInput := new(model.AnalyticsOrderInput)
	gconv.Struct(req, timelineInput)
	timelineRes, err := service.AnalyticsOrder().GetOrderAmount(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(timelineRes, &res)

	return
}

// 获取用户访问量
func (c *cAnalytics) GetVisitor(ctx context.Context, req *analytics.VisitorReq) (res analytics.VisitorRes, err error) {
	topRes := service.AnalyticsSys().GetVisitor(ctx)

	gconv.Struct(topRes, &res)

	return
}

// 获取订单量
func (c *cAnalytics) GetOrderNumToday(ctx context.Context, req *analytics.OrderNumTodayReq) (res analytics.OrderNumTodayRes, err error) {
	timelineInput := new(model.AnalyticsOrderInput)
	gconv.Struct(req, timelineInput)

	topRes, err := service.AnalyticsOrder().GetOrderNumToday(ctx)
	if err != nil {
		return
	}

	gconv.Struct(topRes, &res)

	return
}

// 获取订单量
func (c *cAnalytics) GetOrderNum(ctx context.Context, req *analytics.OrderNumReq) (res analytics.OrderNumRes, err error) {
	timelineInput := new(model.AnalyticsOrderInput)
	gconv.Struct(req, timelineInput)
	timelineRes, err := service.AnalyticsOrder().GetOrderNum(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(timelineRes, &res)

	return
}

// 获取新增用户
func (c *cAnalytics) GetRegUser(ctx context.Context, req *analytics.RegUserReq) (res analytics.RegUserRes, err error) {
	topRes := service.AnalyticsUser().GetRegUser(ctx)

	gconv.Struct(topRes, &res)

	return
}

// 商品数量
func (c *cAnalytics) GetProductNum(ctx context.Context, req *analytics.AnalyticsProductReq) (res analytics.AnalyticsProductRes, err error) {
	timelineInput := new(model.AnalyticsProductInput)
	gconv.Struct(req, timelineInput)
	topRes, err := service.AnalyticsProduct().GetProductNum(ctx, timelineInput)
	if err != nil {
		return
	}
	gconv.Struct(topRes, &res)

	return
}

// 新增用户
func (c *cAnalytics) GetUserTimeLine(ctx context.Context, req *analytics.UserTimeLineReq) (res analytics.UserTimeLineRes, err error) {
	timelineInput := new(model.TimelineInput)
	gconv.Struct(req, timelineInput)
	timelineRes, err := service.AnalyticsUser().GetUserTimeLine(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(timelineRes, &res)

	if res == nil {
		res = make([]*model.TimelineOutput, 0)
	}

	return
}

// 获取新增用户
func (c *cAnalytics) GetUserNum(ctx context.Context, req *analytics.UserNumReq) (res analytics.UserNumRes, err error) {
	timelineInput := new(model.TimelineInput)
	gconv.Struct(req, timelineInput)
	timelineRes, err := service.AnalyticsUser().GetUserNum(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(timelineRes, &res)

	return
}

// 仪表板看板柱形图数据
func (c *cAnalytics) GetDashboardTimeLine(ctx context.Context, req *analytics.DashboardTimeLineReq) (res analytics.DashboardTimeLineRes, err error) {
	timelineInput := new(model.TimelineInput)
	gconv.Struct(req, timelineInput)
	dashBoardTimeLineRes, err := service.AnalyticsOrder().GetDashboardTimeLine(ctx, timelineInput)
	if err != nil {
		return
	}
	gconv.Struct(dashBoardTimeLineRes, &res)

	return
}

// 购买商品顾客数量统计
func (c *cAnalytics) GetOrderCustomerNumTimeline(ctx context.Context, req *analytics.OrderCustomerNumTimelineReq) (res analytics.OrderCustomerNumTimelineRes, err error) {
	timelineInput := new(model.TimelineInput)
	gconv.Struct(req, timelineInput)
	timelineRes, err := service.AnalyticsOrder().GetOrderCustomerNumTimeline(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(timelineRes, &res)

	return
}

// 订单销售金额对比图
func (c *cAnalytics) GetSaleOrderAmount(ctx context.Context, req *analytics.SaleOrderAmountReq) (res analytics.SaleOrderAmountRes, err error) {
	timelineInput := new(model.TimelineInput)
	gconv.Struct(req, timelineInput)
	amountRes, err := service.AnalyticsOrder().GetSaleOrderAmount(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(amountRes, &res)

	if res == nil {
		res = make([]*model.TimelineOutput, 0)
	}

	return
}

// 订单数量统计
func (c *cAnalytics) GetOrderNumTimeline(ctx context.Context, req *analytics.OrderNumTimelineReq) (res analytics.OrderNumTimelineRes, err error) {
	timelineInput := new(model.TimelineInput)
	gconv.Struct(req, timelineInput)
	timelineRes, err := service.AnalyticsOrder().GetOrderNumTimeline(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(timelineRes, &res)

	return
}

// 订单商品销量统计
func (c *cAnalytics) GetOrderItemNumTimeLine(ctx context.Context, req *analytics.OrderItemNumTimeLineReq) (res analytics.OrderItemNumTimeLineRes, err error) {
	timelineInput := new(model.OrderItemNumTimelineInput)
	gconv.Struct(req, timelineInput)
	timelineRes, err := service.AnalyticsOrder().GetOrderItemNumTimeLine(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(timelineRes, &res)

	return
}

// 订单商品销量统计
func (c *cAnalytics) ListOrderItemNum(ctx context.Context, req *analytics.OrderItemNumListReq) (res analytics.OrderItemNumListRes, err error) {
	timelineInput := new(model.OrderItemNumTimelineInput)
	gconv.Struct(req, timelineInput)
	timelineRes, err := service.AnalyticsOrder().ListOrderItemNum(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(timelineRes, &res)

	return
}

// 访问商品统计
func (c *cAnalytics) GetAccessItemTimeLine(ctx context.Context, req *analytics.AccessItemTimelineReq) (res analytics.AccessItemTimelineRes, err error) {
	timelineInput := new(model.AccessItemTimelineInput)
	gconv.Struct(req, timelineInput)
	timelineRes, err := service.AnalyticsSys().GetAccessItemTimeLine(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(timelineRes, &res)

	return
}

// 用户访问商品统计
func (c *cAnalytics) GetAccessItemUserTimeLine(ctx context.Context, req *analytics.AccessItemUserTimeLineReq) (res analytics.AccessItemUserTimeLineRes, err error) {
	timelineInput := new(model.AccessItemTimelineInput)
	gconv.Struct(req, timelineInput)
	timelineRes, err := service.AnalyticsSys().GetAccessItemUserTimeLine(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(timelineRes, &res)

	return
}

// 访客数
func (c *cAnalytics) GetAccessVisitorTimeLine(ctx context.Context, req *analytics.AccessVisitorTimeLineReq) (res analytics.AccessVisitorTimeLineRes, err error) {
	timelineInput := new(model.TimelineInput)
	gconv.Struct(req, timelineInput)
	timelineRes, err := service.AnalyticsSys().GetAccessVisitorTimeLine(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(timelineRes, &res)

	return
}

// 访客数
func (c *cAnalytics) GetAccessVisitorNum(ctx context.Context, req *analytics.AccessVisitorNumReq) (res analytics.AccessVisitorNumRes, err error) {
	timelineInput := &model.AccessItemTimelineInput{}
	gconv.Struct(req, timelineInput)
	timelineRes := service.AnalyticsSys().GetAccessVisitorNum(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(timelineRes, &res)

	return
}

// 浏览量
func (c *cAnalytics) GetAccessNum(ctx context.Context, req *analytics.AccessNumReq) (res analytics.AccessNumRes, err error) {
	timelineInput := &model.AccessItemTimelineInput{}
	gconv.Struct(req, timelineInput)
	timelineRes := service.AnalyticsSys().GetAccessNum(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(timelineRes, &res)

	return
}

// 商品浏览量
func (c *cAnalytics) GetAccessItemNum(ctx context.Context, req *analytics.AccessItemNumReq) (res analytics.AccessItemNumRes, err error) {
	timelineInput := &model.AccessItemTimelineInput{}
	gconv.Struct(req, timelineInput)
	timelineRes := service.AnalyticsSys().GetAccessItemNum(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(timelineRes, &res)

	return
}

// 商品浏览排行
func (c *cAnalytics) ListAccessItem(ctx context.Context, req *analytics.AccessItemListReq) (res analytics.AccessItemListRes, err error) {
	timelineInput := &model.AccessItemTimelineInput{}
	gconv.Struct(req, timelineInput)
	timelineRes, err := service.AnalyticsSys().ListAccessItem(ctx, timelineInput)
	if err != nil {
		return nil, err
	}

	gconv.Struct(timelineRes, &res)

	return
}

// 商品访客数
func (c *cAnalytics) GetAccessItemUserNum(ctx context.Context, req *analytics.AccessItemUserNumReq) (res analytics.AccessItemUserNumRes, err error) {
	timelineInput := &model.AccessItemTimelineInput{}
	gconv.Struct(req, timelineInput)
	timelineRes := service.AnalyticsSys().GetAccessItemUserNum(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(timelineRes, &res)

	return
}

// 商品访客数
func (c *cAnalytics) GetOrderItemNum(ctx context.Context, req *analytics.OrderItemNumReq) (res analytics.OrderItemNumRes, err error) {
	timelineInput := &model.OrderItemNumTimelineInput{}
	gconv.Struct(req, timelineInput)
	timelineRes, err := service.AnalyticsOrder().GetOrderItemNum(ctx, timelineInput)
	if err != nil {
		return
	}

	gconv.Struct(timelineRes, &res)

	return
}
