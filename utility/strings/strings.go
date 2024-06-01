package strings

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func IsEmail(ctx context.Context, email string) (bool, error) {
	// 验证email格式
	if err := g.Validator().Rules("email").Data(email).Run(ctx); err != nil {
		return false, gerror.New("Email不准确！")
	}

	return true, nil
}
