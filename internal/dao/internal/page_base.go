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

// PageBaseDao is the data access object for table sys_page_base.
type PageBaseDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns PageBaseColumns // columns contains all the column names of Table for convenient usage.
}

// PageBaseColumns defines and stores column names for table sys_page_base.
type PageBaseColumns struct {
	PageId         string // 页面编号
	PageName       string // 页面名称
	StoreId        string // 所属店铺
	UserId         string // 所属用户
	SubsiteId      string // 所属分站:0-总站
	PageBuildin    string // 是否内置(BOOL):0-否;1-是
	PageType       string // 类型(ENUM):1-WAP;2-PC;3-APP
	PageTpl        string // 页面布局模板
	AppId          string // 所属APP
	PageCode       string // 页面代码
	PageNav        string // 导航数据
	PageConfig     string // 页面配置
	PageShareTitle string // 分享标题
	PageShareImage string // 分享图片
	PageQrcode     string // 分享二维码
	PageIndex      string // 是否首页(BOOL):0-非首页;1-首页
	PageGb         string // 拼团首页(BOOL):0-非首页;1-首页
	PageActivity   string // 活动首页(BOOL):0-非首页;1-首页
	PagePoint      string // 积分首页(BOOL):0-非首页;1-首页
	PageGbs        string // 团购首页(BOOL):0-非首页;1-首页
	PagePackage    string // 组合套餐(BOOL):0-非首页;1-首页
	PagePfgb       string // 批发团购首页(BOOL):0-非首页;1-首页
	PageSns        string // 社区(BOOL):0-非首页;1-首页
	PageArticle    string // 资讯(BOOL):0-非首页;1-首页
	PageZerobuy    string // 零元购区(BOOL):0-否;1-是
	PageHigharea   string // 高额返区(BOOL):0-否;1-是
	PageTaday      string // 今日爆款(BOOL):0-否;1-是
	PageEveryday   string // 每日好店(BOOL):0-否;1-是
	PageSecondkill string // 整点秒杀(BOOL):0-否;1-是
	PageSecondday  string // 天天秒淘(BOOL):0-否;1-是
	PageRura       string // 设置土特产(BOOL):0-否;1-是
	PageLikeyou    string // 用户页banner(BOOL):0-否;1-是
	PageExchange   string // 兑换专区(BOOL):0-否;1-是
	PageNew        string // 新品首发(BOOL):0-否;1-是
	PageNewperson  string // 新人优惠(BOOL):0-否;1-是
	PageUpgrade    string // 升级VIP(BOOL):0-否;1-是
	PageMessage    string // 信息发布(BOOL):0-否;1-是
	PageRelease    string // 是否发布(BOOL):0-否;1-是
	PrimaryKey     string // 主键
}

// pageBaseColumns holds the columns for table sys_page_base.
var pageBaseColumns = PageBaseColumns{
	PageId:         "page_id",
	PageName:       "page_name",
	StoreId:        "store_id",
	UserId:         "user_id",
	SubsiteId:      "subsite_id",
	PageBuildin:    "page_buildin",
	PageType:       "page_type",
	PageTpl:        "page_tpl",
	AppId:          "app_id",
	PageCode:       "page_code",
	PageNav:        "page_nav",
	PageConfig:     "page_config",
	PageShareTitle: "page_share_title",
	PageShareImage: "page_share_image",
	PageQrcode:     "page_qrcode",
	PageIndex:      "page_index",
	PageGb:         "page_gb",
	PageActivity:   "page_activity",
	PagePoint:      "page_point",
	PageGbs:        "page_gbs",
	PagePackage:    "page_package",
	PagePfgb:       "page_pfgb",
	PageSns:        "page_sns",
	PageArticle:    "page_article",
	PageZerobuy:    "page_zerobuy",
	PageHigharea:   "page_higharea",
	PageTaday:      "page_taday",
	PageEveryday:   "page_everyday",
	PageSecondkill: "page_secondkill",
	PageSecondday:  "page_secondday",
	PageRura:       "page_rura",
	PageLikeyou:    "page_likeyou",
	PageExchange:   "page_exchange",
	PageNew:        "page_new",
	PageNewperson:  "page_newperson",
	PageUpgrade:    "page_upgrade",
	PageMessage:    "page_message",
	PageRelease:    "page_release",
	PrimaryKey:     "page_id",
}

// NewPageBaseDao creates and returns a new DAO object for table data access.
func NewPageBaseDao() *PageBaseDao {
	return &PageBaseDao{
		group:   "sys",
		table:   "sys_page_base",
		columns: pageBaseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PageBaseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PageBaseDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PageBaseDao) Columns() PageBaseColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PageBaseDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PageBaseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PageBaseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *PageBaseDao) Get(ctx context.Context, id any) (one *entity.PageBase, err error) {
	var entitys []*entity.PageBase
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
func (dao *PageBaseDao) Gets(ctx context.Context, id any) (entitys []*entity.PageBase, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *PageBaseDao) Find(ctx context.Context, in *do.PageBaseListInput) (out []*entity.PageBase, err error) {
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
func (dao *PageBaseDao) FindOne(ctx context.Context, in *do.PageBaseListInput) (one *entity.PageBase, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.PageBase
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
func (dao *PageBaseDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.PageBaseListInput) (out gdb.Result, err error) {
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
func (dao *PageBaseDao) FindKey(ctx context.Context, in *do.PageBaseListInput) (out []interface{}, err error) {
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
func (dao *PageBaseDao) List(ctx context.Context, in *do.PageBaseListInput) (out *do.PageBaseListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.PageBaseListOutput{}
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
func (dao *PageBaseDao) Add(ctx context.Context, in *do.PageBase) (lastInsertId int64, err error) {
	data := do.PageBase{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *PageBaseDao) Edit(ctx context.Context, id any, in *do.PageBase) (int64, error) {
	data := do.PageBase{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *PageBaseDao) EditWhere(ctx context.Context, where *do.PageBaseListInput, in *do.PageBase) (int64, error) {
	data := do.PageBase{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *PageBaseDao) Save(ctx context.Context, in *do.PageBase) (affected int64, err error) {
	data := do.PageBase{}
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
func (dao *PageBaseDao) Saves(ctx context.Context, in []*do.PageBase) (affected int64, err error) {
	data := []do.PageBase{}
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
func (dao *PageBaseDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *PageBaseDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *PageBaseDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *PageBaseDao) RemoveWhere(ctx context.Context, where *do.PageBaseListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *PageBaseDao) Count(ctx context.Context, in *do.PageBaseListInput) (count int, err error) {
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
