package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/model"
)

type SignInfoReq struct {
	g.Meta `path:"/front/pay/userResource/getSignInfo" tags:"签到" method:"get" summary:"获取签到基本信息"`
}

type SignInfoRes model.SignInfoOutput

type SignInReq struct {
	g.Meta `path:"/front/pay/userResource/signIn" tags:"签到" method:"post" summary:"签到"`
}

type SignInRes struct {
}

type SignStateReq struct {
	g.Meta `path:"/front/pay/userResource/signState" tags:"签到" method:"get" summary:"签到"`
}

type SignStateRes struct {
}

type DistributionCommissionReq struct {
	g.Meta `path:"/front/pay/userResource/getCommissionInfo" tags:"账户余额信息" method:"get" summary:"账户余额信息"`
}

type DistributionCommissionRes struct {
}
