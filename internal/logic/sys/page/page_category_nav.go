package page

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
)

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

type sPageCategoryNav struct{}

func init() {
	service.RegisterPageCategoryNav(NewPageCategoryNav())
}

func NewPageCategoryNav() *sPageCategoryNav {
	return &sPageCategoryNav{}
}

// Find 查询数据
func (s *sPageCategoryNav) Find(ctx context.Context, in *do.PageCategoryNavListInput) (out []*entity.PageCategoryNav, err error) {
	out, err = dao.PageCategoryNav.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sPageCategoryNav) List(ctx context.Context, in *do.PageCategoryNavListInput) (out *do.PageCategoryNavListOutput, err error) {
	out, err = dao.PageCategoryNav.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sPageCategoryNav) Add(ctx context.Context, in *do.PageCategoryNav) (lastInsertId int64, err error) {
	lastInsertId, err = dao.PageCategoryNav.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sPageCategoryNav) Edit(ctx context.Context, in *do.PageCategoryNav) (affected int64, err error) {
	_, err = dao.PageCategoryNav.Edit(ctx, in.CategoryNavId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sPageCategoryNav) Remove(ctx context.Context, id any) (affected int64, err error) {

	affected, err = dao.PageCategoryNav.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// PcLayout PC头尾数据
func (s *sPageCategoryNav) GetPcLayout(ctx context.Context) (resultSlice []interface{}, err error) {
	userId := service.BizCtx().GetUserId(ctx)

	// 读取页面分类导航
	pageCategoryNavs, err := dao.PageCategoryNav.Find(ctx, &do.PageCategoryNavListInput{
		Where: do.PageCategoryNav{CategoryNavEnable: true},
		BaseList: ml.BaseList{
			Sort: "asc",
			Sidx: dao.PageCategoryNav.Columns().CategoryNavOrder,
		},
	})
	if err != nil {
		return nil, err
	}

	// 无用数据较多，可以独立封装一个数据，只传必须数据
	productCategoryListReq := &do.ProductCategoryListInput{
		Where: do.ProductCategory{
			CategoryIsEnable: true,
		},
	}
	// 获取树状分类
	categoryRes, err := service.ProductCategory().GetTree(ctx, productCategoryListReq, 0)
	if err != nil {
		return nil, err
	}

	// 读取所有商品信息
	itemIds := make([]uint64, 0)
	for _, item := range pageCategoryNavs {
		// 处理商品信息
		if item.CategoryNavType == 2 {
			itemIds = append(itemIds, gconv.SliceUint64(item.ItemIds)...)
		}
	}

	var items []*model.ProductItemVo
	if len(itemIds) > 0 {
		items, err = service.ProductBase().GetItems(ctx, itemIds, userId)
		if err != nil {
			return nil, err
		}
	}

	input := &model.PageCategoryNavVo{}
	for i := range pageCategoryNavs {
		item := pageCategoryNavs[i]
		gconv.Scan(item, input)
		// 处理分类编号
		if input.CategoryNavType == 1 {
			categoryId := gconv.Uint(input.CategoryIds)
			if categoryId != 0 {
				for _, category := range categoryRes {
					if category.CategoryId == categoryId {
						input.ProductCategoryTree = category
						break
					}
				}
			}
		}

		// 处理商品信息
		if input.CategoryNavType == 2 {
			var productItemVoList []*model.ProductItemVo
			itemIdList := gconv.SliceUint64(gstr.Split(input.ItemIds, ","))
			for _, itemId := range itemIdList {
				for _, productItem := range items {
					if productItem.ItemId == itemId {
						productItemVoList = append(productItemVoList, productItem)
						break
					}
				}
			}
			input.ProductItems = productItemVoList
		}
		resultSlice = append(resultSlice, input)
	}

	return resultSlice, nil
}
