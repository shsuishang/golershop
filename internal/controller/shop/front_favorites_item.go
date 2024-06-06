package shop

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/shop"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	FavoritesItem = cFavoritesItem{}
)

type cFavoritesItem struct{}

func (c *cFavoritesItem) List(ctx context.Context, req *shop.UserFavoritesItemListReq) (res *shop.UserFavoritesItemListRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	req.UserId = userId

	input := do.UserFavoritesItemListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.UserFavoritesItem().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增收藏
func (c *cFavoritesItem) Add(ctx context.Context, req *shop.UserFavoritesItemAddReq) (res *shop.UserFavoritesItemEditRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)

	input := do.UserFavoritesItem{}
	gconv.Scan(req, &input)

	input.UserId = userId

	var result, error = service.UserFavoritesItem().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &shop.UserFavoritesItemEditRes{
		FavoritesItemId: uint(result),
	}

	return
}

// Edit 编辑收藏
func (c *cFavoritesItem) Edit(ctx context.Context, req *shop.UserFavoritesItemEditReq) (res *shop.UserFavoritesItemEditRes, err error) {
	input := do.UserFavoritesItem{}
	gconv.Scan(req, &input)

	return
}

// Remove 删除收藏
func (c *cFavoritesItem) Remove(ctx context.Context, req *shop.UserFavoritesItemRemoveReq) (res *shop.UserFavoritesItemRemoveRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)

	input := &do.UserFavoritesItemListInput{Where: do.UserFavoritesItem{UserId: userId, ItemId: req.ItemId}}

	if favoritesItem, err := service.UserFavoritesItem().FindOne(ctx, input); err != nil {
		return nil, err
	} else {
		if favoritesItem == nil {
			return nil, errors.New("无权限！")
		}

		var _, error = service.UserFavoritesItem().Remove(ctx, favoritesItem.FavoritesItemId)

		if error != nil {
			err = error
		}
	}

	res = &shop.UserFavoritesItemRemoveRes{}

	return
}

// List 用户收藏列表
func (c *cFavoritesItem) Lists(ctx context.Context, req *shop.UserFavoritesItemListsReq) (res *shop.UserFavoritesItemListsRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	req.UserId = userId

	input := do.UserFavoritesItemListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	// 调用服务获取列表
	pageList, err := service.UserFavoritesItem().GetList(ctx, &input)

	if err != nil {
		return nil, err
	}

	gconv.Scan(pageList, &res)

	return
}
