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
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mallsuite/gocore/core/ml"
	"golershop.cn/internal/model/entity"
)

// UserPointsHistory is the golang structure of table pay_user_points_history for DAO operations like Where/Data.
type UserPointsHistory struct {
	g.Meta          `orm:"table:pay_user_points_history, do:true"`
	PointsLogId     interface{} // 记录编号
	PointsKindId    interface{} // 类型(ENUM):1-获取积分;2-消费积分;
	PointsTypeId    interface{} // 积分类型(ENUM):1-会员注册;2-会员登录;3-商品评论;4-购买商品;5-管理员操作;7-积分换购商品;8-积分兑换代金券
	UserId          interface{} // 会员编号
	PointsLogPoints interface{} // 可用积分
	UserPoints      interface{} // 当前积分
	PointsLogTime   interface{} // 创建时间
	PointsLogDesc   interface{} // 描述
	StoreId         interface{} // 所属店铺
	PointsLogDate   *gtime.Time // 积分日期
	UserIdOther     interface{} // 交互会员
	PointsLogState  interface{} // 领取状态(BOOL):0-未领取;1-已领取
	ExtId           interface{} // 关联单号
}

type UserPointsHistoryListInput struct {
	ml.BaseList
	Where UserPointsHistory // 查询条件
}

type UserPointsHistoryListOutput struct {
	Items   []*entity.UserPointsHistory // 列表
	Page    int                         // 分页号码
	Total   int                         // 总页数
	Records int                         // 数据总数
	Size    int                         // 单页数量
}

type UserPointsHistoryListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
