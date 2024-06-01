package sys

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/sys"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	MessageTemplate = cMessageTemplate{}
)

type cMessageTemplate struct{}

// =================== 管理端使用 =========================
func (c *cMessageTemplate) List(ctx context.Context, req *sys.MessageTemplateListReq) (res *sys.MessageTemplateListRes, err error) {
	input := do.MessageTemplateListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.MessageTemplate().List(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.MessageTemplateListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 新增类型
func (c *cMessageTemplate) Add(ctx context.Context, req *sys.MessageTemplateAddReq) (res *sys.MessageTemplateEditRes, err error) {

	input := do.MessageTemplate{}
	gconv.Scan(req, &input)

	var result, error = service.MessageTemplate().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.MessageTemplateEditRes{
		MessageId: string(result),
	}

	return
}

// Edit 编辑类型
func (c *cMessageTemplate) Edit(ctx context.Context, req *sys.MessageTemplateEditReq) (res *sys.MessageTemplateEditRes, err error) {

	input := do.MessageTemplate{}
	gconv.Scan(req, &input)

	var result, error = service.MessageTemplate().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &sys.MessageTemplateEditRes{
		MessageId: string(result),
	}

	return
}

// Remove 删除类型
func (c *cMessageTemplate) Remove(ctx context.Context, req *sys.MessageTemplateRemoveReq) (res *sys.MessageTemplateRemoveRes, err error) {

	idSlice := gstr.Split(req.MessageId, ",")
	for _, contractTypeId := range idSlice {
		var _, error = service.MessageTemplate().Remove(ctx, contractTypeId)

		if error != nil {
			err = error
		}
	}

	res = &sys.MessageTemplateRemoveRes{}

	return
}

// EditState 编辑状态
func (c *cMessageTemplate) EditState(ctx context.Context, req *sys.MessageTemplateEditStateReq) (res *sys.MessageTemplateEditStateRes, err error) {
	input := do.MessageTemplate{}
	gconv.Scan(req, &input)

	_, err = service.MessageTemplate().Edit(ctx, &input)

	res = &sys.MessageTemplateEditStateRes{
		MessageId: req.MessageId,
	}

	return
}
