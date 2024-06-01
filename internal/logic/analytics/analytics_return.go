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
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/service"
)

type sAnalyticsReturn struct{}

func init() {
	service.RegisterAnalyticsReturn(NewAnalyticsReturn())
}

func NewAnalyticsReturn() *sAnalyticsReturn {
	return &sAnalyticsReturn{}
}

// GetReturnAmountTimeline 获取退单金额时间线
func (s *sAnalyticsReturn) GetReturnAmountTimeline(ctx context.Context, input *model.TimelineInput) (out []*model.TimelineOutput, err error) {
	// 统计没有取消的退单
	returnStateIds := []uint{
		consts.RETURN_PROCESS_FINISH,
		consts.RETURN_PROCESS_CHECK,
		consts.RETURN_PROCESS_RECEIVED,
		consts.RETURN_PROCESS_REFUND,
		consts.RETURN_PROCESS_RECEIPT_CONFIRMATION,
		consts.RETURN_PROCESS_REFUSED,
		consts.RETURN_PROCESS_SUBMIT,
	}

	returnAmountTimeline, err := dao.AnalyticsReturn.GetReturnAmountTimeline(ctx, input.Stime, input.Etime, returnStateIds)

	if err != nil {
		return nil, err
	}

	gconv.Struct(returnAmountTimeline, &out)

	return
}

// GetReturnNum 获取退单数量
func (s *sAnalyticsReturn) GetReturnNum(ctx context.Context, input *model.AnalyticsReturnInput) (*model.AnalyticsNumOutput, error) {
	topRes := &model.AnalyticsNumOutput{}

	// 统计没有取消的退单
	returnStateIds := []uint{
		consts.RETURN_PROCESS_FINISH,
		consts.RETURN_PROCESS_CHECK,
		consts.RETURN_PROCESS_RECEIVED,
		consts.RETURN_PROCESS_REFUND,
		consts.RETURN_PROCESS_RECEIPT_CONFIRMATION,
		consts.RETURN_PROCESS_REFUSED,
		consts.RETURN_PROCESS_SUBMIT,
	}
	input.ReturnStateId = returnStateIds

	// 获取当前周期内数据
	currentRegNum, err := dao.AnalyticsReturn.GetReturnNum(ctx, input.Stime, input.Etime, input.ReturnStateId)
	if err != nil {
		return nil, err
	}

	if currentRegNum != nil {
		topRes.Current = currentRegNum
	}

	// 上个周期
	preInput := &model.AnalyticsReturnInput{}
	gconv.Struct(input, preInput)

	if !g.IsEmpty(input.Stime) && !g.IsEmpty(input.Etime) {
		preInput.Stime = input.Stime - (input.Etime - input.Stime)
		preInput.Etime = input.Stime

		preRegNum, err := dao.AnalyticsReturn.GetReturnNum(ctx, preInput.Stime, preInput.Etime, preInput.ReturnStateId)
		if err != nil {
			return nil, err
		}

		if !g.IsEmpty(preRegNum) {
			topRes.Pre = preRegNum

			// 计算日环比 日环比 = (当日数据 - 前一日数据) / 前一日数据 * 100%
			topRes.Daym2m = 0.0
			if gconv.Float64(preRegNum) != 0 {
				topRes.Daym2m = fmt.Sprintf("%.2f", (gconv.Float64(currentRegNum)-gconv.Float64(preRegNum))/gconv.Float64(preRegNum)*100)
			}
		}
	}

	return topRes, nil
}

// GetReturnAmount 获取退单金额
func (s *sAnalyticsReturn) GetReturnAmount(ctx context.Context, input *model.AnalyticsReturnInput) (*model.AnalyticsNumOutput, error) {
	topRes := &model.AnalyticsNumOutput{}

	// 统计没有取消的退单
	returnStateIds := []uint{
		consts.RETURN_PROCESS_FINISH,
		consts.RETURN_PROCESS_CHECK,
		consts.RETURN_PROCESS_RECEIVED,
		consts.RETURN_PROCESS_REFUND,
		consts.RETURN_PROCESS_RECEIPT_CONFIRMATION,
		consts.RETURN_PROCESS_REFUSED,
		consts.RETURN_PROCESS_SUBMIT,
	}
	input.ReturnStateId = returnStateIds

	// 获取当前周期内数据
	currentRegNum, err := dao.AnalyticsReturn.GetReturnAmount(ctx, input.Stime, input.Etime, input.ReturnStateId)
	if err != nil {
		return nil, err
	}

	topRes.Current = currentRegNum

	// 上个周期
	preInput := &model.TimelineInput{}
	gconv.Struct(input, preInput)

	if !g.IsEmpty(input.Stime) && !g.IsEmpty(input.Etime) {
		preInput.Stime = input.Stime - (input.Etime - input.Stime)
		preInput.Etime = input.Stime

		preRegNum, err := dao.AnalyticsReturn.GetReturnAmount(ctx, preInput.Stime, preInput.Etime, input.ReturnStateId)
		if err != nil {
			return nil, err
		}

		if !g.IsEmpty(preRegNum) {
			topRes.Pre = preRegNum

			// 计算日环比 日环比 = (当日数据 - 前一日数据) / 前一日数据 * 100%
			topRes.Daym2m = 0.0
			if gconv.Float64(preRegNum) != 0 {
				topRes.Daym2m = fmt.Sprintf("%.2f", (gconv.Float64(currentRegNum)-gconv.Float64(preRegNum))/gconv.Float64(preRegNum)*100)
			}
		}
	}

	return topRes, nil
}

// GetReturnNumTimeline 获取退单数量时间线
func (s *sAnalyticsReturn) GetReturnNumTimeline(ctx context.Context, input *model.TimelineInput) (out []*model.TimelineOutput, err error) {
	// 统计没有取消的退单
	returnStateIds := []uint{
		consts.RETURN_PROCESS_FINISH,
		consts.RETURN_PROCESS_CHECK,
		consts.RETURN_PROCESS_RECEIVED,
		consts.RETURN_PROCESS_REFUND,
		consts.RETURN_PROCESS_RECEIPT_CONFIRMATION,
		consts.RETURN_PROCESS_REFUSED,
		consts.RETURN_PROCESS_SUBMIT,
	}

	returnTimeLine, err := dao.AnalyticsReturn.GetReturnTimeLine(ctx, input.Stime, input.Etime, returnStateIds)

	if err != nil {
		return nil, err
	}

	gconv.Struct(returnTimeLine, &out)

	return
}

// GetReturnItemNumTimeLine 获取退单商品数量时间线
func (s *sAnalyticsReturn) GetReturnItemNumTimeLine(ctx context.Context, input *model.OrderItemNumTimelineInput) (out []*model.TimelineOutput, err error) {
	returnTimeLine, err := dao.AnalyticsReturn.GetReturnItemNumTimeLine(ctx, input)

	if err != nil {
		return nil, err
	}

	gconv.Struct(returnTimeLine, &out)

	return
}

// GetReturnItemNum 获取退单商品数量
func (s *sAnalyticsReturn) GetReturnItemNum(ctx context.Context, input *model.OrderItemNumTimelineInput) (*model.AnalyticsNumOutput, error) {
	topRes := &model.AnalyticsNumOutput{}

	// 获取当前周期内数据
	currentProductNum, err := dao.AnalyticsReturn.GetReturnItemNum(ctx, input)
	if err != nil {
		return nil, err
	}

	if currentProductNum != nil {
		topRes.Current = currentProductNum
	}

	// 上个周期
	preInput := &model.OrderItemNumTimelineInput{}
	gconv.Struct(input, preInput)

	if !g.IsEmpty(input.Stime) && !g.IsEmpty(input.Etime) {
		preInput.Stime = input.Stime - (input.Etime - input.Stime)
		preInput.Etime = input.Stime

		preProductNum, err := dao.AnalyticsReturn.GetReturnItemNum(ctx, preInput)
		if err != nil {
			return nil, err
		}

		if !g.IsEmpty(preProductNum) {
			topRes.Pre = preProductNum

			// 计算日环比 日环比 = (当日数据 - 前一日数据) / 前一日数据 * 100%
			topRes.Daym2m = 0.0
			if gconv.Float64(preProductNum) != 0 {
				topRes.Daym2m = fmt.Sprintf("%.2f", (gconv.Float64(currentProductNum)-gconv.Float64(preProductNum))/gconv.Float64(preProductNum)*100)
			}
		}
	}

	return topRes, nil
}

// ListReturnItemNum 获取退单商品数量列表
func (s *sAnalyticsReturn) ListReturnItemNum(ctx context.Context, input *model.OrderItemNumTimelineInput) (out []*model.AnalyticsOrderItemNumOutput, err error) {
	returnTimeLine, err := dao.AnalyticsReturn.ListReturnItemNum(ctx, input)

	if err != nil {
		return nil, err
	}

	gconv.Struct(returnTimeLine, &out)

	return
}

// GetReturnCustomerNumTimeline 获取退单客户统计
func (s *sAnalyticsReturn) GetReturnCustomerNumTimeline(ctx context.Context, input *model.TimelineInput) (out []*model.TimelineOutput, err error) {
	return
}
