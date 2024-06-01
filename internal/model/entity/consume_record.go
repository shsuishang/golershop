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

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ConsumeRecord is the golang structure for table consume_record.
type ConsumeRecord struct {
	ConsumeRecordId                 int64       `json:"consume_record_id"                  ` // 支付流水号
	OrderId                         string      `json:"order_id"                           ` // 商户订单编号
	UserId                          uint        `json:"user_id"                            ` // 所属用编号
	UserNickname                    string      `json:"user_nickname"                      ` // 昵称
	CurrencyId                      uint        `json:"currency_id"                        ` // 货币编号
	CurrencySymbolLeft              string      `json:"currency_symbol_left"               ` // 左符号
	RecordTotal                     float64     `json:"record_total"                       ` // 金额
	RecordMoney                     float64     `json:"record_money"                       ` // 金额:record_total-佣金
	RecordCommissionFee             float64     `json:"record_commission_fee"              ` // 佣金:平台佣金针对销售收款
	RecordDistributionCommissionFee float64     `json:"record_distribution_commission_fee" ` // 分销佣金:针对销售收款
	RecordDate                      *gtime.Time `json:"record_date"                        ` // 年-月-日
	RecordYear                      uint        `json:"record_year"                        ` // 年
	RecordMonth                     uint        `json:"record_month"                       ` // 月
	RecordDay                       uint        `json:"record_day"                         ` // 日
	RecordTitle                     string      `json:"record_title"                       ` // 标题
	RecordDesc                      string      `json:"record_desc"                        ` // 描述
	RecordTime                      int64       `json:"record_time"                        ` // 支付时间
	TradeTypeId                     uint        `json:"trade_type_id"                      ` // 交易类型(ENUM):1201-购物; 1202-转账; 1203-充值; 1204-提现; 1205-销售; 1206-佣金; 1207-退货付款;1208-退货收款;1209-转账收款
	PaymentTypeId                   uint        `json:"payment_type_id"                    ` // 支付方式(ENUM):1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付;
	PaymentChannelId                uint        `json:"payment_channel_id"                 ` // 支付渠道
	StoreId                         uint        `json:"store_id"                           ` // 所属店铺
	ChainId                         uint        `json:"chain_id"                           ` // 所属门店
	PaymentMetId                    uint        `json:"payment_met_id"                     ` // 消费类型(ENUM):1-余额支付; 2-充值卡支付; 3-积分支付; 4-信用支付; 5-红包支付
	RecordEnable                    bool        `json:"record_enable"                      ` // 状态(BOOL):1-已收款;0-作废
	SubsiteId                       uint        `json:"subsite_id"                         ` // 所属分站:0-总站
	Version                         uint        `json:"version"                            ` // 版本
}
