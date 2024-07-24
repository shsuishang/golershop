package account

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/account"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	UserTagGroup = cUserTagGroup{}
)

type cUserTagGroup struct{}

// ---------------------------- 用户标签 -------------------------------
// List 属性分类列表
func (c *cUserTagGroup) ListAssist(ctx context.Context, req *account.UserTagGroupListReq) (res *account.UserTagGroupListRes, err error) {
	var result, error = service.UserTagGroup().List(ctx, &do.UserTagGroupListInput{
		BaseList: ml.BaseList{Page: req.Page,
			Size: req.Size,
			Sidx: dao.UserTagGroup.Columns().TagGroupSort,
			Sort: "ASC"},
		Where: do.UserTagGroup{},
	})

	if error != nil {
		err = error
	}

	res = &account.UserTagGroupListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 添加属性分类
func (c *cUserTagGroup) AddAssist(ctx context.Context, req *account.UserTagGroupAddReq) (res *account.UserTagGroupEditRes, err error) {
	input := do.UserTagGroup{}
	gconv.Scan(req, &input)

	var result, error = service.UserTagGroup().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &account.UserTagGroupEditRes{
		TagGroupId: result,
	}

	return
}

// Edit 编辑属性分类
func (c *cUserTagGroup) EditAssist(ctx context.Context, req *account.UserTagGroupEditReq) (res *account.UserTagGroupEditRes, err error) {
	input := do.UserTagGroup{}
	gconv.Scan(req, &input)

	_, error := service.UserTagGroup().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &account.UserTagGroupEditRes{
		TagGroupId: req.TagGroupId,
	}

	return
}

// Remove 删除属性分类
func (c *cUserTagGroup) RemoveAssist(ctx context.Context, req *account.UserTagGroupRemoveReq) (res *account.UserTagGroupRemoveRes, err error) {
	var _, error = service.UserTagGroup().Remove(ctx, req.TagGroupId)

	if error != nil {
		err = error
	}

	res = &account.UserTagGroupRemoveRes{}

	return
}

// Tree 标签分组表-树形集合
func (c *cUserTagGroup) Tree(ctx context.Context, req *account.UserTagGroupTreeReq) (res []*account.UserTagGroupTreeRes, err error) {
	// 将请求参数转换为服务层输入结构
	input := &do.UserTagGroupListInput{}
	gconv.Scan(req, input)

	// 调用服务层的树形集合方法
	result, err := service.UserTagGroup().Tree(ctx, input)
	if err != nil {
		return nil, err
	}

	// 将服务层结果转换为响应结构
	res = []*account.UserTagGroupTreeRes{}
	gconv.Scan(result, &res)

	return
}

// ---------------------------- 用户项 -------------------------------

// List 属性项目列表
func (c *cUserTagGroup) ListItem(ctx context.Context, req *account.UserTagBaseListReq) (res *account.UserTagBaseListRes, err error) {
	TagGroupId := req.TagGroupId
	item := do.UserTagBase{TagGroupId: TagGroupId}

	if TagGroupId == 0 {
		item.TagId = nil
	}

	var likes []*ml.WhereExt

	var result, error = service.UserTagBase().List(ctx, &do.UserTagBaseListInput{
		BaseList: ml.BaseList{Page: req.Page,
			Size:      req.Size,
			WhereLike: likes,
			Sidx:      dao.UserTagBase.Columns().TagSort,
			Sort:      "ASC"},
		Where: item,
	})

	if error != nil {
		err = error
	}

	res = &account.UserTagBaseListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 添加属性项目
func (c *cUserTagGroup) AddItem(ctx context.Context, req *account.UserTagBaseAddReq) (res *account.UserTagBaseEditRes, err error) {
	input := do.UserTagBase{}
	gconv.Scan(req, &input)

	var result, error = service.UserTagBase().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &account.UserTagBaseEditRes{
		TagId: result,
	}

	return
}

// Edit 编辑属性项目
func (c *cUserTagGroup) EditItem(ctx context.Context, req *account.UserTagBaseEditReq) (res *account.UserTagBaseEditRes, err error) {
	input := do.UserTagBase{}
	gconv.Scan(req, &input)

	var _, error = service.UserTagBase().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &account.UserTagBaseEditRes{
		TagId: req.TagId,
	}

	return
}

// Remove 删除属性项目
func (c *cUserTagGroup) RemoveItem(ctx context.Context, req *account.UserTagBaseRemoveReq) (res *account.UserTagBaseRemoveRes, err error) {
	var _, error = service.UserTagBase().Remove(ctx, req.TagId)

	if error != nil {
		err = error
	}

	res = &account.UserTagBaseRemoveRes{}

	return
}

// EditState 编辑任务状态
func (c *cUserTagGroup) EditState(ctx context.Context, req *account.UserTagBaseEditStateReq) (res *account.UserTagBaseEditStateRes, err error) {

	input := do.UserTagBase{}
	gconv.Scan(req, &input)

	var result, error = service.UserTagBase().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &account.UserTagBaseEditStateRes{
		TagId: result,
	}

	return
}
