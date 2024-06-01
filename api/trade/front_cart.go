package trade

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model"
)

// 购物车表
type UserCartListReq struct {
	g.Meta `path:"/front/trade/cart/list" tags:"购物车" method:"get" summary:"购物车列表"`

	ml.BaseList

	UserId     uint `json:"user_id"`     // 买家编号
	CartSelect bool `json:"cart_select"` // 选中状态(BOOL):0-未选;1-已选
	UdId       uint `json:"ud_id"`       // 地址编号 或者 地址数据Map
}

type UserCartListRes model.CheckoutOutput

// 购物车表参数
type UserCartAddReq struct {
	g.Meta       `path:"/front/trade/cart/add" tags:"购物车" method:"post" summary:"购物车添加"`
	ItemId       uint64 `json:"item_id" v:"required#请输入商品编号"`          // 商品编号
	CartQuantity uint   `json:"cart_quantity" v:"required#请输入商品数量"`    // 购买商品数量
	CartType     uint   `json:"cart_type" `                            // 购买类型(ENUM):1-购买; 2-积分兑换; 3-赠品; 4-活动促销
	ActivityId   uint   `json:"activity_id"  d:"0" v:"required#活动编号" ` // 活动Id-加价购等等加入购物的需要提示
}

type UserCartAddRes struct {
	CartId uint64 `json:"cart_id"` // 编号
}

// 购物车表参数
type UserCartAddBatReq struct {
	g.Meta       `path:"/front/trade/cart/addBatch" tags:"购物车" method:"post" summary:"购物车添加"`
	ItemId       uint64 `json:"item_id" v:"required#请输入商品编号"`          // 商品编号
	CartQuantity uint   `json:"cart_quantity" v:"required#请输入商品数量"`    // 购买商品数量
	CartType     uint   `json:"cart_type" `                            // 购买类型(ENUM):1-购买; 2-积分兑换; 3-赠品; 4-活动促销
	ActivityId   uint   `json:"activity_id" d:"0"  v:"required#活动编号" ` // 活动Id-加价购等等加入购物的需要提示
}

// 购物车表参数
type UserCartEditReq struct {
	g.Meta       `path:"/front/trade/cart/edit" tags:"购物车" method:"post" summary:"购物车修改"`
	CartId       uint64 `json:"cart_id" v:"required#请输入购物车编号"` // 编号
	CartQuantity uint   `json:"cart_quantity" `                // 购买商品数量
}

type UserCartEditRes struct {
	CartId uint64 `json:"cart_id"` // 编号
}

// 购物车表参数
type UserCartEditQuantityReq struct {
	g.Meta       `path:"/front/trade/cart/editQuantity" tags:"购物车" method:"post" summary:"购物车修改"`
	CartId       uint64 `json:"cart_id" v:"required#请输入购物车编号"` // 编号
	CartQuantity uint   `json:"cart_quantity" `                // 购买商品数量
}

// 购物车表参数
type UserCartRemoveReq struct {
	g.Meta `path:"/front/trade/cart/remove" tags:"购物车" method:"post" summary:"购物车删除"`
	CartId uint64 `json:"cart_id" v:"required#请输入购物车编号"` // 编号
}

type UserCartRemoveRes struct {
	CartId uint64 `json:"cart_id"` // 编号
}

type UserCartRemoveBatchReq struct {
	g.Meta  `path:"/front/trade/cart/removeBatch" tags:"购物车" method:"post" summary:"购物车删除"`
	CartIds string `p:"cart_id" v:"required#请输入购物车编号"` // 编号
}

type UserCartRemoveBatchRes struct {
	CartId uint64 `json:"cart_id"` // 编号
}

// 购物车表参数
type UserCartSelectReq struct {
	g.Meta     `path:"/front/trade/cart/sel" tags:"购物车" method:"post" summary:"选中"`
	Action     string `json:"action" `      // All:全部; store:店铺编号
	CartId     uint64 `json:"cart_id" `     // 编号
	CartSelect bool   `json:"cart_select" ` // 是否选中
}

type UserCartSelectRes struct {
	CartId uint64 `json:"cart_id"` // 编号
}

// OrderCheckoutReq 订单编辑接口
type OrderCheckoutReq struct {
	g.Meta         `path:"/front/trade/cart/checkout" tags:"购物车" method:"get" summary:"checkout"`
	UdId           uint   `json:"ud_id"         d:"0"   v:"required#收货地址编号"`                   // 收货地址编号
	CartId         string `json:"cart_id"       d:""   v:"required#下单商品数据:商品编号|数量,商品编号|数量..."` // 下单商品数据
	ChainId        uint   `json:"chain_id"      d:"0"   v:"required#门店编号"`                     // 门店编号
	ActivityId     uint   `json:"activity_id"    d:"0"  v:"required#活动编号"`                     // 活动编号
	DeliveryTypeId uint   `json:"delivery_type_id" d:"2" v:"required#配送方式"`                    // 配送方式
	ChannelType    uint   `json:"channel_type"   d:"0"   v:"required#来源渠道 0:正常下单;1:直播渠道"`      // 来源渠道
	UserVoucherIds string `json:"user_voucher_ids"`                                            // 优惠券
}

type CheckoutRes model.CheckoutOutput
