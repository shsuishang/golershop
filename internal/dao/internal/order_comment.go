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
	"golershop.cn/internal/dao/global"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
	"golershop.cn/utility/log"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// OrderCommentDao is the data access object for table trade_order_comment.
type OrderCommentDao struct {
	BaseRepository
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns OrderCommentColumns // columns contains all the column names of Table for convenient usage.
}

// OrderCommentColumns defines and stores column names for table trade_order_comment.
type OrderCommentColumns struct {
	OrderId                    string // 订单编号
	StoreId                    string // 卖家店铺编号-冗余
	StoreName                  string // 店铺名称
	UserId                     string // 买家编号
	UserName                   string // 买家姓名
	CommentPoints              string // 获得积分-冗余，独立表记录
	CommentScores              string // 评价星级1-5积分
	CommentContent             string // 评价内容
	CommentImage               string // 评论上传的图片：|分割多张图片
	CommentHelpful             string // 有帮助
	CommentNohelpful           string // 无帮助
	CommentTime                string // 评价时间
	CommentIsAnonymous         string // 匿名评价
	CommentEnable              string // 评价信息的状态(BOOL): 1-正常显示; 0-禁止显示
	CommentStoreDescCredit     string // 描述相符评分 - order_buyer_evaluation_status , 评价状态改变后不需要再次评论，根据订单走
	CommentStoreServiceCredit  string // 服务态度评分 - order_buyer_evaluation_status
	CommentStoreDeliveryCredit string // 发货速度评分 - order_buyer_evaluation_status
	SubsiteId                  string // 所属分站:0-总站
	PrimaryKey                 string // 主键
}

// orderCommentColumns holds the columns for table trade_order_comment.
var orderCommentColumns = OrderCommentColumns{
	OrderId:                    "order_id",
	StoreId:                    "store_id",
	StoreName:                  "store_name",
	UserId:                     "user_id",
	UserName:                   "user_name",
	CommentPoints:              "comment_points",
	CommentScores:              "comment_scores",
	CommentContent:             "comment_content",
	CommentImage:               "comment_image",
	CommentHelpful:             "comment_helpful",
	CommentNohelpful:           "comment_nohelpful",
	CommentTime:                "comment_time",
	CommentIsAnonymous:         "comment_is_anonymous",
	CommentEnable:              "comment_enable",
	CommentStoreDescCredit:     "comment_store_desc_credit",
	CommentStoreServiceCredit:  "comment_store_service_credit",
	CommentStoreDeliveryCredit: "comment_store_delivery_credit",
	SubsiteId:                  "subsite_id",
	PrimaryKey:                 "order_id",
}

// NewOrderCommentDao creates and returns a new DAO object for table data access.
func NewOrderCommentDao() *OrderCommentDao {
	return &OrderCommentDao{
		BaseRepository: BaseRepository{
			Table: "trade_order_comment",
			Group: "trade",
		},
		group:   "trade",
		table:   "trade_order_comment",
		columns: orderCommentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderCommentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderCommentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrderCommentDao) Columns() OrderCommentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderCommentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrderCommentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrderCommentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	service.BizCtx().IncrementTx(ctx)
	err = dao.Ctx(ctx).Transaction(ctx, f)
	service.BizCtx().DecrementTx(ctx)
	return
}

// Get 读取一条记录
func (dao *OrderCommentDao) Get(ctx context.Context, id any) (one *entity.OrderComment, err error) {
	var entitys []*entity.OrderComment
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
func (dao *OrderCommentDao) Gets(ctx context.Context, id any) (entitys []*entity.OrderComment, err error) {
	if !g.IsEmpty(id) {
		items := map[string]gdb.Record{}
		if global.Cache && dao.IsNotInTransaction(ctx) {
			keys := dao.Keys(id)
			rowMap, err := g.Redis().MGet(ctx, keys...)
			if err != nil {
				log.Error(ctx, err)
			}

			var existIds []interface{}
			for _, v := range rowMap {
				if !g.IsEmpty(v) {
					item := gdb.Record{}
					gconv.Struct(v, &item)
					items[item[dao.Columns().PrimaryKey].String()] = item
					existIds = append(existIds, item[dao.columns.PrimaryKey])
				}
			}

			var newIds []any
			if len(items) != len(keys) {
				for _, i := range gconv.SliceAny(id) {
					if !array.InArray(existIds, i) {
						newIds = append(newIds, i)
					}
				}

				if len(newIds) != 0 {
					newRes, err := dao.Ctx(ctx).WherePri(newIds).All()

					if err != nil {
						return nil, err
					}

					for _, record := range newRes {
						if !record[dao.Columns().PrimaryKey].IsEmpty() {
							_, err := g.Redis().Set(ctx, dao.Key(record[dao.Columns().PrimaryKey]), record)
							if err != nil {
								log.Error(ctx, err)
							}

							items[record[dao.Columns().PrimaryKey].String()] = record
						}
					}
				}
			}
		} else {
			res, err := dao.Ctx(ctx).WherePri(id).All()

			if err != nil {
				return nil, err
			}

			for _, record := range res {
				if !record[dao.Columns().PrimaryKey].IsEmpty() {
					items[record[dao.Columns().PrimaryKey].String()] = record
				}
			}
		}

		// 排序
		if len(items) > 0 {
			for _, i := range gconv.SliceStr(id) {
				item := &entity.OrderComment{}
				gconv.Struct(items[i], item)

				entitys = append(entitys, item)
			}
		}
	}

	return entitys, err
}

// Find 查询数据
func (dao *OrderCommentDao) Find(ctx context.Context, in *do.OrderCommentListInput) (out []*entity.OrderComment, err error) {
	keys, err := dao.FindKey(ctx, in)
	if err != nil {
		return nil, err
	}

	if len(keys) > 0 {
		out, err = dao.Gets(ctx, keys)
	}

	return out, err
}

// FindOne 查询一条数据
func (dao *OrderCommentDao) FindOne(ctx context.Context, in *do.OrderCommentListInput) (one *entity.OrderComment, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.OrderComment
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
func (dao *OrderCommentDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.OrderCommentListInput) (out gdb.Result, err error) {
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
func (dao *OrderCommentDao) FindKey(ctx context.Context, in *do.OrderCommentListInput) (out []interface{}, err error) {
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
func (dao *OrderCommentDao) List(ctx context.Context, in *do.OrderCommentListInput) (out *do.OrderCommentListOutput, err error) {
	res, err := dao.ListKey(ctx, in)

	if err != nil {
		return nil, err
	}

	out = &do.OrderCommentListOutput{}
	out.Page = res.Page
	out.Total = res.Total
	out.Records = res.Records
	out.Size = res.Size

	if len(res.Items) > 0 {
		rows, err := dao.Gets(ctx, res.Items)
		if err != nil {
			return nil, err
		}

		out.Items = rows
	}

	return out, nil
}

// List 分页读取
func (dao *OrderCommentDao) ListKey(ctx context.Context, in *do.OrderCommentListInput) (out *do.OrderCommentListKeyOutput, err error) {
	var (
		m = dao.Ctx(ctx).Fields(dao.Columns().PrimaryKey)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.OrderCommentListKeyOutput{}
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
	idRes, err := query.All()
	if err != nil {
		return nil, err
	}

	for _, record := range idRes {
		if !record[dao.Columns().PrimaryKey].IsEmpty() {
			out.Items = append(out.Items, record[dao.Columns().PrimaryKey])
		}
	}

	return out, nil
}

// Add 新增
func (dao *OrderCommentDao) Add(ctx context.Context, in *do.OrderComment) (lastInsertId int64, err error) {
	data := do.OrderComment{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *OrderCommentDao) Edit(ctx context.Context, id any, in *do.OrderComment) (int64, error) {
	data := do.OrderComment{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	num, err := dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
	if err != nil {
		return 0, err
	}

	if num > 0 {
		dao.RemoveCache(ctx, id)
	}

	return num, err
}

// EditWhere 根据Where条件编辑
func (dao *OrderCommentDao) EditWhere(ctx context.Context, where *do.OrderCommentListInput, in *do.OrderComment) (int64, error) {
	ids, err := dao.FindKey(ctx, where)
	if err != nil {
		return 0, err
	}

	if len(ids) > 0 {
		num, err := dao.Edit(ctx, ids, in)

		if err != nil {
			return 0, err
		}

		return num, err
	}

	return 0, nil
}

// Save 保存
func (dao *OrderCommentDao) Save(ctx context.Context, in *do.OrderComment) (affected int64, err error) {
	input := make([]*do.OrderComment, 0)
	input = append(input, in)

	return dao.Saves(ctx, input)
}

// Saves 批量保存
func (dao *OrderCommentDao) Saves(ctx context.Context, in []*do.OrderComment) (affected int64, err error) {
	data := []do.OrderComment{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	res, err := dao.Ctx(ctx).Data(data).OmitNil().Save()

	if err != nil {
		return 0, err
	}

	ids := array.Column(in, dao.columns.PrimaryKey)
	ids = array.DeleteEmpty(ids)
	dao.RemoveCache(ctx, ids)

	return res.RowsAffected()
}

// Increment 增加
func (dao *OrderCommentDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Increment(column, amount)
	if err == nil {
		dao.RemoveCache(ctx, id)
	}

	return res, err
}

// Decrement 减少
func (dao *OrderCommentDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
	if err == nil {
		dao.RemoveCache(ctx, id)
	}

	return res, err
}

// Remove 根据主键删除
func (dao *OrderCommentDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	dao.RemoveCache(ctx, id)

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *OrderCommentDao) RemoveWhere(ctx context.Context, where *do.OrderCommentListInput) (int64, error) {
	ids, err := dao.FindKey(ctx, where)
	if err != nil {
		return 0, err
	}

	if len(ids) > 0 {
		num, err := dao.Remove(ctx, ids)

		if err != nil {
			return 0, err
		}

		return num, err
	}

	return 0, nil
}

// Count 查询数据记录
func (dao *OrderCommentDao) Count(ctx context.Context, in *do.OrderCommentListInput) (count int, err error) {
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
