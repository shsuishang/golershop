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

package global

import (
	"github.com/gogf/gf/v2/frame/g"
	"golershop.cn/internal/consts"
	"golershop.cn/internal/model"
)

var (
	Registry              g.Map  = make(g.Map)
	StateIdRow            []uint //启用的订单状态
	StateIdAll            []uint //所有预设订单状态
	StateIdSelect         []*model.SelectVo
	ReturnStateSelectList []*model.SelectVo

	ReturnStateIdList           []uint //启用的订单状态
	ReturnStateIdAll            []uint //所有预设订单状态
	ReturnStateIdSelect         []*model.SelectVo
	ReturnReturnStateSelectList []*model.SelectVo

	ReturnProcessMap = map[uint]int{
		consts.RETURN_PROCESS_SUBMIT:               1,
		consts.RETURN_PROCESS_CHECK:                2,
		consts.RETURN_PROCESS_RECEIVED:             3,
		consts.RETURN_PROCESS_REFUND:               4,
		consts.RETURN_PROCESS_RECEIPT_CONFIRMATION: 5,
		consts.RETURN_PROCESS_FINISH:               6,
	}

	PaymentChannelSelect []*model.SelectVo //开启的支付方式
	PaymentChannelMap    map[uint]string   //

	UserLevelMap     map[uint]string  //
	UserLevelRateMap map[uint]float64 //

	Debug bool //是否启用调试
	Cache bool //是否启用缓存

	Namespace string //缓存命名空间
	Lk        string

	BaseUrl string // 访问域名
	UrlH5   string // 访问域名
	UrlPc   string // 访问域名

)
