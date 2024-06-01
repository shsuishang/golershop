package sys

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/sys"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	Dict = cDict{}
)

type cDict struct{}

// =================== 管理端使用 =========================
// List 字典分类列表
func (c *cDict) List(ctx context.Context, req *sys.DictBaseListReq) (res *sys.DictBaseListRes, err error) {
	var result, error = service.Dict().List(ctx, &do.DictBaseListInput{
		BaseList: ml.BaseList{Page: req.Page,
			Size: req.Size,
			Sidx: dao.DictBase.Columns().DictSort,
			Sort: "ASC"},
		Where: do.DictBase{},
	})

	if error != nil {
		err = error
	}

	res = &sys.DictBaseListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 添加字典分类
func (c *cDict) Add(ctx context.Context, req *sys.DictAddReq) (res *sys.DictEditRes, err error) {
	input := do.DictBase{}
	gconv.Scan(req, &input)

	var _, error = service.Dict().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.DictEditRes{
		DictId: req.DictId,
	}

	return
}

// Edit 编辑字典分类
func (c *cDict) Edit(ctx context.Context, req *sys.DictEditReq) (res *sys.DictEditRes, err error) {
	input := do.DictBase{}
	gconv.Scan(req, &input)

	var _, error = service.Dict().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.DictEditRes{
		DictId: req.DictId,
	}

	return
}

// Remove 删除字典分类
func (c *cDict) Remove(ctx context.Context, req *sys.DictRemoveReq) (res *sys.DictRemoveRes, err error) {
	var _, error = service.Dict().Remove(ctx, req.DictId)

	if error != nil {
		err = error
	}

	res = &sys.DictRemoveRes{}

	return
}

// List 字典项目列表
func (c *cDict) ListItem(ctx context.Context, req *sys.DictItemListReq) (res *sys.DictItemListRes, err error) {
	dictId := gstr.Trim(req.DictId)
	item := do.DictItem{DictId: req.DictId}
	if dictId == "" {
		item.DictId = nil
	}

	var result, error = service.DictItem().List(ctx, &do.DictItemListInput{
		BaseList: ml.BaseList{Page: req.Page,
			Size: req.Size,
			Sidx: dao.DictItem.Columns().DictItemSort,
			Sort: "ASC"},
		Where: item,
	})

	if error != nil {
		err = error
	}

	res = &sys.DictItemListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 添加字典项目
func (c *cDict) AddItem(ctx context.Context, req *sys.DictItemAddReq) (res *sys.DictItemEditRes, err error) {
	input := do.DictItem{}
	gconv.Scan(req, &input)

	var _, error = service.DictItem().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.DictItemEditRes{
		DictItemId: req.DictItemId,
	}

	return
}

// Edit 编辑字典项目
func (c *cDict) EditItem(ctx context.Context, req *sys.DictItemEditReq) (res *sys.DictItemEditRes, err error) {
	input := do.DictItem{}
	gconv.Scan(req, &input)

	var _, error = service.DictItem().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.DictItemEditRes{
		DictItemId: req.DictItemId,
	}

	return
}

// Remove 删除字典项目
func (c *cDict) RemoveItem(ctx context.Context, req *sys.DictItemRemoveReq) (res *sys.DictItemRemoveRes, err error) {
	var _, error = service.DictItem().Remove(ctx, req.DictItemId)

	if error != nil {
		err = error
	}

	res = &sys.DictItemRemoveRes{}

	return
}
