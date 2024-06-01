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
	"golershop.cn/internal/model"
)

// AnalyticsSysDao is the data access object for table trade_order_info.
type AnalyticsSysDao struct {
	BaseRepository
	table string // table is the underlying table name of the DAO.
	group string // group is the database configuration group name of current DAO
}

// NewAnalyticsSysDao creates and returns a new DAO object for table data access.
func NewAnalyticsSysDao() *AnalyticsSysDao {
	return &AnalyticsSysDao{
		BaseRepository: BaseRepository{
			Table: "trade_order_info",
			Group: "trade",
		},
		group: "trade",
		table: "trade_order_info",
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AnalyticsSysDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Group returns the configuration group name of database of current dao.
func (dao *AnalyticsSysDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AnalyticsSysDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// GetVisitor 查询字段数据
func (dao *AnalyticsSysDao) GetVisitor(ctx context.Context, stime, etime int64) (out interface{}, err error) {
	whereSet := ""
	if !g.IsEmpty(stime) && !g.IsEmpty(etime) {
		whereSet = fmt.Sprintf(" AND access_time BETWEEN %d AND %d", stime, etime)
	}

	sql := fmt.Sprintf(`
		SELECT
            COUNT(*) AS num
        FROM
            sys_access_history
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

// GetAccessNum 查询字段数据
func (dao *AnalyticsSysDao) GetAccessNum(ctx context.Context, stime, etime int64) (out interface{}, err error) {
	whereSet := ""
	if !g.IsEmpty(stime) && !g.IsEmpty(etime) {
		whereSet = fmt.Sprintf(" AND access_time BETWEEN %d AND %d", stime, etime)
	}

	sql := fmt.Sprintf(`
        SELECT
            COUNT(*) AS num
        FROM
            sys_access_history
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

// GetVisitorNum 查询字段数据
func (dao *AnalyticsSysDao) GetVisitorNum(ctx context.Context, stime, etime int64) (out interface{}, err error) {
	whereSet := ""
	if !g.IsEmpty(stime) && !g.IsEmpty(etime) {
		whereSet = fmt.Sprintf(" AND access_time BETWEEN %d AND %d", stime, etime)
	}

	sql := fmt.Sprintf(`
        SELECT
        COUNT(DISTINCT ( access_client_id )) AS num
        FROM
        sys_access_history
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

// GetAccessItemTimeLine 查询字段数据
func (dao *AnalyticsSysDao) GetAccessItemTimeLine(ctx context.Context, stime, etime, itemId int64) (out gdb.Result, err error) {
	whereSet := ""
	if !g.IsEmpty(stime) && !g.IsEmpty(etime) {
		whereSet = fmt.Sprintf(" AND access_time BETWEEN %d AND %d", stime, etime)
	}

	if !g.IsEmpty(itemId) {
		whereSet = fmt.Sprintf(" AND item_id = %d", itemId)
	}

	sql := fmt.Sprintf(`
        SELECT count(*) num,
        from_unixtime(round(access_time / 1000), '%%m-%%d') AS time
        FROM sys_access_history
        WHERE 1 %s `, whereSet)

	out, err = dao.DB().GetAll(ctx, sql)

	if err != nil {
		return out, err
	}

	return out, nil
}

// GetAccessItemNum 查询字段数据
func (dao *AnalyticsSysDao) GetAccessItemNum(ctx context.Context, stime, etime, itemId int64) (out interface{}, err error) {
	whereSet := ""
	if !g.IsEmpty(stime) && !g.IsEmpty(etime) {
		whereSet = fmt.Sprintf(" AND access_time BETWEEN %d AND %d", stime, etime)
	}

	if !g.IsEmpty(itemId) {
		whereSet = fmt.Sprintf(" AND item_id = %d", itemId)
	}

	sql := fmt.Sprintf(`
        SELECT count(*) num
        FROM sys_access_history
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

// GetAccessItemUserTimeLine 查询字段数据
func (dao *AnalyticsSysDao) GetAccessItemUserTimeLine(ctx context.Context, stime, etime, itemId int64) (out gdb.Result, err error) {
	whereSet := ""
	if !g.IsEmpty(stime) && !g.IsEmpty(etime) {
		whereSet = fmt.Sprintf(" AND access_time BETWEEN %d AND %d", stime, etime)
	}

	if !g.IsEmpty(itemId) {
		whereSet = fmt.Sprintf(" AND item_id = %d", itemId)
	}

	sql := fmt.Sprintf(`
		SELECT
        count( DISTINCT ( access_client_id ) ) AS num,
        from_unixtime(round(access_time / 1000), '%%m-%%d') AS time
        FROM sys_access_history
        WHERE 1 %s
        GROUP BY time
        ORDER BY access_time `, whereSet)

	out, err = dao.DB().GetAll(ctx, sql)

	if err != nil {
		return out, err
	}

	return out, nil
}

// GetAccessItemUserNum 查询字段数据
func (dao *AnalyticsSysDao) GetAccessItemUserNum(ctx context.Context, stime, etime, itemId int64) (out interface{}, err error) {
	whereSet := ""
	if !g.IsEmpty(stime) && !g.IsEmpty(etime) {
		whereSet = fmt.Sprintf(" AND access_time BETWEEN %d AND %d", stime, etime)
	}

	if !g.IsEmpty(itemId) {
		whereSet = fmt.Sprintf(" AND item_id = %d", itemId)
	}

	sql := fmt.Sprintf(`
        SELECT count(DISTINCT ( access_client_id )) num
        FROM sys_access_history
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

// GetAccessVisitorTimeLine 查询字段数据
func (dao *AnalyticsSysDao) GetAccessVisitorTimeLine(ctx context.Context, stime, etime int64) (out gdb.Result, err error) {
	whereSet := ""
	if !g.IsEmpty(stime) && !g.IsEmpty(etime) {
		whereSet = fmt.Sprintf(" AND access_time BETWEEN %d AND %d", stime, etime)
	}

	sql := fmt.Sprintf(`
        SELECT count(*) num,
        from_unixtime(round(access_time / 1000), '%%m-%%d') AS time
        FROM sys_access_history
		WHERE 1 %s 
        GROUP BY time
        ORDER BY access_time `, whereSet)

	out, err = dao.DB().GetAll(ctx, sql)

	if err != nil {
		return out, err
	}

	return out, nil
}

// ListAccessItem 查询字段数据
func (dao *AnalyticsSysDao) ListAccessItem(ctx context.Context, in *model.AccessItemTimelineInput) (out gdb.Result, err error) {
	whereSet := ""
	if !g.IsEmpty(in.Stime) && !g.IsEmpty(in.Etime) {
		whereSet = fmt.Sprintf(" AND ash.access_time BETWEEN %d AND %d", in.Stime, in.Etime)
	}

	if !g.IsEmpty(in.ItemId) {
		whereSet = fmt.Sprintf(" AND item_id = %d", in.ItemId)
	}

	sql := fmt.Sprintf(`
        SELECT ash.item_id,ppi.item_name,ppi.product_id,ppi.item_unit_price,ppb.product_name,
               COUNT(*) AS num
        FROM sys_access_history ash
                 INNER JOIN pt_product_item ppi on ash.item_id = ppi.item_id
                 INNER JOIN pt_product_base ppb on ppi.product_id = ppb.product_id
		WHERE 1 %s 
        GROUP BY ash.item_id
        ORDER BY num DESC
            LIMIT 0, 100 `, whereSet)

	out, err = dao.DB().GetAll(ctx, sql)

	if err != nil {
		return out, err
	}

	return out, nil
}
