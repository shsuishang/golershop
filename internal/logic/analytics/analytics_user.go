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
package analytics

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/service"
	"golershop.cn/utility/mtime"
)

type sAnalyticsUser struct{}

func init() {
	service.RegisterAnalyticsUser(NewAnalyticsUser())
}

func NewAnalyticsUser() *sAnalyticsUser {
	return &sAnalyticsUser{}
}

// 获取注册用户数量
func (s *sAnalyticsUser) GetRegUser(ctx context.Context) *model.DashboardTopVo {
	topRes := &model.DashboardTopVo{}
	stime, etime := mtime.Today()

	// 获取当日新增用户量
	todayRegUser, err := dao.AnalyticsUser.GetRegUser(ctx, stime, etime)
	topRes.Today = todayRegUser

	// 昨日
	stime, etime = mtime.Yestoday()

	// 获取当日新增用户量
	yestodayRegUser, err := dao.AnalyticsUser.GetRegUser(ctx, stime, etime)
	if err == nil {
		topRes.Yestoday = yestodayRegUser

		// 计算日环比 日环比 = (当日数据 - 前一日数据) / 前一日数据 * 100%
		topRes.Daym2m = 0.0
		if gconv.Float64(yestodayRegUser) != 0 {
			topRes.Daym2m = fmt.Sprintf("%.2f", (gconv.Float64(todayRegUser)-gconv.Float64(yestodayRegUser))/gconv.Float64(yestodayRegUser)*100)
		}
	}

	// 本月
	stime, etime = mtime.Month()
	monthRegUser, err := dao.AnalyticsUser.GetRegUser(ctx, stime, etime)
	if err != nil {
		topRes.Month = monthRegUser
	}

	return topRes
}

// 获取用户时间线
func (s *sAnalyticsUser) GetUserTimeLine(ctx context.Context, input *model.TimelineInput) (out []*model.TimelineOutput, err error) {
	userTimeLine, err := dao.AnalyticsUser.GetUserTimeLine(ctx, input.Stime, input.Etime)

	if err != nil {
		return nil, err
	}

	gconv.Struct(userTimeLine, &out)

	if out == nil {
		out = make([]*model.TimelineOutput, 0)
	}

	return
}

// 获取用户数量
func (s *sAnalyticsUser) GetUserNum(ctx context.Context, input *model.TimelineInput) (out *model.AnalyticsNumOutput, err error) {
	topRes := &model.AnalyticsNumOutput{}

	// 获取当前周期内数据
	currentRegNum, err := dao.AnalyticsUser.GetRegUser(ctx, input.Stime, input.Etime)
	topRes.Current = currentRegNum

	// 上个周期
	preInput := &model.TimelineInput{}
	gconv.Struct(input, preInput)

	if input.Stime > 0 && input.Etime > 0 {
		preInput.Stime = input.Stime - (input.Etime - input.Stime)
		preInput.Etime = input.Stime

		preRegNum, err := dao.AnalyticsUser.GetRegUser(ctx, preInput.Stime, preInput.Etime)

		if err == nil {
			topRes.Pre = preRegNum

			// 计算日环比 日环比 = (当日数据 - 前一日数据) / 前一日数据 * 100%
			topRes.Daym2m = 0.0
			if gconv.Float64(preRegNum) != 0 {
				topRes.Daym2m = fmt.Sprintf("%.2f", (gconv.Float64(currentRegNum)-gconv.Float64(preRegNum))/gconv.Float64(preRegNum)*100)
			}
		}
	}

	return topRes, err
}
