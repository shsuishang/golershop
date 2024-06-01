package trade

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
)

// start fo front

// start fo manage
type OrderInvoiceAdd struct {
	OrderInvoiceId        int         `json:"order_invoice_id"        ` // 发票编号
	OrderId               string      `json:"order_id"                ` // 订单编号
	UserId                uint        `json:"user_id"                 ` // 所属用户
	StoreId               int         `json:"store_id"                ` // 店铺编号
	InvoiceTitle          string      `json:"invoice_title"           ` // 发票抬头
	InvoiceContent        string      `json:"invoice_content"         ` // 发票内容
	InvoiceAmount         float64     `json:"invoice_amount"          ` // 开票金额
	InvoiceCompanyCode    string      `json:"invoice_company_code"    ` // 纳税人识别号
	InvoiceIsCompany      bool        `json:"invoice_is_company"      ` // 公司开票(BOOL):0-个人;1-公司
	InvoiceIsElectronic   bool        `json:"invoice_is_electronic"   ` // 电子发票(ENUM):0-纸质发票;1-电子发票
	InvoiceType           int         `json:"invoice_type"            ` // 发票类型(ENUM):1-普通发票;2-增值税专用发票
	InvoiceStatus         bool        `json:"invoice_status"          ` // 开票状态(BOOL): 0-未开票; 1-已开票;
	InvoiceDatetime       *gtime.Time `json:"invoice_datetime"        ` // 开票时间
	InvoiceAddress        string      `json:"invoice_address"         ` // 单位地址
	InvoicePhone          string      `json:"invoice_phone"           ` // 单位电话
	InvoiceBankname       string      `json:"invoice_bankname"        ` // 开户银行
	InvoiceBankaccount    string      `json:"invoice_bankaccount"     ` // 银行账号
	InvoiceContactName    string      `json:"invoice_contact_name"    ` // 收票人
	InvoiceContactArea    string      `json:"invoice_contact_area"    ` // 收票人地区
	InvoiceContactAddress string      `json:"invoice_contact_address" ` // 收票详细地址
	UserMobile            string      `json:"user_mobile"             ` // 手机号码(mobile)
	UserEmail             string      `json:"user_email"              ` // 用户邮箱(email)
}

type OrderInvoiceAddReq struct {
	g.Meta `path:"/manage/trade/orderInvoice/add" tags:"发票管理" method:"post" summary:"发票管理编辑接口"`

	OrderInvoiceAdd
}

type OrderInvoiceEditReq struct {
	g.Meta `path:"/manage/trade/orderInvoice/edit" tags:"发票管理" method:"post" summary:"发票管理编辑接口"`

	OrderInvoiceId uint `json:"order_invoice_id"  dc:"发票管理编号"   ` // 发票管理编号
	OrderInvoiceAdd
}

type OrderInvoiceEditRes struct {
	OrderInvoiceId uint `json:"order_invoice_id" dc:"发票管理编号"   ` // 发票管理编号
}

type OrderInvoiceRemoveReq struct {
	g.Meta         `path:"/manage/trade/orderInvoice/remove" tags:"发票管理" method:"post" summary:"发票管理删除接口"`
	OrderInvoiceId uint `json:"order_invoice_id" dc:"发票管理编号"   ` // 发票管理编号
}

type OrderInvoiceRemoveRes struct {
}

type OrderInvoiceListReq struct {
	g.Meta `path:"/manage/trade/orderInvoice/list" tags:"发票管理" method:"get" summary:"发票管理列表接口"`
	ml.BaseList

	UserId        uint `json:"user_id"                   `        // 用户编号
	InvoiceStatus uint `json:"invoice_status"                   ` // 状态

	OrderId      string `json:"order_id"   type:"LIKE"       ` // 订单编号
	InvoiceTitle string `json:"invoice_title"  type:"LIKE"   ` // 发票抬头
}

type OrderInvoiceListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
