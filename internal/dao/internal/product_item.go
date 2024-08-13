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

// ProductItemDao is the data access object for table pt_product_item.
type ProductItemDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ProductItemColumns // columns contains all the column names of Table for convenient usage.
}

// ProductItemColumns defines and stores column names for table pt_product_item.
type ProductItemColumns struct {
	ItemId             string // 商品编号-SKU编号
	ItemName           string // 副标题(DOT):SKU名称
	ItemIndex          string // 索引(DOT)
	ProductId          string // 产品编号
	ColorId            string // 颜色SKU，规格值
	ItemIsDefault      string // 是否为默认展示的商品，必须为item_enable
	ItemNumber         string // SKU商家编码:SKU商家编码为非必填项，若不填写，系统会自动生成一个SKU商家编码。
	ItemBarcode        string // 条形码
	ItemCostPrice      string // 成本价
	ItemUnitPrice      string // 商品价格
	ItemMarketPrice    string // 市场价
	ItemUnitPoints     string // 积分价格
	ItemQuantity       string // 商品库存
	ItemQuantityFrozen string // 商品冻结库存
	ItemWarnQuantity   string // 库存预警值
	ItemSpec           string // 商品规格序列化(JSON):{spec_id:spec_item_id, spec_id:spec_item_id, spec_id:spec_item_id}
	SpecItemIds        string // 商品规格值编号
	ItemEnable         string // 是否启用(LIST):1001-正常;1002-下架仓库中;1000-违规禁售
	ItemIsChange       string // 被改动(BOOL):0-未改动;1-已改动分销使用
	ItemWeight         string // 商品重量:KG
	ItemVolume         string // 商品体积:立方米
	ItemFxCommission   string // 微小店分销佣金
	ItemRebate         string // 返利额度
	ItemSrcId          string // 供应商SKU编号
	CategoryId         string // 商品分类
	CourseCategoryId   string // 课程分类
	StoreId            string // 所属店铺
	Version            string // 版本
	PrimaryKey         string // 主键
}

// productItemColumns holds the columns for table pt_product_item.
var productItemColumns = ProductItemColumns{
	ItemId:             "item_id",
	ItemName:           "item_name",
	ItemIndex:          "item_index",
	ProductId:          "product_id",
	ColorId:            "color_id",
	ItemIsDefault:      "item_is_default",
	ItemNumber:         "item_number",
	ItemBarcode:        "item_barcode",
	ItemCostPrice:      "item_cost_price",
	ItemUnitPrice:      "item_unit_price",
	ItemMarketPrice:    "item_market_price",
	ItemUnitPoints:     "item_unit_points",
	ItemQuantity:       "item_quantity",
	ItemQuantityFrozen: "item_quantity_frozen",
	ItemWarnQuantity:   "item_warn_quantity",
	ItemSpec:           "item_spec",
	SpecItemIds:        "spec_item_ids",
	ItemEnable:         "item_enable",
	ItemIsChange:       "item_is_change",
	ItemWeight:         "item_weight",
	ItemVolume:         "item_volume",
	ItemFxCommission:   "item_fx_commission",
	ItemRebate:         "item_rebate",
	ItemSrcId:          "item_src_id",
	CategoryId:         "category_id",
	CourseCategoryId:   "course_category_id",
	StoreId:            "store_id",
	Version:            "version",
	PrimaryKey:         "item_id",
}

// NewProductItemDao creates and returns a new DAO object for table data access.
func NewProductItemDao() *ProductItemDao {
	return &ProductItemDao{
		group:   "pt",
		table:   "pt_product_item",
		columns: productItemColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductItemDao) Columns() ProductItemColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductItemDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *ProductItemDao) Get(ctx context.Context, id any) (one *entity.ProductItem, err error) {
	var entitys []*entity.ProductItem
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
func (dao *ProductItemDao) Gets(ctx context.Context, id any) (entitys []*entity.ProductItem, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	for _, item := range entitys {
		if item.ItemQuantity >= item.ItemQuantityFrozen {
			item.AvailableQuantity = item.ItemQuantity - item.ItemQuantityFrozen
		} else {
			item.AvailableQuantity = 0
		}
	}

	return entitys, err
}

// Find 查询数据
func (dao *ProductItemDao) Find(ctx context.Context, in *do.ProductItemListInput) (out []*entity.ProductItem, err error) {
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

	for _, item := range out {
		if item.ItemQuantity >= item.ItemQuantityFrozen {
			item.AvailableQuantity = item.ItemQuantity - item.ItemQuantityFrozen
		} else {
			item.AvailableQuantity = 0
		}
	}

	return out, nil
}

// FindOne 查询一条数据
func (dao *ProductItemDao) FindOne(ctx context.Context, in *do.ProductItemListInput) (one *entity.ProductItem, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.ProductItem
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
func (dao *ProductItemDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.ProductItemListInput) (out gdb.Result, err error) {
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
func (dao *ProductItemDao) FindKey(ctx context.Context, in *do.ProductItemListInput) (out []interface{}, err error) {
	idRes, err := dao.FindFields(ctx, dao.Columns().PrimaryKey, in)

	if err != nil {
		return nil, err
	}

	for _, record := range idRes {
		if !record[dao.Columns().PrimaryKey].IsEmpty() {
			out = append(out, record[dao.Columns().PrimaryKey].Uint64())
		}
	}

	return out, err
}

// List 分页读取
func (dao *ProductItemDao) List(ctx context.Context, in *do.ProductItemListInput) (out *do.ProductItemListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.ProductItemListOutput{}
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

	for _, item := range out.Items {
		if item.ItemQuantity >= item.ItemQuantityFrozen {
			item.AvailableQuantity = item.ItemQuantity - item.ItemQuantityFrozen
		} else {
			item.AvailableQuantity = 0
		}
	}

	return out, nil
}

// Add 新增
func (dao *ProductItemDao) Add(ctx context.Context, in *do.ProductItem) (lastInsertId int64, err error) {
	data := do.ProductItem{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *ProductItemDao) Edit(ctx context.Context, id any, in *do.ProductItem) (int64, error) {
	data := do.ProductItem{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *ProductItemDao) EditWhere(ctx context.Context, where *do.ProductItemListInput, in *do.ProductItem) (int64, error) {
	data := do.ProductItem{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *ProductItemDao) Save(ctx context.Context, in *do.ProductItem) (affected int64, err error) {
	data := do.ProductItem{}
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
func (dao *ProductItemDao) Saves(ctx context.Context, in []*do.ProductItem) (affected int64, err error) {
	data := []do.ProductItem{}
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
func (dao *ProductItemDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *ProductItemDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *ProductItemDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *ProductItemDao) RemoveWhere(ctx context.Context, where *do.ProductItemListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *ProductItemDao) Count(ctx context.Context, in *do.ProductItemListInput) (count int, err error) {
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
