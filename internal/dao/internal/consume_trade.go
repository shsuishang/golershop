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

// ConsumeTradeDao is the data access object for table pay_consume_trade.
type ConsumeTradeDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ConsumeTradeColumns // columns contains all the column names of Table for convenient usage.
}

// ConsumeTradeColumns defines and stores column names for table pay_consume_trade.
type ConsumeTradeColumns struct {
	ConsumeTradeId           string // 交易订单编号
	TradeTitle               string // 标题
	OrderId                  string // 商户订单编号
	BuyerId                  string // 买家编号
	BuyerStoreId             string // 买家是否有店铺
	StoreId                  string // 店铺编号
	SubsiteId                string // 所属分站:0-总站
	SellerId                 string // 卖家编号
	ChainId                  string // 门店编号
	TradeIsPaid              string // 支付状态
	TradeTypeId              string // 交易类型(ENUM):1201-购物; 1202-转账; 1203-充值; 1204-提现; 1205-销售; 1206-佣金;
	PaymentChannelId         string // 支付渠道
	TradeModeId              string // 交易模式(ENUM):1-担保交易;  2-直接交易
	RechargeLevelId          string // 充值编号
	CurrencyId               string // 货币编号
	CurrencySymbolLeft       string // 左符号
	OrderPaymentAmount       string // 总付款额度: trade_payment_amount + trade_payment_money + trade_payment_recharge_card + trade_payment_points
	OrderCommissionFee       string // 平台交易佣金
	TradePaymentAmount       string // 实付金额:在线支付金额,此为订单默认需要支付额度。
	TradePaymentMoney        string // 余额支付
	TradePaymentRechargeCard string // 充值卡余额支付
	TradePaymentPoints       string // 积分支付
	TradePaymentSp           string // 众宝支付
	TradePaymentCredit       string // 信用支付
	TradePaymentRedpack      string // 红包支付
	TradeDiscount            string // 折扣优惠
	TradeAmount              string // 总额虚拟:trade_order_amount + trade_discount
	TradeDesc                string // 描述
	TradeRemark              string // 备注
	TradeCreateTime          string // 创建时间
	TradePaidTime            string // 付款时间
	TradeDelete              string // 是否删除
	Version                  string // 版本
	PrimaryKey               string // 主键
}

// consumeTradeColumns holds the columns for table pay_consume_trade.
var consumeTradeColumns = ConsumeTradeColumns{
	ConsumeTradeId:           "consume_trade_id",
	TradeTitle:               "trade_title",
	OrderId:                  "order_id",
	BuyerId:                  "buyer_id",
	BuyerStoreId:             "buyer_store_id",
	StoreId:                  "store_id",
	SubsiteId:                "subsite_id",
	SellerId:                 "seller_id",
	ChainId:                  "chain_id",
	TradeIsPaid:              "trade_is_paid",
	TradeTypeId:              "trade_type_id",
	PaymentChannelId:         "payment_channel_id",
	TradeModeId:              "trade_mode_id",
	RechargeLevelId:          "recharge_level_id",
	CurrencyId:               "currency_id",
	CurrencySymbolLeft:       "currency_symbol_left",
	OrderPaymentAmount:       "order_payment_amount",
	OrderCommissionFee:       "order_commission_fee",
	TradePaymentAmount:       "trade_payment_amount",
	TradePaymentMoney:        "trade_payment_money",
	TradePaymentRechargeCard: "trade_payment_recharge_card",
	TradePaymentPoints:       "trade_payment_points",
	TradePaymentSp:           "trade_payment_sp",
	TradePaymentCredit:       "trade_payment_credit",
	TradePaymentRedpack:      "trade_payment_redpack",
	TradeDiscount:            "trade_discount",
	TradeAmount:              "trade_amount",
	TradeDesc:                "trade_desc",
	TradeRemark:              "trade_remark",
	TradeCreateTime:          "trade_create_time",
	TradePaidTime:            "trade_paid_time",
	TradeDelete:              "trade_delete",
	Version:                  "version",
	PrimaryKey:               "consume_trade_id",
}

// NewConsumeTradeDao creates and returns a new DAO object for table data access.
func NewConsumeTradeDao() *ConsumeTradeDao {
	return &ConsumeTradeDao{
		group:   "pay",
		table:   "pay_consume_trade",
		columns: consumeTradeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ConsumeTradeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ConsumeTradeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ConsumeTradeDao) Columns() ConsumeTradeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ConsumeTradeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ConsumeTradeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ConsumeTradeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *ConsumeTradeDao) Get(ctx context.Context, id any) (one *entity.ConsumeTrade, err error) {
	var entitys []*entity.ConsumeTrade
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
func (dao *ConsumeTradeDao) Gets(ctx context.Context, id any) (entitys []*entity.ConsumeTrade, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *ConsumeTradeDao) Find(ctx context.Context, in *do.ConsumeTradeListInput) (out []*entity.ConsumeTrade, err error) {
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
func (dao *ConsumeTradeDao) FindOne(ctx context.Context, in *do.ConsumeTradeListInput) (one *entity.ConsumeTrade, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.ConsumeTrade
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
func (dao *ConsumeTradeDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.ConsumeTradeListInput) (out gdb.Result, err error) {
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
func (dao *ConsumeTradeDao) FindKey(ctx context.Context, in *do.ConsumeTradeListInput) (out []interface{}, err error) {
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
func (dao *ConsumeTradeDao) List(ctx context.Context, in *do.ConsumeTradeListInput) (out *do.ConsumeTradeListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.ConsumeTradeListOutput{}
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
func (dao *ConsumeTradeDao) Add(ctx context.Context, in *do.ConsumeTrade) (lastInsertId int64, err error) {
	data := do.ConsumeTrade{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *ConsumeTradeDao) Edit(ctx context.Context, id any, in *do.ConsumeTrade) (int64, error) {
	data := do.ConsumeTrade{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *ConsumeTradeDao) EditWhere(ctx context.Context, where *do.ConsumeTradeListInput, in *do.ConsumeTrade) (int64, error) {
	data := do.ConsumeTrade{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *ConsumeTradeDao) Save(ctx context.Context, in *do.ConsumeTrade) (affected int64, err error) {
	data := do.ConsumeTrade{}
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
func (dao *ConsumeTradeDao) Saves(ctx context.Context, in []*do.ConsumeTrade) (affected int64, err error) {
	data := []do.ConsumeTrade{}
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
func (dao *ConsumeTradeDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *ConsumeTradeDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *ConsumeTradeDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *ConsumeTradeDao) RemoveWhere(ctx context.Context, where *do.ConsumeTradeListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *ConsumeTradeDao) Count(ctx context.Context, in *do.ConsumeTradeListInput) (count int, err error) {
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
