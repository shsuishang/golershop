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

package internal

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AnalyticsTradeDao is the data access object for table trade_order_info.
type AnalyticsTradeDao struct {
	table string // table is the underlying table name of the DAO.
	group string // group is the database configuration group name of current DAO
}

// NewAnalyticsTradeDao creates and returns a new DAO object for table data access.
func NewAnalyticsTradeDao() *AnalyticsTradeDao {
	return &AnalyticsTradeDao{
		group: "pay",
		table: "pay_consume_trade",
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AnalyticsTradeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Group returns the configuration group name of database of current dao.
func (dao *AnalyticsTradeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AnalyticsTradeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// TradeAmount 交易总额
func (dao *AnalyticsTradeDao) SalesAmount(ctx context.Context, start int64, end int64, buyerId int64) (res interface{}, err error) {

	/*sql := fmt.Sprintf("select sum(order_payment_amount) as amount from pay_consume_trade where trade_paid_time BETWEEN ? AND ? AND trade_is_paid IN(%d, %d) AND trade_type_id IN(%d, %d)", consts.ORDER_PAID_STATE_PART, consts.ORDER_PAID_STATE_YES, consts.TRADE_TYPE_SHOPPING, consts.TRADE_TYPE_FAVORABLE)

	one, err := dao.DB().GetOne(ctx, sql, start, end)

	if err != nil {
		return nil, err
	}

	if one["amount"] != nil {
		res = one["amount"]
	} else {
		res = 0
	}*/

	whereSet := ""

	whereSet = fmt.Sprintf(" AND trade_is_paid IN ( 3012, 3013 )")

	if !g.IsEmpty(start) && !g.IsEmpty(end) {
		whereSet = fmt.Sprintf(" AND trade_paid_time BETWEEN %d AND %d", start, end)
	}

	whereSet = whereSet + fmt.Sprintf(" AND trade_type_id IN ( 1201, 1214 )")

	if !g.IsEmpty(buyerId) {
		whereSet = whereSet + fmt.Sprintf(" AND buyer_id = %d", buyerId)
	}

	sql := fmt.Sprintf(`
		       SELECT
		        	sum(order_payment_amount) as amount
		        FROM
		        	pay_consume_trade
		        WHERE 1 %s `, whereSet)

	one, err := dao.DB().GetOne(ctx, sql)

	if err != nil {
		return res, err
	}

	if one["amount"] != nil {
		res = one["amount"]
	} else {
		res = 0
	}

	return
}
