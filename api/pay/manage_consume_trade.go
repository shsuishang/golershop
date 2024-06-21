package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model"
)

// start fo manage
type ConsumeTradeListReq struct {
	g.Meta `path:"/manage/pay/consumeTrade/list" tags:"交易单" method:"get" summary:"交易单列表接口"`
	ml.BaseList

	BuyerId        uint `json:"buyer_id"   `                  //买家编号
	ConsumeTradeId uint `json:"consume_trade_id"            ` // 交易订单编号
}

type ConsumeTradeListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}

type ConsumeTradePayReq struct {
	g.Meta `path:"/manage/pay/consumeTrade/pay" tags:"交易单" method:"post" summary:"支付接口"`

	OrderId string `json:"order_id" v:"required#请输入订单编号"  dc:"订单编号"` //商城支付编号

	model.PayMetVo
}

type ConsumeTradePayRes struct {
	OrderId string `json:"order_id"  dc:"订单编号"` //商城支付编号
}
