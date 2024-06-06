package marketing

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/marketing"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

// ListVoucher 分页列表
func (c *cActivityBase) ListVoucher(ctx context.Context, req *marketing.ActivityVoucherListReq) (res *marketing.ActivityVoucherListRes, err error) {
	input := &do.ActivityBaseListInput{}
	gconv.Scan(req, &input)

	// 根据请求的 "met" 字段判断是否是首页请求，若是则设置每页返回数量为 3
	if req.Met == "index" {
		input.Size = 3
	}

	// 获取登录用户的ID
	user := service.BizCtx().GetUser(ctx)

	if user != nil {
		input.Where.UserId = user.UserId
	}

	// 调用活动服务的 voucherList 方法
	activityPage, err := service.ActivityBase().ListVoucher(ctx, input)
	if err != nil {
		return nil, err
	}

	// 将结果转换为响应结构体
	res = &marketing.ActivityVoucherListRes{}
	gconv.Struct(activityPage, res)

	return res, nil
}

// listsUserGroupbookingHistory 我的拼团-分页列表查询
func (c *cActivityBase) ListsUserGroupbookingHistory(ctx context.Context, req *marketing.ActivityGroupbookingHistoryReq) (res *marketing.ActivityGroupbookingHistoryRes, err error) {
	input := &do.ActivityGroupbookingHistoryListInput{
		BaseList: ml.BaseList{
			Page: req.Page,
			Size: req.Size,
			Sidx: dao.ActivityGroupbookingHistory.Columns().GbhId,
			Sort: "DESC",
		},
	}
	// 获取当前登录用户
	user := service.BizCtx().GetUser(ctx)
	input.Where.UserId = user.UserId
	if req.GbEnable == "" {
		input.Where.GbEnable = nil
	} else {
		input.Where.GbEnable = req.GbEnable
	}

	var result, error = service.ActivityGroupbooking().ListsUserGroupbookingHistory(ctx, input)

	if error != nil {
		err = error
	}

	// 将结果转换为响应结构体
	gconv.Scan(result, &res)

	return
}

// GetUserGroupbooking 拼团详情
func (c *cActivityBase) GetUserGroupbooking(ctx context.Context, req *marketing.ActivityGroupbookingReq) (res *marketing.ActivityGroupbookingRes, err error) {
	// 获取当前登录用户ID
	user := service.BizCtx().GetUser(ctx)

	// 将用户ID设置到请求参数中
	req.UserId = user.UserId

	// 调用服务层的拼团详情函数
	result, err := service.ActivityGroupbooking().GetUserGroupbooking(ctx, req)
	if err != nil {
		return nil, err
	}

	// 将服务层的结果转换为响应结构体
	gconv.Scan(result, &res)
	return
}
