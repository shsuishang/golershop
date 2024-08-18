package trade

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/trade"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	OrderBase = cOrderBase{}
)

type cOrderBase struct{}

// =================== 管理端使用 =========================
// Detail 订单详情
func (c *cOrderBase) Detail(ctx context.Context, req *trade.OrderDetailReq) (res *trade.OrderDetailRes, err error) {

	var result, error = service.Order().Detail(ctx, req.OrderId)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// List
func (c *cOrderBase) List(ctx context.Context, req *trade.OrderListReq) (res *trade.OrderListRes, err error) {
	req.Sidx = dao.OrderInfo.Columns().CreateTime
	req.Sort = ml.ORDER_BY_DESC

	input := do.OrderInfoListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.Order().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增订单
func (c *cOrderBase) Add(ctx context.Context, req *trade.OrderBaseAddReq) (res *trade.OrderBaseEditRes, err error) {

	input := &model.CheckoutInput{}
	gconv.Scan(req, input)

	// 如果商品数据为空，抛出业务异常
	if g.IsEmpty(req.ProductItems) {
		return nil, gerror.New("商品数据为空！")
	}

	var checkoutItemVos []*model.CheckoutItemVo
	err = json.Unmarshal([]byte(req.ProductItems), &checkoutItemVos)
	if err != nil || len(checkoutItemVos) == 0 {
		return nil, gerror.New("商品数据错误！")
	}

	// 设置商品项
	input.Items = checkoutItemVos

	// 调用服务替换添加
	output, err := service.Order().Add(ctx, input)
	if err != nil {
		return nil, err
	}

	// 如果订单 ID 不为空，则表示成功
	if len(output.OrderIds) > 0 {
		res.OrderId = output.OrderIds
		return res, nil
	}

	// 返回失败
	return nil, gerror.New("下单失败")
}

// Edit 编辑订单
func (c *cOrderBase) Edit(ctx context.Context, req *trade.OrderBaseEditReq) (res *trade.OrderBaseEditRes, err error) {

	input := do.OrderBase{}
	gconv.Scan(req, &input)

	var result, error = service.OrderBase().Edit(ctx, &input)
	//var result, error = service.OrderBase().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &trade.OrderBaseEditRes{
		OrderId: result,
	}

	return
}

// EditState 编辑订单
func (c *cOrderBase) EditState(ctx context.Context, req *trade.OrderBaseEditStateReq) (res *trade.OrderBaseEditStateRes, err error) {
	input := do.OrderBase{}
	gconv.Scan(req, &input)

	var result, error = service.OrderBase().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &trade.OrderBaseEditStateRes{
		OrderId: result,
	}

	return
}

// Remove 删除订单
func (c *cOrderBase) Remove(ctx context.Context, req *trade.OrderBaseRemoveReq) (res *trade.OrderBaseRemoveRes, err error) {

	var _, error = service.OrderBase().Remove(ctx, req.OrderId)

	if error != nil {
		err = error
	}

	res = &trade.OrderBaseRemoveRes{}

	return
}

// Cancel 取消订单
func (c *cOrderBase) Cancel(ctx context.Context, req *trade.OrderCancelReq) (res *trade.OrderCancelRes, err error) {
	res = &trade.OrderCancelRes{}

	for _, orderId := range req.OrderId {
		var _, error = service.Order().Cancel(ctx, orderId, req.OrderCancelReason)

		if error != nil {
		} else {
			res.OrderId = append(res.OrderId, orderId)
		}
	}

	if len(res.OrderId) == 0 {
		err = errors.New("未更改到符合条件的订单！")
	} else {

	}

	return
}

// Review 审核订单
func (c *cOrderBase) Review(ctx context.Context, req *trade.OrderReviewReq) (res *trade.OrderReviewRes, err error) {
	res = &trade.OrderReviewRes{}

	for _, orderId := range req.OrderId {
		var _, error = service.Order().Review(ctx, orderId, req.OrderReviewReason)

		if error != nil {
		} else {
			res.OrderId = append(res.OrderId, orderId)
		}
	}

	if len(res.OrderId) == 0 {
		err = errors.New("未更改到符合条件的订单！")
	} else {

	}

	return
}

// Finance 财务审核
func (c *cOrderBase) Finance(ctx context.Context, req *trade.OrderFinanceReq) (res *trade.OrderFinanceRes, err error) {
	res = &trade.OrderFinanceRes{}

	for _, orderId := range req.OrderId {
		var _, error = service.Order().Finance(ctx, orderId, req.OrderFinanceReason)

		if error != nil {
		} else {
			res.OrderId = append(res.OrderId, orderId)
		}
	}

	if len(res.OrderId) == 0 {
		err = errors.New("未更改到符合条件的订单！")
	} else {

	}

	return
}

// Picking 出库审核
func (c *cOrderBase) Picking(ctx context.Context, req *trade.OrderPickingReq) (res *trade.OrderPickingRes, err error) {
	res = &trade.OrderPickingRes{}

	for _, orderId := range req.OrderId {
		input := &model.OrderPickingInput{}
		gconv.Scan(req, input)
		input.OrderId = orderId

		var _, error = service.Order().Picking(ctx, input)

		if error != nil {
		} else {
			res.OrderId = append(res.OrderId, orderId)
		}
	}

	if len(res.OrderId) == 0 {
		err = errors.New("未更改到符合条件的订单！")
	} else {

	}

	return
}

// Shipping 发货审核
func (c *cOrderBase) Shipping(ctx context.Context, req *trade.OrderShippingReq) (res *trade.OrderShippingRes, err error) {
	res = &trade.OrderShippingRes{}

	for _, orderId := range req.OrderId {
		input := &model.OrderShippingInput{}
		gconv.Scan(req, input)
		input.OrderId = orderId

		var _, error = service.Order().Shipping(ctx, input)

		if error != nil {
		} else {
			res.OrderId = append(res.OrderId, orderId)
		}
	}

	if len(res.OrderId) == 0 {
		err = errors.New("未更改到符合条件的订单！")
	} else {

	}

	return
}

// Receive 收货审核
func (c *cOrderBase) Receive(ctx context.Context, req *trade.OrderReceiveReq) (res *trade.OrderReceiveRes, err error) {
	res = &trade.OrderReceiveRes{}

	for _, orderId := range req.OrderId {
		var _, error = service.Order().Receive(ctx, orderId, "")

		if error != nil {
		} else {
			res.OrderId = append(res.OrderId, orderId)
		}
	}

	if len(res.OrderId) == 0 {
		err = errors.New("未更改到符合条件的订单！")
	} else {

	}

	return
}

// ListStateLog 操作日志
func (c *cOrderBase) ListStateLog(ctx context.Context, req *trade.OrderStateLogListReq) (res *trade.OrderStateLogListRes, err error) {
	input := do.OrderStateLogListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.OrderStateLog().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}
