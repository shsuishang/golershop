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

// OrderComment is the golang structure for table order_comment.
type OrderComment struct {
	OrderId                    string      `json:"order_id"                      ` // 订单编号
	StoreId                    uint        `json:"store_id"                      ` // 卖家店铺编号-冗余
	StoreName                  string      `json:"store_name"                    ` // 店铺名称
	UserId                     uint        `json:"user_id"                       ` // 买家编号
	UserName                   string      `json:"user_name"                     ` // 买家姓名
	CommentPoints              float64     `json:"comment_points"                ` // 获得积分-冗余，独立表记录
	CommentScores              uint        `json:"comment_scores"                ` // 评价星级1-5积分
	CommentContent             string      `json:"comment_content"               ` // 评价内容
	CommentImage               string      `json:"comment_image"                 ` // 评论上传的图片：|分割多张图片
	CommentHelpful             uint        `json:"comment_helpful"               ` // 有帮助
	CommentNohelpful           uint        `json:"comment_nohelpful"             ` // 无帮助
	CommentTime                *gtime.Time `json:"comment_time"                  ` // 评价时间
	CommentIsAnonymous         uint        `json:"comment_is_anonymous"          ` // 匿名评价
	CommentEnable              uint        `json:"comment_enable"                ` // 评价信息的状态(BOOL): 1-正常显示; 0-禁止显示
	CommentStoreDescCredit     uint        `json:"comment_store_desc_credit"     ` // 描述相符评分 - order_buyer_evaluation_status , 评价状态改变后不需要再次评论，根据订单走
	CommentStoreServiceCredit  uint        `json:"comment_store_service_credit"  ` // 服务态度评分 - order_buyer_evaluation_status
	CommentStoreDeliveryCredit uint        `json:"comment_store_delivery_credit" ` // 发货速度评分 - order_buyer_evaluation_status
	SubsiteId                  uint        `json:"subsite_id"                    ` // 所属分站:0-总站
}
