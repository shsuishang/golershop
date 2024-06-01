// +----------------------------------------------------------------------
// | ShopSuite商城系统 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 随商信息技术（上海）有限公司
// +----------------------------------------------------------------------
// | 未获商业授权前，不得将本软件用于商业用途。禁止整体或任何部分基础上以发展任何派生版本、
// | 修改版本或第三方版本用于重新分发。
// +----------------------------------------------------------------------
// | 官方网站: https://www.shopsuite.cn  https://www.golershop.cn
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本公司对该软件产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细见https://www.golershop.cn/policy
// +----------------------------------------------------------------------

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInvoice is the golang structure for table user_invoice.
type UserInvoice struct {
	UserInvoiceId         int         `json:"user_invoice_id"         ` // 发票编号
	UserId                uint        `json:"user_id"                 ` // 所属用户
	InvoiceTitle          string      `json:"invoice_title"           ` // 发票抬头
	InvoiceCompanyCode    string      `json:"invoice_company_code"    ` // 纳税人识别号
	InvoiceContent        string      `json:"invoice_content"         ` // 发票内容
	InvoiceIsCompany      bool        `json:"invoice_is_company"      ` // 公司开票(BOOL):0-个人;1-公司
	InvoiceIsElectronic   bool        `json:"invoice_is_electronic"   ` // 电子发票(ENUM):0-纸质发票;1-电子发票
	InvoiceType           uint        `json:"invoice_type"            ` // 发票类型(ENUM):1-普通发票;2-增值税专用发票
	InvoiceDatetime       *gtime.Time `json:"invoice_datetime"        ` // 添加时间
	InvoiceAddress        string      `json:"invoice_address"         ` // 单位地址
	InvoicePhone          string      `json:"invoice_phone"           ` // 单位电话
	InvoiceBankname       string      `json:"invoice_bankname"        ` // 开户银行
	InvoiceBankaccount    string      `json:"invoice_bankaccount"     ` // 银行账号
	InvoiceContactMobile  string      `json:"invoice_contact_mobile"  ` // 收票人手机
	InvoiceContactEmail   string      `json:"invoice_contact_email"   ` // 收票人邮箱
	InvoiceIsDefault      int         `json:"invoice_is_default"      ` // 是否默认
	InvoiceContactName    string      `json:"invoice_contact_name"    ` // 收票人
	InvoiceContactArea    string      `json:"invoice_contact_area"    ` // 收票人地区
	InvoiceContactAddress string      `json:"invoice_contact_address" ` // 收票详细地址
}
