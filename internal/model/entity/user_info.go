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

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	UserId               uint        `json:"user_id"                ` // 用户编号
	UserAccount          string      `json:"user_account"           ` // 用户账号
	UserNickname         string      `json:"user_nickname"          ` // 用户昵称
	UserAvatar           string      `json:"user_avatar"            ` // 用户头像
	UserState            uint        `json:"user_state"             ` // 状态(ENUM):0-锁定;1-已激活;2-未激活;
	UserMobile           string      `json:"user_mobile"            ` // 手机号码(mobile)
	UserIntl             string      `json:"user_intl"              ` // 国家编码
	UserGender           uint        `json:"user_gender"            ` // 性别(ENUM):0-保密;1-男;  2-女;
	UserBirthday         *gtime.Time `json:"user_birthday"          ` // 生日(DATE)
	UserEmail            string      `json:"user_email"             ` // 用户邮箱(email)
	UserLevelId          uint        `json:"user_level_id"          ` // 等级编号
	UserRealname         string      `json:"user_realname"          ` // 真实姓名
	UserIdcard           string      `json:"user_idcard"            ` // 身份证
	UserIdcardImages     string      `json:"user_idcard_images"     ` // 身份证图片(DTO)
	UserIsAuthentication uint        `json:"user_is_authentication" ` // 认证状态(ENUM):0-未认证;1-待审核;2-认证通过;3-认证失败
	TagIds               string      `json:"tag_ids"                ` // 用户标签(DOT)
	UserFrom             uint        `json:"user_from"              ` // 用户来源(ENUM):2310-其它;2311-pc;2312-H5;2313-APP;2314-小程序;2315-公众号
	UserNew              bool        `json:"user_new"               ` // 新人标识(BOOL):0-不是;1-是
}
