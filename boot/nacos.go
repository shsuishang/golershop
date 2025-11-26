package boot

import (
	"context"
	"fmt"

	"github.com/gogf/gf/contrib/registry/nacos/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gsvc"
)

// 注册nacos
func initNacosRegistry(ctx context.Context) error {

	nacosIP := g.Cfg().MustGet(ctx, "nacos.server.ip").String()
	nacosPort := g.Cfg().MustGet(ctx, "nacos.server.port").Uint64()

	// Get Nacos registry group configuration
	group := g.Cfg().MustGet(ctx, "nacos.registry.group").String()
	if group == "" {
		group = "GOLERSHOP_GROUP" // Default value
	}

	// Create Nacos registry
	registry := nacos.New(fmt.Sprintf("%s:%d", nacosIP, nacosPort))

	// Set service name and group name
	registry = registry.SetGroupName(group)

	// Set registry as global service registry
	gsvc.SetRegistry(registry)

	return nil
}
