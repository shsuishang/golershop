package account

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type UserDistributionAdd struct {
}
type UserDistributionEditReq struct {
	g.Meta `path:"/manage/account/userDistribution/edit" tags:"粉丝来源关系表" method:"post" summary:"粉丝来源关系表"`

	UserDistributionAdd
}

type UserDistributionAddRes struct {
}

type UserDistributionEditRes struct {
}

type UserDistributionAddReq struct {
	g.Meta `path:"/manage/account/userDistribution/add" tags:"粉丝来源关系表" method:"post" summary:"粉丝来源关系表"`

	UserDistributionAdd
}

type UserDistributionRemoveReq struct {
	g.Meta `path:"/manage/account/userDistribution/remove" tags:"粉丝来源关系表" method:"post" summary:"粉丝来源关系表"`
}

type UserDistributionRemoveRes struct {
}

type UserDistributionListReq struct {
	g.Meta `path:"/manage/account/userDistribution/list" tags:"粉丝来源关系表" method:"get" summary:"粉丝来源关系表"`
	ml.BaseList
}

type UserDistributionListRes struct {
}

type UserDistributionEditStateReq struct {
	g.Meta `path:"/manage/account/userDistribution/editState" tags:"粉丝来源关系表-修改状态" method:"post" summary:"粉丝来源关系表-修改状态"`
}

type UserDistributionEditStateRes struct {
}
