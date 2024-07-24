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

package product

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/util/gutil"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pt"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
)

type sProductBrand struct{}

func init() {
	service.RegisterProductBrand(NewProductBrand())
}

func NewProductBrand() *sProductBrand {
	return &sProductBrand{}
}

// Get 读取品牌
func (s *sProductBrand) Get(ctx context.Context, id any) (out *entity.ProductBrand, err error) {
	var list []*entity.ProductBrand
	list, err = s.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// Gets 读取多条品牌
func (s *sProductBrand) Gets(ctx context.Context, id any) (list []*entity.ProductBrand, err error) {
	err = dao.ProductBrand.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sProductBrand) Find(ctx context.Context, in *do.ProductBrandListInput) (out []*entity.ProductBrand, err error) {
	out, err = dao.ProductBrand.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sProductBrand) List(ctx context.Context, in *do.ProductBrandListInput) (out *do.ProductBrandListOutput, err error) {
	out, err = dao.ProductBrand.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sProductBrand) Add(ctx context.Context, in *do.ProductBrand) (lastInsertId int64, err error) {
	lastInsertId, err = dao.ProductBrand.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sProductBrand) Edit(ctx context.Context, in *do.ProductBrand) (affected int64, err error) {
	_, err = dao.ProductBrand.Edit(ctx, in.BrandId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sProductBrand) Remove(ctx context.Context, id any) (affected int64, err error) {
	//是否有子项
	brandCount, err := dao.ProductType.Ctx(ctx).Count(do.ProductType{BrandIds: id})

	if err != nil {
		return 0, err
	}

	if brandCount > 0 {
		return 0, errors.New(fmt.Sprintf("有 %d 条类型使用，不可删除", brandCount))
	}

	affected, err = dao.ProductBrand.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// GetTree 获取商品品牌树
func (s *sProductBrand) GetTree(ctx context.Context) (brandList []*pt.BrandTreeRes, err error) {
	// 获取品牌列表并按brand_id升序排序
	brandLists, err := dao.ProductBrand.Find(ctx, &do.ProductBrandListInput{
		BaseList: ml.BaseList{
			Sidx: dao.ProductBrand.Columns().BrandId,
			Sort: ml.ORDER_BY_ASC,
		},
	})
	if err != nil {
		return nil, err
	}

	// 提取category_id列表
	categoryIds := array.Column(brandLists, "BrandId")

	// 获取商品分类
	categories, err := service.ProductCategory().Gets(ctx, categoryIds)
	if err != nil {
		return nil, err
	}

	// 构建品牌树
	for _, category := range categories {
		var brands []*entity.ProductBrand
		for _, brand := range brandLists {
			if brand.CategoryId == category.CategoryId {
				brands = append(brands, brand)
			}
		}

		if !gutil.IsEmpty(brands) {
			brandRe := &pt.BrandTreeRes{
				BrandId:   category.CategoryId,
				BrandName: category.CategoryName,
				Children:  brands,
			}
			brandList = append(brandList, brandRe)
		}
	}

	return brandList, nil
}
