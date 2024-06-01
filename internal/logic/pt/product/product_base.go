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
	"encoding/json"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"golershop.cn/utility/array"
	"sort"
)

type sProductBase struct{}

func init() {
	service.RegisterProductBase(NewProductBase())
}

func NewProductBase() *sProductBase {
	return &sProductBase{}
}

// Find 查询数据
func (s *sProductBase) Find(ctx context.Context, in *do.ProductBaseListInput) (out []*entity.ProductBase, err error) {
	out, err = dao.ProductBase.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sProductBase) List(ctx context.Context, in *do.ProductBaseListInput) (out *do.ProductBaseListOutput, err error) {
	out, err = dao.ProductBase.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sProductBase) Add(ctx context.Context, in *do.ProductBase) (lastInsertId int64, err error) {
	lastInsertId, err = dao.ProductBase.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sProductBase) Edit(ctx context.Context, in *do.ProductBase) (affected int64, err error) {
	_, err = dao.ProductBase.Edit(ctx, in.ProductId, in)
	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sProductBase) Remove(ctx context.Context, id any) (affected int64, err error) {
	affected, err = dao.ProductBase.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// SaveProdcut 添加或者编辑商品
func (s *sProductBase) SaveProdcut(ctx context.Context, in *model.SaveProductInput) (productId uint64, err error) {
	//todo 过滤替换词汇

	//todo 是否需要审核

	//判断新增 or 修改
	if g.IsEmpty(in.ProductBase.ProductId) {
		productId, err = service.NumberSeq().GetNextSeqInt(ctx, "product_id")

		if err != nil {
			return 0, err
		}

		//初始化默认值
		if g.IsEmpty(in.ProductIndex.ProductStateId) {
			in.ProductIndex.ProductStateId = consts.PRODUCT_STATE_NORMAL
		}

		if in.ProductIndex.ProductStateId == consts.PRODUCT_STATE_OFF_THE_SHELF {
			//in.ProductIndex.ProductSaleTime
		}

		in.ProductIndex.ProductSpEnable = 0   //供应商是否允许批发市场分销
		in.ProductIndex.ProductDistEnable = 1 //是否允许三级分销
		in.ProductIndex.ProductAddTime = gtime.Timestamp()
		in.ProductIndex.ProductFrom = 1000

	} else {
		productId = gconv.Uint64(in.ProductBase.ProductId)
	}

	//默认商品设置判断
	isSetDefault := false //设置了默认SKU

	for _, v := range in.ProductItems {
		if true == v.ItemIsDefault {
			isSetDefault = true
			break
		}
	}

	//如果未设置， 将第SKU设置为默认
	if !isSetDefault {
		in.ProductItems[0].ItemIsDefault = true
	}

	//处理主图
	productImage := ""

	//根据默认商品获取默认主图
	for _, v := range in.ProductItems {
		if true == v.ItemIsDefault {
			isSetDefault = true
			for _, image := range in.ProductImages {
				if image.ColorId == v.ColorId {
					if g.IsEmpty(image.ItemImageDefault) {
						image.ItemImageDefault = service.ConfigBase().GetDefaultImage(ctx)
					}

					productImage = gconv.String(image.ItemImageDefault)
					break
				}
			}

			break
		}
	}

	//商品价格最大值及最小值
	var productUnitPriceMin float64 = -1
	var productUnitPriceMax float64 = 0

	for _, v := range in.ProductItems {
		if productUnitPriceMin == -1 {
			productUnitPriceMin = v.ItemUnitPrice.(float64)
		}

		if v.ItemUnitPrice.(float64) < productUnitPriceMin {
			productUnitPriceMin = v.ItemUnitPrice.(float64)
		}

		if v.ItemUnitPrice.(float64) > productUnitPriceMax {
			productUnitPriceMax = v.ItemUnitPrice.(float64)
		}
	}

	in.ProductIndex.ProductUnitPriceMin = productUnitPriceMin
	in.ProductIndex.ProductUnitPriceMax = productUnitPriceMax

	//初始化商品状态

	//开启事务
	err = dao.ProductBase.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//商品基础表ProductBase
		in.ProductBase.ProductId = productId
		in.ProductBase.ProductImage = productImage
		_, err = dao.ProductBase.Save(ctx, in.ProductBase)
		if err != nil {
			return err
		}

		//商品辅助属性 Product_AssistIndexModel
		//处理辅助属性
		var productAssistMap map[string][]uint

		if err := json.Unmarshal([]byte(gconv.String(in.ProductInfo.ProductAssist)), &productAssistMap); err != nil {
			return err
		}

		var productAssistData []uint

		for assistId, assistItemIdList := range productAssistMap {
			g.Log().Info(ctx, assistId, assistItemIdList)
			productAssistData = append(productAssistData, assistItemIdList...)
		}

		//商品索引表
		//product_assist_data unit_price_min/max
		in.ProductIndex.ProductId = productId
		in.ProductIndex.ProductVerifyId = 3001
		in.ProductIndex.ProductNameIndex = in.ProductIndex.ProductName

		//商品状态 product_state_id 商品状态判断修正： 是否需要审核等

		productCategory, err := dao.ProductCategory.Get(ctx, in.ProductIndex.CategoryId)
		in.ProductIndex.TypeId = productCategory.TypeId
		in.ProductIndex.ProductAssistData = gstr.JoinAny(productAssistData, ",")
		_, err = dao.ProductIndex.Save(ctx, in.ProductIndex)
		if err != nil {
			return err
		}

		//商品信息表

		//处理product_uniqid
		var productUniqid = make(map[string][]interface{})

		for _, v := range in.ProductItems {
			var specItemIds []uint

			var specs []model.ISpecVo
			err := json.Unmarshal([]byte(v.ItemSpec.(string)), &specs)

			if err != nil {
				panic(err)
			}

			for _, spec := range specs {
				specItemIds = append(specItemIds, spec.Item.Id)
			}

			// Sort the slice in ascending order
			sort.Slice(specItemIds, func(i, j int) bool {
				return specItemIds[i] < specItemIds[j]
			})

			var colorImage string

			for _, image := range in.ProductImages {
				if image.ColorId == v.ColorId {
					colorImage = gconv.String(image.ItemImageDefault)
					break
				}
			}

			//[]interface{}{"ItemId", "item_unit_price", "ItemEnable", "item_quantity", "color_id", "color_img"}
			productUniqid[gstr.JoinAny(specItemIds, "-")] = []interface{}{v.ItemId, v.ItemUnitPrice, v.ItemQuantity, v.ItemEnable, v.ColorId, colorImage}
		}

		in.ProductInfo.ProductId = productId
		in.ProductInfo.ProductUniqid = productUniqid

		var productSpecList []struct{ Id uint }

		if err := json.Unmarshal([]byte(gconv.String(in.ProductInfo.ProductSpec)), &productSpecList); err != nil {
			return err
		}
		column := array.Column(productSpecList, "id")

		in.ProductInfo.SpecIds = gstr.JoinAny(column, ",")

		_, err = dao.ProductInfo.Save(ctx, in.ProductInfo)
		if err != nil {
			return err
		}

		//商品SKU Product_ItemModel
		//读取已经存在的SKU, 需要删除的记录
		itemIds, err := dao.ProductItem.FindKey(ctx, &do.ProductItemListInput{Where: do.ProductItem{ProductId: productId}})

		for _, v := range in.ProductItems {
			v.ProductId = in.ProductBase.ProductId
			v.CategoryId = in.ProductIndex.CategoryId

			if array.InArray(itemIds, v.ItemId) {
				itemIds = array.DeleteSlice(itemIds, v.ItemId).([]interface{})
			}
		}

		if !g.IsEmpty(itemIds) {
			dao.ProductItem.Remove(ctx, itemIds)
		}

		_, err = dao.ProductItem.Saves(ctx, in.ProductItems)
		if err != nil {
			return err
		}

		//商品图片 图片和规格属性一起保存  !如果是编辑，允许增加SKU，需要删除不使用的记录
		for _, v := range in.ProductImages {
			v.ProductId = in.ProductBase.ProductId
		}

		_, err = dao.ProductImage.Saves(ctx, in.ProductImages)
		if err != nil {
			return err
		}

		//Product_ValidPeriodModel
		if in.ProductIndex.KindId == consts.PRODUCT_KIND_FUWU {
			_, err = dao.ProductValidPeriod.Save(ctx, in.ProductValidPeriod)

			if err != nil {
				return err
			}
		}

		return nil
	})

	return productId, err
}

// RemoveProdcut 删除商品
func (s *sProductBase) RemoveProdcut(ctx context.Context, id any) (affected int64, err error) {

	affected, err = dao.ProductInfo.Remove(ctx, id)

	affected, err = dao.ProductIndex.Remove(ctx, id)

	var ProductImage = &do.ProductImageListInput{}
	ProductImage.Where.ProductId = id
	imageId, err := dao.ProductImage.FindFields(ctx, dao.ProductImage.Columns().ProductImageId, ProductImage)
	affected, err = dao.ProductImage.Remove(ctx, imageId)

	var ProductItem = &do.ProductItemListInput{}
	ProductItem.Where.ProductId = id
	itemId, err := dao.ProductItem.FindFields(ctx, dao.ProductItem.Columns().ItemId, ProductItem)
	affected, err = dao.ProductItem.Remove(ctx, itemId)

	affected, err = dao.ProductBase.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	_, err = dao.ProductValidPeriod.Remove(ctx, id)

	return affected, err
}

// GetProduct 读取商品信息
func (s *sProductBase) GetProduct(ctx context.Context, id any) (productData model.ProductDateOutput, err error) {
	//基础表
	productBase, err := dao.ProductBase.Get(ctx, id)

	if err != nil {
		return productData, err
	}

	productData.ProductBase = productBase

	//索引表
	productIndex, err := dao.ProductIndex.Get(ctx, id)

	if err != nil {
		return productData, err
	}

	productData.ProductIndex = productIndex

	//信息表
	productInfo, err := dao.ProductInfo.Get(ctx, id)

	if err != nil {
		return productData, err
	}

	productData.ProductInfo = productInfo

	//SKU表
	in := do.ProductItemListInput{}
	in.Where.ProductId = id

	productItems, err := dao.ProductItem.Find(ctx, &in)

	if err != nil {
		return productData, err
	}

	productData.ProductItem = productItems

	//图片表
	inImage := do.ProductImageListInput{}
	inImage.Where.ProductId = id
	productImages, err := dao.ProductImage.Find(ctx, &inImage)

	if err != nil {
		return productData, err
	}

	productData.ProductImage = productImages

	return productData, err
}

func (s *sProductBase) GetItems(ctx context.Context, itemIds []uint64, userId uint) (out []*model.ProductItemVo, err error) {
	// 获取商品SKU基本信息
	var output []*model.ProductItemVo

	if len(itemIds) > 0 {
		//todo 参与活动信息，格式化活动数据
		activityInfoList, _ := service.ActivityItem().GetActivityInfo(ctx, itemIds)

		itemList, _ := dao.ProductItem.Gets(ctx, itemIds)

		for _, productItem := range itemList {
			itVo := &model.ProductItemVo{
				CartSelect:    true,
				IsOos:         false,
				IsOnSale:      true,
				ItemVoucher:   0,
				ItemReduction: 0,

				PulseGiftCart:     make([]interface{}, 0),
				PulseReduction:    make([]interface{}, 0),
				PulseMultple:      make([]interface{}, 0),
				PulseBargainsCart: make([]interface{}, 0),
				PulseBargains:     make([]interface{}, 0),
			}

			gconv.Struct(productItem, itVo)
			itVo.ItemSalePrice = itVo.ItemUnitPrice

			// 1、如果存在活动，则优先采用活动价、如果没有活动，判断等级折扣价
			//用户等级折扣
			var userLevelRate float64 = 100

			if userId != 0 {
				userInfo, _ := dao.UserInfo.Get(ctx, userId)

				userLevelRateMap := service.UserLevel().GetUserLevelRateMap(ctx)
				userLevelRate = userLevelRateMap[userInfo.UserLevelId]
			}

			//是否有活动信息
			var activityInfoVo *model.ActivityInfoVo
			for _, info := range activityInfoList {
				if info.ItemId == itVo.ItemId {
					activityInfoVo = info
					break
				}
			}

			if activityInfoVo != nil && s.checkSingleActivity(ctx, activityInfoVo.ActivityTypeId) {
				itVo.ActivityInfo = activityInfoVo

				//判断是否执行活动信息
				if true {
					itVo.ActivityId = activityInfoVo.ActivityId
				}

				// 阶梯价， 需要数量计算单价
				if activityInfoVo.ActivityTypeId == consts.ACTIVITY_TYPE_BATDISCOUNT || activityInfoVo.ActivityTypeId == consts.ACTIVITY_TYPE_MULTIPLEDISCOUNT {
					//用户等级判断
					if userLevelRate != 100 {
						itVo.ItemSalePrice = itVo.ItemUnitPrice * float64(userLevelRate) / 100
						itVo.ItemSavePrice = itVo.ItemUnitPrice - itVo.ItemSalePrice
						itVo.ItemDiscountAmount = itVo.ItemSavePrice * float64(itVo.CartQuantity)
					}
				} else {
					itVo.ItemSalePrice = activityInfoVo.ActivityItemPrice
					itVo.ItemSavePrice = itVo.ItemUnitPrice - activityInfoVo.ActivityItemPrice
					itVo.ItemDiscountAmount = itVo.ItemSavePrice * float64(itVo.CartQuantity)
				}
			} else {
				//用户等级判断
				if userLevelRate != 100 {
					itVo.ItemSalePrice = itVo.ItemUnitPrice * float64(userLevelRate) / 100
					itVo.ItemSavePrice = itVo.ItemUnitPrice - itVo.ItemSalePrice
					itVo.ItemDiscountAmount = itVo.ItemSavePrice * float64(itVo.CartQuantity)
				}
			}

			output = append(output, itVo)
		}
	}

	var productIds []uint64
	for _, itemVo := range output {
		productIds = append(productIds, itemVo.ProductId)
	}

	if len(productIds) > 0 {
		productBases, _ := dao.ProductBase.Gets(ctx, productIds)
		productIndexs, _ := dao.ProductIndex.Gets(ctx, productIds)
		productImages, _ := dao.ProductImage.Find(ctx, &do.ProductImageListInput{Where: do.ProductImage{ProductId: productIds}})

		for _, itemVo := range output {
			productId := itemVo.ProductId
			var productName, productTips string
			var productCommissionRate float64
			var transportTypeId uint
			var productTags string
			var productDistEnable bool
			var productStateId, productInventoryLock, kindId uint

			for _, productBase := range productBases {
				if productBase.ProductId == productId {
					productName = productBase.ProductName
					productTips = productBase.ProductTips
					productCommissionRate = productBase.ProductCommissionRate
					transportTypeId = productBase.TransportTypeId
					break
				}
			}

			for _, productIndex := range productIndexs {
				if productIndex.ProductId == productId {
					productTags = productIndex.ProductTags
					productDistEnable = productIndex.ProductDistEnable
					productStateId = productIndex.ProductStateId
					productInventoryLock = productIndex.ProductInventoryLock
					kindId = productIndex.KindId
					break
				}
			}

			itemVo.ProductName = productName
			itemVo.ProductTips = productTips
			itemVo.ProductCommissionRate = productCommissionRate
			itemVo.TransportTypeId = transportTypeId
			itemVo.ProductTags = productTags
			itemVo.ProductDistEnable = productDistEnable
			itemVo.ProductStateId = productStateId
			itemVo.ProductInventoryLock = productInventoryLock
			itemVo.KindId = kindId

			for _, productImage := range productImages {
				if productImage.ProductId == productId && productImage.ColorId == itemVo.ColorId {
					if productImage.ItemImageDefault != "" {
						itemVo.ProductImage = productImage.ItemImageDefault
					}
					break
				}
			}
		}
	}

	return output, err
}

/**
 * 单品直接购买活动，判断是否通过活动
 *
 * @param activityTypeId
 * @return
 */
func (s *sProductBase) checkSingleActivity(ctx context.Context, activityTypeId uint) bool {

	if activityTypeId == consts.ACTIVITY_TYPE_GROUPBOOKING {
		activityId := gconv.Uint(g.RequestFromCtx(ctx).Get("activity_id"))

		return activityId != 0
	}

	return true
}
