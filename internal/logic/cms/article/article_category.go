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

package cms

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
	"time"
)

type sArticleCategory struct{}

var (
	cachePreKey   = "menu_base"
	cacheDuration = time.Hour
	redisAdapter  = gcache.NewAdapterRedis(g.Redis())
	cache         = gcache.NewWithAdapter(redisAdapter)
)

func init() {
	service.RegisterArticleCategory(New())
}

func New() *sArticleCategory {
	return &sArticleCategory{}
}

// 读取商品分类
func (s *sArticleCategory) Get(ctx context.Context, id any) (out *entity.ArticleCategory, err error) {
	var list []*entity.ArticleCategory
	list, err = s.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// 读取多条记录模式
func (s *sArticleCategory) Gets(ctx context.Context, id any) (list []*entity.ArticleCategory, err error) {

	err = dao.ArticleCategory.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// 分页读取
func (s *sArticleCategory) List(ctx context.Context, in *do.ArticleCategoryListInput) (out *do.ArticleCategoryListOutput, err error) {
	list, err := dao.ArticleCategory.List(ctx, in)

	gconv.Scan(list, &out)

	return out, nil
}

// 业务封装

// 查询数据
func (s *sArticleCategory) GetTree(ctx context.Context, in *do.ArticleCategoryListInput) (out []*model.ArticleCategoryTreeNode, err error) {
	//in.Where.CategoryIsEnable = true
	in.Sidx = dao.ArticleCategory.Columns().CategoryOrder
	in.Sort = "ASC"
	//in.Size = consts.MAX_LIST_NUM

	res, err := dao.ArticleCategory.List(ctx, in)

	// 数据转换
	var list []*entity.ArticleCategory
	gconv.Scan(res.Items, &list)

	// 数据处理
	var categoryNode model.ArticleCategoryTreeNode
	s.makeTree(list, &categoryNode)

	columnIds := array.Column(list, dao.ArticleCategory.Columns().CategoryId)

	for _, c := range list {
		if c.CategoryParentId != 0 && !array.InArray(columnIds, c.CategoryParentId) {
			child := &model.ArticleCategoryTreeNode{}
			//child.ArticleCategory = *c
			gconv.Scan(*c, &child.ArticleCategory)

			categoryNode.Children = append(categoryNode.Children, child)
		}
	}

	return categoryNode.Children, nil
}

// 递归生成分类列表
func (s *sArticleCategory) makeTree(list []*entity.ArticleCategory, tn *model.ArticleCategoryTreeNode) {
	for _, c := range list {
		if c.CategoryParentId == tn.CategoryId {
			child := &model.ArticleCategoryTreeNode{}
			//child.ArticleCategory = *c
			gconv.Scan(*c, &child.ArticleCategory)

			tn.Children = append(tn.Children, child)
			s.makeTree(list, child)
		}
	}
}

// 新增
func (s *sArticleCategory) Add(ctx context.Context, in *do.ArticleCategory) (out int64, err error) {
	// 不允许HTML代码
	//if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
	//	return out, err
	//}

	lastInsertId, err := dao.ArticleCategory.Add(ctx, in)
	if err != nil {
		return out, err
	}
	return lastInsertId, err
}

// 编辑
func (s *sArticleCategory) Edit(ctx context.Context, in *do.ArticleCategory) (affected int64, err error) {
	// 不允许HTML代码
	//if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
	//	return out, err
	//}

	affected, err = dao.ArticleCategory.Edit(ctx, in.CategoryId, in)

	return affected, err
}

// 删除多条记录模式
func (s *sArticleCategory) Remove(ctx context.Context, id any) (affected int64, err error) {

	//是否子项
	categoryNum, err := dao.ArticleCategory.Ctx(ctx).Count(do.ArticleCategory{CategoryParentId: id})
	if err != nil {
		return 0, err
	}
	if categoryNum > 0 {
		return 0, errors.New(fmt.Sprintf("该分类下有子分类"))
	}

	//是否被使用
	count, err := dao.ArticleBase.Ctx(ctx).Count(do.ArticleBase{CategoryId: id})

	if err != nil {
		return 0, err
	}

	if count > 0 {
		return 0, errors.New(fmt.Sprintf("有 %d 章文章使用，不可删除", count))
	}

	//系统内置
	one, err := dao.ArticleCategory.Get(ctx, id)

	if one.CategoryBuildin {
		return 0, errors.New("系统内置，不可删除")
	}

	affected, err = dao.ArticleCategory.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}
