package sys

import (
	"context"
	"golershop.cn/api/sys"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

// =========================== 用户端使用 =============================

// Info 站点配置信息

func (c *cConfig) Info(ctx context.Context, req *sys.ConfigInfoReq) (res sys.ConfigInfoRes, err error) {
	return service.ConfigBase().GetSiteInfo(ctx, req.SourceUccCode)
}

func (c *cConfig) PublicKey(ctx context.Context, req *sys.ConfigPublicKeyReq) (res sys.ConfigPublicKeyRes, err error) {
	res.PublicKey = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDBT2vr+dhZElF73FJ6xiP181txKWUSNLPQQlid6DUJhGAOZblluafIdLmnUyKE8mMHhT3R+Ib3ssZcJku6Hn72yHYj/qPkCGFv0eFo7G+GJfDIUeDyalBN0QsuiE/XzPHJBuJDfRArOiWvH0BXOv5kpeXSXM8yTt5Na1jAYSiQ/wIDAQAB"
	return
}

func (c *cConfig) TranslateLang(ctx context.Context, req *sys.TranslateLangReq) (res sys.TranslateLangRes, err error) {
	return
}
func (c *cConfig) GetPcHelp(ctx context.Context, req *sys.GetPcHelpReq) (res sys.GetPcHelpRes, err error) {
	// 定义配置键
	keys := "page_pc_help"

	// 获取配置项
	pagePcHelp := service.ConfigBase().GetStr(ctx, keys, "")

	// 构建返回结果
	res = sys.GetPcHelpRes{}

	res.PagePcHelp = pagePcHelp

	return
}

// SavePcHelp 保存站点帮助
func (c *cConfig) SavePcHelp(ctx context.Context, req *sys.SavePcHelpReq) (res *sys.SavePcHelpRes, err error) {
	configBase := &do.ConfigBase{
		ConfigKey:      "page_pc_help",
		ConfigValue:    req.PcHelp,
		ConfigTypeId:   0,
		ConfigDatatype: "text",
		ConfigBuildin:  true,
	}

	_, err = service.ConfigBase().Save(ctx, configBase)
	if err != nil {
		return nil, err
	}

	res = &sys.SavePcHelpRes{}
	return res, nil
}
