package model

import "golershop.cn/internal/model/entity"

// StoreItemVo 店铺及商品信息结构体
type StoreItemVo struct {
	StoreId            uint                 `json:"store_id"`             // 店铺编号
	StoreName          string               `json:"store_name"`           // 店铺名称
	Items              []*ProductItemVo     `json:"items"`                // 商品信息
	Activitys          ActivitysVo          `json:"activitys"`            // 店铺活动
	ActivityBase       *entity.ActivityBase `json:"activity_base"`        // 过程中的非排他活动
	RedemptionItems    []*ActivitysVo       `json:"redemption_items"`     // 提货券
	VoucherItems       []*UserVoucherRes    `json:"voucher_items"`        // 优惠券
	ProductAmount      float64              `json:"product_amount"`       // 商品原价总价
	FreightAmount      float64              `json:"freight_amount"`       // 运费总价
	FreightFreeBalance float64              `json:"freight_free_balance"` // 还差N免运费
	DiscountAmount     float64              `json:"discount_amount"`      // 优惠总额度
	MoneyItemAmount    float64              `json:"money_item_amount"`    // 单品优惠总价
	MoneyAmount        float64              `json:"money_amount"`         // 商品最终总价
	PointsAmount       float64              `json:"points_amount"`        // 需要总积分
	SpAmount           float64              `json:"sp_amount"`            // 需要总积分2
	UserVoucherId      uint                 `json:"user_voucher_id"`      // 优惠券编号
	VoucherAmount      float64              `json:"voucher_amount"`       // 代金券
	KindId             uint                 `json:"kind_id"`              // 订单类型
	IsVirtual          bool                 `json:"is_virtual"`           // 是否虚拟
}
