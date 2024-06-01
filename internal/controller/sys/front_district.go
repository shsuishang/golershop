package sys

import (
	"context"
	"golershop.cn/api/sys"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

// =========================== 用户端使用 =============================

// Aera 读取区域信息
func (c *cConfig) DistrictTree(ctx context.Context, req *sys.DistrictTreeReq) (res sys.DistrictTreeRes, err error) {
	res, err = service.DistrictBase().Tree(ctx, &do.DistrictBaseListInput{})

	if err != nil {
	}

	return
}
