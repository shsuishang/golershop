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

// FeedbackBaseDao is the data access object for table sys_feedback_base.
type FeedbackBaseDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns FeedbackBaseColumns // columns contains all the column names of Table for convenient usage.
}

// FeedbackBaseColumns defines and stores column names for table sys_feedback_base.
type FeedbackBaseColumns struct {
	FeedbackId                 string // 反馈编号
	FeedbackCategoryId         string // 分类编号
	UserId                     string // 用户编号
	UserNickname               string // 用户昵称
	FeedbackQuestion           string // 反馈问题:在这里描述您遇到的问题
	FeedbackQuestionUrl        string // 页面链接
	FeedbackQuestionAnswer     string // 反馈问题
	FeedbackQuestionTime       string // 反馈时间
	FeedbackQuestionAnswerTime string // 回复时间
	FeedbackQuestionStatus     string // 举报状态(BOOL):0-未处理;1-已处理
	FeedbackQuestionResult     string // 举报处理结果(ENUM):1-无效举报;2-恶意举报;3-有效举报
	ItemId                     string // 产品编号
	AdminId                    string // 回复人员
	PrimaryKey                 string // 主键
}

// feedbackBaseColumns holds the columns for table sys_feedback_base.
var feedbackBaseColumns = FeedbackBaseColumns{
	FeedbackId:                 "feedback_id",
	FeedbackCategoryId:         "feedback_category_id",
	UserId:                     "user_id",
	UserNickname:               "user_nickname",
	FeedbackQuestion:           "feedback_question",
	FeedbackQuestionUrl:        "feedback_question_url",
	FeedbackQuestionAnswer:     "feedback_question_answer",
	FeedbackQuestionTime:       "feedback_question_time",
	FeedbackQuestionAnswerTime: "feedback_question_answer_time",
	FeedbackQuestionStatus:     "feedback_question_status",
	FeedbackQuestionResult:     "feedback_question_result",
	ItemId:                     "item_id",
	AdminId:                    "admin_id",
	PrimaryKey:                 "feedback_id",
}

// NewFeedbackBaseDao creates and returns a new DAO object for table data access.
func NewFeedbackBaseDao() *FeedbackBaseDao {
	return &FeedbackBaseDao{
		group:   "sys",
		table:   "sys_feedback_base",
		columns: feedbackBaseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FeedbackBaseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FeedbackBaseDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FeedbackBaseDao) Columns() FeedbackBaseColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FeedbackBaseDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FeedbackBaseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FeedbackBaseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *FeedbackBaseDao) Get(ctx context.Context, id any) (one *entity.FeedbackBase, err error) {
	var entitys []*entity.FeedbackBase
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
func (dao *FeedbackBaseDao) Gets(ctx context.Context, id any) (entitys []*entity.FeedbackBase, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *FeedbackBaseDao) Find(ctx context.Context, in *do.FeedbackBaseListInput) (out []*entity.FeedbackBase, err error) {
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
func (dao *FeedbackBaseDao) FindOne(ctx context.Context, in *do.FeedbackBaseListInput) (one *entity.FeedbackBase, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.FeedbackBase
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
func (dao *FeedbackBaseDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.FeedbackBaseListInput) (out gdb.Result, err error) {
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
func (dao *FeedbackBaseDao) FindKey(ctx context.Context, in *do.FeedbackBaseListInput) (out []interface{}, err error) {
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
func (dao *FeedbackBaseDao) List(ctx context.Context, in *do.FeedbackBaseListInput) (out *do.FeedbackBaseListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.FeedbackBaseListOutput{}
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
func (dao *FeedbackBaseDao) Add(ctx context.Context, in *do.FeedbackBase) (lastInsertId int64, err error) {
	data := do.FeedbackBase{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *FeedbackBaseDao) Edit(ctx context.Context, id any, in *do.FeedbackBase) (int64, error) {
	data := do.FeedbackBase{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *FeedbackBaseDao) EditWhere(ctx context.Context, where *do.FeedbackBaseListInput, in *do.FeedbackBase) (int64, error) {
	data := do.FeedbackBase{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *FeedbackBaseDao) Save(ctx context.Context, in *do.FeedbackBase) (affected int64, err error) {
	data := do.FeedbackBase{}
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
func (dao *FeedbackBaseDao) Saves(ctx context.Context, in []*do.FeedbackBase) (affected int64, err error) {
	data := []do.FeedbackBase{}
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
func (dao *FeedbackBaseDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *FeedbackBaseDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *FeedbackBaseDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *FeedbackBaseDao) RemoveWhere(ctx context.Context, where *do.FeedbackBaseListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *FeedbackBaseDao) Count(ctx context.Context, in *do.FeedbackBaseListInput) (count int, err error) {
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
