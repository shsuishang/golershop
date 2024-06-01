package trade

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/trade"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	Return = cReturn{}
)

type cReturn struct{}

// List 退款退货列表
func (c *cReturn) List(ctx context.Context, req *trade.ReturnListReq) (res *trade.ReturnListRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	if g.IsEmpty(userId) {
		return nil, gerror.New(g.I18n().Translate(ctx, "请先登录！"))
	}

	req.BuyerUserId = userId

	input := &do.OrderReturnListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	page, err := service.OrderReturn().GetList(ctx, input)
	if err != nil {
		return nil, err
	}

	res = &trade.ReturnListRes{}
	gconv.Struct(page, res)
	return res, nil
}

// Get 读取退款退货
func (c *cReturn) Get(ctx context.Context, req *trade.ReturnDetailReq) (res *trade.ReturnDetailRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	if g.IsEmpty(userId) {
		return nil, gerror.New(g.I18n().Translate(ctx, "请先登录！"))
	}

	orderReturnRes, err := service.OrderReturn().GetReturn(ctx, req.ReturnId)
	if err != nil {
		return nil, err
	}

	res = &trade.ReturnDetailRes{}
	gconv.Struct(orderReturnRes, res)

	return res, nil
}

// Edit 确认退货物流单号
func (c *cReturn) Edit(ctx context.Context, req *trade.ReturnEditReq) (res *trade.ReturnEditRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	if g.IsEmpty(userId) {
		return nil, gerror.New(g.I18n().Translate(ctx, "请先登录！"))
	}

	orderReturn := new(do.OrderReturn)
	gconv.Struct(req, orderReturn)

	returnOrder, err := service.OrderReturn().Get(ctx, req.ReturnId)
	if err != nil {
		return nil, err
	}

	if returnOrder.BuyerUserId == userId {
		success, err := service.OrderReturn().EditReturn(ctx, orderReturn)
		if err != nil || !success {
			return nil, err
		}
		return &trade.ReturnEditRes{ReturnId: orderReturn.ReturnId}, nil
	}

	return nil, gerror.New("编辑失败")
}

// Cancel 取消退款订单
func (c *cReturn) Cancel(ctx context.Context, req *trade.ReturnCancelReq) (res *trade.ReturnCancelRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	if g.IsEmpty(userId) {
		return nil, gerror.New(g.I18n().Translate(ctx, "请先登录！"))
	}

	returnOrder, err := service.OrderReturn().Get(ctx, req.ReturnId)
	if err != nil {
		return nil, err
	}
	if returnOrder.BuyerUserId == userId {
		success, err := service.OrderReturn().Cancel(ctx, req.ReturnId, userId)
		if err != nil || !success {
			return nil, err
		}

		return &trade.ReturnCancelRes{}, nil
	}

	return nil, gerror.New("取消失败")
}

// ReturnItem 订单item详情,列出订单的item，及退款详情
func (c *cReturn) ReturnItem(ctx context.Context, req *trade.ReturnItemReq) (res *trade.ReturnItemRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	if g.IsEmpty(userId) {
		return nil, gerror.New(g.I18n().Translate(ctx, "请先登录！"))
	}

	orderReturnItemVo, err := service.OrderReturn().ReturnItem(ctx, req.OrderId, req.OrderItemId, userId)
	if err != nil {
		return nil, err
	}

	res = &trade.ReturnItemRes{}
	gconv.Struct(orderReturnItemVo, res)

	return res, nil
}

// Add 添加退款退货
func (c *cReturn) Add(ctx context.Context, req *trade.ReturnAddReq) (res *trade.ReturnAddRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	if g.IsEmpty(userId) {
		return nil, gerror.New(g.I18n().Translate(ctx, "请先登录！"))
	}

	orderReturnInput := new(model.OrderReturnInput)
	gconv.Struct(req, orderReturnInput)

	orderReturnInput.UserId = userId
	orderReturnInput.ReturnAllFlag = false

	returnItemInputVo := new(model.OrderReturnItemInputVo)
	returnItemInputVo.OrderItemId = req.OrderItemId
	returnItemInputVo.ReturnRefundAmount = req.ReturnRefundAmount
	returnItemInputVo.ReturnItemNum = req.ReturnItemNum
	orderReturnInput.ReturnItems = append(orderReturnInput.ReturnItems, returnItemInputVo)

	returnId, err := service.OrderReturn().AddItem(ctx, orderReturnInput)
	if err != nil {
		return nil, err
	}

	res = &trade.ReturnAddRes{}
	res.ReturnId = returnId

	return res, nil
}
