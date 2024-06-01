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

// OrderReturnDao is the data access object for table trade_order_return.
type OrderReturnDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns OrderReturnColumns // columns contains all the column names of Table for convenient usage.
}

// OrderReturnColumns defines and stores column names for table trade_order_return.
type OrderReturnColumns struct {
	ReturnId                   string // 退单号
	ServiceTypeId              string // 服务类型(ENUM):1-退款;2-退货;3-换货;4-维修
	OrderId                    string // 订单编号
	ReturnRefundAmount         string // 退款金额
	ReturnRefundPoint          string // 积分部分
	StoreId                    string // 店铺编号
	BuyerUserId                string // 买家编号
	BuyerStoreId               string // 买家是否有店铺
	ReturnAddTime              string // 添加时间
	ReturnReasonId             string // 退款理由编号
	ReturnBuyerMessage         string // 买家退货备注
	ReturnAddrContacter        string // 收货人
	ReturnTel                  string // 联系电话
	ReturnAddr                 string // 收货地址详情
	ReturnPostCode             string // 邮编
	ExpressId                  string // 物流公司编号
	ReturnTrackingName         string // 物流名称
	ReturnTrackingNumber       string // 物流单号
	PlantformReturnStateId     string // 申请状态平台(ENUM):3180-未申请;3181-待处理;3182-为已完成
	ReturnStateId              string // 卖家处理状态(ENUM): 3100-【客户】提交退单;3105-退单审核;3110-收货确认;3115-退款确认;3120-客户】收款确认;3125-完成
	ReturnIsPaid               string // 退款完成
	ReturnIsShippingFee        string // 退货类型(BOOL): 0-退款单;1-退运费单
	ReturnShippingFee          string // 退运费额度
	ReturnFlag                 string // 退货类型(ENUM): 0-不用退货;1-需要退货
	ReturnType                 string // 申请类型(ENUM): 1-退款申请; 2-退货申请; 3-虚拟退款
	ReturnOrderLock            string // 订单锁定类型(BOOL):1-不用锁定;2-需要锁定
	ReturnItemStateId          string // 物流状态(LIST):2030-待发货;2040-已发货/待收货确认;2060-已完成/已签收;2070-已取消/已作废;
	ReturnStoreTime            string // 商家处理时间
	ReturnStoreMessage         string // 商家备注
	ReturnCommisionFee         string // 退还佣金
	ReturnFinishTime           string // 退款完成时间
	ReturnPlatformMessage      string // 平台留言
	ReturnIsSettlemented       string // 订单是否结算(BOOL): 0-未结算; 1-已结算
	ReturnSettlementTime       string // 订单结算时间
	ReturnChannelCode          string // 退款渠道(ENUM):money-余额;alipay-支付宝;wx_native-微信
	ReturnChannelFlag          string // 渠道是否退款(ENUM): 0-待退; 1-已退; 2-异常
	ReturnChannelTime          string // 渠道退款时间
	ReturnChannelTransId       string // 渠道退款单号
	DepositTradeNo             string // 交易号
	PaymentChannelId           string // 支付渠道
	TradePaymentAmount         string // 实付金额:在线支付金额
	ReturnContactName          string // 联系人
	ReturnStoreUserId          string // 审核人员id
	ReturnWithdrawConfirm      string // 提现审核(BOOL):0-未审核; 1-已审核
	ReturnFinancialConfirm     string // 退款财务确认(BOOL):0-未确认; 1-已确认
	ReturnFinancialConfirmTime string // 退款财务确认时间
	SubsiteId                  string // 所属分站:0-总站
	PrimaryKey                 string // 主键
}

// orderReturnColumns holds the columns for table trade_order_return.
var orderReturnColumns = OrderReturnColumns{
	ReturnId:                   "return_id",
	ServiceTypeId:              "service_type_id",
	OrderId:                    "order_id",
	ReturnRefundAmount:         "return_refund_amount",
	ReturnRefundPoint:          "return_refund_point",
	StoreId:                    "store_id",
	BuyerUserId:                "buyer_user_id",
	BuyerStoreId:               "buyer_store_id",
	ReturnAddTime:              "return_add_time",
	ReturnReasonId:             "return_reason_id",
	ReturnBuyerMessage:         "return_buyer_message",
	ReturnAddrContacter:        "return_addr_contacter",
	ReturnTel:                  "return_tel",
	ReturnAddr:                 "return_addr",
	ReturnPostCode:             "return_post_code",
	ExpressId:                  "express_id",
	ReturnTrackingName:         "return_tracking_name",
	ReturnTrackingNumber:       "return_tracking_number",
	PlantformReturnStateId:     "plantform_return_state_id",
	ReturnStateId:              "return_state_id",
	ReturnIsPaid:               "return_is_paid",
	ReturnIsShippingFee:        "return_is_shipping_fee",
	ReturnShippingFee:          "return_shipping_fee",
	ReturnFlag:                 "return_flag",
	ReturnType:                 "return_type",
	ReturnOrderLock:            "return_order_lock",
	ReturnItemStateId:          "return_item_state_id",
	ReturnStoreTime:            "return_store_time",
	ReturnStoreMessage:         "return_store_message",
	ReturnCommisionFee:         "return_commision_fee",
	ReturnFinishTime:           "return_finish_time",
	ReturnPlatformMessage:      "return_platform_message",
	ReturnIsSettlemented:       "return_is_settlemented",
	ReturnSettlementTime:       "return_settlement_time",
	ReturnChannelCode:          "return_channel_code",
	ReturnChannelFlag:          "return_channel_flag",
	ReturnChannelTime:          "return_channel_time",
	ReturnChannelTransId:       "return_channel_trans_id",
	DepositTradeNo:             "deposit_trade_no",
	PaymentChannelId:           "payment_channel_id",
	TradePaymentAmount:         "trade_payment_amount",
	ReturnContactName:          "return_contact_name",
	ReturnStoreUserId:          "return_store_user_id",
	ReturnWithdrawConfirm:      "return_withdraw_confirm",
	ReturnFinancialConfirm:     "return_financial_confirm",
	ReturnFinancialConfirmTime: "return_financial_confirm_time",
	SubsiteId:                  "subsite_id",
	PrimaryKey:                 "return_id",
}

// NewOrderReturnDao creates and returns a new DAO object for table data access.
func NewOrderReturnDao() *OrderReturnDao {
	return &OrderReturnDao{
		group:   "trade",
		table:   "trade_order_return",
		columns: orderReturnColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderReturnDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderReturnDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrderReturnDao) Columns() OrderReturnColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderReturnDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrderReturnDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrderReturnDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *OrderReturnDao) Get(ctx context.Context, id any) (one *entity.OrderReturn, err error) {
	var entitys []*entity.OrderReturn
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
func (dao *OrderReturnDao) Gets(ctx context.Context, id any) (entitys []*entity.OrderReturn, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *OrderReturnDao) Find(ctx context.Context, in *do.OrderReturnListInput) (out []*entity.OrderReturn, err error) {
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
func (dao *OrderReturnDao) FindOne(ctx context.Context, in *do.OrderReturnListInput) (one *entity.OrderReturn, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.OrderReturn
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
func (dao *OrderReturnDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.OrderReturnListInput) (out gdb.Result, err error) {
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
func (dao *OrderReturnDao) FindKey(ctx context.Context, in *do.OrderReturnListInput) (out []interface{}, err error) {
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
func (dao *OrderReturnDao) List(ctx context.Context, in *do.OrderReturnListInput) (out *do.OrderReturnListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.OrderReturnListOutput{}
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
func (dao *OrderReturnDao) Add(ctx context.Context, in *do.OrderReturn) (lastInsertId int64, err error) {
	data := do.OrderReturn{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *OrderReturnDao) Edit(ctx context.Context, id any, in *do.OrderReturn) (int64, error) {
	data := do.OrderReturn{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *OrderReturnDao) EditWhere(ctx context.Context, where *do.OrderReturnListInput, in *do.OrderReturn) (int64, error) {
	data := do.OrderReturn{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *OrderReturnDao) Save(ctx context.Context, in *do.OrderReturn) (affected int64, err error) {
	data := do.OrderReturn{}
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
func (dao *OrderReturnDao) Saves(ctx context.Context, in []*do.OrderReturn) (affected int64, err error) {
	data := []do.OrderReturn{}
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
func (dao *OrderReturnDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *OrderReturnDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *OrderReturnDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *OrderReturnDao) RemoveWhere(ctx context.Context, where *do.OrderReturnListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *OrderReturnDao) Count(ctx context.Context, in *do.OrderReturnListInput) (count int, err error) {
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
