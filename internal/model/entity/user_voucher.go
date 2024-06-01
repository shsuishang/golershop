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

// UserVoucher is the golang structure for table user_voucher.
type UserVoucher struct {
	UserVoucherId         uint        `json:"user_voucher_id"         ` // 代金券编号
	ActivityId            uint        `json:"activity_id"             ` // 代金券模版编号
	ActivityName          string      `json:"activity_name"           ` // 代金券名称
	VoucherStateId        uint        `json:"voucher_state_id"        ` // 代金券状态(ENUM):1501-未用;1502-已用;1503-过期;1504-收回
	UserVoucherTime       *gtime.Time `json:"user_voucher_time"       ` // 代金券发放日期
	UserId                uint        `json:"user_id"                 ` // 所属用户
	OrderId               string      `json:"order_id"                ` // 订单编号
	UserVoucherActivetime *gtime.Time `json:"user_voucher_activetime" ` // 使用时间
	VoucherPrice          float64     `json:"voucher_price"           ` // 优惠券可抵扣价格
	VoucherSubtotal       float64     `json:"voucher_subtotal"        ` // 使用优惠券的订单金额
	VoucherStartDate      int64       `json:"voucher_start_date"      ` // 生效时间
	VoucherEndDate        int64       `json:"voucher_end_date"        ` // 失效时间
	StoreId               uint        `json:"store_id"                ` // 所属店铺编号
	ItemId                string      `json:"item_id"                 ` // 单品优惠商品编号(DOT)
	VoucherType           uint        `json:"voucher_type"            ` // 优惠券类型(ENUM): 0-普通优惠券;1-免拼券
	WriteoffCode          string      `json:"writeoff_code"           ` // 线下活动提货码
	ActivityRule          string      `json:"activity_rule"           ` // 活动规则(JSON):不检索{rule_id:{}, rule_id:{}},统一解析规则{"requirement":{"buy":{"item":[1,2,3],"subtotal":"通过计算修正满足的条件"}},"rule":[{"total":100,"max_num":1,"item":{"1":1,"1200":3}},{"total":200,"max_num":1,"item":{"1":1,"1200":3}}]}
}
