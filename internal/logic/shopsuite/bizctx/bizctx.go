// +----------------------------------------------------------------------
// | ShopSuite商城系统 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 随商信息技术（上海）有限公司
// +----------------------------------------------------------------------
// | 未获商业授权前，不得将本软件用于商业用途。禁止整体或任何部分基础上以发展任何派生版本、
// | 修改版本或第三方版本用于重新分发。
// +----------------------------------------------------------------------
// | 官方网站: https://www.shopsuite.cn  https://www.golershop.cn
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本公司对该软件产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细见https://www.golershop.cn/policy
// +----------------------------------------------------------------------

package bizctx

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/dao/global"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/log"
)

type (
	sBizCtx struct{}
)

func init() {
	service.RegisterBizCtx(New())
}

func New() *sBizCtx {
	return &sBizCtx{}
}

// Init initializes and injects custom business context object into request context.
func (s *sBizCtx) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// Get retrieves and returns the user object from context.
// It returns nil if nothing found in given context.
func (s *sBizCtx) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser injects business user object into context.
func (s *sBizCtx) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser

	if ctxUser.ClientId == 1 {
		//user admin
		userAdmin, _ := dao.UserAdmin.Get(ctx, ctxUser.UserId)

		if !g.IsEmpty(userAdmin) {
			//角色编号:0-用户;2-商家;3-门店;9-平台;
			ctxUser.RoleId = consts.ROLE_ADMIN

			userRole, _ := dao.UserRole.Get(ctx, userAdmin.UserRoleId)

			if !g.IsEmpty(userRole) {
				ctxUser.Roles = make([]*entity.UserRole, 0, 1)
				ctxUser.Roles = append(ctxUser.Roles, userRole)

				entitys, _ := dao.MenuBase.Gets(ctx, gstr.Split(userRole.MenuIds, ","))
				ctxUser.Authorities = entitys
			}
		}
	}
}

func (s *sBizCtx) GetUser(ctx context.Context) *model.ContextUser {
	if v := s.Get(ctx); v != nil && v.User != nil {
		return v.User
	}
	return nil
}

func (s *sBizCtx) GetUserId(ctx context.Context) uint {
	if v := s.Get(ctx); v != nil && v.User != nil {
		return v.User.UserId
	}

	return 0
}

func (s *sBizCtx) GetRequest(ctx context.Context) *ghttp.Request {
	return g.RequestFromCtx(ctx)
}

func (s *sBizCtx) IncrementTx(ctx context.Context) {
	if v := s.Get(ctx); v != nil {
		v.Tx = v.Tx + 1
	}
}

func (s *sBizCtx) DecrementTx(ctx context.Context) {
	if v := s.Get(ctx); v != nil {
		if v.Tx >= 1 {
			v.Tx = v.Tx - 1
		}
	}
}

func (s *sBizCtx) GetTx(ctx context.Context) int {
	if v := s.Get(ctx); v != nil {
		return v.Tx
	}

	return 0
}

func (s *sBizCtx) AddCacheKey(ctx context.Context, keys []string) {
	if v := s.Get(ctx); v != nil {
		v.CacheKeys = append(v.CacheKeys, keys...)
	}
}

func (s *sBizCtx) GetCacheKeys(ctx context.Context) []string {
	if v := s.Get(ctx); v != nil {
		return v.CacheKeys
	}

	return nil
}

func (s *sBizCtx) AfterOutput(r *ghttp.Request) {
	if global.Cache {
		if v := s.Get(r.GetCtx()); v != nil {
			keys := v.CacheKeys

			if !g.IsEmpty(keys) {
				_, err := g.Redis().Del(r.GetCtx(), keys...)
				if err != nil {
					log.Error(r.GetCtx(), err)
				}
			}
		}
	}
}
