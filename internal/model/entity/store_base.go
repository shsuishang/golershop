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

// StoreBase is the golang structure for table store_base.
type StoreBase struct {
	StoreId                uint        `json:"store_id"                  ` // 店铺编号
	StoreName              string      `json:"store_name"                ` // 店铺名称
	StoreGradeId           uint        `json:"store_grade_id"            ` // 店铺等级
	StoreLogo              string      `json:"store_logo"                ` // 店铺logo
	StoreLatitude          float64     `json:"store_latitude"            ` // 纬度
	StoreLongitude         float64     `json:"store_longitude"           ` // 经度
	StoreDeliverDistrictId string      `json:"store_deliver_district_id" ` // 配送区域(DOT)
	StoreIsSelfsupport     bool        `json:"store_is_selfsupport"      ` // 是否自营(ENUM): 1-自营;0-非自营
	StoreType              uint        `json:"store_type"                ` // 店铺类型(ENUM): 1-卖家店铺; 2-供应商店铺
	StoreIsOpen            bool        `json:"store_is_open"             ` // 店铺状态(BOOL):0-关闭;  1-运营中
	StoreCategoryId        uint        `json:"store_category_id"         ` // 店铺分类编号
	StoreO2OTags           string      `json:"store_o_2_o_tags"          ` // 免费服务(DOT)
	StoreO2OFlag           uint        `json:"store_o_2_o_flag"          ` // 是否O2O(BOOL):0-否;1-是
	StoreCircle            string      `json:"store_circle"              ` // 所属商圈(DOT)
	SubsiteId              uint        `json:"subsite_id"                ` // 所属分站:0-总站
	CreateTime             *gtime.Time `json:"create_time"               ` // 加入时间
}
