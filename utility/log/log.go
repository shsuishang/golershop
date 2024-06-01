package log

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// Print 信息
func Print(ctx context.Context, v ...interface{}) {
	g.Log().Print(ctx, v)
}

// Info 信息
func Info(ctx context.Context, v ...interface{}) {
	g.Log().Info(ctx, v)
}

// Debug 信息
func Debug(ctx context.Context, v ...interface{}) {
	g.Log().Debug(ctx, v)
}

// Warning 信息
func Warning(ctx context.Context, v ...interface{}) {
	g.Log().Warning(ctx, v)
}

// Error 错误
func Error(ctx context.Context, v ...interface{}) {
	g.Log().Error(ctx, v)
}
