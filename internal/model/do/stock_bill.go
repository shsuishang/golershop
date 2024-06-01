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

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// StockBill is the golang structure of table invoicing_stock_bill for DAO operations like Where/Data.
type StockBill struct {
	g.Meta               `orm:"table:invoicing_stock_bill, do:true"`
	StockBillId          interface{} // 购货(退货)单编号
	StockBillChecked     interface{} // 是否审核(BOOL):1-已审核;  0-未审核
	StockBillDate        *gtime.Time // 单据日期
	StockBillModifyTime  *gtime.Time // 更新时间
	StockBillTime        interface{} // 创建时间
	BillTypeId           interface{} // 业务类别purchase_type_id, sale_type_id(ENUM):2750-入库;2700-出库;2855-采购订单;2850-销售订单
	StockTransportTypeId interface{} // 库存类型(ENUM)
	StoreId              interface{} // 所属店铺
	WarehouseId          interface{} // 所属仓库
	OrderId              interface{} // 源单号码:一个订单一个出入库记录可以拆单
	StockBillRemark      interface{} // 备注
	EmployeeId           interface{} // 经办人
	AdminId              interface{} // 制单人
	StockBillOtherMoney  interface{} // 其它金额
	StockBillAmount      interface{} // 单据金额
	StockBillEnable      interface{} // 是否有效(BOOL):1-有效; 0-无效
	StockBillSrcId       interface{} // 关联编号
}

type StockBillListInput struct {
	ml.BaseList
	Where StockBill // 查询条件
}

type StockBillListOutput struct {
	Items   []*entity.StockBill // 列表
	Page    int                 // 分页号码
	Total   int                 // 总页数
	Records int                 // 数据总数
	Size    int                 // 单页数量
}

type StockBillListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
