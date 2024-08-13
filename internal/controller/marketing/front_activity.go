package marketing

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/marketing"
	"golershop.cn/api/pt"
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

func (c *cActivityBase) List(ctx context.Context, req *marketing.ActivityListReq) (res *marketing.ActivityListRes, err error) {
	input := do.ActivityBaseListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.ActivityBase().GetList(ctx, &input)

	if error != nil {
		err = error
	}

	res = &marketing.ActivityListRes{}
	gconv.Scan(result, &res)

	return
}

// GetActivityInfo
func (c *cActivityBase) GetActivityInfo(ctx context.Context, req *marketing.ProductItemListReq) (res *marketing.ProductItemListRes, err error) {
	input := &pt.ItemListReq{}
	gconv.Scan(req, &input)

	var result, error = service.ProductIndex().ListItem(ctx, input)

	if error != nil {
		err = error
		return
	}

	res = &marketing.ProductItemListRes{}
	gconv.Scan(result, res)
	return res, nil
}
