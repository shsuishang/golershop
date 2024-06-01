package sys

import (
	"github.com/gogf/gf/v2/frame/g"
)

// =========================== 用户端使用 =============================
type ConfigInfoReq struct {
	g.Meta        `path:"/front/sys/config/info" tags:"配置管理" method:"get" summary:"站点配置信息"`
	SourceUccCode string `json:"source_ucc_code"`
}

type ConfigInfoRes map[string]interface{}

type ConfigPublicKeyReq struct {
	g.Meta `path:"/front/sys/config/publicKey" tags:"配置管理" method:"get" summary:"加密公钥"`
}

type ConfigPublicKeyRes struct {
	PublicKey string `json:"public_key"        `
}

type TranslateLangReq struct {
	g.Meta `path:"/front/sys/config/listTranslateLang" tags:"配置管理" method:"get" summary:"语言包"`
}

type TranslateLangRes struct {
}
