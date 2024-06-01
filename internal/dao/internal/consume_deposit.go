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

// ConsumeDepositDao is the data access object for table pay_consume_deposit.
type ConsumeDepositDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns ConsumeDepositColumns // columns contains all the column names of Table for convenient usage.
}

// ConsumeDepositColumns defines and stores column names for table pay_consume_deposit.
type ConsumeDepositColumns struct {
	DepositId               string // 支付流水号
	DepositNo               string // 商城支付编号
	DepositTradeNo          string // 交易号:支付宝etc
	OrderId                 string // 商户网站唯一订单号(DOT):合并支付则为多个订单号, 没有创建联合支付交易号
	PaymentChannelId        string // 支付渠道
	DepositSubject          string // 商品名称
	DepositPaymentType      string // 支付方式(ENUM):1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付;
	DepositTradeStatus      string // 交易状态
	DepositSellerId         string // 卖家户号:支付宝etc
	DepositSellerEmail      string // 卖家支付账号
	DepositBuyerId          string // 买家支付用户号
	DepositBuyerEmail       string // 买家支付账号
	CurrencyId              string // 货币编号
	CurrencySymbolLeft      string // 左符号
	DepositTotalFee         string // 交易金额
	DepositQuantity         string // 购买数量
	DepositPrice            string // 商品单价
	DepositBody             string // 商品描述
	DepositIsTotalFeeAdjust string // 是否调整总价
	DepositUseCoupon        string // 是否使用红包买家
	DepositDiscount         string // 折扣
	DepositNotifyTime       string // 通知时间
	DepositNotifyType       string // 通知类型
	DepositNotifyId         string // 通知校验编号
	DepositSignType         string // 签名方式
	DepositSign             string // 签名
	DepositExtraParam       string // 额外参数
	DepositService          string // 支付
	DepositState            string // 支付状态:0-默认; 1-接收正确数据处理完逻辑; 9-异常订单
	DepositAsync            string // 是否同步(BOOL):0-同步; 1-异步回调使用
	DepositReview           string // 收款确认(BOOL):0-未确认;1-已确认
	DepositEnable           string // 是否作废(BOOL):1-正常; 2-作废
	StoreId                 string // 所属店铺:直接交易起作用
	UserId                  string // 所属用户
	ChainId                 string // 所属门店:直接交易起作用
	SubsiteId               string // 所属分站:直接交易起作用
	DepositTime             string // 支付时间
	PrimaryKey              string // 主键
}

// consumeDepositColumns holds the columns for table pay_consume_deposit.
var consumeDepositColumns = ConsumeDepositColumns{
	DepositId:               "deposit_id",
	DepositNo:               "deposit_no",
	DepositTradeNo:          "deposit_trade_no",
	OrderId:                 "order_id",
	PaymentChannelId:        "payment_channel_id",
	DepositSubject:          "deposit_subject",
	DepositPaymentType:      "deposit_payment_type",
	DepositTradeStatus:      "deposit_trade_status",
	DepositSellerId:         "deposit_seller_id",
	DepositSellerEmail:      "deposit_seller_email",
	DepositBuyerId:          "deposit_buyer_id",
	DepositBuyerEmail:       "deposit_buyer_email",
	CurrencyId:              "currency_id",
	CurrencySymbolLeft:      "currency_symbol_left",
	DepositTotalFee:         "deposit_total_fee",
	DepositQuantity:         "deposit_quantity",
	DepositPrice:            "deposit_price",
	DepositBody:             "deposit_body",
	DepositIsTotalFeeAdjust: "deposit_is_total_fee_adjust",
	DepositUseCoupon:        "deposit_use_coupon",
	DepositDiscount:         "deposit_discount",
	DepositNotifyTime:       "deposit_notify_time",
	DepositNotifyType:       "deposit_notify_type",
	DepositNotifyId:         "deposit_notify_id",
	DepositSignType:         "deposit_sign_type",
	DepositSign:             "deposit_sign",
	DepositExtraParam:       "deposit_extra_param",
	DepositService:          "deposit_service",
	DepositState:            "deposit_state",
	DepositAsync:            "deposit_async",
	DepositReview:           "deposit_review",
	DepositEnable:           "deposit_enable",
	StoreId:                 "store_id",
	UserId:                  "user_id",
	ChainId:                 "chain_id",
	SubsiteId:               "subsite_id",
	DepositTime:             "deposit_time",
	PrimaryKey:              "deposit_id",
}

// NewConsumeDepositDao creates and returns a new DAO object for table data access.
func NewConsumeDepositDao() *ConsumeDepositDao {
	return &ConsumeDepositDao{
		group:   "pay",
		table:   "pay_consume_deposit",
		columns: consumeDepositColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ConsumeDepositDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ConsumeDepositDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ConsumeDepositDao) Columns() ConsumeDepositColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ConsumeDepositDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ConsumeDepositDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ConsumeDepositDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *ConsumeDepositDao) Get(ctx context.Context, id any) (one *entity.ConsumeDeposit, err error) {
	var entitys []*entity.ConsumeDeposit
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
func (dao *ConsumeDepositDao) Gets(ctx context.Context, id any) (entitys []*entity.ConsumeDeposit, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *ConsumeDepositDao) Find(ctx context.Context, in *do.ConsumeDepositListInput) (out []*entity.ConsumeDeposit, err error) {
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
func (dao *ConsumeDepositDao) FindOne(ctx context.Context, in *do.ConsumeDepositListInput) (one *entity.ConsumeDeposit, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.ConsumeDeposit
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
func (dao *ConsumeDepositDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.ConsumeDepositListInput) (out gdb.Result, err error) {
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
func (dao *ConsumeDepositDao) FindKey(ctx context.Context, in *do.ConsumeDepositListInput) (out []interface{}, err error) {
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
func (dao *ConsumeDepositDao) List(ctx context.Context, in *do.ConsumeDepositListInput) (out *do.ConsumeDepositListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.ConsumeDepositListOutput{}
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
func (dao *ConsumeDepositDao) Add(ctx context.Context, in *do.ConsumeDeposit) (lastInsertId int64, err error) {
	data := do.ConsumeDeposit{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *ConsumeDepositDao) Edit(ctx context.Context, id any, in *do.ConsumeDeposit) (int64, error) {
	data := do.ConsumeDeposit{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *ConsumeDepositDao) EditWhere(ctx context.Context, where *do.ConsumeDepositListInput, in *do.ConsumeDeposit) (int64, error) {
	data := do.ConsumeDeposit{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *ConsumeDepositDao) Save(ctx context.Context, in *do.ConsumeDeposit) (affected int64, err error) {
	data := do.ConsumeDeposit{}
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
func (dao *ConsumeDepositDao) Saves(ctx context.Context, in []*do.ConsumeDeposit) (affected int64, err error) {
	data := []do.ConsumeDeposit{}
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
func (dao *ConsumeDepositDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *ConsumeDepositDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *ConsumeDepositDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *ConsumeDepositDao) RemoveWhere(ctx context.Context, where *do.ConsumeDepositListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *ConsumeDepositDao) Count(ctx context.Context, in *do.ConsumeDepositListInput) (count int, err error) {
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
