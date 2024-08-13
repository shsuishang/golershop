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
	"github.com/gogf/gf/v2/text/gstr"
	"golershop.cn/internal/model"
)

// AnalyticsOrderDao is the data access object for table trade_order_info.
type AnalyticsOrderDao struct {
	table string // table is the underlying table name of the DAO.
	group string // group is the database configuration group name of current DAO
}

// NewAnalyticsOrderDao creates and returns a new DAO object for table data access.
func NewAnalyticsOrderDao() *AnalyticsOrderDao {
	return &AnalyticsOrderDao{
		group: "trade",
		table: "trade_order_info",
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AnalyticsOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Group returns the configuration group name of database of current dao.
func (dao *AnalyticsOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AnalyticsOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// GetOrderTimeLine 查询字段数据
func (dao *AnalyticsOrderDao) GetOrderTimeLine(ctx context.Context, stime, etime int64) (out gdb.Result, err error) {
	whereSet := ""
	if !g.IsEmpty(stime) && !g.IsEmpty(etime) {
		whereSet = fmt.Sprintf(" AND create_time BETWEEN %d AND %d", stime, etime)
	}

	sql := fmt.Sprintf(` SELECT
			FROM_UNIXTIME( create_time / 1000, '%%m-%%d' ) AS time,
			count(*) AS num
        FROM
        	trade_order_info
		WHERE 1 %s
        GROUP BY time
        ORDER BY create_time`, whereSet)

	out, err = dao.DB().GetAll(ctx, sql)

	if err != nil {
		return out, err
	}

	return out, nil
}

// GetPayTimeLine 查询字段数据
func (dao *AnalyticsOrderDao) GetPayTimeLine(ctx context.Context, stime, etime int64) (out gdb.Result, err error) {
	whereSet := ""
	if !g.IsEmpty(stime) && !g.IsEmpty(etime) {
		whereSet = fmt.Sprintf(" AND trade_paid_time BETWEEN %d AND %d", stime, etime)
	}

	sql := fmt.Sprintf(`         SELECT
        FROM_UNIXTIME(trade_paid_time / 1000, '%%m-%%d') AS time,
        sum( order_payment_amount ) num
        FROM
        pay_consume_trade
		WHERE 1 %s
        GROUP BY time
        ORDER BY trade_paid_time`, whereSet)

	out, err = dao.DB().GetAll(ctx, sql)

	if err != nil {
		return out, err
	}

	return out, nil

}

// GetSaleOrderAmount 查询字段数据
func (dao *AnalyticsOrderDao) GetSaleOrderAmount(ctx context.Context, stime, etime int64) (out gdb.Result, err error) {
	whereSet := ""
	if !g.IsEmpty(stime) && !g.IsEmpty(etime) {
		whereSet = fmt.Sprintf(" AND trade_order_info.create_time BETWEEN %d AND %d", stime, etime)
	}

	sql := fmt.Sprintf(`
       SELECT
        FROM_UNIXTIME( trade_order_info.create_time / 1000, '%%m-%%d' ) AS time,
        sum(trade_order_base.order_payment_amount) AS num
        FROM
        trade_order_info left join trade_order_base ON trade_order_info.order_id = trade_order_base.order_id
		WHERE 1 %s AND trade_order_info.order_is_paid IN (3012, 3013)
        GROUP BY time
        ORDER BY trade_order_info.create_time`, whereSet)

	out, err = dao.DB().GetAll(ctx, sql)

	if err != nil {
		return out, err
	}

	return out, nil
}

// GetOrderCustomerNumTimeline 查询字段数据
func (dao *AnalyticsOrderDao) GetOrderCustomerNumTimeline(ctx context.Context, stime, etime int64) (out gdb.Result, err error) {
	whereSet := ""
	if !g.IsEmpty(stime) && !g.IsEmpty(etime) {
		whereSet = fmt.Sprintf(" AND create_time BETWEEN %d AND %d", stime, etime)
	}

	sql := fmt.Sprintf(`
       SELECT
            FROM_UNIXTIME( create_time / 1000, '%%m-%%d' ) AS time,
            count( DISTINCT ( user_id ) ) AS num
        FROM
            trade_order_info
        WHERE 1 %s 
        GROUP BY time
        ORDER BY create_time`, whereSet)

	out, err = dao.DB().GetAll(ctx, sql)

	if err != nil {
		return out, err
	}

	return out, nil
}

// GetOrderNum 查询字段数据
func (dao *AnalyticsOrderDao) GetOrderNum(ctx context.Context, stime, etime int64, orderStateId []uint, orderIsPaid []uint, userId uint, kindId uint) (out interface{}, err error) {
	whereSet := ""
	if !g.IsEmpty(stime) && !g.IsEmpty(etime) {
		whereSet = fmt.Sprintf(" AND create_time BETWEEN %d AND %d", stime, etime)
	}

	if !g.IsEmpty(orderStateId) {
		whereSet = whereSet + fmt.Sprintf(" AND order_state_id IN (%s)", gstr.JoinAny(orderStateId, ","))
	}

	if !g.IsEmpty(userId) {
		whereSet = whereSet + fmt.Sprintf(" AND user_id = %d", userId)
	}

	if !g.IsEmpty(kindId) {
		whereSet = whereSet + fmt.Sprintf(" AND kind_id = %d", kindId)
	}

	if !g.IsEmpty(orderIsPaid) {
		whereSet = whereSet + fmt.Sprintf(" AND order_is_paid IN (%s)", gstr.JoinAny(orderIsPaid, ","))
	}

	sql := fmt.Sprintf(`
       SELECT
            COUNT( * ) AS num
        FROM
            trade_order_info
        WHERE 1 %s `, whereSet)

	one, err := dao.DB().GetOne(ctx, sql)

	if err != nil {
		return out, err
	}

	if one["num"] != nil {
		out = one["num"]
	} else {
		out = 0
	}

	return
}

// GetOrderNum 查询字段数据
func (dao *AnalyticsOrderDao) GetOrderAmount(ctx context.Context, stime, etime int64, orderStateId []uint, orderIsPaid []uint, userId uint, kindId uint) (out interface{}, err error) {
	whereSet := ""
	if !g.IsEmpty(stime) && !g.IsEmpty(etime) {
		whereSet = fmt.Sprintf(" AND trade_order_info.create_time BETWEEN %d AND %d", stime, etime)
	}

	if !g.IsEmpty(orderStateId) {
		whereSet = whereSet + fmt.Sprintf(" AND trade_order_info.order_state_id IN (%s)", gstr.JoinAny(orderStateId, ","))
	}

	if !g.IsEmpty(userId) {
		whereSet = whereSet + fmt.Sprintf(" AND trade_order_info.user_id = %d", userId)
	}

	if !g.IsEmpty(kindId) {
		whereSet = whereSet + fmt.Sprintf(" AND trade_order_info.kind_id = %d", kindId)
	}

	if !g.IsEmpty(orderIsPaid) {
		whereSet = whereSet + fmt.Sprintf(" AND trade_order_info.order_is_paid IN (%s)", gstr.JoinAny(orderIsPaid, ","))
	}

	sql := fmt.Sprintf(`
        SELECT
        sum(trade_order_base.order_payment_amount) AS num
        FROM
        trade_order_info left join trade_order_base ON trade_order_info.order_id = trade_order_base.order_id
        WHERE 1 %s AND trade_order_info.order_is_paid IN (3012, 3013) `, whereSet)

	one, err := dao.DB().GetOne(ctx, sql)

	if err != nil {
		return out, err
	}

	if one["num"] != nil {
		out = one["num"]
	} else {
		out = 0
	}

	return
}

// GetOrderItemNumTimeLine 查询字段数据
func (dao *AnalyticsOrderDao) GetOrderItemNumTimeLine(ctx context.Context, in *model.OrderItemNumTimelineInput) (out gdb.Result, err error) {
	whereSet := ""
	if !g.IsEmpty(in.Stime) && !g.IsEmpty(in.Etime) {
		whereSet = fmt.Sprintf(" AND b.create_time BETWEEN %d AND %d", in.Stime, in.Etime)
	}

	if !g.IsEmpty(in.StoreId) {
		whereSet = fmt.Sprintf(" AND i.store_id =  %d", in.StoreId)
	}

	if !g.IsEmpty(in.ProductId) {
		whereSet = fmt.Sprintf(" AND i.product_id =  %d", in.ProductId)
	}
	if !g.IsEmpty(in.ItemId) {
		whereSet = fmt.Sprintf(" AND i.item_id IN  (%s)", gstr.JoinAny(in.ItemId, ","))
	}
	if !g.IsEmpty(in.CategoryId) {
		whereSet = fmt.Sprintf(" AND i.category_id IN  (%s)", gstr.JoinAny(in.CategoryId, ","))
	}
	if !g.IsEmpty(in.ProductName) {
		whereSet = fmt.Sprintf(" AND i.product_name LIKE  %%%s%%", in.ProductName)
	}
	if !g.IsEmpty(in.StoreType) {
		whereSet = fmt.Sprintf(" AND b.store_type = %d", in.StoreType)
	}
	if !g.IsEmpty(in.KindId) {
		whereSet = fmt.Sprintf(" AND b.kind_id = %d", in.KindId)
	}
	sql := fmt.Sprintf(` 
		SELECT
        FROM_UNIXTIME( b.create_time / 1000, '%%m-%%d' ) AS time,
        count(*) AS num
        FROM trade_order_item i
        LEFT JOIN
        trade_order_info b ON i.order_id = b.order_id
        WHERE 1 %s  
        GROUP BY time
        ORDER BY b.create_time`, whereSet)

	out, err = dao.DB().GetAll(ctx, sql)

	if err != nil {
		return out, err
	}

	return out, nil
}

// GetOrderItemNum 查询字段数据
func (dao *AnalyticsOrderDao) GetOrderItemNum(ctx context.Context, in *model.OrderItemNumTimelineInput) (out interface{}, err error) {
	whereSet := ""
	if !g.IsEmpty(in.Stime) && !g.IsEmpty(in.Etime) {
		whereSet = fmt.Sprintf(" AND b.create_time BETWEEN %d AND %d", in.Stime, in.Etime)
	}

	if !g.IsEmpty(in.StoreId) {
		whereSet = fmt.Sprintf(" AND i.store_id =  %d", in.StoreId)
	}

	if !g.IsEmpty(in.ProductId) {
		whereSet = fmt.Sprintf(" AND i.product_id =  %d", in.ProductId)
	}
	if !g.IsEmpty(in.ItemId) {
		whereSet = fmt.Sprintf(" AND i.item_id IN  (%s)", gstr.JoinAny(in.ItemId, ","))
	}
	if !g.IsEmpty(in.CategoryId) {
		whereSet = fmt.Sprintf(" AND i.category_id IN  (%s)", gstr.JoinAny(in.CategoryId, ","))
	}
	if !g.IsEmpty(in.ProductName) {
		whereSet = fmt.Sprintf(" AND i.product_name LIKE  %%%s%%", in.ProductName)
	}
	if !g.IsEmpty(in.StoreType) {
		whereSet = fmt.Sprintf(" AND b.store_type = %d", in.StoreType)
	}
	if !g.IsEmpty(in.KindId) {
		whereSet = fmt.Sprintf(" AND b.kind_id = %d", in.KindId)
	}
	sql := fmt.Sprintf(` 
        SELECT
        count(*) AS num
        FROM trade_order_item i
        LEFT JOIN
        trade_order_info b ON i.order_id = b.order_id
        WHERE 1 %s  `, whereSet)

	one, err := dao.DB().GetOne(ctx, sql)

	if err != nil {
		return out, err
	}

	if one["num"] != nil {
		out = one["num"]
	} else {
		out = 0
	}

	return
}

// ListOrderItemNum 查询字段数据
func (dao *AnalyticsOrderDao) ListOrderItemNum(ctx context.Context, in *model.OrderItemNumTimelineInput) (out gdb.Result, err error) {
	whereSet := ""
	if !g.IsEmpty(in.Stime) && !g.IsEmpty(in.Etime) {
		whereSet = fmt.Sprintf(" AND b.create_time BETWEEN %d AND %d", in.Stime, in.Etime)
	}

	if !g.IsEmpty(in.StoreId) {
		whereSet = fmt.Sprintf(" AND i.store_id =  %d", in.StoreId)
	}

	if !g.IsEmpty(in.ProductId) {
		whereSet = fmt.Sprintf(" AND i.product_id =  %d", in.ProductId)
	}
	if !g.IsEmpty(in.ItemId) {
		whereSet = fmt.Sprintf(" AND i.item_id IN  (%s)", gstr.JoinAny(in.ItemId, ","))
	}
	if !g.IsEmpty(in.CategoryId) {
		whereSet = fmt.Sprintf(" AND i.category_id IN  (%s)", gstr.JoinAny(in.CategoryId, ","))
	}
	if !g.IsEmpty(in.ProductName) {
		whereSet = fmt.Sprintf(" AND i.product_name LIKE  %%%s%%", in.ProductName)
	}
	if !g.IsEmpty(in.StoreType) {
		whereSet = fmt.Sprintf(" AND b.store_type = %d", in.StoreType)
	}
	if !g.IsEmpty(in.KindId) {
		whereSet = fmt.Sprintf(" AND b.kind_id = %d", in.KindId)
	}
	sql := fmt.Sprintf(` 
		SELECT
        i.product_id,
        i.item_id,
        i.order_item_image,
        i.product_name,
        i.item_name,
        sum(i.order_item_quantity) AS num,
        sum(i.order_item_amount) AS order_item_amount_sum
        FROM trade_order_item i
        LEFT JOIN
        trade_order_info b ON i.order_id = b.order_id     
        WHERE 1 %s 
        GROUP BY i.item_id
        ORDER BY num DESC
        LIMIT 0, 100 `, whereSet)

	out, err = dao.DB().GetAll(ctx, sql)

	if err != nil {
		return out, err
	}

	return out, nil
}
