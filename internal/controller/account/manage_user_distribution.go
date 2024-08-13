package account

import (
	"context"
	"golershop.cn/api/account"
)

var (
	UserDistribution = cUserDistribution{}
)

type cUserDistribution struct{}

// =================== 管理端使用 =========================

func (c *cUserDistribution) List(ctx context.Context, req *account.UserDistributionListReq) (res *account.UserDistributionListRes, err error) {

	return
}

// Add 新增用户等级
func (c *cUserDistribution) Add(ctx context.Context, req *account.UserDistributionAddReq) (res *account.UserDistributionEditRes, err error) {

	return res, nil
}

// Edit 编辑用户等级
func (c *cUserDistribution) Edit(ctx context.Context, req *account.UserDistributionEditReq) (res *account.UserDistributionEditRes, err error) {

	return
}

// Edit 编辑用户等级
func (c *cUserDistribution) EditState(ctx context.Context, req *account.UserDistributionEditStateReq) (res *account.UserDistributionEditStateRes, err error) {

	return
}
