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

// ConsumeRecordDao is the data access object for table pay_consume_record.
type ConsumeRecordDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns ConsumeRecordColumns // columns contains all the column names of Table for convenient usage.
}

// ConsumeRecordColumns defines and stores column names for table pay_consume_record.
type ConsumeRecordColumns struct {
	ConsumeRecordId                 string // 支付流水号
	OrderId                         string // 商户订单编号
	UserId                          string // 所属用编号
	UserNickname                    string // 昵称
	CurrencyId                      string // 货币编号
	CurrencySymbolLeft              string // 左符号
	RecordTotal                     string // 金额
	RecordMoney                     string // 金额:record_total-佣金
	RecordCommissionFee             string // 佣金:平台佣金针对销售收款
	RecordDistributionCommissionFee string // 分销佣金:针对销售收款
	RecordDate                      string // 年-月-日
	RecordYear                      string // 年
	RecordMonth                     string // 月
	RecordDay                       string // 日
	RecordTitle                     string // 标题
	RecordDesc                      string // 描述
	RecordTime                      string // 支付时间
	TradeTypeId                     string // 交易类型(ENUM):1201-购物; 1202-转账; 1203-充值; 1204-提现; 1205-销售; 1206-佣金; 1207-退货付款;1208-退货收款;1209-转账收款
	PaymentTypeId                   string // 支付方式(ENUM):1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付;
	PaymentChannelId                string // 支付渠道
	StoreId                         string // 所属店铺
	ChainId                         string // 所属门店
	PaymentMetId                    string // 消费类型(ENUM):1-余额支付; 2-充值卡支付; 3-积分支付; 4-信用支付; 5-红包支付
	RecordEnable                    string // 状态(BOOL):1-已收款;0-作废
	SubsiteId                       string // 所属分站:0-总站
	Version                         string // 版本
	PrimaryKey                      string // 主键
}

// consumeRecordColumns holds the columns for table pay_consume_record.
var consumeRecordColumns = ConsumeRecordColumns{
	ConsumeRecordId:                 "consume_record_id",
	OrderId:                         "order_id",
	UserId:                          "user_id",
	UserNickname:                    "user_nickname",
	CurrencyId:                      "currency_id",
	CurrencySymbolLeft:              "currency_symbol_left",
	RecordTotal:                     "record_total",
	RecordMoney:                     "record_money",
	RecordCommissionFee:             "record_commission_fee",
	RecordDistributionCommissionFee: "record_distribution_commission_fee",
	RecordDate:                      "record_date",
	RecordYear:                      "record_year",
	RecordMonth:                     "record_month",
	RecordDay:                       "record_day",
	RecordTitle:                     "record_title",
	RecordDesc:                      "record_desc",
	RecordTime:                      "record_time",
	TradeTypeId:                     "trade_type_id",
	PaymentTypeId:                   "payment_type_id",
	PaymentChannelId:                "payment_channel_id",
	StoreId:                         "store_id",
	ChainId:                         "chain_id",
	PaymentMetId:                    "payment_met_id",
	RecordEnable:                    "record_enable",
	SubsiteId:                       "subsite_id",
	Version:                         "version",
	PrimaryKey:                      "consume_record_id",
}

// NewConsumeRecordDao creates and returns a new DAO object for table data access.
func NewConsumeRecordDao() *ConsumeRecordDao {
	return &ConsumeRecordDao{
		group:   "pay",
		table:   "pay_consume_record",
		columns: consumeRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ConsumeRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ConsumeRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ConsumeRecordDao) Columns() ConsumeRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ConsumeRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ConsumeRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ConsumeRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *ConsumeRecordDao) Get(ctx context.Context, id any) (one *entity.ConsumeRecord, err error) {
	var entitys []*entity.ConsumeRecord
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
func (dao *ConsumeRecordDao) Gets(ctx context.Context, id any) (entitys []*entity.ConsumeRecord, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *ConsumeRecordDao) Find(ctx context.Context, in *do.ConsumeRecordListInput) (out []*entity.ConsumeRecord, err error) {
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
func (dao *ConsumeRecordDao) FindOne(ctx context.Context, in *do.ConsumeRecordListInput) (one *entity.ConsumeRecord, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.ConsumeRecord
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
func (dao *ConsumeRecordDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.ConsumeRecordListInput) (out gdb.Result, err error) {
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
func (dao *ConsumeRecordDao) FindKey(ctx context.Context, in *do.ConsumeRecordListInput) (out []interface{}, err error) {
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
func (dao *ConsumeRecordDao) List(ctx context.Context, in *do.ConsumeRecordListInput) (out *do.ConsumeRecordListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.ConsumeRecordListOutput{}
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
func (dao *ConsumeRecordDao) Add(ctx context.Context, in *do.ConsumeRecord) (lastInsertId int64, err error) {
	data := do.ConsumeRecord{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *ConsumeRecordDao) Edit(ctx context.Context, id any, in *do.ConsumeRecord) (int64, error) {
	data := do.ConsumeRecord{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *ConsumeRecordDao) EditWhere(ctx context.Context, where *do.ConsumeRecordListInput, in *do.ConsumeRecord) (int64, error) {
	data := do.ConsumeRecord{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *ConsumeRecordDao) Save(ctx context.Context, in *do.ConsumeRecord) (affected int64, err error) {
	data := do.ConsumeRecord{}
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
func (dao *ConsumeRecordDao) Saves(ctx context.Context, in []*do.ConsumeRecord) (affected int64, err error) {
	data := []do.ConsumeRecord{}
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
func (dao *ConsumeRecordDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *ConsumeRecordDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *ConsumeRecordDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *ConsumeRecordDao) RemoveWhere(ctx context.Context, where *do.ConsumeRecordListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *ConsumeRecordDao) Count(ctx context.Context, in *do.ConsumeRecordListInput) (count int, err error) {
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
