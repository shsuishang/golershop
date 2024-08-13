package model

import "golershop.cn/internal/model/entity"

// StoreTransportItemVo 店铺运输项目视图对象
type StoreTransportItemVo struct {
	entity.StoreTransportType
	Item *entity.StoreTransportItem `json:"item,omitempty"`
}

type StoreInfoVo struct {
	entity.StoreInfo
	StoreId                int     `json:"store_id" dc:"店铺编号"`                               // 店铺编号
	StoreName              string  `json:"store_name" dc:"店铺名称"`                             // 店铺名称
	StoreGradeId           int     `json:"store_grade_id" dc:"店铺等级"`                         // 店铺等级
	StoreLogo              string  `json:"store_logo" dc:"店铺logo"`                           // 店铺logo
	StoreLatitude          float64 `json:"store_latitude" dc:"纬度"`                           // 纬度
	StoreLongitude         float64 `json:"store_longitude" dc:"经度"`                          // 经度
	StoreDeliverDistrictId string  `json:"store_deliver_district_id" dc:"配送区域(DOT)"`         // 配送区域(DOT)
	StoreIsSelfsupport     bool    `json:"store_is_selfsupport" dc:"是否自营(ENUM): 1-自营;0-非自营"` // 是否自营(ENUM): 1-自营;0-非自营
	StoreType              int     `json:"store_type" dc:"店铺类型(ENUM): 1-卖家店铺; 2-供应商店铺"`      // 店铺类型(ENUM): 1-卖家店铺; 2-供应商店铺
	StoreIsOpen            bool    `json:"store_is_open" dc:"店铺状态(BOOL):0-关闭;  1-运营中"`       // 店铺状态(BOOL):0-关闭;  1-运营中
	StoreCategoryId        int     `json:"store_category_id" dc:"店铺分类编号"`                    // 店铺分类编号
	StoreO2oTags           string  `json:"store_o2o_tags" dc:"免费服务(DOT)"`                    // 免费服务(DOT)
	StoreO2oFlag           bool    `json:"store_o2o_flag" dc:"是否O2O(BOOL):0-否;1-是"`          // 是否O2O(BOOL):0-否;1-是
	StoreCircle            string  `json:"store_circle" dc:"所属商圈(DOT)"`                      // 所属商圈(DOT)
	SubsiteId              int     `json:"subsite_id" dc:"所属分站:0-总站"`                        // 所属分站:0-总站
}
