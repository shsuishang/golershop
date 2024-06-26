package marketing

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/marketing"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ActivityBase = cActivityBase{}
)

type cActivityBase struct{}

// =================== 管理端使用 =========================

// List 活动分页列表
func (c *cActivityBase) List(ctx context.Context, req *marketing.ActivityBaseListReq) (res *marketing.ActivityBaseListRes, err error) {
	input := do.ActivityBaseListInput{}
	gconv.Scan(req, &input)

	input.Where.ActivityTypeId = req.ActivityTypeId

	var result, error = service.ActivityBase().GetList(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增活动
func (c *cActivityBase) Add(ctx context.Context, req *marketing.ActivityBaseAddReq) (res *marketing.ActivityBaseEditRes, err error) {

	input := do.ActivityBase{}
	gconv.Scan(req, &input)

	input.ActivityName = req.ActivityTitle

	var result, error = service.ActivityBase().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &marketing.ActivityBaseEditRes{
		ActivityId: result,
	}

	return
}

// Edit 编辑活动
func (c *cActivityBase) Edit(ctx context.Context, req *marketing.ActivityBaseEditReq) (res *marketing.ActivityBaseEditRes, err error) {

	input := do.ActivityBase{}
	gconv.Scan(req, &input)

	input.ActivityName = req.ActivityTitle

	var result, error = service.ActivityBase().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &marketing.ActivityBaseEditRes{
		ActivityId: result,
	}

	return
}

// Remove 删除活动
func (c *cActivityBase) Remove(ctx context.Context, req *marketing.ActivityBaseRemoveReq) (res *marketing.ActivityBaseRemoveRes, err error) {
	var _, error = service.ActivityBase().Remove(ctx, req.ActivityId)

	if error != nil {
		err = error
	}

	res = &marketing.ActivityBaseRemoveRes{}

	return
}

// EditState 活动编辑
func (c *cActivityBase) EditState(ctx context.Context, req *marketing.ActivityBaseEditStateReq) (res *marketing.ActivityBaseEditStateRes, err error) {
	// 获取活动详情
	activityBase, err := service.ActivityBase().Get(ctx, req.ActivityId)
	if err != nil {
		return nil, err
	}

	// 获取当前登录用户
	user := service.BizCtx().GetUser(ctx)
	if user == nil {
		return nil, gerror.New("需要登录")
	}

	// 更新活动状态
	activityBase.ActivityState = req.ActivityState

	input := &do.ActivityBase{}
	gconv.Scan(activityBase, input)

	// 编辑活动
	success, err := service.ActivityBase().EditActivityBase(ctx, req.ActivityId, input)
	if err != nil {
		return nil, err
	}

	if success == false {
		return nil, gerror.New("操作失败")
	}

	return
}
