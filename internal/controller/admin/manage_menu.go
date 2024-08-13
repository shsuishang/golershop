package admin

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/admin"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	Menu = cMenu{}
)

type cMenu struct{}

// =================== 管理端使用 =========================

// 查询商品
func (c *cMenu) Detail(ctx context.Context, req *admin.MenuDetailReq) (res *admin.MenuDetailRes, err error) {

	var result, error = service.Menu().Get(ctx, req.MenuId)

	if error != nil {
		err = error
	}

	res = &admin.MenuDetailRes{
		MenuBase: result,
	}

	return
}

func (c *cMenu) List(ctx context.Context, req *admin.MenuListReq) (res *admin.MenuListRes, err error) {
	var result, error = service.Menu().List(ctx, &do.MenuBaseListInput{
		BaseList: ml.BaseList{
			Page: req.Page,
			Size: req.Size},
		Where: do.MenuBase{
			MenuHidden: 0,
		},
	})

	if error != nil {
		err = error
	}

	res = &admin.MenuListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Tree 树形菜单
func (c *cMenu) Tree(ctx context.Context, req *admin.MenuTreeReq) (res admin.MenuTreeRes, err error) {
	var menuType any

	//是否读取全部类型菜单
	if req.Type == 2 {
		menuType = nil
	} else {
		menuType = req.Type
	}

	var likes = []*ml.WhereExt{{
		Column: dao.MenuBase.Columns().MenuTitle,
		Val:    "%" + req.Title + "%",
		Symbol: ml.LIKE,
	}}

	var result1, error1 = service.Menu().GetTree(ctx, &do.MenuBaseListInput{
		Where: do.MenuBase{
			MenuType: menuType,
		},
		BaseList: ml.BaseList{WhereExt: likes},
	})

	if error1 != nil {
		err = error1
	}

	gconv.Scan(result1, &res)

	return
}

// Add 新增菜单
func (c *cMenu) Add(ctx context.Context, req *admin.MenuAddReq) (res *admin.MenuEditRes, err error) {

	input := do.MenuBase{}
	gconv.Scan(req, &input)
	input.MenuTime = gtime.Now()

	var result, error = service.Menu().Add(ctx, &input)
	//var result, error = service.Menu().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &admin.MenuEditRes{
		MenuId: result,
	}

	return
}

// Edit 编辑菜单
func (c *cMenu) Edit(ctx context.Context, req *admin.MenuEditReq) (res *admin.MenuEditRes, err error) {

	input := do.MenuBase{}
	gconv.Scan(req, &input)
	input.MenuTime = gtime.Now()

	var result, error = service.Menu().Edit(ctx, &input)
	//var result, error = service.Menu().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &admin.MenuEditRes{
		MenuId: result,
	}

	return
}

// Remove 删除菜单
func (c *cMenu) Remove(ctx context.Context, req *admin.MenuRemoveReq) (res *admin.MenuRemoveRes, err error) {

	var _, error = service.Menu().Remove(ctx, req.MenuId)

	/*
		input := do.MenuBase{}
		input.MenuTime = gtime.Now()
		input.MenuId = req.MenuId[0]
		input.MenuSort = 0

		var _, error = service.Menu().Edit(ctx, &input)
	*/

	if error != nil {
		err = error
	}

	res = &admin.MenuRemoveRes{}

	return
}

// EditState 编辑菜单状态
func (c *cMenu) EditState(ctx context.Context, req *admin.MenuEditStateReq) (res *admin.MenuEditStateRes, err error) {
	input := do.MenuBase{}
	gconv.Scan(req, &input)

	var result, error = service.Menu().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &admin.MenuEditStateRes{
		MenuId: result,
	}

	return
}
