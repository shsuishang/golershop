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
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/pt"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
	"strings"
)

type sProductIndex struct{}

func init() {
	service.RegisterProductIndex(NewProductIndex())
}

func NewProductIndex() *sProductIndex {
	return &sProductIndex{}
}

// GetList 商品搜索查询列表
func (s *sProductIndex) GetList(ctx context.Context, in *do.ProductIndexListInput) (out *model.ProductListOutput, err error) {
	in.Order = []*ml.BaseOrder{{Sidx: dao.ProductIndex.Columns().ProductOrder, Sort: ml.ORDER_BY_ASC}, {Sidx: dao.ProductIndex.Columns().ProductId, Sort: ml.ORDER_BY_DESC}}

	productIndexList, err := dao.ProductIndex.List(ctx, in)
	gconv.Scan(productIndexList, &out)

	//补全商品基础表信息
	ids := array.Column(productIndexList.Items, dao.ProductIndex.Columns().ProductId)

	if len(ids) > 0 {
		productBaseList, err := dao.ProductBase.Gets(ctx, ids)
		productInfoList, err := dao.ProductInfo.Gets(ctx, ids)
		productItems, err := dao.ProductItem.Find(ctx, &do.ProductItemListInput{Where: do.ProductItem{ProductId: ids}})

		// 处理为 map
		itemMap := make(map[uint64][]*entity.ProductItem)
		defaultItemMap := make(map[uint64]uint64)

		if len(productItems) > 0 {
			for _, item := range productItems {
				if _, ok := itemMap[item.ProductId]; !ok {
					itemMap[item.ProductId] = make([]*entity.ProductItem, 0)
				}

				itemMap[item.ProductId] = append(itemMap[item.ProductId], item)

				// 默认 item
				if item.ItemIsDefault {
					defaultItemMap[item.ProductId] = item.ItemId
				}
			}
		}

		//
		for _, item := range out.Items {
			for _, base := range productBaseList {
				if item.ProductId == base.ProductId {
					item.ProductTips = base.ProductTips
					item.ProductImage = base.ProductImage
					item.ProductVideo = base.ProductVideo
					item.TransportTypeId = base.TransportTypeId
					item.ProductBuyLimit = base.ProductBuyLimit
					item.ProductCommissionRate = base.ProductCommissionRate
				}
			}

			for _, info := range productInfoList {
				if item.ProductId == info.ProductId {
					item.ProductSpec = info.ProductSpec
					item.ProductUniqid = info.ProductUniqid
				}
			}

			// 默认
			item.ItemId = defaultItemMap[item.ProductId]
			item.Items = itemMap[item.ProductId]
		}

		// 活动价格
		var itemIds []uint64
		for _, record := range out.Items {
			itemIds = append(itemIds, record.ItemId)
		}
		activityItemList, err := dao.ActivityItem.Find(ctx, &do.ActivityItemListInput{Where: do.ActivityItem{ItemId: itemIds, ActivityItemState: consts.ACTIVITY_STATE_NORMAL}})
		if err != nil {
			// 处理错误
			return out, err
		}

		itemPriceMap := make(map[uint64]float64)
		for _, activityItem := range activityItemList {
			itemPriceMap[activityItem.ItemId] = activityItem.ActivityItemPrice
		}

		for _, productOutput := range out.Items {
			if price, ok := itemPriceMap[productOutput.ItemId]; ok {
				productOutput.ProductUnitPriceMin = price
			}
		}
	}

	return out, err
}

// Find 查询数据
func (s *sProductIndex) Find(ctx context.Context, in *do.ProductIndexListInput) (out []*entity.ProductIndex, err error) {
	out, err = dao.ProductIndex.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sProductIndex) List(ctx context.Context, in *do.ProductIndexListInput) (out *do.ProductIndexListOutput, err error) {
	out, err = dao.ProductIndex.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sProductIndex) Add(ctx context.Context, in *do.ProductIndex) (lastInsertId int64, err error) {
	lastInsertId, err = dao.ProductIndex.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sProductIndex) Edit(ctx context.Context, in *do.ProductIndex) (affected int64, err error) {
	_, err = dao.ProductIndex.Edit(ctx, in.ProductId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sProductIndex) Remove(ctx context.Context, id any) (affected int64, err error) {
	affected, err = dao.ProductIndex.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// ListItem 读取SKU商品
func (s *sProductIndex) ListItem(ctx context.Context, req *pt.ItemListReq) (out *model.ItemListOutput, err error) {
	output := &model.ItemListOutput{}

	if !g.IsEmpty(req.ActivityId) {
		activityBase, _ := dao.ActivityBase.Get(ctx, req.ActivityId)

		if activityBase != nil {
			req.ItemId = activityBase.ActivityItemIds
			output.ActivityBase = activityBase
		}
	}

	lists, err := service.ProductItem().ListItemKey(ctx, req)
	itemIds := lists.Items

	output.Records = lists.Records
	output.Size = lists.Size
	output.Page = lists.Page
	output.Total = lists.Total

	output.Items = []*model.ItemOutput{}

	if len(itemIds) > 0 {
		productItems, _ := dao.ProductItem.Gets(ctx, itemIds)
		uniqueProductIds := garray.NewArray()
		for _, item := range productItems {
			if !uniqueProductIds.Contains(item.ProductId) {
				uniqueProductIds.Append(item.ProductId)
			}
		}
		productIds := uniqueProductIds.Slice()

		// SKU图片
		productImages, _ := dao.ProductImage.Find(ctx, &do.ProductImageListInput{Where: do.ProductImage{
			ProductId: productIds,
		}})

		// 基础表数据
		productBases, _ := dao.ProductBase.Gets(ctx, productIds)

		// 产品状态
		productIndices, _ := dao.ProductIndex.Gets(ctx, productIds)
		productIndexMap := make(map[uint64]*entity.ProductIndex)
		for _, idx := range productIndices {
			productIndexMap[idx.ProductId] = idx
		}

		for _, vo := range productItems {
			var productImage *entity.ProductImage
			for _, img := range productImages {
				if img.ProductId == vo.ProductId && img.ColorId == vo.ColorId {
					productImage = img
					break
				}
			}

			if g.IsEmpty(productImage) {
				productImage = &entity.ProductImage{}
			}

			for _, base := range productBases {
				if vo.ProductId == base.ProductId {
					itemName := vo.ItemName
					itemSpecName := strings.ReplaceAll(itemName, ",", " ")

					productItemName := base.ProductName + " " + itemSpecName

					it := &model.ItemOutput{
						ProductItem:     *vo,
						ProductId:       vo.ProductId,
						ProductNumber:   base.ProductNumber,
						ItemSpecName:    productItemName,
						ProductName:     base.ProductName,
						ProductTips:     base.ProductTips,
						StoreId:         base.StoreId,
						ProductVideo:    base.ProductVideo,
						TransportTypeId: base.TransportTypeId,
						ProductBuyLimit: base.ProductBuyLimit,
						ProductImage:    base.ProductImage,
						ProductStateId:  productIndexMap[base.ProductId].ProductStateId,
					}

					if !g.IsEmpty(productImage.ItemImageDefault) {
						it.ProductImage = productImage.ItemImageDefault
					}

					// 商品状态
					if !g.IsEmpty(productIndexMap) {
						if !g.IsEmpty(productIndexMap[base.ProductId]) {
							it.ProductStateId = productIndexMap[base.ProductId].ProductStateId
						}
					}

					output.Items = append(output.Items, it)
				}
			}
		}
	}

	return output, err
}

// Detail 商品详情
func (s *sProductIndex) Detail(ctx context.Context, input *model.ProductDetailInput) (*model.ProductDetailOutput, error) {
	itemId := input.ItemId
	districtId := input.DistrictId
	//gbId := input.GbId

	out := &model.ProductDetailOutput{}
	out.ItemId = itemId

	productItem, err := dao.ProductItem.Get(ctx, itemId)
	if err != nil {
		return nil, err
	}

	if productItem == nil {
		return nil, gerror.New("商品SKU不存在!")
	}

	out.ItemRow = productItem

	// 设置销售价
	productItem.ItemSalePrice = productItem.ItemUnitPrice

	// 读取活动信息
	activityInfoVoList, err := service.ActivityItem().GetActivityInfo(ctx, []uint64{itemId})
	if err != nil {
		return nil, err
	}
	if len(activityInfoVoList) > 0 {
		activityInfoVo := activityInfoVoList[0]
		productItem.ActivityId = activityInfoVo.ActivityId
		productItem.ActivityInfo = activityInfoVo

		if g.IsEmpty(activityInfoVo.ActivityItemPrice) {
			productItem.ItemSalePrice = activityInfoVo.ActivityItemPrice
		}
	}

	productId := productItem.ProductId

	productIndex, err := dao.ProductIndex.Get(ctx, productId)
	if err != nil {
		return nil, err
	}
	productBase, err := dao.ProductBase.Get(ctx, productId)
	if err != nil {
		return nil, err
	}
	productInfo, err := dao.ProductInfo.Get(ctx, productId)
	if err != nil {
		return nil, err
	}

	gconv.Struct(productIndex, out)
	gconv.Struct(productBase, out)
	gconv.Struct(productInfo, out)

	// SKU图片

	image, err := dao.ProductImage.FindOne(ctx, &do.ProductImageListInput{
		Where: do.ProductImage{
			ProductId: productId,
			ColorId:   productItem.ColorId,
		},
	})
	if err != nil {
		return nil, err
	}
	out.Image = image
	out.ProductImage = image.ItemImageDefault

	itemName := productItem.ItemName
	itemSpecName := gstr.Replace(itemName, ",", " ")

	productItemName := fmt.Sprintf("%s %s", productBase.ProductName, itemSpecName)
	out.ProductItemName = productItemName

	// 是否可销售
	if productItem.AvailableQuantity > 0 {
		out.IfStore = true

		// 可售区域
		if !g.IsEmpty(districtId) {
			// 商品也，默认三级分类
			if !g.IsEmpty(districtId) {
				// 读取上级分类信息
				districtBase, err := dao.DistrictBase.Get(ctx, districtId)
				if err != nil {
					return nil, err
				}
				if districtBase != nil {
					districtId = districtBase.DistrictParentId
				}
			}

			storeTransportItemVo, err := service.StoreTransportType().GetFreight(ctx, productBase.TransportTypeId, districtId)
			if err != nil {
				return nil, err
			}

			if storeTransportItemVo == nil {
				out.IfStore = false
			} else {
				if !storeTransportItemVo.TransportTypeFree && storeTransportItemVo.Item == nil {
					out.IfStore = false
				}
				if transportItem := storeTransportItemVo.Item; transportItem != nil {
					out.Freight = transportItem.TransportItemDefaultPrice
				}
			}
		}
	} else {
		out.IfStore = false
	}

	categoryId := out.CategoryId

	if !g.IsEmpty(categoryId) {
		// 读取上级分类信息
		parentCategoryListById, err := service.ProductCategory().GetParentCategory(ctx, categoryId)
		if err != nil {
			return nil, err
		}
		out.ProductCategorys = parentCategoryListById
	}

	// 服务
	contractTypes, err := dao.ContractType.Find(ctx, &do.ContractTypeListInput{
		Where: do.ContractType{
			ContractTypeEnable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	//空数据返回 []
	if len(contractTypes) == 0 {
		out.Contracts = make([]*entity.ContractType, 0)
	} else {
		out.Contracts = contractTypes
	}

	// 商品评论
	commentInput := &do.ProductCommentListInput{
		Where: do.ProductComment{
			ProductId:     productId,
			CommentEnable: true,
		},
	}
	commentInput.Sidx = dao.ProductComment.Columns().CommentId
	commentInput.Sort = ml.ORDER_BY_DESC
	commentInput.Size = 5

	productCommentList, err := dao.ProductComment.Find(ctx, commentInput)
	if err != nil {
		return nil, err
	}

	userIds := make([]uint, 0)
	for _, comment := range productCommentList {
		userIds = append(userIds, comment.UserId)
	}
	userInfos, err := dao.UserInfo.Gets(ctx, userIds)
	if err != nil {
		return nil, err
	}

	for _, comment := range productCommentList {
		if g.IsEmpty(comment.CommentContent) {
			comment.CommentContent = "无评论"
		}

		for _, userInfo := range userInfos {
			if comment.UserId == userInfo.UserId {
				comment.UserAvatar = userInfo.UserAvatar
				break
			}
		}
		commentImage := comment.CommentImage
		if commentImage != "" {
			replaceImg := gstr.Replace(commentImage, "[", "", -1)
			replaceImg = gstr.Replace(replaceImg, "]", "", -1)
			comment.CommentImages = gconv.SliceStr(replaceImg)
		}
	}

	if len(productCommentList) > 0 {
		out.LastComments = productCommentList
		out.LastComment = productCommentList[0]
	}

	if !g.IsEmpty(input.UserId) {
		// 添加浏览记录
		service.UserProductBrowse().AddBrowser(ctx, input.ItemId, input.UserId)

		// 是否收藏
		count, _ := dao.UserFavoritesItem.Count(ctx, &do.UserFavoritesItemListInput{
			Where: do.UserFavoritesItem{
				UserId: input.UserId,
				ItemId: input.ItemId,
			},
		})

		if count > 0 {
			out.IsFavorite = true
		}
	}

	return out, nil
}
