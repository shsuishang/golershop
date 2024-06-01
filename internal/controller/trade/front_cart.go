package trade

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/trade"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
	"strconv"
	"strings"
)

var (
	Cart = cCart{}
)

// Cart 购物车控制器
type cCart struct{}

// List 购物车分页列表查询
func (c *cCart) List(ctx context.Context, req *trade.UserCartListReq) (*trade.UserCartListRes, error) {
	userId := service.BizCtx().GetUserId(ctx)

	input := &do.UserCartListInput{}
	gconv.Struct(req, input)
	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	input.Where.UserId = userId
	input.Size = consts.MAX_LIST_NUM

	list, err := service.UserCart().GetList(ctx, input)

	res := &trade.UserCartListRes{}
	gconv.Scan(list, res)

	return res, err
}

// Add 购物车添加
func (c *cCart) Add(ctx context.Context, req *trade.UserCartAddReq) (*trade.UserCartAddRes, error) {
	userId := service.BizCtx().GetUserId(ctx)

	userCart := &model.CartAddInput{}
	gconv.Struct(req, userCart)
	userCart.UserId = userId
	userCart.CartSelect = true

	_, err := service.UserCart().AddCart(ctx, userCart)

	return &trade.UserCartAddRes{}, err
}

// Edit 购物车编辑
func (c *cCart) Edit(ctx context.Context, req *trade.UserCartEditReq) (*trade.UserCartEditRes, error) {
	userCart := &do.UserCart{}
	gconv.Struct(req, userCart)

	out, err := service.UserCart().Edit(ctx, userCart)

	return &trade.UserCartEditRes{CartId: gconv.Uint64(out)}, err
}

// Remove 通过购物车ID删除购物车
func (c *cCart) Remove(ctx context.Context, req *trade.UserCartRemoveReq) (*trade.UserCartRemoveRes, error) {
	out, err := service.UserCart().Remove(ctx, req.CartId)

	return &trade.UserCartRemoveRes{CartId: gconv.Uint64(out)}, err
}

// RemoveBatch 批量删除购物车
func (c *cCart) RemoveBatch(ctx context.Context, req *trade.UserCartRemoveBatchReq) (*trade.UserCartRemoveBatchRes, error) {
	cartIdList := gconv.SliceInt64(strings.Split(req.CartIds, ","))

	out, err := service.UserCart().Remove(ctx, cartIdList)

	return &trade.UserCartRemoveBatchRes{CartId: gconv.Uint64(out)}, err
}

// EditQuantity 修改购物车数量
func (c *cCart) EditQuantity(ctx context.Context, req *trade.UserCartEditQuantityReq) (*trade.UserCartEditRes, error) {
	userId := service.BizCtx().GetUserId(ctx)

	userCart := &do.UserCart{}
	gconv.Struct(req, userCart)

	out, err := service.UserCart().EditQuantity(ctx, userCart, userId)

	return &trade.UserCartEditRes{CartId: gconv.Uint64(out)}, err
}

// Sel 购物车选中商品
func (c *cCart) Sel(ctx context.Context, req *trade.UserCartSelectReq) (*trade.UserCartSelectRes, error) {
	userId := service.BizCtx().GetUserId(ctx)

	input := &model.UserCartSelectInput{}
	gconv.Struct(req, input)
	input.UserId = userId

	_, err := service.UserCart().Sel(ctx, input)

	return &trade.UserCartSelectRes{}, err
}

// AddBatch 批量添加购物车
func (c *cCart) AddBatch(ctx context.Context, req *trade.UserCartAddBatReq) (*trade.UserCartAddRes, error) {
	userId := service.BizCtx().GetUserId(ctx)

	userCart := &do.UserCart{}
	gconv.Struct(req, userCart)
	userCart.UserId = userId

	out, err := service.UserCart().Add(ctx, userCart)

	return &trade.UserCartAddRes{CartId: gconv.Uint64(out)}, err
}

// Checkout 结算
func (c *cCart) Checkout(ctx context.Context, req *trade.OrderCheckoutReq) (*trade.CheckoutRes, error) {
	userId := service.BizCtx().GetUserId(ctx)

	input := &model.CheckoutInput{}
	gconv.Struct(req, input)

	// 优惠券
	if !g.IsEmpty(req.UserVoucherIds) {
		ts := gconv.SliceUint(gstr.Split(req.UserVoucherIds, ","))
		input.UserVoucherIds = ts
	}
	// 处理数据
	input.UserId = userId

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
	out, err := service.UserCart().Checkout(ctx, input)

	res := &trade.CheckoutRes{}
	gconv.Scan(out, res)

	return res, err
}
