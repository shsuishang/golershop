package account

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/model"
)

type IndexReq struct {
	g.Meta    `path:"/front/account/wechat/index" tags:"login" method:"post" method:"post,get" summary:"验证签名"`
	Timestamp string `json:"timestamp"` //
	Signature string `json:"signature"` //
	Nonce     string `json:"nonce"`     //
	Echostr   string `json:"echostr"`   //

}

type IndexRes string

type WxCodeReq struct {
	g.Meta `path:"/front/account/wechat/redirectToWxCode" tags:"login" method:"post" summary:"公众号登录 - 获取code请求"`
}

type WxCodeRes string

type CallbackMpReq struct {
	g.Meta     `path:"/front/account/wechat/callbackMp" tags:"login" method:"get" summary:"公众号登录 - 公众号授权回调"`
	ActivityId uint   `json:"activity_id"` //
	Code       string `json:"code"`        //
}

type CallbackMpRes model.LoginOutput

type CheckAppLoginReq struct {
	g.Meta `path:"/front/account/wechat/checkAppLogin" tags:"login" method:"get" summary:"用户登录验证"`

	Code string `json:"code"` //
}

type CheckAppLoginRes model.LoginOutput

type WxConfigReq struct {
	g.Meta `path:"/front/account/wechat/wxConfig" tags:"login" method:"get" summary:"获取微信配置请求"`

	Href string `json:"href"` //
}

type WxConfigRes interface {
}

type JsCode2SessionReq struct {
	g.Meta `path:"/front/account/wechat/jsCode2Session" tags:"login" method:"get" summary:"用户注册"`

	Code          string `json:"code"`           //
	EncryptedData string `json:"encryptedData"`  //
	Iv            string `json:"iv"`             //
	UserInfo      string `json:"user_info"`      //
	ActivityId    uint   `json:"activity_id"`    //
	SourceUserId  uint   `json:"source_user_id"` //

}

type JsCode2SessionRes model.LoginOutput

type UserPhoneNumberReq struct {
	g.Meta `path:"/front/account/wechat/getUserPhoneNumber" tags:"login" method:"get" summary:"小程序获取手机号"`

	Code string `json:"code"` //
}

type UserPhoneNumberRes model.LoginOutput

type OpenIdByCodeReq struct {
	g.Meta `path:"/front/account/wechat/getOpenIdByCode" tags:"login" method:"get" summary:"根据code 获取openid"`

	Code string `json:"code"` //
}

type OpenIdByCodeRes map[string]interface{}

type QrCodeReq struct {
	g.Meta `path:"/front/account/wechat/getQrCode" tags:"login" method:"get" summary:"微信网页登录生成二维码"`

	Code string `json:"code"` //
}

type QrCodeRes string

type CallbackPcReq struct {
	g.Meta `path:"/front/account/wechat/callbackPc" tags:"login" method:"get" summary:"二维码登录 - 微信网页回调"`

	Code string `json:"code"` //
}

type CallbackPcRes string
