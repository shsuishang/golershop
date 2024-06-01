package shop

import (
	"context"
	"golershop.cn/api/pt"
	"golershop.cn/internal/service"
)

func (c *cShop) List(ctx context.Context, req *pt.SearchInfoReq) (res *pt.SearchInfoRes, err error) {

	res, err = service.UserSearchHistory().GetSearchInfo(ctx)

	return
}
