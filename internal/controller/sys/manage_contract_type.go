package sys

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/sys"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ContractType = cContractType{}
)

type cContractType struct{}

// =================== 管理端使用 =========================

func (c *cContractType) List(ctx context.Context, req *sys.ContractTypeListReq) (res *sys.ContractTypeListRes, err error) {
	var result, error = service.ContractType().List(ctx, &do.ContractTypeListInput{
		BaseList: ml.BaseList{Page: req.Page,
			Size: req.Size,
			Sidx: dao.ContractType.Columns().ContractTypeOrder,
			Sort: "ASC"},
		Where: do.ContractType{},
	})

	if error != nil {
		err = error
	}

	res = &sys.ContractTypeListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}

	return
}

// Add 新增菜单
func (c *cContractType) Add(ctx context.Context, req *sys.ContractTypeAddReq) (res *sys.ContractTypeEditRes, err error) {

	input := do.ContractType{}
	gconv.Scan(req, &input)

	var result, error = service.ContractType().Add(ctx, &input)
	//var result, error = service.ContractType().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &sys.ContractTypeEditRes{
		ContractTypeId: result,
	}

	return
}

// Edit 编辑菜单
func (c *cContractType) Edit(ctx context.Context, req *sys.ContractTypeEditReq) (res *sys.ContractTypeEditRes, err error) {

	input := do.ContractType{}
	gconv.Scan(req, &input)

	var result, error = service.ContractType().Edit(ctx, &input)
	//var result, error = service.ContractType().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &sys.ContractTypeEditRes{
		ContractTypeId: result,
	}

	return
}

// Remove 删除菜单
func (c *cContractType) Remove(ctx context.Context, req *sys.ContractTypeRemoveReq) (res *sys.ContractTypeRemoveRes, err error) {

	idSlice := gstr.Split(req.ContractTypeId, ",")
	for _, contractTypeId := range idSlice {
		var _, error = service.ContractType().Remove(ctx, contractTypeId)

		if error != nil {
			err = error
		}
	}

	res = &sys.ContractTypeRemoveRes{}

	return
}
