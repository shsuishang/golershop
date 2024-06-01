package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/sys"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	LogAction = cLogAction{}
)

type cLogAction struct{}

// =================== 管理端使用 =========================

// List 活动分页列表
func (c *cLogAction) ListAction(ctx context.Context, req *sys.LogActionListReq) (res *sys.LogActionListRes, err error) {
	input := do.LogActionListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	input.Sidx = dao.LogAction.Columns().LogId
	input.Sort = ml.ORDER_BY_DESC

	var result, error = service.LogAction().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// List 活动分页列表
func (c *cLogAction) ListError(ctx context.Context, req *sys.LogErrorListReq) (res *sys.LogErrorListRes, err error) {
	input := do.LogErrorListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	input.Sidx = dao.LogError.Columns().LogErrorId
	input.Sort = ml.ORDER_BY_DESC

	var result, error = service.LogError().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}
