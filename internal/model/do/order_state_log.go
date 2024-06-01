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

// OrderStateLog is the golang structure of table trade_order_state_log for DAO operations like Where/Data.
type OrderStateLog struct {
	g.Meta           `orm:"table:trade_order_state_log, do:true"`
	OrderStateLogId  interface{} // 状态编号
	OrderId          interface{} // 订单编号
	OrderStateId     interface{} // 订单状态:2010-待付款;2020-待配货;2030-待发货;2040-已发货;2050-已签收;2060-已完成;2070-已取消;
	OrderStatePreId  interface{} // 订单状态:2010-待付款;2020-待配货;2030-待发货;2040-已发货;2050-已签收;2060-已完成;2070-已取消;
	OrderStateType   interface{} // 操作类别
	OrderStateTime   *gtime.Time // 操作时间
	UserId           interface{} // 操作用户
	UserAccount      interface{} // 操作账号
	OrderStateNote   interface{} // 操作备注
	OrderStateIsSync interface{} // 是否同步(BOOL):0-否;1-是
}

type OrderStateLogListInput struct {
	ml.BaseList
	Where OrderStateLog // 查询条件
}

type OrderStateLogListOutput struct {
	Items   []*entity.OrderStateLog // 列表
	Page    int                     // 分页号码
	Total   int                     // 总页数
	Records int                     // 数据总数
	Size    int                     // 单页数量
}

type OrderStateLogListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
