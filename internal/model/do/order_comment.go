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

// OrderComment is the golang structure of table trade_order_comment for DAO operations like Where/Data.
type OrderComment struct {
	g.Meta                     `orm:"table:trade_order_comment, do:true"`
	OrderId                    interface{} // 订单编号
	StoreId                    interface{} // 卖家店铺编号-冗余
	StoreName                  interface{} // 店铺名称
	UserId                     interface{} // 买家编号
	UserName                   interface{} // 买家姓名
	CommentPoints              interface{} // 获得积分-冗余，独立表记录
	CommentScores              interface{} // 评价星级1-5积分
	CommentContent             interface{} // 评价内容
	CommentImage               interface{} // 评论上传的图片：|分割多张图片
	CommentHelpful             interface{} // 有帮助
	CommentNohelpful           interface{} // 无帮助
	CommentTime                *gtime.Time // 评价时间
	CommentIsAnonymous         interface{} // 匿名评价
	CommentEnable              interface{} // 评价信息的状态(BOOL): 1-正常显示; 0-禁止显示
	CommentStoreDescCredit     interface{} // 描述相符评分 - order_buyer_evaluation_status , 评价状态改变后不需要再次评论，根据订单走
	CommentStoreServiceCredit  interface{} // 服务态度评分 - order_buyer_evaluation_status
	CommentStoreDeliveryCredit interface{} // 发货速度评分 - order_buyer_evaluation_status
	SubsiteId                  interface{} // 所属分站:0-总站
}

type OrderCommentListInput struct {
	ml.BaseList
	Where OrderComment // 查询条件
}

type OrderCommentListOutput struct {
	Items   []*entity.OrderComment // 列表
	Page    int                    // 分页号码
	Total   int                    // 总页数
	Records int                    // 数据总数
	Size    int                    // 单页数量
}

type OrderCommentListKeyOutput struct {
	Items   []interface{} // 列表
	Page    int           // 分页号码
	Total   int           // 总页数
	Records int           // 数据总数
	Size    int           // 单页数量
}
