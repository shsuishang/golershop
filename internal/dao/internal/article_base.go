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

// ArticleBaseDao is the data access object for table cms_article_base.
type ArticleBaseDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ArticleBaseColumns // columns contains all the column names of Table for convenient usage.
}

// ArticleBaseColumns defines and stores column names for table cms_article_base.
type ArticleBaseColumns struct {
	ArticleId             string // 文章编号
	ArticleTitle          string // 文章标题
	ArticleName           string // 文章别名
	ArticleExcerpt        string // 文章摘要
	ArticleContent        string // 文章内容(HTML)
	ArticleUrl            string // 调用网址:默认为本页面构造的网址
	CategoryId            string // 所属分类
	ArticleTemplate       string // 文章模板
	ArticleSeoTitle       string // SEO标题
	ArticleSeoKeywords    string // SEO关键字
	ArticleSeoDescription string // SEO描述
	ArticleReplyFlag      string // 启用问答(BOOL):0-否;1-是
	ArticleLang           string // 文章语言
	ArticleType           string // 文章类型(ENUM):1-文章;2-公告
	ArticleSort           string // 文章排序
	ArticleStatus         string // 文章状态(BOOL):0-关闭;1-启用
	ArticleAddTime        string // 添加时间
	ArticleImage          string // 文章图片
	UserId                string // 文章作者
	ArticleTags           string // 文章标签(DOT):文章标签
	ArticleIsPopular      string // 是否热门(BOOL):0-否;1-是
	PrimaryKey            string // 主键
}

// articleBaseColumns holds the columns for table cms_article_base.
var articleBaseColumns = ArticleBaseColumns{
	ArticleId:             "article_id",
	ArticleTitle:          "article_title",
	ArticleName:           "article_name",
	ArticleExcerpt:        "article_excerpt",
	ArticleContent:        "article_content",
	ArticleUrl:            "article_url",
	CategoryId:            "category_id",
	ArticleTemplate:       "article_template",
	ArticleSeoTitle:       "article_seo_title",
	ArticleSeoKeywords:    "article_seo_keywords",
	ArticleSeoDescription: "article_seo_description",
	ArticleReplyFlag:      "article_reply_flag",
	ArticleLang:           "article_lang",
	ArticleType:           "article_type",
	ArticleSort:           "article_sort",
	ArticleStatus:         "article_status",
	ArticleAddTime:        "article_add_time",
	ArticleImage:          "article_image",
	UserId:                "user_id",
	ArticleTags:           "article_tags",
	ArticleIsPopular:      "article_is_popular",
	PrimaryKey:            "article_id",
}

// NewArticleBaseDao creates and returns a new DAO object for table data access.
func NewArticleBaseDao() *ArticleBaseDao {
	return &ArticleBaseDao{
		group:   "cms",
		table:   "cms_article_base",
		columns: articleBaseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ArticleBaseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ArticleBaseDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ArticleBaseDao) Columns() ArticleBaseColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ArticleBaseDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ArticleBaseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ArticleBaseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *ArticleBaseDao) Get(ctx context.Context, id any) (one *entity.ArticleBase, err error) {
	var entitys []*entity.ArticleBase
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
func (dao *ArticleBaseDao) Gets(ctx context.Context, id any) (entitys []*entity.ArticleBase, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *ArticleBaseDao) Find(ctx context.Context, in *do.ArticleBaseListInput) (out []*entity.ArticleBase, err error) {
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
func (dao *ArticleBaseDao) FindOne(ctx context.Context, in *do.ArticleBaseListInput) (one *entity.ArticleBase, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.ArticleBase
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
func (dao *ArticleBaseDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.ArticleBaseListInput) (out gdb.Result, err error) {
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
func (dao *ArticleBaseDao) FindKey(ctx context.Context, in *do.ArticleBaseListInput) (out []interface{}, err error) {
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
func (dao *ArticleBaseDao) List(ctx context.Context, in *do.ArticleBaseListInput) (out *do.ArticleBaseListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.ArticleBaseListOutput{}
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
func (dao *ArticleBaseDao) Add(ctx context.Context, in *do.ArticleBase) (lastInsertId int64, err error) {
	data := do.ArticleBase{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *ArticleBaseDao) Edit(ctx context.Context, id any, in *do.ArticleBase) (int64, error) {
	data := do.ArticleBase{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *ArticleBaseDao) EditWhere(ctx context.Context, where *do.ArticleBaseListInput, in *do.ArticleBase) (int64, error) {
	data := do.ArticleBase{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *ArticleBaseDao) Save(ctx context.Context, in *do.ArticleBase) (affected int64, err error) {
	data := do.ArticleBase{}
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
func (dao *ArticleBaseDao) Saves(ctx context.Context, in []*do.ArticleBase) (affected int64, err error) {
	data := []do.ArticleBase{}
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
func (dao *ArticleBaseDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *ArticleBaseDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *ArticleBaseDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *ArticleBaseDao) RemoveWhere(ctx context.Context, where *do.ArticleBaseListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *ArticleBaseDao) Count(ctx context.Context, in *do.ArticleBaseListInput) (count int, err error) {
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
