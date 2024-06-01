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

// AccessHistory is the golang structure for table access_history.
type AccessHistory struct {
	AccessId             int64       `json:"access_id"              ` // 编号
	UserId               uint        `json:"user_id"                ` // 用户编号
	ItemId               uint64      `json:"item_id"                ` // SKU编号
	AccessClientId       string      `json:"access_client_id"       ` // 唯一客户编号
	AccessOs             string      `json:"access_os"              ` // 操作系统
	AccessBrowserName    string      `json:"access_browser_name"    ` // 浏览器名称
	AccessBrowserVersion string      `json:"access_browser_version" ` // 浏览器版本
	AccessSpider         string      `json:"access_spider"          ` // 搜索引擎
	AccessCountry        string      `json:"access_country"         ` // 国家
	AccessProvince       string      `json:"access_province"        ` // 省份
	AccessCity           string      `json:"access_city"            ` // 市
	AccessCounty         string      `json:"access_county"          ` // 区
	AccessLang           string      `json:"access_lang"            ` // 语言
	AccessIp             string      `json:"access_ip"              ` // 访问IP
	AccessMethod         string      `json:"access_method"          ` // 请求方法
	AccessUrl            string      `json:"access_url"             ` // 访问地址
	AccessTime           uint64      `json:"access_time"            ` // 访问时间
	AccessYear           uint        `json:"access_year"            ` // 年
	AccessMonth          uint        `json:"access_month"           ` // 月
	AccessDay            uint        `json:"access_day"             ` // 日
	AccessHour           uint        `json:"access_hour"            ` // 时
	AccessDate           *gtime.Time `json:"access_date"            ` // 年月日
	AccessDatetime       *gtime.Time `json:"access_datetime"        ` // 时间
	AccessReferDomain    string      `json:"access_refer_domain"    ` // 来源
	AccessReferUrl       string      `json:"access_refer_url"       ` // 来源
	AccessMobile         bool        `json:"access_mobile"          ` // 是否手机
	AccessPad            bool        `json:"access_pad"             ` // 是否平板
	AccessPc             bool        `json:"access_pc"              ` // 是否PC
	AccessDevice         uint        `json:"access_device"          ` // 终端(ENUM):1-Phone;2-Pad;3-Pc
	AccessType           uint        `json:"access_type"            ` // 终端来源(ENUM):2310-其它;2311-pc;2312-H5;2313-APP;2314-小程序
	AccessFrom           uint        `json:"access_from"            ` // 终端来源(ENUM):2320-其它;2321-微信;2322-百度;2323-支付宝;2324-头条
	AccessData           string      `json:"access_data"            ` // 请求数据
}
