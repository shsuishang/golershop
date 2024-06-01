package model

import "golershop.cn/internal/model/entity"

// 退货申请对象
type OrderReturnInput struct {
	ReturnId           string                    `json:"return_id,omitempty"`            // 退单号
	OrderId            string                    `json:"order_id,omitempty"`             // 订单编号
	ReturnItems        []*OrderReturnItemInputVo `json:"return_items,omitempty"`         // 订单商品表编号
	ReturnBuyerMessage string                    `json:"return_buyer_message,omitempty"` // 买家退货备注
	ReturnTel          string                    `json:"return_tel,omitempty"`           // 联系电话
	ReturnReasonId     uint                      `json:"return_reason_id,omitempty"`     // 退款理由编号
	StoreId            uint                      `json:"store_id,omitempty"`             // 店铺编号
	ReturnItemImage    string                    `json:"return_item_image,omitempty"`    // 退款凭据(DOT)
	UserId             uint                      `json:"user_id,omitempty"`              // 用户编号
	ReturnAllFlag      bool                      `json:"return_all_flag,omitempty"`      // 退货标记(BOOL):true-全退， false-单品退
	ReturnFlag         uint                      `json:"return_flag,omitempty"`          // 退货类型(ENUM): 0-不用退货;1-需要退货
	ReviewFlag         bool                      `json:"review_flag,omitempty"`          // 自动审核(BOOL): 0-不审核;1-自动审核
}

// 退货申请对象
type OrderReturnItemInputVo struct {
	OrderItemId        uint64  `json:"order_item_id,omitempty"`        // 订单商品表编号
	ReturnRefundAmount float64 `json:"return_refund_amount,omitempty"` // 退款金额
	ReturnItemNum      uint    `json:"return_item_num,omitempty"`      // 退货商品数量
}

// OrderReturnRes 售后服务
type OrderReturnVo struct {
	entity.OrderReturn

	ReturnNum                uint                 `json:"return_num" description:"退货商品总数量"`               // 退货商品总数量
	SubmitReturnRefundAmount float64              `json:"submit_return_refund_amount" description:"退款金额"` // 退款金额
	ReturnReasonName         string               `json:"return_reason_name" description:"售后理由"`          // 售后理由
	Items                    []*OrderReturnItemVo `json:"items" description:"退货订单商品信息"`                   // 退货订单商品信息

	BuyerUserName string `json:"buyer_user_name" dc:"客户名称"` // 客户名称
	DaName        string `json:"da_name" dc:"联系人"`          // 联系人
	DaMobile      string `json:"da_mobile" dc:"手机号码"`       // 手机号码
	DaProvince    string `json:"da_province" dc:"省份"`       // 省份
	DaCity        string `json:"da_city" dc:"市"`            // 市
	DaCounty      string `json:"da_county" dc:"县区"`         // 县区
	DaAddress     string `json:"da_address" dc:"详细地址"`      // 详细地址
}

// OrderReturnItemVo 退货订单商品
type OrderReturnItemVo struct {
	OrderId             string                      `json:"order_id" description:"订单编号"`                   // 订单编号
	OrderItemId         uint64                      `json:"order_item_id" description:"订单item_id"`         // 订单item_id
	ProductItemName     string                      `json:"product_item_name" description:"商品名称"`          // 商品名称
	UnitName            string                      `json:"unit_name" description:"单位"`                    // 单位
	ItemUnitPrice       float64                     `json:"item_unit_price" description:"单价"`              // 单价
	ReturnItemImageList []string                    `json:"return_item_image_list" description:"退款凭据"`     // 退款凭据
	ReturnItemNum       uint                        `json:"return_item_num" description:"退货商品数量"`          // 退货商品数量
	ReturnItemSubtotal  float64                     `json:"return_item_subtotal" description:"退款总额"`       // 退款总额
	OrderItemImage      string                      `json:"order_item_image" description:"商品图片"`           // 商品图片
	OrderItemSalePrice  float64                     `json:"order_item_sale_price" description:"商品实际成交价单价"` // 商品实际成交价单价
	OrderItemQuantity   uint                        `json:"order_item_quantity" description:"商品数量"`        // 商品数量
	ItemId              uint64                      `json:"item_id" description:"货品编号"`                    // 货品编号
	ItemName            string                      `json:"item_name" description:"商品名称"`                  // 商品名称
	ProductId           uint64                      `json:"product_id" description:"产品编号"`                 // 产品编号
	ProductName         string                      `json:"product_name" description:"商品名称"`               // 商品名称
	CanRefundAmount     float64                     `json:"can_refund_amount" description:"可退金额"`          // 可退金额
	CanRefundNum        uint                        `json:"can_refund_num" description:"可退数量"`             // 可退数量
	ReturnReasonList    []*entity.OrderReturnReason `json:"return_reason_list" description:"退货原因集合"`       // 退货原因集合
}

type OrderReturnOutput struct {
	Items   []*OrderReturnVo `json:"items"    dc:"分页数据内容"`
	Page    int              `json:"page"`    // 分页号码
	Total   int              `json:"total"`   // 总页数
	Records int              `json:"records"` // 数据总数
	Size    int              `json:"size"`    // 单页数量
}
