package pay

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PayPasswdReq struct {
	g.Meta `path:"/front/pay/index/getPayPasswd" tags:"用户" method:"get" summary:"支付密码"`
}

type PayPasswdRes struct {
}

type ChangeUserPayReq struct {
	g.Meta `path:"/front/pay/index/changePayPassword" tags:"修改支付密码" method:"post" summary:"修改支付密码"`

	OlePayPassword string `json:"old_pay_password" `
	NewPayPassword string `json:"new_pay_password" `
	PayPasssord    string `json:"pay_password" `
}

type ChangeUserPayRes struct {
}
