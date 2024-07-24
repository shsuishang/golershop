package shop

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model"
	"golershop.cn/internal/model/entity"
)

// UserVoucherListReq 用户优惠券表分页查询
type UserVoucherListReq struct {
	g.Meta `path:"/front/shop/userVoucher/list" tags:"用户优惠券" method:"get" summary:"用户优惠券表分页查询"`
	ml.BaseList

	model.UserVoucherRes
}

type UserVoucherListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量                                                                                                                                                                                                                            // 优惠券是否生效
}

type GetVoucherNumReq struct {
	g.Meta `path:"/front/shop/userVoucher/getEachVoucherNum" tags:"优惠券" method:"get" summary:"自定义不同优惠券数量"`

	VoucherStateId uint `form:"voucher_state_id"  dc:"优惠券状态"`
}

type GetVoucherNumRes struct {
	VoucherAllNum       int `json:"voucher_all_num"      dc:"所有优惠券"`  // 所有优惠券
	VoucherOfflinedNum  int `json:"voucher_offlined_num" dc:"线下优惠券"`  // 线下优惠券
	VoucherOnlinedNum   int `json:"voucher_onlined_num"  dc:"线上优惠券"`  // 线上优惠券
	VoucherCarcouponNum int `json:"voucher_carcoupon_num" dc:"附加优惠券"` // 附加优惠券
	VoucherUnusedNum    int `json:"voucher_unused_num"   dc:"未使用优惠券"` // 未使用优惠券
	VoucherUsedNum      int `json:"voucher_used_num"     dc:"已使用优惠券"` // 已使用优惠券
	VoucherTimeoutNum   int `json:"voucher_timeout_num"  dc:"已过期优惠券"` // 已过期优惠券
}
type UserVoucherAddReq struct {
	g.Meta `path:"/front/shop/userVoucher/add" tags:"领取代金券" method:"post" summary:"领取代金券"`

	ActivityId uint `form:"activity_id"  dc:"活动id"`
}
type UserVoucherAddRes struct {
	*entity.UserVoucher
}
