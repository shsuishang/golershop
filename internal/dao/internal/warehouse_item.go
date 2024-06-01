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

// WarehouseItemDao is the data access object for table invoicing_warehouse_item.
type WarehouseItemDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns WarehouseItemColumns // columns contains all the column names of Table for convenient usage.
}

// WarehouseItemColumns defines and stores column names for table invoicing_warehouse_item.
type WarehouseItemColumns struct {
	WarehouseItemId           string // 仓库商品编号
	ProductId                 string // 商品编号（冗余）
	UnitId                    string // 商品单位id（冗余）
	ItemId                    string // SKU
	CategoryId                string // 所属分类
	WarehouseId               string // 仓库编号
	WarehouseItemCost         string // 成本
	WarehouseItemQuantity     string // 数量
	WarehouseItemLockQuantity string // 锁定库存:实际库存-锁定库存=可售库存
	WarehouseItemQuantityMin  string // 最小库存
	WarehouseItemQuantityMax  string // 最大库存
	StoreId                   string // 所属店铺
	PrimaryKey                string // 主键
}

// warehouseItemColumns holds the columns for table invoicing_warehouse_item.
var warehouseItemColumns = WarehouseItemColumns{
	WarehouseItemId:           "warehouse_item_id",
	ProductId:                 "product_id",
	UnitId:                    "unit_id",
	ItemId:                    "item_id",
	CategoryId:                "category_id",
	WarehouseId:               "warehouse_id",
	WarehouseItemCost:         "warehouse_item_cost",
	WarehouseItemQuantity:     "warehouse_item_quantity",
	WarehouseItemLockQuantity: "warehouse_item_lock_quantity",
	WarehouseItemQuantityMin:  "warehouse_item_quantity_min",
	WarehouseItemQuantityMax:  "warehouse_item_quantity_max",
	StoreId:                   "store_id",
	PrimaryKey:                "warehouse_item_id",
}

// NewWarehouseItemDao creates and returns a new DAO object for table data access.
func NewWarehouseItemDao() *WarehouseItemDao {
	return &WarehouseItemDao{
		group:   "invoicing",
		table:   "invoicing_warehouse_item",
		columns: warehouseItemColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WarehouseItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WarehouseItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WarehouseItemDao) Columns() WarehouseItemColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WarehouseItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WarehouseItemDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WarehouseItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *WarehouseItemDao) Get(ctx context.Context, id any) (one *entity.WarehouseItem, err error) {
	var entitys []*entity.WarehouseItem
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
func (dao *WarehouseItemDao) Gets(ctx context.Context, id any) (entitys []*entity.WarehouseItem, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *WarehouseItemDao) Find(ctx context.Context, in *do.WarehouseItemListInput) (out []*entity.WarehouseItem, err error) {
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
func (dao *WarehouseItemDao) FindOne(ctx context.Context, in *do.WarehouseItemListInput) (one *entity.WarehouseItem, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.WarehouseItem
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
func (dao *WarehouseItemDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.WarehouseItemListInput) (out gdb.Result, err error) {
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
func (dao *WarehouseItemDao) FindKey(ctx context.Context, in *do.WarehouseItemListInput) (out []interface{}, err error) {
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
func (dao *WarehouseItemDao) List(ctx context.Context, in *do.WarehouseItemListInput) (out *do.WarehouseItemListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.WarehouseItemListOutput{}
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
func (dao *WarehouseItemDao) Add(ctx context.Context, in *do.WarehouseItem) (lastInsertId int64, err error) {
	data := do.WarehouseItem{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *WarehouseItemDao) Edit(ctx context.Context, id any, in *do.WarehouseItem) (int64, error) {
	data := do.WarehouseItem{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *WarehouseItemDao) EditWhere(ctx context.Context, where *do.WarehouseItemListInput, in *do.WarehouseItem) (int64, error) {
	data := do.WarehouseItem{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *WarehouseItemDao) Save(ctx context.Context, in *do.WarehouseItem) (affected int64, err error) {
	data := do.WarehouseItem{}
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
func (dao *WarehouseItemDao) Saves(ctx context.Context, in []*do.WarehouseItem) (affected int64, err error) {
	data := []do.WarehouseItem{}
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
func (dao *WarehouseItemDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *WarehouseItemDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *WarehouseItemDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *WarehouseItemDao) RemoveWhere(ctx context.Context, where *do.WarehouseItemListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *WarehouseItemDao) Count(ctx context.Context, in *do.WarehouseItemListInput) (count int, err error) {
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
