package trade

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/trade"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
)

var (
	OrderReturn = cOrderReturn{}
)

type cOrderReturn struct{}

// =================== 管理端使用 =========================
// Detail 订单详情
func (c *cOrderReturn) Detail(ctx context.Context, req *trade.OrderReturnDetailReq) (res *trade.OrderReturnDetailRes, err error) {

	var result, error = service.Order().Detail(ctx, req.ReturnId)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// List
func (c *cOrderReturn) List(ctx context.Context, req *trade.OrderReturnListReq) (res *trade.OrderReturnListRes, err error) {
	input := do.OrderReturnListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.OrderReturn().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// EditState 编辑订单
func (c *cOrderReturn) EditState(ctx context.Context, req *trade.OrderReturnEditStateReq) (res *trade.OrderReturnEditStateRes, err error) {
	input := do.OrderReturn{}
	gconv.Scan(req, &input)

	var result, error = service.OrderReturn().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &trade.OrderReturnEditStateRes{
		OrderId: result,
	}

	return
}

// Cancel 取消订单
func (c *cOrderReturn) Cancel(ctx context.Context, req *trade.OrderReturnCancelReq) (res *trade.OrderReturnCancelRes, err error) {
	res = &trade.OrderReturnCancelRes{}

	return
}

// Review 审核订单
func (c *cOrderReturn) Review(ctx context.Context, req *trade.OrderReturnReviewReq) (res *trade.OrderReturnReviewRes, err error) {
	// 将请求参数转换为OrderReturn实体
	input := &do.OrderReturn{}
	gconv.Scan(req, input)

	// 获取当前登录用户
	user := service.BizCtx().GetUser(ctx)

	if user == nil {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "需要登录")
	}

	// 设置店铺ID
	input.StoreId = user.StoreId

	// 调用审核服务
	success, err := service.OrderReturn().Review(ctx, input, req.ReceivingAddress)

	if err != nil {
		return nil, err
	}

	if success {
		return nil, nil
	}

	return nil, gerror.New("审核失败")

}

// GetByReturnId 获取退款退货详情
func (c *cOrderReturn) GetByReturnId(ctx context.Context, req *trade.GetByReturnIdReq) (res *trade.GetByReturnIdRes, err error) {
	// 根据returnId获取退款退货详情
	orderReturnVo, err := service.OrderReturn().GetByReturnId(ctx, req.ReturnId)
	if err != nil {
		return nil, err
	}
	res = &trade.GetByReturnIdRes{}
	// 将结果转换为响应结构
	gconv.Scan(orderReturnVo, &res)

	return
}

// Receive 退货单审核-确认收货
func (c *cOrderReturn) Receive(ctx context.Context, req *trade.OrderReturnReceiveReq) (res *trade.OrderReturnReceiveRes, err error) {
	// 获取当前登录用户
	user := service.BizCtx().GetUser(ctx)

	if user == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "需要登录")
	}

	storeId := user.StoreId
	orderReturn, err := service.OrderReturn().Get(ctx, req.ReturnId)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "退货单未找到")
	}

	err = service.OrderReturn().DealWithReturn(ctx, []string{req.ReturnId}, storeId, consts.RETURN_PROCESS_RECEIVED, []*entity.OrderReturn{orderReturn}, 0)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "处理退货单失败")
	}

	return &trade.OrderReturnReceiveRes{}, nil
}

// Refund 确认收款
func (c *cOrderReturn) Refund(ctx context.Context, req *trade.OrderReturnRefundReq) (res *trade.OrderReturnRefundRes, err error) {
	// 获取当前登录用户
	user := service.BizCtx().GetUser(ctx)

	if user == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "需要登录")
	}

	storeId := user.StoreId

	// 获取退货单信息
	orderReturn, err := dao.OrderReturn.Get(ctx, req.ReturnId)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "退货单未找到")
	}

	err = service.OrderReturn().DealWithReturn(ctx, []string{req.ReturnId}, storeId, consts.RETURN_PROCESS_REFUND, []*entity.OrderReturn{orderReturn}, 0)
	if err != nil {
		return nil, gerror.New("处理退货单失败")
	}

	return &trade.OrderReturnRefundRes{}, nil
}

// Refused 卖家拒绝退款/退货
func (c *cOrderReturn) Refused(ctx context.Context, req *trade.OrderReturnRefusedReq) (res *trade.OrderReturnRefusedRes, err error) {
	// 将请求参数转换为 OrderReturn 数据结构
	input := &do.OrderReturn{}
	gconv.Scan(req, input)

	// 获取当前登录用户
	user := service.BizCtx().GetUser(ctx)

	// 如果用户未登录，抛出需要登录的异常
	if user == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "需要登录")
	}

	// 设置订单所属店铺ID为当前用户的店铺ID
	input.StoreId = user.StoreId

	// 调用拒绝退款/退货服务
	success, err := service.OrderReturn().Refused(ctx, input)
	if err != nil {
		return nil, err
	}

	// 如果操作成功，返回成功响应
	if success {
		return &trade.OrderReturnRefusedRes{}, nil
	}

	// 操作失败，返回失败响应
	return nil, gerror.New("操作失败")
}
