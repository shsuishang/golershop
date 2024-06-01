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

// ActivityBaseDao is the data access object for table marketing_activity_base.
type ActivityBaseDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ActivityBaseColumns // columns contains all the column names of Table for convenient usage.
}

// ActivityBaseColumns defines and stores column names for table marketing_activity_base.
type ActivityBaseColumns struct {
	ActivityId                string // 活动编号
	StoreId                   string // 店铺编号
	UserId                    string // 用户编号
	ActivityName              string // 活动名称
	ActivityTitle             string // 活动标题
	ActivityRemark            string // 活动说明
	ActivityTypeId            string // 活动类型
	ActivityStarttime         string // 活动开始时间
	ActivityEndtime           string // 活动结束时间
	ActivityState             string // 活动状态(ENUM):0-未开启;1-正常;2-已结束;3-管理员关闭;4-商家关闭
	ActivityRule              string // 活动规则(JSON):不检索{rule_id:{}, rule_id:{}},统一解析规则{"requirement":{"buy":{"item":[1,2,3],"subtotal":"通过计算修正满足的条件"}},"rule":[{"total":100,"max_num":1,"item":{"1":1,"1200":3}},{"total":200,"max_num":1,"item":{"1":1,"1200":3}}]}
	ActivityEffectiveQuantity string // 已经参与数量
	ActivityType              string // 参与类型(ENUM):1-免费参与;2-积分参与;3-购买参与;4-分享参与
	ActivitySort              string // 活动排序
	ActivityIsFinish          string // 活动是否完成(ENUM):0-未完成;1-已完成;2-已解散(目前用于团购)
	SubsiteId                 string // 分站编号
	ActivityUseLevel          string // 使用等级(DOT)
	ActivityItemIds           string // 活动SKU(DOT):activity_rule中数据冗余
	ActivityAddtime           string // 添加时间
	Version                   string // 版本
	PrimaryKey                string // 主键
}

// activityBaseColumns holds the columns for table marketing_activity_base.
var activityBaseColumns = ActivityBaseColumns{
	ActivityId:                "activity_id",
	StoreId:                   "store_id",
	UserId:                    "user_id",
	ActivityName:              "activity_name",
	ActivityTitle:             "activity_title",
	ActivityRemark:            "activity_remark",
	ActivityTypeId:            "activity_type_id",
	ActivityStarttime:         "activity_starttime",
	ActivityEndtime:           "activity_endtime",
	ActivityState:             "activity_state",
	ActivityRule:              "activity_rule",
	ActivityEffectiveQuantity: "activity_effective_quantity",
	ActivityType:              "activity_type",
	ActivitySort:              "activity_sort",
	ActivityIsFinish:          "activity_is_finish",
	SubsiteId:                 "subsite_id",
	ActivityUseLevel:          "activity_use_level",
	ActivityItemIds:           "activity_item_ids",
	ActivityAddtime:           "activity_addtime",
	Version:                   "version",
	PrimaryKey:                "activity_id",
}

// NewActivityBaseDao creates and returns a new DAO object for table data access.
func NewActivityBaseDao() *ActivityBaseDao {
	return &ActivityBaseDao{
		group:   "marketing",
		table:   "marketing_activity_base",
		columns: activityBaseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivityBaseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivityBaseDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivityBaseDao) Columns() ActivityBaseColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivityBaseDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivityBaseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivityBaseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *ActivityBaseDao) Get(ctx context.Context, id any) (one *entity.ActivityBase, err error) {
	var entitys []*entity.ActivityBase
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
func (dao *ActivityBaseDao) Gets(ctx context.Context, id any) (entitys []*entity.ActivityBase, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *ActivityBaseDao) Find(ctx context.Context, in *do.ActivityBaseListInput) (out []*entity.ActivityBase, err error) {
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
func (dao *ActivityBaseDao) FindOne(ctx context.Context, in *do.ActivityBaseListInput) (one *entity.ActivityBase, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.ActivityBase
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
func (dao *ActivityBaseDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.ActivityBaseListInput) (out gdb.Result, err error) {
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
func (dao *ActivityBaseDao) FindKey(ctx context.Context, in *do.ActivityBaseListInput) (out []interface{}, err error) {
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
func (dao *ActivityBaseDao) List(ctx context.Context, in *do.ActivityBaseListInput) (out *do.ActivityBaseListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.ActivityBaseListOutput{}
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
func (dao *ActivityBaseDao) Add(ctx context.Context, in *do.ActivityBase) (lastInsertId int64, err error) {
	data := do.ActivityBase{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *ActivityBaseDao) Edit(ctx context.Context, id any, in *do.ActivityBase) (int64, error) {
	data := do.ActivityBase{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *ActivityBaseDao) EditWhere(ctx context.Context, where *do.ActivityBaseListInput, in *do.ActivityBase) (int64, error) {
	data := do.ActivityBase{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *ActivityBaseDao) Save(ctx context.Context, in *do.ActivityBase) (affected int64, err error) {
	data := do.ActivityBase{}
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
func (dao *ActivityBaseDao) Saves(ctx context.Context, in []*do.ActivityBase) (affected int64, err error) {
	data := []do.ActivityBase{}
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
func (dao *ActivityBaseDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *ActivityBaseDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *ActivityBaseDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *ActivityBaseDao) RemoveWhere(ctx context.Context, where *do.ActivityBaseListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *ActivityBaseDao) Count(ctx context.Context, in *do.ActivityBaseListInput) (count int, err error) {
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
