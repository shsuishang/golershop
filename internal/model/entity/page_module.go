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

// PageModule is the golang structure for table page_module.
type PageModule struct {
	PmId       uint        `json:"pm_id"       ` //
	PmName     string      `json:"pm_name"     ` // 模块名称
	PageId     int64       `json:"page_id"     ` //
	UserId     uint        `json:"user_id"     ` // 所属用户
	PmColor    string      `json:"pm_color"    ` // 颜色
	PmType     string      `json:"pm_type"     ` // 所在页面
	ModuleId   string      `json:"module_id"   ` // 模版
	PmUtime    *gtime.Time `json:"pm_utime"    ` // 更新时间
	PmOrder    uint        `json:"pm_order"    ` // 排序
	PmEnable   uint        `json:"pm_enable"   ` // 是否显示
	PmHtml     string      `json:"pm_html"     ` // 模块html代码
	PmJson     string      `json:"pm_json"     ` // 模块JSON代码(JSON)
	SubsiteId  uint        `json:"subsite_id"  ` // 所属分站Id:0-总站
	PmPosition string      `json:"pm_position" ` // column_left:content_top
}
