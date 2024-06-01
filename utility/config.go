package utility

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func UrlBase(ctx context.Context) string {
	// 附件存储路径
	url, _ := g.Cfg().Get(ctx, "shopSuite.urlBase")
	if url.String() != "" {
		return url.String()
	} else {
	}

	return ""
}

func UrlH5(ctx context.Context) string {
	url, _ := g.Cfg().Get(ctx, "shopSuite.urlH5")
	if url.String() != "" {
		return url.String()
	} else {
	}

	return ""
}

func UrlPc(ctx context.Context) string {
	url, _ := g.Cfg().Get(ctx, "shopSuite.urlPc")
	if url.String() != "" {
		return url.String()
	} else {
	}

	return ""
}

// 缓存模式
func CacheEnable(ctx context.Context) bool {
	debug, _ := g.Cfg().Get(ctx, "shopSuite.cacheEnable")
	if debug.String() == "true" {
		return true
	} else {
		return false
	}
}

// 空间
func Namespace(ctx context.Context) string {
	namespace, _ := g.Cfg().Get(ctx, "shopSuite.cacheNamespace")
	if namespace.String() != "" {
		return namespace.String()
	} else {
		return ""
	}
}

// 调试模式
func AppDebug(ctx context.Context) bool {
	debug, _ := g.Cfg().Get(ctx, "shopSuite.appDebug")
	if debug.String() == "true" {
		return true
	} else {
		return false
	}
}

// 系统版本号
func GetVersion(ctx context.Context) string {
	version, _ := g.Cfg().Get(ctx, "shopSuite.version")
	if version.String() != "" {
		return version.String()
	} else {
		return ""
	}
}
