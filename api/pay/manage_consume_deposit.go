package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

type ConsumeDepositOfflinePayReq struct {
	g.Meta `path:"/manage/pay/consumeDeposit/offline" tags:"交易单" method:"post" summary:"线下支付接口"`

	OrderId            string      `json:"order_id" v:"required#请输入订单编号"  dc:"订单编号"` //商城支付编号
	PaymentChannelId   uint        `json:"payment_channel_id"  d:"1422"`             //支付渠道
	DepositPaymentType uint        `json:"deposit_payment_type"  d:"1305"`           //支付方式(ENUM):1301-货到付款; 1302-在线支付; 1303-白条支付; 1304-现金支付; 1305-线下支付;
	DepositTradeNo     string      `json:"deposit_trade_no"  v:"required#请输入交易凭证号" ` // 交易号:支付宝etc
	DepositNotifyTime  *gtime.Time `json:"deposit_notify_time" `                     //时间
	DepositTotalFee    float64     `json:"deposit_total_fee"           `             // 交易金额

	PmMoney        float64 `json:"pm_money" `         //余额
	PmRechargeCard float64 `json:"pm_recharge_card" ` //充值卡
	PmPoints       float64 `json:"pm_points" `        //积分
	PmCredit       float64 `json:"pm_credit" `        //信用账户
	PmRedpack      float64 `json:"pm_redpack" `       //红包账户
}

type ConsumeDepositOfflinePayRes struct {
	OrderId string `json:"order_id"  dc:"订单编号"`
}

type ConsumeDepositListReq struct {
	g.Meta `path:"/manage/pay/consumeDeposit/list" tags:"交易单" method:"get" summary:"交易单列表接口"`
	ml.BaseList

	UserId           uint   `json:"user_id"    `                             // 用户编号
	DepositSubject   string `json:"deposit_subject"  type:"LIKE"           ` // 商品名称
	ConsumeDepositId uint   `json:"deposit_id"            `                  // 交易订单编号
}

type ConsumeDepositListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ConsumeDepositEditReviewReq struct {
	g.Meta `path:"/manage/pay/consumeDeposit/editReview" tags:"支付表-收款确认" method:"post" summary:"支付表-收款确认"`
	ml.BaseList

	ConsumeDepositId uint `json:"deposit_id"            ` // 交易订单编号
	DepositReview    bool `json:"deposit_review"        ` // 收款确认(BOOL):0-未确认;1-已确认
}

type ConsumeDepositEditReviewRes struct {
	ConsumeDepositId uint `json:"deposit_id"            ` // 交易订单编号
}
