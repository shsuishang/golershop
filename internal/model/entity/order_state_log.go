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

// OrderStateLog is the golang structure for table order_state_log.
type OrderStateLog struct {
	OrderStateLogId  uint        `json:"order_state_log_id"  ` // 状态编号
	OrderId          string      `json:"order_id"            ` // 订单编号
	OrderStateId     uint        `json:"order_state_id"      ` // 订单状态:2010-待付款;2020-待配货;2030-待发货;2040-已发货;2050-已签收;2060-已完成;2070-已取消;
	OrderStatePreId  uint        `json:"order_state_pre_id"  ` // 订单状态:2010-待付款;2020-待配货;2030-待发货;2040-已发货;2050-已签收;2060-已完成;2070-已取消;
	OrderStateType   string      `json:"order_state_type"    ` // 操作类别
	OrderStateTime   *gtime.Time `json:"order_state_time"    ` // 操作时间
	UserId           uint        `json:"user_id"             ` // 操作用户
	UserAccount      string      `json:"user_account"        ` // 操作账号
	OrderStateNote   string      `json:"order_state_note"    ` // 操作备注
	OrderStateIsSync bool        `json:"order_state_is_sync" ` // 是否同步(BOOL):0-否;1-是
}
