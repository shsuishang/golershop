package marketing

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/marketing"
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

	// 将结果转换为响应结构体
	res = &marketing.ActivityVoucherListRes{}

	return res, nil
}
