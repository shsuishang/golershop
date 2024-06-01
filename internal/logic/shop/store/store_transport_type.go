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

package store

import (
	"context"
	"errors"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/model/entity"
	"golershop.cn/internal/service"
	"math"
)

type sStoreTransportType struct{}

func init() {
	service.RegisterStoreTransportType(NewStoreTransportType())
}

func NewStoreTransportType() *sStoreTransportType {
	return &sStoreTransportType{}
}

// Find 查询数据
func (s *sStoreTransportType) Find(ctx context.Context, in *do.StoreTransportTypeListInput) (out []*entity.StoreTransportType, err error) {
	out, err = dao.StoreTransportType.Find(ctx, in)

	return out, err
}

// List 分页读取
func (s *sStoreTransportType) List(ctx context.Context, in *do.StoreTransportTypeListInput) (out *do.StoreTransportTypeListOutput, err error) {
	out, err = dao.StoreTransportType.List(ctx, in)

	return out, err
}

// Add 新增
func (s *sStoreTransportType) Add(ctx context.Context, in *do.StoreTransportType) (lastInsertId int64, err error) {
	lastInsertId, err = dao.StoreTransportType.Add(ctx, in)
	if err != nil {
		return 0, err
	}
	return lastInsertId, err
}

// Edit 编辑
func (s *sStoreTransportType) Edit(ctx context.Context, in *do.StoreTransportType) (affected int64, err error) {
	_, err = dao.StoreTransportType.Edit(ctx, in.TransportTypeId, in)

	if err != nil {
		return 0, err
	}
	return
}

// Remove 删除多条记录模式
func (s *sStoreTransportType) Remove(ctx context.Context, id any) (affected int64, err error) {
	affected, err = dao.StoreTransportType.Remove(ctx, id)

	if err != nil {
		return 0, err
	}

	return affected, err
}

// GetFreight 获取配送区域信息及运费
func (s *sStoreTransportType) GetFreight(ctx context.Context, transportTypeId, districtId uint) (*model.StoreTransportItemVo, error) {
	var storeTransportItemVo *model.StoreTransportItemVo

	if transportTypeId > 0 {
		storeTransportType, err := dao.StoreTransportType.Get(ctx, transportTypeId)
		if err != nil {
			return nil, err
		}

		if storeTransportType != nil {
			storeTransportItemVo = &model.StoreTransportItemVo{
				StoreTransportType: *storeTransportType,
			}

			// 全部免运费，任何地区都配送
			transportTypeFree := storeTransportType.TransportTypeFree

			if transportTypeFree {
				storeTransportItem := &entity.StoreTransportItem{
					TransportItemDefaultPrice: 0,
				}
				storeTransportItemVo.Item = storeTransportItem
			} else {
				itemQuery := &do.StoreTransportItemListInput{
					Where: do.StoreTransportItem{TransportTypeId: transportTypeId},
				}

				if districtId > 0 {
					itemQuery.WhereExt = append(itemQuery.WhereExt, &ml.WhereExt{
						Column: dao.StoreTransportItem.Columns().TransportItemCityIds,
						Val:    []uint64{uint64(districtId)},
						Symbol: ml.FIND_IN_SET,
					})
				}

				storeTransportItem, err := dao.StoreTransportItem.FindOne(ctx, itemQuery)
				if err != nil {
					return nil, err
				}

				storeTransportItemVo.Item = storeTransportItem
			}
		}
	}

	return storeTransportItemVo, nil
}

// CalFreight 运费计算
func (s *sStoreTransportType) CalFreight(ctx context.Context, transportTypeId, districtId, quantity uint, orderTotal, postFreeMax float64) (*model.OrderFreightVo, error) {
	data := &model.OrderFreightVo{
		CanDelivery:    true,
		FreightFreeMin: postFreeMax,
		Freight:        0,
	}

	if transportTypeId == 0 {
		return nil, errors.New("transportTypeId is wrong")
	}

	// 获取配送区域及运费
	storeTransportItemVo, err := s.GetFreight(ctx, transportTypeId, districtId)
	if err != nil {
		return nil, err
	}

	if storeTransportItemVo != nil {
		item := storeTransportItemVo.Item

		// 可售区域
		if !storeTransportItemVo.TransportTypeFree && item == nil {
			data.CanDelivery = false
		}

		if item != nil {
			transportTypeFreightFree := storeTransportItemVo.TransportTypeFreightFree
			postFreeMax = math.Max(postFreeMax, transportTypeFreightFree)
			data.FreightFreeMin = postFreeMax

			if transportTypeFreightFree > 0 && orderTotal >= transportTypeFreightFree {
				// 订单免运费
				return data, nil
			}

			transportItemDefaultNum := item.TransportItemDefaultNum
			transportItemAddNum := item.TransportItemAddNum

			addNum := math.Max(0, float64(quantity-transportItemDefaultNum))
			transportItemDefaultPrice := item.TransportItemDefaultPrice
			if addNum > 0 && transportItemAddNum > 0 {
				transportItemAddPrice := item.TransportItemAddPrice
				sum := transportItemDefaultPrice + (transportItemAddPrice * (addNum / float64(transportItemAddNum)))

				data.Freight = sum
			} else {
				// 默认运费
				data.Freight = transportItemDefaultPrice
			}

			return data, nil
		}
	} else {
		data.Freight = 0
	}

	return data, nil
}
