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

type QrcodeReq struct {
	g.Meta `path:"/front/sys/config/getQrcode" tags:"配置管理" method:"get" summary:"根据提货码获取二维码"`
}

type QrcodeRes struct {
}

type GuideReq struct {
	g.Meta `path:"/front/sys/config/guide" tags:"配置管理" method:"get" summary:"加载广告导航"`
}

type GuideRes struct {
}

type TranslateLangReq struct {
	g.Meta `path:"/front/sys/config/listTranslateLang" tags:"配置管理" method:"get" summary:"语言包"`
}

type TranslateLangRes struct {
}

type GetPcHelpReq struct {
	g.Meta `path:"/front/sys/config/getPcHelp" tags:"站点帮助" method:"get" summary:"站点帮助"`
}

type GetPcHelpRes struct {
	PagePcHelp string `json:"page_pc_help"  `
}
type SavePcHelpReq struct {
	g.Meta `path:"/manage/sys/config/savePcHelp" tags:"站点帮助" method:"post" summary:"站点帮助"`
	PcHelp string `json:"pc_help"  `
}

type SavePcHelpRes struct {
}
