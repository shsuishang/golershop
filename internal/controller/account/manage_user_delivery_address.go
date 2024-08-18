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
	DeliveryAddress = cDeliveryAddress{}
)

type cDeliveryAddress struct{}

func (c *cDeliveryAddress) List(ctx context.Context, req *account.DeliveryAddressListReq) (res *account.DeliveryAddressListRes, err error) {
	input := do.UserDeliveryAddressListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.UserDeliveryAddress().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增
func (c *cDeliveryAddress) Add(ctx context.Context, req *account.DeliveryAddressAddReq) (res *account.DeliveryAddressEditRes, err error) {

	userDeliveryAddress := &do.UserDeliveryAddress{}
	gconv.Struct(req, userDeliveryAddress)
	userDeliveryAddress.UserId = req.UserId
	udId, err := service.UserDeliveryAddress().Add(ctx, userDeliveryAddress)
	if err != nil {
		return nil, err
	}

	userDeliveryAddress.UdId = udId
	res = &account.DeliveryAddressEditRes{}
	gconv.Struct(userDeliveryAddress, res)

	return res, nil
}

// Edit 编辑
func (c *cDeliveryAddress) Edit(ctx context.Context, req *account.DeliveryAddressEditReq) (res *account.DeliveryAddressEditRes, err error) {

	input := do.UserDeliveryAddress{}
	gconv.Scan(req, &input)

	var result, error = service.UserDeliveryAddress().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &account.DeliveryAddressEditRes{
		UdId: uint(result),
	}

	return
}

// Remove 删除
func (c *cDeliveryAddress) Remove(ctx context.Context, req *account.DeliveryAddressRemoveReq) (res *account.DeliveryAddressRemoveRes, err error) {

	var _, error = service.UserDeliveryAddress().Remove(ctx, req.UdId)

	if error != nil {
		err = error
	}

	res = &account.DeliveryAddressRemoveRes{}

	return
}
