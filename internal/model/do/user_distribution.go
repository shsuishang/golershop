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

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// UserDistribution is the golang structure of table account_user_distribution for DAO operations like Where/Data.
type UserDistribution struct {
	g.Meta             `orm:"table:account_user_distribution, do:true"`
	UserId             interface{} // 用户编号
	UserParentId       interface{} // 上级用户编号
	UserPartnerId      interface{} // 所属城市合伙人
	UserTeamCount      interface{} // 团队人数
	UserProvinceTeamId interface{} // 所属省公司
	UserCityTeamId     interface{} // 所属市公司
	UserCountyTeamId   interface{} // 所属区公司
	RoleLevelId        interface{} // 角色等级
	UccId              interface{} // 渠道编号
	ActivityId         interface{} // 活动编号
	UserTime           interface{} // 注册时间
	UserFansNum        interface{} // 粉丝数量:冗余
	UserIsSp           interface{} // 服务商(BOOL):0-否;1-是;
	UserIsDa           interface{} // 区代理(BOOL):0-否;1-是为区Id;
	UserIsCa           interface{} // 市代理(BOOL):0-否;1-是为市Id;
	UserIsPa           interface{} // 省代理(BOOL):0-否;1-是为省Id;
	UserIsPt           interface{} // 城市合伙人(BOOL):0-否;1-是;
	UserActive         interface{} // 是否有效(BOOL):0-未生效;1-有效
	UserVoucherIds     interface{} // 分销优惠券
}

type UserDistributionListInput struct {
	ml.BaseList
	Where UserDistribution // 查询条件
}

type UserDistributionListOutput struct {
	Items   []*entity.UserDistribution // 列表
	Page    int                        // 分页号码
	Total   int                        // 总页数
	Records int                        // 数据总数
	Size    int                        // 单页数量
}

type UserDistributionListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
