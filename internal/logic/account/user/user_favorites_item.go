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

package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/api/shop"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
)

type sUserFavoritesItem struct{}

func init() {
	service.RegisterUserFavoritesItem(NewUserFavoritesItem())
}

func NewUserFavoritesItem() *sUserFavoritesItem {
	return &sUserFavoritesItem{}
}

// Get 读取兑换码
func (s *sUserFavoritesItem) Get(ctx context.Context, id any) (out *entity.UserFavoritesItem, err error) {
	var list []*entity.UserFavoritesItem
	list, err = s.Gets(ctx, id)

	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return out, nil
}

// Gets 读取多条兑换码
func (s *sUserFavoritesItem) Gets(ctx context.Context, id any) (list []*entity.UserFavoritesItem, err error) {
	err = dao.UserFavoritesItem.Ctx(ctx).WherePri(id).Scan(&list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// Find 查询数据
func (s *sUserFavoritesItem) Find(ctx context.Context, in *do.UserFavoritesItemListInput) (out []*entity.UserFavoritesItem, err error) {
	out, err = dao.UserFavoritesItem.Find(ctx, in)

	return out, err
}

// Find 查询数据
func (s *sUserFavoritesItem) FindOne(ctx context.Context, in *do.UserFavoritesItemListInput) (out *entity.UserFavoritesItem, err error) {
	out, err = dao.UserFavoritesItem.FindOne(ctx, in)

	return out, err
}

// List 分页读取
func (s *sUserFavoritesItem) List(ctx context.Context, in *do.UserFavoritesItemListInput) (out *do.UserFavoritesItemListOutput, err error) {
	out, err = dao.UserFavoritesItem.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sUserFavoritesItem) Add(ctx context.Context, in *do.UserFavoritesItem) (lastInsertId int64, err error) {
	lastInsertId, err = dao.UserFavoritesItem.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sUserFavoritesItem) Edit(ctx context.Context, in *do.UserFavoritesItem) (affected int64, err error) {
	_, err = dao.UserFavoritesItem.Edit(ctx, in.FavoritesItemId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sUserFavoritesItem) Remove(ctx context.Context, id any) (affected int64, err error) {
	affected, err = dao.UserFavoritesItem.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// GetList 读取用户收藏列表
func (s *sUserFavoritesItem) GetList(ctx context.Context, req *do.UserFavoritesItemListInput) (res *shop.UserFavoritesItemListsRes, err error) {
	userId := service.BizCtx().GetUserId(ctx)

	// 获取收藏项目列表
	favoritesItemPage, err := dao.UserFavoritesItem.List(ctx, req)
	if err != nil {
		return nil, err
	}

	if !g.IsEmpty(favoritesItemPage.Items) {
		res = &shop.UserFavoritesItemListsRes{}
		gconv.Scan(favoritesItemPage, res)
		favoritesItems := favoritesItemPage.Items
		itemIdsInterface := array.Column(favoritesItems, "ItemId")
		itemIds := make([]uint64, len(itemIdsInterface))
		for i, v := range itemIdsInterface {
			switch id := v.(type) {
			case uint64:
				itemIds[i] = id
			case int64:
				itemIds[i] = uint64(id)
			}
		}

		// 获取商品项目信息
		productItemVos, err := service.ProductBase().GetItems(ctx, itemIds, userId)
		if err != nil {
			return nil, err
		}

		if !g.IsEmpty(productItemVos) {
			productItemVoMap := make(map[uint64]*model.ProductItemVo)
			for _, productItemVo := range productItemVos {
				productItemVoMap[productItemVo.ItemId] = productItemVo
			}

			itemRes := make([]map[string]interface{}, 0)
			// 遍历收藏夹商品，合并商品详细信息
			for _, favoritesItem := range favoritesItems {
				userFavoritesItemRes := make(map[string]interface{})
				gconv.Struct(favoritesItem, &userFavoritesItemRes)

				if productItemVo, found := productItemVoMap[favoritesItem.ItemId]; found {

					userFavoritesItemRes["product_item_name"] = productItemVo.ProductName
					userFavoritesItemRes["item_unit_price"] = productItemVo.ItemUnitPrice
					userFavoritesItemRes["product_image"] = productItemVo.ProductImage
				}

				itemRes = append(itemRes, userFavoritesItemRes)
			}

			res.Items = itemRes
		}
	}

	return res, err
}
