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
	UserDeliveryAddress = cUserDeliveryAddress{}
)

type cUserDeliveryAddress struct{}

// List 用户地址表
func (c *cUserDeliveryAddress) List(ctx context.Context, req *account.UserDeliveryAddressListReq) (res *account.UserDeliveryAddressListRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)

	var input = &do.UserDeliveryAddressListInput{
		Where: do.UserDeliveryAddress{
			UserId: userId,
		},
		BaseList: ml.BaseList{Sidx: "ud_time", Sort: "DESC"},
	}

	result, err := service.UserDeliveryAddress().List(ctx, input)

	if err != nil {
		return nil, err
	}

	res = &account.UserDeliveryAddressListRes{}
	gconv.Struct(result, res)

	return res, err
}

// Get 通过ud_id查询用户地址表
func (c *cUserDeliveryAddress) Get(ctx context.Context, req *account.UserDeliveryAddressGetReq) (*account.UserDeliveryAddressGetRes, error) {
	address, err := service.UserDeliveryAddress().Get(ctx, req.UdId)
	if err != nil {
		return nil, err
	}

	userId := service.BizCtx().GetUserId(ctx)
	if address.UserId == userId {
		res := &account.UserDeliveryAddressGetRes{}
		gconv.Struct(address, res)
		return res, err
	}

	return nil, nil
}

// Add 添加用户地址表
func (c *cUserDeliveryAddress) Add(ctx context.Context, req *account.UserDeliveryAddressAddReq) (*account.UserDeliveryAddressAddRes, error) {
	userId := service.BizCtx().GetUserId(ctx)
	userDeliveryAddress := &do.UserDeliveryAddress{}
	gconv.Struct(req, userDeliveryAddress)
	userDeliveryAddress.UserId = userId
	udId, err := service.UserDeliveryAddress().Add(ctx, userDeliveryAddress)
	if err != nil {
		return nil, err
	}

	userDeliveryAddress.UdId = udId
	res := &account.UserDeliveryAddressAddRes{}
	gconv.Struct(userDeliveryAddress, res)

	return res, nil

}

// Edit 编辑用户地址表
func (c *cUserDeliveryAddress) Edit(ctx context.Context, req *account.UserDeliveryAddressEditReq) (*account.UserDeliveryAddressEditRes, error) {
	userId := service.BizCtx().GetUserId(ctx)
	userDeliveryAddress := &do.UserDeliveryAddress{}
	gconv.Struct(req, userDeliveryAddress)
	userDeliveryAddress.UserId = userId
	address, err := service.UserDeliveryAddress().Get(ctx, req.UdId)
	if address == nil || err != nil {
		return nil, err
	}

	if address.UserId == userId {
		_, err = service.UserDeliveryAddress().Edit(ctx, userDeliveryAddress)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

// Remove 通过ud_id删除用户地址表
func (c *cUserDeliveryAddress) Remove(ctx context.Context, req *account.UserDeliveryAddressRemoveReq) (*account.UserDeliveryAddressEditRes, error) {
	userId := service.BizCtx().GetUserId(ctx)
	address, err := service.UserDeliveryAddress().Get(ctx, req.UdId)
	if err != nil {
		return nil, err
	}

	if address.UserId == userId {
		_, err = service.UserDeliveryAddress().Remove(ctx, address.UdId)
		if err != nil {
			return nil, err
		}
	}

	return nil, err
}
