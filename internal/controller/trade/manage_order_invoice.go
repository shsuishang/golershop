package trade

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/trade"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	OrderInvoice = cOrderInvoice{}
)

type cOrderInvoice struct{}

// =================== 管理端使用 =========================
func (c *cOrderInvoice) List(ctx context.Context, req *trade.OrderInvoiceListReq) (res *trade.OrderInvoiceListRes, err error) {
	input := do.OrderInvoiceListInput{}
	gconv.Scan(req, &input)

	var result, error = service.OrderInvoice().List(ctx, &input)

	if error != nil {
		err = error
	}

	gconv.Scan(result, &res)

	return
}

// Add 新增
func (c *cOrderInvoice) Add(ctx context.Context, req *trade.OrderInvoiceAddReq) (res *trade.OrderInvoiceEditRes, err error) {
	input := do.OrderInvoice{}
	gconv.Scan(req, &input)

	var result, error = service.OrderInvoice().Add(ctx, &input)

	if error != nil {
		err = error
	}

	res = &trade.OrderInvoiceEditRes{
		OrderInvoiceId: uint(result),
	}

	return
}

// Edit 编辑
func (c *cOrderInvoice) Edit(ctx context.Context, req *trade.OrderInvoiceEditReq) (res *trade.OrderInvoiceEditRes, err error) {

	input := do.OrderInvoice{}
	gconv.Scan(req, &input)

	var result, error = service.OrderInvoice().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &trade.OrderInvoiceEditRes{
		OrderInvoiceId: uint(result),
	}

	return
}

// Remove 删除
func (c *cOrderInvoice) Remove(ctx context.Context, req *trade.OrderInvoiceRemoveReq) (res *trade.OrderInvoiceRemoveRes, err error) {
	var _, error = service.OrderInvoice().Remove(ctx, req.OrderInvoiceId)

	if error != nil {
		err = error
	}

	res = &trade.OrderInvoiceRemoveRes{}

	return
}
