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
	"golershop.cn/utility/mtime"
)

type sAnalyticsOrder struct{}

func init() {
	service.RegisterAnalyticsOrder(NewAnalyticsOrder())
}

func NewAnalyticsOrder() *sAnalyticsOrder {
	return &sAnalyticsOrder{}
}

// 获取后台仪表盘中部柱形数据，包含[订单数据，用户数据，商品数据，销售额数据]
func (s *sAnalyticsOrder) GetDashboardTimeLine(ctx context.Context, input *model.TimelineInput) (model.DashBoardTimelineOutput, error) {
	var dashBoardTimeLineRes model.DashBoardTimelineOutput

	orderTimelineOutput, err := dao.AnalyticsOrder.GetOrderTimeLine(ctx, input.Stime, input.Etime)
	if err != nil {
		return dashBoardTimeLineRes, err
	}
	userTimelineOutput, err := dao.AnalyticsUser.GetUserTimeLine(ctx, input.Stime, input.Etime)
	if err != nil {
		return dashBoardTimeLineRes, err
	}
	ptTimelineOutput, err := dao.AnalyticsProduct.GetProductTimeLine(ctx, input.Stime, input.Etime)
	if err != nil {
		return dashBoardTimeLineRes, err
	}

	payTimelineOutput, err := dao.AnalyticsOrder.GetPayTimeLine(ctx, input.Stime, input.Etime)
	if err != nil {
		return dashBoardTimeLineRes, err
	}

	gconv.Struct(orderTimelineOutput, &dashBoardTimeLineRes.OrderTimeLine)
	gconv.Struct(userTimelineOutput, &dashBoardTimeLineRes.UserTimeLine)
	gconv.Struct(ptTimelineOutput, &dashBoardTimeLineRes.PtTimeLine)
	gconv.Struct(payTimelineOutput, &dashBoardTimeLineRes.PayTimeLine)

	if dashBoardTimeLineRes.OrderTimeLine == nil {
		dashBoardTimeLineRes.OrderTimeLine = make([]*model.TimelineOutput, 0)
	}

	if dashBoardTimeLineRes.UserTimeLine == nil {
		dashBoardTimeLineRes.UserTimeLine = make([]*model.TimelineOutput, 0)
	}

	if dashBoardTimeLineRes.PtTimeLine == nil {
		dashBoardTimeLineRes.PtTimeLine = make([]*model.TimelineOutput, 0)
	}

	if dashBoardTimeLineRes.PayTimeLine == nil {
		dashBoardTimeLineRes.PayTimeLine = make([]*model.TimelineOutput, 0)
	}

	return dashBoardTimeLineRes, nil
}

func (s *sAnalyticsOrder) GetSaleOrderAmount(ctx context.Context, input *model.TimelineInput) (out []*model.TimelineOutput, err error) {
	out = make([]*model.TimelineOutput, 0)

	saleOrderAmount, err := dao.AnalyticsOrder.GetSaleOrderAmount(ctx, input.Stime, input.Etime)
	if err != nil {
		return out, err
	}

	gconv.Struct(saleOrderAmount, &out)

	if out == nil {
		out = make([]*model.TimelineOutput, 0)
	}

	return
}

/*
// 获取总交易额和近一周交易额增长率
func (s *sAnalyticsOrder) GetTradeData(ctx context.Context, dashBoardRes *model.AdminDashBoardVo, recentOneWeek, recentTwoWeek time.Time) {
	// 总交易额
	tradeQueryWrapper := dao.ConsumeTrade.Query().Where("trade_is_paid", consts.ORDER_PAID_STATE_YES).WhereIn("trade_type_id", []uint{consts.TRADE_TYPE_SHOPPING, consts.TRADE_TYPE_FAVORABLE})
	tradeAmount, err := tradeQueryWrapper.Count()
	if err != nil {
		return
	}
	dashBoardRes.TradeAmount = tradeAmount

	// 总交易额近一周增长率
	tradeQueryWrapper.Where("trade_paid_time > ?", recentOneWeek)
	recentOneWeekTradeAmount, err := tradeQueryWrapper.Count()
	if err != nil {
		return
	}
	tradeQueryWrapper.Where("trade_paid_time > ?", recentTwoWeek)
	recentTwoWeekTradeAmount, err := tradeQueryWrapper.Count()
	if err != nil {
		return
	}

	tradeAmountRate := 0.0
	if tradeAmount != 0 && recentOneWeekTradeAmount != 0 {
		tradeAmountRate = float64((recentOneWeekTradeAmount - recentTwoWeekTradeAmount) / recentOneWeekTradeAmount * 100)
	}
	dashBoardRes.TradeAmountIncreaseRate = tradeAmountRate
}

// 获取总订单成交和订单成交近一周增长率
func (s *sAnalyticsOrder) GetOrderFinishData(ctx context.Context, dashBoardRes *model.AdminDashBoardVo, recentOneWeek, recentTwoWeek time.Time) {
	// 订单完成总成交
	infoQueryWrapper := dao.OrderInfo.Query().WhereIn("order_state_id", consts.ORDER_STATE_FINISH)
	orderFinishNum, err := infoQueryWrapper.Count()
	if err != nil {
		return
	}
	dashBoardRes.OrderFinishNum = orderFinishNum

	// 订单完成成交增长率
	infoQueryWrapper.Where("order_received_time > ?", recentOneWeek)
	recentOneWeekOrderFinishNum, err := infoQueryWrapper.Count()
	if err != nil {
		return
	}
	infoQueryWrapper.Where("order_received_time > ?", recentTwoWeek)
	recentTwoWeekOrderFinishNum, err := infoQueryWrapper.Count()
	if err != nil {
		return
	}

	orderFinishNumIncreaseRate := 0.0
	if orderFinishNum != 0 && recentOneWeekOrderFinishNum != 0 {
		orderFinishNumIncreaseRate = float64((recentOneWeekOrderFinishNum - recentTwoWeekOrderFinishNum) / recentOneWeekOrderFinishNum * 100)
	}
	dashBoardRes.OrderFinishNumIncreaseRate = orderFinishNumIncreaseRate
}

// 获取订单总量和订单近一周增长率
func (s *sAnalyticsOrder) GetOrderData(ctx context.Context, dashBoardRes *model.AdminDashBoardVo, recentOneWeek, recentTwoWeek time.Time) {
	// 订单总成交
	infoQueryWrapper := dao.OrderInfo.Query().WhereIn("order_state_id", s.getOrderStates())
	orderNum, err := infoQueryWrapper.Count()
	if err != nil {
		return
	}
	dashBoardRes.OrderNum = orderNum

	// 订单总成交近一周增长率
	infoQueryWrapper.Where("order_received_time > ?", recentOneWeek)
	recentOneWeekOrderNum, err := infoQueryWrapper.Count()
	if err != nil {
		return
	}
	infoQueryWrapper.Where("order_received_time > ?", recentTwoWeek)
	recentTwoWeekOrderNum, err := infoQueryWrapper.Count()
	if err != nil {
		return
	}

	orderNumIncreaseRate := 0.0
	if orderNum != 0 && recentOneWeekOrderNum != 0 {
		orderNumIncreaseRate = float64((recentOneWeekOrderNum - recentTwoWeekOrderNum) / recentOneWeekOrderNum * 100)
	}
	dashBoardRes.OrderNumIncreaseRate = orderNumIncreaseRate
}

// 获取用户总数和用户近一周增长率
func (s *sAnalyticsOrder) GetUserCertificationData(ctx context.Context, dashBoardRes *model.AdminDashBoardVo, recentOneWeek, recentTwoWeek time.Time) {
	// 会员总数
	userCertificationNum, err := dao.UserInfo.Query().Count()
	if err != nil {
		return
	}
	dashBoardRes.UserCertificationNum = userCertificationNum

	// 会员近一周增长率
	loginQueryWrapper := dao.UserLogin.Query().Where("user_reg_time > ?", recentOneWeek)
	recentOneWeekCertificationNum, err := loginQueryWrapper.Count()
	if err != nil {
		return
	}
	loginQueryWrapper.Where("user_reg_time > ?", recentTwoWeek)
	recentTwoWeekCertificationNum, err := loginQueryWrapper.Count()
	if err != nil {
		return
	}

	userCertificationNumIncreaseRate := 0.0
	if userCertificationNum != 0 && recentOneWeekCertificationNum != 0 {
		userCertificationNumIncreaseRate = float64((recentOneWeekCertificationNum - recentTwoWeekCertificationNum) / recentOneWeekCertificationNum * 100)
	}
	dashBoardRes.UserCertificationNumIncreaseRate = userCertificationNumIncreaseRate
}

// 获取所有订单状态
func (s *sAnalyticsOrder) getOrderStates() []int {
	return []uint{
		consts.ORDER_STATE_WAIT_PAY,
		consts.ORDER_STATE_WAIT_PAID,
		consts.ORDER_STATE_WAIT_REVIEW,
		consts.ORDER_STATE_WAIT_FINANCE_REVIEW,
		consts.ORDER_STATE_PICKING,
		consts.ORDER_STATE_WAIT_SHIPPING,
		consts.ORDER_STATE_SHIPPED,
		consts.ORDER_STATE_RECEIVED,
		consts.ORDER_STATE_FINISH,
		consts.ORDER_STATE_CANCEL,
		consts.ORDER_STATE_SELF_PICKUP,
	}
}
*/

func (s *sAnalyticsOrder) GetOrderCustomerNumTimeline(ctx context.Context, input *model.TimelineInput) (out []*model.TimelineOutput, err error) {
	orderCustomerNumTimeline, err := dao.AnalyticsOrder.GetOrderCustomerNumTimeline(ctx, input.Stime, input.Etime)

	if err != nil {
		return nil, err
	}

	gconv.Struct(orderCustomerNumTimeline, &out)

	return
}

// 计算本月订单量，今日订单量，昨日订单量，以及日环比
func (s *sAnalyticsOrder) GetOrderNumToday(ctx context.Context) (topRes model.DashboardTopVo, err error) {
	// 获取当日订单量
	stime, etime := mtime.Today()

	//统计没有取消的订单
	orderStateId := []uint{consts.ORDER_STATE_WAIT_PAY, consts.ORDER_STATE_WAIT_PAID, consts.ORDER_STATE_WAIT_REVIEW, consts.ORDER_STATE_WAIT_FINANCE_REVIEW, consts.ORDER_STATE_PICKING, consts.ORDER_STATE_WAIT_SHIPPING, consts.ORDER_STATE_SHIPPED, consts.ORDER_STATE_RECEIVED, consts.ORDER_STATE_FINISH, consts.ORDER_STATE_SELF_PICKUP}
	paidStateId := []uint{consts.ORDER_PAID_STATE_NO, consts.ORDER_PAID_STATE_PART, consts.ORDER_PAID_STATE_YES}
	todayOrderNum, err := dao.AnalyticsOrder.GetOrderNum(ctx, stime, etime, orderStateId, paidStateId, 0, 0)

	if err != nil {
		return topRes, err
	}

	topRes.Today = todayOrderNum

	// 获取昨日订单量
	stime, etime = mtime.Yestoday()
	yesterdayOrderNum, err := dao.AnalyticsOrder.GetOrderNum(ctx, stime, etime, orderStateId, paidStateId, 0, 0)

	if err != nil {
		return topRes, err
	}
	topRes.Yestoday = yesterdayOrderNum

	// 获取本月订单量
	stime, etime = mtime.Month()
	monthOrderNum, err := dao.AnalyticsOrder.GetOrderNum(ctx, stime, etime, orderStateId, paidStateId, 0, 0)

	if err != nil {
		return topRes, err
	}
	topRes.Month = monthOrderNum

	// 获取日环比
	topRes.Daym2m = 0.0
	if gconv.Float64(yesterdayOrderNum) != 0 {
		topRes.Daym2m = fmt.Sprintf("%.2f", (gconv.Float64(todayOrderNum)-gconv.Float64(yesterdayOrderNum))/gconv.Float64(yesterdayOrderNum)*100)
	}
	return topRes, nil
}

// 分析订单数据
func (s *sAnalyticsOrder) GetOrderNum(ctx context.Context, input *model.AnalyticsOrderInput) (topRes model.AnalyticsNumOutput, err error) {
	// 获取当日订单量
	stime := input.Stime
	etime := input.Etime

	//统计没有取消的订单
	orderStateId := []uint{consts.ORDER_STATE_WAIT_PAY, consts.ORDER_STATE_WAIT_PAID, consts.ORDER_STATE_WAIT_REVIEW, consts.ORDER_STATE_WAIT_FINANCE_REVIEW, consts.ORDER_STATE_PICKING, consts.ORDER_STATE_WAIT_SHIPPING, consts.ORDER_STATE_SHIPPED, consts.ORDER_STATE_RECEIVED, consts.ORDER_STATE_FINISH, consts.ORDER_STATE_SELF_PICKUP}
	paidStateId := []uint{consts.ORDER_PAID_STATE_NO, consts.ORDER_PAID_STATE_PART, consts.ORDER_PAID_STATE_YES}
	currentRegNum, err := dao.AnalyticsOrder.GetOrderNum(ctx, stime, etime, orderStateId, paidStateId, 0, 0)

	if err != nil {
		return topRes, err
	}

	topRes.Current = currentRegNum

	// 上个周期
	preInput := &model.AnalyticsOrderInput{}
	gconv.Struct(input, preInput)

	if input.Stime > 0 && input.Etime > 0 {
		preInput.Stime = input.Stime - (input.Etime - input.Stime)
		preInput.Etime = input.Stime

		stime = preInput.Stime
		etime = preInput.Etime

		preRegNum, _ := dao.AnalyticsOrder.GetOrderNum(ctx, stime, etime, orderStateId, paidStateId, 0, 0)

		if g.IsEmpty(preRegNum) {
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

// 获取订单金额
func (s *sAnalyticsOrder) GetOrderAmount(ctx context.Context, in *model.AnalyticsOrderInput) (*model.AnalyticsNumOutput, error) {
	topRes := &model.AnalyticsNumOutput{}

	// 统计没有取消的订单
	orderState := []uint{
		consts.ORDER_STATE_WAIT_PAY,
		consts.ORDER_STATE_WAIT_PAID,
		consts.ORDER_STATE_WAIT_REVIEW,
		consts.ORDER_STATE_WAIT_FINANCE_REVIEW,
		consts.ORDER_STATE_PICKING,
		consts.ORDER_STATE_WAIT_SHIPPING,
		consts.ORDER_STATE_SHIPPED,
		consts.ORDER_STATE_RECEIVED,
		consts.ORDER_STATE_FINISH,
		consts.ORDER_STATE_SELF_PICKUP,
	}

	paidState := []uint{
		consts.ORDER_PAID_STATE_PART,
		consts.ORDER_PAID_STATE_YES,
	}

	orderState = nil
	in.OrderStateId = orderState
	in.OrderIsPaid = paidState

	// 获取当前周期内数据
	currentRegNum, err := dao.AnalyticsOrder.GetOrderAmount(ctx, in.Stime, in.Etime, in.OrderStateId, in.OrderIsPaid, in.UserId, in.KindId)
	if err != nil {
		return nil, err
	}

	topRes.Current = currentRegNum

	// 上个周期
	preInput := &model.AnalyticsOrderInput{}
	gconv.Struct(in, preInput)

	if in.Stime > 0 && in.Etime > 0 {
		preInput.Stime = in.Stime - (in.Etime - in.Stime)
		preInput.Etime = in.Stime

		preRegNum, err := dao.AnalyticsOrder.GetOrderAmount(ctx, preInput.Stime, preInput.Etime, preInput.OrderStateId, preInput.OrderIsPaid, preInput.UserId, preInput.KindId)
		if err != nil {
			return nil, err
		}
		if preRegNum != nil {
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

// 获取订单时间线
func (s *sAnalyticsOrder) GetOrderNumTimeline(ctx context.Context, in *model.TimelineInput) (out []*model.TimelineOutput, err error) {
	orderTimeLine, err := dao.AnalyticsOrder.GetOrderTimeLine(ctx, in.Stime, in.Etime)

	if err != nil {
		return nil, err
	}

	gconv.Struct(orderTimeLine, &out)

	return
}

// 获取订单商品数量时间线
func (s *sAnalyticsOrder) GetOrderItemNumTimeLine(ctx context.Context, in *model.OrderItemNumTimelineInput) (out []*model.TimelineOutput, err error) {
	orderItemNumTimeLine, err := dao.AnalyticsOrder.GetOrderItemNumTimeLine(ctx, in)
	if err != nil {
		return nil, err
	}

	gconv.Struct(orderItemNumTimeLine, &out)

	return
}

// 获取订单商品数量
func (s *sAnalyticsOrder) GetOrderItemNum(ctx context.Context, in *model.OrderItemNumTimelineInput) (*model.AnalyticsNumOutput, error) {
	topRes := &model.AnalyticsNumOutput{}

	// 获取当前周期内数据
	currentProductNum, err := dao.AnalyticsOrder.GetOrderItemNum(ctx, in)
	if err != nil {
		return nil, err
	}
	if currentProductNum != nil {
		topRes.Current = currentProductNum
	}

	// 上个周期
	preInput := &model.OrderItemNumTimelineInput{}
	gconv.Struct(in, preInput)

	if in.Stime > 0 && in.Etime > 0 {
		preInput.Stime = in.Stime - (in.Etime - in.Stime)
		preInput.Etime = in.Stime

		preProductNum, err := dao.AnalyticsOrder.GetOrderItemNum(ctx, preInput)
		if err != nil {
			return nil, err
		}
		if preProductNum != nil {
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

// 列出订单商品数量
func (s *sAnalyticsOrder) ListOrderItemNum(ctx context.Context, in *model.OrderItemNumTimelineInput) (out []*model.AnalyticsOrderItemNumOutput, err error) {
	orderItemNumList, err := dao.AnalyticsOrder.ListOrderItemNum(ctx, in)
	if err != nil {
		return nil, err
	}

	gconv.Struct(orderItemNumList, &out)

	return
}
