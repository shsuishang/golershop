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
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/shopspring/decimal"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"sort"
	"strings"
)

type sConsumeDeposit struct{}

func init() {
	service.RegisterConsumeDeposit(NewConsumeDeposit())
}

func NewConsumeDeposit() *sConsumeDeposit {
	return &sConsumeDeposit{}
}

// Find 查询数据
func (s *sConsumeDeposit) Find(ctx context.Context, in *do.ConsumeDepositListInput) (out []*entity.ConsumeDeposit, err error) {
	out, err = dao.ConsumeDeposit.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sConsumeDeposit) List(ctx context.Context, in *do.ConsumeDepositListInput) (out *do.ConsumeDepositListOutput, err error) {
	out, err = dao.ConsumeDeposit.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sConsumeDeposit) Add(ctx context.Context, in *do.ConsumeDeposit) (lastInsertId int64, err error) {
	lastInsertId, err = dao.ConsumeDeposit.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sConsumeDeposit) Edit(ctx context.Context, in *do.ConsumeDeposit) (affected int64, err error) {
	_, err = dao.ConsumeDeposit.Edit(ctx, in.DepositId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sConsumeDeposit) Remove(ctx context.Context, id any) (affected int64, err error) {

	affected, err = dao.ConsumeDeposit.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// ProcessDeposit 新增
func (s *sConsumeDeposit) ProcessDeposit(ctx context.Context, in *do.ConsumeDeposit) (lastInsertId int64, err error) {
	//判断记录是否存在
	deposits, err := dao.ConsumeDeposit.Find(ctx, &do.ConsumeDepositListInput{Where: do.ConsumeDeposit{DepositNo: in.DepositNo, DepositTradeNo: in.DepositTradeNo}})

	if err != nil {
		return 0, err
	}

	var deposit *entity.ConsumeDeposit

	now := gtime.Now()

	//开启事务
	err = dao.ConsumeDeposit.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {

		if len(deposits) > 0 {
			deposit = deposits[0]
		} else {
			lastInsertId, err = dao.ConsumeDeposit.Add(ctx, in)
			if err != nil {
				return err
			}

			deposit, err = dao.ConsumeDeposit.Get(ctx, lastInsertId)

			if err != nil {
				return err
			}
		}

		if deposit.DepositState == 0 {
			trades, err := dao.ConsumeTrade.Find(ctx, &do.ConsumeTradeListInput{Where: do.ConsumeTrade{OrderId: gstr.Split(deposit.OrderId, ",")}})

			if err != nil {
				return err
			}

			var trade *entity.ConsumeTrade

			if len(trades) > 0 {
				trade = trades[0]
			}

			//用户账户增加充值额度
			_, err = dao.UserResource.Increment(ctx, trade.BuyerId, dao.UserResource.Columns().UserMoney, deposit.DepositTotalFee)
			if err != nil {
				return err
			}

			//写入充值流水
			record := do.ConsumeRecord{}
			record.OrderId = trade.OrderId
			record.UserId = trade.BuyerId
			record.UserNickname = ""
			record.StoreId = trade.BuyerStoreId
			record.ChainId = 0
			record.RecordTotal = deposit.DepositTotalFee
			record.RecordMoney = deposit.DepositTotalFee
			record.RecordDate = now
			record.RecordYear = now.Format("Y")
			record.RecordMonth = now.Format("n")
			record.RecordDay = now.Format("j")
			record.RecordTitle = deposit.DepositSubject
			record.RecordDesc = deposit.DepositBody
			record.RecordTime = now.UnixMilli()
			record.TradeTypeId = consts.TRADE_TYPE_DEPOSIT
			record.PaymentMetId = consts.PAYMENT_MET_MONEY
			record.PaymentTypeId = deposit.DepositPaymentType
			record.PaymentChannelId = deposit.PaymentChannelId

			_, err = dao.ConsumeRecord.Add(ctx, &record)
			if err != nil {
				return err
			}
			//end 写入充值流水

			//修改充值成功状态
			_, err = dao.ConsumeDeposit.Edit(ctx, deposit.DepositId, &do.ConsumeDeposit{DepositState: 1})
			if err != nil {
				return err
			}

			//处理订单支付结果
			payInfo := model.PayMetVo{}
			payInfo.PaymentMetId = consts.PAYMENT_MET_MONEY
			payInfo.PaymentChannelId = deposit.PaymentChannelId
			payInfo.PaymentTypeId = deposit.DepositPaymentType
			payInfo.PmMoney = deposit.DepositTotalFee

			_, err = service.ConsumeTrade().ProcessPay(ctx, deposit.OrderId, payInfo)

			if err != nil {
				return err
			}

			/*
				depositTotalFee := big.NewFloat(deposit.DepositTotalFee)

				for _, trade := range trades {
					orderId := trade.OrderId

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

							_, err := dao.ConsumeTrade.Edit(ctx, trade.ConsumeTradeId, &tradeData)

							if err != nil {
								return err
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

							depositTotalFee.Sub(depositTotalFee, tradePaymentAmount)
						} else {
							tradePaymentAmount = depositTotalFee

							//订单处理
							//不够支付完成
							tradeData.TradeIsPaid = consts.ORDER_PAID_STATE_PART
							//tradeData.PaymentChannelId = deposit.PaymentChannelId
							tradeData.TradePaymentAmount, _ = big.NewFloat(trade.TradePaymentAmount).Sub(big.NewFloat(trade.TradePaymentAmount), tradePaymentAmount).Float64()
							tradeData.TradePaymentMoney, _ = big.NewFloat(trade.TradePaymentMoney).Add(big.NewFloat(trade.TradePaymentMoney), tradePaymentAmount).Float64()

							_, err := dao.ConsumeTrade.Edit(ctx, trade.ConsumeTradeId, &tradeData)

							if err != nil {
								return err
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
								return err
							}

							f, _ := tradePaymentAmount.Float64()
							_, err = dao.UserResource.Decrement(ctx, trade.BuyerId, dao.UserResource.Columns().UserMoney, f)

							if err != nil {
								return err
							}

							//2. 卖家订单流水增加
							record.UserId = trade.SellerId
							record.UserNickname = ""
							record.StoreId = trade.StoreId
							record.ChainId = trade.ChainId
							record.TradeTypeId = consts.TRADE_TYPE_SALES
							record.PaymentTypeId = deposit.DepositPaymentType
							record.PaymentChannelId = deposit.PaymentChannelId
							record.RecordTotal, _ = tradePaymentAmount.Float64()

							//卖家收益涉及佣金问题， 可以分多次付款，支付完成才扣佣金
							if tradeData.TradeIsPaid == consts.ORDER_PAID_STATE_YES {
								//卖家收益，进入冻结中?
								if consts.PAYMENT_TYPE_OFFLINE == deposit.DepositPaymentType {
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
								return err
							}

							//卖家收益，进入冻结中?
							if consts.PAYMENT_TYPE_OFFLINE == deposit.DepositPaymentType {
								//线下支付，需要扣除商家交易佣金？
							} else {
								_, err = dao.UserResource.Increment(ctx, trade.SellerId, dao.UserResource.Columns().UserMoney, record.RecordMoney)
								if err != nil {
									return err
								}
							}
						}
					} else {
						tradeData.TradeIsPaid = consts.ORDER_PAID_STATE_YES
					}

					if consts.TRADE_TYPE_DEPOSIT == tradeData.TradeTypeId {
						//$flag_row[] = $this->notifyDeposit($order_id, $trade_row);
					}

					if consts.ORDER_PAID_STATE_YES == tradeData.TradeIsPaid {
						_, err = service.Order().SetPaidYes(ctx, orderId)
						if err != nil {
							return err
						}
					}
				}
			*/
		} else {
			//只说明本次充值已经操作完成
			/*
				for _, orderId := range gstr.Split(deposit.OrderId, ",") {
					_, err = service.Order().SetPaidYes(ctx, orderId)

					if err != nil {
						return err
					}
				}
			*/
		}

		return nil
	})

	return lastInsertId, err
}

// GetTradeNo 获取交易编号
func (s *sConsumeDeposit) GetTradeNo(ctx context.Context, orderIdList []string) (tradeNo string, err error) {

	if len(orderIdList) > 1 {
		sort.Slice(orderIdList, func(i, j int) bool {
			return orderIdList[i] < orderIdList[j]
		})

		orderIds := gstr.JoinAny(orderIdList, ",")
		tradeNo, _ = gmd5.Encrypt(orderIds)
		tradeNo = "TC-" + tradeNo

		_, err = dao.ConsumeCombine.Save(ctx, &do.ConsumeCombine{CtcId: tradeNo, OrderIds: orderIds})
	} else {
		tradeNo = orderIdList[0]
	}

	return
}

// GetOrderId 根据交易编号 获取 订单编号
func (s *sConsumeDeposit) GetOrderId(ctx context.Context, tradeNo string) (orderIds string, err error) {
	consumeCombine, err := dao.ConsumeCombine.Get(ctx, tradeNo)
	if err != nil {
		return "", err
	}

	if consumeCombine != nil {
		orderIds = consumeCombine.OrderIds
	} else {
		orderIds = tradeNo
	}

	return
}

// WechatApplet 微信小程序
func (s *sConsumeDeposit) WechatApplet(ctx context.Context, in *do.ConsumeDeposit) (lastInsertId int64, err error) {

	return
}

// WechatApp App
func (s *sConsumeDeposit) WechatApp(ctx context.Context, in *do.ConsumeDeposit) (lastInsertId int64, err error) {

	return
}

// WechatJSAPI 微信JSAPI
func (s *sConsumeDeposit) WechatJSAPI(ctx context.Context, in *do.ConsumeDeposit) (lastInsertId int64, err error) {

	return
}

// OfflinePay 离线支付
func (s *sConsumeDeposit) OfflinePay(ctx context.Context, in *do.ConsumeDeposit) (lastInsertId int64, err error) {
	if g.IsEmpty(in.DepositSubject) {
		orderInfo, _ := dao.OrderInfo.Get(ctx, in.OrderId)
		//in.DepositSubject = in.DepositNo
		in.DepositSubject = orderInfo.OrderTitle
	}

	lastInsertId, err = s.ProcessDeposit(ctx, in)

	if err != nil {
		return 0, err
	}

	return
}

// 获取业务返回数据
func (s *sConsumeDeposit) GetPayResult(ctx context.Context, in *model.PaymentInput) (*model.PaymentOutput, error) {
	consumeTrades, err := s.checkTrade(ctx, in)
	if err != nil {
		return nil, err
	}

	// 订单id
	// 是否为联合支付
	tradeNo, err := s.GetTradeNo(ctx, in.OrderId)

	// 标题
	var titles []string
	for _, v := range consumeTrades {
		titles = append(titles, v.TradeTitle)
	}
	title := strings.Join(titles, ",")

	// 付款金额
	var tradePaymentAmount decimal.Decimal
	for _, v := range consumeTrades {
		tradePaymentAmount = tradePaymentAmount.Add(decimal.NewFromFloat(v.TradePaymentAmount))
	}

	if tradePaymentAmount.Cmp(decimal.Zero) <= 0 {
		return nil, errors.New("无需支付！")
	}

	// 设置订单金额 单位为分且最小为1
	amount, _ := tradePaymentAmount.Mul(decimal.NewFromInt(100)).Float64()

	return &model.PaymentOutput{
		TradeNo: tradeNo,
		Title:   title,
		Amount:  amount,
	}, nil
}

func (s *sConsumeDeposit) checkTrade(ctx context.Context, input *model.PaymentInput) ([]*entity.ConsumeTrade, error) {
	tradeQueryWrapper := &do.ConsumeTradeListInput{Where: do.ConsumeTrade{
		OrderId: input.OrderId,
	}}
	consumeTrades, err := dao.ConsumeTrade.Find(ctx, tradeQueryWrapper)
	if err != nil {
		return nil, err
	}
	if len(consumeTrades) == 0 {
		return nil, errors.New("交易订单不存在")
	}

	// 判断是否可以联合支付

	// 是否包含已付款
	for _, trade := range consumeTrades {
		if trade.TradeIsPaid == consts.ORDER_PAID_STATE_YES {
			return nil, errors.New(fmt.Sprintf("%s 订单状态不为待付款状态！", trade.OrderId))
		}
	}

	return consumeTrades, nil
}
