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

// OrderItemDao is the data access object for table trade_order_item.
type OrderItemDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns OrderItemColumns // columns contains all the column names of Table for convenient usage.
}

// OrderItemColumns defines and stores column names for table trade_order_item.
type OrderItemColumns struct {
	OrderItemId                  string // 编号
	OrderId                      string // 订单编号
	UserId                       string // 买家编号
	StoreId                      string // 店铺编号
	ProductId                    string // 产品编号
	ProductName                  string // 商品名称
	ItemId                       string // 货品编号
	ItemName                     string // 商品名称
	CategoryId                   string // 分类编号
	ItemCostPrice                string // 成本价
	ItemUnitPrice                string // 商品价格单价
	ItemUnitPoints               string // 资源1单价
	ItemUnitSp                   string // 资源2单价
	OrderItemSalePrice           string // 商品实际成交价单价
	OrderItemQuantity            string // 商品数量
	OrderItemInventoryLock       string // 库存锁定(ENUM):1001-下单锁定;1002-支付锁定;
	OrderItemImage               string // 商品图片
	OrderItemReturnNum           string // 退货数量
	OrderItemReturnSubtotal      string // 退款总额
	OrderItemReturnAgreeAmount   string // 退款金额:同意额度
	OrderItemAmount              string // 商品实际总金额: order_item_sale_price * order_item_quantity
	OrderItemDiscountAmount      string // 优惠金额:只考虑单品的，订单及店铺总活动优惠不影响
	OrderItemAdjustFee           string // 手工调整金额
	OrderItemPointsFee           string // 积分费用
	OrderItemPointsAdd           string // 赠送积分
	OrderItemPaymentAmount       string // 实付金额: order_item_payment_amount =  order_item_amount - order_item_discount_amount - order_item_adjust_fee - order_item_point_fee
	OrderItemEvaluationStatus    string // 评价状态(ENUM): 0-未评价;1-已评价;2-失效评价
	ActivityTypeId               string // 活动类型(ENUM):0-默认;1101-加价购=搭配宝;1102-店铺满赠-小礼品;1103-限时折扣;1104-优惠套装;1105-店铺代金券coupon优惠券;1106-拼团;1107-满减送;1108-阶梯价;1109-积分换购
	ActivityId                   string // 促销活动ID:与activity_type_id搭配使用, 团购ID/限时折扣ID/优惠套装ID/积分兑换编号
	ActivityCode                 string // 礼包活动对应兑换码code
	OrderItemCommissionRate      string // 分佣金比例百分比
	OrderItemCommissionFee       string // 佣金
	OrderItemCommissionFeeRefund string // 退款佣金
	PolicyDiscountrate           string // 价格策略折扣率
	OrderItemVoucher             string // 分配优惠券额度
	OrderItemReduce              string // 分配满减额度
	OrderItemNote                string // 备注
	OrderItemFile                string // 订单附件
	OrderItemConfirmFile         string // 商家附件
	OrderItemConfirmStatus       string // 买家确认状态(BOOL):0-为确认;1-已确认
	OrderItemSalerId             string // 单品分销者编号
	ItemSrcId                    string // 分销商品编号
	OrderItemSupplierSync        string // 拆单同步状态(BOOL):0-未同步;1-已同步
	SrcOrderId                   string // 来源订单
	OrderItemReturnAgreeNum      string // 同意退货数量
	OrderGiveId                  string // 满返优惠券id
	Version                      string // 版本
	PrimaryKey                   string // 主键
}

// orderItemColumns holds the columns for table trade_order_item.
var orderItemColumns = OrderItemColumns{
	OrderItemId:                  "order_item_id",
	OrderId:                      "order_id",
	UserId:                       "user_id",
	StoreId:                      "store_id",
	ProductId:                    "product_id",
	ProductName:                  "product_name",
	ItemId:                       "item_id",
	ItemName:                     "item_name",
	CategoryId:                   "category_id",
	ItemCostPrice:                "item_cost_price",
	ItemUnitPrice:                "item_unit_price",
	ItemUnitPoints:               "item_unit_points",
	ItemUnitSp:                   "item_unit_sp",
	OrderItemSalePrice:           "order_item_sale_price",
	OrderItemQuantity:            "order_item_quantity",
	OrderItemInventoryLock:       "order_item_inventory_lock",
	OrderItemImage:               "order_item_image",
	OrderItemReturnNum:           "order_item_return_num",
	OrderItemReturnSubtotal:      "order_item_return_subtotal",
	OrderItemReturnAgreeAmount:   "order_item_return_agree_amount",
	OrderItemAmount:              "order_item_amount",
	OrderItemDiscountAmount:      "order_item_discount_amount",
	OrderItemAdjustFee:           "order_item_adjust_fee",
	OrderItemPointsFee:           "order_item_points_fee",
	OrderItemPointsAdd:           "order_item_points_add",
	OrderItemPaymentAmount:       "order_item_payment_amount",
	OrderItemEvaluationStatus:    "order_item_evaluation_status",
	ActivityTypeId:               "activity_type_id",
	ActivityId:                   "activity_id",
	ActivityCode:                 "activity_code",
	OrderItemCommissionRate:      "order_item_commission_rate",
	OrderItemCommissionFee:       "order_item_commission_fee",
	OrderItemCommissionFeeRefund: "order_item_commission_fee_refund",
	PolicyDiscountrate:           "policy_discountrate",
	OrderItemVoucher:             "order_item_voucher",
	OrderItemReduce:              "order_item_reduce",
	OrderItemNote:                "order_item_note",
	OrderItemFile:                "order_item_file",
	OrderItemConfirmFile:         "order_item_confirm_file",
	OrderItemConfirmStatus:       "order_item_confirm_status",
	OrderItemSalerId:             "order_item_saler_id",
	ItemSrcId:                    "item_src_id",
	OrderItemSupplierSync:        "order_item_supplier_sync",
	SrcOrderId:                   "src_order_id",
	OrderItemReturnAgreeNum:      "order_item_return_agree_num",
	OrderGiveId:                  "order_give_id",
	Version:                      "version",
	PrimaryKey:                   "order_item_id",
}

// NewOrderItemDao creates and returns a new DAO object for table data access.
func NewOrderItemDao() *OrderItemDao {
	return &OrderItemDao{
		group:   "trade",
		table:   "trade_order_item",
		columns: orderItemColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrderItemDao) Columns() OrderItemColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrderItemDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrderItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *OrderItemDao) Get(ctx context.Context, id any) (one *entity.OrderItem, err error) {
	var entitys []*entity.OrderItem
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
func (dao *OrderItemDao) Gets(ctx context.Context, id any) (entitys []*entity.OrderItem, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *OrderItemDao) Find(ctx context.Context, in *do.OrderItemListInput) (out []*entity.OrderItem, err error) {
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
func (dao *OrderItemDao) FindOne(ctx context.Context, in *do.OrderItemListInput) (one *entity.OrderItem, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.OrderItem
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
func (dao *OrderItemDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.OrderItemListInput) (out gdb.Result, err error) {
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
func (dao *OrderItemDao) FindKey(ctx context.Context, in *do.OrderItemListInput) (out []interface{}, err error) {
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
func (dao *OrderItemDao) List(ctx context.Context, in *do.OrderItemListInput) (out *do.OrderItemListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.OrderItemListOutput{}
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
func (dao *OrderItemDao) Add(ctx context.Context, in *do.OrderItem) (lastInsertId int64, err error) {
	data := do.OrderItem{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *OrderItemDao) Edit(ctx context.Context, id any, in *do.OrderItem) (int64, error) {
	data := do.OrderItem{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *OrderItemDao) EditWhere(ctx context.Context, where *do.OrderItemListInput, in *do.OrderItem) (int64, error) {
	data := do.OrderItem{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *OrderItemDao) Save(ctx context.Context, in *do.OrderItem) (affected int64, err error) {
	data := do.OrderItem{}
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
func (dao *OrderItemDao) Saves(ctx context.Context, in []*do.OrderItem) (affected int64, err error) {
	data := []do.OrderItem{}
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
func (dao *OrderItemDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *OrderItemDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *OrderItemDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *OrderItemDao) RemoveWhere(ctx context.Context, where *do.OrderItemListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *OrderItemDao) Count(ctx context.Context, in *do.OrderItemListInput) (count int, err error) {
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
