package shop

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/api/shop"
	"golershop.cn/internal/dao"
	"golershop.cn/internal/model/do"
	"golershop.cn/internal/service"
)

var (
	StoreTransportItem = cStoreTransportItem{}
)

type cStoreTransportItem struct{}

// =================== 管理端使用 =========================

func (c *cStoreTransportItem) List(ctx context.Context, req *shop.StoreTransportItemListReq) (res *shop.StoreTransportItemListRes, err error) {
	var likes []*ml.WhereExt
	item := do.StoreTransportItem{TransportTypeId: req.TransportTypeId}

	if item.TransportTypeId == 0 {
		item.TransportTypeId = nil
	}

	result, err := service.StoreTransportItem().List(ctx, &do.StoreTransportItemListInput{
		BaseList: ml.BaseList{
			Page:      req.Page,
			Size:      req.Size,
			WhereLike: likes,
			Sidx:      dao.StoreTransportItem.Columns().TransportItemId,
			Sort:      "ASC",
		},
		Where: item,
	})

	if err != nil {
		return nil, err
	}

	/*
		for _, item := range result.Items {
			// 获取城市ID列表
			var cityIdList []int
			cityIds := strings.Split(item.TransportItemCityIds, ",")

			for _, cityIdStr := range cityIds {
				cityId, _ := strconv.Atoi(cityIdStr)
				cityIdList = append(cityIdList, cityId)
			}

			//城市列表
			var cityNames []map[string]interface{}
			if item.TransportItemCityName == "" {
				cityNames = []map[string]interface{}{}
			} else {
				err = json.Unmarshal([]byte(item.TransportItemCityName), &cityNames)
				if err != nil {
					return nil, err
				}
			}

			var cityNameList []map[string]interface{}
			if len(cityNames) == 0 {
				cityNameList = []map[string]interface{}{}
			} else {
				for _, city := range cityNames {
					provinces, _ := city["provinces"].(string)
					citys, _ := city["citys"].([]interface{})

					cityMap := map[string]interface{}{
						"provinces": provinces,
						"citys":     citys,
					}
					cityNameList = append(cityNameList, cityMap)
				}
				//item.TransportItemCityIdlist = cityIdList
				//item.TransportItemCityNameList = cityNameList
			}
		}
	*/

	res = &shop.StoreTransportItemListRes{
		Items:   result.Items,
		Page:    result.Page,
		Records: result.Records,
		Total:   result.Total,
		Size:    result.Size,
	}
	return res, nil
}

// Add 新增菜单
func (c *cStoreTransportItem) Add(ctx context.Context, req *shop.StoreTransportItemAddReq) (res *shop.StoreTransportItemEditRes, err error) {

	input := do.StoreTransportItem{}
	gconv.Scan(req, &input)

	var result, error = service.StoreTransportItem().Add(ctx, &input)
	//var result, error = service.StoreTransportItem().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &shop.StoreTransportItemEditRes{
		TransportItemId: result,
	}

	return
}

// Edit 编辑菜单
func (c *cStoreTransportItem) Edit(ctx context.Context, req *shop.StoreTransportItemEditReq) (res *shop.StoreTransportItemEditRes, err error) {

	input := do.StoreTransportItem{}
	gconv.Scan(req, &input)

	var result, error = service.StoreTransportItem().Edit(ctx, &input)
	//var result, error = service.StoreTransportItem().Edit(ctx, req)

	if error != nil {
		err = error
	}

	res = &shop.StoreTransportItemEditRes{
		TransportItemId: result,
	}

	return
}

// Remove 删除菜单
func (c *cStoreTransportItem) Remove(ctx context.Context, req *shop.StoreTransportItemRemoveReq) (res *shop.StoreTransportItemRemoveRes, err error) {

	var _, error = service.StoreTransportItem().Remove(ctx, req.TransportItemId)

	/*
		input := do.StoreTransportItem{}
		input.StoreTransportItemTime = gtime.Now()
		input.StoreTransportItemId = req.StoreTransportItemId[0]
		input.StoreTransportItemSort = 0

		var _, error = service.StoreTransportItem().Edit(ctx, &input)
	*/

	if error != nil {
		err = error
	}

	res = &shop.StoreTransportItemRemoveRes{}

	return
}
