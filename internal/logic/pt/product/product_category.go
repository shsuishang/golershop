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
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/pt"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"time"
)

type sProductCategory struct{}

var (
	cachePreKey   = "menu_base"
	cacheDuration = time.Hour
	redisAdapter  = gcache.NewAdapterRedis(g.Redis())
	cache         = gcache.NewWithAdapter(redisAdapter)
)

func init() {
	service.RegisterProductCategory(New())
}

func New() *sProductCategory {
	return &sProductCategory{}
}

// 读取商品分类
func (s *sProductCategory) Get(ctx context.Context, id any) (out *entity.ProductCategory, err error) {
	var list []*entity.ProductCategory
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
func (s *sProductCategory) Gets(ctx context.Context, id any) (list []*entity.ProductCategory, err error) {

	err = dao.ProductCategory.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// 分页读取
func (s *sProductCategory) List(ctx context.Context, in *do.ProductCategoryListInput) (out *do.ProductCategoryListOutput, err error) {
	list, err := dao.ProductCategory.List(ctx, in)

	gconv.Scan(list, &out)

	return out, nil
}

// 业务封装

// 查询数据
func (s *sProductCategory) GetTree(ctx context.Context, in *do.ProductCategoryListInput, pid uint) (out []*model.CategoryTreeNode, err error) {
	//in.Where.CategoryIsEnable = true
	in.Sidx = dao.ProductCategory.Columns().CategorySort
	in.Sort = "ASC"
	in.Size = consts.MAX_LIST_NUM

	res, err := dao.ProductCategory.List(ctx, in)

	// 数据转换
	var list []*entity.ProductCategory
	gconv.Scan(res.Items, &list)

	// 数据处理
	var categoryNode model.CategoryTreeNode
	categoryNode.CategoryId = pid
	categoryNode.Children = make([]*model.CategoryTreeNode, 0)

	s.makeTree(list, &categoryNode)

	/*
		columnIds := array.Column(list, dao.ProductCategory.Columns().CategoryId)

			for _, c := range list {
				if c.CategoryParentId != 0 && !array.InArray(columnIds, c.CategoryParentId) {
					child := &model.CategoryTreeNode{}
					child.Children = make([]*model.CategoryTreeNode, 0)

					//child.ProductCategory = *c
					gconv.Scan(*c, &child.ProductCategory)

					categoryNode.Children = append(categoryNode.Children, child)
				}
			}

	*/

	return categoryNode.Children, nil
}

// 递归生成分类列表
func (s *sProductCategory) makeTree(list []*entity.ProductCategory, tn *model.CategoryTreeNode) {
	for _, c := range list {
		if c.CategoryParentId == tn.CategoryId {
			child := &model.CategoryTreeNode{}
			child.Children = make([]*model.CategoryTreeNode, 0)

			//child.ProductCategory = *c
			gconv.Scan(*c, &child.ProductCategory)

			tn.Children = append(tn.Children, child)
			s.makeTree(list, child)
		}
	}
}

// 读取下级分类叶节点分类编号
func (s *sProductCategory) GetCategoryLeafs(ctx context.Context, pid uint) (out *[]uint, err error) {
	input := &do.ProductCategoryListInput{}
	//input.Where.CategoryParentId = pid
	//input.Where.CategoryIsEnable = false
	//input.Where.CategoryName = category_name

	// 定义存储叶子节点ID的切片
	ids := &[]uint{}
	// 调用获取分类树的函数，获取指定父节点下的分类树
	tree, err := s.GetTree(ctx, input, pid)
	if err != nil {
		return nil, err
	}

	// 遍历分类树数据
	s.TraverseTreeAndGetIds(ctx, tree, ids)

	// 返回叶子节点ID列表
	return ids, nil
}

// GetParentCategory 根据 ID 查询所有父级（包含自身）
func (s *sProductCategory) GetParentCategory(ctx context.Context, categoryId uint) (output []*entity.ProductCategory, err error) {
	var categoryList []*entity.ProductCategory
	i := 5 // 设置一个循环次数防止死循环

	for categoryId != 0 && i != 0 {
		category, err := s.Get(ctx, categoryId)
		if err != nil {
			// 处理错误
			break
		}

		if category != nil {
			categoryList = append(categoryList, category)
			categoryId = category.CategoryParentId
		} else {
			break
		}

		i--
	}

	return categoryList, err
}

// 遍历树形结构，将所有节点的 ID 放入数组中
func (s *sProductCategory) TraverseTreeAndGetIds(ctx context.Context, root []*model.CategoryTreeNode, ids *[]uint) {
	if root == nil {
		return
	}

	// 遍历子节点
	for _, child := range root {
		// 将当前节点的 ID 放入数组中
		*ids = append(*ids, child.CategoryId)

		s.TraverseTreeAndGetIds(ctx, child.Children, ids)
	}
}

// GetSearchFilter 分类过滤搜索选项
func (s *sProductCategory) GetSearchFilter(ctx context.Context, categoryId uint) (output *pt.SearchFilterRes, err error) {
	output = &pt.SearchFilterRes{}
	output.Contracts = make([]*entity.ContractType, 0)
	output.Markets = make([]interface{}, 0)

	output.Brands = make([]*entity.ProductBrand, 0)
	output.Parent = make([]*entity.ProductCategory, 0)
	output.Children = make([]*entity.ProductCategory, 0)
	output.Assists = make([]*model.ProductAssistOutput, 0)

	if categoryId != 0 {
		// 获取分类信息
		productCategory, err := s.Get(ctx, categoryId)
		if err != nil || productCategory == nil {
			// 处理错误
			return output, err
		}

		output.Info = productCategory

		// 获取上级分类
		parentCategoryListById, err := s.GetParentCategory(ctx, categoryId)
		if err != nil {
			// 处理错误
			return output, err
		}
		output.Parent = parentCategoryListById

		// 获取下级分类
		childCategorys, err := dao.ProductCategory.Find(ctx, &do.ProductCategoryListInput{
			Where: do.ProductCategory{CategoryParentId: categoryId, CategoryIsEnable: true},
		})
		if err != nil {
			// 处理错误
			return output, err
		}
		output.Children = childCategorys

		// 获取辅助属性
		productType, err := dao.ProductType.Get(ctx, productCategory.TypeId)
		if err != nil {
			// 处理错误
			return output, err
		}

		assists, err := service.ProductAssist().GetAssists(ctx, productType.AssistIds)
		if err != nil {
			// 处理错误
			return output, err
		}

		output.Assists = assists

		// 获取品牌
		if len(productType.BrandIds) > 0 {
			brandIds := gconv.SliceInt(productType.BrandIds)
			brandList, err := dao.ProductBrand.Gets(ctx, brandIds)
			if err != nil {
				// 处理错误
				return output, err
			}

			output.Brands = brandList
		} else {
			output.Brands = make([]*entity.ProductBrand, 0)
		}
	}

	return
}

// 新增
func (s *sProductCategory) Add(ctx context.Context, in *do.ProductCategory) (out int64, err error) {
	// 不允许HTML代码
	//if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
	//	return out, err
	//}

	lastInsertId, err := dao.ProductCategory.Add(ctx, in)
	if err != nil {
		return out, err
	}
	return lastInsertId, err
}

// 编辑
func (s *sProductCategory) Edit(ctx context.Context, in *do.ProductCategory) (affected int64, err error) {
	// 不允许HTML代码
	//if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
	//	return out, err
	//}

	affected, err = dao.ProductCategory.Edit(ctx, in.CategoryId, in)

	return affected, err
}

// 删除多条记录模式
func (s *sProductCategory) Remove(ctx context.Context, id any) (affected int64, err error) {

	count, err := dao.ProductCategory.Ctx(ctx).Count(do.ProductCategory{CategoryParentId: id})

	if err != nil {
		return 0, err
	}

	if count > 0 {
		return 0, errors.New(fmt.Sprintf("有 %d 个子级商品分类，不可删除", count))
	}

	//是否有子项
	typeCount, err := dao.ProductType.Ctx(ctx).Count(do.ProductType{CategoryId: id})

	if err != nil {
		return 0, err
	}

	if typeCount > 0 {
		return 0, errors.New(fmt.Sprintf("有 %d 条类型使用，不可删除", typeCount))
	}

	brandCount, err := dao.ProductBrand.Ctx(ctx).Count(do.ProductBrand{CategoryId: id})

	if err != nil {
		return 0, err
	}

	if brandCount > 0 {
		return 0, errors.New(fmt.Sprintf("有 %d 条品牌使用，不可删除", brandCount))
	}

	affected, err = dao.ProductCategory.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}
