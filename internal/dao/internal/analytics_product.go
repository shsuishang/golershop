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

// AnalyticsProductDao is the data access object for table trade_order_info.
type AnalyticsProductDao struct {
	table string // table is the underlying table name of the DAO.
	group string // group is the database configuration group name of current DAO
}

// NewAnalyticsProductDao creates and returns a new DAO object for table data access.
func NewAnalyticsProductDao() *AnalyticsProductDao {
	return &AnalyticsProductDao{
		group: "pt",
		table: "pt_product_index",
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AnalyticsProductDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Group returns the configuration group name of database of current dao.
func (dao *AnalyticsProductDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AnalyticsProductDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// GetProductTimeLine 查询字段数据
func (dao *AnalyticsProductDao) GetProductTimeLine(ctx context.Context, stime, etime int64) (out gdb.Result, err error) {
	whereSet := ""
	if !g.IsEmpty(stime) && !g.IsEmpty(etime) {
		whereSet = fmt.Sprintf(" AND product_add_time BETWEEN %d AND %d", stime, etime)
	}

	sql := fmt.Sprintf(`         SELECT count(*) num,
               from_unixtime(round(product_add_time / 1000), '%%m-%%d') AS time
        FROM pt_product_index
		WHERE 1 %s
        GROUP BY time
        ORDER BY product_add_time`, whereSet)

	out, err = dao.DB().GetAll(ctx, sql)

	if err != nil {
		return out, err
	}

	return out, nil
}

// GetProductNum 查询字段数据
func (dao *AnalyticsProductDao) GetProductNum(ctx context.Context, input *model.AnalyticsProductInput) (out interface{}, err error) {
	whereSet := ""
	if !g.IsEmpty(input.Stime) && !g.IsEmpty(input.Etime) {
		whereSet = fmt.Sprintf(" AND product_add_time BETWEEN %d AND %d", input.Stime, input.Etime)
	}

	if !g.IsEmpty(input.ProductStateId) {
		whereSet = fmt.Sprintf(" AND product_state_id = %d", input.ProductStateId)
	}

	if !g.IsEmpty(input.CategoryId) {
		whereSet = fmt.Sprintf(" AND category_id = %d", input.CategoryId)
	}

	sql := fmt.Sprintf(` 
        SELECT count(*) num
        FROM pt_product_index
		WHERE 1 %s`, whereSet)

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
