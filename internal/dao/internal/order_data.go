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

// OrderDataDao is the data access object for table trade_order_data.
type OrderDataDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns OrderDataColumns // columns contains all the column names of Table for convenient usage.
}

// OrderDataColumns defines and stores column names for table trade_order_data.
type OrderDataColumns struct {
	OrderId                  string // 订单编号
	OrderDesc                string // 订单描述
	OrderDelayTime           string // 延迟时间,默认为0 - 收货确认
	DeliveryTypeId           string // 配送方式
	DeliveryTimeId           string // 配送时间:要求，不限、周一~周五、周末等等
	DeliveryTime             string // 配送日期
	DeliveryIstimer          string // 是否定时配送(BOOL):0-不定时;1-定时
	OrderMessage             string // 买家订单留言
	OrderItemAmount          string // 商品总价格/商品金额, 不包含运费
	OrderDiscountAmount      string // 折扣价格/优惠总金额
	OrderAdjustFee           string // 手工调整费用店铺优惠
	OrderPointsFee           string // 积分抵扣费用
	OrderShippingFeeAmount   string // 运费价格/运费金额
	OrderShippingFee         string // 实际运费金额-卖家可修改
	VoucherId                string // 代金券id/优惠券/返现:发放选择使用
	VoucherNumber            string // 代金券编码
	VoucherPrice             string // 代金券面额
	RedpacketId              string // 红包id-平台代金券
	RedpacketNumber          string // 红包编码
	RedpacketPrice           string // 红包面额
	OrderRedpacketPrice      string // 红包抵扣订单金额
	OrderResourceExt1        string // 第二需要支付资源例如积分
	OrderResourceExt2        string // 众宝
	OrderResourceExt3        string // 金宝
	TradePaymentMoney        string // 余额支付
	TradePaymentRechargeCard string // 充值卡支付
	TradePaymentCredit       string // 信用支付
	OrderRefundStatus        string // 退款状态:0-是无退款;1-是部分退款;2-是全部退款
	OrderRefundAmount        string // 退款金额:申请额度
	OrderRefundAgreeAmount   string // 退款金额:同意额度
	OrderRefundAgreeCash     string // 已同意退的现金
	OrderRefundAgreePoints   string // 已退的积分额度
	OrderReturnStatus        string // 退货状态(ENUM):0-是无退货;1-是部分退货;2-是全部退货
	OrderReturnNum           string // 退货数量
	OrderReturnIds           string // 退货单编号s(DOT):冗余
	OrderCommissionFee       string // 平台交易佣金
	OrderCommissionFeeRefund string // 交易佣金-退款
	OrderPointsAdd           string // 订单赠送积分
	OrderActivityData        string // 促销信息
	OrderCancelIdentity      string // 订单取消者身份(ENUM):1-买家; 2-卖家; 3-系统
	OrderCancelReason        string // 订单取消原因
	OrderCancelTime          string // 订单取消时间
	OrderBpAdd               string // 赠送资源2
	OrderRebate              string // 订单返利
	BuyerMobile              string // 手机号码
	BuyerContacter           string // 联系人
	ActivityManhuiId         string // 满返优惠券活动id(DOT)
	ActivityDoublePointsId   string // 活动-多倍积分id
	OrderDoublePointsAdd     string // 活动-多倍积分
	ActivityVoucherId        string // 满返用户优惠券id(DOT)
	OrderActivityManhuiState string // 满返优惠券发放状态(ENUM):1000-无需发放;1001-待发放; 1002-已发放; 1003-发放异常
	Version                  string // 版本
	PrimaryKey               string // 主键
}

// orderDataColumns holds the columns for table trade_order_data.
var orderDataColumns = OrderDataColumns{
	OrderId:                  "order_id",
	OrderDesc:                "order_desc",
	OrderDelayTime:           "order_delay_time",
	DeliveryTypeId:           "delivery_type_id",
	DeliveryTimeId:           "delivery_time_id",
	DeliveryTime:             "delivery_time",
	DeliveryIstimer:          "delivery_istimer",
	OrderMessage:             "order_message",
	OrderItemAmount:          "order_item_amount",
	OrderDiscountAmount:      "order_discount_amount",
	OrderAdjustFee:           "order_adjust_fee",
	OrderPointsFee:           "order_points_fee",
	OrderShippingFeeAmount:   "order_shipping_fee_amount",
	OrderShippingFee:         "order_shipping_fee",
	VoucherId:                "voucher_id",
	VoucherNumber:            "voucher_number",
	VoucherPrice:             "voucher_price",
	RedpacketId:              "redpacket_id",
	RedpacketNumber:          "redpacket_number",
	RedpacketPrice:           "redpacket_price",
	OrderRedpacketPrice:      "order_redpacket_price",
	OrderResourceExt1:        "order_resource_ext1",
	OrderResourceExt2:        "order_resource_ext2",
	OrderResourceExt3:        "order_resource_ext3",
	TradePaymentMoney:        "trade_payment_money",
	TradePaymentRechargeCard: "trade_payment_recharge_card",
	TradePaymentCredit:       "trade_payment_credit",
	OrderRefundStatus:        "order_refund_status",
	OrderRefundAmount:        "order_refund_amount",
	OrderRefundAgreeAmount:   "order_refund_agree_amount",
	OrderRefundAgreeCash:     "order_refund_agree_cash",
	OrderRefundAgreePoints:   "order_refund_agree_points",
	OrderReturnStatus:        "order_return_status",
	OrderReturnNum:           "order_return_num",
	OrderReturnIds:           "order_return_ids",
	OrderCommissionFee:       "order_commission_fee",
	OrderCommissionFeeRefund: "order_commission_fee_refund",
	OrderPointsAdd:           "order_points_add",
	OrderActivityData:        "order_activity_data",
	OrderCancelIdentity:      "order_cancel_identity",
	OrderCancelReason:        "order_cancel_reason",
	OrderCancelTime:          "order_cancel_time",
	OrderBpAdd:               "order_bp_add",
	OrderRebate:              "order_rebate",
	BuyerMobile:              "buyer_mobile",
	BuyerContacter:           "buyer_contacter",
	ActivityManhuiId:         "activity_manhui_id",
	ActivityDoublePointsId:   "activity_double_points_id",
	OrderDoublePointsAdd:     "order_double_points_add",
	ActivityVoucherId:        "activity_voucher_id",
	OrderActivityManhuiState: "order_activity_manhui_state",
	Version:                  "version",
	PrimaryKey:               "order_id",
}

// NewOrderDataDao creates and returns a new DAO object for table data access.
func NewOrderDataDao() *OrderDataDao {
	return &OrderDataDao{
		group:   "trade",
		table:   "trade_order_data",
		columns: orderDataColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrderDataDao) Columns() OrderDataColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrderDataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrderDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *OrderDataDao) Get(ctx context.Context, id any) (one *entity.OrderData, err error) {
	var entitys []*entity.OrderData
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
func (dao *OrderDataDao) Gets(ctx context.Context, id any) (entitys []*entity.OrderData, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *OrderDataDao) Find(ctx context.Context, in *do.OrderDataListInput) (out []*entity.OrderData, err error) {
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
func (dao *OrderDataDao) FindOne(ctx context.Context, in *do.OrderDataListInput) (one *entity.OrderData, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.OrderData
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
func (dao *OrderDataDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.OrderDataListInput) (out gdb.Result, err error) {
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
func (dao *OrderDataDao) FindKey(ctx context.Context, in *do.OrderDataListInput) (out []interface{}, err error) {
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
func (dao *OrderDataDao) List(ctx context.Context, in *do.OrderDataListInput) (out *do.OrderDataListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.OrderDataListOutput{}
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
func (dao *OrderDataDao) Add(ctx context.Context, in *do.OrderData) (lastInsertId int64, err error) {
	data := do.OrderData{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *OrderDataDao) Edit(ctx context.Context, id any, in *do.OrderData) (int64, error) {
	data := do.OrderData{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *OrderDataDao) EditWhere(ctx context.Context, where *do.OrderDataListInput, in *do.OrderData) (int64, error) {
	data := do.OrderData{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *OrderDataDao) Save(ctx context.Context, in *do.OrderData) (affected int64, err error) {
	data := do.OrderData{}
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
func (dao *OrderDataDao) Saves(ctx context.Context, in []*do.OrderData) (affected int64, err error) {
	data := []do.OrderData{}
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
func (dao *OrderDataDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *OrderDataDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *OrderDataDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *OrderDataDao) RemoveWhere(ctx context.Context, where *do.OrderDataListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *OrderDataDao) Count(ctx context.Context, in *do.OrderDataListInput) (count int, err error) {
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
