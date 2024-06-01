package sys

import (
	"context"
	"golershop.cn/api/sys"
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
