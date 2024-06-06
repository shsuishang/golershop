package shop

import (
	"context"
	"golershop.cn/api/shop"
	"golershop.cn/internal/service"
)

var (
	UserProductBrowse = cUserProductBrowse{}
)

type cUserProductBrowse struct{}

// List 商品浏览历史表-SPU-不应该直接存数据库-分页列表查询
func (c *cUserProductBrowse) List(ctx context.Context, req *shop.UserProductBrowseListReq) (res []*shop.UserProductBrowseListRes, err error) {
	// 获取当前登录用户ID
	userId := service.BizCtx().GetUserId(ctx)
	if err != nil {
		return nil, err
	}

	// 获取浏览历史列表
	browseResList, err := service.UserProductBrowse().GetList(ctx, userId)
	if err != nil {
		return nil, err
	}

	return browseResList, nil
}

func (c *cUserProductBrowse) Remove(ctx context.Context, req *shop.UserProductBrowseRemoveReq) (res *shop.UserProductBrowseRemoveRes, err error) {
	input := &shop.UserProductBrowseRemoveReq{}
	// 获取当前登录用户ID
	userId := service.BizCtx().GetUserId(ctx)
	if err != nil {
		return nil, err
	}
	input.UserId = userId
	input.ItemId = req.ItemId
	// 获取浏览历史列表
	_, err = service.UserProductBrowse().RemoveBrowser(ctx, input)

	if err != nil {
		return nil, err
	}

	return
}
