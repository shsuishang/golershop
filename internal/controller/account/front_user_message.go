package account

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/account"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

// IM配置
// @summary IM配置
// @description IM配置接口
// @tags 配置
// @produce json
// @router /getImConfig [get]
func (c *cUser) GetImConfig(ctx context.Context, req *account.ImConfigReq) (res *account.ImConfigRes, err error) {
	return &account.ImConfigRes{}, nil
}

// 客服配置
// @summary 客服配置
// @description 客服配置接口
// @tags 配置
// @produce json
// @router /getKefuConfig [get]
func (c *cUser) GetKefuConfig(ctx context.Context, req *account.KefuConfigReq) (res *account.KefuConfigRes, err error) {
	return &account.KefuConfigRes{}, nil
}

func (c *cUser) GetNotice(ctx context.Context, req *account.NoticeListReq) (res *account.NoticeListRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	req.UserId = userId

	input := do.UserMessageListInput{}
	gconv.Scan(req, &input)

	if !g.IsEmpty(req.MessageTitle) {
		var likes = []*ml.WhereExt{{
			Column: dao.UserMessage.Columns().MessageTitle,
			Val:    "%" + req.MessageTitle + "%",
			Symbol: ml.LIKE,
		}}

		input.WhereExt = likes
	}

	var result, error = service.UserMessage().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

func (c *cUser) GetMsgCount(ctx context.Context, req *account.MsgCountReq) (res *account.MsgCountRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	req.UserId = userId

	input := model.MsgCountInput{}
	gconv.Scan(req, &input)

	if !g.IsEmpty(req.RecentlyFlag) {
		input.RecentlyFlag = req.RecentlyFlag
	}

	var result, error = service.UserMessage().GetMsgCount(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// 短消息列表数据
// @summary 短消息列表数据
// @description 短消息列表数据
// @tags 短消息
// @produce json
// @param userMessageListReq query do.UserMessageListReq true "用户消息列表请求"
// @router /list [get]
func (c *cUser) ListMessage(ctx context.Context, req *account.MessageListReq) (res *account.MessageListRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	req.UserId = userId

	input := do.UserMessageListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	input.Order = []*ml.BaseOrder{{Sidx: dao.UserMessage.Columns().MessageTime, Sort: ml.ORDER_BY_DESC}, {Sidx: dao.UserMessage.Columns().MessageIsRead, Sort: ml.ORDER_BY_ASC}}

	if input.Where.MessageKind == 0 {
		input.Where.MessageKind = 2
	}

	var result, error = service.UserMessage().GetList(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// 读取短消息
// @summary 读取短消息
// @description 读取短消息
// @tags 短消息
// @produce json
// @param message_id query int true "消息ID"
// @router /get [get]
func (c *cUser) Get(ctx context.Context, req *account.MessageGetReq) (res *account.MessageGetRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)

	userMessage, err := service.UserMessage().GetById(ctx, req.MessageId, userId)
	if err != nil {
		return nil, err
	}

	res = &account.MessageGetRes{}
	gconv.Struct(userMessage, res)

	return res, nil
}

// 设置为已读
// @summary 设置为已读
// @description 设置为已读
// @tags 短消息
// @produce json
// @param message_id formData int false "消息ID"
// @param user_other_id formData int false "其他用户ID"
// @router /setRead [post]
func (c *cUser) SetRead(ctx context.Context, req *account.MsgReadReq) (res *account.MsgReadRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	success, err := service.UserMessage().SetRead(ctx, req.MessageId, req.UserOtherId, userId)
	if err != nil || !success {
		return nil, err
	}
	return &account.MsgReadRes{}, nil
}

// 添加短消息
// @summary 添加短消息
// @description 添加短消息
// @tags 短消息
// @produce json
// @param messageAddInput body do.UserMessageAddInput true "消息添加输入"
// @router /add [post]
func (c *cUser) Add(ctx context.Context, req *account.MessageAddReq) (res *account.MessageAddRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	input := &model.UserMessageAddInput{}
	gconv.Scan(req, &input)

	userMessageRes, err := service.UserMessage().AddMessage(ctx, input, userId)
	if err != nil {
		return nil, err
	}
	res = &account.MessageAddRes{}
	gconv.Struct(userMessageRes, res)

	return res, nil
}

// 读取分页列表
// @summary 读取分页列表
// @description 读取分页列表
// @tags 短消息
// @produce json
// @param userMessageListReq query do.UserMessageListReq true "用户消息列表请求"
// @router /listChatMsg [get]
func (c *cUser) ListChatMsg(ctx context.Context, req *account.ChatMsgListReq) (res *account.ChatMsgListRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	req.UserId = userId

	input := &do.UserMessageListInput{}
	gconv.Scan(req, input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	if !g.IsEmpty(req.SourceType) && req.SourceType != 2311 {
		input.Sidx = dao.UserMessage.Columns().MessageTime
		input.Sort = ml.ORDER_BY_ASC
	} else {
		input.Sidx = dao.UserMessage.Columns().MessageTime
		input.Sort = ml.ORDER_BY_DESC
	}

	pageList, err := service.UserMessage().ListChatMsg(ctx, input)
	if err != nil {
		return nil, err
	}

	res = &account.ChatMsgListRes{}
	gconv.Struct(pageList, res)

	return res, nil
}

// 消息中心-信息数
// @summary 消息中心-信息数
// @description 消息中心-信息数
// @tags 短消息
// @produce json
// @param userMessageListReq query do.UserMessageListReq true "用户消息列表请求"
// @router /getMessageNum [get]
func (c *cUser) GetMessageNum(ctx context.Context, req *account.MessageNumReq) (res *account.MessageNumRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)
	req.UserId = userId
	input := &do.UserMessageListInput{}
	gconv.Scan(req, input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	userMessageRes, err := service.UserMessage().GetMessageNum(ctx, input)
	if err != nil {
		return nil, err
	}

	res = &account.MessageNumRes{}
	gconv.Struct(userMessageRes, res)

	return res, nil
}
