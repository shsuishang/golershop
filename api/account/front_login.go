package account

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/model"
)

type LoginReq struct {
	g.Meta   `path:"/front/account/login/login" tags:"login" method:"post" summary:"登录接口"`
	UserName string `p:"username" v:"required|length:4,30#请输入登录账号|账号长度为：min-max位"`
	Password string `p:"password" v:"required|length:5,9999#请输入密码|密码长度为：min-max位"`
	Captcha  string `p:"verify_code"`
	IdKey    string `p:"verify_key"`
	Encrypt  bool   `p:"encrypt" d:"false" dc:"密码是否加密"`
}

type LoginRes model.LoginOutput

type RegisterReq struct {
	g.Meta      `path:"/front/account/login/register" tags:"login" method:"post" summary:"注册接口"`
	UserAccount string `p:"user_account" v:"required|length:4,50#请输入登录账号|账号长度为：min-max位"`
	Password    string `p:"password" v:"required|length:5,12#请输入密码|密码长度为：min-max位"`
	//Captcha  string `p:"verify_code" v:"required|length:4,6#请输入验证码|验证码长度不够"`
	//IdKey    string `p:"verify_key" v:"required#验证码KEY不能为空"`

	VerifyCode    string `json:"verify_code"`                   // 验证码
	VerifyKey     string `json:"verify_key"`                    // 验证码KEY
	Encrypt       bool   `json:"encrypt"`                       // 密码是否加密
	BindType      int    `json:"bind_type" p:"bind_type" d:"3"` // 注册方式=>BindConnectCode
	ActivityId    int    `json:"activity_id"`                   // 活动编号
	SourceUserId  int    `json:"source_user_id"`                // 来源用户编号
	SourceUccCode string `json:"source_ucc_code"`               // 渠道码
}

type RegisterRes model.LoginOutput

type LogoutReq struct {
	g.Meta `path:"/front/account/login/logout" tags:"login" method:"get" summary:"退出接口"`
}

type LogoutRes struct {
	Token string `json:"token"`
}
type LoginPolicyReq struct {
	g.Meta `path:"/front/account/login/protocol" tags:"login" method:"get" summary:"隐私政策"`

	DocumentType string `json:"document_type"`
	State        uint   `json:"state"`
	ProtocolsKey string `json:"protocols_key" d:"reg_protocols_description" `
}

type LoginPolicyRes struct {
	Document string `json:"document"`
	State    uint   `json:"state"`
}
