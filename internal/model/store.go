package model

import "golershop.cn/internal/model/entity"

// StoreTransportItemVo 店铺运输项目视图对象
type StoreTransportItemVo struct {
	entity.StoreTransportType
	Item *entity.StoreTransportItem `json:"item,omitempty"`
}
