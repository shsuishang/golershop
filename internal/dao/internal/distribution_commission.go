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
	"database/sql"
	"math"

	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// DistributionCommissionDao is the data access object for table pay_distribution_commission.
type DistributionCommissionDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns DistributionCommissionColumns // columns contains all the column names of Table for convenient usage.
}

// DistributionCommissionColumns defines and stores column names for table pay_distribution_commission.
type DistributionCommissionColumns struct {
	UserId                             string // 店铺编号
	CommissionAmount                   string // 佣金总额:历史总额度
	CommissionDirectsellerAmount0      string // 销售员佣金
	CommissionDirectsellerAmount1      string // 二级销售员
	CommissionDirectsellerAmount2      string // 三级销售员
	CommissionBuyAmount0               string // 推广消费佣金
	CommissionBuyAmount1               string // 消费佣金
	CommissionBuyAmount2               string // 消费佣金
	CommissionClickAmount0             string // 本店流量佣金
	CommissionClickAmount1             string // 一级流量佣金
	CommissionClickAmount2             string // 二级流量佣金
	CommissionRegAmount0               string // 本店注册佣金
	CommissionRegAmount1               string // 一级注册佣金
	CommissionRegAmount2               string // 二级注册佣金
	CommissionSettled                  string // 已经结算佣金
	CommissionDirectsellerSettled      string // 销售员已经结算
	CommissionBuySettled               string // 推广员已经结算
	CommissionBuyDa                    string // 区代理收益
	CommissionBuyCa                    string // 市代理收益
	CommissionDirectsellerDa           string // 区代理收益
	CommissionDirectsellerCa           string // 市代理收益
	CommissionBuyTrade0                string // 交易总额
	CommissionBuyTrade1                string // 交易总额
	CommissionBuyTrade2                string // 交易总额
	CommissionBuyDaTrade               string // 交易总额
	CommissionBuyCaTrade               string // 交易总额
	CommissionDirectsellerTrade0       string // 交易总额
	CommissionDirectsellerTrade1       string // 交易总额
	CommissionDirectsellerTrade2       string // 交易总额
	CommissionDirectsellerDaTrade      string // 交易总额
	CommissionDirectsellerCaTrade      string // 交易总额
	CommissionPartnerBuyTrade          string // 合伙人交易总额
	CommissionPartnerDirectsellerTrade string // 合伙人交易总额
	CommissionPartnerDepositTrade      string // 合伙人充值总额
	CommissionDistributorAmount        string // 分销商收益
	CommissionSalespersonAmount        string // 销售员收益
	CommissionRefundAmount             string // 退款总佣金
	Version                            string // 版本
	PrimaryKey                         string // 主键
}

// distributionCommissionColumns holds the columns for table pay_distribution_commission.
var distributionCommissionColumns = DistributionCommissionColumns{
	UserId:                             "user_id",
	CommissionAmount:                   "commission_amount",
	CommissionDirectsellerAmount0:      "commission_directseller_amount_0",
	CommissionDirectsellerAmount1:      "commission_directseller_amount_1",
	CommissionDirectsellerAmount2:      "commission_directseller_amount_2",
	CommissionBuyAmount0:               "commission_buy_amount_0",
	CommissionBuyAmount1:               "commission_buy_amount_1",
	CommissionBuyAmount2:               "commission_buy_amount_2",
	CommissionClickAmount0:             "commission_click_amount_0",
	CommissionClickAmount1:             "commission_click_amount_1",
	CommissionClickAmount2:             "commission_click_amount_2",
	CommissionRegAmount0:               "commission_reg_amount_0",
	CommissionRegAmount1:               "commission_reg_amount_1",
	CommissionRegAmount2:               "commission_reg_amount_2",
	CommissionSettled:                  "commission_settled",
	CommissionDirectsellerSettled:      "commission_directseller_settled",
	CommissionBuySettled:               "commission_buy_settled",
	CommissionBuyDa:                    "commission_buy_da",
	CommissionBuyCa:                    "commission_buy_ca",
	CommissionDirectsellerDa:           "commission_directseller_da",
	CommissionDirectsellerCa:           "commission_directseller_ca",
	CommissionBuyTrade0:                "commission_buy_trade_0",
	CommissionBuyTrade1:                "commission_buy_trade_1",
	CommissionBuyTrade2:                "commission_buy_trade_2",
	CommissionBuyDaTrade:               "commission_buy_da_trade",
	CommissionBuyCaTrade:               "commission_buy_ca_trade",
	CommissionDirectsellerTrade0:       "commission_directseller_trade_0",
	CommissionDirectsellerTrade1:       "commission_directseller_trade_1",
	CommissionDirectsellerTrade2:       "commission_directseller_trade_2",
	CommissionDirectsellerDaTrade:      "commission_directseller_da_trade",
	CommissionDirectsellerCaTrade:      "commission_directseller_ca_trade",
	CommissionPartnerBuyTrade:          "commission_partner_buy_trade",
	CommissionPartnerDirectsellerTrade: "commission_partner_directseller_trade",
	CommissionPartnerDepositTrade:      "commission_partner_deposit_trade",
	CommissionDistributorAmount:        "commission_distributor_amount",
	CommissionSalespersonAmount:        "commission_salesperson_amount",
	CommissionRefundAmount:             "commission_refund_amount",
	Version:                            "version",
	PrimaryKey:                         "user_id",
}

// NewDistributionCommissionDao creates and returns a new DAO object for table data access.
func NewDistributionCommissionDao() *DistributionCommissionDao {
	return &DistributionCommissionDao{
		group:   "pay",
		table:   "pay_distribution_commission",
		columns: distributionCommissionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DistributionCommissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DistributionCommissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DistributionCommissionDao) Columns() DistributionCommissionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DistributionCommissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DistributionCommissionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DistributionCommissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *DistributionCommissionDao) Get(ctx context.Context, id any) (one *entity.DistributionCommission, err error) {
	var entitys []*entity.DistributionCommission
	entitys, err = dao.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(entitys) > 0 {
		one = entitys[0]
	}

	return one, err
}

// Gets 读取多条记录
func (dao *DistributionCommissionDao) Gets(ctx context.Context, id any) (entitys []*entity.DistributionCommission, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *DistributionCommissionDao) Find(ctx context.Context, in *do.DistributionCommissionListInput) (out []*entity.DistributionCommission, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	// 排序
	query = ml.BuildOrder(query, in.Sidx, in.Sort)
	if len(in.Order) > 0 {
		for _, it := range in.Order {
			query = ml.BuildOrder(query, it.Sidx, it.Sort)
		}
	}

	// 对象转换
	if err := query.Scan(&out); err != nil {
		return out, err
	}

	return out, nil
}

// FindOne 查询一条数据
func (dao *DistributionCommissionDao) FindOne(ctx context.Context, in *do.DistributionCommissionListInput) (one *entity.DistributionCommission, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.DistributionCommission
	entitys, err = dao.Find(ctx, in)

	if err != nil {
		return nil, err
	}

	if len(entitys) > 0 {
		one = entitys[0]
	}

	return one, err
}

// Find 查询字段数据
func (dao *DistributionCommissionDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.DistributionCommissionListInput) (out gdb.Result, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Fields(fieldNamesOrMapStruct).Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	// 排序
	query = ml.BuildOrder(query, in.Sidx, in.Sort)
	if len(in.Order) > 0 {
		for _, it := range in.Order {
			query = ml.BuildOrder(query, it.Sidx, it.Sort)
		}
	}

	out, err = query.All()

	if err != nil {
		return out, err
	}

	return out, nil
}

// FindKey 查询主键数据
func (dao *DistributionCommissionDao) FindKey(ctx context.Context, in *do.DistributionCommissionListInput) (out []interface{}, err error) {
	idRes, err := dao.FindFields(ctx, dao.Columns().PrimaryKey, in)

	if err != nil {
		return nil, err
	}

	for _, record := range idRes {
		if !record[dao.Columns().PrimaryKey].IsEmpty() {
			out = append(out, record[dao.Columns().PrimaryKey])
		}
	}

	return out, err
}

// List 分页读取
func (dao *DistributionCommissionDao) List(ctx context.Context, in *do.DistributionCommissionListInput) (out *do.DistributionCommissionListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.DistributionCommissionListOutput{}
	out.Page = in.Page
	out.Size = in.Size

	// 查询记录总数
	count, err1 := query.Count()
	if err1 != nil {
		return nil, err1
	}

	out.Records = count
	out.Total = int(math.Ceil(float64(count) / float64(out.Size)))

	// 排序
	query = ml.BuildOrder(query, in.Sidx, in.Sort)
	if len(in.Order) > 0 {
		for _, it := range in.Order {
			query = ml.BuildOrder(query, it.Sidx, it.Sort)
		}
	}

	// 分页
	query = query.Page(in.Page, in.Size)

	// 对象转换
	if err := query.Scan(&out.Items); err != nil {
		return out, err
	}

	return out, nil
}

// Add 新增
func (dao *DistributionCommissionDao) Add(ctx context.Context, in *do.DistributionCommission) (lastInsertId int64, err error) {
	data := do.DistributionCommission{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *DistributionCommissionDao) Edit(ctx context.Context, id any, in *do.DistributionCommission) (int64, error) {
	data := do.DistributionCommission{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *DistributionCommissionDao) EditWhere(ctx context.Context, where *do.DistributionCommissionListInput, in *do.DistributionCommission) (int64, error) {
	data := do.DistributionCommission{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *DistributionCommissionDao) Save(ctx context.Context, in *do.DistributionCommission) (affected int64, err error) {
	data := do.DistributionCommission{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	res, err := dao.Ctx(ctx).Data(data).OmitNil().Save()

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Saves 批量保存
func (dao *DistributionCommissionDao) Saves(ctx context.Context, in []*do.DistributionCommission) (affected int64, err error) {
	data := []do.DistributionCommission{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	res, err := dao.Ctx(ctx).Data(data).OmitNil().Save()

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Increment 增加
func (dao *DistributionCommissionDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *DistributionCommissionDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *DistributionCommissionDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *DistributionCommissionDao) RemoveWhere(ctx context.Context, where *do.DistributionCommissionListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *DistributionCommissionDao) Count(ctx context.Context, in *do.DistributionCommissionListInput) (count int, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	//记录数
	count, err = query.Count()

	if err != nil {
		return 0, err
	}

	return count, nil
}
