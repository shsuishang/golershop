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

// UserInfoDao is the data access object for table account_user_info.
type UserInfoDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns UserInfoColumns // columns contains all the column names of Table for convenient usage.
}

// UserInfoColumns defines and stores column names for table account_user_info.
type UserInfoColumns struct {
	UserId               string // 用户编号
	UserAccount          string // 用户账号
	UserNickname         string // 用户昵称
	UserAvatar           string // 用户头像
	UserState            string // 状态(ENUM):0-锁定;1-已激活;2-未激活;
	UserMobile           string // 手机号码(mobile)
	UserIntl             string // 国家编码
	UserGender           string // 性别(ENUM):0-保密;1-男;  2-女;
	UserBirthday         string // 生日(DATE)
	UserEmail            string // 用户邮箱(email)
	UserLevelId          string // 等级编号
	UserRealname         string // 真实姓名
	UserIdcard           string // 身份证
	UserIdcardImages     string // 身份证图片(DTO)
	UserIsAuthentication string // 认证状态(ENUM):0-未认证;1-待审核;2-认证通过;3-认证失败
	TagIds               string // 用户标签(DOT)
	UserFrom             string // 用户来源(ENUM):2310-其它;2311-pc;2312-H5;2313-APP;2314-小程序;2315-公众号
	UserNew              string // 新人标识(BOOL):0-不是;1-是
	PrimaryKey           string // 主键
}

// userInfoColumns holds the columns for table account_user_info.
var userInfoColumns = UserInfoColumns{
	UserId:               "user_id",
	UserAccount:          "user_account",
	UserNickname:         "user_nickname",
	UserAvatar:           "user_avatar",
	UserState:            "user_state",
	UserMobile:           "user_mobile",
	UserIntl:             "user_intl",
	UserGender:           "user_gender",
	UserBirthday:         "user_birthday",
	UserEmail:            "user_email",
	UserLevelId:          "user_level_id",
	UserRealname:         "user_realname",
	UserIdcard:           "user_idcard",
	UserIdcardImages:     "user_idcard_images",
	UserIsAuthentication: "user_is_authentication",
	TagIds:               "tag_ids",
	UserFrom:             "user_from",
	UserNew:              "user_new",
	PrimaryKey:           "user_id",
}

// NewUserInfoDao creates and returns a new DAO object for table data access.
func NewUserInfoDao() *UserInfoDao {
	return &UserInfoDao{
		group:   "account",
		table:   "account_user_info",
		columns: userInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserInfoDao) Columns() UserInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Get 读取一条记录
func (dao *UserInfoDao) Get(ctx context.Context, id any) (one *entity.UserInfo, err error) {
	var entitys []*entity.UserInfo
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
func (dao *UserInfoDao) Gets(ctx context.Context, id any) (entitys []*entity.UserInfo, err error) {
	if !g.IsEmpty(id) {
		err = dao.Ctx(ctx).WherePri(id).Scan(&entitys)
	}

	return entitys, err
}

// Find 查询数据
func (dao *UserInfoDao) Find(ctx context.Context, in *do.UserInfoListInput) (out []*entity.UserInfo, err error) {
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
func (dao *UserInfoDao) FindOne(ctx context.Context, in *do.UserInfoListInput) (one *entity.UserInfo, err error) {
	in.BaseList.Size = 1

	var entitys []*entity.UserInfo
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
func (dao *UserInfoDao) FindFields(ctx context.Context, fieldNamesOrMapStruct interface{}, in *do.UserInfoListInput) (out gdb.Result, err error) {
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
func (dao *UserInfoDao) FindKey(ctx context.Context, in *do.UserInfoListInput) (out []interface{}, err error) {
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
func (dao *UserInfoDao) List(ctx context.Context, in *do.UserInfoListInput) (out *do.UserInfoListOutput, err error) {
	var (
		m = dao.Ctx(ctx)
	)

	query := m.Where(in.Where).OmitNil()
	query = ml.BuildWhere(query, in.WhereLike, in.WhereExt)

	out = &do.UserInfoListOutput{}
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
func (dao *UserInfoDao) Add(ctx context.Context, in *do.UserInfo) (lastInsertId int64, err error) {
	data := do.UserInfo{}
	if err = gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	return dao.Ctx(ctx).Data(data).OmitNil().InsertAndGetId()
}

// Edit 编辑
func (dao *UserInfoDao) Edit(ctx context.Context, id any, in *do.UserInfo) (int64, error) {
	data := do.UserInfo{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	//FieldsEx(dao.Columns().Id)
	return dao.Ctx(ctx).Data(data).OmitNil().WherePri(id).UpdateAndGetAffected()
}

// EditWhere 根据Where条件编辑
func (dao *UserInfoDao) EditWhere(ctx context.Context, where *do.UserInfoListInput, in *do.UserInfo) (int64, error) {
	data := do.UserInfo{}
	if err := gconv.Scan(in, &data); err != nil {
		return 0, err
	}

	query := dao.Ctx(ctx).Data(data).OmitNil().Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	return query.UpdateAndGetAffected()
}

// Save 保存
func (dao *UserInfoDao) Save(ctx context.Context, in *do.UserInfo) (affected int64, err error) {
	data := do.UserInfo{}
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
func (dao *UserInfoDao) Saves(ctx context.Context, in []*do.UserInfo) (affected int64, err error) {
	data := []do.UserInfo{}
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
func (dao *UserInfoDao) Increment(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Increment(column, amount)
}

// Decrement 减少
func (dao *UserInfoDao) Decrement(ctx context.Context, id any, column string, amount interface{}) (sql.Result, error) {
	return dao.Ctx(ctx).WherePri(id).Decrement(column, amount)
}

// Remove 根据主键删除
func (dao *UserInfoDao) Remove(ctx context.Context, id any) (int64, error) {
	res, err := dao.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Remove 根据Where条件删除
func (dao *UserInfoDao) RemoveWhere(ctx context.Context, where *do.UserInfoListInput) (int64, error) {
	query := dao.Ctx(ctx).Where(where.Where)
	query = ml.BuildWhere(query, where.WhereLike, where.WhereExt)

	res, err := query.Delete()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Count 查询数据记录
func (dao *UserInfoDao) Count(ctx context.Context, in *do.UserInfoListInput) (count int, err error) {
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

// GetUserInfoMap 获取用户信息映射
func (dao *UserInfoDao) GetUserInfoMap(ctx context.Context, userIds []uint) (map[int]*entity.UserInfo, error) {
	// 获取用户信息列表
	userInfos, err := dao.Gets(ctx, userIds)
	if err != nil {
		return nil, err
	}

	// 初始化用户信息映射
	userInfoMap := make(map[int]*entity.UserInfo)

	if len(userInfos) > 0 {
		// 将用户信息列表转换为映射
		for _, userInfo := range userInfos {
			userInfoMap[int(userInfo.UserId)] = userInfo
		}
	}

	return userInfoMap, nil
}
