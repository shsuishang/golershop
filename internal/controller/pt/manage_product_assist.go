package pt

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pt"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ProductAssist = cProductAssist{}
)

type cProductAssist struct{}

// =========================== 用户端使用 =============================

// =========================== 管理端使用 =============================

// ---------------------------- 属性分类 -------------------------------
// List 属性分类列表
func (c *cProductAssist) ListAssist(ctx context.Context, req *pt.ProductAssistListReq) (res *pt.ProductAssistListRes, err error) {
	typeId := req.TypeId
	item := do.ProductAssist{TypeId: typeId}

	if typeId == 0 {
		item.AssistId = nil
	}

	var likes []*ml.WhereExt

	if req.AssistName != "" {
		likes = []*ml.WhereExt{{
			Column: dao.ProductAssist.Columns().AssistName,
			Val:    "%" + req.AssistName + "%",
		}}
	}

	var result, error = service.ProductAssist().List(ctx, &do.ProductAssistListInput{
		BaseList: ml.BaseList{Page: req.Page,
			Size:      req.Size,
			WhereLike: likes,
			Sidx:      dao.ProductAssist.Columns().AssistSort,
			Sort:      "ASC"},
		Where: item,
	})

	if error != nil {
		err = error
	}

	res = &pt.ProductAssistListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 添加属性分类
func (c *cProductAssist) AddAssist(ctx context.Context, req *pt.ProductAssistAddReq) (res *pt.ProductAssistEditRes, err error) {
	input := do.ProductAssist{}
	gconv.Scan(req, &input)

	var result, error = service.ProductAssist().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &pt.ProductAssistEditRes{
		AssistId: uint64(result),
	}

	return
}

// Edit 编辑属性分类
func (c *cProductAssist) EditAssist(ctx context.Context, req *pt.ProductAssistEditReq) (res *pt.ProductAssistEditRes, err error) {
	input := do.ProductAssist{}
	gconv.Scan(req, &input)

	_, error := service.ProductAssist().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &pt.ProductAssistEditRes{
		AssistId: req.AssistId,
	}

	return
}

// Remove 删除属性分类
func (c *cProductAssist) RemoveAssist(ctx context.Context, req *pt.ProductAssistRemoveReq) (res *pt.ProductAssistRemoveRes, err error) {
	var _, error = service.ProductAssist().Remove(ctx, req.AssistId)

	if error != nil {
		err = error
	}

	res = &pt.ProductAssistRemoveRes{}

	return
}

//====================================================

// List 属性项目列表
func (c *cProductAssist) ListItem(ctx context.Context, req *pt.ProductAssistItemListReq) (res *pt.ProductAssistItemListRes, err error) {
	assistId := req.AssistId
	item := do.ProductAssistItem{AssistId: assistId}

	if assistId == 0 {
		item.AssistItemId = nil
	}

	var likes []*ml.WhereExt

	if req.AssistItemName != "" {
		likes = []*ml.WhereExt{{
			Column: dao.ProductAssistItem.Columns().AssistItemName,
			Val:    "%" + req.AssistItemName + "%",
		}}
	}

	var result, error = service.ProductAssistItem().List(ctx, &do.ProductAssistItemListInput{
		BaseList: ml.BaseList{Page: req.Page,
			Size:      req.Size,
			WhereLike: likes,
			Sidx:      dao.ProductAssistItem.Columns().AssistItemSort,
			Sort:      "ASC"},
		Where: item,
	})

	if error != nil {
		err = error
	}

	res = &pt.ProductAssistItemListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 添加属性项目
func (c *cProductAssist) AddItem(ctx context.Context, req *pt.ProductAssistItemAddReq) (res *pt.ProductAssistItemEditRes, err error) {
	input := do.ProductAssistItem{}
	gconv.Scan(req, &input)

	var result, error = service.ProductAssistItem().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &pt.ProductAssistItemEditRes{
		AssistItemId: uint64(result),
	}

	return
}

// Edit 编辑属性项目
func (c *cProductAssist) EditItem(ctx context.Context, req *pt.ProductAssistItemEditReq) (res *pt.ProductAssistItemEditRes, err error) {
	input := do.ProductAssistItem{}
	gconv.Scan(req, &input)

	var _, error = service.ProductAssistItem().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &pt.ProductAssistItemEditRes{
		AssistItemId: req.AssistItemId,
	}

	return
}

// Remove 删除属性项目
func (c *cProductAssist) RemoveItem(ctx context.Context, req *pt.ProductAssistItemRemoveReq) (res *pt.ProductAssistItemRemoveRes, err error) {
	var _, error = service.ProductAssistItem().Remove(ctx, req.AssistItemId)

	if error != nil {
		err = error
	}

	res = &pt.ProductAssistItemRemoveRes{}

	return
}
