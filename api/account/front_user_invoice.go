package account

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// 用户发票
type UserInvoiceAdd struct {
	UserId                uint   `json:"user_id"                 ` // 所属用户
	InvoiceTitle          string `json:"invoice_title"           ` // 发票抬头
	InvoiceCompanyCode    string `json:"invoice_company_code"    ` // 纳税人识别号
	InvoiceContent        string `json:"invoice_content"         ` // 发票内容
	InvoiceIsCompany      bool   `json:"invoice_is_company"      ` // 公司开票(BOOL):0-个人;1-公司
	InvoiceIsElectronic   bool   `json:"invoice_is_electronic"   ` // 电子发票(ENUM):0-纸质发票;1-电子发票
	InvoiceType           uint   `json:"invoice_type"            ` // 发票类型(ENUM):1-普通发票;2-增值税专用发票
	InvoiceAddress        string `json:"invoice_address"         ` // 单位地址
	InvoicePhone          string `json:"invoice_phone"           ` // 单位电话
	InvoiceBankname       string `json:"invoice_bankname"        ` // 开户银行
	InvoiceBankaccount    string `json:"invoice_bankaccount"     ` // 银行账号
	InvoiceContactMobile  string `json:"invoice_contact_mobile"  ` // 收票人手机
	InvoiceContactEmail   string `json:"invoice_contact_email"   ` // 收票人邮箱
	InvoiceIsDefault      int    `json:"invoice_is_default"      ` // 是否默认
	InvoiceContactName    string `json:"invoice_contact_name"    ` // 收票人
	InvoiceContactArea    string `json:"invoice_contact_area"    ` // 收票人地区
	InvoiceContactAddress string `json:"invoice_contact_address" ` // 收票详细地址

}

type UserInvoiceAddReq struct {
	g.Meta `path:"/front/account/userInvoice/add" tags:"user" method:"post" summary:"用户发票接口"`

	UserInvoiceAdd
}

type UserInvoiceAddRes entity.UserInvoice

type UserInvoiceEditReq struct {
	g.Meta `path:"/front/account/userInvoice/edit" tags:"user" method:"post" summary:"用户发票接口"`

	UserInvoiceId int `json:"user_invoice_id"  v:"required#请输入发票编号"        ` // 发票编号
	UserInvoiceAdd
}

type UserInvoiceEditRes struct {
	UserInvoiceId int `json:"user_invoice_id"    ` // 发票编号
}

type UserInvoiceGetReq struct {
	g.Meta `path:"/front/account/userInvoice/get" tags:"用户发票" method:"get" summary:"用户发票接口"`

	UserInvoiceId int `json:"user_invoice_id"  v:"required#请输入发票编号"        ` // 发票编号
}

type UserInvoiceGetRes entity.UserInvoice

type UserInvoiceRemoveReq struct {
	g.Meta `path:"/front/account/userInvoice/remove" tags:"用户发票" method:"post" summary:"用户发票接口"`

	UserInvoiceId int `json:"user_invoice_id"  v:"required#请输入发票编号"        ` // 发票编号
}

type UserInvoiceRemoveRes struct {
}

type UserInvoiceListReq struct {
	g.Meta `path:"/front/account/userInvoice/list" tags:"用户发票" method:"get" summary:"发票列表接口"`
	ml.BaseList

	UdName string `json:"ud_name" type:"LIKE"        ` // 联系人
}

type UserInvoiceListRes struct {
	Items   interface{} `json:"items"    dc:"分页数据内容"`
	Page    int         `json:"page"`    // 分页号码
	Total   int         `json:"total"`   // 总页数
	Records int         `json:"records"` // 数据总数
	Size    int         `json:"size"`    // 单页数量
}
