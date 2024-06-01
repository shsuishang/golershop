package account

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/account"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	UserInvoice = cUserInvoice{}
)

type cUserInvoice struct{}

// List 用户发票表
func (c *cUserInvoice) List(ctx context.Context, req *account.UserInvoiceListReq) (res *account.UserInvoiceListRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)

	var input = &do.UserInvoiceListInput{
		Where: do.UserInvoice{
			UserId: userId,
		},
		BaseList: ml.BaseList{Sidx: dao.UserInvoice.Columns().InvoiceDatetime, Sort: "DESC"},
	}

	result, err := service.UserInvoice().List(ctx, input)

	if err != nil {
		return nil, err
	}

	res = &account.UserInvoiceListRes{}
	gconv.Struct(result, res)

	return res, err
}

// Get 通过ud_id查询用户发票表
func (c *cUserInvoice) Get(ctx context.Context, req *account.UserInvoiceGetReq) (*account.UserInvoiceGetRes, error) {
	userInvoice, err := service.UserInvoice().Get(ctx, req.UserInvoiceId)
	if err != nil {
		return nil, err
	}

	userId := service.BizCtx().GetUserId(ctx)
	if userInvoice.UserId == userId {
		res := &account.UserInvoiceGetRes{}
		gconv.Struct(userInvoice, res)
		return res, err
	}

	return nil, nil
}

// Add 添加用户发票表
func (c *cUserInvoice) Add(ctx context.Context, req *account.UserInvoiceAddReq) (*account.UserInvoiceAddRes, error) {
	userId := service.BizCtx().GetUserId(ctx)
	userInvoice := &do.UserInvoice{}
	gconv.Struct(req, userInvoice)
	userInvoice.UserId = userId
	udId, err := service.UserInvoice().Add(ctx, userInvoice)
	if err != nil {
		return nil, err
	}

	userInvoice.UserInvoiceId = udId
	res := &account.UserInvoiceAddRes{}
	gconv.Struct(userInvoice, res)

	return res, nil

}

// Edit 编辑用户发票表
func (c *cUserInvoice) Edit(ctx context.Context, req *account.UserInvoiceEditReq) (*account.UserInvoiceEditRes, error) {
	userId := service.BizCtx().GetUserId(ctx)
	userInvoice := &do.UserInvoice{}
	gconv.Struct(req, userInvoice)
	userInvoice.UserId = userId

	invoice, err := service.UserInvoice().Get(ctx, req.UserInvoiceId)
	if invoice == nil || err != nil {
		return nil, err
	}

	if invoice.UserId == userId {
		_, err = service.UserInvoice().Edit(ctx, userInvoice)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

// Remove 通过ud_id删除用户发票表
func (c *cUserInvoice) Remove(ctx context.Context, req *account.UserInvoiceRemoveReq) (*account.UserInvoiceEditRes, error) {
	userId := service.BizCtx().GetUserId(ctx)
	userInvoice, err := service.UserInvoice().Get(ctx, req.UserInvoiceId)
	if err != nil {
		return nil, err
	}

	if userInvoice.UserId == userId {
		_, err = service.UserInvoice().Remove(ctx, userInvoice.UserInvoiceId)
		if err != nil {
			return nil, err
		}
	}

	return nil, err
}
