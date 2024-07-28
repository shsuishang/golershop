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
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"strings"
)

type sArticleBase struct{}

func init() {
	service.RegisterArticleBase(NewArticleBase())
}

func NewArticleBase() *sArticleBase {
	return &sArticleBase{}
}

// Find 查询数据
func (s *sArticleBase) Find(ctx context.Context, in *do.ArticleBaseListInput) (out []*entity.ArticleBase, err error) {
	out, err = dao.ArticleBase.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sArticleBase) List(ctx context.Context, in *do.ArticleBaseListInput) (out *model.ArticleBaseOutput, err error) {
	list, err := dao.ArticleBase.List(ctx, in)
	gconv.Scan(list, &out)

	if !g.IsEmpty(out) {
		for _, item := range out.Items {
			if gstr.Trim(item.ArticleTags) != "" {
				tags := gstr.Split(item.ArticleTags, ",")
				tagList, err := dao.ArticleTag.Gets(ctx, tags)
				if err != nil {
					return nil, err
				}
				item.ArticleTagList = tagList
			}
		}
	}

	return out, err
}

// Add 新增
func (s *sArticleBase) Add(ctx context.Context, in *do.ArticleBase) (lastInsertId int64, err error) {
	lastInsertId, err = dao.ArticleBase.Add(ctx, in)

	if err != nil {
		return 0, err
	}

	//判断所属分类，为分类内容数量+1
	articleCategory, err := dao.ArticleCategory.Get(ctx, in.CategoryId)
	if err != nil {
		return 0, err
	}
	if articleCategory != nil {
		articleCategory.CategoryCount++
		NewArticleCategory := do.ArticleCategory{}
		gconv.Scan(articleCategory, &NewArticleCategory)
		_, err = dao.ArticleCategory.Edit(ctx, articleCategory.CategoryId, &NewArticleCategory)
		if err != nil {
			return 0, err
		}
	}

	//文章标签, 为每个标签内容数量+1
	err = s.manageArticleTags(ctx, in.ArticleTags, consts.ADD)
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

func (s *sArticleBase) manageArticleTags(ctx context.Context, tag interface{}, action string) error {
	if !g.IsEmpty(tag) {
		tags := strings.Split(tag.(string), ",")
		for _, t := range tags {
			articleTag, err := dao.ArticleTag.Get(ctx, t)
			if err != nil {
				return err
			}
			if articleTag != nil {
				if action == consts.ADD {
					articleTag.TagCount++
				} else if action == consts.SUBTRACT {
					if articleTag.TagCount > 0 {
						articleTag.TagCount--
					}
				}
				NewArticleTag := do.ArticleTag{}
				gconv.Scan(articleTag, &NewArticleTag)
				_, err = dao.ArticleTag.Edit(ctx, articleTag.TagId, &NewArticleTag)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Edit 编辑
func (s *sArticleBase) Edit(ctx context.Context, in *do.ArticleBase) (affected int64, err error) {
	//判断所属分类，为分类内容数量+1
	articleBase, err := dao.ArticleBase.Get(ctx, in.ArticleId)
	if err != nil {
		return 0, err
	}
	if articleBase != nil {
		if articleBase.CategoryId != in.CategoryId {
			lastArticleCategory, err := dao.ArticleCategory.Get(ctx, articleBase.CategoryId)
			if err != nil {
				return 0, err
			}
			if lastArticleCategory != nil {
				lastArticleCategory.CategoryCount--
				articleCategory := do.ArticleCategory{}
				gconv.Scan(lastArticleCategory, &articleCategory)
				_, err = dao.ArticleCategory.Edit(ctx, lastArticleCategory.CategoryId, &articleCategory)
				if err != nil {
					return 0, err
				}
			}

			newArticleCategory, err := dao.ArticleCategory.Get(ctx, in.CategoryId)
			if err != nil {
				return 0, err
			}
			if newArticleCategory != nil {
				newArticleCategory.CategoryCount++
				articleCategory := do.ArticleCategory{}
				gconv.Scan(newArticleCategory, &articleCategory)
				_, err = dao.ArticleCategory.Edit(ctx, newArticleCategory.CategoryId, &articleCategory)
				if err != nil {
					return 0, err
				}
			}
		}

		if articleBase.ArticleTags != in.ArticleTags {
			err = s.manageArticleTags(ctx, articleBase.ArticleTags, consts.SUBTRACT)
			if err != nil {
				return 0, err
			}

			err = s.manageArticleTags(ctx, in.ArticleTags, consts.ADD)
			if err != nil {
				return 0, err
			}
		}
	}
	_, err = dao.ArticleBase.Edit(ctx, in.ArticleId, in)

	if err != nil {
		return 0, err
	}

	return
}

// Remove 删除多条记录模式
func (s *sArticleBase) Remove(ctx context.Context, id any) (affected int64, err error) {

	err = s.articleBaseRelevance(ctx, id)
	if err != nil {
		return 0, err
	}

	affected, err = dao.ArticleBase.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// RemoveBatch 删除多条记录模式
func (s *sArticleBase) RemoveBatch(ctx context.Context, id string) (affected int64, err error) {
	idSlice := gstr.Split(id, ",")
	for _, articleBaseId := range idSlice {
		err = s.articleBaseRelevance(ctx, articleBaseId)
		if err != nil {
			return 0, err
		}
		affected, err = dao.ArticleBase.Remove(ctx, articleBaseId)
		if err != nil {
			return 0, err
		}
	}

	return affected, err
}

func (s *sArticleBase) articleBaseRelevance(ctx context.Context, id any) error {
	articleBase, err := dao.ArticleBase.Get(ctx, id)
	if err != nil {
		return err
	}

	//判断所属分类，为分类内容数量-1
	if articleBase != nil {
		articleCategory, err := dao.ArticleCategory.Get(ctx, articleBase.CategoryId)
		if err != nil {
			return err
		}
		if articleCategory != nil {
			articleCategory.CategoryCount--
			NewArticleCategory := do.ArticleCategory{}
			gconv.Scan(articleCategory, &NewArticleCategory)
			_, err = dao.ArticleCategory.Edit(ctx, articleCategory.CategoryId, &NewArticleCategory)
			if err != nil {
				return err
			}
		}

		//文章标签, 为每个标签内容数量-1
		err = s.manageArticleTags(ctx, articleBase.ArticleTags, consts.SUBTRACT)
		if err != nil {
			return err
		}

		// 删除该文章下所有评论
		articleComment := do.ArticleCommentListInput{}
		articleComment.Where.ArticleId = articleBase.ArticleId
		commentIds, _ := dao.ArticleComment.FindKey(ctx, &articleComment)
		_, err = dao.ArticleComment.Remove(ctx, commentIds)
	}

	return nil
}
