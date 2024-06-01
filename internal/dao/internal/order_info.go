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

// OrderInfoDao is the data access object for table trade_order_info.
type OrderInfoDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns OrderInfoColumns // columns contains all the column names of Table for convenient usage.
}

// OrderInfoColumns defines and stores column names for table trade_order_info.
type OrderInfoColumns struct {
	OrderId                     string // 订单编号
	OrderTitle                  string // 订单标题
	StoreId                     string // 卖家店铺编号
	SubsiteId                   string // 所属分站:0-总站
	UserId                      string // 买家编号
	KindId                      string // 订单种类(ENUM): 1201-实物 ; 1202-教育类 ; 1203-电子卡券  ; 1204-其它
	OrderLockStatus             string // 锁定状态(BOOL):0-正常;1-锁定,退款退货
	OrderIsSettlemented         string // 订单是否结算(BOOL):0-未结算; 1-已结算
	OrderSettlementTime         string // 订单结算时间
	OrderBuyerEvaluationStatus  string // 买家针对订单对店铺评价(ENUM): 0-未评价;1-已评价;  2-已过期未评价
	OrderSellerEvaluationStatus string // 卖家评价状态(ENUM):0-未评价;1-已评价;  2-已过期未评价
	OrderBuyerHidden            string // 买家删除(BOOL): 1-是; 0-否
	OrderShopHidden             string // 店铺删除(BOOL): 1-是; 0-否
	PaymentTypeId               string // 支付方式(ENUM): 1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付;
	PaymentTime                 string // 付款时间
	OrderStateId                string // 订单状态(LIST):2011-待订单审核;2013-待财务审核;2020-待配货/待出库审核;2030-待发货;2040-已发货/待收货确认;2060-已完成/已签收;2070-已取消/已作废;
	OrderIsReview               string // 订单审核(BOOL):0-未审核;1-已审核;
	OrderFinanceReview          string // 财务状态(BOOL):0-未审核;1-已审核
	OrderIsPaid                 string // 付款状态(ENUM):3010-未付款;3011-付款待审核;3012-部分付款;3013-已付款
	OrderIsOut                  string // 出库状态(ENUM):3020-未出库;3021-部分出库通过拆单解决这种问题;3022-已出库
	OrderIsShipped              string // 发货状态(ENUM):3030-未发货;3032-已发货;3031-部分发货
	OrderIsReceived             string // 收货状态(BOOL):0-未收货;1-已收货
	OrderReceivedTime           string // 订单签收时间
	ChainId                     string // 门店编号
	DeliveryTypeId              string // 配送方式
	OrderIsOffline              string // 线下订单(BOOL):0-线上;1-线下
	OrderExpressPrint           string // 是否打印(BOOL):0-未打印;1-已打印
	ActivityId                  string // 活动编号(DOT)
	ActivityTypeId              string // 活动类型(DOT)
	SalespersonId               string // 销售员编号:用户编号
	OrderIsSync                 string // 是否ERP同步(BOOL):0-未同步; 1-已同步
	StoreIsSelfsupport          string // 是否自营(ENUM): 1-自营;0-非自营
	StoreType                   string // 店铺类型(ENUM): 1-卖家店铺; 2-供应商店铺
	OrderErpId                  string // ERP订单编号
	DistributorUserId           string // 分销商编号:用户编号
	OrderIsCb                   string // 跨境订单(BOOL):0-否; 1-是
	OrderIsCbSync               string // 是否报关(BOOL):0-否; 1-是
	SrcOrderId                  string // 来源订单
	OrderIsTransfer             string // 是否代发(BOOL):0-否; 1-是
	OrderIsTransferNote         string // 转单执行结果
	OrderFxIsSettlemented       string // 佣金是否发放(BOOL):0 -未发放;1 -已发放
	OrderFxSettlementTime       string // 佣金结算时间
	OrderType                   string // 订单类型(ENUM)
	OrderWithdrawConfirm        string // 提现审核(BOOL):0-未审核; 1-已审核
	PaymentFormId               string // 支付方式(BOOL):1-先预约后支付;0-先支付后预约
	CartTypeId                  string // 购买类型(ENUM):1-购买; 2-积分兑换; 3-赠品; 4-活动促销
	CardKindId                  string // 商品绑定卡片类型(ENUM): 1001-次卡商品; 1002-优惠券商品;1003-券码商品;
	CreateTime                  string // 下单时间:检索使用
	UpdateTime                  string // 当前状态的处理时间
	Version                     string // 乐观锁
	PrimaryKey                  string // 主键
}

// orderInfoColumns holds the columns for table trade_order_info.
var orderInfoColumns = OrderInfoColumns{
	OrderId:                     "order_id",
	OrderTitle:                  "order_title",
	StoreId:                     "store_id",
	SubsiteId:                   "subsite_id",
	UserId:                      "user_id",
	KindId:                      "kind_id",
	OrderLockStatus:             "order_lock_status",
	OrderIsSettlemented:         "order_is_settlemented",
	OrderSettlementTime:         "order_settlement_time",
	OrderBuyerEvaluationStatus:  "order_buyer_evaluation_status",
	OrderSellerEvaluationStatus: "order_seller_evaluation_status",
	OrderBuyerHidden:            "order_buyer_hidden",
	OrderShopHidden:             "order_shop_hidden",
	PaymentTypeId:               "payment_type_id",
	PaymentTime:                 "payment_time",
	OrderStateId:                "order_state_id",
	OrderIsReview:               "order_is_review",
	OrderFinanceReview:          "order_finance_review",
	OrderIsPaid:                 "order_is_paid",
	OrderIsOut:                  "order_is_out",
	OrderIsShipped:              "order_is_shipped",
	OrderIsReceived:             "order_is_received",
	OrderReceivedTime:           "order_received_time",
	ChainId:                     "chain_id",
	DeliveryTypeId:              "delivery_type_id",
	OrderIsOffline:              "order_is_offline",
	OrderExpressPrint:           "order_express_print",
	ActivityId:                  "activity_id",
	ActivityTypeId:              "activity_type_id",
	SalespersonId:               "salesperson_id",
	OrderIsSync:                 "order_is_sync",
	StoreIsSelfsupport:          "store_is_selfsupport",
	StoreType:                   "store_type",
	OrderErpId:                  "order_erp_id",
	DistributorUserId:           "distributor_user_id",
	OrderIsCb:                   "order_is_cb",
	OrderIsCbSync:               "order_is_cb_sync",
	SrcOrderId:                  "src_order_id",
	OrderIsTransfer:             "order_is_transfer",
	OrderIsTransferNote:         "order_is_transfer_note",
	OrderFxIsSettlemented:       "order_fx_is_settlemented",
	OrderFxSettlementTime:       "order_fx_settlement_time",
	OrderType:                   "order_type",
	OrderWithdrawConfirm:        "order_withdraw_confirm",
	PaymentFormId:               "payment_form_id",
	CartTypeId:                  "cart_type_id",
	CardKindId:                  "card_kind_id",
	CreateTime:                  "create_time",
	UpdateTime:                  "update_time",
	Version:                     "version",
	PrimaryKey:                  "order_id",
}

// NewOrderInfoDao creates and returns a new DAO object for table data access.
func NewOrderInfoDao() *OrderInfoDao {
	return &OrderInfoDao{
		group:   "trade",
		table:   "trade_order_info",
		columns: orderInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrderInfoDao) Columns() OrderInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrderInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrderInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *OrderInfoDao) Get(ctx context.Context, id any) (one *entity.OrderInfo, err error) {
	var entitys []*entity.OrderInfo
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
func (dao *OrderInfoDao) Gets(ctx context.Context, id any) (entitys []*entity.OrderInfo, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *OrderInfoDao) Find(ctx context.Context, in *do.OrderInfoListInput) (out []*entity.OrderInfo, err error) {
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
func (dao *OrderInfoDao) FindOne(ctx context.Context, in *do.OrderInfoListInput) (one *entity.OrderInfo, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.OrderInfo
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
func (dao *OrderInfoDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.OrderInfoListInput) (out gdb.Result, err error) {
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
func (dao *OrderInfoDao) FindKey(ctx context.Context, in *do.OrderInfoListInput) (out []interface{}, err error) {
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
func (dao *OrderInfoDao) List(ctx context.Context, in *do.OrderInfoListInput) (out *do.OrderInfoListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.OrderInfoListOutput{}
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
func (dao *OrderInfoDao) Add(ctx context.Context, in *do.OrderInfo) (lastInsertId int64, err error) {
	data := do.OrderInfo{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *OrderInfoDao) Edit(ctx context.Context, id any, in *do.OrderInfo) (int64, error) {
	data := do.OrderInfo{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *OrderInfoDao) EditWhere(ctx context.Context, where *do.OrderInfoListInput, in *do.OrderInfo) (int64, error) {
	data := do.OrderInfo{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *OrderInfoDao) Save(ctx context.Context, in *do.OrderInfo) (affected int64, err error) {
	data := do.OrderInfo{}
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
func (dao *OrderInfoDao) Saves(ctx context.Context, in []*do.OrderInfo) (affected int64, err error) {
	data := []do.OrderInfo{}
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
func (dao *OrderInfoDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *OrderInfoDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *OrderInfoDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *OrderInfoDao) RemoveWhere(ctx context.Context, where *do.OrderInfoListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *OrderInfoDao) Count(ctx context.Context, in *do.OrderInfoListInput) (count int, err error) {
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
