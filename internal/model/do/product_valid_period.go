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
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// ProductValidPeriod is the golang structure of table pt_product_valid_period for DAO operations like Where/Data.
type ProductValidPeriod struct {
	g.Meta                      `orm:"table:pt_product_valid_period, do:true"`
	ProductId                   interface{} // 产品编号
	ProductValidPeriod          interface{} // 有效期:1001-长期有效;1002-自定义有效期;1003-购买起有效时长年单位
	ProductValidityStart        interface{} // 开始时间
	ProductValidityEnd          interface{} // 失效时间
	ProductValidityDuration     interface{} // 有效时长单位为天
	ProductValidType            interface{} // 服务类型(ENUM):1001-到店服务;1002-上门服务
	ProductServiceDateFlag      interface{} // 填写预约日期(BOOL):0-否;1-是
	ProductServiceContactorFlag interface{} // 填写联系人(BOOL):0-否;1-是
	ProductValidRefundFlag      interface{} // 支持过期退款(BOOL):0-否;1-是
}

type ProductValidPeriodListInput struct {
	ml.BaseList
	Where ProductValidPeriod // 查询条件
}

type ProductValidPeriodListOutput struct {
	Items   []*entity.ProductValidPeriod // 列表
	Page    int                          // 分页号码
	Total   int                          // 总页数
	Records int                          // 数据总数
	Size    int                          // 单页数量
}

type ProductValidPeriodListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
