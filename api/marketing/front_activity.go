package marketing

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model"
)

type ActivityVoucherListReq struct {
	g.Meta `path:"/front/marketing/activityBase/listVoucher" tags:"活动管理" method:"get" summary:"活动列表接口"`
	ml.BaseList
	ActivityReqVo
}

type ActivityVoucherListRes model.ActivityListOutput
