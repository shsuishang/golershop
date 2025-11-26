package boot

import (
	"github.com/gogf/gf/v2/os/gctx"
)

// It tries to use Nacos configuration first, but falls back to local configuration if Nacos is unavailable.
func init() {
	var (
		ctx = gctx.GetInitCtx()
	)

	// Try to initialize Nacos configuration
	if err := initNacosConfig(ctx); err != nil {
	}
}
