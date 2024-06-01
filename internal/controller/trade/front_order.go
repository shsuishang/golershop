package trade

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/trade"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
	"strconv"
)

var (
	Order = cOrder{}
)

type cOrder struct{}

// List 订单列表信息
func (c *cOrder) List(ctx context.Context, req *trade.UserOrderListReq) (res *trade.UserOrderListRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	input := do.OrderInfoListInput{}
	req.Sidx = dao.OrderInfo.Columns().CreateTime
	req.Sort = ml.ORDER_BY_DESC

	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	input.Where.UserId = userId
	var result, error = service.Order().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Detail 订单详细信息
func (c *cOrder) Detail(ctx context.Context, req *trade.UserOrderDetailReq) (res *trade.UserOrderDetailRes, err error) {
	var result, error = service.Order().Detail(ctx, req.OrderId)

	if error != nil {
		err = error
	}

	//只能访问自己的订单
	userId := service.BizCtx().GetUserId(ctx)
	if result.UserId == userId {
		gconv.Scan(result, &res)
	}

	return
}

// Add 添加订单
func (c *cOrder) Add(ctx context.Context, req *trade.UserOrderAddReq) (res *trade.UserOrderAddRes, err error) {
	user := service.BizCtx().GetUser(ctx)

	input := &model.CheckoutInput{
		UserId:       user.UserId,
		UserNickname: user.UserNickname,
		Items:        []*model.CheckoutItemVo{},
	}

	if !g.IsEmpty(req.UserVoucherIds) {
		input.UserVoucherIds = gconv.SliceUint(gstr.Split(req.UserVoucherIds, ","))
	}

	if !g.IsEmpty(req.OrderMessage) {
		var msg map[uint]string
		err := json.Unmarshal([]byte(req.OrderMessage), &msg)
		if err == nil {
			input.Message = msg
		}
	}

	items := make([]*model.CheckoutItemVo, 0)

	itemInfoRow := gstr.Split(req.CartId, ",")
	for _, item := range itemInfoRow {
		itemRow := gstr.Split(item, "|")
		if cartQuantity, err := strconv.Atoi(itemRow[1]); err != nil || cartQuantity <= 0 {
			return nil, errors.New("购买数量最低为 1 哦~")
		}

		checkoutItemVo := &model.CheckoutItemVo{
			ItemId:       gconv.Uint64(itemRow[0]),
			CartQuantity: gconv.Uint(itemRow[1]),
			CartId:       gconv.Uint64(itemRow[2]),
			CartSelect:   true,
			CartType:     1,
		}

		items = append(items, checkoutItemVo)
	}

	input.Items = items

	res = &trade.UserOrderAddRes{}

	success, err := service.Order().Add(ctx, input)
	if err == nil && len(success.OrderIds) > 0 {
		gconv.Struct(success, res)
	}

	return
}

// Cancel 取消订单
func (c *cOrder) Cancel(ctx context.Context, req *trade.UserOrderCancelReq) (res *trade.UserOrderCancelRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	res = &trade.UserOrderCancelRes{}

	for _, orderId := range req.OrderId {
		orderBase, _ := service.OrderBase().Get(ctx, orderId)
		if orderBase == nil {
			panic("拒绝访问！")
		}

		if orderBase.UserId == userId {
			var _, error = service.Order().Cancel(ctx, orderId, req.OrderCancelReason)

			if error != nil {
				err = error
			} else {
				res.OrderId = append(res.OrderId, orderId)
			}
		}
	}

	if len(res.OrderId) == 0 {
		//err = errors.New("未更改到符合条件的订单！")
	} else {

	}

	return
}

// Receive 确认收货
func (c *cOrder) Receive(ctx context.Context, req *trade.UserOrderReceiveReq) (res *trade.UserOrderReceiveRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	res = &trade.UserOrderReceiveRes{}

	for _, orderId := range req.OrderId {
		orderBase, _ := service.OrderBase().Get(ctx, orderId)
		if orderBase == nil {
			panic("拒绝访问！")
		}

		if orderBase.UserId == userId {
			itemQueryWrapper := &do.OrderReturnItemListInput{Where: do.OrderReturnItem{OrderId: orderId}}
			orderReturnItems, _ := service.OrderReturnItem().Find(ctx, itemQueryWrapper)
			if len(orderReturnItems) > 0 {
				for _, item := range orderReturnItems {
					if item.ReturnStateId != consts.RETURN_PROCESS_FINISH && item.ReturnStateId != consts.RETURN_PROCESS_CANCEL {
						return nil, errors.New("有订单在售后处理中，不能确认收货")
					}
				}
			}

			var _, error = service.Order().Receive(ctx, orderId, "")

			if error != nil {
			} else {
				res.OrderId = append(res.OrderId, orderId)
			}
		}
	}

	if len(res.OrderId) == 0 {
		err = errors.New("未更改到符合条件的订单！")
	} else {

	}

	return
}

/*
// StoreEvaluationWithContent 读取订单商品
func (c *cOrder) StoreEvaluationWithContent(ctx context.Context, orderId string) (*model.CommonRes, error) {
	return service.ProductComment.StoreEvaluationWithContent(orderId)
}

// AddOrderComment 添加订单评论
func (c *cOrder) AddOrderComment(ctx context.Context, req *trade.UserOrderCommentReq) (*trade.UserOrderCommentRes, error) {
	userId := service.ContextUtil.CheckLoginUserId(ctx)
	orderBase, err := service.OrderBase.Get(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}

	if !gutil.CheckDataRights(userId, orderBase, "UserId") {
		return nil, model.BusinessException(ResultCode.FORBIDDEN)
	}

	err = service.ProductComment.AddOrderComment(req.OrderId)
	if err != nil {
		return nil, err
	}

	return model.Success(), nil
}
*/

// OrderNum 获取用户中心订单数量
func (c *cOrder) OrderNum(ctx context.Context, req *trade.UserOrderNumReq) (*trade.UserOrderNumRes, error) {
	userId := service.BizCtx().GetUserId(ctx)
	info, err := service.Order().GetOrderStatisticsInfo(ctx, userId)

	res := &trade.UserOrderNumRes{}
	gconv.Struct(info, res)

	return res, err
}

// ListInvoice 订单发票管理表-分页列表查询
func (c *cOrder) ListInvoice(ctx context.Context, req *trade.UserOrderInvoiceListReq) (res *trade.UserOrderInvoiceListRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	req.UserId = userId

	input := do.OrderInvoiceListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.OrderInvoice().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// AddOrderInvoice 申请订单发票
func (c *cOrder) AddOrderInvoice(ctx context.Context, req *trade.UserOrderInvoiceAddReq) (res *trade.UserOrderInvoiceAddRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	req.UserId = userId

	input := do.OrderInvoice{}
	gconv.Scan(req, &input)

	var result, error = service.OrderInvoice().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &trade.UserOrderInvoiceAddRes{
		OrderInvoiceId: uint(result),
	}

	return
}
