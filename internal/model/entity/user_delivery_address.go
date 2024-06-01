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

// UserDeliveryAddress is the golang structure for table user_delivery_address.
type UserDeliveryAddress struct {
	UdId         uint        `json:"ud_id"          ` // 地址编号
	UserId       uint        `json:"user_id"        ` // 用户编号
	UdName       string      `json:"ud_name"        ` // 联系人
	UdIntl       string      `json:"ud_intl"        ` // 国家编码
	UdMobile     string      `json:"ud_mobile"      ` // 手机号码
	UdTelephone  string      `json:"ud_telephone"   ` // 联系电话
	UdProvinceId uint        `json:"ud_province_id" ` // 省编号
	UdProvince   string      `json:"ud_province"    ` // 省份
	UdCityId     uint        `json:"ud_city_id"     ` // 市编号
	UdCity       string      `json:"ud_city"        ` // 市
	UdCountyId   uint        `json:"ud_county_id"   ` // 县
	UdCounty     string      `json:"ud_county"      ` // 县区
	UdAddress    string      `json:"ud_address"     ` // 详细地址
	UdPostalcode string      `json:"ud_postalcode"  ` // 邮政编码
	UdTagName    string      `json:"ud_tag_name"    ` // 地址标签(ENUM):1001-家里;1002-公司
	UdLongitude  float64     `json:"ud_longitude"   ` // 经度
	UdLatitude   float64     `json:"ud_latitude"    ` // 纬读
	UdTime       *gtime.Time `json:"ud_time"        ` // 添加时间
	UdIsDefault  bool        `json:"ud_is_default"  ` // 是否默认(BOOL):0-非默认;1-默认
}
