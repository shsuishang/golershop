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

package log

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
	"golershop.cn/utility"
	"net/url"
)

type sLogAction struct {
	Pool *grpool.Pool
}

func init() {
	service.RegisterLogAction(NewLogAction())
}

func NewLogAction() *sLogAction {
	pool := grpool.New()
	return &sLogAction{Pool: pool}
}

// List 分页读取
func (s *sLogAction) List(ctx context.Context, in *do.LogActionListInput) (out *do.LogActionListOutput, err error) {
	out, err = dao.LogAction.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sLogAction) Add(ctx context.Context, in *do.LogAction) (lastInsertId int64, err error) {
	lastInsertId, err = dao.LogAction.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// OperateLog  异步记录操作日志
func (s *sLogAction) OperateLog(r *ghttp.Request) {
	if r.Method == "POST" {
		u, _ := url.Parse(r.RequestURI)
		user := service.BizCtx().GetUser(r.GetCtx())

		if g.IsEmpty(user) {
			return
		}

		action := do.LogAction{
			UserId:      user.UserId,
			UserAccount: user.UserAccount,
			UserName:    user.UserNickname,
			//Menu:         menu,
			LogUrl:    u.Path,
			LogParam:  r.GetMap(),
			LogMethod: r.Method,
			LogIp:     utility.GetClientIp(r),
			LogDate:   gtime.Now(),
			LogTime:   gtime.Now(),
		}

		s.Pool.Add(r.GetCtx(), func(ctx context.Context) {
			//整理权限信息
			out, err := service.Menu().Find(ctx, &do.MenuBaseListInput{Where: do.MenuBase{MenuPermission: action.LogUrl, MenuType: 0}})

			if err == nil {
				if len(out) > 0 {
					get, err := service.Menu().Get(ctx, out[0].MenuParentId)
					if err == nil {
						action.LogName = fmt.Sprintf("%s - %s", get.MenuTitle, out[0].MenuTitle)
					}
				}
			}

			//写入日志数据
			s.Add(r.GetCtx(), &action)
		})
	}
}
