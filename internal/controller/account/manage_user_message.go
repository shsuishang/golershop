package account

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/account"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	UserMessage = cUserMessage{}
)

type cUserMessage struct{}

// =================== 管理端使用 =========================

func (c *cUserMessage) List(ctx context.Context, req *account.UserMessageListReq) (res *account.UserMessageListRes, err error) {
	input := do.UserMessageListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.UserMessage().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增消息
func (c *cUserMessage) Add(ctx context.Context, req *account.UserMessageAddReq) (res *account.UserMessageEditRes, err error) {

	input := do.UserMessage{}
	gconv.Scan(req, &input)
	input.MessageTime = gtime.Now()

	var result, error = service.UserMessage().Add(ctx, &input)
	//var result, error = service.UserMessage().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &account.UserMessageEditRes{
		MessageId: result,
	}

	return
}

// Edit 编辑消息
func (c *cUserMessage) Edit(ctx context.Context, req *account.UserMessageEditReq) (res *account.UserMessageEditRes, err error) {

	input := do.UserMessage{}
	gconv.Scan(req, &input)
	input.MessageTime = gtime.Now()

	var result, error = service.UserMessage().Edit(ctx, &input)
	//var result, error = service.UserMessage().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &account.UserMessageEditRes{
		MessageId: result,
	}

	return
}

// Remove 删除消息
func (c *cUserMessage) Remove(ctx context.Context, req *account.UserMessageRemoveReq) (res *account.UserMessageRemoveRes, err error) {

	var _, error = service.UserMessage().Remove(ctx, req.MessageId)

	/*
		input := do.UserMessage{}
		input.MessageTime = gtime.Now()
		input.MessageId = req.MessageId[0]
		input.UserMessageSort = 0

		var _, error = service.UserMessage().Edit(ctx, &input)
	*/

	if error != nil {
		err = error
	}

	res = &account.UserMessageRemoveRes{}

	return
}

// MessageNoticeReq
func (c *cUserMessage) MessageNoticeReq(ctx context.Context, req *account.MessageNoticeReq) (res *account.MessageNoticeRes, err error) {

	res = &account.MessageNoticeRes{}

	return
}

// EditState 编辑任务状态
func (c *cUserMessage) EditState(ctx context.Context, req *account.UserMessageEditStateReq) (res *account.UserMessageEditStateRes, err error) {

	input := do.UserMessage{}
	gconv.Scan(req, &input)

	var result, error = service.UserMessage().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &account.UserMessageEditStateRes{
		MessageId: result,
	}

	return
}
