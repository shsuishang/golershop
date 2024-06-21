package account

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/account"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	UserBindConnect = cUserBindConnect{}
)

type cUserBindConnect struct{}

// List 用户绑定表-平台-分页列表查询
func (c *cUserBindConnect) List(ctx context.Context, req *account.UserBindConnectListReq) (res *account.UserBindConnectListRes, err error) {
	input := &do.UserBindConnectListInput{}
	gconv.Scan(req, input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.BaseList.WhereExt)

	pageList, err := service.UserBindConnect().List(ctx, input)
	if err != nil {
		return nil, err
	}

	res = &account.UserBindConnectListRes{}
	gconv.Scan(pageList, res)
	return
}
