package cms

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/cms"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	Article = cArticle{}
)

type cArticle struct{}

// ListCategory 文章分类-分页列表查询
func (c *cArticle) ListCategory(ctx context.Context, req *cms.ArticleListCategoryReq) (res *cms.ArticleListCategoryRes, err error) {

	input := do.ArticleCategoryListInput{}
	gconv.Scan(req, &input)

	// 查询文章分类分页列表
	result, err := service.ArticleCategory().List(ctx, &input)
	if err != nil {
		return nil, err
	}

	// 将查询结果转换为响应参数
	gconv.Scan(result, &res)
	return res, nil
}

// List 获取文章基础列表
func (c *cArticle) List(ctx context.Context, req *cms.ArticleListReq) (res *cms.ArticleListRes, err error) {

	input := do.ArticleBaseListInput{}
	gconv.Scan(req, &input)

	input.Where.CategoryId = req.CategoryId

	result, err := service.ArticleBase().List(ctx, &input)
	if err != nil {
		return nil, err
	}

	gconv.Scan(result, &res)

	return res, nil
}

// Detail 获取文章基础信息
func (c *cArticle) Detail(ctx context.Context, req *cms.ArticleDetailReq) (res *cms.ArticleDetailRes, err error) {
	articleBaseRes := &cms.ArticleDetailRes{}
	// 获取文章基础信息
	articleBase, err := dao.ArticleBase.Get(ctx, req.ArticleId)
	if err != nil {
		return nil, err
	}

	// 如果文章基础信息存在，进行属性复制
	if articleBase != nil {
		gconv.Scan(articleBase, articleBaseRes)

		// 获取用户信息
		userInfo, err := dao.UserInfo.Get(ctx, articleBase.UserId)
		if err != nil {
			return nil, err
		}

		// 如果用户信息存在，设置用户昵称
		if userInfo != nil {
			articleBaseRes.UserNickname = userInfo.UserNickname
		}
	}

	return articleBaseRes, nil
}
