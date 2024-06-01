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
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/service"
	"golershop.cn/utility/mtime"
)

type sAnalyticsSys struct{}

func init() {
	service.RegisterAnalyticsSys(NewAnalyticsSys())
}

func NewAnalyticsSys() *sAnalyticsSys {
	return &sAnalyticsSys{}
}

// 计算本月访问量，今日访问量，昨日访问量，以及日环比
func (s *sAnalyticsSys) GetVisitor(ctx context.Context) *model.DashboardTopVo {
	topRes := &model.DashboardTopVo{}

	// 获取当日访问量
	stime, etime := mtime.Today()
	todayVisits, err := dao.AnalyticsSys.GetVisitor(ctx, stime, etime)
	topRes.Today = todayVisits

	// 昨日
	stime, etime = mtime.Yestoday()
	yesterdayVisits, err := dao.AnalyticsSys.GetVisitor(ctx, stime, etime)
	if err == nil {
		topRes.Yestoday = yesterdayVisits

		// 计算日环比 日环比 = (当日数据 - 前一日数据) / 前一日数据 * 100%
		topRes.Daym2m = 0.0
		if !g.IsEmpty(gconv.Float64(yesterdayVisits)) {
			topRes.Daym2m = fmt.Sprintf("%.2f", (gconv.Float64(todayVisits)-gconv.Float64(yesterdayVisits))/gconv.Float64(yesterdayVisits)*100)
		}
	}

	// 本月
	stime, etime = mtime.Month()
	monthVisits, err := dao.AnalyticsSys.GetVisitor(ctx, stime, etime)
	if err != nil {
		topRes.Month = monthVisits
	}

	return topRes
}

// 获取访问访问者数量
func (s *sAnalyticsSys) GetAccessVisitorNum(ctx context.Context, input *model.AccessItemTimelineInput) *model.AnalyticsNumOutput {
	topRes := &model.AnalyticsNumOutput{}

	// 获取当前周期内数据
	currentProductNum, _ := dao.AnalyticsSys.GetVisitorNum(ctx, input.Stime, input.Etime)
	topRes.Current = currentProductNum

	// 上个周期
	preInput := &model.TimelineInput{}
	gconv.Struct(input, preInput)

	if input.Stime != 0 && input.Etime != 0 {
		preInput.Stime -= input.Etime - input.Stime
		preInput.Etime = input.Stime

		preProductNum, err := dao.AnalyticsSys.GetVisitorNum(ctx, preInput.Stime, preInput.Etime)

		if err == nil {
			topRes.Pre = preProductNum

			// 计算日环比 日环比 = (当日数据 - 前一日数据) / 前一日数据 * 100%
			topRes.Daym2m = 0.0
			if gconv.Float64(preProductNum) != 0 {
				topRes.Daym2m = fmt.Sprintf("%.2f", (gconv.Float64(currentProductNum)-gconv.Float64(preProductNum))/gconv.Float64(preProductNum)*100)
			}
		}
	}

	return topRes
}

// 获取访问数量
func (s *sAnalyticsSys) GetAccessNum(ctx context.Context, input *model.AccessItemTimelineInput) *model.AnalyticsNumOutput {
	topRes := &model.AnalyticsNumOutput{}

	// 获取当前周期内数据
	currentProductNum, _ := dao.AnalyticsSys.GetAccessNum(ctx, input.Stime, input.Etime)
	topRes.Current = currentProductNum

	// 上个周期
	preInput := &model.TimelineInput{}
	gconv.Struct(input, preInput)

	if input.Stime != 0 && input.Etime != 0 {
		preInput.Stime -= input.Etime - input.Stime
		preInput.Etime = input.Stime

		preProductNum, err := dao.AnalyticsSys.GetAccessNum(ctx, preInput.Stime, preInput.Etime)

		if err == nil {
			topRes.Pre = preProductNum

			// 计算日环比 日环比 = (当日数据 - 前一日数据) / 前一日数据 * 100%
			topRes.Daym2m = 0.0
			if gconv.Float64(preProductNum) != 0 {
				topRes.Daym2m = fmt.Sprintf("%.2f", (gconv.Float64(currentProductNum)-gconv.Float64(preProductNum))/gconv.Float64(preProductNum)*100)
			}
		}
	}

	return topRes
}

// 获取访问项目时间线
func (s *sAnalyticsSys) GetAccessItemTimeLine(ctx context.Context, input *model.AccessItemTimelineInput) (out []*model.TimelineOutput, err error) {
	accessItemTimeLine, err := dao.AnalyticsSys.GetAccessItemTimeLine(ctx, input.Stime, input.Etime, input.ItemId)

	if err != nil {
		return nil, err
	}

	gconv.Struct(accessItemTimeLine, &out)

	return
}

// 获取访问项目数量
func (s *sAnalyticsSys) GetAccessItemNum(ctx context.Context, input *model.AccessItemTimelineInput) *model.AnalyticsNumOutput {
	topRes := &model.AnalyticsNumOutput{}

	// 获取当前周期内数据
	currentProductNum, _ := dao.AnalyticsSys.GetAccessItemNum(ctx, input.Stime, input.Etime, input.ItemId)
	topRes.Current = currentProductNum

	// 上个周期
	preInput := &model.AccessItemTimelineInput{}
	gconv.Struct(input, preInput)

	if input.Stime != 0 && input.Etime != 0 {
		preInput.Stime -= input.Etime - input.Stime
		preInput.Etime = input.Stime

		preProductNum, err := dao.AnalyticsSys.GetAccessItemNum(ctx, preInput.Stime, preInput.Etime, preInput.ItemId)

		if err == nil {
			topRes.Pre = preProductNum

			// 计算日环比 日环比 = (当日数据 - 前一日数据) / 前一日数据 * 100%
			topRes.Daym2m = 0.0
			if gconv.Float64(preProductNum) != 0 {
				topRes.Daym2m = fmt.Sprintf("%.2f", (gconv.Float64(currentProductNum)-gconv.Float64(preProductNum))/gconv.Float64(preProductNum)*100)
			}
		}
	}

	return topRes
}

// 获取访问项目用户时间线
func (s *sAnalyticsSys) GetAccessItemUserTimeLine(ctx context.Context, input *model.AccessItemTimelineInput) (out []*model.TimelineOutput, err error) {
	accessItemUserTimeLine, err := dao.AnalyticsSys.GetAccessItemUserTimeLine(ctx, input.Stime, input.Etime, input.ItemId)

	if err != nil {
		return nil, err
	}

	gconv.Struct(accessItemUserTimeLine, &out)

	return
}

// 获取访问物品用户数量
func (s *sAnalyticsSys) GetAccessItemUserNum(ctx context.Context, input *model.AccessItemTimelineInput) *model.AnalyticsNumOutput {
	topRes := &model.AnalyticsNumOutput{}

	// 获取当前周期内数据
	currentProductNum, _ := dao.AnalyticsSys.GetAccessItemUserNum(ctx, input.Stime, input.Etime, input.ItemId)
	topRes.Current = currentProductNum

	// 上个周期
	preInput := &model.TimelineInput{}
	gconv.Struct(input, preInput)

	if input.Stime != 0 && input.Etime != 0 {
		preInput.Stime = input.Stime - (input.Etime - input.Stime)
		preInput.Etime = input.Stime

		preProductNum, err := dao.AnalyticsSys.GetAccessItemUserNum(ctx, input.Stime, input.Etime, input.ItemId)

		if err == nil {
			topRes.Pre = preProductNum

			// 计算日环比 日环比 = (当日数据 - 前一日数据) / 前一日数据 * 100%
			topRes.Daym2m = 0.0
			if gconv.Float64(preProductNum) != 0 {
				topRes.Daym2m = fmt.Sprintf("%.2f", (gconv.Float64(currentProductNum)-gconv.Float64(preProductNum))/gconv.Float64(preProductNum)*100)
			}
		}
	}

	return topRes
}

// 获取访问者时间线
func (s *sAnalyticsSys) GetAccessVisitorTimeLine(ctx context.Context, input *model.TimelineInput) (out []*model.TimelineOutput, err error) {
	accessVisitorTimeLine, err := dao.AnalyticsSys.GetAccessVisitorTimeLine(ctx, input.Stime, input.Etime)

	if err != nil {
		return nil, err
	}

	gconv.Struct(accessVisitorTimeLine, &out)

	return
}

// 列出访问物品
func (s *sAnalyticsSys) ListAccessItem(ctx context.Context, timelineInput *model.AccessItemTimelineInput) (out []*model.AnalyticsAccessItemOutput, err error) {
	accessItem, err := dao.AnalyticsSys.ListAccessItem(ctx, timelineInput)

	if err != nil {
		return nil, err
	}

	gconv.Struct(accessItem, &out)

	return
}
