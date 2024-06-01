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

package pt

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/pt"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	ProductCategory = cProductCategory{}
)

type cProductCategory struct{}

// =================== 管理端使用 =========================

// Tree 树形分类
func (c *cProductCategory) Tree(ctx context.Context, req *pt.ProductCategoryTreeReq) (res pt.ProductCategoryTreeRes, err error) {

	input := do.ProductCategoryListInput{}
	gconv.Scan(req, &input)

	//ml.ConvertReqToInputWhere(req, &input.Where, &input.WhereExt)
	if !g.IsEmpty(req.CategoryName) {
		input.Where.CategoryName = req.CategoryName
	}

	var result1, error1 = service.ProductCategory().GetTree(ctx, &input, 0)

	if error1 != nil {
		err = error1
	}

	gconv.Scan(result1, &res)

	return
}

// Add 新增分类
func (c *cProductCategory) Add(ctx context.Context, req *pt.ProductCategoryAddReq) (res *pt.ProductCategoryEditRes, err error) {

	input := do.ProductCategory{}
	gconv.Scan(req, &input)

	var result, error = service.ProductCategory().Add(ctx, &input)
	//var result, error = service.ProductCategory().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &pt.ProductCategoryEditRes{
		CategoryId: result,
	}

	return
}

// Edit 编辑分类
func (c *cProductCategory) Edit(ctx context.Context, req *pt.ProductCategoryEditReq) (res *pt.ProductCategoryEditRes, err error) {

	input := do.ProductCategory{}
	gconv.Scan(req, &input)

	var result, error = service.ProductCategory().Edit(ctx, &input)
	//var result, error = service.ProductCategory().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &pt.ProductCategoryEditRes{
		CategoryId: result,
	}

	return
}

// EditState 编辑分类
func (c *cProductCategory) EditState(ctx context.Context, req *pt.ProductCategoryEditStateReq) (res *pt.ProductCategoryEditStateRes, err error) {
	input := do.ProductCategory{}
	gconv.Scan(req, &input)

	var result, error = service.ProductCategory().Edit(ctx, &input)

	if error != nil {
		err = error
	}

	res = &pt.ProductCategoryEditStateRes{
		CategoryId: result,
	}

	return
}

// Remove 删除分类
func (c *cProductCategory) Remove(ctx context.Context, req *pt.ProductCategoryRemoveReq) (res *pt.ProductCategoryRemoveRes, err error) {

	var _, error = service.ProductCategory().Remove(ctx, req.CategoryId)

	/*
		input := do.ProductCategory{}
		input.ProductCategoryTime = gtime.Now()
		input.ProductCategoryId = req.ProductCategoryId[0]
		input.ProductCategorySort = 0

		var _, error = service.ProductCategory().Edit(ctx, &input)
	*/

	if error != nil {
		err = error
	}

	res = &pt.ProductCategoryRemoveRes{}

	return
}
