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
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/cms"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ArticleCategory = cArticleCategory{}
)

type cArticleCategory struct{}

// =================== 管理端使用 =========================

// Tree 树形菜单
func (c *cArticleCategory) Tree(ctx context.Context, req *cms.ArticleCategoryTreeReq) (res cms.ArticleCategoryTreeRes, err error) {

	input := do.ArticleCategoryListInput{}
	gconv.Scan(req, &input)

	ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)

	var result, error = service.ArticleCategory().GetTree(ctx, &input)

	if error != nil {
		err = error
	}

	res, err = service.ArticleCategory().GetTree(ctx, &input)

	gconv.Scan(result, &res)

	return
}

// Add 新增菜单
func (c *cArticleCategory) Add(ctx context.Context, req *cms.ArticleCategoryAddReq) (res *cms.ArticleCategoryEditRes, err error) {

	input := do.ArticleCategory{}
	gconv.Scan(req, &input)

	var result, error = service.ArticleCategory().Add(ctx, &input)
	//var result, error = service.ArticleCategory().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &cms.ArticleCategoryEditRes{
		CategoryId: result,
	}

	return
}

// Edit 编辑菜单
func (c *cArticleCategory) Edit(ctx context.Context, req *cms.ArticleCategoryEditReq) (res *cms.ArticleCategoryEditRes, err error) {

	input := do.ArticleCategory{}
	gconv.Scan(req, &input)

	var result, error = service.ArticleCategory().Edit(ctx, &input)
	//var result, error = service.ArticleCategory().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &cms.ArticleCategoryEditRes{
		CategoryId: result,
	}

	return
}

// EditState 编辑菜单
func (c *cArticleCategory) EditState(ctx context.Context, req *cms.ArticleCategoryEditStateReq) (res *cms.ArticleCategoryEditStateRes, err error) {

	input := do.ArticleCategory{}
	gconv.Scan(req, &input)

	var result, error = service.ArticleCategory().Edit(ctx, &input)
	//var result, error = service.ArticleCategory().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &cms.ArticleCategoryEditStateRes{
		CategoryId: result,
	}

	return
}

// Remove 删除菜单
func (c *cArticleCategory) Remove(ctx context.Context, req *cms.ArticleCategoryRemoveReq) (res *cms.ArticleCategoryRemoveRes, err error) {

	var _, error = service.ArticleCategory().Remove(ctx, req.CategoryId)

	/*
		input := do.ArticleCategory{}
		input.ArticleCategoryTime = gtime.Now()
		input.ArticleCategoryId = req.ArticleCategoryId[0]
		input.ArticleCategorySort = 0

		var _, error = service.ArticleCategory().Edit(ctx, &input)
	*/

	if error != nil {
		err = error
	}

	res = &cms.ArticleCategoryRemoveRes{}

	return
}
