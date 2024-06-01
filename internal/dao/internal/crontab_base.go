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

// CrontabBaseDao is the data access object for table sys_crontab_base.
type CrontabBaseDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns CrontabBaseColumns // columns contains all the column names of Table for convenient usage.
}

// CrontabBaseColumns defines and stores column names for table sys_crontab_base.
type CrontabBaseColumns struct {
	CrontabId          string // 任务编号
	CrontabName        string // 任务名称
	CrontabFile        string // 任务脚本
	CrontabLastExeTime string // 上次执行时间
	CrontabNextExeTime string // 下次执行时间
	CrontabMinute      string // 分钟(LIST):*-每分; 0-0;1-1; 2-2; 3-3; 4-4; 5-5; 6-6; 7-7; 8-8; 9-9; 10-10; 11-11; 12-12; 13-13; 14-14; 15-15; 16-16; 17-17; 18-18; 19-19; 20-20; 21-21; 22-22; 23-23; 24-24; 25-25; 26-26; 27-27; 28-28; 29-29; 30-30; 31-31; 32-32; 33-33; 34-34; 35-35; 36-36; 37-37; 38-38; 39-39; 40-40; 41-41; 42-42; 43-43; 44-44; 45-45; 46-46; 47-47; 48-48; 49-49; 50-50; 51-51; 52-52; 53-53; 54-54; 55-55; 56-56; 57-57; 58-58; 59-59
	CrontabHour        string // 小时(LIST):*-任意; 0-0;1-1; 2-2; 3-3; 4-4; 5-5; 6-6; 7-7; 8-8; 9-9; 10-10; 11-11; 12-12; 13-13; 14-14; 15-15; 16-16; 17-17; 18-18; 19-19; 20-20; 21-21; 22-22; 23-23
	CrontabDay         string // 每天(LIST):*-任意;1-1; 2-2; 3-3; 4-4; 5-5; 6-6; 7-7; 8-8; 9-9; 10-10; 11-11; 12-12; 13-13; 14-14; 15-15; 16-16; 17-17; 18-18; 19-19; 20-20; 21-21; 22-22; 23-23; 24-24; 25-25; 26-26; 27-27; 28-28; 29-29; 30-30; 31-31
	CrontabMonth       string // 每月(LIST):*-任意;1-1月; 2-2月; 3-3月; 4-4月; 5-5月; 6-6月; 7-7月; 8-8月; 9-9月; 10-10月; 11-11月; 12-12月
	CrontabWeek        string // 每周(LIST):*-每周;0-周日; 1-周一;2-周二;3-周三;4-周四;5-周五;6-周六
	CrontabEnable      string // 是否启用(BOOL):0-禁用; 1-启用
	CrontabBuildin     string // 是否内置(BOOL):0-否; 1-是
	CrontabRemark      string // 任务备注
	PrimaryKey         string // 主键
}

// crontabBaseColumns holds the columns for table sys_crontab_base.
var crontabBaseColumns = CrontabBaseColumns{
	CrontabId:          "crontab_id",
	CrontabName:        "crontab_name",
	CrontabFile:        "crontab_file",
	CrontabLastExeTime: "crontab_last_exe_time",
	CrontabNextExeTime: "crontab_next_exe_time",
	CrontabMinute:      "crontab_minute",
	CrontabHour:        "crontab_hour",
	CrontabDay:         "crontab_day",
	CrontabMonth:       "crontab_month",
	CrontabWeek:        "crontab_week",
	CrontabEnable:      "crontab_enable",
	CrontabBuildin:     "crontab_buildin",
	CrontabRemark:      "crontab_remark",
	PrimaryKey:         "crontab_id",
}

// NewCrontabBaseDao creates and returns a new DAO object for table data access.
func NewCrontabBaseDao() *CrontabBaseDao {
	return &CrontabBaseDao{
		group:   "sys",
		table:   "sys_crontab_base",
		columns: crontabBaseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CrontabBaseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CrontabBaseDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CrontabBaseDao) Columns() CrontabBaseColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CrontabBaseDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CrontabBaseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CrontabBaseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *CrontabBaseDao) Get(ctx context.Context, id any) (one *entity.CrontabBase, err error) {
	var entitys []*entity.CrontabBase
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
func (dao *CrontabBaseDao) Gets(ctx context.Context, id any) (entitys []*entity.CrontabBase, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *CrontabBaseDao) Find(ctx context.Context, in *do.CrontabBaseListInput) (out []*entity.CrontabBase, err error) {
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
func (dao *CrontabBaseDao) FindOne(ctx context.Context, in *do.CrontabBaseListInput) (one *entity.CrontabBase, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.CrontabBase
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
func (dao *CrontabBaseDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.CrontabBaseListInput) (out gdb.Result, err error) {
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
func (dao *CrontabBaseDao) FindKey(ctx context.Context, in *do.CrontabBaseListInput) (out []interface{}, err error) {
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
func (dao *CrontabBaseDao) List(ctx context.Context, in *do.CrontabBaseListInput) (out *do.CrontabBaseListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.CrontabBaseListOutput{}
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
func (dao *CrontabBaseDao) Add(ctx context.Context, in *do.CrontabBase) (lastInsertId int64, err error) {
	data := do.CrontabBase{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *CrontabBaseDao) Edit(ctx context.Context, id any, in *do.CrontabBase) (int64, error) {
	data := do.CrontabBase{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *CrontabBaseDao) EditWhere(ctx context.Context, where *do.CrontabBaseListInput, in *do.CrontabBase) (int64, error) {
	data := do.CrontabBase{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *CrontabBaseDao) Save(ctx context.Context, in *do.CrontabBase) (affected int64, err error) {
	data := do.CrontabBase{}
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
func (dao *CrontabBaseDao) Saves(ctx context.Context, in []*do.CrontabBase) (affected int64, err error) {
	data := []do.CrontabBase{}
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
func (dao *CrontabBaseDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *CrontabBaseDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *CrontabBaseDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *CrontabBaseDao) RemoveWhere(ctx context.Context, where *do.CrontabBaseListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *CrontabBaseDao) Count(ctx context.Context, in *do.CrontabBaseListInput) (count int, err error) {
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
