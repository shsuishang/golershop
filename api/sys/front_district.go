package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/model"
)

// =========================== 用户端使用 =============================
type DistrictTreeReq struct {
	g.Meta `path:"/front/sys/district/tree" tags:"区域管理" method:"get" summary:"获得地址区域"`
}

type DistrictTreeRes []*model.DistrictTreeNode
