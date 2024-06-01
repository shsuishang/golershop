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

// UserExpHistory is the golang structure for table user_exp_history.
type UserExpHistory struct {
	ExpLogId    uint        `json:"exp_log_id"    ` //
	ExpKindId   uint        `json:"exp_kind_id"   ` // 类型(ENUM):1-获取;2-消耗
	ExpTypeId   uint        `json:"exp_type_id"   ` // 积分类型(ENUM):1-会员注册;2-会员登录;3-商品评论;4-购买商品;5-管理员操作;7-积分换购商品;8-积分兑换代金券
	UserId      uint        `json:"user_id"       ` // 会员编号
	ExpLogValue uint        `json:"exp_log_value" ` // 可用经验
	UserExp     uint        `json:"user_exp"      ` // 当前经验
	ExpLogTime  *gtime.Time `json:"exp_log_time"  ` // 创建时间
	ExpLogDesc  string      `json:"exp_log_desc"  ` // 描述
	StoreId     uint        `json:"store_id"      ` // 所属店铺
	ExpLogDate  *gtime.Time `json:"exp_log_date"  ` // 获得经验的日期
}
