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
	"golershop.cn/api/pt"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
)

type sUserProductBrowse struct{}

func init() {
	service.RegisterUserProductBrowse(NewUserProductBrowse())
}

func NewUserProductBrowse() *sUserProductBrowse {
	return &sUserProductBrowse{}
}

// Find 查询数据
func (s *sUserProductBrowse) Find(ctx context.Context, in *do.UserProductBrowseListInput) (out []*entity.UserProductBrowse, err error) {
	out, err = dao.UserProductBrowse.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sUserProductBrowse) List(ctx context.Context, in *do.UserProductBrowseListInput) (out *do.UserProductBrowseListOutput, err error) {
	out, err = dao.UserProductBrowse.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sUserProductBrowse) Add(ctx context.Context, in *do.UserProductBrowse) (lastInsertId int64, err error) {
	lastInsertId, err = dao.UserProductBrowse.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sUserProductBrowse) Edit(ctx context.Context, in *do.UserProductBrowse) (affected int64, err error) {
	_, err = dao.UserProductBrowse.Edit(ctx, in.ProductBrowseId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sUserProductBrowse) Remove(ctx context.Context, id any) (affected int64, err error) {
	affected, err = dao.UserProductBrowse.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

/*
// GetList 获取用户浏览记录列表
func (s *sUserProductBrowse) GetList(ctx context.Context, userProductBrowseListReq *pt.UserProductBrowseListReq) ([]*pt.UserProductBrowseRes, error) {
	userProductBrowseRes := []*pt.UserProductBrowseRes{}
	userId := userProductBrowseListReq.UserId
	cacheKey := fmt.Sprintf("user_id|%d", userId)
	productBrowses := redis.GetVar(ctx, cacheKey).([]*model.UserProductBrowse)

	if len(productBrowses) > 0 {
		// 提取浏览商品的itemIds
		var itemIds []uint64
		visitedItemMap := make(map[uint64]bool)
		for _, pb := range productBrowses {
			if _, exists := visitedItemMap[pb.ItemId]; !exists {
				visitedItemMap[pb.ItemId] = true
				itemIds = append(itemIds, pb.ItemId)
			}
		}

		// 查询商品信息
		productItemVos, err := service.ProductBase().GetItems(ctx, itemIds, nil)
		if err != nil {
			return nil, err
		}

		// 构建商品ID到商品信息的映射
		productItemVoMap := make(map[uint64]*entity.ProductItemVo)
		for _, productItemVo := range productItemVos {
			productItemVoMap[productItemVo.ItemId] = productItemVo
		}

		// 构建返回结果
		for _, pb := range productBrowses {
			browseRes := &model.UserProductBrowseRes{}
			gconv.Struct(pb, browseRes)

			if productItemVo, ok := productItemVoMap[pb.ItemId]; ok {
				browseRes.ProductImage = productItemVo.ProductImage
				browseRes.ItemSalePrice = productItemVo.ItemSalePrice
				browseRes.ProductItemName = productItemVo.ProductItemName
				browseRes.ProductName = productItemVo.ProductName
				browseRes.ItemName = productItemVo.ItemName

				if activityInfo := productItemVo.ActivityInfo; activityInfo != nil {
					browseRes.ActivityTypeId = activityInfo.ActivityTypeId
					browseRes.ActivityTypeName = activityInfo.ActivityTypeName
				}
			}

			userProductBrowseRes = append(userProductBrowseRes, browseRes)
		}
	}

	return userProductBrowseRes, nil
}

*/

// AddBrowser 添加浏览记录
func (s *sUserProductBrowse) AddBrowser(ctx context.Context, itemId uint64, userId uint) (productBrowses []*entity.UserProductBrowse, err error) {

	/*	productBrowse := &entity.UserProductBrowse{
			ItemId:     itemId,
			UserId:     userId,
			BrowseTime: time.Now().Unix(),
		}
		cacheKey := fmt.Sprintf("user_id|%d", userId)

			// 从缓存中获取浏览记录
			productBrowses = gconv.SliceAnyToEntity(dao.UserProductBrowse.Get(ctx, cacheKey), new(entity.UserProductBrowse)).([]*entity.UserProductBrowse)
			if len(productBrowses) > 0 {
				// 去除重复记录
				for i := 0; i < len(productBrowses); i++ {
					if productBrowses[i].ItemId == itemId {
						productBrowses = append(productBrowses[:i], productBrowses[i+1:]...)
						break
					}
				}
				// 如果浏览记录达到上限，删除最后一条记录
				if len(productBrowses) == 10 {
					productBrowses = productBrowses[:9]
				}
				// 添加新记录到首位
				productBrowses = append([]*entity.UserProductBrowse{productBrowse}, productBrowses...)
			} else {
				productBrowses = []*entity.UserProductBrowse{productBrowse}
			}

			// 更新缓存
			dao.UserProductBrowse.Set(ctx, cacheKey, productBrowses)

	*/

	return productBrowses, nil
}

// RemoveBrowser 删除浏览记录
func (s *sUserProductBrowse) RemoveBrowser(ctx context.Context, userProductBrowseListReq *pt.UserProductBrowseListReq) (success bool, err error) {

	/*	userId := userProductBrowseListReq.UserId
		itemId := userProductBrowseListReq.ItemId
		cacheKey := fmt.Sprintf("user_id|%d", userId)

			// 从缓存中获取浏览记录
			productBrowses := gconv.SliceAnyToEntity(dao.UserProductBrowse.Get(ctx, cacheKey), new(entity.UserProductBrowse)).([]*entity.UserProductBrowse)

			// 删除匹配的浏览记录
			if len(productBrowses) > 0 {
				for i := 0; i < len(productBrowses); i++ {
					if productBrowses[i].ItemId == itemId {
						productBrowses = append(productBrowses[:i], productBrowses[i+1:]...)
						break
					}
				}
			}

			// 更新缓存
			dao.UserProductBrowse.Set(ctx, cacheKey, productBrowses)


	*/
	return true, nil
}
