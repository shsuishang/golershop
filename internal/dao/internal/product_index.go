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

// ProductIndexDao is the data access object for table pt_product_index.
type ProductIndexDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ProductIndexColumns // columns contains all the column names of Table for convenient usage.
}

// ProductIndexColumns defines and stores column names for table pt_product_index.
type ProductIndexColumns struct {
	ProductId                string // 产品编号:定为SPU编号
	ProductNumber            string // SPU商家编码:货号
	ProductName              string // 产品名称
	ProductNameIndex         string // 名称索引关键字(DOT)
	StoreId                  string // 店铺编号
	StoreIsOpen              string // 店铺状态(BOOL):0-关闭;1-运营中
	StoreType                string // 店铺类型(ENUM): 1-卖家店铺; 2-供应商店铺
	StoreCategoryIds         string // 店铺分类(DOT)
	CategoryId               string // 商品分类
	CourseCategoryId         string // 课程分类
	TypeId                   string // 类型编号:冗余检索
	ProductQuantity          string // 商品库存:冗余计算
	ProductWarnQuantity      string // 预警数量
	BrandId                  string // 品牌编号
	ProductServiceTypeIds    string // 售后服务(DOT)
	ProductStateId           string // 商品状态(ENUM):1001-正常;1002-下架仓库中;1003-待审核; 1000-违规禁售
	ProductSaleDistrictIds   string // 销售区域(DOT): district_id=1000全部区域
	ProductVerifyId          string // 商品审核(ENUM):3001-审核通过;3002-审核中;3000-审核未通过
	ProductIsInvoices        string // 是否开票(BOOL): 1-是; 0-否
	ProductIsReturn          string // 允许退换货(BOOL): 1-是; 0-否
	ProductIsRecommend       string // 商品推荐(BOOL):1-是; 0-否
	ProductStockStatus       string // 缺货状态(ENUM):1-有现货;2-预售商品;3-缺货;4-2至3天
	KindId                   string // 商品种类:1201-实物;1202-虚拟
	ActivityTypeIds          string // 参与活动(DOT)
	ContractTypeIds          string // 消费者保障(DOT):由店铺映射到商品
	ProductAssistData        string // 辅助属性值列(DOT):assist_item_id每个都不用 , setFilter(tagid, array(2,3,4));是表示含有标签值2,3,4中的任意一个即符合筛选，这里是or关系。 setFilter(‘tagid’, array(2)); setFilter(‘tagid’, array(3)); 形成and关系| msyql where FIND_IN_SET('1', product_assist_data)
	ProductUnitPriceMin      string // 最低单价
	ProductUnitPriceMax      string // 最高单价
	ProductUnitPointsMin     string // 商品积分
	ProductUnitPointsMax     string // 商品积分
	ProductSaleNum           string // 销售数量
	ProductFavoriteNum       string // 收藏数量
	ProductClick             string // 点击数量
	ProductEvaluationNum     string // 评价次数
	ProductRegionDistrictIds string // 所属区域(DOT)
	ProductFreight           string // 运费:包邮为0
	ProductTags              string // 商品标签(DOT)
	StoreIsSelfsupport       string // 是否自营(BOOL):1-自营;0-非自营
	ProductSpEnable          string // 允许分销(BOOL):1-启用分销;0-禁用分销
	ProductDistEnable        string // 三级分销允许分销(BOOL):1-启用分销;0-禁用分销
	ProductAddTime           string // 添加时间
	ProductSaleTime          string // 上架时间:预设上架时间,可以动态修正状态
	ProductOrder             string // 排序:越小越靠前
	ProductSrcId             string // 产品来源编号
	MarketCategoryId         string // 所属商圈(DOT)
	StoreLatitude            string // 纬度
	StoreLongitude           string // 经度
	ProductIsVideo           string // 是否视频(BOOL):1-有视频;0-无视频
	ProductTransportId       string // 配送服务(ENUM):1001-快递发货;1002-到店自提;1003-上门服务
	SubsiteId                string // 所属分站:0-总站
	ProductIsLock            string // 是否锁定(BOOL):0-未锁定; 1-锁定,参加团购的商品不予许修改
	ProductInventoryLock     string // 库存锁定(ENUM):1001-下单锁定;1002-支付锁定;
	ProductFrom              string // 商品来源(ENUM):1000-发布;1001-天猫;1002-淘宝;1003-阿里巴巴;1004-京东;
	Version                  string // 乐观锁
	PrimaryKey               string // 主键
}

// productIndexColumns holds the columns for table pt_product_index.
var productIndexColumns = ProductIndexColumns{
	ProductId:                "product_id",
	ProductNumber:            "product_number",
	ProductName:              "product_name",
	ProductNameIndex:         "product_name_index",
	StoreId:                  "store_id",
	StoreIsOpen:              "store_is_open",
	StoreType:                "store_type",
	StoreCategoryIds:         "store_category_ids",
	CategoryId:               "category_id",
	CourseCategoryId:         "course_category_id",
	TypeId:                   "type_id",
	ProductQuantity:          "product_quantity",
	ProductWarnQuantity:      "product_warn_quantity",
	BrandId:                  "brand_id",
	ProductServiceTypeIds:    "product_service_type_ids",
	ProductStateId:           "product_state_id",
	ProductSaleDistrictIds:   "product_sale_district_ids",
	ProductVerifyId:          "product_verify_id",
	ProductIsInvoices:        "product_is_invoices",
	ProductIsReturn:          "product_is_return",
	ProductIsRecommend:       "product_is_recommend",
	ProductStockStatus:       "product_stock_status",
	KindId:                   "kind_id",
	ActivityTypeIds:          "activity_type_ids",
	ContractTypeIds:          "contract_type_ids",
	ProductAssistData:        "product_assist_data",
	ProductUnitPriceMin:      "product_unit_price_min",
	ProductUnitPriceMax:      "product_unit_price_max",
	ProductUnitPointsMin:     "product_unit_points_min",
	ProductUnitPointsMax:     "product_unit_points_max",
	ProductSaleNum:           "product_sale_num",
	ProductFavoriteNum:       "product_favorite_num",
	ProductClick:             "product_click",
	ProductEvaluationNum:     "product_evaluation_num",
	ProductRegionDistrictIds: "product_region_district_ids",
	ProductFreight:           "product_freight",
	ProductTags:              "product_tags",
	StoreIsSelfsupport:       "store_is_selfsupport",
	ProductSpEnable:          "product_sp_enable",
	ProductDistEnable:        "product_dist_enable",
	ProductAddTime:           "product_add_time",
	ProductSaleTime:          "product_sale_time",
	ProductOrder:             "product_order",
	ProductSrcId:             "product_src_id",
	MarketCategoryId:         "market_category_id",
	StoreLatitude:            "store_latitude",
	StoreLongitude:           "store_longitude",
	ProductIsVideo:           "product_is_video",
	ProductTransportId:       "product_transport_id",
	SubsiteId:                "subsite_id",
	ProductIsLock:            "product_is_lock",
	ProductInventoryLock:     "product_inventory_lock",
	ProductFrom:              "product_from",
	Version:                  "version",
	PrimaryKey:               "product_id",
}

// NewProductIndexDao creates and returns a new DAO object for table data access.
func NewProductIndexDao() *ProductIndexDao {
	return &ProductIndexDao{
		group:   "pt",
		table:   "pt_product_index",
		columns: productIndexColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProductIndexDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProductIndexDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProductIndexDao) Columns() ProductIndexColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProductIndexDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProductIndexDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProductIndexDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *ProductIndexDao) Get(ctx context.Context, id any) (one *entity.ProductIndex, err error) {
	var entitys []*entity.ProductIndex
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
func (dao *ProductIndexDao) Gets(ctx context.Context, id any) (entitys []*entity.ProductIndex, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *ProductIndexDao) Find(ctx context.Context, in *do.ProductIndexListInput) (out []*entity.ProductIndex, err error) {
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
func (dao *ProductIndexDao) FindOne(ctx context.Context, in *do.ProductIndexListInput) (one *entity.ProductIndex, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.ProductIndex
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
func (dao *ProductIndexDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.ProductIndexListInput) (out gdb.Result, err error) {
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
func (dao *ProductIndexDao) FindKey(ctx context.Context, in *do.ProductIndexListInput) (out []interface{}, err error) {
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
func (dao *ProductIndexDao) List(ctx context.Context, in *do.ProductIndexListInput) (out *do.ProductIndexListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.ProductIndexListOutput{}
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
func (dao *ProductIndexDao) Add(ctx context.Context, in *do.ProductIndex) (lastInsertId int64, err error) {
	data := do.ProductIndex{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *ProductIndexDao) Edit(ctx context.Context, id any, in *do.ProductIndex) (int64, error) {
	data := do.ProductIndex{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *ProductIndexDao) EditWhere(ctx context.Context, where *do.ProductIndexListInput, in *do.ProductIndex) (int64, error) {
	data := do.ProductIndex{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *ProductIndexDao) Save(ctx context.Context, in *do.ProductIndex) (affected int64, err error) {
	data := do.ProductIndex{}
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
func (dao *ProductIndexDao) Saves(ctx context.Context, in []*do.ProductIndex) (affected int64, err error) {
	data := []do.ProductIndex{}
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
func (dao *ProductIndexDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *ProductIndexDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *ProductIndexDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *ProductIndexDao) RemoveWhere(ctx context.Context, where *do.ProductIndexListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *ProductIndexDao) Count(ctx context.Context, in *do.ProductIndexListInput) (count int, err error) {
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
