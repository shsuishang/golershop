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

// StockBill is the golang structure for table stock_bill.
type StockBill struct {
	StockBillId          string      `json:"stock_bill_id"           ` // 购货(退货)单编号
	StockBillChecked     bool        `json:"stock_bill_checked"      ` // 是否审核(BOOL):1-已审核;  0-未审核
	StockBillDate        *gtime.Time `json:"stock_bill_date"         ` // 单据日期
	StockBillModifyTime  *gtime.Time `json:"stock_bill_modify_time"  ` // 更新时间
	StockBillTime        uint64      `json:"stock_bill_time"         ` // 创建时间
	BillTypeId           uint        `json:"bill_type_id"            ` // 业务类别purchase_type_id, sale_type_id(ENUM):2750-入库;2700-出库;2855-采购订单;2850-销售订单
	StockTransportTypeId uint        `json:"stock_transport_type_id" ` // 库存类型(ENUM)
	StoreId              uint        `json:"store_id"                ` // 所属店铺
	WarehouseId          uint        `json:"warehouse_id"            ` // 所属仓库
	OrderId              string      `json:"order_id"                ` // 源单号码:一个订单一个出入库记录可以拆单
	StockBillRemark      string      `json:"stock_bill_remark"       ` // 备注
	EmployeeId           uint        `json:"employee_id"             ` // 经办人
	AdminId              uint        `json:"admin_id"                ` // 制单人
	StockBillOtherMoney  float64     `json:"stock_bill_other_money"  ` // 其它金额
	StockBillAmount      float64     `json:"stock_bill_amount"       ` // 单据金额
	StockBillEnable      bool        `json:"stock_bill_enable"       ` // 是否有效(BOOL):1-有效; 0-无效
	StockBillSrcId       string      `json:"stock_bill_src_id"       ` // 关联编号
}
