package marketing

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/marketing"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
	"strconv"
)

var (
	ActivityBase = cActivityBase{}
)

type cActivityBase struct{}

// =================== 管理端使用 =========================

// GetList 活动分页列表
func (c *cActivityBase) GetList(ctx context.Context, req *marketing.ActivityBaseListReq) (res *marketing.ActivityBaseListRes, err error) {
	input := do.ActivityBaseListInput{
		Where: do.ActivityBase{
			ActivityTypeId: req.ActivityTypeId,
		},
		BaseList: ml.BaseList{
			Sidx: dao.ActivityBase.Columns().ActivityId,
			Sort: ml.ORDER_BY_DESC,
		},
	}

	whereExt := []*ml.WhereExt{}

	if !g.IsEmpty(req.ActivityName) {
		whereExt = append(whereExt, &ml.WhereExt{
			Column: dao.ActivityBase.Columns().ActivityName,
			Val:    "%" + req.ActivityName + "%",
			Symbol: ml.LIKE,
		})
	}

	if !g.IsEmpty(req.ActivityState) {
		whereExt = append(whereExt, &ml.WhereExt{
			Column: dao.ActivityBase.Columns().ActivityState,
			Val:    req.ActivityState,
		})
	}

	input.WhereExt = whereExt

	var result, error = service.ActivityBase().GetList(ctx, &input)

	if error != nil {
		err = error
	}

	res = &marketing.ActivityBaseListRes{}
	gconv.Scan(result, &res)

	return
}

// Add 新增活动
func (c *cActivityBase) Add(ctx context.Context, req *marketing.ActivityBaseAddReq) (res *marketing.ActivityBaseEditRes, err error) {
	// 获取当前登录用户
	user := service.BizCtx().GetUser(ctx)
	if user == nil {
		return nil, gerror.New("需要登录")
	}

	// 设置用户相关信息
	req.UserId = user.UserId
	req.StoreId = user.StoreId
	req.SubsiteId = strconv.Itoa(int(user.SiteId))

	// 调用服务添加活动
	success, err := service.ActivityBase().AddActivityBase(ctx, req)
	if err != nil {
		return nil, err
	}

	if !success {
		return nil, gerror.New("添加活动失败")
	}

	res = &marketing.ActivityBaseEditRes{}

	return res, nil
}

// Edit 编辑活动
func (c *cActivityBase) Edit(ctx context.Context, req *marketing.ActivityBaseEditReq) (res *marketing.ActivityBaseEditRes, err error) {

	// 获取当前登录用户
	user := service.BizCtx().GetUser(ctx)
	if user == nil {
		return nil, gerror.New("需要登录")
	}

	req.StoreId = user.StoreId

	_, err = service.ActivityBase().UpdateActivityBase(ctx, req)
	if err != nil {
		return nil, err
	}

	return nil, err
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
