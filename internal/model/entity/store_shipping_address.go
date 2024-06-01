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

// StoreShippingAddress is the golang structure for table store_shipping_address.
type StoreShippingAddress struct {
	SsId         uint        `json:"ss_id"          ` // 地址编号
	SsName       string      `json:"ss_name"        ` // 联系人
	SsIntl       string      `json:"ss_intl"        ` // 国家编码
	SsMobile     string      `json:"ss_mobile"      ` // 手机号码
	SsTelephone  string      `json:"ss_telephone"   ` // 联系电话
	SsContacter  string      `json:"ss_contacter"   ` // 联系人(未启用)
	SsPostalcode string      `json:"ss_postalcode"  ` // 邮编
	SsProvinceId uint        `json:"ss_province_id" ` // 省编号
	SsProvince   string      `json:"ss_province"    ` // 省份
	SsCityId     uint        `json:"ss_city_id"     ` // 市编号
	SsCity       string      `json:"ss_city"        ` // 市
	SsCountyId   uint        `json:"ss_county_id"   ` // 县
	SsCounty     string      `json:"ss_county"      ` // 县区
	SsAddress    string      `json:"ss_address"     ` // 详细地址:不必重复填写地区
	SsTime       *gtime.Time `json:"ss_time"        ` // 添加时间
	SsIsDefault  bool        `json:"ss_is_default"  ` // 默认地址(ENUM):0-否;1-是
	StoreId      uint        `json:"store_id"       ` // 所属店铺
}
