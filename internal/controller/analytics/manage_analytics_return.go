package analytics

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/analytics"
	"golershop.cn/internal/model"
	"golershop.cn/internal/service"
)

var (
	AnalyticsReturn = cAnalyticsReturn{}
)

type cAnalyticsReturn struct{}

// List 退单额
func (c *cAnalyticsReturn) GetReturnAmount(ctx context.Context, req *analytics.ReturnAmountReq) (res analytics.ReturnAmountRes, err error) {
	input := &model.AnalyticsReturnInput{}
	gconv.Scan(req, &input)

	result, error := service.AnalyticsReturn().GetReturnAmount(ctx, input)
	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)
	return
}

// List 退单量
func (c *cAnalyticsReturn) GetReturnNum(ctx context.Context, req *analytics.ReturnNumReq) (res analytics.ReturnNumRes, err error) {
	input := &model.AnalyticsReturnInput{}
	gconv.Scan(req, &input)

	result, error := service.AnalyticsReturn().GetReturnNum(ctx, input)
	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)
	return
}

// List 退单金额对比图
func (c *cAnalyticsReturn) GetReturnAmountTimeline(ctx context.Context, req *analytics.ReturnAmountTimelineReq) (res analytics.ReturnAmountTimelineRes, err error) {
	input := &model.TimelineInput{}
	gconv.Scan(req, &input)

	result, error := service.AnalyticsReturn().GetReturnAmountTimeline(ctx, input)
	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)
	return
}

// List 退单数量统计
func (c *cAnalyticsReturn) GetReturnNumTimeline(ctx context.Context, req *analytics.ReturnNumTimelineReq) (res analytics.ReturnNumTimelineRes, err error) {
	input := &model.TimelineInput{}
	gconv.Scan(req, &input)

	result, error := service.AnalyticsReturn().GetReturnNumTimeline(ctx, input)
	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)
	return
}

// List 退单商品销量统计
func (c *cAnalyticsReturn) GetReturnItemNumTimeline(ctx context.Context, req *analytics.ReturnItemNumTimelineReq) (res analytics.ReturnItemNumTimelineRes, err error) {
	input := &model.OrderItemNumTimelineInput{}
	gconv.Scan(req, &input)

	result, error := service.AnalyticsReturn().GetReturnItemNumTimeLine(ctx, input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)
	return
}

// List 退单商品销量统计
func (c *cAnalyticsReturn) ListReturnItemNum(ctx context.Context, req *analytics.ReturnItemNumListReq) (res analytics.ReturnItemNumListRes, err error) {
	input := &model.OrderItemNumTimelineInput{}
	gconv.Scan(req, &input)

	result, error := service.AnalyticsReturn().ListReturnItemNum(ctx, input)
	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)
	return
}

// List 商品访客数
func (c *cAnalyticsReturn) GetReturnItemNum(ctx context.Context, req *analytics.ReturnItemNumReq) (res analytics.ReturnItemNumRes, err error) {
	input := &model.OrderItemNumTimelineInput{}
	gconv.Scan(req, &input)

	result, error := service.AnalyticsReturn().GetReturnItemNum(ctx, input)
	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)
	return
}
