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

	"golershop.cn/api/pt"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
)

type (
	IProductAssistItem interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.ProductAssistItemListInput) (out []*entity.ProductAssistItem, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ProductAssistItemListInput) (out *do.ProductAssistItemListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.ProductAssistItem) (lastInsertId int64, err error)
		UpdateAssist(ctx context.Context, assistId interface{}) (flag bool, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ProductAssistItem) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IProductBase interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.ProductBaseListInput) (out []*entity.ProductBase, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ProductBaseListInput) (out *do.ProductBaseListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.ProductBase) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ProductBase) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// SaveProdcut 添加或者编辑商品
		SaveProdcut(ctx context.Context, in *model.SaveProductInput) (productId uint64, err error)
		// RemoveProdcut 删除商品
		RemoveProdcut(ctx context.Context, id any) (affected int64, err error)
		// GetProduct 读取商品信息
		GetProduct(ctx context.Context, id any) (productData model.ProductDateOutput, err error)
		// GetItems 读取商品信息
		GetItems(ctx context.Context, itemIds []uint64, userId uint) (out []*model.ProductItemVo, err error)
		// BatchEditState 批量编辑商品状态
		BatchEditState(ctx context.Context, productIds []uint64, productStateId uint) (result bool, err error)
	}
	IProductCategory interface {
		// 读取商品分类
		Get(ctx context.Context, id any) (out *entity.ProductCategory, err error)
		// 读取多条记录模式
		Gets(ctx context.Context, id any) (list []*entity.ProductCategory, err error)
		// 分页读取
		List(ctx context.Context, in *do.ProductCategoryListInput) (out *do.ProductCategoryListOutput, err error)
		// 查询数据
		GetTree(ctx context.Context, in *do.ProductCategoryListInput, pid uint) (out []*model.CategoryTreeNode, err error)
		// 读取下级分类叶节点分类编号
		GetCategoryLeafs(ctx context.Context, pid uint) (out *[]uint, err error)
		// GetParentCategory 根据 ID 查询所有父级（包含自身）
		GetParentCategory(ctx context.Context, categoryId uint) (output []*entity.ProductCategory, err error)
		// 遍历树形结构，将所有节点的 ID 放入数组中
		TraverseTreeAndGetIds(ctx context.Context, root []*model.CategoryTreeNode, ids *[]uint)
		// GetSearchFilter 分类过滤搜索选项
		GetSearchFilter(ctx context.Context, categoryId uint) (output *pt.SearchFilterRes, err error)
		// 新增
		Add(ctx context.Context, in *do.ProductCategory) (out int64, err error)
		// 编辑
		Edit(ctx context.Context, in *do.ProductCategory) (affected int64, err error)
		// 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IProductIndex interface {
		// GetList 商品搜索查询列表
		GetList(ctx context.Context, in *do.ProductIndexListInput) (out *model.ProductListOutput, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.ProductIndexListInput) (out []*entity.ProductIndex, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ProductIndexListInput) (out *do.ProductIndexListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.ProductIndex) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ProductIndex) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// ListItem 读取SKU商品
		ListItem(ctx context.Context, req *pt.ItemListReq) (out *model.ItemListOutput, err error)
		// Detail 商品详情
		Detail(ctx context.Context, input *model.ProductDetailInput) (*model.ProductDetailOutput, error)
	}
	IProductItem interface {
		// Get 读取SKU
		Get(ctx context.Context, id any) (out *entity.ProductItem, err error)
		// Gets 读取多条SKU
		Gets(ctx context.Context, id any) (list []*entity.ProductItem, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.ProductItemListInput) (out []*entity.ProductItem, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ProductItemListInput) (out *do.ProductItemListOutput, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ProductItem) (affected int64, err error)
		// LockSkuStock 锁库存
		LockSkuStock(ctx context.Context, itemId uint64, cartQuantity uint) (affected int64, err error)
		// LockSkuStock 锁库存
		EditStock(ctx context.Context, in *model.ProductEditStockInput) (success bool, err error)
		// ReleaseSkuStock 锁库存
		ReleaseSkuStock(ctx context.Context, itemId uint64, releaseQuantity uint) (affected int64, err error)
		// List 分页读取
		ListItemKey(ctx context.Context, req *pt.ItemListReq) (out *do.ItemListKeyOutput, err error)
		IfOnSale(ctx context.Context, item *model.ProductItemVo) bool
	}
	IProductSpec interface {
		// Get 读取规格
		Get(ctx context.Context, id any) (out *entity.ProductSpec, err error)
		// Gets 读取多条规格
		Gets(ctx context.Context, id any) (list []*entity.ProductSpec, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.ProductSpecListInput) (out []*entity.ProductSpec, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ProductSpecListInput) (out *do.ProductSpecListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.ProductSpec) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ProductSpec) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IProductTag interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.ProductTagListInput) (out []*entity.ProductTag, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ProductTagListInput) (out *do.ProductTagListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.ProductTag) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ProductTag) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IProductComment interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.ProductCommentListInput) (out []*entity.ProductComment, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ProductCommentListInput) (out *do.ProductCommentListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.ProductComment) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ProductComment) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// GetList
		GetList(ctx context.Context, productCommentListReq *do.ProductCommentListInput) (productCommentPage *do.ProductCommentListOutput, err error)
	}
	IProductAssist interface {
		// Get 读取属性
		Get(ctx context.Context, id any) (out *entity.ProductAssist, err error)
		// Gets 读取多条属性
		Gets(ctx context.Context, id any) (list []*entity.ProductAssist, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.ProductAssistListInput) (out []*entity.ProductAssist, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ProductAssistListInput) (out *do.ProductAssistListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.ProductAssist) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ProductAssist) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// GetAssists 获取商品辅助属性
		GetAssists(ctx context.Context, assistIds string) (out []*model.ProductAssistOutput, err error)
	}
	IProductBrand interface {
		// Get 读取品牌
		Get(ctx context.Context, id any) (out *entity.ProductBrand, err error)
		// Gets 读取多条品牌
		Gets(ctx context.Context, id any) (list []*entity.ProductBrand, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.ProductBrandListInput) (out []*entity.ProductBrand, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ProductBrandListInput) (out *do.ProductBrandListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.ProductBrand) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ProductBrand) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
	}
	IProductSpecItem interface {
		// Find 查询数据
		Find(ctx context.Context, in *do.ProductSpecItemListInput) (out []*entity.ProductSpecItem, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ProductSpecItemListInput) (out *do.ProductSpecItemListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.ProductSpecItem) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ProductSpecItem) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// EditState 编辑任务状态
		EditState(ctx context.Context, req *pt.ProductSpecItemEditStateReq) (res *pt.ProductSpecItemEditStateRes, err error)
	}
	IProductType interface {
		// 读取类型
		Get(ctx context.Context, id any) (out *entity.ProductType, err error)
		// 读取多条类型
		Gets(ctx context.Context, id any) (list []*entity.ProductType, err error)
		// Find 查询数据
		Find(ctx context.Context, in *do.ProductTypeListInput) (out []*entity.ProductType, err error)
		// List 分页读取
		List(ctx context.Context, in *do.ProductTypeListInput) (out *do.ProductTypeListOutput, err error)
		// Add 新增
		Add(ctx context.Context, in *do.ProductType) (lastInsertId int64, err error)
		// Edit 编辑
		Edit(ctx context.Context, in *do.ProductType) (affected int64, err error)
		// Update
		UpdateAssistIds(ctx context.Context, typeId interface{}) (affected int64, err error)
		// Remove 删除多条记录模式
		Remove(ctx context.Context, id any) (affected int64, err error)
		// Info 读取类型信息
		Info(ctx context.Context, id any) (out *model.ProductTypeInfoOutput, err error)
	}
)

var (
	localProductAssist     IProductAssist
	localProductBrand      IProductBrand
	localProductSpecItem   IProductSpecItem
	localProductType       IProductType
	localProductSpec       IProductSpec
	localProductTag        IProductTag
	localProductComment    IProductComment
	localProductAssistItem IProductAssistItem
	localProductBase       IProductBase
	localProductCategory   IProductCategory
	localProductIndex      IProductIndex
	localProductItem       IProductItem
)

func ProductSpecItem() IProductSpecItem {
	if localProductSpecItem == nil {
		panic("implement not found for interface IProductSpecItem, forgot register?")
	}
	return localProductSpecItem
}

func RegisterProductSpecItem(i IProductSpecItem) {
	localProductSpecItem = i
}

func ProductType() IProductType {
	if localProductType == nil {
		panic("implement not found for interface IProductType, forgot register?")
	}
	return localProductType
}

func RegisterProductType(i IProductType) {
	localProductType = i
}

func ProductAssist() IProductAssist {
	if localProductAssist == nil {
		panic("implement not found for interface IProductAssist, forgot register?")
	}
	return localProductAssist
}

func RegisterProductAssist(i IProductAssist) {
	localProductAssist = i
}

func ProductBrand() IProductBrand {
	if localProductBrand == nil {
		panic("implement not found for interface IProductBrand, forgot register?")
	}
	return localProductBrand
}

func RegisterProductBrand(i IProductBrand) {
	localProductBrand = i
}

func ProductCategory() IProductCategory {
	if localProductCategory == nil {
		panic("implement not found for interface IProductCategory, forgot register?")
	}
	return localProductCategory
}

func RegisterProductCategory(i IProductCategory) {
	localProductCategory = i
}

func ProductIndex() IProductIndex {
	if localProductIndex == nil {
		panic("implement not found for interface IProductIndex, forgot register?")
	}
	return localProductIndex
}

func RegisterProductIndex(i IProductIndex) {
	localProductIndex = i
}

func ProductItem() IProductItem {
	if localProductItem == nil {
		panic("implement not found for interface IProductItem, forgot register?")
	}
	return localProductItem
}

func RegisterProductItem(i IProductItem) {
	localProductItem = i
}

func ProductSpec() IProductSpec {
	if localProductSpec == nil {
		panic("implement not found for interface IProductSpec, forgot register?")
	}
	return localProductSpec
}

func RegisterProductSpec(i IProductSpec) {
	localProductSpec = i
}

func ProductTag() IProductTag {
	if localProductTag == nil {
		panic("implement not found for interface IProductTag, forgot register?")
	}
	return localProductTag
}

func RegisterProductTag(i IProductTag) {
	localProductTag = i
}

func ProductComment() IProductComment {
	if localProductComment == nil {
		panic("implement not found for interface IProductComment, forgot register?")
	}
	return localProductComment
}

func RegisterProductComment(i IProductComment) {
	localProductComment = i
}

func ProductAssistItem() IProductAssistItem {
	if localProductAssistItem == nil {
		panic("implement not found for interface IProductAssistItem, forgot register?")
	}
	return localProductAssistItem
}

func RegisterProductAssistItem(i IProductAssistItem) {
	localProductAssistItem = i
}

func ProductBase() IProductBase {
	if localProductBase == nil {
		panic("implement not found for interface IProductBase, forgot register?")
	}
	return localProductBase
}

func RegisterProductBase(i IProductBase) {
	localProductBase = i
}
