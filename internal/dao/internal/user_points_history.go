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

// UserPointsHistoryDao is the data access object for table pay_user_points_history.
type UserPointsHistoryDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns UserPointsHistoryColumns // columns contains all the column names of Table for convenient usage.
}

// UserPointsHistoryColumns defines and stores column names for table pay_user_points_history.
type UserPointsHistoryColumns struct {
	PointsLogId     string // 记录编号
	PointsKindId    string // 类型(ENUM):1-获取积分;2-消费积分;
	PointsTypeId    string // 积分类型(ENUM):1-会员注册;2-会员登录;3-商品评论;4-购买商品;5-管理员操作;7-积分换购商品;8-积分兑换代金券
	UserId          string // 会员编号
	PointsLogPoints string // 可用积分
	UserPoints      string // 当前积分
	PointsLogTime   string // 创建时间
	PointsLogDesc   string // 描述
	StoreId         string // 所属店铺
	PointsLogDate   string // 积分日期
	UserIdOther     string // 交互会员
	PointsLogState  string // 领取状态(BOOL):0-未领取;1-已领取
	ExtId           string // 关联单号
	PrimaryKey      string // 主键
}

// userPointsHistoryColumns holds the columns for table pay_user_points_history.
var userPointsHistoryColumns = UserPointsHistoryColumns{
	PointsLogId:     "points_log_id",
	PointsKindId:    "points_kind_id",
	PointsTypeId:    "points_type_id",
	UserId:          "user_id",
	PointsLogPoints: "points_log_points",
	UserPoints:      "user_points",
	PointsLogTime:   "points_log_time",
	PointsLogDesc:   "points_log_desc",
	StoreId:         "store_id",
	PointsLogDate:   "points_log_date",
	UserIdOther:     "user_id_other",
	PointsLogState:  "points_log_state",
	ExtId:           "ext_id",
	PrimaryKey:      "points_log_id",
}

// NewUserPointsHistoryDao creates and returns a new DAO object for table data access.
func NewUserPointsHistoryDao() *UserPointsHistoryDao {
	return &UserPointsHistoryDao{
		group:   "pay",
		table:   "pay_user_points_history",
		columns: userPointsHistoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserPointsHistoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserPointsHistoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserPointsHistoryDao) Columns() UserPointsHistoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserPointsHistoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserPointsHistoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserPointsHistoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *UserPointsHistoryDao) Get(ctx context.Context, id any) (one *entity.UserPointsHistory, err error) {
	var entitys []*entity.UserPointsHistory
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
func (dao *UserPointsHistoryDao) Gets(ctx context.Context, id any) (entitys []*entity.UserPointsHistory, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *UserPointsHistoryDao) Find(ctx context.Context, in *do.UserPointsHistoryListInput) (out []*entity.UserPointsHistory, err error) {
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
func (dao *UserPointsHistoryDao) FindOne(ctx context.Context, in *do.UserPointsHistoryListInput) (one *entity.UserPointsHistory, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.UserPointsHistory
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
func (dao *UserPointsHistoryDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.UserPointsHistoryListInput) (out gdb.Result, err error) {
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
func (dao *UserPointsHistoryDao) FindKey(ctx context.Context, in *do.UserPointsHistoryListInput) (out []interface{}, err error) {
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
func (dao *UserPointsHistoryDao) List(ctx context.Context, in *do.UserPointsHistoryListInput) (out *do.UserPointsHistoryListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.UserPointsHistoryListOutput{}
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
func (dao *UserPointsHistoryDao) Add(ctx context.Context, in *do.UserPointsHistory) (lastInsertId int64, err error) {
	data := do.UserPointsHistory{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *UserPointsHistoryDao) Edit(ctx context.Context, id any, in *do.UserPointsHistory) (int64, error) {
	data := do.UserPointsHistory{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *UserPointsHistoryDao) EditWhere(ctx context.Context, where *do.UserPointsHistoryListInput, in *do.UserPointsHistory) (int64, error) {
	data := do.UserPointsHistory{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *UserPointsHistoryDao) Save(ctx context.Context, in *do.UserPointsHistory) (affected int64, err error) {
	data := do.UserPointsHistory{}
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
func (dao *UserPointsHistoryDao) Saves(ctx context.Context, in []*do.UserPointsHistory) (affected int64, err error) {
	data := []do.UserPointsHistory{}
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
func (dao *UserPointsHistoryDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *UserPointsHistoryDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *UserPointsHistoryDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *UserPointsHistoryDao) RemoveWhere(ctx context.Context, where *do.UserPointsHistoryListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *UserPointsHistoryDao) Count(ctx context.Context, in *do.UserPointsHistoryListInput) (count int, err error) {
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
