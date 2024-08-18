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

package pay

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"math/big"
)

type sConsumeTrade struct{}

func init() {
	service.RegisterConsumeTrade(NewConsumeTrade())
}

func NewConsumeTrade() *sConsumeTrade {
	return &sConsumeTrade{}
}

// Find 查询数据
func (s *sConsumeTrade) Find(ctx context.Context, in *do.ConsumeTradeListInput) (out []*entity.ConsumeTrade, err error) {
	out, err = dao.ConsumeTrade.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sConsumeTrade) List(ctx context.Context, in *do.ConsumeTradeListInput) (out *do.ConsumeTradeListOutput, err error) {
	out, err = dao.ConsumeTrade.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sConsumeTrade) Add(ctx context.Context, in *do.ConsumeTrade) (lastInsertId int64, err error) {
	lastInsertId, err = dao.ConsumeTrade.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sConsumeTrade) Edit(ctx context.Context, in *do.ConsumeTrade) (affected int64, err error) {
	_, err = dao.ConsumeTrade.Edit(ctx, in.ConsumeTradeId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sConsumeTrade) Remove(ctx context.Context, id any) (affected int64, err error) {

	affected, err = dao.ConsumeTrade.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// ProcessPay 余额支付
func (s *sConsumeTrade) ProcessPay(ctx context.Context, ids string, deposit model.PayMetVo) (out model.ProcessPayOutput, err error) {
	out.OrderId = ids

	trades, err := dao.ConsumeTrade.Find(ctx, &do.ConsumeTradeListInput{Where: do.ConsumeTrade{OrderId: gstr.Split(ids, ",")}})
	if err != nil {
		return out, err
	}

	now := gtime.Now()
	depositTotalFee := big.NewFloat(0)

	switch deposit.PaymentMetId {
	case consts.PAYMENT_MET_MONEY:
		depositTotalFee = big.NewFloat(deposit.PmMoney)
	case consts.PAYMENT_MET_POINTS:
		depositTotalFee = big.NewFloat(deposit.PmPoints)
	case consts.PAYMENT_MET_CREDIT:
		depositTotalFee = big.NewFloat(deposit.PmCredit)
	case consts.PAYMENT_MET_RECHARGE_CARD:
		depositTotalFee = big.NewFloat(deposit.PmRechargeCard)
	default:
		return out, gerror.New("支付渠道不合法")
	}

	if true {
		for _, trade := range trades {
			orderId := trade.OrderId

			//写入充值流水
			record := do.ConsumeRecord{}
			record.OrderId = trade.OrderId
			record.UserId = trade.BuyerId
			record.UserNickname = ""
			record.StoreId = trade.BuyerStoreId
			record.ChainId = 0
			record.RecordTotal = depositTotalFee
			record.RecordMoney = depositTotalFee
			record.RecordDate = now
			record.RecordYear = now.Format("Y")
			record.RecordMonth = now.Format("n")
			record.RecordDay = now.Format("j")
			record.RecordTitle = trade.TradeTitle
			record.RecordDesc = trade.TradeDesc
			record.RecordTime = now.UnixMilli()
			record.TradeTypeId = consts.TRADE_TYPE_DEPOSIT
			record.PaymentMetId = consts.PAYMENT_MET_MONEY
			record.PaymentTypeId = deposit.PaymentTypeId
			record.PaymentChannelId = deposit.PaymentChannelId

			var tradeData do.ConsumeTrade

			if depositTotalFee.Cmp(big.NewFloat(0)) > 0 && trade.TradeIsPaid != consts.ORDER_PAID_STATE_YES {
				//trade.SellerId

				//当前订单需要支付额度
				tradePaymentAmount := big.NewFloat(trade.TradePaymentAmount)

				if depositTotalFee.Cmp(tradePaymentAmount) >= 0 {
					//订单处理
					//更改订单状态, 可以完成订单支付状态

					tradeData.TradeIsPaid = consts.ORDER_PAID_STATE_YES
					tradeData.PaymentChannelId = deposit.PaymentChannelId
					tradeData.TradePaymentAmount = 0
					tradeData.TradePaymentMoney, _ = big.NewFloat(trade.TradePaymentMoney).Add(big.NewFloat(trade.TradePaymentMoney), tradePaymentAmount).Float64()
					tradeData.TradePaidTime = now.UnixMilli()

					_, err := dao.ConsumeTrade.Edit(ctx, trade.ConsumeTradeId, &tradeData)

					if err != nil {
						return out, err
					}

					////不是充值订单, 订单支付完成
					//if (StateCode::TRADE_TYPE_SHOPPING == $trade_row['trade_type_id'])
					//{
					//$paid_order_id_row[] = $order_id;
					//}
					//
					//if (StateCode::TRADE_TYPE_XQ_BUY == $trade_row['trade_type_id'])
					//{
					//$paid_order_id_service_row[] = $order_id;
					//}

					depositTotalFee = depositTotalFee.Sub(depositTotalFee, tradePaymentAmount)
				} else {
					tradePaymentAmount = depositTotalFee

					//订单处理
					//不够支付完成
					tradeData.TradeIsPaid = consts.ORDER_PAID_STATE_PART
					//tradeData.PaymentChannelId = deposit.PaymentChannelId
					tradeData.TradePaymentAmount, _ = big.NewFloat(trade.TradePaymentAmount).Sub(big.NewFloat(trade.TradePaymentAmount), tradePaymentAmount).Float64()
					tradeData.TradePaymentMoney, _ = big.NewFloat(trade.TradePaymentMoney).Add(big.NewFloat(trade.TradePaymentMoney), tradePaymentAmount).Float64()
					tradeData.TradePaidTime = now.UnixMilli()
					_, err := dao.ConsumeTrade.Edit(ctx, trade.ConsumeTradeId, &tradeData)

					if err != nil {
						return out, err
					}

					depositTotalFee = big.NewFloat(0)
				}

				//订单扣除流水
				//订单消费流水
				//涉及佣金结算问题
				if consts.TRADE_TYPE_SHOPPING == trade.TradeTypeId {

					//1. 买家流水及订单扣除
					record.RecordTitle = trade.TradeTitle
					record.UserId = trade.BuyerId
					record.UserNickname = ""
					record.StoreId = trade.BuyerStoreId
					record.ChainId = 0
					record.TradeTypeId = trade.TradeTypeId
					record.RecordTotal, _ = big.NewFloat(0).Sub(big.NewFloat(0), tradePaymentAmount).Float64()
					record.RecordMoney = record.RecordTotal

					_, err = dao.ConsumeRecord.Add(ctx, &record)

					if err != nil {
						return out, err
					}

					f, _ := tradePaymentAmount.Float64()
					_, err = dao.UserResource.Decrement(ctx, trade.BuyerId, dao.UserResource.Columns().UserMoney, f)

					if err != nil {
						return out, err
					}

					//2. 卖家订单流水增加
					record.UserId = trade.SellerId
					record.UserNickname = ""
					record.StoreId = trade.StoreId
					record.ChainId = trade.ChainId
					record.TradeTypeId = consts.TRADE_TYPE_SALES
					record.PaymentTypeId = deposit.PaymentTypeId
					record.PaymentChannelId = deposit.PaymentChannelId
					record.RecordTotal, _ = tradePaymentAmount.Float64()

					//卖家收益涉及佣金问题， 可以分多次付款，支付完成才扣佣金
					if tradeData.TradeIsPaid == consts.ORDER_PAID_STATE_YES {
						//卖家收益，进入冻结中?
						if consts.PAYMENT_TYPE_OFFLINE == deposit.PaymentTypeId {
							record.RecordMoney = record.RecordTotal //佣金平台获取。 是否需要加入一个统计字段中？
							record.RecordCommissionFee = 0
						} else {
							record.RecordMoney, _ = big.NewFloat(0).Sub(tradePaymentAmount, big.NewFloat(trade.OrderCommissionFee)).Float64() //佣金平台获取。 是否需要加入一个统计字段中？
							record.RecordCommissionFee = trade.OrderCommissionFee                                                             //佣金平台获取

							//平台佣金总额
							//$plantform_resource_row = array();
							//$plantform_resource_row['plantform_resource_id'] = DATA_ID;
							//$plantform_resource_row['plantform_commission_fee'] = $trade_row['order_commission_fee'];
							//$flag_row[] = Plantform_ResourceModel::getInstance()->save($plantform_resource_row, true, true);

						}
					} else {
						record.RecordMoney, _ = tradePaymentAmount.Float64()
					}

					_, err = dao.ConsumeRecord.Add(ctx, &record)

					if err != nil {
						return out, err
					}

					//卖家收益，进入冻结中?
					if consts.PAYMENT_TYPE_OFFLINE == deposit.PaymentTypeId {
						//线下支付，需要扣除商家交易佣金？
					} else {
						_, err = dao.UserResource.Increment(ctx, trade.SellerId, dao.UserResource.Columns().UserMoney, record.RecordMoney)
						if err != nil {
							return out, err
						}
					}
				}
			} else {
				tradeData.TradeIsPaid = consts.ORDER_PAID_STATE_YES
			}

			if consts.TRADE_TYPE_DEPOSIT == trade.TradeTypeId {
				//$flag_row[] = $this->notifyDeposit($order_id, $trade_row);
			}

			if consts.ORDER_PAID_STATE_YES == tradeData.TradeIsPaid.(uint) {
				orderInfo, _ := dao.OrderInfo.Get(ctx, orderId)

				if orderInfo.OrderStateId == consts.ORDER_STATE_WAIT_PAY {
					_, err = service.Order().SetPaidYes(ctx, orderId)
					if err != nil {
						return out, err
					}
					out.Paid = true
				} else {
					if orderInfo.PaymentTypeId == consts.PAYMENT_TYPE_OFFLINE {

					}

					//判断是否线下支付
					if consts.PAYMENT_TYPE_OFFLINE == deposit.PaymentTypeId {
                        //直接处理订单支付状态， 不处理订单状态
						_, err = dao.OrderInfo.Edit(ctx, orderId, &do.OrderInfo{OrderIsPaid: consts.ORDER_PAID_STATE_YES})

						if err != nil {
							return out, err
						}

						out.Paid = true
					} else {
						_, err = service.Order().SetPaidYes(ctx, orderId)
						if err != nil {
							return out, err
						}
						out.Paid = true
					}
				}
			} else if consts.ORDER_PAID_STATE_PART == tradeData.TradeIsPaid.(uint) {
				_, err = dao.OrderInfo.Edit(ctx, orderId, &do.OrderInfo{OrderIsPaid: consts.ORDER_PAID_STATE_PART})
				if err != nil {
					return out, err
				}

				out.Paid = false
			}
		}
	}

	return
}
