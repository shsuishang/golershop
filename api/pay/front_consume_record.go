package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type RecordListReq struct {
	g.Meta `path:"/front/pay/consumeRecord/list" tags:"交易明细" method:"get" summary:"交易明细表-账户收支明细-资金流水表-账户金额变化流水分页查询"`

	ChangeType                      uint       `json:"change_type" dc:"支出收入:1-支出;2-收入"`                                                                                     // 支出收入:1-支出;2-收入
	ConsumeRecordId                 int64      `json:"consume_record_id" dc:"支付流水号"`                                                                                        // 支付流水号
	OrderId                         string     `json:"order_id" dc:"商户订单编号"`                                                                                                // 商户订单编号
	UserId                          uint       `json:"user_id" dc:"所属用户编号"`                                                                                                 // 所属用户编号
	UserNickname                    string     `json:"user_nickname" dc:"昵称"`                                                                                               // 昵称
	CurrencyId                      int        `json:"currency_id" dc:"货币编号"`                                                                                               // 货币编号
	CurrencySymbolLeft              string     `json:"currency_symbol_left" dc:"左符号"`                                                                                       // 左符号
	RecordTotal                     float64    `json:"record_total" dc:"金额"`                                                                                                // 金额
	RecordMoney                     float64    `json:"record_money" dc:"金额:record_total-佣金"`                                                                                // 金额:record_total-佣金
	RecordCommissionFee             float64    `json:"record_commission_fee" dc:"佣金:平台佣金针对销售收款"`                                                                            // 佣金:平台佣金针对销售收款
	RecordDistributionCommissionFee float64    `json:"record_distribution_commission_fee" dc:"分销佣金:针对销售收款"`                                                                 // 分销佣金:针对销售收款
	RecordDate                      gtime.Time `json:"record_date" dc:"年-月-日"`                                                                                              // 年-月-日
	RecordYear                      int        `json:"record_year" dc:"年"`                                                                                                  // 年
	RecordMonth                     int        `json:"record_month" dc:"月"`                                                                                                 // 月
	RecordDay                       int        `json:"record_day" dc:"日"`                                                                                                   // 日
	RecordTitle                     string     `json:"record_title" dc:"标题"`                                                                                                // 标题
	RecordDesc                      string     `json:"record_desc" dc:"描述"`                                                                                                 // 描述
	RecordTime                      int64      `json:"record_time" dc:"支付时间"`                                                                                               // 支付时间
	TradeTypeId                     []uint     `json:"trade_type_id" dc:"交易类型(ENUM):1201-购物; 1202-转账; 1203-充值; 1204-提现; 1205-销售; 1206-佣金; 1207-退货付款; 1208-退货收款; 1209-转账收款"` // 交易类型
	PaymentTypeId                   int        `json:"payment_type_id" dc:"支付方式(ENUM):1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付;"`                              // 支付方式
	PaymentChannelId                int        `json:"payment_channel_id" dc:"支付渠道"`                                                                                        // 支付渠道
	StoreId                         int        `json:"store_id" dc:"所属店铺"`                                                                                                  // 所属店铺
	ChainId                         int        `json:"chain_id" dc:"所属门店"`                                                                                                  // 所属门店
	PaymentMetId                    int        `json:"payment_met_id" dc:"消费类型(ENUM):1-余额支付; 2-充值卡支付; 3-积分支付; 4-信用支付; 5-红包支付"`                                              // 消费类型
	RecordEnable                    bool       `json:"record_enable" dc:"状态(BOOL):1-已收款;0-作废"`                                                                              // 状态
	SubsiteId                       int        `json:"subsite_id" dc:"所属分站:0-总站"`                                                                                           // 所属分站
	RecordTimeStart                 int64      `json:"record_time_start" dc:"支付时间-开始"`                                                                                      // 支付时间-开始
	RecordTimeEnd                   int64      `json:"record_time_end" dc:"支付时间-结束"`                                                                                        // 支付时间-结束
}
type RecordListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
