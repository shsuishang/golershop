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

package service

import (
	"context"

	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
)

type (
	IArticleBase interface {
		Find(ctx context.Context, in *do.ArticleBaseListInput) (out []*entity.ArticleBase, err error)
		List(ctx context.Context, in *do.ArticleBaseListInput) (out *model.ArticleBaseOutput, err error)
		Add(ctx context.Context, in *do.ArticleBase) (lastInsertId int64, err error)
		Edit(ctx context.Context, in *do.ArticleBase) (affected int64, err error)
		Remove(ctx context.Context, id any) (affected int64, err error)
		RemoveBatch(ctx context.Context, id string) (affected int64, err error)
	}
	IArticleCategory interface {
		Get(ctx context.Context, id any) (out *entity.ArticleCategory, err error)
		Gets(ctx context.Context, id any) (list []*entity.ArticleCategory, err error)
		List(ctx context.Context, in *do.ArticleCategoryListInput) (out *do.ArticleCategoryListOutput, err error)
		GetTree(ctx context.Context, in *do.ArticleCategoryListInput) (out []*model.ArticleCategoryTreeNode, err error)
		Add(ctx context.Context, in *do.ArticleCategory) (out int64, err error)
		Edit(ctx context.Context, in *do.ArticleCategory) (affected int64, err error)
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IArticleTag interface {
		Find(ctx context.Context, in *do.ArticleTagListInput) (out []*entity.ArticleTag, err error)
		List(ctx context.Context, in *do.ArticleTagListInput) (out *do.ArticleTagListOutput, err error)
		Add(ctx context.Context, in *do.ArticleTag) (lastInsertId int64, err error)
		Edit(ctx context.Context, in *do.ArticleTag) (affected int64, err error)
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IArticleComment interface {
		Find(ctx context.Context, in *do.ArticleCommentListInput) (out []*entity.ArticleComment, err error)
		List(ctx context.Context, in *do.ArticleCommentListInput) (out *do.ArticleCommentListOutput, err error)
		Add(ctx context.Context, in *do.ArticleComment) (lastInsertId int64, err error)
		Edit(ctx context.Context, in *do.ArticleComment) (affected int64, err error)
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
)

var (
	localArticleBase     IArticleBase
	localArticleCategory IArticleCategory
	localArticleTag      IArticleTag
	localArticleComment  IArticleComment
)

func ArticleBase() IArticleBase {
	if localArticleBase == nil {
		panic("implement not found for interface IArticleBase, forgot register?")
	}
	return localArticleBase
}

func RegisterArticleBase(i IArticleBase) {
	localArticleBase = i
}

func ArticleCategory() IArticleCategory {
	if localArticleCategory == nil {
		panic("implement not found for interface IArticleCategory, forgot register?")
	}
	return localArticleCategory
}

func RegisterArticleCategory(i IArticleCategory) {
	localArticleCategory = i
}
func ArticleTag() IArticleTag {
	if localArticleTag == nil {
		panic("implement not found for interface IArticleTag, forgot register?")
	}
	return localArticleTag
}

func RegisterArticleTag(i IArticleTag) {
	localArticleTag = i
}
func ArticleComment() IArticleComment {
	if localArticleComment == nil {
		panic("implement not found for interface IArticleComment, forgot register?")
	}
	return localArticleComment
}

func RegisterArticleComment(i IArticleComment) {
	localArticleComment = i
}
