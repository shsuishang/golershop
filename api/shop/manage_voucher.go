package shop

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

type UserVoucherReq struct {
	g.Meta `path:"/manage/shop/userVoucher/list" tags:"用户优惠券" method:"get" summary:"用户优惠券表分页查询"`
	ml.BaseList

	UserVoucherId         int         `json:"user_voucher_id" dc:"代金券编号"`                                        // 代金券编号
	ActivityId            int         `json:"activity_id" dc:"代金券模版编号"`                                          // 代金券模版编号
	VoucherStateId        int         `json:"voucher_state_id" dc:"代金券状态(ENUM):1501-未用;1502-已用;1503-过期;1504-收回"` // 代金券状态(ENUM)
	UserVoucherTime       *gtime.Time `json:"user_voucher_time" dc:"代金券发放日期"`                                    // 代金券发放日期
	UserId                int         `json:"user_id" dc:"所属用户"`                                                 // 所属用户
	OrderId               string      `json:"order_id" dc:"订单编号"`                                                // 订单编号
	UserVoucherActivetime *gtime.Time `json:"user_voucher_activetime" dc:"使用时间"`                                 // 使用时间
	VoucherPrice          float64     `json:"voucher_price" dc:"优惠券可抵扣价格"`                                       // 优惠券可抵扣价格
	VoucherSubtotal       float64     `json:"voucher_subtotal" dc:"使用优惠券的订单金额"`                                  // 使用优惠券的订单金额
	VoucherEndDate        *gtime.Time `json:"voucher_end_date" dc:"失效日期"`                                        // 失效日期
	StoreId               int         `json:"store_id" dc:"所属店铺编号"`                                              // 所属店铺编号
	ItemId                string      `json:"item_id" dc:"单品优惠商品编号(DOT)"`                                        // 单品优惠商品编号(DOT)
	VoucherType           int         `json:"voucher_type" dc:"优惠券类型(ENUM): 0-普通优惠券;1-免拼券"`                      // 优惠券类型(ENUM)
	VoucherStartDate      *gtime.Time `json:"voucher_start_date" dc:"到期使用时间"`                                    // 到期使用时间
	WriteoffCode          string      `json:"writeoff_code" dc:"线下活动提货码"`                                        // 线下活动提货码
	VoucherUserWay        int         `json:"voucher_user_way" dc:"使用方式"`                                        // 使用方式
	VoucherEffect         bool        `json:"voucher_effect" dc:"优惠券是否生效(BOOL): false-未生效;true-生效"`              // 优惠券是否生效(BOOL)
}

type UserVoucherRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
