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

// ActivityCutprice is the golang structure of table marketing_activity_cutprice for DAO operations like Where/Data.
type ActivityCutprice struct {
	g.Meta          `orm:"table:marketing_activity_cutprice, do:true"`
	AcId            interface{} // 砍价编号
	ActivityId      interface{} // 砍价商品编号
	UserId          interface{} // 用户编号
	AcSalePrice     interface{} // 砍价后价格
	AcDatetime      interface{} // 创建时间
	OrderId         interface{} // 订单号
	AcMixLimitPrice interface{} // 砍价底价
	AcNum           interface{} // 砍价人数
	Version         interface{} // 版本
}

type ActivityCutpriceListInput struct {
	ml.BaseList
	Where ActivityCutprice // 查询条件
}

type ActivityCutpriceListOutput struct {
	Items   []*entity.ActivityCutprice // 列表
	Page    int                        // 分页号码
	Total   int                        // 总页数
	Records int                        // 数据总数
	Size    int                        // 单页数量
}

type ActivityCutpriceListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
